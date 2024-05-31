# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.22-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /go-ethereum/
COPY go.sum /go-ethereum/
RUN cd /go-ethereum && go mod download

ADD . /go-ethereum
RUN cd /go-ethereum && go run build/ci.go install -static ./cmd/geth

# Pull Geth into a second stage deploy alpine container
# Create a non-root user and set up directory permissions
FROM alpine:latest

# Add a new user to avoid running as root
RUN adduser -D -H geth

# Set the working directory and create data directory
WORKDIR /geth
RUN mkdir -p /geth/ && chown -R geth /geth

# Copy the Geth binary from the builder stage
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/

# Expose the necessary ports
EXPOSE 8545 8546 30303 30303/udp

# Switch to the non-root user
USER geth

# Define the entrypoint
ENTRYPOINT ["geth", "--datadir=/geth/"]

# Add some metadata labels to help programmatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"