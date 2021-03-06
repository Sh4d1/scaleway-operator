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
	ctrl "sigs.k8s.io/controller-runtime"

	rdbv1alpha1 "github.com/scaleway/scaleway-operator/apis/rdb/v1alpha1"
	"github.com/scaleway/scaleway-operator/controllers"
)

// RDBInstanceReconciler reconciles a RDBInstance object
type RDBInstanceReconciler struct {
	ScalewayReconciler *controllers.ScalewayReconciler
}

// +kubebuilder:rbac:groups=rdb.scaleway.com,resources=rdbinstances,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rdb.scaleway.com,resources=rdbinstances/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch

// Reconcile reconciles the RDB Instance
func (r *RDBInstanceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.ScalewayReconciler.Reconcile(req, &rdbv1alpha1.RDBInstance{})
}

// SetupWithManager registers the RDB Instance Controller
func (r *RDBInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rdbv1alpha1.RDBInstance{}).
		Complete(r)
}
