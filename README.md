# Go Pokedex REST API

<p align="center">
  <img src="https://lirp.cdn-website.com/61e0cb41/dms3rep/multi/opt/pngimg.com+-+pokeball_PNG24-1920w.png" width="200" alt="Nest Logo" />
</p>

Pokedex REST API allows users to keep track of caught pokemon and their stats: `TYPE`, `HP`, `ATTACK`, and `DEFENSE`. The API utilizes JWT authentification for additional security as well as a number of other tools (look below). The projects structure implements `clean architecture` and `dependecy injection` principles 

## Endpoints

#### API
- **/v1/pokemon &ensp;** `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Add pokemon
- **/v1/pokemon &ensp;** `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Get all pokemon
- **/v1/pokemon/{id} &ensp;** `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Get pokemon by id
- **/v1/pokemon/{id} &ensp;** `=>`  &ensp; **PUT** &ensp;  `=>` &ensp; Update pokemon by id
- **/v1/pokemon/{id} &ensp;** `=>`  &ensp; **DELETE** &ensp;  `=>` &ensp; Delete pokemon by id
  
#### Other
- **/health &ensp;**  `=>` &ensp; **GET** &ensp; `=>` &ensp; Ping the database connection
- **/auth/sign-up &ensp;**  `=>` &ensp; **POST** &ensp; `=>` &ensp; Create new pokemon trainer
- **/auth/sign-in &ensp;** `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Sign in with existing profile to generate JWT authentification token

## Tools used

- `App configuration` &nbsp; **=>**  &nbsp; [Viper](https://github.com/spf13/viper)
- `Logging` &nbsp; **=>**  &nbsp; [Logrus](https://github.com/sirupsen/logrus)
- `Routing` &nbsp; **=>**  &nbsp;  [Chi](https://github.com/go-chi/chi)
- `Database` &nbsp; **=>**  &nbsp;  Postgres + [pgx](https://github.com/jackc/pgx/)
- `Database migrations` &nbsp; **=>**  &nbsp; [Goose](https://github.com/pressly/goose#sql-migrations)
- `Containerization` &nbsp; **=>**  &nbsp; [Docker](http://docker.com/) + Docker Compose
- `Authentification and middleware` &nbsp; **=>**  &nbsp;  [JWT Go](https://github.com/golang-jwt/jwt)
- `Testing` &nbsp; **=>**  &nbsp; [Sqlmock](https://github.com/DATA-DOG/go-sqlmock) + [Gomock](https://github.com/uber-go/mock) + [Testify](https://github.com/stretchr/testify)

## Installation
```
make initUp
```

## Running the app

```bash
# rebuild containers
make build

# start the app
make run

# run psql utility
make startPsql
```
