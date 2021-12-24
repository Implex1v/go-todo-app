FROM golang:1.17.5-bullseye as base
WORKDIR /app

FROM base as build

COPY . ./

RUN go test -v ./...
RUN go build

FROM base as release
ENV PORT 8080
ENV GIN_MODE release

COPY --from=build /app/api-server ./

CMD ["./api-server"]