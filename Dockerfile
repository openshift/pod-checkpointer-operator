FROM openshift/origin-release:golang-1.10
COPY . /go/src/github.com/openshift/pod-checkpointer-operator
RUN cd /go/src/github.com/openshift/pod-checkpointer-operator && make build

FROM centos:7
COPY manifests /manifests

LABEL io.openshift.release.operator true
