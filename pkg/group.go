package pkg

import (
	"context"
	"fmt"

	turso "github.com/pierskarsenbarg/pulumi-turso/internal"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Group struct{}

type GroupArgs struct {
	Name            string   `pulumi:"name,optional"`
	PrimaryLocation string   `pulumi:"primaryLocation"`
	Organization    string   `pulumi:"organization"`
	Locations       []string `pulumi:"locations,optional"`
}

func (ga *GroupArgs) Annotate(a infer.Annotator) {
	a.Describe(&ga.Name, "The name of the new group.")
	a.Describe(&ga.PrimaryLocation, "The primary location key for the new group.")
	a.Describe(&ga.Organization, "The name of the organization or user.")
	a.Describe(&ga.Locations, "An array of location keys the group is located.")
}

type GroupState struct {
	Id              string   `pulumi:"groupId"`
	PrimaryLocation string   `pulumi:"primaryLocation"`
	DbVersion       string   `pulumi:"dbVersion"`
	Name            string   `pulumi:"name"`
	Locations       []string `pulumi:"locations,optional"`
	Organization    string   `pulumi:"organization"`
}

func (gs *GroupState) Annotate(a infer.Annotator) {
	a.Describe(&gs.Id, "The group universal unique identifier (UUID).")
	a.Describe(&gs.PrimaryLocation, "The primary location key.")
	a.Describe(&gs.DbVersion, "The current libSQL server version the databases in that group are running.")
	a.Describe(&gs.Name, "The group name, unique across your organization.")
	a.Describe(&gs.Locations, "An array of location keys the group is located.")
	a.Describe(&gs.Organization, "The name of the organization or user.")
}

type GetGroup struct{}

type GetGroupArgs struct {
	GroupName        string `pulumi:"groupName"`
	OrganizationName string `pulumi:"organizationName"`
}

func (g *Group) Create(ctx context.Context, name string, input GroupArgs, preview bool) (id string, output GroupState, err error) {

	if preview {
		return "", GroupState{}, nil
	}
	groupName, _ := buildName(name)
	if len(input.Name) > 0 {
		groupName = input.Name
	}
	config := infer.GetConfig[Config](ctx)
	res, err := g.createGroup(ctx, groupName, input.PrimaryLocation, input.Organization, config)

	if err != nil {
		return "", GroupState{}, err
	}

	for _, location := range input.Locations {
		err = g.addLocationToGroup(ctx, groupName, input.Organization, location, config)
		if err != nil {
			p.GetLogger(ctx).Warning(fmt.Sprintf("couldn't add location to group: %s", err))
		}
	}

	return name, GroupState{
		Id:              res.Group.Uuid,
		PrimaryLocation: res.Group.PrimaryLocation,
		DbVersion:       res.Group.Version,
		Name:            res.Group.Name,
		Locations:       input.Locations,
		Organization:    input.Organization,
	}, nil
}

func (*Group) createGroup(ctx context.Context, name string, location string, organization string, config Config) (*turso.CreateGroupResponse, error) {
	group, err := config.Client.CreateGroup(ctx, turso.CreateGroupRequest{
		Name:     name,
		Location: location,
	}, organization)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *Group) Delete(ctx context.Context, id string, props GroupState) error {
	config := infer.GetConfig[Config](ctx)
	err := g.deleteGroup(ctx, props.Name, props.Organization, config)
	if err != nil {
		return err
	}
	return nil
}

func (*Group) deleteGroup(ctx context.Context, name string, organization string, config Config) error {
	err, res := config.Client.DeleteGroup(ctx, turso.DeleteGroupRequest{
		Name:         name,
		Organization: organization,
	})
	if err != nil {
		if res.StatusCode == 404 {
			return nil
		}
		return err
	}
	return nil
}

func (g *Group) Diff(ctx context.Context, id string, olds GroupState, news GroupArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	if len(olds.Name) > 0 && len(news.Name) == 0 {
		// name removed
		diff["name"] = p.PropertyDiff{Kind: p.Delete}
	} else if len(news.Name) > 0 && olds.Name != news.Name {
		// name updated
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}

	if !sliceCompare(olds.Locations, news.Locations) {
		diff["locations"] = p.PropertyDiff{Kind: p.Update}
	}

	if olds.PrimaryLocation != news.PrimaryLocation {
		diff["primaryLocation"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (g *Group) Update(ctx context.Context, id string, olds GroupState, news GroupArgs, preview bool) (GroupState, error) {
	updatedGroupState := GroupState{
		Id:              olds.Id,
		PrimaryLocation: olds.PrimaryLocation,
		DbVersion:       olds.DbVersion,
		Name:            olds.Name,
		Locations:       olds.Locations,
		Organization:    olds.Organization,
	}
	config := infer.GetConfig[Config](ctx)
	var err error
	if !preview {
		// Check if name of group has changed
		if news.Name != olds.Name && len(news.Name) == 0 {
			// this happens if they had the name arg set but removed it so we fall back to the resource name arg
			updatedGroupState.Name, err = buildName(id)
			if err != nil {
				return GroupState{}, fmt.Errorf("there was an issue creating the new resource name: %s", err)
			}
		} else if news.Name != olds.Name {
			// this happens if they just change the name
			updatedGroupState.Name = news.Name
		}

		// check that group exists
		_, err, res := config.Client.GetGroup(ctx, turso.GetGroupRequest{
			OrganizationName: updatedGroupState.Organization,
			GroupName:        updatedGroupState.Name,
		})

		if err != nil {
			if res.StatusCode == 404 {
				return GroupState{}, nil
			}
			return GroupState{}, err
		}

		// check if the old and new locations match and if they don't then delete them all and add new ones
		if !sliceCompare(updatedGroupState.Locations, news.Locations) {
			updatedGroupState.Locations = make([]string, 0)
			// delete all locations from group
			for _, oldLocation := range olds.Locations {
				err = g.removeLocationFromGroup(ctx, updatedGroupState.Name, updatedGroupState.Organization, oldLocation, config)
				if err != nil {
					return GroupState{}, fmt.Errorf("there as an issue removing location %s from group %s: %s", oldLocation, updatedGroupState.Name, err)
				}
			}

			// add locations back
			for _, newLocation := range news.Locations {
				err = g.addLocationToGroup(ctx, updatedGroupState.Name, updatedGroupState.Organization, newLocation, config)
				if err != nil {
					return GroupState{}, fmt.Errorf("there was an issue adding location %s to group %s: %s", newLocation, updatedGroupState.Name, err)
				}
			}
			updatedGroupState.Locations = news.Locations
		}
	}

	return updatedGroupState, nil
}

func (g *Group) Read(ctx context.Context, id string, inputs GroupArgs, state GroupState) (
	string, GroupArgs, GroupState, error) {
	config := infer.GetConfig[Config](ctx)
	group, err, res := config.Client.GetGroup(ctx, turso.GetGroupRequest{
		OrganizationName: inputs.Organization,
		GroupName:        inputs.Name,
	})
	if err != nil {
		if res.StatusCode == 404 {
			return "", GroupArgs{}, GroupState{}, nil
		}
		return "", GroupArgs{}, GroupState{}, err
	}

	return id, GroupArgs{
			Name:            inputs.Name,
			PrimaryLocation: inputs.PrimaryLocation,
			Organization:    inputs.Organization,
			Locations:       inputs.Locations,
		}, GroupState{
			Id:              group.Group.Uuid,
			PrimaryLocation: group.Group.PrimaryLocation,
			DbVersion:       group.Group.Version,
			Name:            group.Group.Name,
			Locations:       group.Group.Locations,
			Organization:    inputs.Organization,
		}, nil
}

func (g *GetGroup) Call(ctx context.Context, args GetGroupArgs) (GroupState, error) {
	config := infer.GetConfig[Config](ctx)
	group, err, res := config.Client.GetGroup(ctx, turso.GetGroupRequest{
		OrganizationName: args.OrganizationName,
		GroupName:        args.GroupName,
	})

	if err != nil {
		if res.StatusCode == 404 {
			return GroupState{}, nil
		}
		return GroupState{}, err
	}

	return GroupState{
		Id:              group.Group.Uuid,
		PrimaryLocation: group.Group.PrimaryLocation,
		DbVersion:       group.Group.Version,
		Name:            group.Group.Name,
		Locations:       group.Group.Locations,
		Organization:    args.OrganizationName,
	}, nil
}

func (*Group) addLocationToGroup(ctx context.Context, name string, organization string, location string, config Config) error {
	_, err := config.Client.AddLocationToGroup(ctx, turso.GroupLocationRequest{
		Organization: organization,
		GroupName:    name,
		Location:     location,
	})
	if err != nil {
		return err
	}
	return nil
}

func (*Group) removeLocationFromGroup(ctx context.Context, name string, organization string, location string, config Config) error {
	_, err := config.Client.RemoveLocationFromGroup(ctx, turso.GroupLocationRequest{
		Organization: organization,
		GroupName:    name,
		Location:     location,
	})
	if err != nil {
		return err
	}
	return nil
}

// func (g *Group) getGroup(ctx context.Context, name string, organization string, config Config) (*turso.GetGroupResponse, error) {
// 	req := turso.GetGroupRequest{
// 		OrganizationName: organization,
// 		GroupName:        name,
// 	}
// 	group, err, _ := config.Client.GetGroup(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return group, nil
// }
