package stub

import (
	"context"

	"github.com/openshift/pod-checkpointer-operator/pkg/apis/pod/v1alpha1"
	"github.com/openshift/pod-checkpointer-operator/pkg/manifests"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	defaultNamespace = "openshift-pod-checkpointer"
	defaultImage     = "quay.io/coreos/pod-checkpointer:9dc83e1ab3bc36ca25c9f7c18ddef1b91d4a0558"
)

func NewHandler(m *Metrics, factory *manifests.Factory) *Handler {
	return &Handler{
		metrics:          m,
		manifestsFactory: factory,
	}
}

type Metrics struct {
	operatorErrors prometheus.Counter
}

type Handler struct {
	metrics          *Metrics
	manifestsFactory *manifests.Factory
}

func (h *Handler) EnsureObjects() error {
	return nil
	/*
		cr, err := h.manifestsFactory.New
		if err != nil {
			return fmt.Errorf("couldn't build router cluster role: %v", err)
		}
		err = sdk.Create(cr)
		if err == nil {
			logrus.Infof("created router cluster role %q", cr.Name)
		} else if !errors.IsAlreadyExists(err) {
			return fmt.Errorf("couldn't create router cluster role: %v", err)
		}

		ns, err := h.manifestsFactory.RouterNamespace()
		if err != nil {
			return fmt.Errorf("couldn't build router namespace: %v", err)
		}
		err = sdk.Create(ns)
		if err == nil {
			logrus.Infof("created router namespace %q", ns.Name)
		} else if !errors.IsAlreadyExists(err) {
			return fmt.Errorf("couldn't create router namespace %q: %v", ns.Name, err)
		}
	*/
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.PodCheckpointerOperator:
		err := sdk.Create(newCheckpointerDaemonSet(o))
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create checkpointer pod : %v", err)
			h.metrics.operatorErrors.Inc()
			return err
		}
	}
	return nil
}

// newCheckpointerDaemonSet demonstrates how to create a busybox pod
func newCheckpointerDaemonSet(cr *v1alpha1.PodCheckpointerOperator) *appsv1.DaemonSet {
	return &appsv1.DaemonSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "DaemonSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod-checkpointer",
			Namespace: defaultNamespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "PodCheckpointerOperator",
				}),
			},
			Labels: map[string]string{
				"tier":    "control-plane",
				"k8s-app": "pod-checkpointer",
			},
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"tier":    "control-plane",
					"k8s-app": "pod-checkpointer",
				},
			},
			UpdateStrategy: appsv1.DaemonSetUpdateStrategy{
				Type: appsv1.RollingUpdateDaemonSetStrategyType,
				RollingUpdate: &appsv1.RollingUpdateDaemonSet{
					MaxUnavailable: func() *intstr.IntOrString { i := intstr.FromInt(1); return &i }(),
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"tier":    "control-plane",
						"k8s-app": "pod-checkpointer",
					},
					Annotations: map[string]string{
						"checkpointer.alpha.coreos.com/checkpoint": "true",
					},
				},
				Spec: corev1.PodSpec{
					HostNetwork: true,
					Volumes: []corev1.Volume{
						corev1.Volume{
							Name: "kubeconfig",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "kubeconfig-in-cluster",
									},
								},
							},
						},
						corev1.Volume{
							Name: "etc-kubernetes",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/etc/kubernetes",
								},
							},
						},
						corev1.Volume{
							Name: "var-run",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/var/run",
								},
							},
						},
					},
					Containers: []corev1.Container{
						{
							Name:  "pod-checkpointer",
							Image: defaultImage,
							VolumeMounts: []corev1.VolumeMount{
								corev1.VolumeMount{
									Name:      "kubeconfig",
									MountPath: "/etc/checkpointer",
								},
								corev1.VolumeMount{
									Name:      "etc-kubernetes",
									MountPath: "/etc/kubernetes",
								},
								corev1.VolumeMount{
									Name:      "var-run",
									MountPath: "/var/run",
								},
							},
							Env: []corev1.EnvVar{
								corev1.EnvVar{
									Name: "NODE_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "spec.nodeName",
										},
									},
								},
								corev1.EnvVar{
									Name: "POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
								corev1.EnvVar{
									Name: "POD_NAMESPACE",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.namespace",
										},
									},
								},
							},
							Command: []string{
								"/checkpoint",
								"--lock-file=/var/run/lock/pod-checkpointer.lock",
								"--kubeconfig=/etc/checkpointer/kubeconfig",
								"--checkpoint-grace-period=5m",
							},
						},
					},
					NodeSelector: map[string]string{
						"node-role.kubernetes.io/master": "",
					},
					RestartPolicy:      corev1.RestartPolicyAlways,
					ServiceAccountName: "pod-checkpointer",
					Tolerations: []corev1.Toleration{
						corev1.Toleration{
							Key:      "node-role.kubernetes.io/master",
							Operator: corev1.TolerationOpExists,
							Effect:   corev1.TaintEffectNoSchedule,
						},
					},
				},
			},
		},
	}
}

func RegisterOperatorMetrics() (*Metrics, error) {
	operatorErrors := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "pod_checkpointer_operator_reconcile_errors_total",
		Help: "Number of errors that occurred while reconciling the pod checkpointer operator",
	})
	err := prometheus.Register(operatorErrors)
	if err != nil {
		return nil, err
	}
	return &Metrics{operatorErrors: operatorErrors}, nil
}