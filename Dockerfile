FROM golang:1.14-alpine
WORKDIR /jobCrawler
ADD . /jobCrawler
RUN cd /jobCrawler \
    && cp .env.example .env \
    && go build
ENTRYPOINT ["./jobCrawler"]
