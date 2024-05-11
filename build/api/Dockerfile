# Start from golang base image
FROM golang:1.20

# Add Maintainer info
LABEL maintainer="Vishal Anarase"

# Set destination for COPY
WORKDIR /app

# Copy code
COPY . .

# Download all the dependencies
RUN go mod tidy

# Build
RUN CGO_ENABLE=0 GOOS=linux go build -o bin/app cmd/api/main.go

# Expose port
EXPOSE 8080

# Run
CMD ["bin/app"]