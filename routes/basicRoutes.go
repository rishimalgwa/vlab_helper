package routes

import (
	"vlab/controllers"

	"github.com/gofiber/fiber/v2"
)

func BasicRoutes(incomingRoutes *fiber.App) {

	incomingRoutes.Get("/get-links", controllers.GetLinks)
	//incomingRoutes.Post("/users/signup", RegisterUser)

}
