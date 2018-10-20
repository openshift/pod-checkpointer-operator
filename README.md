# Pod Checkpointer Operator

The Pod Checkpointer Operator manages the Checkpointer running on master
Kubernetes nodes. The checkpointer writes static pod manifests to be able to
recover the Kubernetes control plane (kube-apiserver, kube-controller-manager, kube-scheduler, and the checkpointer itself).

## Development

See [HACKING](HACKING.md) for development topics.