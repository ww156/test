FROM golang
COPY main.go /app/
WORKDIR /app
EXPOSE 8080
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
CMD ["go", "run", "main.go"]
