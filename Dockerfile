FROM golang:1.14-alpine
ARG keyword
ENV keyword=$keyword
WORKDIR /jobCrawler
ADD . /jobCrawler
RUN cd /jobCrawler \
    && go build
ENTRYPOINT ["sh","-c","./jobCrawler -keyword ${keyword}"]
