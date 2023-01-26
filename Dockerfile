FROM golang:1.19.4-buster AS build

WORKDIR /app
# Install necessary packages
# RUN apt-get update && apt-get install -y curl git tar

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# RUN curl -L https://github.com/kyleconroy/sqlc/releases/download/v1.16.0/sqlc_1.16.0_linux_amd64.tar.gz -o sqlc.tar.gz \
#   && tar -xzf sqlc.tar.gz \
#   && mv sqlc /usr/local/bin/ \
#   && rm sqlc.tar.gz

# RUN chmod +x /usr/local/bin/sqlc

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

#COPY --from=build /usr/local/bin/sqlc /usr/local/bin/sqlc

EXPOSE 8080

ENTRYPOINT ["/web"]