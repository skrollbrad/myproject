FROM golang:latest
WORKDIR /app 
#workdir errors gopath
COPY ./ ./

RUN go build -o main .
CMD ["./main"]

# FROM golang:latest

# # Set destination for COPY
# WORKDIR /app

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./

# RUN go build -o main .




# # Run
# CMD ["./main"]