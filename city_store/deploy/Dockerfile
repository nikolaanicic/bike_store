FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/citystore_backend ./city_store/cmd

FROM alpine:latest AS run
RUN apk add --no-cache curl bash mysql-client

WORKDIR /app
COPY --from=build /app/citystore_backend .
COPY --from=build /app/city_store/data/migrations/1_init.up.sql .
COPY city_store/deploy/run.sh .
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz \
  | tar xz && mv migrate /usr/local/bin/migrate

RUN chmod +x run.sh
EXPOSE 8080
ENTRYPOINT ["./run.sh"]