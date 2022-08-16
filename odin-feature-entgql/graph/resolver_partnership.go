package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"odin/db"
	"odin/ent"
	"odin/mail"
	"regexp"
)

func (r *mutationResolver) AskForPartnership(ctx context.Context, data PartnershipInput) (*ent.Partnership, error) {
	d, err := db.Conn()
	if err != nil {
		log.Printf("failed connecting to postgres: %v", err)
	}

	var emailRegexp = regexp.MustCompile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")

	if data.Name == "" {
		return nil, errors.New("name field is required")
	} else if data.Company == "" {
		return nil, errors.New("company field is required")
	} else if data.Email == "" {
		return nil, errors.New("email field is required")
	} else if !emailRegexp.MatchString(data.Email) {
		return nil, errors.New("unexpected ema	il address")
	}

	p, err := d.Partnership.
		Create().
		SetName(data.Name).
		SetCompany(data.Company).
		SetEmail(data.Email).
		SetContent(*data.Content).
		Save(ctx)
	if err != nil {
		log.Printf("failed creating to partnership: %v", err)
	}

	go mail.SendPartnershipMail(data.Name, data.Company, data.Email, *data.Content, p.CreatedAt)

	return p, nil
}

func (r *queryResolver) Partnerships(ctx context.Context) (*ent.Partnership, error) {
	panic(fmt.Errorf("not implemented"))
}
