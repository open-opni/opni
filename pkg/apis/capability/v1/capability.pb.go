// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v1.0.0
// source: github.com/rancher/opni/pkg/apis/capability/v1/capability.proto

package v1

import (
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CapabilityName string `protobuf:"bytes,1,opt,name=capabilityName,proto3" json:"capabilityName,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescGZIP(), []int{0}
}

func (x *InfoResponse) GetCapabilityName() string {
	if x != nil {
		return x.CapabilityName
	}
	return ""
}

type InstallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *v1.Reference `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *InstallRequest) Reset() {
	*x = InstallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallRequest) ProtoMessage() {}

func (x *InstallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallRequest.ProtoReflect.Descriptor instead.
func (*InstallRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescGZIP(), []int{1}
}

func (x *InstallRequest) GetCluster() *v1.Reference {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type UninstallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *v1.Reference `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Options string        `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *UninstallRequest) Reset() {
	*x = UninstallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UninstallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UninstallRequest) ProtoMessage() {}

func (x *UninstallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UninstallRequest.ProtoReflect.Descriptor instead.
func (*UninstallRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescGZIP(), []int{2}
}

func (x *UninstallRequest) GetCluster() *v1.Reference {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *UninstallRequest) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

type InstallerTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template string `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *InstallerTemplateResponse) Reset() {
	*x = InstallerTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallerTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallerTemplateResponse) ProtoMessage() {}

func (x *InstallerTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallerTemplateResponse.ProtoReflect.Descriptor instead.
func (*InstallerTemplateResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescGZIP(), []int{3}
}

func (x *InstallerTemplateResponse) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

var File_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f,
	0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x36, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x22, 0x57, 0x0a, 0x10, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x37, 0x0a,
	0x19, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x32, 0xc9, 0x03, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x12, 0x38, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a,
	0x43, 0x61, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x07, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x09, 0x55, 0x6e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2e, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x0f,
	0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x55, 0x6e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52,
	0x0a, 0x11, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x63, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescData = file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDesc
)

func file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescData)
	})
	return file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDescData
}

var file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_goTypes = []interface{}{
	(*InfoResponse)(nil),              // 0: capability.InfoResponse
	(*InstallRequest)(nil),            // 1: capability.InstallRequest
	(*UninstallRequest)(nil),          // 2: capability.UninstallRequest
	(*InstallerTemplateResponse)(nil), // 3: capability.InstallerTemplateResponse
	(*v1.Reference)(nil),              // 4: core.Reference
	(*emptypb.Empty)(nil),             // 5: google.protobuf.Empty
	(*v1.TaskStatus)(nil),             // 6: core.TaskStatus
}
var file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_depIdxs = []int32{
	4, // 0: capability.InstallRequest.cluster:type_name -> core.Reference
	4, // 1: capability.UninstallRequest.cluster:type_name -> core.Reference
	5, // 2: capability.Backend.Info:input_type -> google.protobuf.Empty
	5, // 3: capability.Backend.CanInstall:input_type -> google.protobuf.Empty
	1, // 4: capability.Backend.Install:input_type -> capability.InstallRequest
	2, // 5: capability.Backend.Uninstall:input_type -> capability.UninstallRequest
	4, // 6: capability.Backend.UninstallStatus:input_type -> core.Reference
	4, // 7: capability.Backend.CancelUninstall:input_type -> core.Reference
	5, // 8: capability.Backend.InstallerTemplate:input_type -> google.protobuf.Empty
	0, // 9: capability.Backend.Info:output_type -> capability.InfoResponse
	5, // 10: capability.Backend.CanInstall:output_type -> google.protobuf.Empty
	5, // 11: capability.Backend.Install:output_type -> google.protobuf.Empty
	5, // 12: capability.Backend.Uninstall:output_type -> google.protobuf.Empty
	6, // 13: capability.Backend.UninstallStatus:output_type -> core.TaskStatus
	5, // 14: capability.Backend.CancelUninstall:output_type -> google.protobuf.Empty
	3, // 15: capability.Backend.InstallerTemplate:output_type -> capability.InstallerTemplateResponse
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_init() }
func file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_init() {
	if File_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallRequest); i {
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
		file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UninstallRequest); i {
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
		file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallerTemplateResponse); i {
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
			RawDescriptor: file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto = out.File
	file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_rawDesc = nil
	file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_goTypes = nil
	file_github_com_rancher_opni_pkg_apis_capability_v1_capability_proto_depIdxs = nil
}