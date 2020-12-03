
# Sttart from base image 1.12.13:
FROM golang:1.12.13

ENV MYSQL_HOSTS=localhost:8080
ENV LOG_LEVEL=info

# Configure the repo url so we can configure our work directory:
ENV REPO_URL=https://github.com/TimKotowski/bookstore_users-api

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/go/$REPO_URL

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o users-api .

# Expose port 8081 to the world:
EXPOSE 8081

CMD ["./users-api"]
