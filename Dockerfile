FROM golang:1.14.3-alpine

# Set the working dir
WORKDIR /app/onhub-web

COPY . .

RUN ls -la

# Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download

CMD [ "go", "run", "main.go" ]