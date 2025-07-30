FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/central_store_backend ./central_store/cmd

FROM alpine:latest AS run
WORKDIR /app
COPY --from=build /app/central_store_backend .
COPY --from=build /app/central_store/data/migrations/1_create_bikes.up.sql /app/migrations/1_init.up.sql

EXPOSE 5000
CMD ["./central_store_backend"]