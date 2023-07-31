#!/bin/bash

REPOSITORY=9ilbert
SERVICE=member-go
VERSION=0.0

docker buildx rm multiarch-builder
docker buildx create --name multiarch-builder default --use
docker buildx build --platform linux/amd64 -t $REPOSITORY/$SERVICE:amd-$VERSION . --push
docker buildx build --platform linux/arm64 -t $REPOSITORY/$SERVICE:arm-$VERSION . --push
docker buildx build --platform linux/amd64,linux/arm64 -t $REPOSITORY/$SERVICE:multiarch-$VERSION . --push