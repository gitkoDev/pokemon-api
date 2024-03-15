# Go Pokedex REST API

Pokedex REST API allows users to keep track of caught pokemon and their stats: `TYPE`, `HP`, `ATTACK`, and `DEFENSE`. The API utilizes JWT authentification for additional security.

## Endpoints

### API
- **/v1/pokemon &ensp;** `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Add pokemon
- **/v1/pokemon &ensp;** `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Get all pokemon
- **/v1/pokemon/{name} &ensp;** `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Get pokemon by name
- **/v1/pokemon/{name} &ensp;** `=>`  &ensp; **PUT** &ensp;  `=>` &ensp; Update pokemon
- **/v1/pokemon/{name} &ensp;** `=>`  &ensp; **DELETE** &ensp;  `=>` &ensp; Delete pokemon by name
  
### Other
- **/health &ensp;**  `=>` &ensp; **GET** &ensp; `=>` &ensp; Ping the database connection
- **/auth/sign-up &ensp;**  `=>` &ensp; **POS** &ensp; `=>` &ensp; Create new pokemon trainer
- **/auth/sign-in &ensp;** `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Sign in with existing profile to generate JWT authentification token

## Tools used

- `Routing`: [Chi](https://github.com/go-chi/chi)
- `Database`: Postgres + [pgx](https://github.com/jackc/pgx/)
- `Containerization`: [Docker](http://docker.com/) + Docker Compose
- `Authentification and middleware`: [JWT Go](https://github.com/golang-jwt/jwt)
