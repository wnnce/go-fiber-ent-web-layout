package data

import (
	"context"
	"go-fiber-ent-web-layout/ent"
	"go-fiber-ent-web-layout/internal/factory"
	"go-fiber-ent-web-layout/internal/usercase"
	"log/slog"
	"strconv"
	"time"
)

type ExampleRepo struct {
	data   *Data
	logger *slog.Logger
}

var (
	redisPrefix     = "example:"
	cacheExpireTime = 24 * time.Hour
)

func NewExampleRepo(data *Data) usercase.IExampleRepo {
	return &ExampleRepo{
		data:   data,
		logger: factory.GetLogger("example-repo"),
	}
}

func (ep *ExampleRepo) QueryExampleById(id int) (*ent.Example, error) {
	cacheKey := redisPrefix + strconv.Itoa(id)
	example := &ent.Example{}
	err := ep.data.Rc.Get(context.Background(), cacheKey, example)
	if err == nil {
		return example, err
	} else {
		ep.logger.Warn("redis缓存获取失败，key:%s", cacheKey)
	}
	example, err = ep.data.Ec.Example.Get(context.Background(), id)
	if err != nil {
		ep.logger.Warn("查询example失败，错误信息：%v", err)
		return nil, err
	}
	if example != nil {
		err = ep.data.Rc.Set(context.Background(), cacheKey, example, cacheExpireTime)
		if err != nil {
			ep.logger.Error("查询结果缓存到Redis失败，错误信息：%v", err)
		}
	}
	return example, nil
}

func (ep *ExampleRepo) ListAllExample() ([]*ent.Example, error) {
	redisKey := redisPrefix + "list"
	var examples []*ent.Example
	err := ep.data.Rc.Get(context.Background(), redisKey, &examples)
	if err == nil {
		return examples, nil
	}
	examples, err = ep.data.Ec.Example.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	if len(examples) > 0 {
		err = ep.data.Rc.Set(context.Background(), redisKey, examples, cacheExpireTime)
		if err != nil {
			ep.logger.Warn("redis缓存example切片失败")
		}
	}
	return examples, nil
}

func (ep *ExampleRepo) CreateExample(example *ent.Example) error {
	_, err := ep.data.Ec.Example.Create().SetName(example.Name).SetSummary(example.Summary).SetPrice(example.Price).Save(context.Background())
	return err
}

func (ep *ExampleRepo) UpdateExampleById(example *ent.Example) error {
	updateBuilder := ep.data.Ec.Example.Update()
	if len(example.Name) > 0 {
		updateBuilder.SetName(example.Name)
	}
	if len(example.Summary) > 0 {
		updateBuilder.SetSummary(example.Summary)
	}
	if example.Price > 0 {
		updateBuilder.SetPrice(example.Price)
	}
	_, err := updateBuilder.Save(context.Background())
	if err != nil {
		return err
	}
	err = ep.data.Rc.Remove(context.Background(), redisPrefix+strconv.Itoa(example.ID))
	if err != nil {
		ep.logger.Error("删除redis缓存失败，错误消息：%v", err)
	}
	return nil
}
