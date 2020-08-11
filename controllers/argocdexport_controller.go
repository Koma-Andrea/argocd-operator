/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	batchv1 "k8s.io/api/batch/v1"
	batchv1b1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	argoprojiov1alpha1 "github.com/argoproj-labs/argocd-operator/api/v1alpha1"
)

// ArgoCDExportReconciler reconciles a ArgoCDExport object
type ArgoCDExportReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:namespace=placeholder,groups=argoproj.io,resources=argocdexports;argocdexports/finalizers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:namespace=placeholder,groups=argoproj.io,resources=argocdexports/status,verbs=get;update;patch
// +kubebuilder:rbac:namespace=placeholder,groups=batch,resources=cronjobs;jobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:namespace=placeholder,groups=core,resources=persistentvolumeclaims;secrets,verbs=get;list;watch;create;update;patch;delete

func (r *ArgoCDExportReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("argocdexport", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *ArgoCDExportReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&argoprojiov1alpha1.ArgoCDExport{}).
		Owns(&batchv1b1.CronJob{}).
		Owns(&batchv1.Job{}).
		Owns(&corev1.PersistentVolumeClaim{}).
		Owns(&corev1.Secret{}).
		Complete(r)
}
