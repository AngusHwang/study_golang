package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"odin/ent"
	"odin/ent/user"
	"odin/graph/directives"
	"odin/securities"
)

func (r *mutationResolver) Login(ctx context.Context, data UserInput) (string, error) {
	u, err := r.client.User.
		Query().
		Where(user.UsernameEQ(data.Username)).
		Only(ctx)

	if err != nil {
		return "", errors.New("아이디 또는 비밀번호를 잘못 입력했습니다.")
	}

	if err = securities.ComparePassword(u.Password, data.Password); err != nil {
		return "", errors.New("아이디 또는 비밀번호를 잘못 입력했습니다.")
	}

	return securities.GenerateWithUserTokens(u)
}

func (r *mutationResolver) Register(ctx context.Context, data UserInput) (string, error) {
	_, err := r.client.User.
		Query().
		Where(user.UsernameEQ(data.Username)).
		Only(ctx)

	if err == nil {
		return "", errors.New("이미 사용중이거나 탈퇴한 아이디입니다.")
	}

	hashedPassword := securities.HashPassword(data.Password)

	_, err = r.client.User.
		Create().
		SetUsername(data.Username).
		SetPassword(hashedPassword).
		Save(ctx)

	if err != nil {
		return "", err
	}

	return "성공적으로 생성되었습니다!", nil
}

func (r *queryResolver) AuthPing(ctx context.Context) (string, error) {
	claim, ok := ctx.Value(directives.UserCtxKey).(*securities.JwtCustomClaim)
	if !ok {
		return "", errors.New("asdf")
	}

	fmt.Println(claim.ID)

	return "authPong", nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *userResolver) Name(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

type userResolver struct{ *Resolver }
