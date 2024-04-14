FROM  golang:1.21.9-alpine as builder

WORKDIR /project/capstone

COPY go.mod ./
COPY go.sum ./

RUN go mod download

# fix "gcc" not found
RUN apk add build-base

COPY . .
RUN CGO_ENABLED=1 go build -tags http -o /project/capstone/build/capstone .

FROM alpine:latest

# to fix timezone not loaded
RUN apk add --no-cache tzdata

# fix "gcc" not found
RUN apk add build-base

COPY --from=builder /project/capstone/build/capstone /project/capstone/build/capstone

WORKDIR /project/capstone/build/

EXPOSE 8080
CMD [ "./capstone", "http" ]
