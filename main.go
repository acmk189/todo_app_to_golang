package main

import "github.com/acmk189/golang_udemy_todo_app/app/models"

func main() {
	/*
					fmt.Println(config.Config.Port)
					fmt.Println(config.Config.SQLDriver)
					fmt.Println(config.Config.DbName)
					fmt.Println(config.Config.LogFile)

					log.Println("test")

				fmt.Println(models.Db)

			u := &models.User{}
			u.Name = "test"
			u.Email = "test@example.com"
			u.Password = "testtest"
			fmt.Println(u)

			u.CreateUser()

				u, _ := models.GetUser(1)

				fmt.Println(u)

				u.Name = "Test2"
				u.Email = "test2@example.com"
				u.UpdateUser()

				u, _ = models.GetUser(1)
				u.UpdateUser()
				fmt.Println(u)

				u.DeleteUser()
				u, _ = models.GetUser(1)
				fmt.Println(u)

			user, _ := models.GetUser(2)
			user.CreateTodo("First Todo")

				t, _ := models.GetTodos
				fmt.Println(t)

			user, _ := models.GetUser(1)

				user.CreateTodo("Final Todo")

					todos, _ := models.GetTodos()
					for _, v := range todos {
						fmt.Println(v)
					}


		todos, _ := user.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		t, _ := models.GetTodo(1)
		t.Content = "Update Todo"
		t.UpdateTodo()
	*/

	t, _ := models.GetTodo(1)
	t.DeleteTodo()

}
