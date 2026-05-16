FROM golang:1.26.3

# Set the working directory inside the container
# All commands below run from here
WORKDIR /app 


# Install Air globally inside the container
# Air binary goes to $GOPATH/bin which is already in PATH
RUN go installl github.com/air-verse/air@latest


# Copy only go.mod and go.sum first
# why ? Docker layer caching - if these don't change,
# it won't re-run go mod download on every build

COPY go.mod go.sum ./
RUN go mod download

# Notice : we do NOT do "COPY . ." here
# Because the bind mount will provide the source files at runtime
# If we copied here, bind mount wold override it anyway

# Tell Docker this container willl listen on port 8080
EXPOSE 8080

# When container starts, run Air (not your app directory)
# Air will compile and run your app, and watch for changes
CMD ["go/bin/air"]