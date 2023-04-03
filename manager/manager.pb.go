// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/manager.proto

package manager

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description        string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status             string               `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Owner              string               `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	StartTime          *timestamp.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Datasets           []string             `protobuf:"bytes,8,rep,name=datasets,proto3" json:"datasets,omitempty"`
	Algorithms         []string             `protobuf:"bytes,9,rep,name=algorithms,proto3" json:"algorithms,omitempty"`
	DatasetProviders   []string             `protobuf:"bytes,10,rep,name=dataset_providers,json=datasetProviders,proto3" json:"dataset_providers,omitempty"`
	AlgorithmProviders []string             `protobuf:"bytes,11,rep,name=algorithm_providers,json=algorithmProviders,proto3" json:"algorithm_providers,omitempty"`
	ResultConsumers    []string             `protobuf:"bytes,13,rep,name=result_consumers,json=resultConsumers,proto3" json:"result_consumers,omitempty"`
	Ttl                int32                `protobuf:"varint,12,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_proto_manager_proto_rawDescGZIP(), []int{0}
}

func (x *RunRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RunRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RunRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RunRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *RunRequest) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *RunRequest) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *RunRequest) GetDatasets() []string {
	if x != nil {
		return x.Datasets
	}
	return nil
}

func (x *RunRequest) GetAlgorithms() []string {
	if x != nil {
		return x.Algorithms
	}
	return nil
}

func (x *RunRequest) GetDatasetProviders() []string {
	if x != nil {
		return x.DatasetProviders
	}
	return nil
}

func (x *RunRequest) GetAlgorithmProviders() []string {
	if x != nil {
		return x.AlgorithmProviders
	}
	return nil
}

func (x *RunRequest) GetResultConsumers() []string {
	if x != nil {
		return x.ResultConsumers
	}
	return nil
}

func (x *RunRequest) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_proto_manager_proto_rawDescGZIP(), []int{1}
}

func (x *RunResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type CreateDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pool   string `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	Volume string `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *CreateDomainRequest) Reset() {
	*x = CreateDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDomainRequest) ProtoMessage() {}

func (x *CreateDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDomainRequest.ProtoReflect.Descriptor instead.
func (*CreateDomainRequest) Descriptor() ([]byte, []int) {
	return file_proto_manager_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDomainRequest) GetPool() string {
	if x != nil {
		return x.Pool
	}
	return ""
}

func (x *CreateDomainRequest) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *CreateDomainRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type CreateDomainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateDomainResponse) Reset() {
	*x = CreateDomainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDomainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDomainResponse) ProtoMessage() {}

func (x *CreateDomainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDomainResponse.ProtoReflect.Descriptor instead.
func (*CreateDomainResponse) Descriptor() ([]byte, []int) {
	return file_proto_manager_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDomainResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_manager_proto protoreflect.FileDescriptor

var file_proto_manager_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x03, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x12, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74,
	0x6c, 0x22, 0x1d, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x59, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x2a, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xab, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x52, 0x75,
	0x6e, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x22, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_manager_proto_rawDescOnce sync.Once
	file_proto_manager_proto_rawDescData = file_proto_manager_proto_rawDesc
)

func file_proto_manager_proto_rawDescGZIP() []byte {
	file_proto_manager_proto_rawDescOnce.Do(func() {
		file_proto_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_manager_proto_rawDescData)
	})
	return file_proto_manager_proto_rawDescData
}

var file_proto_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_manager_proto_goTypes = []interface{}{
	(*RunRequest)(nil),           // 0: manager_proto.RunRequest
	(*RunResponse)(nil),          // 1: manager_proto.RunResponse
	(*CreateDomainRequest)(nil),  // 2: manager_proto.CreateDomainRequest
	(*CreateDomainResponse)(nil), // 3: manager_proto.CreateDomainResponse
	(*timestamp.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_proto_manager_proto_depIdxs = []int32{
	4, // 0: manager_proto.RunRequest.start_time:type_name -> google.protobuf.Timestamp
	4, // 1: manager_proto.RunRequest.end_time:type_name -> google.protobuf.Timestamp
	0, // 2: manager_proto.ManagerService.Run:input_type -> manager_proto.RunRequest
	2, // 3: manager_proto.ManagerService.CreateDomain:input_type -> manager_proto.CreateDomainRequest
	1, // 4: manager_proto.ManagerService.Run:output_type -> manager_proto.RunResponse
	3, // 5: manager_proto.ManagerService.CreateDomain:output_type -> manager_proto.CreateDomainResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_manager_proto_init() }
func file_proto_manager_proto_init() {
	if File_proto_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunRequest); i {
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
		file_proto_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
		file_proto_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDomainRequest); i {
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
		file_proto_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDomainResponse); i {
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
			RawDescriptor: file_proto_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_manager_proto_goTypes,
		DependencyIndexes: file_proto_manager_proto_depIdxs,
		MessageInfos:      file_proto_manager_proto_msgTypes,
	}.Build()
	File_proto_manager_proto = out.File
	file_proto_manager_proto_rawDesc = nil
	file_proto_manager_proto_goTypes = nil
	file_proto_manager_proto_depIdxs = nil
}
