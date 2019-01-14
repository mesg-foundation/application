FROM golang:1.11.4-stretch AS build
WORKDIR /core
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG version
RUN go build -o mesg-core \
      -ldflags="-X 'github.com/mesg-foundation/core/version.Version=$version'" \
      core/main.go

FROM ubuntu:18.04
RUN apt-get update && \
      apt-get install -y --no-install-recommends ca-certificates=20180409 && \
      apt-get clean && \
      rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=build /core/mesg-core .
CMD ["./mesg-core"]
