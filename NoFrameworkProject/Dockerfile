FROM golang:latest

RUN mkdir -p /user/app
WORKDIR /user/app
RUN go get -u github.com/lib/pq

COPY . .
CMD ["go", "run", "Main.go"]