# syntax=docker/dockerfile:1

FROM golang:1.22.0 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /nutri-plans ./cmd/server/main.go

FROM gcr.io/distroless/base-debian11 AS build-release

WORKDIR /

COPY --from=build /nutri-plans .
COPY --from=build /app/storages/ storages/
COPY --from=build /app/static/ static/
COPY --from=build /app/docs/openapi.yml docs/openapi.yml

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/nutri-plans" ]