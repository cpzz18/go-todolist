package main

import (
	"github.com/cpzz18/go-todo/controller"
	"github.com/cpzz18/go-todo/database"
	"github.com/labstack/echo"
)

func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
	e := echo.New()

	controller.NewGetAllTodosController(e, db)
	controller.NewCreateTodosController(e, db)
	controller.NewDeleteTodosController(e, db)
	controller.NewCheckTodosController(e, db)
	controller.NewUpdateTodosController(e, db)

	e.Start(":8000")

}
