// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: client/api/system/v1alpha1/api.proto

package v1alpha1

import (
	context "context"
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

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/ns-winsvc-service_status#members
type ServiceStatus int32

const (
	ServiceStatus_UNKNOWN          ServiceStatus = 0
	ServiceStatus_STOPPED          ServiceStatus = 1
	ServiceStatus_START_PENDING    ServiceStatus = 2
	ServiceStatus_STOP_PENDING     ServiceStatus = 3
	ServiceStatus_RUNNING          ServiceStatus = 4
	ServiceStatus_CONTINUE_PENDING ServiceStatus = 5
	ServiceStatus_PAUSE_PENDING    ServiceStatus = 6
	ServiceStatus_PAUSED           ServiceStatus = 7
)

// Enum value maps for ServiceStatus.
var (
	ServiceStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "STOPPED",
		2: "START_PENDING",
		3: "STOP_PENDING",
		4: "RUNNING",
		5: "CONTINUE_PENDING",
		6: "PAUSE_PENDING",
		7: "PAUSED",
	}
	ServiceStatus_value = map[string]int32{
		"UNKNOWN":          0,
		"STOPPED":          1,
		"START_PENDING":    2,
		"STOP_PENDING":     3,
		"RUNNING":          4,
		"CONTINUE_PENDING": 5,
		"PAUSE_PENDING":    6,
		"PAUSED":           7,
	}
)

func (x ServiceStatus) Enum() *ServiceStatus {
	p := new(ServiceStatus)
	*p = x
	return p
}

func (x ServiceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_client_api_system_v1alpha1_api_proto_enumTypes[0].Descriptor()
}

func (ServiceStatus) Type() protoreflect.EnumType {
	return &file_client_api_system_v1alpha1_api_proto_enumTypes[0]
}

func (x ServiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceStatus.Descriptor instead.
func (ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_client_api_system_v1alpha1_api_proto_rawDescGZIP(), []int{0}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-changeserviceconfiga
type StartType int32

const (
	StartType_BOOT      StartType = 0
	StartType_SYSTEM    StartType = 1
	StartType_AUTOMATIC StartType = 2
	StartType_MANUAL    StartType = 3
	StartType_DISABLED  StartType = 4
)

// Enum value maps for StartType.
var (
	StartType_name = map[int32]string{
		0: "BOOT",
		1: "SYSTEM",
		2: "AUTOMATIC",
		3: "MANUAL",
		4: "DISABLED",
	}
	StartType_value = map[string]int32{
		"BOOT":      0,
		"SYSTEM":    1,
		"AUTOMATIC": 2,
		"MANUAL":    3,
		"DISABLED":  4,
	}
)

func (x StartType) Enum() *StartType {
	p := new(StartType)
	*p = x
	return p
}

func (x StartType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StartType) Descriptor() protoreflect.EnumDescriptor {
	return file_client_api_system_v1alpha1_api_proto_enumTypes[1].Descriptor()
}

func (StartType) Type() protoreflect.EnumType {
	return &file_client_api_system_v1alpha1_api_proto_enumTypes[1]
}

func (x StartType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StartType.Descriptor instead.
func (StartType) EnumDescriptor() ([]byte, []int) {
	return file_client_api_system_v1alpha1_api_proto_rawDescGZIP(), []int{1}
}

type GetBIOSSerialNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBIOSSerialNumberRequest) Reset() {
	*x = GetBIOSSerialNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBIOSSerialNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBIOSSerialNumberRequest) ProtoMessage() {}

func (x *GetBIOSSerialNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBIOSSerialNumberRequest.ProtoReflect.Descriptor instead.
func (*GetBIOSSerialNumberRequest) Descriptor() ([]byte, []int) {
	return file_client_api_system_v1alpha1_api_proto_rawDescGZIP(), []int{0}
}

type GetBIOSSerialNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Serial number
	SerialNumber string `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
}

func (x *GetBIOSSerialNumberResponse) Reset() {
	*x = GetBIOSSerialNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBIOSSerialNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBIOSSerialNumberResponse) ProtoMessage() {}

func (x *GetBIOSSerialNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBIOSSerialNumberResponse.ProtoReflect.Descriptor instead.
func (*GetBIOSSerialNumberResponse) Descriptor() ([]byte, []int) {
	return file_client_api_system_v1alpha1_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetBIOSSerialNumberResponse) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

type StartServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service name (as listed in System\CCS\Services keys)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StartServiceRequest) Reset() {
	*x = StartServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartServiceRequest) ProtoMessage() {}

func (x *StartServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartServiceRequest.ProtoReflect.Descriptor instead.
func (*StartServiceRequest) Descriptor() ([]byte, []int) {
	return file_client_api_system_v1alpha1_api_proto_rawDescGZIP(), []int{2}
}

func (x *StartServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StartServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartServiceResponse) Reset() {
	*x = StartServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartServiceResponse) ProtoMessage() {}

func (x *StartServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartServiceResponse.ProtoReflect.Descriptor instead.
func (*StartServiceResponse) Descriptor() ([]byte, []int) {
	return file_client_api_system_v1alpha1_api_proto_rawDescGZIP(), []int{3}
}

type StopServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service name (as listed in System\CCS\Services keys)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Forces stopping of services that has dependent services
	Force bool `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *StopServiceRequest) Reset() {
	*x = StopServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopServiceRequest) ProtoMessage() {}

func (x *StopServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopServiceRequest.ProtoReflect.Descriptor instead.
func (*StopServiceRequest) Descriptor() ([]byte, []int) {
	return file_client_api_system_v1alpha1_api_proto_rawDescGZIP(), []int{4}
}

func (x *StopServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StopServiceRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type StopServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopServiceResponse) Reset() {
	*x = StopServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopServiceResponse) ProtoMessage() {}

func (x *StopServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopServiceResponse.ProtoReflect.Descriptor instead.
func (*StopServiceResponse) Descriptor() ([]byte, []int) {
	return file_client_api_system_v1alpha1_api_proto_rawDescGZIP(), []int{5}
}

type GetServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service name (as listed in System\CCS\Services keys)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetServiceRequest) Reset() {
	*x = GetServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceRequest) ProtoMessage() {}

func (x *GetServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceRequest.ProtoReflect.Descriptor instead.
func (*GetServiceRequest) Descriptor() ([]byte, []int) {
	return file_client_api_system_v1alpha1_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service display name
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Service start type.
	// Used to control whether a service will start on boot, and if so on which
	// boot phase.
	StartType StartType `protobuf:"varint,2,opt,name=start_type,json=startType,proto3,enum=v1alpha1.StartType" json:"start_type,omitempty"`
	// Service status, e.g. stopped, running, paused
	Status ServiceStatus `protobuf:"varint,3,opt,name=status,proto3,enum=v1alpha1.ServiceStatus" json:"status,omitempty"`
}

func (x *GetServiceResponse) Reset() {
	*x = GetServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceResponse) ProtoMessage() {}

func (x *GetServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_api_system_v1alpha1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceResponse.ProtoReflect.Descriptor instead.
func (*GetServiceResponse) Descriptor() ([]byte, []int) {
	return file_client_api_system_v1alpha1_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetServiceResponse) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *GetServiceResponse) GetStartType() StartType {
	if x != nil {
		return x.StartType
	}
	return StartType_BOOT
}

func (x *GetServiceResponse) GetStatus() ServiceStatus {
	if x != nil {
		return x.Status
	}
	return ServiceStatus_UNKNOWN
}

var File_client_api_system_v1alpha1_api_proto protoreflect.FileDescriptor

var file_client_api_system_v1alpha1_api_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x42, 0x49, 0x4f, 0x53, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x42, 0x49, 0x4f, 0x53, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x29, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a,
	0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x32, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2a, 0x90, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x5f, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x41, 0x55, 0x53,
	0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x50,
	0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x07, 0x2a, 0x4a, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x54, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x55,
	0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4e,
	0x55, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45,
	0x44, 0x10, 0x04, 0x32, 0xd8, 0x02, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x64,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x49, 0x4f, 0x53, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x49, 0x4f, 0x53, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x49, 0x4f, 0x53, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74,
	0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x40,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x63, 0x73, 0x69, 0x2f, 0x63, 0x73, 0x69, 0x2d,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_api_system_v1alpha1_api_proto_rawDescOnce sync.Once
	file_client_api_system_v1alpha1_api_proto_rawDescData = file_client_api_system_v1alpha1_api_proto_rawDesc
)

func file_client_api_system_v1alpha1_api_proto_rawDescGZIP() []byte {
	file_client_api_system_v1alpha1_api_proto_rawDescOnce.Do(func() {
		file_client_api_system_v1alpha1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_api_system_v1alpha1_api_proto_rawDescData)
	})
	return file_client_api_system_v1alpha1_api_proto_rawDescData
}

var file_client_api_system_v1alpha1_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_client_api_system_v1alpha1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_client_api_system_v1alpha1_api_proto_goTypes = []interface{}{
	(ServiceStatus)(0),                  // 0: v1alpha1.ServiceStatus
	(StartType)(0),                      // 1: v1alpha1.StartType
	(*GetBIOSSerialNumberRequest)(nil),  // 2: v1alpha1.GetBIOSSerialNumberRequest
	(*GetBIOSSerialNumberResponse)(nil), // 3: v1alpha1.GetBIOSSerialNumberResponse
	(*StartServiceRequest)(nil),         // 4: v1alpha1.StartServiceRequest
	(*StartServiceResponse)(nil),        // 5: v1alpha1.StartServiceResponse
	(*StopServiceRequest)(nil),          // 6: v1alpha1.StopServiceRequest
	(*StopServiceResponse)(nil),         // 7: v1alpha1.StopServiceResponse
	(*GetServiceRequest)(nil),           // 8: v1alpha1.GetServiceRequest
	(*GetServiceResponse)(nil),          // 9: v1alpha1.GetServiceResponse
}
var file_client_api_system_v1alpha1_api_proto_depIdxs = []int32{
	1, // 0: v1alpha1.GetServiceResponse.start_type:type_name -> v1alpha1.StartType
	0, // 1: v1alpha1.GetServiceResponse.status:type_name -> v1alpha1.ServiceStatus
	2, // 2: v1alpha1.System.GetBIOSSerialNumber:input_type -> v1alpha1.GetBIOSSerialNumberRequest
	4, // 3: v1alpha1.System.StartService:input_type -> v1alpha1.StartServiceRequest
	6, // 4: v1alpha1.System.StopService:input_type -> v1alpha1.StopServiceRequest
	8, // 5: v1alpha1.System.GetService:input_type -> v1alpha1.GetServiceRequest
	3, // 6: v1alpha1.System.GetBIOSSerialNumber:output_type -> v1alpha1.GetBIOSSerialNumberResponse
	5, // 7: v1alpha1.System.StartService:output_type -> v1alpha1.StartServiceResponse
	7, // 8: v1alpha1.System.StopService:output_type -> v1alpha1.StopServiceResponse
	9, // 9: v1alpha1.System.GetService:output_type -> v1alpha1.GetServiceResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_client_api_system_v1alpha1_api_proto_init() }
func file_client_api_system_v1alpha1_api_proto_init() {
	if File_client_api_system_v1alpha1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_api_system_v1alpha1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBIOSSerialNumberRequest); i {
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
		file_client_api_system_v1alpha1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBIOSSerialNumberResponse); i {
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
		file_client_api_system_v1alpha1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartServiceRequest); i {
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
		file_client_api_system_v1alpha1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartServiceResponse); i {
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
		file_client_api_system_v1alpha1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopServiceRequest); i {
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
		file_client_api_system_v1alpha1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopServiceResponse); i {
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
		file_client_api_system_v1alpha1_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceRequest); i {
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
		file_client_api_system_v1alpha1_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceResponse); i {
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
			RawDescriptor: file_client_api_system_v1alpha1_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_api_system_v1alpha1_api_proto_goTypes,
		DependencyIndexes: file_client_api_system_v1alpha1_api_proto_depIdxs,
		EnumInfos:         file_client_api_system_v1alpha1_api_proto_enumTypes,
		MessageInfos:      file_client_api_system_v1alpha1_api_proto_msgTypes,
	}.Build()
	File_client_api_system_v1alpha1_api_proto = out.File
	file_client_api_system_v1alpha1_api_proto_rawDesc = nil
	file_client_api_system_v1alpha1_api_proto_goTypes = nil
	file_client_api_system_v1alpha1_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SystemClient is the client API for System service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SystemClient interface {
	// GetBIOSSerialNumber returns the device's serial number
	GetBIOSSerialNumber(ctx context.Context, in *GetBIOSSerialNumberRequest, opts ...grpc.CallOption) (*GetBIOSSerialNumberResponse, error)
	// StartService starts a Windows service
	// NOTE: This method affects global node state and should only be used
	//
	//	with consideration to other CSI drivers that run concurrently.
	StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceResponse, error)
	// StopService stops a Windows service
	// NOTE: This method affects global node state and should only be used
	//
	//	with consideration to other CSI drivers that run concurrently.
	StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*StopServiceResponse, error)
	// GetService queries a Windows service state
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
}

type systemClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemClient(cc grpc.ClientConnInterface) SystemClient {
	return &systemClient{cc}
}

func (c *systemClient) GetBIOSSerialNumber(ctx context.Context, in *GetBIOSSerialNumberRequest, opts ...grpc.CallOption) (*GetBIOSSerialNumberResponse, error) {
	out := new(GetBIOSSerialNumberResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.System/GetBIOSSerialNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceResponse, error) {
	out := new(StartServiceResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.System/StartService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*StopServiceResponse, error) {
	out := new(StopServiceResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.System/StopService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.System/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServer is the server API for System service.
type SystemServer interface {
	// GetBIOSSerialNumber returns the device's serial number
	GetBIOSSerialNumber(context.Context, *GetBIOSSerialNumberRequest) (*GetBIOSSerialNumberResponse, error)
	// StartService starts a Windows service
	// NOTE: This method affects global node state and should only be used
	//
	//	with consideration to other CSI drivers that run concurrently.
	StartService(context.Context, *StartServiceRequest) (*StartServiceResponse, error)
	// StopService stops a Windows service
	// NOTE: This method affects global node state and should only be used
	//
	//	with consideration to other CSI drivers that run concurrently.
	StopService(context.Context, *StopServiceRequest) (*StopServiceResponse, error)
	// GetService queries a Windows service state
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
}

// UnimplementedSystemServer can be embedded to have forward compatible implementations.
type UnimplementedSystemServer struct {
}

func (*UnimplementedSystemServer) GetBIOSSerialNumber(context.Context, *GetBIOSSerialNumberRequest) (*GetBIOSSerialNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBIOSSerialNumber not implemented")
}
func (*UnimplementedSystemServer) StartService(context.Context, *StartServiceRequest) (*StartServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartService not implemented")
}
func (*UnimplementedSystemServer) StopService(context.Context, *StopServiceRequest) (*StopServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopService not implemented")
}
func (*UnimplementedSystemServer) GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}

func RegisterSystemServer(s *grpc.Server, srv SystemServer) {
	s.RegisterService(&_System_serviceDesc, srv)
}

func _System_GetBIOSSerialNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBIOSSerialNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).GetBIOSSerialNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.System/GetBIOSSerialNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).GetBIOSSerialNumber(ctx, req.(*GetBIOSSerialNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.System/StartService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).StartService(ctx, req.(*StartServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_StopService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).StopService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.System/StopService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).StopService(ctx, req.(*StopServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.System/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _System_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.System",
	HandlerType: (*SystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBIOSSerialNumber",
			Handler:    _System_GetBIOSSerialNumber_Handler,
		},
		{
			MethodName: "StartService",
			Handler:    _System_StartService_Handler,
		},
		{
			MethodName: "StopService",
			Handler:    _System_StopService_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _System_GetService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/api/system/v1alpha1/api.proto",
}
