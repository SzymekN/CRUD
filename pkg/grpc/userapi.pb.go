// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pkg/grpc/userapi.proto

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\pkg\grpc\userapi.proto

package grpc

import (
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

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_userapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_userapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_userapi_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Id        int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Age       int64  `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UserPost) Reset() {
	*x = UserPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_userapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPost) ProtoMessage() {}

func (x *UserPost) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_userapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPost.ProtoReflect.Descriptor instead.
func (*UserPost) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_userapi_proto_rawDescGZIP(), []int{1}
}

func (x *UserPost) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserPost) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserPost) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserPost) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserPost) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type UserPut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Id        int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Age       int64  `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UserPut) Reset() {
	*x = UserPut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_userapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPut) ProtoMessage() {}

func (x *UserPut) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_userapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPut.ProtoReflect.Descriptor instead.
func (*UserPut) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_userapi_proto_rawDescGZIP(), []int{2}
}

func (x *UserPut) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserPut) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserPut) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserPut) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserPut) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type UserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Id        int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Age       int64  `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UserReply) Reset() {
	*x = UserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_userapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReply) ProtoMessage() {}

func (x *UserReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_userapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReply.ProtoReflect.Descriptor instead.
func (*UserReply) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_userapi_proto_rawDescGZIP(), []int{3}
}

func (x *UserReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserReply) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserReply) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserReply) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_pkg_grpc_userapi_proto protoreflect.FileDescriptor

var file_pkg_grpc_userapi_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x1d,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x80, 0x01,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x22, 0x7f, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0x81, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x32, 0xfe, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x11,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2d, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x07, 0x50, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x74, 0x1a,
	0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_userapi_proto_rawDescOnce sync.Once
	file_pkg_grpc_userapi_proto_rawDescData = file_pkg_grpc_userapi_proto_rawDesc
)

func file_pkg_grpc_userapi_proto_rawDescGZIP() []byte {
	file_pkg_grpc_userapi_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_userapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_userapi_proto_rawDescData)
	})
	return file_pkg_grpc_userapi_proto_rawDescData
}

var file_pkg_grpc_userapi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_grpc_userapi_proto_goTypes = []interface{}{
	(*UserRequest)(nil), // 0: grpc.UserRequest
	(*UserPost)(nil),    // 1: grpc.UserPost
	(*UserPut)(nil),     // 2: grpc.UserPut
	(*UserReply)(nil),   // 3: grpc.UserReply
}
var file_pkg_grpc_userapi_proto_depIdxs = []int32{
	0, // 0: grpc.Userapi.GetUser:input_type -> grpc.UserRequest
	0, // 1: grpc.Userapi.GetUsers:input_type -> grpc.UserRequest
	1, // 2: grpc.Userapi.PostUser:input_type -> grpc.UserPost
	2, // 3: grpc.Userapi.PutUser:input_type -> grpc.UserPut
	0, // 4: grpc.Userapi.DeleteUser:input_type -> grpc.UserRequest
	3, // 5: grpc.Userapi.GetUser:output_type -> grpc.UserReply
	3, // 6: grpc.Userapi.GetUsers:output_type -> grpc.UserReply
	3, // 7: grpc.Userapi.PostUser:output_type -> grpc.UserReply
	3, // 8: grpc.Userapi.PutUser:output_type -> grpc.UserReply
	3, // 9: grpc.Userapi.DeleteUser:output_type -> grpc.UserReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_grpc_userapi_proto_init() }
func file_pkg_grpc_userapi_proto_init() {
	if File_pkg_grpc_userapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_userapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_pkg_grpc_userapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPost); i {
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
		file_pkg_grpc_userapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPut); i {
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
		file_pkg_grpc_userapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReply); i {
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
			RawDescriptor: file_pkg_grpc_userapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_userapi_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_userapi_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_userapi_proto_msgTypes,
	}.Build()
	File_pkg_grpc_userapi_proto = out.File
	file_pkg_grpc_userapi_proto_rawDesc = nil
	file_pkg_grpc_userapi_proto_goTypes = nil
	file_pkg_grpc_userapi_proto_depIdxs = nil
}
