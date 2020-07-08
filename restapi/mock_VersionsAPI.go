// Code generated by mockery v1.0.0. DO NOT EDIT.

package restapi

import (
	context "context"

	middleware "github.com/go-openapi/runtime/middleware"
	mock "github.com/stretchr/testify/mock"

	versions "github.com/filanov/bm-inventory/restapi/operations/versions"
)

// MockVersionsAPI is an autogenerated mock type for the VersionsAPI type
type MockVersionsAPI struct {
	mock.Mock
}

// ListComponentVersions provides a mock function with given fields: ctx, params
func (_m *MockVersionsAPI) ListComponentVersions(ctx context.Context, params versions.ListComponentVersionsParams) middleware.Responder {
	ret := _m.Called(ctx, params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(context.Context, versions.ListComponentVersionsParams) middleware.Responder); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}
