package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/presenters"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/domain"
	"strconv"

	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/graph/generated"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	userPresenter := presenters.NewUserPresenter(r.DB)
	user := domain.User{
		Name:   input.Name,
		Email:  input.Email,
		Gender: input.Gender,
	}
	createdUser, err := userPresenter.CreateUser(user)
	if err != nil {
		panic(fmt.Errorf("panic in CreateUser() at resolvers.go"))
	}
	return &model.User{
		ID:     strconv.FormatUint(uint64(user.ID),10),
		Name:   createdUser.Name,
		Email:  createdUser.Email,
		Gender: createdUser.Gender,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	userPresenter := presenters.NewUserPresenter(r.DB)
	users, err := userPresenter.GetAllUsers()
	if err != nil {
		panic(fmt.Errorf("panic in Users() at resolvers.go"))
	}

	var modelUsers []*model.User
	for _, user := range users {
		modelUser := model.User{
			ID:     strconv.FormatUint(uint64(user.ID),10),
			Name:   user.Name,
			Email:  user.Email,
			Gender: user.Gender,
		}
		modelUsers = append(modelUsers, &modelUser)
	}
	return modelUsers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
