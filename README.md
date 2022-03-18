# Go Rest Clean Boilerplate

Clean boilerplate for Go REST APIs.

## Libraries

- [gofiber/fiber](https://github.com/gofiber/fiber): Web framework
- [goccy/go-json](https://github.com/goccy/go-json): Fast JSON tool
- [uber/zap](https://github.com/uber/zap): Fast logging tool
- [spf13/cobra](https://github.com/spf13/cobra): Cli commands
- [swaggo/swag](https://github.com/swaggo/swag): Generating Swagger documentation
- [stretchr/testify](https://github.com/stretchr/testify): Testing
- [mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver): Database driver

## Run

Simply use `docker compose up`

## Commands

    make Command

| Command  | Description                                            |
|----------|--------------------------------------------------------|
| all      | Runs fmt, lint, test, swagger and build with order     |
| build    | Builds API Server to build folder                      |
| clean    | Removes ./bin folder                                   |
| coverage | Runs all tests and shows total coverage                |
| fmt      | Checks code format                                     |
| help     | Shows this message                                     |
| install  | Downloads dependencies                                 |
| lint     | Runs lint tools                                        |
| run      | Runs API server                                        |
| run-swag | Generaes Swagger API documents and runs API Server     |
| swagger  | Generates Swagger API documents                        |
| test     | Runs tests without coverage                            |


## Project Structure 

```
├── cmd
│   └── server
├── docs
├── pkg
│   ├── api
│   │   ├── controller
│   │   ├── middleware
│   │   ├── repository
│   │   └── service
│   ├── configs
│   ├── db
│   ├── model
│   └── utils
│       └── helper
└── test
    └── unit
        ├── controller
        ├── repository
        └── service
```
