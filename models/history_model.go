package models

type History struct {
	Id     int      `json:"id"`
	UserId int      `json:"userId"`
	Desc   string   `json:"desc"`
	Viewer []Viewer `json:"viewer"`
}
