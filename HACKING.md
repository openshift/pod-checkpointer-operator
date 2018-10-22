# Pod Checkpointer Operator Hacking

## Local development

It's possible (and useful) to develop the operator locally targeting a remote cluster.

### Prerequisites

* An OpenShift or Kubernetes cluster with at least a master, infra, and compute
node. An admin-scoped * `KUBECONFIG` for the cluster. The [operator-sdk](https://github.com/operator-framework/operator-sdk).

### Building

To build the operator during development, use the standard Go toolchain:

```
$ make build
```

### Container image

To build the operator docker image:

##### Building using operator-sdk

```
operator-sdk build quay.io/[some_username]/pod-checkpointer-operator
docker push quay.io/[some_username]/pod-checkpointer-operator
```

### Running

To run the operator, first deploy the custom resource definitions:

```
oc apply -f manifests
```

#### Running with Operator SDK

Use the operator-sdk to launch the operator:

```
$ operator-sdk up local namespace kube-system --kubeconfig=$KUBECONFIG
```

#### Running a Containerized Operator

Edit manifests/04-deployment.yaml to point the operator to your own image.

```
oc create -f manifests/04-deployment.yaml
```

*In order to run as a container you will need to upload the Pod Checkpointer Operator image to your registry*
