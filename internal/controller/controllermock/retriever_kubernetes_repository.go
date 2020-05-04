// Code generated by mockery v1.0.0. DO NOT EDIT.

package controllermock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/slok/bilrost/pkg/apis/auth/v1"

	v1beta1 "k8s.io/api/networking/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// RetrieverKubernetesRepository is an autogenerated mock type for the RetrieverKubernetesRepository type
type RetrieverKubernetesRepository struct {
	mock.Mock
}

// ListIngressAuths provides a mock function with given fields: ctx, ns, labelSelector
func (_m *RetrieverKubernetesRepository) ListIngressAuths(ctx context.Context, ns string, labelSelector map[string]string) (*v1.IngressAuthList, error) {
	ret := _m.Called(ctx, ns, labelSelector)

	var r0 *v1.IngressAuthList
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) *v1.IngressAuthList); ok {
		r0 = rf(ctx, ns, labelSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.IngressAuthList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, ns, labelSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIngresses provides a mock function with given fields: ctx, ns, labelSelector
func (_m *RetrieverKubernetesRepository) ListIngresses(ctx context.Context, ns string, labelSelector map[string]string) (*v1beta1.IngressList, error) {
	ret := _m.Called(ctx, ns, labelSelector)

	var r0 *v1beta1.IngressList
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) *v1beta1.IngressList); ok {
		r0 = rf(ctx, ns, labelSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.IngressList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, ns, labelSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchIngressAuths provides a mock function with given fields: ctx, ns, labelSelector
func (_m *RetrieverKubernetesRepository) WatchIngressAuths(ctx context.Context, ns string, labelSelector map[string]string) (watch.Interface, error) {
	ret := _m.Called(ctx, ns, labelSelector)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) watch.Interface); ok {
		r0 = rf(ctx, ns, labelSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, ns, labelSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchIngresses provides a mock function with given fields: ctx, ns, labelSelector
func (_m *RetrieverKubernetesRepository) WatchIngresses(ctx context.Context, ns string, labelSelector map[string]string) (watch.Interface, error) {
	ret := _m.Called(ctx, ns, labelSelector)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) watch.Interface); ok {
		r0 = rf(ctx, ns, labelSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, ns, labelSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}