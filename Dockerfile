FROM node:18.18.0-alpine AS node-build
WORKDIR /etc/logger
COPY  package.json \
      package-lock.json \
      postcss.config.cjs \
      vite.config.ts \
      tsconfig.json \
      tsconfig.node.json \
      ./
RUN npm install
COPY client ./client
RUN npm run build

FROM docker.io/golang:1.21 AS go-build
ENV CGO_ENABLED=1
WORKDIR /etc/logger
COPY go.mod go.sum main.go ./
COPY internal ./internal/
COPY sql ./sql/
COPY --from=node-build /etc/logger/dist ./dist
RUN go build -o ./dist/logger

FROM debian:12
ENV LOGGER_DB_DIR=/var/logger/
ENV LOGGER_PORT=80
EXPOSE 80
RUN mkdir -p /var/logger
COPY --from=go-build /etc/logger/dist/logger /usr/bin/logger
COPY .env /
CMD ["/usr/bin/logger"]
