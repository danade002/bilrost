// Code generated by mockery v1.0.0. DO NOT EDIT.

package securitymock

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/slok/bilrost/internal/model"

// KubeServiceTranslator is an autogenerated mock type for the KubeServiceTranslator type
type KubeServiceTranslator struct {
	mock.Mock
}

// GetServiceHostAndPort provides a mock function with given fields: ctx, svc
func (_m *KubeServiceTranslator) GetServiceHostAndPort(ctx context.Context, svc model.KubernetesService) (string, int, error) {
	ret := _m.Called(ctx, svc)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, model.KubernetesService) string); ok {
		r0 = rf(ctx, svc)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, model.KubernetesService) int); ok {
		r1 = rf(ctx, svc)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, model.KubernetesService) error); ok {
		r2 = rf(ctx, svc)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
