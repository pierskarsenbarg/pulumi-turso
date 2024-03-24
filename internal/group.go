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

/*
"archived": true,
    "locations": [
      "lhr",
      "ams",
      "bos"
    ],
    "name": "default",
    "primary": "lhr",
    "uuid": "0a28102d-6906-11ee-8553-eaa7715aeaf2",
    "version": "v0.23.7"
*/

func (c *Client) CreateGroup(ctx context.Context, name string, location string) {

}
