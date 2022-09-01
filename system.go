package ecs

type System interface {
	Update(dt float64, entities []*Entity)
	Debug() string
}
