package pkg

import (
	"context"
	"fmt"

	turso "github.com/pierskarsenbarg/pulumi-turso/internal"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
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

func (g *Group) Create(ctx p.Context, name string, input GroupArgs, preview bool) (id string, output GroupState, err error) {

	if preview {
		return "", GroupState{}, nil
	}
	groupName, _ := buildName(name)
	if len(input.Name) > 0 {
		groupName = input.Name
	}
	config := infer.GetConfig[Config](ctx)
	res, err := g.createGroup(groupName, input.PrimaryLocation, input.Organization, config)

	if err != nil {
		return "", GroupState{}, err
	}

	for _, location := range input.Locations {
		err = g.addLocationToGroup(groupName, input.Organization, location, config)
		if err != nil {
			ctx.Log(diag.Warning, fmt.Sprintf("couldn't add location to group: %s", err))
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

func (*Group) createGroup(name string, location string, organization string, config Config) (*turso.CreateGroupResponse, error) {
	ctx := context.Background()
	group, err := config.Client.CreateGroup(ctx, turso.CreateGroupRequest{
		Name:     name,
		Location: location,
	}, organization)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *Group) Delete(ctx p.Context, id string, props GroupState) error {
	config := infer.GetConfig[Config](ctx)
	err := g.deleteGroup(props.Name, props.Organization, config)
	if err != nil {
		return err
	}
	return nil
}

func (*Group) deleteGroup(name string, organization string, config Config) error {
	ctx := context.Background()
	err := config.Client.DeleteGroup(ctx, turso.DeleteGroupRequest{
		Name:         name,
		Organization: organization,
	})
	if err != nil {
		return err
	}
	return nil
}

func (*Group) addLocationToGroup(name string, organization string, location string, config Config) error {
	ctx := context.Background()
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

func (g *Group) Diff(ctx p.Context, id string, olds GroupState, news GroupArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	if len(olds.Name) > 0 && len(news.Name) == 0 {
		// name removed
		ctx.Log(diag.Info, fmt.Sprint("name removed"))
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	} else if len(news.Name) > 0 && olds.Name != news.Name {
		// name updated
		ctx.Log(diag.Info, fmt.Sprint("name updated"))
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}

	if sliceCompare(olds.Locations, news.Locations) {
		ctx.Log(diag.Info, fmt.Sprint("locations are not the same"))
		diff["locations"] = p.PropertyDiff{Kind: p.Update}
	}

	if olds.PrimaryLocation != news.PrimaryLocation {
		ctx.Log(diag.Info, fmt.Sprint("primary location changed"))
		diff["primaryLocation"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (o *Group) Update(ctx p.Context, id string, olds GroupState, news GroupArgs, preview bool) (GroupState, error) {
	return GroupState{}, nil
}

func (*Group) removeLocationFromGroup(name string, organization string, location string, config Config) error {
	ctx := context.Background()
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