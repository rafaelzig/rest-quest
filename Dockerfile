FROM golang:1.15.5-alpine AS build
LABEL maintainer="rafaeldasilvacosta@hotmail.co.uk"
WORKDIR /usr/src

COPY go.* ./
RUN go mod download

COPY internal internal
COPY cmd cmd

# Statically compile the binary (resulting binary will not be linked to any C libraries)
ENV CGO_ENABLED=0
RUN go build -o /usr/bin cmd/quest/quest.go

FROM scratch
COPY --chown=1001:1001 --from=build /usr/bin /usr/bin
USER 1001
ENV SERVER_PORT 8080
ENTRYPOINT ["quest"]