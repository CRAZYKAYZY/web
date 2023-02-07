FROM golang:1.19.4-buster AS build

WORKDIR /app

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN go get -u && go mod tidy

RUN go build -o /web

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /web /web

EXPOSE 8080

ENTRYPOINT ["/web"]