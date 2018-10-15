FROM alpine:3.6

RUN adduser -D pod-checkpointer-operator
USER pod-checkpointer-operator

ADD tmp/_output/bin/pod-checkpointer-operator /usr/local/bin/pod-checkpointer-operator
