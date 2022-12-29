## Build
FROM golang:1.19.4-buster AS build

WORKDIR /app

# Add Maintainer info
LABEL maintainer="crazykayzy"

#RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

COPY .  .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN go mod tidy

RUN go build -o /web

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /web /web

EXPOSE 8080

ENTRYPOINT ["/web"]