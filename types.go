package main

import (
	"github.com/drone/drone-go/drone"
)

type Params struct {
	Notify   bool            `json:"notify"`
	From     string          `json:"from"`
	Room     drone.StringInt `json:"room_id_or_name"`
	Token    string          `json:"auth_token"`
	Template string          `json:"template"`
}
