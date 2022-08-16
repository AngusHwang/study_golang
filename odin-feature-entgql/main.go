package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
	"odin/db"
	"odin/ent/migrate"
	"odin/handlers"
	"odin/tasks"
	"time"
)

func main() {
	d, err := db.Conn()

	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	err = d.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	c := cron.New(cron.WithSeconds())
	_, err = c.AddFunc("0 30 12 * * MON-FRI", func() {
		if err = tasks.TodayLunch(ctx); err != nil {
			log.Println(err)
		}
	})

	c.Start()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/graphql", handlers.Playground())
	r.POST("/graphql", handlers.GraphQL(d))
	r.POST("/slack/command", handlers.SlackCommand)

	log.Fatalln(r.Run(":3000"))
}
