// Code generated by goa v3.7.12, DO NOT EDIT.
//
// secret endpoints
//
// Command:
// $ goa gen secretsvc/design

package secret

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "secret" service endpoints.
type Endpoints struct {
	GetSecret    goa.Endpoint
	CreateSecret goa.Endpoint
}

// NewEndpoints wraps the methods of the "secret" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetSecret:    NewGetSecretEndpoint(s),
		CreateSecret: NewCreateSecretEndpoint(s),
	}
}

// Use applies the given middleware to all the "secret" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetSecret = m(e.GetSecret)
	e.CreateSecret = m(e.CreateSecret)
}

// NewGetSecretEndpoint returns an endpoint function that calls the method
// "getSecret" of service "secret".
func NewGetSecretEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetSecret(ctx)
	}
}

// NewCreateSecretEndpoint returns an endpoint function that calls the method
// "createSecret" of service "secret".
func NewCreateSecretEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateSecretRequestBody)
		return s.CreateSecret(ctx, p)
	}
}
