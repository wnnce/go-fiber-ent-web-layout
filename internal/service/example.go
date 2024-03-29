package service

import (
	"go-fiber-ent-web-layout/ent"
	"go-fiber-ent-web-layout/internal/common"
	"go-fiber-ent-web-layout/internal/factory"
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
		logger: factory.GetLogger("example-service"),
		epRepo: epRepo,
	}
}

func (es *ExampleService) QueryExampleInfo(id int) (*ent.Example, error) {
	example, err := es.epRepo.QueryExampleById(id)
	if err != nil {
		return nil, common.FiberServerError("查询失败")
	}
	return example, nil
}

func (es *ExampleService) ListExample() ([]*ent.Example, error) {
	example, err := es.epRepo.ListAllExample()
	if err != nil {
		return nil, common.FiberServerError("查询失败")
	}
	return example, nil
}

func (es *ExampleService) SaveExample(example *ent.Example) error {
	err := es.epRepo.CreateExample(example)
	if err != nil {
		return common.FiberServerError("保存失败")
	}
	return nil
}

func (es *ExampleService) UpdateExample(example *ent.Example) error {
	currenTime := time.Now()
	example.UpdateTime = &currenTime
	err := es.epRepo.UpdateExampleById(example)
	if err != nil {
		return common.FiberServerError("更新失败")
	}
	return nil
}
