package main

import (
	context "context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	k8scorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func deleteService(ctx context.Context, name string, cl k8scorev1.ServiceInterface) error {
	return cl.Delete(ctx, name, metav1.DeleteOptions{})
}

func newService(ctx context.Context, namespace, name string) *corev1.Service {
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind: "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "cscaler",
			Labels:    labels(name),
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:     "standard",
					Protocol: corev1.ProtocolTCP,
					Port:     8080,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 8080,
					},
				},
			},
			Selector: labels(name),
			Type:     corev1.ServiceTypeClusterIP,
		},
	}
}
