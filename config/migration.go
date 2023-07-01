package config

import "startupfundinggolang/models"

type Model struct {
	Model interface{}
}

func Migration() []Model {
	return []Model{
		{Model: models.User{}},
	}
}
