FROM golang:1.6

# Add directory for app source code, make it the working dir.
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# Copy files from current host directory into container.
COPY . /go/src/app

# Install codegangsta/gin to get auto-reload during development
# (not to be confused with the gin-gonic/gin web framework).
# Other dependencies imported by app code will also be installed here
# during build (do not add extra `go get` commands here).
RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install

# PORT to be used by app running in container.
ENV PORT 8080

# Port 3000 exposed to host, it's how to get to the gin proxy.
EXPOSE 3000

# Have gin run my app so it'll handle auto-reload during dev.
CMD gin run
