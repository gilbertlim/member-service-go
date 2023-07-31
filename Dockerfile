FROM --platform=$BUILDPLATFORM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=$TARGETOS \
    GOARCH=$TARGETARCH

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main .

WORKDIR /dist
RUN cp /build/main .
RUN cp /build/*.yaml .

FROM scratch
COPY --from=builder /dist .
ENTRYPOINT ["/main"]