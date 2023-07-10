package services

import (

	"proton/internal/common"
	"proton/internal/router"
	"github.com/spf13/viper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func Run() error {
	// init env
	// err := common.LoadEnv()
	// if err != nil {
	// 	return err
	// }

	// init db
	err := common.InitDB()
	if err != nil {
		return err
	}

	// defer closing db
	// defer common.CloseDB()

	// create app
	app := fiber.New()

	// add basic middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// add routes
	router.AddcourseGroup(app)

	// start server
	port := viper.GetString("mongo.port")
	err = app.Listen(":" + port)

	return err
}
