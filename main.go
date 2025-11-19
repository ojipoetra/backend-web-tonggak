package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ojipoetra/backend-web-tonggak/app"
	"github.com/ojipoetra/backend-web-tonggak/controllers"
	"github.com/ojipoetra/backend-web-tonggak/repository"
	"github.com/ojipoetra/backend-web-tonggak/services"
)

func main(){
	err := godotenv.Load()
	
	if err != nil {
		log.Println("Env variable failed to load")
	}
	log.Println("Env load successfully")
	
	db, err := app.Init()
	
	if err != nil {
		panic(err)
	}
	defer db.Close()
	KamarRepository := repository.NewKamarRepository()

	kamarService := services.NewKamarService(KamarRepository, db)

	kamarController := controllers.NewKamarController(kamarService)

	router := app.NewRouter(kamarController)

	
	// route.GET("api/webtonggak/kamar", kamarcontroller.Index)
	router.Run()
}
