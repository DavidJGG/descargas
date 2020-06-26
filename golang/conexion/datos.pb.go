// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: datos.proto

package p2

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Datos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre        string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Departamento  string `protobuf:"bytes,2,opt,name=departamento,proto3" json:"departamento,omitempty"`
	Edad          string `protobuf:"bytes,3,opt,name=edad,proto3" json:"edad,omitempty"`
	FormaContagio string `protobuf:"bytes,4,opt,name=forma_contagio,json=formaContagio,proto3" json:"forma_contagio,omitempty"`
	Estado        string `protobuf:"bytes,5,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *Datos) Reset() {
	*x = Datos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datos) ProtoMessage() {}

func (x *Datos) ProtoReflect() protoreflect.Message {
	mi := &file_datos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datos.ProtoReflect.Descriptor instead.
func (*Datos) Descriptor() ([]byte, []int) {
	return file_datos_proto_rawDescGZIP(), []int{0}
}

func (x *Datos) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Datos) GetDepartamento() string {
	if x != nil {
		return x.Departamento
	}
	return ""
}

func (x *Datos) GetEdad() string {
	if x != nil {
		return x.Edad
	}
	return ""
}

func (x *Datos) GetFormaContagio() string {
	if x != nil {
		return x.FormaContagio
	}
	return ""
}

func (x *Datos) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type Respuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enviado bool   `protobuf:"varint,1,opt,name=Enviado,proto3" json:"Enviado,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Respuesta) Reset() {
	*x = Respuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Respuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Respuesta) ProtoMessage() {}

func (x *Respuesta) ProtoReflect() protoreflect.Message {
	mi := &file_datos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Respuesta.ProtoReflect.Descriptor instead.
func (*Respuesta) Descriptor() ([]byte, []int) {
	return file_datos_proto_rawDescGZIP(), []int{1}
}

func (x *Respuesta) GetEnviado() bool {
	if x != nil {
		return x.Enviado
	}
	return false
}

func (x *Respuesta) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_datos_proto protoreflect.FileDescriptor

var file_datos_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x32, 0x22, 0x96, 0x01, 0x0a, 0x05, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x64, 0x61, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x64, 0x61, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x67, 0x69, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x67,
	0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x37, 0x0a, 0x09, 0x52, 0x65,
	0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x76, 0x69, 0x61,
	0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x64,
	0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x32, 0x39, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x69, 0x6f, 0x44,
	0x61, 0x74, 0x6f, 0x73, 0x12, 0x28, 0x0a, 0x0c, 0x6f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x6f, 0x73, 0x12, 0x09, 0x2e, 0x70, 0x32, 0x2e, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x1a,
	0x0d, 0x2e, 0x70, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datos_proto_rawDescOnce sync.Once
	file_datos_proto_rawDescData = file_datos_proto_rawDesc
)

func file_datos_proto_rawDescGZIP() []byte {
	file_datos_proto_rawDescOnce.Do(func() {
		file_datos_proto_rawDescData = protoimpl.X.CompressGZIP(file_datos_proto_rawDescData)
	})
	return file_datos_proto_rawDescData
}

var file_datos_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_datos_proto_goTypes = []interface{}{
	(*Datos)(nil),     // 0: p2.Datos
	(*Respuesta)(nil), // 1: p2.Respuesta
}
var file_datos_proto_depIdxs = []int32{
	0, // 0: p2.ServicioDatos.obtenerDatos:input_type -> p2.Datos
	1, // 1: p2.ServicioDatos.obtenerDatos:output_type -> p2.Respuesta
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_datos_proto_init() }
func file_datos_proto_init() {
	if File_datos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datos); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Respuesta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_datos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datos_proto_goTypes,
		DependencyIndexes: file_datos_proto_depIdxs,
		MessageInfos:      file_datos_proto_msgTypes,
	}.Build()
	File_datos_proto = out.File
	file_datos_proto_rawDesc = nil
	file_datos_proto_goTypes = nil
	file_datos_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServicioDatosClient is the client API for ServicioDatos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServicioDatosClient interface {
	ObtenerDatos(ctx context.Context, in *Datos, opts ...grpc.CallOption) (*Respuesta, error)
}

type servicioDatosClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioDatosClient(cc grpc.ClientConnInterface) ServicioDatosClient {
	return &servicioDatosClient{cc}
}

func (c *servicioDatosClient) ObtenerDatos(ctx context.Context, in *Datos, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/p2.ServicioDatos/obtenerDatos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioDatosServer is the server API for ServicioDatos service.
type ServicioDatosServer interface {
	ObtenerDatos(context.Context, *Datos) (*Respuesta, error)
}

// UnimplementedServicioDatosServer can be embedded to have forward compatible implementations.
type UnimplementedServicioDatosServer struct {
}

func (*UnimplementedServicioDatosServer) ObtenerDatos(context.Context, *Datos) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerDatos not implemented")
}

func RegisterServicioDatosServer(s *grpc.Server, srv ServicioDatosServer) {
	s.RegisterService(&_ServicioDatos_serviceDesc, srv)
}

func _ServicioDatos_ObtenerDatos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Datos)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioDatosServer).ObtenerDatos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2.ServicioDatos/ObtenerDatos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioDatosServer).ObtenerDatos(ctx, req.(*Datos))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServicioDatos_serviceDesc = grpc.ServiceDesc{
	ServiceName: "p2.ServicioDatos",
	HandlerType: (*ServicioDatosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "obtenerDatos",
			Handler:    _ServicioDatos_ObtenerDatos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datos.proto",
}