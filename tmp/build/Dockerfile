FROM alpine:3.6

RUN adduser -D control-plane-operator
USER control-plane-operator

ADD tmp/_output/bin/control-plane-operator /usr/local/bin/control-plane-operator
