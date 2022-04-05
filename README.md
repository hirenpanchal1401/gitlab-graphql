**_Project Setup_**

- clone repo : `git clone https://github.com/hirenpanchal1401/gitlab-graphql.git`
- in same dir, run `cd gitlab-graphql`
- run `go install`
- set .env
- run `go run ./server.go`
- visit `http://localhost:8080/`

_Example Query_

```
query last_projects($n : Int = 5) {
  last_projects(last:$n) {
		names
    sumOfFork
  }
}
```
