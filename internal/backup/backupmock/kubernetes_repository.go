// Code generated by mockery v1.0.0. DO NOT EDIT.

package backupmock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/api/networking/v1beta1"
)

// KubernetesRepository is an autogenerated mock type for the KubernetesRepository type
type KubernetesRepository struct {
	mock.Mock
}

// GetIngress provides a mock function with given fields: ctx, ns, name
func (_m *KubernetesRepository) GetIngress(ctx context.Context, ns string, name string) (*v1beta1.Ingress, error) {
	ret := _m.Called(ctx, ns, name)

	var r0 *v1beta1.Ingress
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *v1beta1.Ingress); ok {
		r0 = rf(ctx, ns, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Ingress)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, ns, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIngress provides a mock function with given fields: ctx, ingress
func (_m *KubernetesRepository) UpdateIngress(ctx context.Context, ingress *v1beta1.Ingress) error {
	ret := _m.Called(ctx, ingress)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Ingress) error); ok {
		r0 = rf(ctx, ingress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
