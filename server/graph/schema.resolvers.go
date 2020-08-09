package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	"github.com/redshiftkeying/zen-go/graph/generated"
	"github.com/redshiftkeying/zen-go/model"
)

func (r *mutationResolver) CreateRouter(ctx context.Context, input model.NewRouter) (*model.Router, error) {
	router := &model.Router{
		ID:        uuid.Must(uuid.NewV4()).String(),
		Title:     input.Theme,
		Icon:      input.Icon,
		Order:     input.Order,
		Path:      input.Path,
		Exact:     input.Exact,
		Component: input.Component,
	}
	err := model.Create(router)
	return router, err
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "hello,world!", nil
}

func (r *queryResolver) Story(ctx context.Context) ([]*model.UserStory, error) {
	return nil, nil
}

func (r *queryResolver) Routers(ctx context.Context) ([]*model.Router, error) {
	var rs []model.Router
	err := model.GetAll(&rs)
	if err != nil {
		log.Println(err)
	}
	var res []*model.Router
	for _, v := range rs {
		res = append(res, &v)
	}
	return res, err
}

func (r *userStoryResolver) Dependency(ctx context.Context, obj *model.UserStory) ([]*model.UserStory, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// UserStory returns generated.UserStoryResolver implementation.
func (r *Resolver) UserStory() generated.UserStoryResolver { return &userStoryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userStoryResolver struct{ *Resolver }
