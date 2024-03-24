package turso

import (
	"context"
	"fmt"
	"net/http"
)

type CreateDatabaseRequest struct {
	Name  string `json:"name"`
	Group string `json:"group"`
}

type CreateDatabaseResponse struct {
	Database CreateDatabaseResponseDatabase `json:"database"`
}

type CreateDatabaseResponseDatabase struct {
	Id       string `json:"DbId"`
	HostName string `json:"Hostname"`
	Name     string `json:"Name"`
}

type ListDatabaseResponse struct {
	Database []DatabaseResponse `json:"databases"`
}

type DatabaseResponse struct {
	DbId          string   `json:"DbId"`
	Hostname      string   `json:"Hostname"`
	Name          string   `json:"Name"`
	AllowAttach   bool     `json:"allow_attach"`
	BlockReads    bool     `json:"block_reads"`
	BlockWrites   bool     `json:"block_writes"`
	Group         string   `json:"group"`
	IsSchema      bool     `json:"is_schema"`
	PrimaryRegion string   `json:"primaryRegion"`
	Regions       []string `json:"regions"`
	Schema        string   `json:"schema"`
	Sleeping      bool     `json:"sleeping"`
	Type          string   `json:"type"`
	Version       string   `json:"version"`
}

type GetDatabaseResponse struct {
	Database DatabaseResponse `json:"database"`
}

type DeleteDatabaseResponse struct {
	Database string `json:"database"`
}

func (c *Client) CreateDatabase(ctx context.Context, name string, group string, organisationName string) (*CreateDatabaseResponse, error) {
	requestPath := fmt.Sprintf("/v1/organizations/%s/databases", organisationName)
	req := CreateDatabaseRequest{
		Name:  name,
		Group: group,
	}
	var res CreateDatabaseResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListDatabases(ctx context.Context, organisationName string) (*ListDatabaseResponse, error) {
	requestPath := fmt.Sprintf("/v1/organizations/%s/databases", organisationName)
	var res ListDatabaseResponse
	_, err := c.do(ctx, http.MethodGet, requestPath, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetDatabase(ctx context.Context, organisationName string, databaseName string) (*GetDatabaseResponse, error) {
	requestPath := fmt.Sprintf("/v1/organizations/%s/databases/%s", organisationName, databaseName)
	var res GetDatabaseResponse
	_, err := c.do(ctx, http.MethodGet, requestPath, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteDatabase(ctx context.Context, organisationName string, databaseName string) (*DeleteDatabaseResponse, error) {
	requestPath := fmt.Sprintf("/v1/organizations/%s/databases/%s", organisationName, databaseName)
	var res DeleteDatabaseResponse
	_, err := c.do(ctx, http.MethodDelete, requestPath, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
