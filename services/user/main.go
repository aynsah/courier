package main

import (
	"courier/services/user/config"
	"courier/services/user/database"
	"courier/services/user/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error

	err = config.LoadConfig()
	if err != nil {
		return err
	}

	err = database.LoadDatabase()
	if err != nil {
		return err
	}

	m.router = gin.Default()

	return err
}

// @title UserManagement Service API Document
// @version 1.0
// @description List APIs of UserManagement Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8801
// @BasePath /api/v1/users
func main() {
	m := Main{}

	if err := m.initServer(); err != nil {
		fmt.Print(err)
	}

	defer database.DB.Close()

	routes.V1Application(m.router)

	m.router.Run(config.Config.Port)
}
