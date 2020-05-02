# go-docker-test
Little repo using go + go modules, docker, docker-compose, mongoDB and live reloading

## Requirements

1. Docker
2. Docker-compose

## Install
```
1. Clone the project
2. Set the env vars in .env file (mongoDB instance will use the mongo credentials to config)
```

## Deploy

```
1. Execute `docker-compose up`
```

## Build

```
1. If you update the project source code, you can rebuild and deploy executing `docker-compose up --build`
```


## Endpoints

> [GET] `/users` fetch users in DB.
> List all users in a DB.

> [GET] `/users/create` create a new user in DB.
> Create a new user in DB
> Pass user info in query params

> Example url request - Next url request:
```
    /users/create?name=Julian&age=25&phone=123456&email=julian@docker.com
```
> Will create the next user:
```
    {
        "id": "5eada85d8a5e99345f372350",
        "name": "Julian",
        "age": 25,
        "phone": "123456"
        "email": "julian@docker.com"
    }
```
