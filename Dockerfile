FROM golang:alpine

ENV DB_PROVIDER = mysql\
    DATABASE_URL = root:1234@tcp(localhost:3306)\
    DELAY_MS = 200\
    URL_LENGTH = 15\
    MAX_BID = 999999\
    MIN_BID = 100000\
    SOCKET = localhost:3001

WORKDIR /auction_system

COPY . .

RUN go build -o main .

WORKDIR /auction_binary

RUN cp /auction_system/main .

EXPOSE 3001

CMD ["/auction_binary/main"]