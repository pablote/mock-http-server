## Build
FROM golang:1.19-bullseye AS build

WORKDIR /app

COPY go.mod ./
#COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /server

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /server /server

EXPOSE 9090

USER nonroot:nonroot

ENTRYPOINT ["/server"]