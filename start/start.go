package start

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"todolist/config"
	"todolist/db"
	"todolist/model"
	"todolist/handler"
)

func StartApplication() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config : %v", err)
	}

	db, err := db.NewMysqlConn()
	if err != nil {
		log.Fatalf("Error connect database : %v", err)
	}
	db.AutoMigrate(&model.Activity{}, &model.Todo{})

	app := fiber.New()
	handler.MapHandler(app, db)

	log.Fatal(app.Listen(":3030"))
}