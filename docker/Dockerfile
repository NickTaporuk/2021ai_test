FROM ubuntu:latest

WORKDIR /usr/local/bin
COPY . 2021ai_test
COPY . testdata

RUN ls -la /go/src/2021ai && chmod +x /go/src/2021ai/2021ai_test
RUN ls -la /go/src/2021ai/testdata

CMD ["/usr/local/bin/2021ai_test", "[ SUM [ DIF ./testdata/a.txt /testdata/b.txt /testdata/c.txt ] [ INT /testdata/b.txt /testdata/c.txt ] ]"]