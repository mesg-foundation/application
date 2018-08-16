FROM golang:1.10.3-stretch AS build
WORKDIR src/github.com/mesg-foundation/core
COPY . .

ARG version
RUN go build -o mesg-core \
      -ldflags="-X 'github.com/mesg-foundation/core/version.Version=$version'" \
      core/main.go

FROM ubuntu:18.04
WORKDIR /app
COPY --from=build /go/src/github.com/mesg-foundation/core/mesg-core .
CMD ["./mesg-core"]
