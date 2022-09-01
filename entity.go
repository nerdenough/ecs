package ecs

import "fmt"

type Entity struct {
	id         string
	components map[string]Component
}

func NewEntity() *Entity {
	return &Entity{
		id:         "",
		components: map[string]Component{},
	}
}

func (e *Entity) ID() string {
	return e.id
}

func (e *Entity) Component(componentName string) *Component {
	if cmp, ok := e.components[componentName]; ok {
		return &cmp
	}
	return nil
}

func (e *Entity) UpdateComponent(componentName string, data interface{}) {
	cmp := e.components[componentName]
	cmp.Data = data
	e.components[componentName] = cmp
}

func (e *Entity) WithID(id string) *Entity {
	e.id = id
	return e
}

func (e *Entity) WithComponent(component Component) *Entity {
	e.components[component.Name] = component
	return e
}

func (e *Entity) WithComponents(components ...Component) *Entity {
	for _, component := range components {
		e.components[component.Name] = component
	}
	return e
}

func (e *Entity) RemoveComponent(componentName string) {
	delete(e.components, componentName)
}

func (e *Entity) Debug() string {
	cmpStr := ""
	for _, cmp := range e.components {
		cmpStr += fmt.Sprintf("%s,", cmp.Name)
	}

	return fmt.Sprintf("Entity:%s, Components:%s", e.id, cmpStr)
}
