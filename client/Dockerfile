FROM golang
COPY . /go/src/client
WORKDIR /go/src/client
ENV GIT_TERMINAL_PROMPT=1
RUN git config --global user.email pankaj.razdan@exium.net
RUN git config --global user.name prazdan79
RUN go get . 
ENTRYPOINT go run main.go
EXPOSE 8080
