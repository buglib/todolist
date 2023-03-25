FROM golang:latest

ADD . /src
WORKDIR /src
VOLUME [ "/data" ]
EXPOSE 8080
RUN mkdir -p /bin && \
     go env -w GOPROXY=https://goproxy.cn,direct && \
     go build -o /bin/todolist.exe main.go
# CMD [ "/bin/todolist.exe", \
#      "-host", "web", \
#      "-port", "8080", \
#      "-mysqlHost", "db", \
#      "-mysqlPort", "3306", \
#      "-userName", "buglib", \
#      "-passwd", "123456", \
#      "-db", "todolist"]
# CMD ["/bin/todolist.exe", "-mysqlHost", "db"]
# CMD /bin/todolist.exe
ENTRYPOINT [ "/bin/todolist.exe" ]
