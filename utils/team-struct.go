package utils

type Team struct {
	Name  string   `json:"team" form:"team" query:"team"`
	Mates []string `json:"mates" form:"mates" query:"mates"`
}
