FROM golang:1.12.3-alpine AS build
RUN apk add -u --no-cache bash curl git
WORKDIR /go/src/github.com/obukhov/pantomime/server
COPY ./go.* ./
RUN export GO111MODULE=on && go mod tidy
COPY . .
RUN export GO111MODULE=on && go build -o /go/bin/server

EXPOSE 80
CMD ["/go/src/github.com/obukhov/pantomime/server"]

FROM alpine AS deploy
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /go/bin/api-gateway /app/
COPY --from=build /go/src/github.com/obukhov/teamspirit/api-gateway/schema.graphql /app/
EXPOSE 80

CMD ["/app/api-gateway"]
