# Build the go application from a base Image
FROM golang:1.22-alpline as base
# Set the working directory inside the container
WORKDIR /app
# Copy source file to the container
COPY  go.mod go.sum ./
# Doenload all dependency 
RUN go mod download
# Copy the source code into the container
COPY . .

RUN go build -o /app/main ./cmd/main.go

FROM alpline:latest

WORKDIR /root/


COPY --from=base /app/main .

EXPOSE 4000

CMD ["./main"]
