package main

import (
	"fmt"
	"os"

	"github.com/pierskarsenbarg/pulumi-turso/pkg"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var Version string

func main() {
	err := p.RunProvider("turso", "0.1.0", provider())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func provider() p.Provider {
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			DisplayName: "turso",
			Description: "Turso provider",
			LanguageMap: map[string]any{
				"go": gen.GoPackageInfo{
					Generics:       gen.GenericsSettingGenericsOnly,
					ImportBasePath: "github.com/pierskarsenbarg/pulumi-turso/sdk/go/turso",
				},
				"nodejs": nodejsgen.NodePackageInfo{
					PackageName: "@pierskarsenbarg/pulumi-turso",
					Dependencies: map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
					DevDependencies: map[string]string{
						"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
						"@types/mime": "^2.0.0",
					},
				},
				"csharp": dotnetgen.CSharpPackageInfo{
					RootNamespace: "PiersKarsenbarg",
					PackageReferences: map[string]string{
						"Pulumi": "3.*",
					},
				},
				"python": pythongen.PackageInfo{
					Requires: map[string]string{
						"pulumi": ">=3.0.0,<4.0.0",
					},
					PackageName: "pierskarsenbarg_pulumi_turso",
				},
			},
			PluginDownloadURL: "github://api.github.com/pierskarsenbarg/pulumi-turso",
			Publisher:         "Piers Karsenbarg",
		},
		Resources: []infer.InferredResource{
			infer.Resource[*pkg.Database, pkg.DatabaseArgs, pkg.DatabaseState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"pkg": "index",
		},
		Functions: []infer.InferredFunction{
			// infer.Function[*pkg.GetOrganization, pkg.GetOrganizationArgs, pkg.OrganizationState](),
			// infer.Function[*pkg.GetWorkspace, pkg.GetWorkspaceArgs, pkg.WorkspaceState](),
		},
		Config: infer.Config[*pkg.Config](),
	})
}
