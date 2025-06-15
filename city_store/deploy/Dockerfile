FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/citystore_backend ./city_store/cmd

FROM alpine:latest AS run
WORKDIR /app
COPY --from=build /app/citystore_backend .
COPY city_store/deploy/config.yaml .
EXPOSE 8080
CMD ["./citystore_backend"]