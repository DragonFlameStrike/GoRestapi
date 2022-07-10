FROM golang:1.18-alpine


# Set the Current Working Directory inside the container
WORKDIR /restapimain

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app

RUN go build -o /apiserver ./cmd/apiserver/main.go


# This container exposes port 8080 to the outside world
EXPOSE 8080

RUN pwd
# Run the binary program produced by `go install`
CMD ["pwd"]

CMD ["/apiserver"]