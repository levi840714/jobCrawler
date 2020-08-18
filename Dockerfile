FROM golang:1.14-alpine
WORKDIR /crawlerJob
ADD . /crawlerJob
RUN cd /crawlerJob \
    && cp .env.example .env \
    && go build
ENTRYPOINT ["./crawlerJob"]
