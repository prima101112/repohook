FROM alpine:latest

RUN apk add git curl
COPY repohook /

CMD ["repohook"]
