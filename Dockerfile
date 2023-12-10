# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY . .

# Install any Go dependencies if needed
# RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port that your Go application will listen on
EXPOSE 8080

# Set environment variables for PostgreSQL
ENV POSTGRES_HOST=postgres
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=Admin2030
ENV POSTGRES_DB=api_users

# Run the Go application when the container starts
CMD ["./main"]
