# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.21 as builder

RUN apt-get update && apt-get install -y --no-install-recommends gcc musl-dev git

# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /go-ethereum/
COPY go.sum /go-ethereum/
RUN cd /go-ethereum && go mod download

ADD . /go-ethereum
RUN cd /go-ethereum && go run build/ci.go install -static ./cmd/geth

# Pull Geth into a second stage deploy alpine container
FROM ubuntu:latest

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/

# Add new group and user to avoid running as the root user
RUN useradd geth && \
    mkdir /geth && \
    chown -R geth:geth /geth

USER geth
WORKDIR /geth

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["geth", "--datadir=/geth"]

# Add some metadata labels to help programmatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"