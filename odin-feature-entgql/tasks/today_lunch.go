package tasks

import (
	"context"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/slack-go/slack"
	"math/rand"
	"odin/db"
	"odin/ent"
	"odin/ent/store"
	"odin/utils"
	"os"
	"strings"
	"time"
)

func TodayLunch(ctx context.Context) error {
	c, err := db.Conn()

	if err != nil {
		return err
	}

	hasData, err := c.Store.
		Query().
		Exist(ctx)

	if err != nil {
		return err
	}

	if hasData == false {
		return errors.New("데이터가 없어요 👻")
	}

	stores, err := c.Store.
		Query().
		Where(
			store.Or(
				store.SentAtIsNil(),
				store.SentAtLT(time.Now().AddDate(0, 0, -3)),
			),
		).
		All(ctx)

	if err != nil {
		return err
	}

	if len(stores) <= 0 {
		return errors.New("데이터가 부족해요 👻")
	}

	rand.Seed(time.Now().UnixNano())
	shuffledStores := lo.Shuffle(stores)

	end := 4
	if len(shuffledStores) < end {
		end = len(shuffledStores)
	}

	slicedStores := shuffledStores[0:end]

	blocks := []slack.Block{
		slack.SectionBlock{
			Type: "section",
			Text: slack.NewTextBlockObject("mrkdwn", "*⭐️ 오늘의 추천 가게입니다 ⭐*", false, false),
		},
	}

	for _, s := range slicedStores {
		foods := strings.Join(s.Food, ", ")
		middleTemplate := fmt.Sprintf("*<https://map.naver.com/v5/search/성수+%s|%s>* ⇒ *%s*  🚶%d분", s.Name, s.Name, s.Location, s.OnFoot)
		foodTemplate := fmt.Sprintf("> %s", foods)

		blocks = append(blocks,
			slack.SectionBlock{
				Type: "section",
				Text: slack.NewTextBlockObject("mrkdwn", middleTemplate, false, false),
			},
			slack.SectionBlock{
				Type: "section",
				Text: slack.NewTextBlockObject("mrkdwn", foodTemplate, false, false),
			},
			slack.DividerBlock{
				Type: "divider",
			},
		)
	}

	msg := &slack.WebhookMessage{
		Blocks: &slack.Blocks{
			BlockSet: blocks,
		},
	}

	tx, err := c.Tx(ctx)

	if err != nil {
		return err
	}

	_, err = tx.Store.
		Update().
		Where(
			store.IDIn(lo.Map(slicedStores, func(s *ent.Store, key int) int { return s.ID })...),
		).
		SetSentAt(time.Now()).
		Save(ctx)

	if err != nil {
		return utils.TxRollback(tx, err)
	}

	err = slack.PostWebhook(os.Getenv("WEBHOOK_URL"), msg)

	if err != nil {
		return utils.TxRollback(tx, err)
	}

	return tx.Commit()
}
