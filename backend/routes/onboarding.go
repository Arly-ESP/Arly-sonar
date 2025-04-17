package routes

import (
	"github.com/arly/arlyApi/controllers"
	"github.com/arly/arlyApi/middleware"
	"github.com/gofiber/fiber/v2"
)

func SurveyRoutes(app *fiber.App) {
	// Survey Management
	app.Post("/api/surveys", middleware.JWTMiddleware, middleware.AdminMiddleware, controllers.CreateSurvey)  
	app.Put("/api/surveys/:id", middleware.JWTMiddleware, middleware.AdminMiddleware, controllers.UpdateSurvey)
	app.Delete("/api/surveys/:id", middleware.JWTMiddleware, middleware.AdminMiddleware, controllers.DeleteSurvey)

	// Public Survey Routes
	app.Get("/api/surveys", controllers.GetAllSurveys) 
	app.Get("/api/surveys/:id", controllers.GetSurveyDetails) 


	//Get survey by slud 
	app.Get("/api/surveys/slug/:slug", middleware.JWTMiddleware, controllers.GetSurveyBySlug)


	// User Survey Responses
	app.Post("/api/surveys/:id/responses", middleware.JWTMiddleware, controllers.SubmitSurveyResponse)
	app.Get("/api/surveys/:id/responses", middleware.JWTMiddleware, middleware.AdminMiddleware, controllers.GetAllSurveyResponses) 
	app.Get("/api/surveys/:id/responses/:user_id", middleware.JWTMiddleware, controllers.GetUserSurveyResponse)
}
