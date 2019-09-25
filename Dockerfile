FROM golang:1.13-alpine

# git
RUN apk add --no-cache git

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR /app
RUN mkdir -p $APP_DIR

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./ithome-iron-beego)
ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 go build -ldflags '-d -w -s'

EXPOSE 8080
