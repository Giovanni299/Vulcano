FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Giovanni Sanabria <giovanni299@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

RUN go get github.com/labstack/echo
RUN go get github.com/joho/godotenv
RUN go get github.com/lib/pq
RUN go get github.com/swaggo/echo-swagger
RUN go get github.com/Giovanni299/Vulcano/database
RUN go get github.com/Giovanni299/Vulcano/docs

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8084

# Command to run the executable
CMD ["./main"]