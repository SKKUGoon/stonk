FROM golang:1.21.5-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    COOS=linux \
    GOARCH=amd64

# Go building compile phase
WORKDIR /app

RUN apk add --no-cache git

# Copy custom utility package folder
COPY /kis /app/kis
COPY /binance /app/binance

# Copy other go module file
COPY /wqexec /app/wqexec

# Change the working directory. 
WORKDIR /app/wqexec
RUN go mod download

# Build the file as wqback
RUN go build -o wqback .

# Use minimal alpine image to run the compiled binary
FROM alpine:latest

WORKDIR /root/
RUN apk update
RUN apk upgrade 
RUN apk add curl

# Copy developing environment
COPY --from=builder /app/wqexec/.env /root
COPY --from=builder /app/wqexec/.env.test /root
COPY --from=builder /app/wqexec/wqback /root

CMD [ "./wqback" ]
