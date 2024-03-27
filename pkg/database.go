package pkg

import (
	p "github.com/pulumi/pulumi-go-provider"
)

type Database struct{}

type DatabaseArgs struct{}

type DatabaseState struct{}

func (o *Database) Create(ctx p.Context, name string, input DatabaseArgs, preview bool) (id string, output DatabaseState, err error) {
	// config := infer.GetConfig[Config](ctx)

	if preview {
		return "", DatabaseState{}, nil
	}

	return "", DatabaseState{}, nil
}
