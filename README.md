# Go_VueJS_Web_Application
Web Application(Golang, VueJS, Cognito,  Cloud Front, S3, EC2)


## run container 
```
docker-compose down --volumes
docker-compose up -d --build
```

## Backend
```
//folder
cd backend

//formatting
go fmt ./...

```



## URL
```
//get all users
http://localhost:8080/users


//create user
http://localhost:8080/createUser?name=hogehoge&email=hogehoge@com&gender=other
```
