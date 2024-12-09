package main

func (app *Application) routes() {
	app.s.GET("/", app.GetProperties)
	app.s.GET("/:id", app.GetPropertyById)
}
