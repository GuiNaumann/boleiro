package entities

import "boleiro/util"

type Relation struct {
	IdPlayers  int64         `json:"idPlayers,omitempty"`
	IdSponsor  string        `json:"idSponsor,omitempty"`
	StatusCode int8          `json:"statusCode"`
	ModifiedAt util.DateTime `json:"modifiedAt,omitempty"`
	CreatedAt  util.DateTime `json:"createdAt,omitempty"`
}
