// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.24.0
// source: comment/service/v1/comment.proto

package v1

import (
	context "context"

	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCommentServiceCommentAction = "/comment.service.v1.CommentService/CommentAction"
const OperationCommentServiceGetCommentList = "/comment.service.v1.CommentService/GetCommentList"

type CommentServiceHTTPServer interface {
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionReply, error)
	GetCommentList(context.Context, *CommentListRequest) (*CommentListReply, error)
}

func RegisterCommentServiceHTTPServer(s *http.Server, srv CommentServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/douyin/comment/list", _CommentService_GetCommentList0_HTTP_Handler(srv))
	r.POST("/douyin/comment/action", _CommentService_CommentAction0_HTTP_Handler(srv))
}

func _CommentService_GetCommentList0_HTTP_Handler(srv CommentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CommentListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentServiceGetCommentList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCommentList(ctx, req.(*CommentListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommentListReply)
		return ctx.Result(200, reply)
	}
}

func _CommentService_CommentAction0_HTTP_Handler(srv CommentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CommentActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentServiceCommentAction)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CommentAction(ctx, req.(*CommentActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommentActionReply)
		return ctx.Result(200, reply)
	}
}

type CommentServiceHTTPClient interface {
	CommentAction(ctx context.Context, req *CommentActionRequest, opts ...http.CallOption) (rsp *CommentActionReply, err error)
	GetCommentList(ctx context.Context, req *CommentListRequest, opts ...http.CallOption) (rsp *CommentListReply, err error)
}

type CommentServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCommentServiceHTTPClient(client *http.Client) CommentServiceHTTPClient {
	return &CommentServiceHTTPClientImpl{client}
}

func (c *CommentServiceHTTPClientImpl) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...http.CallOption) (*CommentActionReply, error) {
	var out CommentActionReply
	pattern := "/douyin/comment/action"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCommentServiceCommentAction))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentServiceHTTPClientImpl) GetCommentList(ctx context.Context, in *CommentListRequest, opts ...http.CallOption) (*CommentListReply, error) {
	var out CommentListReply
	pattern := "/douyin/comment/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCommentServiceGetCommentList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
