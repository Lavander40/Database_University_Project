package model

type Student struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	GroupId string `json:"group_id"`
}

type Equipment struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Audience struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
