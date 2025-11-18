package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ojipoetra/backend-web-tonggak/app"
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
	db.Close()
	route := gin.Default()

	
	// route.GET("api/webtonggak/kamar", kamarcontroller.Index)
	route.Run()
}
