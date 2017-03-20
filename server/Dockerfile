# # Start from a Debian image with the latest version of Go installed
# # and a workspace (GOPATH) configured at /go.
# FROM golang:1.8

# # Copy the local package files to the container's workspace.
# ADD . /go/src/github.com/navkar/user_api

# # Build the outyet command inside the container.
# # (You may fetch or manage dependencies here,
# # either manually or with a tool like "godep".)
# RUN go get github.com/gin-gonic/gin
# RUN go get github.com/jinzhu/gorm
# RUN go get github.com/lib/pq

# RUN go install github.com/navkar/user_api

# # Run the outyet command by default when the container starts.
# ENTRYPOINT /go/bin/user_api

# # Document that the service listens on port 8080.
# EXPOSE 8080

# Refer - https://blog.golang.org/docker

FROM golang:onbuild
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/lib/pq
RUN go build main.go
EXPOSE 8080