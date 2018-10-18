package stub

import (
	"context"

	"github.com/openshift/pod-checkpointer-operator/pkg/apis/pod/v1alpha1"
	"github.com/openshift/pod-checkpointer-operator/pkg/manifests"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
)

const (
	defaultImage = "quay.io/coreos/pod-checkpointer:018007e77ccd61e8e59b7e15d7fc5e318a5a2682"
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
	switch event.Object.(type) {
	case *v1alpha1.PodCheckpointerOperator:
		ds, err := h.manifestsFactory.DefaultDaemonSet(defaultImage)
		if err != nil {
			logrus.Errorf("failed to load daemonset object: %v", err)
			h.metrics.operatorErrors.Inc()
			return err
		}
		err = sdk.Create(ds)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create checkpointer daemonset: %v", err)
			h.metrics.operatorErrors.Inc()
			return err
		}
	}
	return nil
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
