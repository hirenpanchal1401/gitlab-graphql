package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	api "gitLab-graphql/api/models"
	"gitLab-graphql/graph/generated"
	"gitLab-graphql/graph/model"

	"github.com/shurcooL/graphql"
)

func (r *queryResolver) LastProjects(ctx context.Context, last int) (*model.CustomResult, error) {
	variables := map[string]interface{}{
		"last": graphql.Int(last),
	}
	var q = &api.LastProjects
	err := api.Client.Query(context.Background(), q, variables)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var names string
	var sumOfFork int
	for i, v := range q.Projects.Nodes {
		if i == len(q.Projects.Nodes)-1 {
			names += v.Name
		} else {
			names += fmt.Sprintf("%s, ", v.Name)
		}
		sumOfFork += v.ForksCount
	}

	return &model.CustomResult{
		Names:     names,
		SumOfFork: sumOfFork,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Projects(ctx context.Context) (*model.ProjectConnection, error) {
	var q = &api.ProjectQuery
	err := api.Client.Query(context.Background(), q, nil)
	if err != nil {
		fmt.Println(err.Error())
		return &q.Projects, err
	}
	return &q.Projects, nil
}
