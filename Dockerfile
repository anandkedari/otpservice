FROM golang:latest


RUN export GO111MODULE=on
RUN go get github.com/anandkedari/otpservice
RUN git clone https://github.com/anandkedari/otpservice.git

RUN cd otpservice && go build

EXPOSE 8080

ENTRYPOINT ["otpservice"]