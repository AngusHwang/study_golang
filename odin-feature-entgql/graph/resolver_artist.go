package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"
	"odin/ent"
	"odin/ent/artist"
	"odin/mail"
	"regexp"
)

func (r *mutationResolver) ApplyArtist(ctx context.Context, data ArtistInput) (*ent.Artist, error) {
	tx, err := r.client.Tx(ctx)

	if err != nil {
		return nil, err
	}

	hasPhoneNumber, err := tx.Artist.
		Query().
		Where(artist.PhoneNumberEQ(data.PhoneNumber)).
		Exist(ctx)

	if err != nil {
		log.Println(err)
		return nil, tx.Rollback()
	}

	var urlRegexp = regexp.MustCompile("^https?://")
	var phoneNumberRegexp = regexp.MustCompile("^\\d{3}\\d{3}\\d{4,6}$")

	if hasPhoneNumber == true {
		return nil, errors.New("phone number already exists")
	} else if data.Name == "" {
		return nil, errors.New("name field is required")
	} else if data.ExternalURL == "" {
		return nil, errors.New("external_url field is required")
	} else if data.PhoneNumber == "" {
		return nil, errors.New("phone_number field is required")
	} else if !urlRegexp.MatchString(data.ExternalURL) {
		return nil, errors.New("unexpected external url")
	} else if !phoneNumberRegexp.MatchString(data.PhoneNumber) {
		return nil, errors.New("unexpected phone number")
	}

	newArtist, err := tx.Artist.
		Create().
		SetName(data.Name).
		SetExternalURL(data.ExternalURL).
		SetPhoneNumber(data.PhoneNumber).
		SetDiscord(*data.Discord).
		SetRecommender(*data.Recommender).
		Save(ctx)

	if err != nil {
		log.Println(err)
		return nil, tx.Rollback()
	}

	go mail.SendArtistMail(data.Name, data.ExternalURL, data.PhoneNumber, *data.Discord, *data.Recommender, newArtist.CreatedAt)

	return newArtist, tx.Commit()
}
