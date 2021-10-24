package main

import (
	"courier/services/logistic/config"
	"courier/services/logistic/database"
	"courier/services/logistic/routes"
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

// @title LogisticManagement Service API Document
// @version 1.0
// @description List APIs of LogisticManagement Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8802
// @BasePath /api/v1/logistics
func main() {
	m := Main{}

	if err := m.initServer(); err != nil {
		fmt.Print(err)
	}

	defer database.DB.Close()

	routes.V1Application(m.router)

	m.router.Run(config.Config.Port)
}
