package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"
	"odin/ent"
	"odin/ent/store"
	"odin/utils"
)

func (r *mutationResolver) CreateStore(ctx context.Context, data StoreInput) (*ent.Store, error) {
	storeNames, err := r.client.Store.
		Query().
		Select(store.FieldName).
		Strings(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	for _, n := range storeNames {
		if utils.RemoveSpacing(n) == utils.RemoveSpacing(data.Name) {
			return nil, errors.New("이미 등록되어 있는 가게에요\U0001F979")
		}
	}

	return r.client.
		Store.
		Create().
		SetName(data.Name).
		SetLocation(data.Location).
		SetFood(data.Food).
		SetOnFoot(*data.OnFoot).
		Save(ctx)
}

func (r *mutationResolver) UpdateStore(ctx context.Context, data StoreInput) (*ent.Store, error) {
	bulk := r.client.Store.
		Update().
		Where(store.ID(*data.ID))

	matchedStore, err := r.client.Store.
		Query().
		Where(store.ID(*data.ID)).
		Only(ctx)

	if err != nil {
		return nil, errors.New("해당 가게를 찾을 수 없어요\U0001F97A")
	}

	storeNames, err := r.client.Store.
		Query().
		Select(store.FieldName).
		Strings(ctx)

	if err != nil {
		return nil, err
	}

	for _, n := range storeNames {
		if utils.RemoveSpacing(n) == utils.RemoveSpacing(data.Name) {
			if n == data.Name {
				return nil, errors.New("이미 등록되어 있는 가게에요\U0001F979")
			}
		}
	}

	if matchedStore.Name != data.Name {
		bulk = bulk.SetName(data.Name)
	}

	if matchedStore.Location != data.Location {
		bulk = bulk.SetLocation(data.Location)
	}

	if matchedStore.OnFoot != *data.OnFoot {
		bulk = bulk.SetOnFoot(*data.OnFoot)
	}

	_, err = bulk.
		SetFood(data.Food).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return r.client.Store.
		Query().
		Where(store.ID(*data.ID)).
		Only(ctx)
}

func (r *mutationResolver) DeleteStore(ctx context.Context, id int) (*ent.Store, error) {
	matchedStore, err := r.client.Store.
		Query().
		Where(store.ID(id)).
		Only(ctx)

	if matchedStore == nil {
		return nil, errors.New("해당 가게를 찾을 수 없어요\U0001F97A")
	}

	_, err = r.client.Store.
		Delete().
		Where(store.ID(id)).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	log.Printf("%s 가 삭제되었습니다", matchedStore.Name)

	return matchedStore, err
}

func (r *queryResolver) Store(ctx context.Context, id int) (*ent.Store, error) {
	oneStore, err := r.client.Store.
		Query().
		Where(store.ID(id)).
		Only(ctx)

	if oneStore == nil {
		return nil, errors.New("해당 가게를 찾을 수 없어요\U0001F97A")
	}

	if err != nil {
		return nil, err
	}

	return oneStore, err
}

func (r *queryResolver) Stores(ctx context.Context) ([]*ent.Store, error) {
	stores, err := r.client.Store.
		Query().
		Order(ent.Desc(store.FieldID)).
		All(ctx)

	if len(stores) <= 0 {
		return nil, errors.New("데이터를 찾을 수 없어요. 데이터 등록좀..🙇")
	}

	if err != nil {
		return nil, err
	}

	return stores, err
}
