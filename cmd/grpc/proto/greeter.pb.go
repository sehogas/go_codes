// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: cmd/grpc/proto/greeter.proto

package proto

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

type Language int32

const (
	Language_ENGLISH Language = 0
	Language_SPANISH Language = 1
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "ENGLISH",
		1: "SPANISH",
	}
	Language_value = map[string]int32{
		"ENGLISH": 0,
		"SPANISH": 1,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_grpc_proto_greeter_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_cmd_grpc_proto_greeter_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_cmd_grpc_proto_greeter_proto_rawDescGZIP(), []int{0}
}

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Language Language `protobuf:"varint,2,opt,name=language,proto3,enum=Language" json:"language,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_grpc_proto_greeter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_grpc_proto_greeter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_cmd_grpc_proto_greeter_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_ENGLISH
}

// The response message containing the greetings
type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Language Language `protobuf:"varint,2,opt,name=language,proto3,enum=Language" json:"language,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_grpc_proto_greeter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_grpc_proto_greeter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_cmd_grpc_proto_greeter_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HelloResponse) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_ENGLISH
}

type ListRegardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*HelloResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListRegardsResponse) Reset() {
	*x = ListRegardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_grpc_proto_greeter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRegardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRegardsResponse) ProtoMessage() {}

func (x *ListRegardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_grpc_proto_greeter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRegardsResponse.ProtoReflect.Descriptor instead.
func (*ListRegardsResponse) Descriptor() ([]byte, []int) {
	return file_cmd_grpc_proto_greeter_proto_rawDescGZIP(), []int{2}
}

func (x *ListRegardsResponse) GetItems() []*HelloResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_cmd_grpc_proto_greeter_proto protoreflect.FileDescriptor

var file_cmd_grpc_proto_greeter_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6d, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x0d, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x24, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x47, 0x4c, 0x49, 0x53, 0x48, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x50, 0x41, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x01, 0x32, 0x9e,
	0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x53, 0x61,
	0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0d, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x41, 0x67, 0x61, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x53, 0x61, 0x79,
	0x52, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x73, 0x12, 0x0d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x67, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65,
	0x68, 0x6f, 0x67, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cmd_grpc_proto_greeter_proto_rawDescOnce sync.Once
	file_cmd_grpc_proto_greeter_proto_rawDescData = file_cmd_grpc_proto_greeter_proto_rawDesc
)

func file_cmd_grpc_proto_greeter_proto_rawDescGZIP() []byte {
	file_cmd_grpc_proto_greeter_proto_rawDescOnce.Do(func() {
		file_cmd_grpc_proto_greeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_grpc_proto_greeter_proto_rawDescData)
	})
	return file_cmd_grpc_proto_greeter_proto_rawDescData
}

var file_cmd_grpc_proto_greeter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_grpc_proto_greeter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmd_grpc_proto_greeter_proto_goTypes = []interface{}{
	(Language)(0),               // 0: Language
	(*HelloRequest)(nil),        // 1: HelloRequest
	(*HelloResponse)(nil),       // 2: HelloResponse
	(*ListRegardsResponse)(nil), // 3: ListRegardsResponse
}
var file_cmd_grpc_proto_greeter_proto_depIdxs = []int32{
	0, // 0: HelloRequest.language:type_name -> Language
	0, // 1: HelloResponse.language:type_name -> Language
	2, // 2: ListRegardsResponse.items:type_name -> HelloResponse
	1, // 3: Greeter.SayHello:input_type -> HelloRequest
	1, // 4: Greeter.SayHelloAgain:input_type -> HelloRequest
	1, // 5: Greeter.SayReguards:input_type -> HelloRequest
	2, // 6: Greeter.SayHello:output_type -> HelloResponse
	2, // 7: Greeter.SayHelloAgain:output_type -> HelloResponse
	3, // 8: Greeter.SayReguards:output_type -> ListRegardsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cmd_grpc_proto_greeter_proto_init() }
func file_cmd_grpc_proto_greeter_proto_init() {
	if File_cmd_grpc_proto_greeter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_grpc_proto_greeter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_cmd_grpc_proto_greeter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_cmd_grpc_proto_greeter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRegardsResponse); i {
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
			RawDescriptor: file_cmd_grpc_proto_greeter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_grpc_proto_greeter_proto_goTypes,
		DependencyIndexes: file_cmd_grpc_proto_greeter_proto_depIdxs,
		EnumInfos:         file_cmd_grpc_proto_greeter_proto_enumTypes,
		MessageInfos:      file_cmd_grpc_proto_greeter_proto_msgTypes,
	}.Build()
	File_cmd_grpc_proto_greeter_proto = out.File
	file_cmd_grpc_proto_greeter_proto_rawDesc = nil
	file_cmd_grpc_proto_greeter_proto_goTypes = nil
	file_cmd_grpc_proto_greeter_proto_depIdxs = nil
}
