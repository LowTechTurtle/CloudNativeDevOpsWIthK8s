FROM golang:1.17-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/demo

# multistage for the small binary size( doesnt need golang and alpine,
# just need to run the binary)

# also reduce program in our image, reduce attack surface, less security risk
FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]