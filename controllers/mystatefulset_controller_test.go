package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	appsv1 "github.com/solate/devops-golang-test/api/v1"
)

func TestReconcile(t *testing.T) {
	// Create a fake client
	fakeClient := fake.NewClientBuilder().Build()

	// Create a MyStatefulSet instance
	mystatefulset := &appsv1.MyStatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-mystatefulset",
			Namespace: "default",
		},
		Spec: appsv1.MyStatefulSetSpec{
			Replicas: 3,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test-container",
							Image: "test-image",
						},
					},
				},
			},
		},
	}

	// Create the MyStatefulSet in the fake client
	err := fakeClient.Create(context.TODO(), mystatefulset)
	assert.NoError(t, err)

	// Initialize the controller
	reconciler := &MyStatefulSetReconciler{
		Client: fakeClient,
		Scheme: nil,
	}

	// Reconcile the MyStatefulSet
	_, err = reconciler.Reconcile(context.TODO(), reconcile.Request{NamespacedName: types.NamespacedName{Name: "test-mystatefulset", Namespace: "default"}})
	assert.NoError(t, err)

	// Check if the pods were created
	podList := &corev1.PodList{}
	err = fakeClient.List(context.TODO(), podList, client.InNamespace("default"))
	assert.NoError(t, err)
	assert.Equal(t, 3, len(podList.Items))
}
