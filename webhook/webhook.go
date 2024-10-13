// webhook/webhook.go
package webhook

import (
	"context"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	appsv1 "github.com/solate/devops-golang-test/api/v1"
)

var _ admission.Handler = &MyStatefulSetValidator{}

// MyStatefulSetValidator implements the ValidatingWebhook interface.
type MyStatefulSetValidator struct {
	Client  client.Client
	Decoder runtime.Decoder
}

// NewMyStatefulSetValidator returns a new instance of MyStatefulSetValidator.
func NewMyStatefulSetValidator(client client.Client, decoder runtime.Decoder) *MyStatefulSetValidator {
	return &MyStatefulSetValidator{Client: client, Decoder: decoder}
}

// Handle implements the admission.Handler interface.
func (v *MyStatefulSetValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	var mystatefulset appsv1.MyStatefulSet
	gvk := schema.GroupVersionKind{} // 初始化一个空的 GVK

	// 解码请求体
	if _, _, err := v.Decoder.Decode(req.Object.Raw, &gvk, &mystatefulset); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	// 添加验证逻辑
	if mystatefulset.Spec.Replicas <= 0 {
		return admission.Denied("Replicas must be greater than zero")
	}

	return admission.Allowed("Validation passed")
}

// InjectClient implements the webhook.Validator interface.
func (v *MyStatefulSetValidator) InjectClient(c client.Client) error {
	v.Client = c
	return nil
}

// InjectDecoder implements the webhook.Validator interface.
func (v *MyStatefulSetValidator) InjectDecoder(d runtime.Decoder) error {
	v.Decoder = d
	return nil
}
