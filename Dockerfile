FROM golang:1.19.0

WORKDIR /usr/src/app

# add hot reload
RUN go install github.com/air-verse/air@latest

# copy all files from host to container 
COPY . .
# package properly install
RUN go mod tidy

