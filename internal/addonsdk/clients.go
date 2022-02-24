package addonsdk

import (
	"context"

	addonsv1alpha1 "github.com/openshift/addon-operator/apis/addons/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// our tenants need to build a struct to implement this (here in reference addon, that struct is in `internal/referenceaddoninteractor/`)
// The tenant can choose to implement the following methods in whatever way they want depending on the client-go/controller-runtime they possess
type client interface {
	// the following GetAddonInstance method should be backed by a cache
	GetAddonInstance(context.Context, types.NamespacedName, *addonsv1alpha1.AddonInstance) error
	UpdateAddonInstanceStatus(ctx context.Context, addonInstance *addonsv1alpha1.AddonInstance) error
}

// if the tenant wants a super-custom way to report heartbeats or just like DIY stuff, they can implement their own `AddonInstanceReporterClient` and utilise it
type StatusReporterClient interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	SetConditions(ctx context.Context, conditions []metav1.Condition) error
	ReportAddonInstanceSpecChange(ctx context.Context, newAddonInstance *addonsv1alpha1.AddonInstance) error
}
