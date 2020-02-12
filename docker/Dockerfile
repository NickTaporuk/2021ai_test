FROM ubuntu:latest

WORKDIR /usr/local/bin
COPY ./scalc scalc
COPY ./testdata ./testdata

RUN ls -la /usr/local/bin
RUN ls -la /usr/local/bin/testdata

CMD ["/usr/local/bin/scalc", "[ SUM [ DIF ./testdata/a.txt ./testdata/b.txt ./testdata/c.txt ] [ INT ./testdata/b.txt ./testdata/c.txt ] ]"]
#CMD ["/bin/bash"]