package ent

import (
	"context"
	"fmt"
	"log"
)

func QueryAllMember(ctx context.Context, client *Client) []*Member {
	mArr, err := client.Member.
		Query().
		All(ctx)
	if err != nil {
		fmt.Errorf("failed querying user: %w", err)
		return nil
	}
	log.Println("user returned: ", mArr)
	return mArr
}

func QueryAllId(ctx context.Context, client *Client) []int {
	iArr, err := client.Member.
		Query().
		IDs(ctx)
	if err != nil {
		fmt.Errorf("failed querying user: %w", err)
		return nil
	}
	log.Println("user returned: ", iArr)
	return iArr
}

func CreateMember(ctx context.Context, client *Client, nm Member) (*Member, error) {
	iArr := QueryAllId(ctx, client)
	id := 0
	for i := 0; i < len(iArr); i++ {
		if iArr[i] > id {
			id = iArr[i]
		}
	}
	id++

	m, err := client.Member.
		Create().
		SetID(id).
		SetName(nm.Name).
		SetNick(nm.Nick).
		SetTeam(nm.Team).
		SetDetail(nm.Detail).
		SetImg(nm.Img).
		Save(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", m)
	return m, nil
}

func DeleteMember(ctx context.Context, client *Client, id int) {
	err := client.Member.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		fmt.Errorf("failed delete user: %w", err)
	}
}

func UpdateMember(ctx context.Context, client *Client, m Member) {
	err := client.Member.UpdateOneID(m.ID).
		SetName(m.Name).
		SetNick(m.Nick).
		SetTeam(m.Team).
		SetDetail(m.Detail).
		SetImg(m.Img).
		Exec(ctx)

	if err != nil {
		fmt.Errorf("faled update user: %w", err)
	}
}
