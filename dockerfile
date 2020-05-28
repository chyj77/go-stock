FROM golang

# RUN yum -y update

# 安装wget
# RUN yum -y install golang

#安装git
# RUN yum -y install git

#安装go
# RUN wget -P /home https://studygolang.com/dl/golang/go1.11.linux-amd64.tar.gz &&  tar -C /usr/local/ -zxf /home/go1.11.linux-amd64.tar.gz && rm -rf /home/go1.11.linux-amd64.tar.gz
ENV GOROOT /usr/local/go
# RUN mkdir -p /root/gopath
# config GOPATH
RUN mkdir -p /root/gopath
RUN mkdir -p /root/gopath/src
RUN mkdir -p /root/gopath/pkg
RUN mkdir -p /root/gopath/bin
ENV GOPATH /root/gopath
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin 

#安装beego
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get github.com/astaxie/beego/orm
RUN go get github.com/go-sql-driver/mysql
#  RUN mkdir /app 
 ADD . $GOPATH/src/go-stock 
 WORKDIR $GOPATH/src/go-stock 
EXPOSE 8080
entrypoint bee run