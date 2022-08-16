package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-crud/ent"
	"gqlgen-crud/ent/member"
)

// CreateMember is the resolver for the createMember field.
func (r *mutationResolver) CreateMember(ctx context.Context, input MemberInput) (*ent.Member, error) {
	cl := r.Client
	//defer r.Close()
	id := 0
	iArr, err := cl.Member.
		Query().
		IDs(ctx)

	for i := 0; i < len(iArr); i++ {
		if iArr[i] > id {
			id = iArr[i]
		}
	}
	id++

	if err != nil {
		fmt.Println(err)
	}
	member, err := cl.Member.
		Create().
		SetID(id).
		SetName(input.Name).
		SetNick(input.Nick).
		SetTeam(input.Team).
		SetDetail(input.Detail).
		SetImg(input.Img).
		Save(ctx)

	fmt.Println(member)
	fmt.Println(err)
	return member, err
}

// UpdateMember is the resolver for the updateMember field.
func (r *mutationResolver) UpdateMember(ctx context.Context, input MemberInput) (*ent.Member, error) {
	bulk := r.Client.Member.
		Update().
		Where(member.ID(*input.ID))

	bulk.SetName(input.Name)
	bulk.SetNick(input.Nick)
	bulk.SetTeam(input.Team)
	bulk.SetDetail(input.Detail)
	bulk.SetImg(input.Img).Save(ctx)

	return &ent.Member{ID: *input.ID, Name: input.Name, Nick: input.Nick, Team: input.Team, Detail: input.Detail, Img: input.Img}, nil
}

// DeleteMember is the resolver for the deleteMember field.
func (r *mutationResolver) DeleteMember(ctx context.Context, id *int) (*ent.Member, error) {
	matched, err := r.Client.Member.
		Query().
		Where(member.ID(*id)).
		Only(ctx)

	if err != nil {
		fmt.Println(err)
	}

	err = r.Client.Member.
		DeleteOneID(*id).
		Exec(ctx)
	return matched, err
}

// Members is the resolver for the members field.
func (r *queryResolver) Members(ctx context.Context) ([]*ent.Member, error) {
	members, err := r.Client.Member.
		Query().
		All(ctx)

	return members, err
}

// Member is the resolver for the member field.
func (r *queryResolver) Member(ctx context.Context, id int) (*ent.Member, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
