// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: proto.proto

package log

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

type Status int32

const (
	Status_Ok  Status = 0
	Status_Err Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Ok",
		1: "Err",
	}
	Status_value = map[string]int32{
		"Ok":  0,
		"Err": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_proto_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{0}
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value [][]byte `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetValue() [][]byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ProduceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img *Image `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
}

func (x *ProduceRequest) Reset() {
	*x = ProduceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProduceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProduceRequest) ProtoMessage() {}

func (x *ProduceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProduceRequest.ProtoReflect.Descriptor instead.
func (*ProduceRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{1}
}

func (x *ProduceRequest) GetImg() *Image {
	if x != nil {
		return x.Img
	}
	return nil
}

type ProduceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *ProduceResponse) Reset() {
	*x = ProduceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProduceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProduceResponse) ProtoMessage() {}

func (x *ProduceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProduceResponse.ProtoReflect.Descriptor instead.
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{2}
}

func (x *ProduceResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Ok
}

var File_proto_proto protoreflect.FileDescriptor

var file_proto_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2a, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x22, 0x32, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x19, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x45, 0x72, 0x72, 0x10, 0x01, 0x32, 0x3e, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x12, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x62, 0x2f, 0x6c, 0x6f,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proto_rawDescOnce sync.Once
	file_proto_proto_rawDescData = file_proto_proto_rawDesc
)

func file_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_rawDescData)
	})
	return file_proto_proto_rawDescData
}

var file_proto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_proto_goTypes = []interface{}{
	(Status)(0),             // 0: Status
	(*Image)(nil),           // 1: Image
	(*ProduceRequest)(nil),  // 2: ProduceRequest
	(*ProduceResponse)(nil), // 3: ProduceResponse
}
var file_proto_proto_depIdxs = []int32{
	1, // 0: ProduceRequest.img:type_name -> Image
	0, // 1: ProduceResponse.status:type_name -> Status
	2, // 2: ImageService.Produce:input_type -> ProduceRequest
	3, // 3: ImageService.Produce:output_type -> ProduceResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_proto_init() }
func file_proto_proto_init() {
	if File_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProduceRequest); i {
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
		file_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProduceResponse); i {
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
			RawDescriptor: file_proto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_depIdxs,
		EnumInfos:         file_proto_proto_enumTypes,
		MessageInfos:      file_proto_proto_msgTypes,
	}.Build()
	File_proto_proto = out.File
	file_proto_proto_rawDesc = nil
	file_proto_proto_goTypes = nil
	file_proto_proto_depIdxs = nil
}
