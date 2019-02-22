# base Go image version.
FROM golang:1.11.4-stretch AS build

WORKDIR /project

# install dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG version
RUN go build -o mesg-core \
      -ldflags="-X 'github.com/mesg-foundation/core/version.Version=$version'" \
      core/main.go

# NOTE: docker uses xz-utils for uncompress tar.xz files
FROM ubuntu:18.04
RUN apt-get update && \
      apt-get install -y --no-install-recommends ca-certificates=20180409 xz-utils && \
      apt-get clean && \
      rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=build /project/mesg-core .
CMD ["./mesg-core"]
