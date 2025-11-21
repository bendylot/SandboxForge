# build
FROM golang:1.25.4 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /api .

# run
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=build /api /app/api
COPY .env /app/.env
EXPOSE 3001
ENTRYPOINT ["/app/api"]
