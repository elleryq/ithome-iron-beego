#
# base
#
FROM alpine:3.10 AS base
RUN apk add --no-cache curl wget

#
# builder
#
FROM golang:1.13-alpine3.10 AS builder

# git
RUN apk add --no-cache git

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR /app
RUN mkdir -p $APP_DIR

# Set the entrypoint
ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 go build -ldflags '-d -w -s'

#
# final
#
FROM base AS final

ENV APP_DIR /app
RUN mkdir -p $APP_DIR
COPY --from=builder /app/ithome-iron-beego $APP_DIR/ithome-iron-beego
COPY --from=builder /app/conf $APP_DIR/conf
COPY --from=builder /app/views $APP_DIR/views

ENTRYPOINT (cd $APP_DIR && ./ithome-iron-beego)
EXPOSE 8080
