// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "user": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/user-microservice/design
// --out=$(GOPATH)/src/github.com/JormungandrK/user-microservice
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *UserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created(r *Users) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateUserContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// GetUserContext provides the user get action context.
type GetUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID string
}

// NewGetUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller get action.
func NewGetUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		rctx.UserID = rawUserID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetUserContext) OK(r *Users) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetUserContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// GetMeUserContext provides the user getMe action context.
type GetMeUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID *string
}

// NewGetMeUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller getMe action.
func NewGetMeUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetMeUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetMeUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		rctx.UserID = &rawUserID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetMeUserContext) OK(r *Users) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetMeUserContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID  string
	Payload *UserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		rctx.UserID = rawUserID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateUserContext) OK(r *Users) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
