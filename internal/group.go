package turso

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type CreateGroupRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type DeleteGroupRequest struct {
	Name         string `json:"name"`
	Organization string `json:"organization"`
}

type DeleteGroupResponse struct {
	Group Group `json:"group"`
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

type GroupLocationRequest struct {
	Organization string `json:"organization"`
	GroupName    string `json:"groupName"`
	Location     string `json:"location"`
}

type GroupLocationResponse struct {
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

func (c *Client) CreateGroup(ctx context.Context, req CreateGroupRequest, organization string) (*CreateGroupResponse, error) {
	requestPath := fmt.Sprintf("/v1/organizations/%s/groups", organization)
	var res CreateGroupResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, req, &res)
	if err != nil {
		log.Printf("error: %s", err)
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteGroup(ctx context.Context, req DeleteGroupRequest) error {
	requestPath := fmt.Sprintf("/v1/organizations/%s/groups/%s", req.Organization, req.Name)
	var res DeleteGroupResponse
	_, err := c.do(ctx, http.MethodDelete, requestPath, nil, &res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddLocationToGroup(ctx context.Context, req GroupLocationRequest) (*GroupLocationResponse, error) {
	requestPath := fmt.Sprintf("/v1/organizations/%s/groups/%s/locations/%s", req.Organization, req.GroupName, req.Location)
	var res GroupLocationResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, req, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) RemoveLocationFromGroup(ctx context.Context, req GroupLocationRequest) (*GroupLocationResponse, error) {
	requestpath := fmt.Sprintf("/v1/organizations/%s/groups/%s/locations/%s", req.Organization, req.GroupName, req.Location)
	var res GroupLocationResponse
	_, err := c.do(ctx, http.MethodDelete, requestpath, req, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
