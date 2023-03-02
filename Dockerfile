FROM golang:1.18-alpine

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -v -o /bin/armactl cmd/armactl/main.go

LABEL author="Britton Hayes"
LABEL github="https://github.com/brittonhayes/armactl"

ENTRYPOINT ["/bin/armactl"]

CMD [ "/bin/armactl" ]