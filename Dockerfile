# Create Builder Image, to compile the source code into an executable
FROM golang:1.20 AS build

WORKDIR /app

COPY . .
# COPY . /app

# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o cloudrun cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCCH=amd64 go build -o cloudrun

# Create the final image, running the API and exposing port 8081
# FROM alpine:latest
FROM scratch

WORKDIR /app

# COPY --from=build /app/cloudrun ./
COPY --from=build /app/cloudrun .

# EXPOSE 8081

ENTRYPOINT [ "./cloudrun" ]
# CMD [ "./cloudrun" ]