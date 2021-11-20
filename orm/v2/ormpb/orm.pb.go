// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: regen/orm/v1alpha1/orm.proto

package ormpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TableDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimaryKey *PrimaryKeyDescriptor       `protobuf:"bytes,1,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	Index      []*SecondaryIndexDescriptor `protobuf:"bytes,2,rep,name=index,proto3" json:"index,omitempty"`
	// id is an integer ID for the table that must be unique within the kv-store
	// used for this ORM instance. It will be deprecated in the future when this
	// can be auto-generated.
	Id uint32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TableDescriptor) Reset() {
	*x = TableDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_orm_v1alpha1_orm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableDescriptor) ProtoMessage() {}

func (x *TableDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_regen_orm_v1alpha1_orm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableDescriptor.ProtoReflect.Descriptor instead.
func (*TableDescriptor) Descriptor() ([]byte, []int) {
	return file_regen_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{0}
}

func (x *TableDescriptor) GetPrimaryKey() *PrimaryKeyDescriptor {
	if x != nil {
		return x.PrimaryKey
	}
	return nil
}

func (x *TableDescriptor) GetIndex() []*SecondaryIndexDescriptor {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *TableDescriptor) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PrimaryKeyDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fields is a comma-separated list of fields in the primary key.
	Fields string `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields,omitempty"`
	// auto_increment specifies that the primary key is generated by an
	// auto-incrementing integer. If this is set to true fields must only
	// contain one field of that is of an integral type.
	AutoIncrement bool `protobuf:"varint,2,opt,name=auto_increment,json=autoIncrement,proto3" json:"auto_increment,omitempty"`
}

func (x *PrimaryKeyDescriptor) Reset() {
	*x = PrimaryKeyDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_orm_v1alpha1_orm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryKeyDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryKeyDescriptor) ProtoMessage() {}

func (x *PrimaryKeyDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_regen_orm_v1alpha1_orm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryKeyDescriptor.ProtoReflect.Descriptor instead.
func (*PrimaryKeyDescriptor) Descriptor() ([]byte, []int) {
	return file_regen_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{1}
}

func (x *PrimaryKeyDescriptor) GetFields() string {
	if x != nil {
		return x.Fields
	}
	return ""
}

func (x *PrimaryKeyDescriptor) GetAutoIncrement() bool {
	if x != nil {
		return x.AutoIncrement
	}
	return false
}

type SecondaryIndexDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fields is a comma-separated list of fields in the index.
	Fields string `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields,omitempty"`
	// id is an integer ID for the index within the indexes for this table.
	// It will be deprecated in the future when this can be auto-generated.
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SecondaryIndexDescriptor) Reset() {
	*x = SecondaryIndexDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_orm_v1alpha1_orm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecondaryIndexDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecondaryIndexDescriptor) ProtoMessage() {}

func (x *SecondaryIndexDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_regen_orm_v1alpha1_orm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecondaryIndexDescriptor.ProtoReflect.Descriptor instead.
func (*SecondaryIndexDescriptor) Descriptor() ([]byte, []int) {
	return file_regen_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{2}
}

func (x *SecondaryIndexDescriptor) GetFields() string {
	if x != nil {
		return x.Fields
	}
	return ""
}

func (x *SecondaryIndexDescriptor) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SingletonDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is an integer ID for the singleton that must be unique within the kv-store
	// used for this ORM instance. It will be deprecated in the future when this
	// can be auto-generated.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SingletonDescriptor) Reset() {
	*x = SingletonDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_orm_v1alpha1_orm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingletonDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingletonDescriptor) ProtoMessage() {}

func (x *SingletonDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_regen_orm_v1alpha1_orm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingletonDescriptor.ProtoReflect.Descriptor instead.
func (*SingletonDescriptor) Descriptor() ([]byte, []int) {
	return file_regen_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{3}
}

func (x *SingletonDescriptor) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var file_regen_orm_v1alpha1_orm_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*TableDescriptor)(nil),
		Field:         100000000,
		Name:          "regen.orm.v1alpha1.table",
		Tag:           "bytes,100000000,opt,name=table",
		Filename:      "regen/orm/v1alpha1/orm.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*SingletonDescriptor)(nil),
		Field:         100000001,
		Name:          "regen.orm.v1alpha1.singleton",
		Tag:           "bytes,100000001,opt,name=singleton",
		Filename:      "regen/orm/v1alpha1/orm.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional regen.orm.v1alpha1.TableDescriptor table = 100000000;
	E_Table = &file_regen_orm_v1alpha1_orm_proto_extTypes[0]
	// optional regen.orm.v1alpha1.SingletonDescriptor singleton = 100000001;
	E_Singleton = &file_regen_orm_v1alpha1_orm_proto_extTypes[1]
)

var File_regen_orm_v1alpha1_orm_proto protoreflect.FileDescriptor

var file_regen_orm_v1alpha1_orm_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x49, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x14, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x5f,
	0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x61, 0x75, 0x74, 0x6f, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x42,
	0x0a, 0x18, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x5d, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x80, 0xc2, 0xd7, 0x2f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x3a, 0x69, 0x0a, 0x09, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x81, 0xc2, 0xd7, 0x2f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x74, 0x6f, 0x6e, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x6d,
	0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_orm_v1alpha1_orm_proto_rawDescOnce sync.Once
	file_regen_orm_v1alpha1_orm_proto_rawDescData = file_regen_orm_v1alpha1_orm_proto_rawDesc
)

func file_regen_orm_v1alpha1_orm_proto_rawDescGZIP() []byte {
	file_regen_orm_v1alpha1_orm_proto_rawDescOnce.Do(func() {
		file_regen_orm_v1alpha1_orm_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_orm_v1alpha1_orm_proto_rawDescData)
	})
	return file_regen_orm_v1alpha1_orm_proto_rawDescData
}

var file_regen_orm_v1alpha1_orm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_regen_orm_v1alpha1_orm_proto_goTypes = []interface{}{
	(*TableDescriptor)(nil),             // 0: regen.orm.v1alpha1.TableDescriptor
	(*PrimaryKeyDescriptor)(nil),        // 1: regen.orm.v1alpha1.PrimaryKeyDescriptor
	(*SecondaryIndexDescriptor)(nil),    // 2: regen.orm.v1alpha1.SecondaryIndexDescriptor
	(*SingletonDescriptor)(nil),         // 3: regen.orm.v1alpha1.SingletonDescriptor
	(*descriptorpb.MessageOptions)(nil), // 4: google.protobuf.MessageOptions
}
var file_regen_orm_v1alpha1_orm_proto_depIdxs = []int32{
	1, // 0: regen.orm.v1alpha1.TableDescriptor.primary_key:type_name -> regen.orm.v1alpha1.PrimaryKeyDescriptor
	2, // 1: regen.orm.v1alpha1.TableDescriptor.index:type_name -> regen.orm.v1alpha1.SecondaryIndexDescriptor
	4, // 2: regen.orm.v1alpha1.table:extendee -> google.protobuf.MessageOptions
	4, // 3: regen.orm.v1alpha1.singleton:extendee -> google.protobuf.MessageOptions
	0, // 4: regen.orm.v1alpha1.table:type_name -> regen.orm.v1alpha1.TableDescriptor
	3, // 5: regen.orm.v1alpha1.singleton:type_name -> regen.orm.v1alpha1.SingletonDescriptor
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	4, // [4:6] is the sub-list for extension type_name
	2, // [2:4] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_regen_orm_v1alpha1_orm_proto_init() }
func file_regen_orm_v1alpha1_orm_proto_init() {
	if File_regen_orm_v1alpha1_orm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_orm_v1alpha1_orm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableDescriptor); i {
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
		file_regen_orm_v1alpha1_orm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryKeyDescriptor); i {
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
		file_regen_orm_v1alpha1_orm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecondaryIndexDescriptor); i {
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
		file_regen_orm_v1alpha1_orm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingletonDescriptor); i {
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
			RawDescriptor: file_regen_orm_v1alpha1_orm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_regen_orm_v1alpha1_orm_proto_goTypes,
		DependencyIndexes: file_regen_orm_v1alpha1_orm_proto_depIdxs,
		MessageInfos:      file_regen_orm_v1alpha1_orm_proto_msgTypes,
		ExtensionInfos:    file_regen_orm_v1alpha1_orm_proto_extTypes,
	}.Build()
	File_regen_orm_v1alpha1_orm_proto = out.File
	file_regen_orm_v1alpha1_orm_proto_rawDesc = nil
	file_regen_orm_v1alpha1_orm_proto_goTypes = nil
	file_regen_orm_v1alpha1_orm_proto_depIdxs = nil
}