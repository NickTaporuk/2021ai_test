FROM ubuntu:latest

WORKDIR /usr/local/bin
COPY ./scalc scalc
COPY ./testdata ./testdata

CMD ["/usr/local/bin/scalc", "[ SUM [ DIF ./testdata/a.txt ./testdata/b.txt ./testdata/c.txt ] [ INT ./testdata/b.txt ./testdata/c.txt ] ]"]