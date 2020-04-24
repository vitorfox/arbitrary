FROM golang:1.12-alpine AS build
RUN apk add --no-cache git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o arbitrary ./cmd/http-server/main.go
RUN ls
RUN ls /app

FROM scratch
COPY --from=build /app/arbitrary .
CMD ["./arbitrary"]