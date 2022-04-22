FROM golang:1.18 as base
WORKDIR /app

FROM base as build

COPY . ./

RUN go test -v ./...
RUN go build

FROM base as release
ARG BUILD_VERSION=0.1.0
ARG GIT_COMMIT=unknown

LABEL org.opencontainers.image.version="${BUILD_VERSION}"\
      org.opencontainers.image.revision="${GIT_COMMIT}"

ENV PORT 8080
ENV GIN_MODE release
ENV VERSION ${BUILD_VERSION}

COPY --from=build /app/api-server ./

CMD ["./api-server"]