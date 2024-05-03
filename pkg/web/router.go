package web

import "github.com/gin-gonic/gin"

func CreateWebServer() *gin.Engine {
	app := gin.Default()

	initRouter(app)

	return app
}

func initRouter(app *gin.Engine) {
	app.POST("/translate", CreateTranslationHandler)
}
