package main

import web "genaidemo/pkg/web"

func main() {

	application := web.CreateWebServer()

	application.Run("localhost:8080")
}
