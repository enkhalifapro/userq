# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

# install git
RUN apk add --update --no-cache git
############ Copy the local package files to the container's workspace.
ADD . /go/src/gitlab.com/enkhalifapro/reppy-api
WORKDIR /go/src/gitlab.com/enkhalifapro/reppy-api
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build
ENTRYPOINT ./reppy-api run --cfg-name $ENVNAME
# CMD ["./reppy-api run"]
EXPOSE 3000