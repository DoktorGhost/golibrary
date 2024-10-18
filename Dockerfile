FROM golang:1.22-alpine as app-builder
RUN apk update && apk add curl make git

WORKDIR /src
COPY . .

RUN go mod download
RUN go build ./cmd/app

FROM alpine:latest
RUN apk update && apk add --no-cache curl
WORKDIR /src
COPY --from=app-builder /src/app .
COPY --from=app-builder /src/.env .
COPY --from=app-builder /src/migrations migrations

CMD ["./app"]