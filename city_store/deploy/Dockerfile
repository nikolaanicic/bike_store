FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN apk add --no-cache curl bash mysql-client
RUN chmod +x city_store/deploy/run_migrations_city.sh
RUN city_store/deploy/run_migrations_city.sh

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/citystore_backend ./city_store/cmd

FROM alpine:latest AS run
WORKDIR /app
COPY --from=build /app/citystore_backend .
EXPOSE 8080
CMD ["./citystore_backend"]



RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz \
  | tar xz && mv migrate /usr/local/bin/migrate

COPY ../city_store/data/migrations/1_init.up.sql city_store/data/migrations/1_init.up.sql
COPY ./k8s/run_migrations_city.sh run_migrations_city.sh

RUN chmod +x run_migrations_city.sh

ENTRYPOINT ["./run_migrations_city.sh"]