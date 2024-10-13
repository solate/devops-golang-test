/*
Copyright 2024.

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
	"fmt"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "github.com/solate/devops-golang-test/api/v1"
)

const (
	podNameTemplate = "%s-%s"
)

// MyStatefulSetReconciler reconciles a MyStatefulSet object
type MyStatefulSetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.example.com,resources=mystatefulsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.example.com,resources=mystatefulsets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.example.com,resources=mystatefulsets/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MyStatefulSet object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *MyStatefulSetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// _ = log.FromContext(ctx)

	// TODO(user): your logic here

	logger := log.FromContext(ctx)

	var mystatefulset appsv1.MyStatefulSet
	if err := r.Get(ctx, req.NamespacedName, &mystatefulset); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// Ensure the number of replicas
	for i := int32(0); i < mystatefulset.Spec.Replicas; i++ {
		podName := fmt.Sprintf(podNameTemplate, mystatefulset.Name, strconv.Itoa(int(i)))
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: mystatefulset.Namespace,
				Name:      podName,
				Labels:    map[string]string{"app": mystatefulset.Name},
			},
			Spec: mystatefulset.Spec.Template.Spec,
		}

		if err := r.Create(ctx, pod); err != nil {
			logger.Error(err, "Failed to create pod")
			return ctrl.Result{}, err
		}
	}

	// Update status
	mystatefulset.Status.ReadyReplicas = mystatefulset.Spec.Replicas
	if err := r.Status().Update(ctx, &mystatefulset); err != nil {
		logger.Error(err, "Failed to update MyStatefulSet status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MyStatefulSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.MyStatefulSet{}).
		Complete(r)
}
