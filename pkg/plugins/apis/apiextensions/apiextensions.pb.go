// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v1.0.0
// source: github.com/rancher/opni/pkg/plugins/apis/apiextensions/apiextensions.proto

package apiextensions

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/anypb"
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

type CertConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ca       string `protobuf:"bytes,1,opt,name=ca,proto3" json:"ca,omitempty"`
	CaData   string `protobuf:"bytes,2,opt,name=caData,proto3" json:"caData,omitempty"`
	Cert     string `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
	CertData string `protobuf:"bytes,4,opt,name=certData,proto3" json:"certData,omitempty"`
	Key      string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	KeyData  string `protobuf:"bytes,6,opt,name=keyData,proto3" json:"keyData,omitempty"`
}

func (x *CertConfig) Reset() {
	*x = CertConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertConfig) ProtoMessage() {}

func (x *CertConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertConfig.ProtoReflect.Descriptor instead.
func (*CertConfig) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{0}
}

func (x *CertConfig) GetCa() string {
	if x != nil {
		return x.Ca
	}
	return ""
}

func (x *CertConfig) GetCaData() string {
	if x != nil {
		return x.CaData
	}
	return ""
}

func (x *CertConfig) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

func (x *CertConfig) GetCertData() string {
	if x != nil {
		return x.CertData
	}
	return ""
}

func (x *CertConfig) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CertConfig) GetKeyData() string {
	if x != nil {
		return x.KeyData
	}
	return ""
}

type GatewayAPIExtensionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpAddr string       `protobuf:"bytes,1,opt,name=httpAddr,proto3" json:"httpAddr,omitempty"`
	Routes   []*RouteInfo `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *GatewayAPIExtensionConfig) Reset() {
	*x = GatewayAPIExtensionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayAPIExtensionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayAPIExtensionConfig) ProtoMessage() {}

func (x *GatewayAPIExtensionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayAPIExtensionConfig.ProtoReflect.Descriptor instead.
func (*GatewayAPIExtensionConfig) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayAPIExtensionConfig) GetHttpAddr() string {
	if x != nil {
		return x.HttpAddr
	}
	return ""
}

func (x *GatewayAPIExtensionConfig) GetRoutes() []*RouteInfo {
	if x != nil {
		return x.Routes
	}
	return nil
}

type ServiceDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceDescriptor *descriptorpb.ServiceDescriptorProto `protobuf:"bytes,1,opt,name=serviceDescriptor,proto3" json:"serviceDescriptor,omitempty"`
	Options           *ServiceOptions                      `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ServiceDescriptor) Reset() {
	*x = ServiceDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDescriptor) ProtoMessage() {}

func (x *ServiceDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDescriptor.ProtoReflect.Descriptor instead.
func (*ServiceDescriptor) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceDescriptor) GetServiceDescriptor() *descriptorpb.ServiceDescriptorProto {
	if x != nil {
		return x.ServiceDescriptor
	}
	return nil
}

func (x *ServiceDescriptor) GetOptions() *ServiceOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type ServiceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, the service will only be available to clusters that have this
	// capability.
	RequireCapability string `protobuf:"bytes,1,opt,name=requireCapability,proto3" json:"requireCapability,omitempty"`
}

func (x *ServiceOptions) Reset() {
	*x = ServiceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOptions) ProtoMessage() {}

func (x *ServiceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOptions.ProtoReflect.Descriptor instead.
func (*ServiceOptions) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceOptions) GetRequireCapability() string {
	if x != nil {
		return x.RequireCapability
	}
	return ""
}

type ServiceDescriptorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Descriptors []*ServiceDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
}

func (x *ServiceDescriptorList) Reset() {
	*x = ServiceDescriptorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDescriptorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDescriptorList) ProtoMessage() {}

func (x *ServiceDescriptorList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDescriptorList.ProtoReflect.Descriptor instead.
func (*ServiceDescriptorList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceDescriptorList) GetDescriptors() []*ServiceDescriptor {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

type RouteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *RouteInfo) Reset() {
	*x = RouteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteInfo) ProtoMessage() {}

func (x *RouteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteInfo.ProtoReflect.Descriptor instead.
func (*RouteInfo) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{5}
}

func (x *RouteInfo) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RouteInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x0a, 0x43, 0x65, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x63, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x22, 0x69, 0x0a, 0x19, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x55, 0x0a, 0x11, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x11, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x12, 0x37, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x5b, 0x0a, 0x15, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x37, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x32,
	0x67, 0x0a, 0x16, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x50, 0x49,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x0a, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x67, 0x0a, 0x13, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x50, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0x19, 0x2e, 0x61,
	0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x65, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41,
	0x50, 0x49, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x32, 0x5e, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x50, 0x49, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x32, 0x67, 0x0a, 0x11, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData = file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc
)

func file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData)
	})
	return file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData
}

var file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_goTypes = []interface{}{
	(*CertConfig)(nil),                          // 0: apiextensions.CertConfig
	(*GatewayAPIExtensionConfig)(nil),           // 1: apiextensions.GatewayAPIExtensionConfig
	(*ServiceDescriptor)(nil),                   // 2: apiextensions.ServiceDescriptor
	(*ServiceOptions)(nil),                      // 3: apiextensions.ServiceOptions
	(*ServiceDescriptorList)(nil),               // 4: apiextensions.ServiceDescriptorList
	(*RouteInfo)(nil),                           // 5: apiextensions.RouteInfo
	(*descriptorpb.ServiceDescriptorProto)(nil), // 6: google.protobuf.ServiceDescriptorProto
	(*emptypb.Empty)(nil),                       // 7: google.protobuf.Empty
}
var file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_depIdxs = []int32{
	5, // 0: apiextensions.GatewayAPIExtensionConfig.routes:type_name -> apiextensions.RouteInfo
	6, // 1: apiextensions.ServiceDescriptor.serviceDescriptor:type_name -> google.protobuf.ServiceDescriptorProto
	3, // 2: apiextensions.ServiceDescriptor.options:type_name -> apiextensions.ServiceOptions
	2, // 3: apiextensions.ServiceDescriptorList.descriptors:type_name -> apiextensions.ServiceDescriptor
	7, // 4: apiextensions.ManagementAPIExtension.Descriptor:input_type -> google.protobuf.Empty
	0, // 5: apiextensions.GatewayAPIExtension.Configure:input_type -> apiextensions.CertConfig
	7, // 6: apiextensions.StreamAPIExtension.Services:input_type -> google.protobuf.Empty
	7, // 7: apiextensions.UnaryAPIExtension.UnaryDescriptor:input_type -> google.protobuf.Empty
	6, // 8: apiextensions.ManagementAPIExtension.Descriptor:output_type -> google.protobuf.ServiceDescriptorProto
	1, // 9: apiextensions.GatewayAPIExtension.Configure:output_type -> apiextensions.GatewayAPIExtensionConfig
	4, // 10: apiextensions.StreamAPIExtension.Services:output_type -> apiextensions.ServiceDescriptorList
	6, // 11: apiextensions.UnaryAPIExtension.UnaryDescriptor:output_type -> google.protobuf.ServiceDescriptorProto
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_init() }
func file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_init() {
	if File_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertConfig); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayAPIExtensionConfig); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDescriptor); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOptions); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDescriptorList); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteInfo); i {
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
			RawDescriptor: file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto = out.File
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc = nil
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_goTypes = nil
	file_github_com_rancher_opni_pkg_plugins_apis_apiextensions_apiextensions_proto_depIdxs = nil
}
