FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o pokemonapi cmd/main.go

ENTRYPOINT [ "./pokemonapi" ]

