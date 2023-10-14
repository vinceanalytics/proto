// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: vince/cluster/v1/cluster.proto

package v1

import (
	v1 "github.com/vinceanalytics/proto/gen/go/vince/config/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ApplyClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *v1.ClusterConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ApplyClusterRequest) Reset() {
	*x = ApplyClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_cluster_v1_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyClusterRequest) ProtoMessage() {}

func (x *ApplyClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vince_cluster_v1_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyClusterRequest.ProtoReflect.Descriptor instead.
func (*ApplyClusterRequest) Descriptor() ([]byte, []int) {
	return file_vince_cluster_v1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *ApplyClusterRequest) GetConfig() *v1.ClusterConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type ApplyClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok string `protobuf:"bytes,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ApplyClusterResponse) Reset() {
	*x = ApplyClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_cluster_v1_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyClusterResponse) ProtoMessage() {}

func (x *ApplyClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vince_cluster_v1_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyClusterResponse.ProtoReflect.Descriptor instead.
func (*ApplyClusterResponse) Descriptor() ([]byte, []int) {
	return file_vince_cluster_v1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *ApplyClusterResponse) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

type GetClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClusterRequest) Reset() {
	*x = GetClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_cluster_v1_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterRequest) ProtoMessage() {}

func (x *GetClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vince_cluster_v1_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterRequest.ProtoReflect.Descriptor instead.
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return file_vince_cluster_v1_cluster_proto_rawDescGZIP(), []int{2}
}

type GetClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *v1.ClusterConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetClusterResponse) Reset() {
	*x = GetClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_cluster_v1_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterResponse) ProtoMessage() {}

func (x *GetClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vince_cluster_v1_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterResponse.ProtoReflect.Descriptor instead.
func (*GetClusterResponse) Descriptor() ([]byte, []int) {
	return file_vince_cluster_v1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *GetClusterResponse) GetConfig() *v1.ClusterConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_vince_cluster_v1_cluster_proto protoreflect.FileDescriptor

var file_vince_cluster_v1_cluster_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x40, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x26, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x3f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x32, 0xb9, 0x01, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x0c,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x50, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12,
	0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x77, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2,
	0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vince_cluster_v1_cluster_proto_rawDescOnce sync.Once
	file_vince_cluster_v1_cluster_proto_rawDescData = file_vince_cluster_v1_cluster_proto_rawDesc
)

func file_vince_cluster_v1_cluster_proto_rawDescGZIP() []byte {
	file_vince_cluster_v1_cluster_proto_rawDescOnce.Do(func() {
		file_vince_cluster_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_vince_cluster_v1_cluster_proto_rawDescData)
	})
	return file_vince_cluster_v1_cluster_proto_rawDescData
}

var file_vince_cluster_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_vince_cluster_v1_cluster_proto_goTypes = []interface{}{
	(*ApplyClusterRequest)(nil),  // 0: v1.ApplyClusterRequest
	(*ApplyClusterResponse)(nil), // 1: v1.ApplyClusterResponse
	(*GetClusterRequest)(nil),    // 2: v1.GetClusterRequest
	(*GetClusterResponse)(nil),   // 3: v1.GetClusterResponse
	(*v1.ClusterConfig)(nil),     // 4: v1.ClusterConfig
}
var file_vince_cluster_v1_cluster_proto_depIdxs = []int32{
	4, // 0: v1.ApplyClusterRequest.config:type_name -> v1.ClusterConfig
	4, // 1: v1.GetClusterResponse.config:type_name -> v1.ClusterConfig
	0, // 2: v1.Cluster.ApplyCluster:input_type -> v1.ApplyClusterRequest
	2, // 3: v1.Cluster.GetCluster:input_type -> v1.GetClusterRequest
	1, // 4: v1.Cluster.ApplyCluster:output_type -> v1.ApplyClusterResponse
	3, // 5: v1.Cluster.GetCluster:output_type -> v1.GetClusterResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vince_cluster_v1_cluster_proto_init() }
func file_vince_cluster_v1_cluster_proto_init() {
	if File_vince_cluster_v1_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vince_cluster_v1_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyClusterRequest); i {
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
		file_vince_cluster_v1_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyClusterResponse); i {
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
		file_vince_cluster_v1_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterRequest); i {
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
		file_vince_cluster_v1_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterResponse); i {
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
			RawDescriptor: file_vince_cluster_v1_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vince_cluster_v1_cluster_proto_goTypes,
		DependencyIndexes: file_vince_cluster_v1_cluster_proto_depIdxs,
		MessageInfos:      file_vince_cluster_v1_cluster_proto_msgTypes,
	}.Build()
	File_vince_cluster_v1_cluster_proto = out.File
	file_vince_cluster_v1_cluster_proto_rawDesc = nil
	file_vince_cluster_v1_cluster_proto_goTypes = nil
	file_vince_cluster_v1_cluster_proto_depIdxs = nil
}
