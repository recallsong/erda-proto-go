// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: authentication.proto

package pb

import (
	http "github.com/erda-project/erda-infra/pkg/transport/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// AuthenticationServiceHandler is the server API for AuthenticationService service.
type AuthenticationServiceHandler interface {
}

// RegisterAuthenticationServiceHandler register AuthenticationServiceHandler to http.Router.
func RegisterAuthenticationServiceHandler(r http.Router, srv AuthenticationServiceHandler, opts ...http.HandleOption) {
}