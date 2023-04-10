FROM golang:alpine

RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .

EXPOSE  8090
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]
