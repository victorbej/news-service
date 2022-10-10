// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/api/news.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContentCheckServiceClient is the client API for ContentCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentCheckServiceClient interface {
	// Проверка жизнеспособности сервиса
	CheckHealth(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type contentCheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentCheckServiceClient(cc grpc.ClientConnInterface) ContentCheckServiceClient {
	return &contentCheckServiceClient{cc}
}

func (c *contentCheckServiceClient) CheckHealth(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/api.ContentCheckService/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentCheckServiceServer is the server API for ContentCheckService service.
// All implementations must embed UnimplementedContentCheckServiceServer
// for forward compatibility
type ContentCheckServiceServer interface {
	// Проверка жизнеспособности сервиса
	CheckHealth(context.Context, *EmptyRequest) (*HealthResponse, error)
	mustEmbedUnimplementedContentCheckServiceServer()
}

// UnimplementedContentCheckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentCheckServiceServer struct {
}

func (UnimplementedContentCheckServiceServer) CheckHealth(context.Context, *EmptyRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (UnimplementedContentCheckServiceServer) mustEmbedUnimplementedContentCheckServiceServer() {}

// UnsafeContentCheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentCheckServiceServer will
// result in compilation errors.
type UnsafeContentCheckServiceServer interface {
	mustEmbedUnimplementedContentCheckServiceServer()
}

func RegisterContentCheckServiceServer(s grpc.ServiceRegistrar, srv ContentCheckServiceServer) {
	s.RegisterService(&ContentCheckService_ServiceDesc, srv)
}

func _ContentCheckService_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentCheckServiceServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ContentCheckService/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentCheckServiceServer).CheckHealth(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentCheckService_ServiceDesc is the grpc.ServiceDesc for ContentCheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentCheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ContentCheckService",
	HandlerType: (*ContentCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckHealth",
			Handler:    _ContentCheckService_CheckHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/news.proto",
}

// NewsServiceClient is the client API for NewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsServiceClient interface {
	// Получение списка новостей
	GetNews(ctx context.Context, in *NewsRequestParams, opts ...grpc.CallOption) (*NewsList, error)
	// Получение детальной информации по новости
	GetOne(ctx context.Context, in *ObjectId, opts ...grpc.CallOption) (*NewsObject, error)
	// Получение детальной информации по новости для отображения при переходе по письму
	GetOneBySlug(ctx context.Context, in *ObjectSlug, opts ...grpc.CallOption) (*NewsObject, error)
	// Создание новости
	Create(ctx context.Context, in *RequestNewsObject, opts ...grpc.CallOption) (*BaseResponse, error)
	// Обновление новости
	Update(ctx context.Context, in *RequestNewsObject, opts ...grpc.CallOption) (*BaseResponse, error)
	// Удаление новости
	Delete(ctx context.Context, in *ObjectId, opts ...grpc.CallOption) (*BaseResponse, error)
	// Получить ссылку на файл для скачивания
	GetFileLink(ctx context.Context, in *FileId, opts ...grpc.CallOption) (*FileLink, error)
}

type newsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsServiceClient(cc grpc.ClientConnInterface) NewsServiceClient {
	return &newsServiceClient{cc}
}

func (c *newsServiceClient) GetNews(ctx context.Context, in *NewsRequestParams, opts ...grpc.CallOption) (*NewsList, error) {
	out := new(NewsList)
	err := c.cc.Invoke(ctx, "/api.NewsService/GetNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetOne(ctx context.Context, in *ObjectId, opts ...grpc.CallOption) (*NewsObject, error) {
	out := new(NewsObject)
	err := c.cc.Invoke(ctx, "/api.NewsService/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetOneBySlug(ctx context.Context, in *ObjectSlug, opts ...grpc.CallOption) (*NewsObject, error) {
	out := new(NewsObject)
	err := c.cc.Invoke(ctx, "/api.NewsService/GetOneBySlug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) Create(ctx context.Context, in *RequestNewsObject, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/api.NewsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) Update(ctx context.Context, in *RequestNewsObject, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/api.NewsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) Delete(ctx context.Context, in *ObjectId, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/api.NewsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetFileLink(ctx context.Context, in *FileId, opts ...grpc.CallOption) (*FileLink, error) {
	out := new(FileLink)
	err := c.cc.Invoke(ctx, "/api.NewsService/GetFileLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServiceServer is the server API for NewsService service.
// All implementations must embed UnimplementedNewsServiceServer
// for forward compatibility
type NewsServiceServer interface {
	// Получение списка новостей
	GetNews(context.Context, *NewsRequestParams) (*NewsList, error)
	// Получение детальной информации по новости
	GetOne(context.Context, *ObjectId) (*NewsObject, error)
	// Получение детальной информации по новости для отображения при переходе по письму
	GetOneBySlug(context.Context, *ObjectSlug) (*NewsObject, error)
	// Создание новости
	Create(context.Context, *RequestNewsObject) (*BaseResponse, error)
	// Обновление новости
	Update(context.Context, *RequestNewsObject) (*BaseResponse, error)
	// Удаление новости
	Delete(context.Context, *ObjectId) (*BaseResponse, error)
	// Получить ссылку на файл для скачивания
	GetFileLink(context.Context, *FileId) (*FileLink, error)
	mustEmbedUnimplementedNewsServiceServer()
}

// UnimplementedNewsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNewsServiceServer struct {
}

func (UnimplementedNewsServiceServer) GetNews(context.Context, *NewsRequestParams) (*NewsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNews not implemented")
}
func (UnimplementedNewsServiceServer) GetOne(context.Context, *ObjectId) (*NewsObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedNewsServiceServer) GetOneBySlug(context.Context, *ObjectSlug) (*NewsObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneBySlug not implemented")
}
func (UnimplementedNewsServiceServer) Create(context.Context, *RequestNewsObject) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNewsServiceServer) Update(context.Context, *RequestNewsObject) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNewsServiceServer) Delete(context.Context, *ObjectId) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNewsServiceServer) GetFileLink(context.Context, *FileId) (*FileLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileLink not implemented")
}
func (UnimplementedNewsServiceServer) mustEmbedUnimplementedNewsServiceServer() {}

// UnsafeNewsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsServiceServer will
// result in compilation errors.
type UnsafeNewsServiceServer interface {
	mustEmbedUnimplementedNewsServiceServer()
}

func RegisterNewsServiceServer(s grpc.ServiceRegistrar, srv NewsServiceServer) {
	s.RegisterService(&NewsService_ServiceDesc, srv)
}

func _NewsService_GetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsRequestParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsService/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetNews(ctx, req.(*NewsRequestParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsService/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetOne(ctx, req.(*ObjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetOneBySlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectSlug)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetOneBySlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsService/GetOneBySlug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetOneBySlug(ctx, req.(*ObjectSlug))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNewsObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).Create(ctx, req.(*RequestNewsObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNewsObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).Update(ctx, req.(*RequestNewsObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).Delete(ctx, req.(*ObjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetFileLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetFileLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsService/GetFileLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetFileLink(ctx, req.(*FileId))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsService_ServiceDesc is the grpc.ServiceDesc for NewsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.NewsService",
	HandlerType: (*NewsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNews",
			Handler:    _NewsService_GetNews_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _NewsService_GetOne_Handler,
		},
		{
			MethodName: "GetOneBySlug",
			Handler:    _NewsService_GetOneBySlug_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _NewsService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NewsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NewsService_Delete_Handler,
		},
		{
			MethodName: "GetFileLink",
			Handler:    _NewsService_GetFileLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/news.proto",
}

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	// Получение тега
	Get(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TagList, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) Get(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TagList, error) {
	out := new(TagList)
	err := c.cc.Invoke(ctx, "/api.TagService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	// Получение тега
	Get(context.Context, *EmptyRequest) (*TagList, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) Get(context.Context, *EmptyRequest) (*TagList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s grpc.ServiceRegistrar, srv TagServiceServer) {
	s.RegisterService(&TagService_ServiceDesc, srv)
}

func _TagService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TagService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).Get(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagService_ServiceDesc is the grpc.ServiceDesc for TagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TagService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/news.proto",
}
