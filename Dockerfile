FROM openshift/origin-release:golang-1.10 AS builder
COPY . /go/src/github.com/openshift/pod-checkpointer-operator
RUN cd /go/src/github.com/openshift/pod-checkpointer-operator && make build

FROM centos:7
COPY manifests /manifests

COPY --from=builder /go/src/github.com/openshift/pod-checkpointer-operator/pod-checkpointer-operator /usr/bin/pod-checkpointer-operator

LABEL io.openshift.release.operator true
