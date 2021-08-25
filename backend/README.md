# Backend
1. execute `go run main.go` to activate server
2. access localhost
   ```
   get all users:
   http://localhost:8080/users
   
   create new users:
   http://localhost:8080/createUser?name=newUserName&email=hogehoge@com&gender=male
   ```


# Graphql
1. *.graphqls files is Graphql definition file. You can modify them.
2. After modifying *.graphqls files, run `gqlgen init` to generated golang codes.
3. Implement functions in *.resolver.go files. 

```
# sample query / mutation

## get all users
query {
  users {
    id
    name
    email
    gender
  }
}

mutation {
    createUser(input : {name: "kohekohe", email: "koehkoe@kohe.net",gender: "other"}
    ){
    name
    email
    gender
  }
}

## s


```
