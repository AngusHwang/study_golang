package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/slack-go/slack"
	"log"
	"net/http"
	"odin/db"
	"strings"
)

func SlackCommand(ctx *gin.Context) {
	c, err := db.Conn()

	if err != nil {
		log.Fatalln(err)
	}

	s, err := slack.SlashCommandParse(ctx.Request)

	if err != nil {
		log.Fatalln(err)
	}

	switch s.Command {
	case "/오점무":
		stores, err := c.Store.
			Query().
			All(ctx)

		if err != nil {
			log.Println(err)
		}

		if len(stores) <= 0 {
			log.Println("데이터가 없어요 👻")
		}

		store := lo.Shuffle(stores)[0]

		foods := strings.Join(store.Food, ", ")
		middleTemplate := fmt.Sprintf("*<https://map.naver.com/v5/search/성수+%s|%s>* ⇒ *%s*", store.Name, store.Name, store.Location)
		foodTemplate := fmt.Sprintf("> %s", foods)

		blocks := []slack.Block{
			slack.SectionBlock{
				Type: "section",
				Text: slack.NewTextBlockObject("mrkdwn", middleTemplate, false, false),
			},
			slack.SectionBlock{
				Type: "section",
				Text: slack.NewTextBlockObject("mrkdwn", foodTemplate, false, false),
			},
			slack.DividerBlock{Type: "divider"},
		}

		msg := &slack.WebhookMessage{
			ResponseType: slack.ResponseTypeInChannel,
			Blocks: &slack.Blocks{
				BlockSet: blocks,
			},
		}

		ctx.JSON(http.StatusOK, msg)
	}
}
