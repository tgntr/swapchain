# Install dependencies and build binary
FROM golang:1.19-alpine AS swapchain-builder

WORKDIR /swapchain
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o=build/swapchaind ./cmd/swapchaind

# Run binary
FROM alpine:3.16 AS swapchain-runner

COPY --from=swapchain-builder /swapchain/build/swapchaind /bin/swapchaind

ENV HOME /swapchain
WORKDIR $HOME

EXPOSE 26656 26657 1317 9090

ENTRYPOINT ["swapchaind"]
