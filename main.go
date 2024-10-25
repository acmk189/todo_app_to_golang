package main

import (
	"fmt"

	"github.com/acmk189/golang_udemy_todo_app/app/controllers"
	"github.com/acmk189/golang_udemy_todo_app/app/models"
)

func main() {

	fmt.Println(models.Db)

	controllers.StartMainServer()

}
