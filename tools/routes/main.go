package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/kryptonpust/go-echo-test/cmd/api/di"
	"github.com/kryptonpust/go-echo-test/config"
	"github.com/kryptonpust/go-echo-test/infrastructure/datastore"
	"github.com/kryptonpust/go-echo-test/pkg/logger"
	"github.com/labstack/echo/v4"
)

func main() {
	generateRoutes()
}

func generateRoutes() {
	e := echo.New()
	container := di.BuildDIContainer(
		&datastore.MasterDbInstance{},
		&datastore.SlaveDbInstance{},
		&config.AppConfig{},
	)
	di.RegisterModules(e, container)

	mapRoutes := map[string]map[string]string{}
	count := 0
	for _, r := range e.Routes() {
		if strings.HasPrefix(r.Name, "github.com") {
			continue
		}
		count++
		acl := mapRoutes[r.Path]
		if len(acl) == 0 {
			acl = map[string]string{}
		}
		acl[r.Method] = r.Name
		mapRoutes[r.Path] = acl
	}

	logger.Log().Info("Generated routes: ", count)
	data, err := json.MarshalIndent(mapRoutes, "", "  ")
	if err != nil {
		logger.Log().Fatalf("error json marshal: %v", err)
	}
	ioutil.WriteFile("./pkg/authz/routes.json", data, 0644)
}
