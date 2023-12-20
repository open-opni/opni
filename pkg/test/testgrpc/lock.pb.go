// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        (unknown)
// source: github.com/rancher/opni/pkg/test/testgrpc/lock.proto

package testgrpc

import (
	_ "github.com/rancher/opni/internal/codegen/cli"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type LockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string               `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Dur *durationpb.Duration `protobuf:"bytes,2,opt,name=dur,proto3" json:"dur,omitempty"`
}

func (x *LockRequest) Reset() {
	*x = LockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockRequest) ProtoMessage() {}

func (x *LockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockRequest.ProtoReflect.Descriptor instead.
func (*LockRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescGZIP(), []int{0}
}

func (x *LockRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LockRequest) GetDur() *durationpb.Duration {
	if x != nil {
		return x.Dur
	}
	return nil
}

type UnlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UnlockRequest) Reset() {
	*x = UnlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockRequest) ProtoMessage() {}

func (x *UnlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockRequest.ProtoReflect.Descriptor instead.
func (*UnlockRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescGZIP(), []int{1}
}

func (x *UnlockRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type LockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acquired bool   `protobuf:"varint,1,opt,name=acquired,proto3" json:"acquired,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LockResponse) Reset() {
	*x = LockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockResponse) ProtoMessage() {}

func (x *LockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockResponse.ProtoReflect.Descriptor instead.
func (*LockResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescGZIP(), []int{2}
}

func (x *LockResponse) GetAcquired() bool {
	if x != nil {
		return x.Acquired
	}
	return false
}

func (x *LockResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UnlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unlocked bool   `protobuf:"varint,1,opt,name=unlocked,proto3" json:"unlocked,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UnlockResponse) Reset() {
	*x = UnlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockResponse) ProtoMessage() {}

func (x *UnlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockResponse.ProtoReflect.Descriptor instead.
func (*UnlockResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescGZIP(), []int{3}
}

func (x *UnlockResponse) GetUnlocked() bool {
	if x != nil {
		return x.Unlocked
	}
	return false
}

func (x *UnlockResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ListLocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ListLocksResponse) Reset() {
	*x = ListLocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocksResponse) ProtoMessage() {}

func (x *ListLocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocksResponse.ProtoReflect.Descriptor instead.
func (*ListLocksResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescGZIP(), []int{4}
}

func (x *ListLocksResponse) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_github_com_rancher_opni_pkg_test_testgrpc_lock_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x63,
	0x6b, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x03, 0x64, 0x75, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x03, 0x64, 0x75, 0x72, 0x22, 0x21, 0x0a, 0x0d, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x42, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x44, 0x0a, 0x0e, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x27, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x32, 0x83, 0x02, 0x0a, 0x0a, 0x54,
	0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x6f, 0x63,
	0x6b, 0x12, 0x16, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x54, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x31, 0x82, 0xc0, 0x0c, 0x02, 0x08, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e,
	0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescData = file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDesc
)

func file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescData)
	})
	return file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDescData
}

var file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_goTypes = []interface{}{
	(*LockRequest)(nil),         // 0: test.lock.LockRequest
	(*UnlockRequest)(nil),       // 1: test.lock.UnlockRequest
	(*LockResponse)(nil),        // 2: test.lock.LockResponse
	(*UnlockResponse)(nil),      // 3: test.lock.UnlockResponse
	(*ListLocksResponse)(nil),   // 4: test.lock.ListLocksResponse
	(*durationpb.Duration)(nil), // 5: google.protobuf.Duration
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_depIdxs = []int32{
	5, // 0: test.lock.LockRequest.dur:type_name -> google.protobuf.Duration
	0, // 1: test.lock.TestLocker.Lock:input_type -> test.lock.LockRequest
	0, // 2: test.lock.TestLocker.TryLock:input_type -> test.lock.LockRequest
	6, // 3: test.lock.TestLocker.ListLocks:input_type -> google.protobuf.Empty
	1, // 4: test.lock.TestLocker.Unlock:input_type -> test.lock.UnlockRequest
	2, // 5: test.lock.TestLocker.Lock:output_type -> test.lock.LockResponse
	2, // 6: test.lock.TestLocker.TryLock:output_type -> test.lock.LockResponse
	4, // 7: test.lock.TestLocker.ListLocks:output_type -> test.lock.ListLocksResponse
	3, // 8: test.lock.TestLocker.Unlock:output_type -> test.lock.UnlockResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_init() }
func file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_init() {
	if File_github_com_rancher_opni_pkg_test_testgrpc_lock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockRequest); i {
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
		file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockRequest); i {
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
		file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockResponse); i {
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
		file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockResponse); i {
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
		file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocksResponse); i {
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
			RawDescriptor: file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_pkg_test_testgrpc_lock_proto = out.File
	file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_rawDesc = nil
	file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_goTypes = nil
	file_github_com_rancher_opni_pkg_test_testgrpc_lock_proto_depIdxs = nil
}
