FROM golang:latest

ADD . /src
WORKDIR /src
VOLUME [ "/data" ]
EXPOSE 8080
RUN mkdir -p /bin && \
     go env -w GOPROXY=https://goproxy.cn,direct && \
     go build -o /bin/todolist.exe main.go
ENTRYPOINT [ "/bin/todolist.exe" ]
