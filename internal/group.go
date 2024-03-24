package turso

import "context"

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

func (c *Client) CreateGroup(ctx context.Context, name string, location string) {

}
