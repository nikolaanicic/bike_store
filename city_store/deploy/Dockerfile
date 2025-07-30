FROM golang:1.23.11 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/citystore_backend ./city_store/cmd

FROM alpine:latest AS run

WORKDIR /app
COPY --from=build /app/citystore_backend .
COPY --from=build /app/city_store/data/migrations/1_create_bikes.up.sql /app/migrations/1_create_bikes.up.sql
COPY --from=build /app/city_store/data/migrations/2_create_rentals.up.sql /app/migrations/2_create_rentals.up.sql


EXPOSE 8080
ENTRYPOINT ["./citystore_backend"]