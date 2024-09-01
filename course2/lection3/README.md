# Lection 3
## Create API best practives

- [Done] init go (go mod) relatively (with path)
- [Done] create enter point to the app ```cmd/<app_name>/main.go```
- [Done] create basic main.go with fmt.Println
- [Done] build the project ```go build -v ./cmd/<app_name>/```
- [Done] init server core and co:fig ```internal/app/<app_name>/<app_name>.go``` and ```internal/app/<app_name>/config.go```
- [Done] create config file```configs/[<app_name>.toml | .env]```
- [Done] add a config to the API constructor as a private field
- [Done] read config in cmd/../app.go using https://github.com/BurntSushi/toml
- [Done] use ```flag.StringVar(...``` to get path of the config
- [Done] implement running config as ```api.exe -format [.env|<app_name>.toml] -path configs/[env|toml]```
- [Done] add log_level into config
- [Done] add logger to config and install https://github.com/sirupsen/logrus
- [Done] create ```internal/app_name/helper.go``` and add configureLoggerField() error
- init mux.router to API obj
- add func (a *API) configureRouterField

## Connect DB
- install a lib for working with SQL. Default is ```database/sql```
- init ```storage/storage.go``` it should have Open, Close, New. ```Store struct {config: *Config}```
- add storage.Storage into API object
- create storage.Config with constructor new Config{DatabaseUri: string}
- add configureStorageField to API helper
- add db: *sql.DB into Storage struct
- inside ```Storage.open``` implement connection to postgre. Add ping to make the connection active

## Migrations
- install scoop on windows (analog for MAC homebrew but for windows)
- install golang migrate
- ```scoop install migrate```
- ```migrate create -ext sql -dir migrations UsersCreationMigration``` should create two migrations
- add a migration up/down create/delete table users with id, name, email (unique), age, password_hash
- Run the up migration

## Homework
- read golang-standards/project-layout
- use/implement make (makefile) for running the app
