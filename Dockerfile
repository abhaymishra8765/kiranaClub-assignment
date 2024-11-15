
FROM golang:1.23-alpine AS build


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go build -o main .


FROM gcr.io/distroless/base-debian11


WORKDIR /app


COPY --from=build /app/main .


CMD ["./main"]
