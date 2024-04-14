FROM  golang:1.21.9 as builder

WORKDIR /project/capstone

COPY go.mod ./
COPY go.sum ./

RUN go mod download

# fix "gcc" not found
RUN apt-get install gcc

COPY . .
RUN CGO_ENABLED=1 go build -tags http -o /project/capstone/build/capstone .

FROM alpine:latest

# to fix timezone not loaded
RUN apt-get install -y tzdata


COPY --from=builder /project/capstone/build/capstone /project/capstone/build/capstone

WORKDIR /project/capstone/build/

EXPOSE 8080
CMD [ "./capstone", "http" ]
