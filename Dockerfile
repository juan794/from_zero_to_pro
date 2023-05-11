# Build SPA Stage
FROM node:14 AS spa-builder
WORKDIR /app
COPY ./spa/package*.json ./
RUN npm install
COPY ./spa .
RUN npm run build

# Build Golang App Stage
FROM golang:1.20 AS go-builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY main.go ./
# Copy the built SPA files from the previous stage
COPY --from=spa-builder /app/dist ./public
RUN go build -o server

# Final Stage
FROM debian:buster-slim
COPY --from=go-builder /app/server /server
COPY --from=go-builder /app/public /public
CMD ["/server"]