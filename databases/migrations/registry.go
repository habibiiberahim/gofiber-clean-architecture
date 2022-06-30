package databases

import "github.com/habibiiberahim/gofiber-clean-architecture/entities"

type Entity struct {
	Entity interface{}
}

func RegisterEntities() []Entity {
	return []Entity{
		{Entity: entities.User{}},
	}
}
