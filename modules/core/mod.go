package core

import (
	"github.com/kryptonpust/go-echo-test/config"
	"github.com/kryptonpust/go-echo-test/modules/core/handlers"
	"github.com/kryptonpust/go-echo-test/modules/core/repositories"
	"github.com/kryptonpust/go-echo-test/modules/core/usecases"
	"github.com/kryptonpust/go-echo-test/pkg/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Module ModuleInstance = &coreModule{}

type coreModule struct{}

func (coreModule) RegisterRepositories(container *dig.Container) error {
	container.Provide(repositories.NewPgsqlUserRepository)
	container.Provide(repositories.NewPgsqlOrgRepository)
	container.Provide(repositories.NewPgsqlUserOrgRepository)
	return nil
}

func (coreModule) RegisterUseCases(container *dig.Container) error {
	container.Provide(usecases.NewUserUsecase)
	container.Provide(usecases.NewOrgUsecase)
	return nil
}

func (coreModule) RegisterHandlers(g *echo.Group, container *dig.Container) error {
	return container.Invoke(func(
		appConf *config.AppConfig,
		middManager *middlewares.MiddlewareManager,
		userUsecase usecases.UserUsecase,
		orgUsecase usecases.OrgUsecase,
	) {
		handlers.NewOrgHandler(g, middManager, orgUsecase)
		handlers.NewUserHandler(g, middManager, userUsecase)
		handlers.NewAuthHandler(g, middManager, userUsecase, appConf)
	})
}
