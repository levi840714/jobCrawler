FROM golang:1.14-alpine
ARG keyword
ENV keyword=$keyword
WORKDIR /jobCrawler
ADD . /jobCrawler
RUN cd /jobCrawler \
    && cp .env.example .env \
    && go build
ENTRYPOINT ["sh","-c","./jobCrawler -keyword ${keyword}"]
