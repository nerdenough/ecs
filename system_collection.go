package ecs

type SystemCollection struct {
	systems []System
}

func NewSystemCollection() *SystemCollection {
	return &SystemCollection{
		systems: []System{},
	}
}

func (sc *SystemCollection) Add(system System) *SystemCollection {
	sc.systems = append(sc.systems, system)
	return sc
}

func (sc *SystemCollection) Range(fn func(System)) {
	for _, system := range sc.systems {
		fn(system)
	}
}
