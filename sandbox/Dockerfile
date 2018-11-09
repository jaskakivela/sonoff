FROM alpine:latest
RUN apk add --no-cache busybox-extras wget netcat-openbsd
ENTRYPOINT [ "top", "-d", "60" ]