FROM golang:1.17.1 as development
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Install Reflex for development
RUN go get github.com/cosmtrek/air
# Expose port
EXPOSE 8000
# Start app
CMD air