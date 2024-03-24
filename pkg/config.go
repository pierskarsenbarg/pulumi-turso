package pkg

import (
	"net/http"
	"time"

	turso "github.com/pierskarsenbarg/pulumi-turso/internal"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	ApiToken         string `pulumi:"apiToken,optional" provider:"secret"`
	OrganizationName string `pulumi:"organizationName,optional"`
	Client           turso.Client
}

var _ = (infer.Annotated)((*Config)(nil))

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.ApiToken, "Your Turso API token")
	a.Describe(&c.OrganizationName, "Organisation name")
	a.SetDefault(&c.ApiToken, "", "TURSO_APITOKEN")
	a.SetDefault(&c.OrganizationName, "", "TURSO_ORGANISATIONNAME")
}

var _ = (infer.CustomConfigure)((*Config)(nil))

func (c *Config) Configure(ctx p.Context) error {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	client, err := turso.NewClient(&httpClient, c.ApiToken, "")
	if err != nil {
		return err
	}

	c.Client = *client

	return nil
}
