FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/citystore_backend ./city_store/cmd

FROM alpine:latest AS run
ENV MYSQL_VERSION=8.0.36
ENV MYSQL_BASE_URL=https://downloads.mysql.com/archives/get/p/23/file

ENV MYSQL_VERSION=8.0.36
ENV MYSQL_BASE_URL=https://downloads.mysql.com/archives/get/p/23/file

RUN apk add --no-cache bash curl libaio ncurses-terminfo openssl

RUN curl -L ${MYSQL_BASE_URL}/mysql-${MYSQL_VERSION}-linux-glibc2.28-x86_64.tar.xz -o mysql.tar.xz \
    && mkdir -p /opt/mysql \
    && tar -xf mysql.tar.xz --strip-components=1 -C /opt/mysql \
    && rm mysql.tar.xz \
    # Remove everything except the client binaries to save space
    && rm -rf /opt/mysql/bin/mysqld /opt/mysql/bin/mysql_install_db /opt/mysql/bin/mysql_secure_installation /opt/mysql/lib/plugin /opt/mysql/share \
    && ln -s /opt/mysql/bin/mysql /usr/local/bin/mysql \
    && ln -s /opt/mysql/bin/mysqladmin /usr/local/bin/mysqladmin \
    && ln -s /opt/mysql/bin/mysql_config /usr/local/bin/mysql_config \
    && ln -s /opt/mysql/bin/mysqlshow /usr/local/bin/mysqlshow
RUN export PATH="/opt/mysql/bin:$PATH"

WORKDIR /app
COPY --from=build /app/citystore_backend .
COPY --from=build /app/city_store/data/migrations/1_init.up.sql .
COPY city_store/deploy/run.sh .
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz \
  | tar xz && mv migrate /usr/local/bin/migrate

RUN chmod +x run.sh
EXPOSE 8080
ENTRYPOINT ["./run.sh"]