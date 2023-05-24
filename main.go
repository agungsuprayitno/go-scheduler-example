package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-rest-postgres/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

var (
	server              *gin.Engine
)

func init() {
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("")
	router.POST("/scheduler", KafkaMessage)

	log.Fatal(server.Run(":" + config.ServerPort))
}

func KafkaMessage(ctx *gin.Context) {
	s := gocron.NewScheduler(time.UTC)

	// Every starts the job immediately and then runs at the 
	// specified interval
	_, err := s.Every(5).Seconds().Do(func(){ 
		fmt.Println("Hello this message will be appear every 5 seconds.")
	})
	if err != nil {
		// handle the error related to setting up the job
	}

	// strings parse to duration
	// s.Every("5m").Do(func(){ ... })

	// s.Every(5).Days().Do(func(){ ... })

	// s.Every(1).Month(1, 2, 3).Do(func(){ ... })

	// // set time
	// s.Every(1).Day().At("10:30").Do(func(){ ... })

	// // set multiple times
	// s.Every(1).Day().At("10:30;08:00").Do(func(){ ... })

	// s.Every(1).Day().At("10:30").At("08:00").Do(func(){ ... })

	// // Schedule each last day of the month
	// s.Every(1).MonthLastDay().Do(func(){ ... })

	// // Or each last day of every other month
	// s.Every(2).MonthLastDay().Do(func(){ ... })

	// // cron expressions supported
	// s.Cron("*/1 * * * *").Do(task) // every minute

	// // cron second-level expressions supported
	// s.CronWithSeconds("*/1 * * * * *").Do(task) // every second
	


	// you can start running the scheduler in two different ways:
	// starts the scheduler asynchronously
	// ### s.StartAsync()

	// or starts the scheduler and blocks current execution path
	s.StartBlocking()

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success push data into kafka",
		"data":    "Scheduler is Running now.",
	})
}

