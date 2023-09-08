# Reproduction Steps

1. Open main.go file in `cmd/todo/main.go` and change the `dsn` variable to your Turso db link
2. Run the app `go run ./cmd/todo`
3. Open https://localhost:8081 on your browser
4. Execute the below mutation to create exercise and muscles group resources

```gql
mutation {
  createExercise(
    input: {
    name: "Plank", 
    image: "plank image", 
    createMusclesGroup: [
    {name: "Abdominal", image: "abdominal image"}, 
    {name: "Back", image: "back image"}, 
    {name: "Hip Flexor", image: "hip flexor image"}, 
    {name: "Posterior (Back) Leg", image: "posterior back leg image"}
    ]
    }
  ) {
    id
    name
    musclesGroups {
      edges {
        node {
          id
          name
          image
        }
      }
    }
  }
}
```
    
5. Execute the below query to query exercise resources

```gql
{
  exercises(first: 1) {
    edges {
      node {
        musclesGroups(first: 1) {
          edges {
            node {
              id
              name
            }
          }
          pageInfo{
            hasNextPage
            endCursor
          }
        }
      }
    }
  }
}
```
6. You will get an error message like this 
`"sql/scan: missing struct field for column: COUNT (*) (COUNT )"`
after executing the above query
7. If you want to try the above operations with sqlite3 client, change the branch
from master to `with-sqlite3-client` branch and redo from the second step
