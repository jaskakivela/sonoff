FROM alpine:latest
RUN apk add --no-cache busybox-extras wget netcat-openbsd tcpdump
# ENTRYPOINT [ "nc", "-n", "-k", "-l", "8080", "-w", "1" ]
ENTRYPOINT [ "top", "-d", "60" ]