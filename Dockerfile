FROM golang AS build

WORKDIR /aileg

# Copy go.mod and go.sum files to the workspace separately and download dependecies.
# Doing this separately will cache these as its own separate layer
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# Copy the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 go build -o aileg-server.bin main.go

# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# Then we copy and run it from a slim image
FROM alpine
WORKDIR /aileg

COPY --from=build /aileg/aileg-server.bin .

EXPOSE 8081

ENTRYPOINT ["/aileg/aileg-server.bin"]