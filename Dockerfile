# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:latest


# Add Maintainer Info
LABEL maintainer="Abish"

# Set the Current Working Directory inside the container
WORKDIR /Finleap


# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed

# RUN go mod download

ENV   MYSQL_USER  docker
ENV   MYSQL_PASSWORD  docker

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go mod download

RUN go test ./... -v --cover

RUN rm -r /Finleap/mysql_data/*


# Build the Go app
RUN go build -o Finleap .

# Expose port 8080 to the outside world
EXPOSE 8080 

# Run the executable
CMD ["./Finleap"]