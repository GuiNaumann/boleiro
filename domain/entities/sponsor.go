package entities

import "boleiro/util"

type Sponsor struct {
	Id         int64         `json:"id,omitempty"`
	Name       string        `json:"name,omitempty"`
	StatusCode int8          `json:"statusCode"`
	ModifiedAt util.DateTime `json:"modifiedAt,omitempty"`
	CreatedAt  util.DateTime `json:"createdAt,omitempty"`
}
