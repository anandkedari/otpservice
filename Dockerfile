FROM golang:latest

WORKDIR /go/src

RUN export GO111MODULE=on
RUN git clone https://github.com/anandkedari/otpservice.git
RUN ls
RUN cd otpservice && go build

EXPOSE 8000

ENTRYPOINT ["./otpservice/otpservice"]