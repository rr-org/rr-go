package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naflafadia/go-trivia-app/controllers"
)

func SetupRoutes(f *fiber.App) {

	app := fiber.New()
	f.Mount("/api", app)

	app.Get("/users", controllers.GetUsers)
	app.Get("/user/:id", controllers.GetUser)

	app.Get("/avatars", controllers.GetAvatars)
	app.Get("/avatar/:id", controllers.GetAvatar)

	app.Get("/quizzes", controllers.GetQuizzes)
	app.Get("/quiz/:id", controllers.GetQuiz)
}
