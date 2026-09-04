FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o music-api .

FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/music-api .
EXPOSE 8080
CMD ["./music-api"]
