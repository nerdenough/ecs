package ecs

type Component struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
