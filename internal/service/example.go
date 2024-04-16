package service

import (
	"go-fiber-ent-web-layout/ent"
	"go-fiber-ent-web-layout/internal/common/res"
	"go-fiber-ent-web-layout/internal/usercase"
	"log/slog"
	"time"
)

type ExampleService struct {
	logger *slog.Logger
	epRepo usercase.IExampleRepo
}

func NewExampleService(epRepo usercase.IExampleRepo) usercase.IExampleService {
	return &ExampleService{
		logger: slog.Default().With("trace-name", "example-service"),
		epRepo: epRepo,
	}
}

func (es *ExampleService) QueryExampleInfo(id int) (*ent.Example, error) {
	example, err := es.epRepo.QueryExampleById(id)
	if err != nil {
		return nil, res.FiberServerError("查询失败")
	}
	return example, nil
}

func (es *ExampleService) ListExample() ([]*ent.Example, error) {
	example, err := es.epRepo.ListAllExample()
	if err != nil {
		return nil, res.FiberServerError("查询失败")
	}
	return example, nil
}

func (es *ExampleService) SaveExample(example *ent.Example) error {
	err := es.epRepo.CreateExample(example)
	if err != nil {
		return res.FiberServerError("保存失败")
	}
	return nil
}

func (es *ExampleService) UpdateExample(example *ent.Example) error {
	currenTime := time.Now()
	example.UpdateTime = &currenTime
	err := es.epRepo.UpdateExampleById(example)
	if err != nil {
		return res.FiberServerError("更新失败")
	}
	return nil
}
