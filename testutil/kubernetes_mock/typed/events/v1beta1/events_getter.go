// Code generated by mockery v1.1.2. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
)

// EventsGetter is an autogenerated mock type for the EventsGetter type
type EventsGetter struct {
	mock.Mock
}

// Events provides a mock function with given fields: namespace
func (_m *EventsGetter) Events(namespace string) v1beta1.EventInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.EventInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.EventInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.EventInterface)
		}
	}

	return r0
}
