// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CustomResult struct {
	Names     string `json:"names"`
	SumOfFork int    `json:"sumOfFork"`
}

type Project struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	ForksCount  int    `json:"forksCount"`
}

type ProjectConnection struct {
	Nodes []*Project `json:"nodes"`
}