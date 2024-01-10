## Dating App Service

Backend service for dating mobile apps. We're using port adapter pattern

### Stacks
1. Go 1.20
   1. ORM: Gorm
   2. HTTP Framework: fiber v2
   3. Config reader: viper
   4. Mocking: go-mock
2. PostgreSQL
3. Redis

### Code structure
```bash
.
├── config                   # init config
├── db                       # db schema migrations
├── internal                 # source files
│   ├── adapter              # mainly business logic goes here 
│   │   ├── cache            # caching data client, we're using redis
│   │   ├── db               # db client, we're using postgreSQL
│   │   ├── repositories     # bridge for services and DB
│   │   │   └── model        # represent our tables   
│   │   └── usecase          # business logic goes here
│   ├── bcrypt               # for hashing password
│   ├── constants            # constant values
│   ├── entity               # struct of business
│   ├── handler              # http handler method
│   ├── interfaces           # blueprint of methods
│   ├── logger               # for logging purpose
│   └── mocks                # mock of dependencies
├── utils                    # source files support 
├── go.mod                   
├── go.sum
├── main.go
└── README.md
```


### How to run

#### Pre-requisite
- create database first on your local machine. Either installed on your machine or docker
- run sql command on db/schema.sql for build required tables


1. copy config.yaml.example to config.yaml
2. fill up config.yaml based on available configs
3. run this command on your favorite terminal
```go
    go run main.go
```