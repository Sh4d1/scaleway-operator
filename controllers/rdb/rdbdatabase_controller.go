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

// RDBDatabaseReconciler reconciles a RDBDatabase object
type RDBDatabaseReconciler struct {
	ScalewayReconciler *controllers.ScalewayReconciler
}

// +kubebuilder:rbac:groups=rdb.scaleway.com,resources=rdbdatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rdb.scaleway.com,resources=rdbdatabases/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// Reconcile reconsiles the RDB Database
func (r *RDBDatabaseReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.ScalewayReconciler.Reconcile(req, &rdbv1alpha1.RDBDatabase{})
}

// SetupWithManager registers the RDB Database controller
func (r *RDBDatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rdbv1alpha1.RDBDatabase{}).
		Complete(r)
}
