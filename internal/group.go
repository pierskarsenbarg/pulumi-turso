package turso

import (
	"context"
	"fmt"
	"net/http"
)

type CreateGroupRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type CreateGroupResponse struct {
	Group Group `json:"group"`
}

type ListGroupResponse struct {
	Groups []Group `json:"groups"`
}

type GetGroupResponse struct {
	Group Group `json:"group"`
}

type Group struct {
	Archived        bool     `json:"archived"`
	Locations       []string `json:"locations"`
	Name            string   `json:"name"`
	PrimaryLocation string   `json:"primary"`
	Uuid            string   `json:"uuid"`
	Version         string   `json:"version"`
}

func (c *Client) CreateGroup(ctx context.Context, organisationName string, groupName string, location string) (*CreateGroupResponse, error) {
	requestPath := fmt.Sprintf("/v1/organizations/%s/groups", organisationName)
	req := CreateGroupRequest{
		Name:     groupName,
		Location: location,
	}
	var res CreateGroupResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, req, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
