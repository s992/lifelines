FROM node:18.18.0-alpine AS node-build
WORKDIR /etc/lifelines
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
WORKDIR /etc/lifelines
COPY go.mod go.sum main.go ./
COPY internal ./internal/
COPY sql ./sql/
COPY --from=node-build /etc/lifelines/dist ./dist
RUN go build -o ./dist/lifelines

FROM debian:12
ENV LOGGER_DB_DIR=/var/lifelines/
ENV LOGGER_PORT=80
EXPOSE 80
RUN mkdir -p /var/lifelines
COPY --from=go-build /etc/lifelines/dist/lifelines /usr/bin/lifelines
COPY .env /
CMD ["/usr/bin/lifelines"]
