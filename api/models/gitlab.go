package api

import (
	"context"
	"gitLab-graphql/graph/model"
	"os"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

var ProjectQuery struct {
	Projects model.ProjectConnection `graphql:"projects"`
}

var LastProjects struct {
	Projects model.ProjectConnection `graphql:"projects(last: $last)"`
}

var Client *graphql.Client

func init() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	Client = graphql.NewClient("https://gitlab.com/api/graphql", httpClient)
}
