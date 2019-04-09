FROM golang:latest
WORKDIR /app
COPY . /app
RUN cd server && go build
CMD ["./server/server"]