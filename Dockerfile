FROM golang:1.19.5-alpine3.17 AS build
WORKDIR /usr/local/src/toggle
COPY . .
# Install Taskfile to reuse build commands
RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN task setup
RUN task build

# NOTE(mlee): We want to use `scratch` as the base image, but there was an issue
# with detecting the binary.
# TODO(mlee): Convert to use scratch when I get some time to save a few bytes
FROM alpine:3.17 as run
WORKDIR /opt/bin
COPY --from=build /usr/local/src/toggle/bin/toggle .
CMD ["/opt/bin/toggle", "server"]