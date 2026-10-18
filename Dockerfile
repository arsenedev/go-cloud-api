
FROM golang:1.25-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY main.go main_test.go ./

RUN CGO_ENABLED=0 go test ./...
RUN CGO_ENABLED=0 go build -trimpath -o /api main.go

FROM scratch

COPY --from=build /api /api

EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/api"]
