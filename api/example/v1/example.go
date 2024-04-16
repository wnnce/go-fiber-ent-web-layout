package example

import (
	"github.com/gofiber/fiber/v3"
	"go-fiber-ent-web-layout/ent"
	"go-fiber-ent-web-layout/internal/common"
	"go-fiber-ent-web-layout/internal/common/res"
	"go-fiber-ent-web-layout/internal/usercase"
	"log/slog"
)

type ExampleApi struct {
	logger  *slog.Logger
	service usercase.IExampleService
}

func NewExampleApi(epService usercase.IExampleService) *ExampleApi {
	return &ExampleApi{
		logger:  slog.Default().With("trace-name", "example-api"),
		service: epService,
	}
}

func (e *ExampleApi) QueryExample(ctx fiber.Ctx) error {
	id := fiber.Params[int](ctx, "id")
	example, err := e.service.QueryExampleInfo(id)
	if err != nil {
		return err
	}
	return ctx.JSON(res.OkByData(example))
}

func (e *ExampleApi) ListExample(ctx fiber.Ctx) error {
	exampleList, err := e.service.ListExample()
	if err != nil {
		return err
	}
	return ctx.JSON(res.OkByData(exampleList))
}

func (e *ExampleApi) SaveExample(ctx fiber.Ctx) error {
	example := &ent.Example{}
	if err := ctx.Bind().JSON(example); err != nil {
		return err
	}
	if errorMessage := common.StructFieldValidation(example); len(errorMessage) > 0 {
		return ctx.JSON(res.FiberRequestError(errorMessage))
	}
	if err := e.service.SaveExample(example); err != nil {
		return err
	}
	return ctx.JSON(res.OkByMessage("ok"))
}

func (e *ExampleApi) UpdateExample(ctx fiber.Ctx) error {
	example := &ent.Example{}
	if err := ctx.Bind().JSON(example); err != nil {
		return err
	}
	if errorMessage := common.StructFieldValidation(example); len(errorMessage) > 0 {
		return ctx.JSON(res.FiberRequestError(errorMessage))
	}
	if example.ID <= 0 {
		return ctx.JSON(res.FiberRequestError("example不存在"))
	}
	if err := e.service.UpdateExample(example); err != nil {
		return err
	}
	return ctx.JSON(res.OkByMessage("ok"))
}
