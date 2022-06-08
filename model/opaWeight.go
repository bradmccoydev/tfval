package model

type Weight struct {
	Service string `json:"service"`
	Create  int    `json:"create"`
	Delete  int    `json:"delete"`
	Modify  int    `json:"modify"`
}
