// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.2
// source: proto/screamer.proto

package compiled

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

type MType int32

const (
	MType_COUNTER MType = 0
	MType_GAUGE   MType = 1
)

// Enum value maps for MType.
var (
	MType_name = map[int32]string{
		0: "COUNTER",
		1: "GAUGE",
	}
	MType_value = map[string]int32{
		"COUNTER": 0,
		"GAUGE":   1,
	}
)

func (x MType) Enum() *MType {
	p := new(MType)
	*p = x
	return p
}

func (x MType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_screamer_proto_enumTypes[0].Descriptor()
}

func (MType) Type() protoreflect.EnumType {
	return &file_proto_screamer_proto_enumTypes[0]
}

func (x MType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MType.Descriptor instead.
func (MType) EnumDescriptor() ([]byte, []int) {
	return file_proto_screamer_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mtype         MType                  `protobuf:"varint,2,opt,name=mtype,proto3,enum=demo.MType" json:"mtype,omitempty"`
	Delta         int64                  `protobuf:"varint,3,opt,name=delta,proto3" json:"delta,omitempty"`
	Value         float32                `protobuf:"fixed32,4,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_proto_screamer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_screamer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_screamer_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Request) GetMtype() MType {
	if x != nil {
		return x.Mtype
	}
	return MType_COUNTER
}

func (x *Request) GetDelta() int64 {
	if x != nil {
		return x.Delta
	}
	return 0
}

func (x *Request) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Ident struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ident) Reset() {
	*x = Ident{}
	mi := &file_proto_screamer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ident) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ident) ProtoMessage() {}

func (x *Ident) ProtoReflect() protoreflect.Message {
	mi := &file_proto_screamer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ident.ProtoReflect.Descriptor instead.
func (*Ident) Descriptor() ([]byte, []int) {
	return file_proto_screamer_proto_rawDescGZIP(), []int{1}
}

func (x *Ident) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Ident) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ident         *Ident                 `protobuf:"bytes,1,opt,name=ident,proto3" json:"ident,omitempty"`
	Value         float32                `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_proto_screamer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_screamer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_screamer_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetIdent() *Ident {
	if x != nil {
		return x.Ident
	}
	return nil
}

func (x *Response) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_proto_screamer_proto protoreflect.FileDescriptor

var file_proto_screamer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x68, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x6d, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x05, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2f, 0x0a, 0x05, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x1f, 0x0a, 0x05,
	0x4d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x41, 0x55, 0x47, 0x45, 0x10, 0x01, 0x32, 0x6a, 0x0a,
	0x0f, 0x53, 0x63, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x0d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0d, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x73, 0x63, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_screamer_proto_rawDescOnce sync.Once
	file_proto_screamer_proto_rawDescData = file_proto_screamer_proto_rawDesc
)

func file_proto_screamer_proto_rawDescGZIP() []byte {
	file_proto_screamer_proto_rawDescOnce.Do(func() {
		file_proto_screamer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_screamer_proto_rawDescData)
	})
	return file_proto_screamer_proto_rawDescData
}

var file_proto_screamer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_screamer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_screamer_proto_goTypes = []any{
	(MType)(0),       // 0: demo.MType
	(*Request)(nil),  // 1: demo.Request
	(*Ident)(nil),    // 2: demo.Ident
	(*Response)(nil), // 3: demo.Response
}
var file_proto_screamer_proto_depIdxs = []int32{
	0, // 0: demo.Request.mtype:type_name -> demo.MType
	2, // 1: demo.Response.ident:type_name -> demo.Ident
	1, // 2: demo.ScreamerService.UpdateValue:input_type -> demo.Request
	1, // 3: demo.ScreamerService.GetValue:input_type -> demo.Request
	3, // 4: demo.ScreamerService.UpdateValue:output_type -> demo.Response
	3, // 5: demo.ScreamerService.GetValue:output_type -> demo.Response
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_screamer_proto_init() }
func file_proto_screamer_proto_init() {
	if File_proto_screamer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_screamer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_screamer_proto_goTypes,
		DependencyIndexes: file_proto_screamer_proto_depIdxs,
		EnumInfos:         file_proto_screamer_proto_enumTypes,
		MessageInfos:      file_proto_screamer_proto_msgTypes,
	}.Build()
	File_proto_screamer_proto = out.File
	file_proto_screamer_proto_rawDesc = nil
	file_proto_screamer_proto_goTypes = nil
	file_proto_screamer_proto_depIdxs = nil
}