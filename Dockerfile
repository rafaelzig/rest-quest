FROM golang:1.15.5-alpine AS build
LABEL maintainer="rafaeldasilvacosta@hotmail.co.uk"
WORKDIR /usr/src

COPY go.* ./
RUN go mod download

COPY internal internal
COPY cmd cmd

# Statically compile the binary (resulting binary will not be linked to any C libraries)
ENV CGO_ENABLED=0
RUN go build -o /bin cmd/quest/quest.go \
&&  install -d -m 0744 -o 1001 -g 1001 /var/lib/quest

FROM scratch
COPY --chown=1001:1001 --from=build /bin /bin
COPY --chown=1001:1001 --from=build /var/lib/quest /var/lib/quest
USER 1001
WORKDIR /var/lib/quest
ENV SERVER_PORT 8080
ENTRYPOINT ["quest"]