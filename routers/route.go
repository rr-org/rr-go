package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naflafadia/go-trivia-app/controllers"
)

func SetupRoutes(f *fiber.App) {

	app := fiber.New()
	f.Mount("/api", app)

	app.Get("/user", controllers.GetUsers)
	app.Get("/user/:id", controllers.GetUser)

	app.Get("/avatar", controllers.GetAvatars)
	app.Get("/avatar/:id", controllers.GetAvatar)

	app.Get("/quiz", controllers.GetQuizzes)
	app.Get("/quiz/:id", controllers.GetQuiz)
}
