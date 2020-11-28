package routes

import (
	"github.com/aandrade1980/golang-fiber-basic-todo-app/controllers"
	"github.com/gofiber/fiber/v2"
)

// TodoRoute manage all routes
func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
	route.Post("", controllers.CreateTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Get("/:id", controllers.GetTodo)
}
