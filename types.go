package main

import (
	"github.com/drone/drone-go/drone"
)

// Params are the parameters that the HipChat plugin can parse.
type Params struct {
	Notify            bool            `json:"notify"`
	URL               string          `json:"url"`
	From              string          `json:"from"`
	Room              drone.StringInt `json:"room_id_or_name"`
	Token             string          `json:"auth_token"`
	Template          string          `json:"template"`
	UseCard           bool            `json:"use_card"`
	TitleTemplate     string          `json:"card_title_template"`
	DescTemplate      string          `json:"card_desc_template"`
	ActivityTemplate  string          `json:"card_activity_template"`
	Icon              string          `json:"card_icon"`
}
