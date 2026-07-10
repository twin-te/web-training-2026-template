FROM golang:1.26 AS base
WORKDIR /app

FROM base AS dev
RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 3000
CMD ["air"]

FROM base AS build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /server .

FROM gcr.io/distroless/static-debian12 AS prod
COPY --from=build /server /server
EXPOSE 3000
CMD ["/server"]
