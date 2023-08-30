# Use a lightweight Python image to serve HTML content
FROM python:3.9-alpine AS html-server


# Copy the HTML file from your local machine to the container
COPY index.html .

# Expose port 8000 for the HTML server
EXPOSE 8000


# Use the official Go image as the base image for Go application
FROM golang:1.16-alpine AS go-app

# Copy the Go files from your local machine to the container
COPY hello.go .
COPY go.mod .
COPY go.sum .

# Build the Go application
RUN go build
