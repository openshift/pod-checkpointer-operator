package manifests

import (
	"bytes"
	"io"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"k8s.io/apimachinery/pkg/util/yaml"
)

const (
	PCOClusterRole    = "assets/pod-checkpointer-operator/cluster-role.yaml"
	PCOCustomResource = "assets/pod-checkpointer-operator/custom-resource.yaml"
	PCONamespace      = "assets/pod-checkpointer-operator/namespace.yaml"
	PCOOperator       = "assets/pod-checkpointer-operator/operator.yaml"
	PCORbac           = "assets/pod-checkpointer-operator/rbac.yaml"
	PCORoleBinding    = "assets/pod-checkpointer-operator/role-binding.yaml"
	PCOServiceAccount = "assets/pod-checkpointer-operator/service-account.yaml"
)

func MustAssetReader(asset string) io.Reader {
	return bytes.NewReader(MustAsset(asset))
}

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (*Factory) DefaultPCOClusterRole() (*rbacv1.ClusterRole, error) {
	return NewClusterRole(MustAssetReader(PCOClusterRole))
}

func NewClusterRole(manifest io.Reader) (*rbacv1.ClusterRole, error) {
	cr := rbacv1.ClusterRole{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

func (*Factory) DefaultPCOCustomResourceDefinition() (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	return NewCustomResourceDefinition(MustAssetReader(PCOCustomResource))
}

func NewCustomResourceDefinition(manifest io.Reader) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	crd := apiextensionsv1beta1.CustomResourceDefinition{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&crd); err != nil {
		return nil, err
	}
	return &crd, nil
}

func (*Factory) DefaultPCONamespace() (*corev1.Namespace, error) {
	return NewNamespace(MustAssetReader(PCONamespace))
}

func NewNamespace(manifest io.Reader) (*corev1.Namespace, error) {
	ns := corev1.Namespace{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&ns); err != nil {
		return nil, err
	}
	return &ns, nil
}

func (*Factory) DefaultPCORBAC() (*rbacv1.Role, error) {
	return NewRbac(MustAssetReader(PCORbac))
}

func NewRbac(manifest io.Reader) (*rbacv1.Role, error) {
	r := rbacv1.Role{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (*Factory) DefaultPCORoleBinding() (*rbacv1.RoleBinding, error) {
	return NewRoleBinding(MustAssetReader(PCORoleBinding))
}

func NewRoleBinding(manifest io.Reader) (*rbacv1.RoleBinding, error) {
	rb := rbacv1.RoleBinding{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&rb); err != nil {
		return nil, err
	}
	return &rb, nil
}

func (*Factory) DefaultServiceAccount() (*corev1.ServiceAccount, error) {
	return NewServiceAccount(MustAssetReader(PCOServiceAccount))
}

func NewServiceAccount(manifest io.Reader) (*corev1.ServiceAccount, error) {
	sa := corev1.ServiceAccount{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&sa); err != nil {
		return nil, err
	}
	return &sa, nil
}
