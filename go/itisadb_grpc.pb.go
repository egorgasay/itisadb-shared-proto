// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: api/proto/itisadb.proto

package itisadb_go_proto

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

const (
	ItisaDB_Set_FullMethodName            = "/api.ItisaDB/Set"
	ItisaDB_Get_FullMethodName            = "/api.ItisaDB/Get"
	ItisaDB_Delete_FullMethodName         = "/api.ItisaDB/Delete"
	ItisaDB_DeleteIfExists_FullMethodName = "/api.ItisaDB/DeleteIfExists"
	ItisaDB_NewObject_FullMethodName      = "/api.ItisaDB/NewObject"
	ItisaDB_Object_FullMethodName         = "/api.ItisaDB/Object"
	ItisaDB_SetToObject_FullMethodName    = "/api.ItisaDB/SetToObject"
	ItisaDB_GetFromObject_FullMethodName  = "/api.ItisaDB/GetFromObject"
	ItisaDB_AttachToObject_FullMethodName = "/api.ItisaDB/AttachToObject"
	ItisaDB_ObjectToJSON_FullMethodName   = "/api.ItisaDB/ObjectToJSON"
	ItisaDB_IsObject_FullMethodName       = "/api.ItisaDB/IsObject"
	ItisaDB_Size_FullMethodName           = "/api.ItisaDB/Size"
	ItisaDB_DeleteObject_FullMethodName   = "/api.ItisaDB/DeleteObject"
	ItisaDB_DeleteAttr_FullMethodName     = "/api.ItisaDB/DeleteAttr"
	ItisaDB_Authenticate_FullMethodName   = "/api.ItisaDB/Authenticate"
	ItisaDB_CreateUser_FullMethodName     = "/api.ItisaDB/CreateUser"
	ItisaDB_DeleteUser_FullMethodName     = "/api.ItisaDB/DeleteUser"
	ItisaDB_ChangePassword_FullMethodName = "/api.ItisaDB/ChangePassword"
	ItisaDB_ChangeLevel_FullMethodName    = "/api.ItisaDB/ChangeLevel"
	ItisaDB_Connect_FullMethodName        = "/api.ItisaDB/Connect"
	ItisaDB_Disconnect_FullMethodName     = "/api.ItisaDB/Disconnect"
	ItisaDB_Servers_FullMethodName        = "/api.ItisaDB/Servers"
	ItisaDB_GetRam_FullMethodName         = "/api.ItisaDB/GetRam"
)

// ItisaDBClient is the client API for ItisaDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItisaDBClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	DeleteIfExists(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	NewObject(ctx context.Context, in *NewObjectRequest, opts ...grpc.CallOption) (*NewObjectResponse, error)
	Object(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectResponse, error)
	SetToObject(ctx context.Context, in *SetToObjectRequest, opts ...grpc.CallOption) (*SetToObjectResponse, error)
	GetFromObject(ctx context.Context, in *GetFromObjectRequest, opts ...grpc.CallOption) (*GetFromObjectResponse, error)
	AttachToObject(ctx context.Context, in *AttachToObjectRequest, opts ...grpc.CallOption) (*AttachToObjectResponse, error)
	ObjectToJSON(ctx context.Context, in *ObjectToJSONRequest, opts ...grpc.CallOption) (*ObjectToJSONResponse, error)
	IsObject(ctx context.Context, in *IsObjectRequest, opts ...grpc.CallOption) (*IsObjectResponse, error)
	Size(ctx context.Context, in *ObjectSizeRequest, opts ...grpc.CallOption) (*ObjectSizeResponse, error)
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error)
	DeleteAttr(ctx context.Context, in *DeleteAttrRequest, opts ...grpc.CallOption) (*DeleteAttrResponse, error)
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	ChangeLevel(ctx context.Context, in *ChangeLevelRequest, opts ...grpc.CallOption) (*ChangeLevelResponse, error)
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	Servers(ctx context.Context, in *ServersRequest, opts ...grpc.CallOption) (*ServersResponse, error)
	GetRam(ctx context.Context, in *GetRamRequest, opts ...grpc.CallOption) (*GetRamResponse, error)
}

type itisaDBClient struct {
	cc grpc.ClientConnInterface
}

func NewItisaDBClient(cc grpc.ClientConnInterface) ItisaDBClient {
	return &itisaDBClient{cc}
}

func (c *itisaDBClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, ItisaDB_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, ItisaDB_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, ItisaDB_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) DeleteIfExists(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, ItisaDB_DeleteIfExists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) NewObject(ctx context.Context, in *NewObjectRequest, opts ...grpc.CallOption) (*NewObjectResponse, error) {
	out := new(NewObjectResponse)
	err := c.cc.Invoke(ctx, ItisaDB_NewObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) Object(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectResponse, error) {
	out := new(ObjectResponse)
	err := c.cc.Invoke(ctx, ItisaDB_Object_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) SetToObject(ctx context.Context, in *SetToObjectRequest, opts ...grpc.CallOption) (*SetToObjectResponse, error) {
	out := new(SetToObjectResponse)
	err := c.cc.Invoke(ctx, ItisaDB_SetToObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) GetFromObject(ctx context.Context, in *GetFromObjectRequest, opts ...grpc.CallOption) (*GetFromObjectResponse, error) {
	out := new(GetFromObjectResponse)
	err := c.cc.Invoke(ctx, ItisaDB_GetFromObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) AttachToObject(ctx context.Context, in *AttachToObjectRequest, opts ...grpc.CallOption) (*AttachToObjectResponse, error) {
	out := new(AttachToObjectResponse)
	err := c.cc.Invoke(ctx, ItisaDB_AttachToObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) ObjectToJSON(ctx context.Context, in *ObjectToJSONRequest, opts ...grpc.CallOption) (*ObjectToJSONResponse, error) {
	out := new(ObjectToJSONResponse)
	err := c.cc.Invoke(ctx, ItisaDB_ObjectToJSON_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) IsObject(ctx context.Context, in *IsObjectRequest, opts ...grpc.CallOption) (*IsObjectResponse, error) {
	out := new(IsObjectResponse)
	err := c.cc.Invoke(ctx, ItisaDB_IsObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) Size(ctx context.Context, in *ObjectSizeRequest, opts ...grpc.CallOption) (*ObjectSizeResponse, error) {
	out := new(ObjectSizeResponse)
	err := c.cc.Invoke(ctx, ItisaDB_Size_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error) {
	out := new(DeleteObjectResponse)
	err := c.cc.Invoke(ctx, ItisaDB_DeleteObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) DeleteAttr(ctx context.Context, in *DeleteAttrRequest, opts ...grpc.CallOption) (*DeleteAttrResponse, error) {
	out := new(DeleteAttrResponse)
	err := c.cc.Invoke(ctx, ItisaDB_DeleteAttr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, ItisaDB_Authenticate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, ItisaDB_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, ItisaDB_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, ItisaDB_ChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) ChangeLevel(ctx context.Context, in *ChangeLevelRequest, opts ...grpc.CallOption) (*ChangeLevelResponse, error) {
	out := new(ChangeLevelResponse)
	err := c.cc.Invoke(ctx, ItisaDB_ChangeLevel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, ItisaDB_Connect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, ItisaDB_Disconnect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) Servers(ctx context.Context, in *ServersRequest, opts ...grpc.CallOption) (*ServersResponse, error) {
	out := new(ServersResponse)
	err := c.cc.Invoke(ctx, ItisaDB_Servers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itisaDBClient) GetRam(ctx context.Context, in *GetRamRequest, opts ...grpc.CallOption) (*GetRamResponse, error) {
	out := new(GetRamResponse)
	err := c.cc.Invoke(ctx, ItisaDB_GetRam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItisaDBServer is the server API for ItisaDB service.
// All implementations must embed UnimplementedItisaDBServer
// for forward compatibility
type ItisaDBServer interface {
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	DeleteIfExists(context.Context, *DeleteRequest) (*DeleteResponse, error)
	NewObject(context.Context, *NewObjectRequest) (*NewObjectResponse, error)
	Object(context.Context, *ObjectRequest) (*ObjectResponse, error)
	SetToObject(context.Context, *SetToObjectRequest) (*SetToObjectResponse, error)
	GetFromObject(context.Context, *GetFromObjectRequest) (*GetFromObjectResponse, error)
	AttachToObject(context.Context, *AttachToObjectRequest) (*AttachToObjectResponse, error)
	ObjectToJSON(context.Context, *ObjectToJSONRequest) (*ObjectToJSONResponse, error)
	IsObject(context.Context, *IsObjectRequest) (*IsObjectResponse, error)
	Size(context.Context, *ObjectSizeRequest) (*ObjectSizeResponse, error)
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)
	DeleteAttr(context.Context, *DeleteAttrRequest) (*DeleteAttrResponse, error)
	Authenticate(context.Context, *AuthRequest) (*AuthResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	ChangeLevel(context.Context, *ChangeLevelRequest) (*ChangeLevelResponse, error)
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	Servers(context.Context, *ServersRequest) (*ServersResponse, error)
	GetRam(context.Context, *GetRamRequest) (*GetRamResponse, error)
	mustEmbedUnimplementedItisaDBServer()
}

// UnimplementedItisaDBServer must be embedded to have forward compatible implementations.
type UnimplementedItisaDBServer struct {
}

func (UnimplementedItisaDBServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedItisaDBServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedItisaDBServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedItisaDBServer) DeleteIfExists(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIfExists not implemented")
}
func (UnimplementedItisaDBServer) NewObject(context.Context, *NewObjectRequest) (*NewObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewObject not implemented")
}
func (UnimplementedItisaDBServer) Object(context.Context, *ObjectRequest) (*ObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Object not implemented")
}
func (UnimplementedItisaDBServer) SetToObject(context.Context, *SetToObjectRequest) (*SetToObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetToObject not implemented")
}
func (UnimplementedItisaDBServer) GetFromObject(context.Context, *GetFromObjectRequest) (*GetFromObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFromObject not implemented")
}
func (UnimplementedItisaDBServer) AttachToObject(context.Context, *AttachToObjectRequest) (*AttachToObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachToObject not implemented")
}
func (UnimplementedItisaDBServer) ObjectToJSON(context.Context, *ObjectToJSONRequest) (*ObjectToJSONResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObjectToJSON not implemented")
}
func (UnimplementedItisaDBServer) IsObject(context.Context, *IsObjectRequest) (*IsObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsObject not implemented")
}
func (UnimplementedItisaDBServer) Size(context.Context, *ObjectSizeRequest) (*ObjectSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Size not implemented")
}
func (UnimplementedItisaDBServer) DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedItisaDBServer) DeleteAttr(context.Context, *DeleteAttrRequest) (*DeleteAttrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAttr not implemented")
}
func (UnimplementedItisaDBServer) Authenticate(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedItisaDBServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedItisaDBServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedItisaDBServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedItisaDBServer) ChangeLevel(context.Context, *ChangeLevelRequest) (*ChangeLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeLevel not implemented")
}
func (UnimplementedItisaDBServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedItisaDBServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedItisaDBServer) Servers(context.Context, *ServersRequest) (*ServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Servers not implemented")
}
func (UnimplementedItisaDBServer) GetRam(context.Context, *GetRamRequest) (*GetRamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRam not implemented")
}
func (UnimplementedItisaDBServer) mustEmbedUnimplementedItisaDBServer() {}

// UnsafeItisaDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItisaDBServer will
// result in compilation errors.
type UnsafeItisaDBServer interface {
	mustEmbedUnimplementedItisaDBServer()
}

func RegisterItisaDBServer(s grpc.ServiceRegistrar, srv ItisaDBServer) {
	s.RegisterService(&ItisaDB_ServiceDesc, srv)
}

func _ItisaDB_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_DeleteIfExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).DeleteIfExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_DeleteIfExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).DeleteIfExists(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_NewObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).NewObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_NewObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).NewObject(ctx, req.(*NewObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_Object_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).Object(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_Object_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).Object(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_SetToObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetToObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).SetToObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_SetToObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).SetToObject(ctx, req.(*SetToObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_GetFromObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFromObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).GetFromObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_GetFromObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).GetFromObject(ctx, req.(*GetFromObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_AttachToObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachToObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).AttachToObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_AttachToObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).AttachToObject(ctx, req.(*AttachToObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_ObjectToJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectToJSONRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).ObjectToJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_ObjectToJSON_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).ObjectToJSON(ctx, req.(*ObjectToJSONRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_IsObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).IsObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_IsObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).IsObject(ctx, req.(*IsObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_Size_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).Size(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_Size_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).Size(ctx, req.(*ObjectSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_DeleteObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_DeleteAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAttrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).DeleteAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_DeleteAttr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).DeleteAttr(ctx, req.(*DeleteAttrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_ChangeLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).ChangeLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_ChangeLevel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).ChangeLevel(ctx, req.(*ChangeLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_Disconnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_Servers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).Servers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_Servers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).Servers(ctx, req.(*ServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItisaDB_GetRam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItisaDBServer).GetRam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItisaDB_GetRam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItisaDBServer).GetRam(ctx, req.(*GetRamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItisaDB_ServiceDesc is the grpc.ServiceDesc for ItisaDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItisaDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ItisaDB",
	HandlerType: (*ItisaDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _ItisaDB_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ItisaDB_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ItisaDB_Delete_Handler,
		},
		{
			MethodName: "DeleteIfExists",
			Handler:    _ItisaDB_DeleteIfExists_Handler,
		},
		{
			MethodName: "NewObject",
			Handler:    _ItisaDB_NewObject_Handler,
		},
		{
			MethodName: "Object",
			Handler:    _ItisaDB_Object_Handler,
		},
		{
			MethodName: "SetToObject",
			Handler:    _ItisaDB_SetToObject_Handler,
		},
		{
			MethodName: "GetFromObject",
			Handler:    _ItisaDB_GetFromObject_Handler,
		},
		{
			MethodName: "AttachToObject",
			Handler:    _ItisaDB_AttachToObject_Handler,
		},
		{
			MethodName: "ObjectToJSON",
			Handler:    _ItisaDB_ObjectToJSON_Handler,
		},
		{
			MethodName: "IsObject",
			Handler:    _ItisaDB_IsObject_Handler,
		},
		{
			MethodName: "Size",
			Handler:    _ItisaDB_Size_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _ItisaDB_DeleteObject_Handler,
		},
		{
			MethodName: "DeleteAttr",
			Handler:    _ItisaDB_DeleteAttr_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _ItisaDB_Authenticate_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _ItisaDB_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _ItisaDB_DeleteUser_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _ItisaDB_ChangePassword_Handler,
		},
		{
			MethodName: "ChangeLevel",
			Handler:    _ItisaDB_ChangeLevel_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _ItisaDB_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _ItisaDB_Disconnect_Handler,
		},
		{
			MethodName: "Servers",
			Handler:    _ItisaDB_Servers_Handler,
		},
		{
			MethodName: "GetRam",
			Handler:    _ItisaDB_GetRam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/itisadb.proto",
}
