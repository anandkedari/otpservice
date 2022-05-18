FROM golang:latest

RUN mkdir /build
WORKDIR /build
RUN chmod -R 777 /build/otpservice

RUN export GO111MODULE=on
RUN go get github.com/anandkedari/otpservice
RUN git clone https://github.com/anandkedari/otpservice.git

RUN cd /build/otpservice && go build

EXPOSE 8080

ENTRYPOINT ["/build/otpservice"]