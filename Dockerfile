# parent image
FROM golang:1.15.6-alpine3.12 AS builder

# workspace directory
WORKDIR /app

# copy `go.mod` and `go.sum`
ADD go.mod go.sum ./

# install dependencies
RUN go mod download

# copy source code
COPY . .

# build executable
RUN go build -o ./bin/aquastatistic .

##################################

# parent image
FROM alpine:3.19.3

# workspace directory
WORKDIR /app

# copy binary file from the `builder` stage
COPY --from=builder /app/bin/aquastatistic ./

# set entrypoint
ENTRYPOINT [ "./aquastatistic" ]

