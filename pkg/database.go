package pkg

import (
	"context"

	turso "github.com/pierskarsenbarg/pulumi-turso/internal"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Database struct{}

func (d *Database) Annotate(a infer.Annotator) {
	a.Describe(&d, "Manage a Turso database")
}

type DatabaseArgs struct {
	Name             string `pulumi:"name,optional"`
	GroupName        string `pulumi:"groupName"`
	OrganizationName string `pulumi:"organizationName,optional"`
}

func (da *DatabaseArgs) Annotate(a infer.Annotator) {
	a.Describe(&da.Name, "The name of the new database. Must contain only lowercase letters, numbers, dashes. No longer than 32 characters.")
	a.Describe(&da.GroupName, "The name of the group where the database should be created. **The group must already exist.**")
	a.Describe(&da.OrganizationName, "The name of the organization or user.")
	a.SetDefault(&da.OrganizationName, "", "TURSO_ORGANISATIONNAME")
}

type DatabaseState struct {
	DbId             string `pulumi:"dbId"`
	GroupName        string `pulumi:"groupName"`
	HostName         string `pulumi:"hostName" provider:"secret"`
	Name             string `pulumi:"name"`
	OrganizationName string `pulumi:"organizationName"`
}

func (ds *DatabaseState) Annotate(a infer.Annotator) {
	a.Describe(&ds.DbId, "The database universal unique identifier (UUID).")
	a.Describe(&ds.GroupName, "The name of the group where the database was created.")
	a.Describe(&ds.HostName, "The DNS hostname used for client libSQL and HTTP connections.")
	a.Describe(&ds.Name, "The database name, unique across your organization.")
	a.Describe(&ds.OrganizationName, "The name of the organization or user that created the database.")
}

type GetDatabase struct{}

type GetDatabaseArgs struct {
	DatabaseName     string `pulumi:"databaseName"`
	OrganizationName string `pulumi:"organizationName"`
}

func (d *Database) Create(ctx context.Context, name string, input DatabaseArgs, preview bool) (id string, output DatabaseState, err error) {

	if preview {
		return "", DatabaseState{}, nil
	}

	config := infer.GetConfig[Config](ctx)
	orgName := config.OrganizationName
	if len(input.OrganizationName) > 0 {
		orgName = input.OrganizationName
	}

	databaseName, _ := buildName(name)
	if len(input.Name) > 0 {
		databaseName = input.Name
	}
	res, err := d.createDatabase(databaseName, orgName, input.GroupName, config, ctx)

	if err != nil {
		return "", DatabaseState{}, err
	}

	return name, DatabaseState{
		DbId:      res.Database.Id,
		Name:      res.Database.Name,
		GroupName: input.GroupName,
		HostName:  res.Database.HostName,
	}, nil
}

func (*Database) createDatabase(name string, organization string, groupName string, config Config, ctx context.Context) (*turso.CreateDatabaseResponse, error) {
	database, err := config.Client.CreateDatabase(ctx, name, groupName, organization)
	if err != nil {
		return nil, err
	}
	return database, nil
}

func (d *Database) Delete(ctx context.Context, id string, props DatabaseState) error {
	config := infer.GetConfig[Config](ctx)
	orgName := config.OrganizationName
	if len(props.OrganizationName) > 0 {
		orgName = props.OrganizationName
	}
	err := d.deleteDatabase(orgName, props.Name, config, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (*Database) deleteDatabase(organizationName string, databaseName string, config Config, ctx context.Context) error {
	err := config.Client.DeleteDatabase(ctx, organizationName, databaseName)
	if err != nil {
		return err
	}

	return nil
}

func (*Database) Diff(ctx context.Context, id string, olds DatabaseState, news DatabaseArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	if len(olds.Name) > 0 && len(news.Name) == 0 {
		// name removed
		diff["name"] = p.PropertyDiff{Kind: p.DeleteReplace}
	} else if len(news.Name) > 0 && olds.Name != news.Name {
		// name updated
		diff["name"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}

	if olds.GroupName != news.GroupName {
		diff["groupName"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}

	if olds.OrganizationName != news.OrganizationName {
		diff["organizationName"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Database) Read(ctx context.Context, id string, inputs DatabaseArgs, state DatabaseState) (
	string, DatabaseArgs, DatabaseState, error) {
	config := infer.GetConfig[Config](ctx)

	database, err, res := config.Client.GetDatabase(ctx, state.OrganizationName, state.Name)
	if err != nil {
		if res.StatusCode == 404 {
			// database doesn't exist
			return "", DatabaseArgs{}, DatabaseState{}, nil
		}
		return "", DatabaseArgs{}, DatabaseState{}, err
	}
	return database.Database.DbId, DatabaseArgs{
			Name:             database.Database.Name,
			GroupName:        database.Database.Group,
			OrganizationName: inputs.OrganizationName,
		}, DatabaseState{
			DbId:             database.Database.DbId,
			GroupName:        database.Database.Group,
			HostName:         database.Database.Hostname,
			Name:             database.Database.Name,
			OrganizationName: inputs.OrganizationName,
		}, nil
}

func (GetDatabase) Call(ctx context.Context, args GetDatabaseArgs) (DatabaseState, error) {
	config := infer.GetConfig[Config](ctx)

	database, err, res := config.Client.GetDatabase(ctx, args.OrganizationName, args.DatabaseName)
	if err != nil {
		if res.StatusCode == 404 {
			return DatabaseState{}, nil
		}
		return DatabaseState{}, err
	}
	return DatabaseState{
		DbId:             database.Database.DbId,
		GroupName:        database.Database.Group,
		HostName:         database.Database.Hostname,
		Name:             args.DatabaseName,
		OrganizationName: args.OrganizationName,
	}, nil
}
