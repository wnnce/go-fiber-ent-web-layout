package usercase

import "go-fiber-ent-web-layout/ent"

type IExampleRepo interface {
	QueryExampleById(id int) (*ent.Example, error)
	ListAllExample() ([]*ent.Example, error)
	CreateExample(example *ent.Example) error
	UpdateExampleById(example *ent.Example) error
}

type IExampleService interface {
	QueryExampleInfo(id int) (*ent.Example, error)
	ListExample() ([]*ent.Example, error)
	SaveExample(example *ent.Example) error
	UpdateExample(example *ent.Example) error
}
