FROM golang:1.12.5-alpine

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
RUN apk add --no-cache ca-certificates git

WORKDIR /go/src/app

# Import the code from the context.
COPY . .

RUN go get -d -v ./...

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build -v -installsuffix 'static' -o /app/hello

EXPOSE 8080

USER nobody:nobody

ENTRYPOINT ["/app/hello"]