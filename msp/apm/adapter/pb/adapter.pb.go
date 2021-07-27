// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: adapter.proto

package pb

import (
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdapterLanguage int32

const (
	AdapterLanguage_JAVA    AdapterLanguage = 0
	AdapterLanguage_GO      AdapterLanguage = 1
	AdapterLanguage_PHP     AdapterLanguage = 2
	AdapterLanguage_DOT_NET AdapterLanguage = 3
	AdapterLanguage_NODEJS  AdapterLanguage = 4
)

// Enum value maps for AdapterLanguage.
var (
	AdapterLanguage_name = map[int32]string{
		0: "JAVA",
		1: "GO",
		2: "PHP",
		3: "DOT_NET",
		4: "NODEJS",
	}
	AdapterLanguage_value = map[string]int32{
		"JAVA":    0,
		"GO":      1,
		"PHP":     2,
		"DOT_NET": 3,
		"NODEJS":  4,
	}
)

func (x AdapterLanguage) Enum() *AdapterLanguage {
	p := new(AdapterLanguage)
	*p = x
	return p
}

func (x AdapterLanguage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdapterLanguage) Descriptor() protoreflect.EnumDescriptor {
	return file_adapter_proto_enumTypes[0].Descriptor()
}

func (AdapterLanguage) Type() protoreflect.EnumType {
	return &file_adapter_proto_enumTypes[0]
}

func (x AdapterLanguage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdapterLanguage.Descriptor instead.
func (AdapterLanguage) EnumDescriptor() ([]byte, []int) {
	return file_adapter_proto_rawDescGZIP(), []int{0}
}

type AdapterStrategyKey int32

const (
	AdapterStrategyKey_JAVA_AGENT        AdapterStrategyKey = 0
	AdapterStrategyKey_APACHE_SKYWALKING AdapterStrategyKey = 1
	AdapterStrategyKey_JAEGER            AdapterStrategyKey = 2
	AdapterStrategyKey_OPEN_TELEMETRY    AdapterStrategyKey = 3
	AdapterStrategyKey_NODEJS_AGENT      AdapterStrategyKey = 4
)

// Enum value maps for AdapterStrategyKey.
var (
	AdapterStrategyKey_name = map[int32]string{
		0: "JAVA_AGENT",
		1: "APACHE_SKYWALKING",
		2: "JAEGER",
		3: "OPEN_TELEMETRY",
		4: "NODEJS_AGENT",
	}
	AdapterStrategyKey_value = map[string]int32{
		"JAVA_AGENT":        0,
		"APACHE_SKYWALKING": 1,
		"JAEGER":            2,
		"OPEN_TELEMETRY":    3,
		"NODEJS_AGENT":      4,
	}
)

func (x AdapterStrategyKey) Enum() *AdapterStrategyKey {
	p := new(AdapterStrategyKey)
	*p = x
	return p
}

func (x AdapterStrategyKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdapterStrategyKey) Descriptor() protoreflect.EnumDescriptor {
	return file_adapter_proto_enumTypes[1].Descriptor()
}

func (AdapterStrategyKey) Type() protoreflect.EnumType {
	return &file_adapter_proto_enumTypes[1]
}

func (x AdapterStrategyKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdapterStrategyKey.Descriptor instead.
func (AdapterStrategyKey) EnumDescriptor() ([]byte, []int) {
	return file_adapter_proto_rawDescGZIP(), []int{1}
}

type GetAdaptersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID int64 `protobuf:"varint,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
}

func (x *GetAdaptersRequest) Reset() {
	*x = GetAdaptersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdaptersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdaptersRequest) ProtoMessage() {}

func (x *GetAdaptersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adapter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdaptersRequest.ProtoReflect.Descriptor instead.
func (*GetAdaptersRequest) Descriptor() ([]byte, []int) {
	return file_adapter_proto_rawDescGZIP(), []int{0}
}

func (x *GetAdaptersRequest) GetProjectID() int64 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

type GetAdaptersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Adapters `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAdaptersResponse) Reset() {
	*x = GetAdaptersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdaptersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdaptersResponse) ProtoMessage() {}

func (x *GetAdaptersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adapter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdaptersResponse.ProtoReflect.Descriptor instead.
func (*GetAdaptersResponse) Descriptor() ([]byte, []int) {
	return file_adapter_proto_rawDescGZIP(), []int{1}
}

func (x *GetAdaptersResponse) GetData() []*Adapters {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAdapterDocsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID   int64  `protobuf:"varint,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
	ServiceName string `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetAdapterDocsRequest) Reset() {
	*x = GetAdapterDocsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdapterDocsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdapterDocsRequest) ProtoMessage() {}

func (x *GetAdapterDocsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adapter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdapterDocsRequest.ProtoReflect.Descriptor instead.
func (*GetAdapterDocsRequest) Descriptor() ([]byte, []int) {
	return file_adapter_proto_rawDescGZIP(), []int{2}
}

func (x *GetAdapterDocsRequest) GetProjectID() int64 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *GetAdapterDocsRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *GetAdapterDocsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetAdapterDocsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Adapters `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAdapterDocsResponse) Reset() {
	*x = GetAdapterDocsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdapterDocsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdapterDocsResponse) ProtoMessage() {}

func (x *GetAdapterDocsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adapter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdapterDocsResponse.ProtoReflect.Descriptor instead.
func (*GetAdapterDocsResponse) Descriptor() ([]byte, []int) {
	return file_adapter_proto_rawDescGZIP(), []int{3}
}

func (x *GetAdapterDocsResponse) GetData() []*Adapters {
	if x != nil {
		return x.Data
	}
	return nil
}

type AdapterStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string `protobuf:"bytes,1,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Strategy    string `protobuf:"bytes,2,opt,name=strategy,proto3" json:"strategy,omitempty"`
	Enable      bool   `protobuf:"varint,3,opt,name=enable,proto3" json:"enable,omitempty"`
}

func (x *AdapterStrategy) Reset() {
	*x = AdapterStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdapterStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdapterStrategy) ProtoMessage() {}

func (x *AdapterStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_adapter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdapterStrategy.ProtoReflect.Descriptor instead.
func (*AdapterStrategy) Descriptor() ([]byte, []int) {
	return file_adapter_proto_rawDescGZIP(), []int{4}
}

func (x *AdapterStrategy) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AdapterStrategy) GetStrategy() string {
	if x != nil {
		return x.Strategy
	}
	return ""
}

func (x *AdapterStrategy) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

type Adapters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language    string             `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	DisplayName string             `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Strategies  []*AdapterStrategy `protobuf:"bytes,3,rep,name=strategies,proto3" json:"strategies,omitempty"`
}

func (x *Adapters) Reset() {
	*x = Adapters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adapters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adapters) ProtoMessage() {}

func (x *Adapters) ProtoReflect() protoreflect.Message {
	mi := &file_adapter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adapters.ProtoReflect.Descriptor instead.
func (*Adapters) Descriptor() ([]byte, []int) {
	return file_adapter_proto_rawDescGZIP(), []int{5}
}

func (x *Adapters) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Adapters) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Adapters) GetStrategies() []*AdapterStrategy {
	if x != nil {
		return x.Strategies
	}
	return nil
}

type AdapterDocs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName     string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Type            string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Content         string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AdapterLanguage string `protobuf:"bytes,4,opt,name=adapterLanguage,proto3" json:"adapterLanguage,omitempty"`
	AdapterStrategy string `protobuf:"bytes,5,opt,name=adapterStrategy,proto3" json:"adapterStrategy,omitempty"`
}

func (x *AdapterDocs) Reset() {
	*x = AdapterDocs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdapterDocs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdapterDocs) ProtoMessage() {}

func (x *AdapterDocs) ProtoReflect() protoreflect.Message {
	mi := &file_adapter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdapterDocs.ProtoReflect.Descriptor instead.
func (*AdapterDocs) Descriptor() ([]byte, []int) {
	return file_adapter_proto_rawDescGZIP(), []int{6}
}

func (x *AdapterDocs) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AdapterDocs) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AdapterDocs) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AdapterDocs) GetAdapterLanguage() string {
	if x != nil {
		return x.AdapterLanguage
	}
	return ""
}

func (x *AdapterDocs) GetAdapterStrategy() string {
	if x != nil {
		return x.AdapterStrategy
	}
	return ""
}

var File_adapter_proto protoreflect.FileDescriptor

var file_adapter_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65,
	0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x61, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x73, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x44, 0x6f, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x67, 0x0a, 0x0f, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x8f,
	0x01, 0x0a, 0x08, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x22, 0xb1, 0x01, 0x0a, 0x0b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2a, 0x45, 0x0a, 0x0f, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x41, 0x56, 0x41, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x4f, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x48, 0x50,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x4f, 0x54, 0x5f, 0x4e, 0x45, 0x54, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x44, 0x45, 0x4a, 0x53, 0x10, 0x04, 0x2a, 0x6d, 0x0a, 0x12, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x4b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x0a, 0x4a, 0x41, 0x56, 0x41, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x50, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x53, 0x4b, 0x59, 0x57,
	0x41, 0x4c, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4a, 0x41, 0x45, 0x47,
	0x45, 0x52, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x54, 0x45, 0x4c,
	0x45, 0x4d, 0x45, 0x54, 0x52, 0x59, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x44, 0x45,
	0x4a, 0x53, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x32, 0xa5, 0x02, 0x0a, 0x0e, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x28, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x8e, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x44, 0x6f, 0x63, 0x73, 0x12, 0x2b, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x61, 0x70, 0x6d, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x70, 0x6d,
	0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73,
	0x70, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x64, 0x6f,
	0x63, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72,
	0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x73, 0x70, 0x2f,
	0x61, 0x70, 0x6d, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adapter_proto_rawDescOnce sync.Once
	file_adapter_proto_rawDescData = file_adapter_proto_rawDesc
)

func file_adapter_proto_rawDescGZIP() []byte {
	file_adapter_proto_rawDescOnce.Do(func() {
		file_adapter_proto_rawDescData = protoimpl.X.CompressGZIP(file_adapter_proto_rawDescData)
	})
	return file_adapter_proto_rawDescData
}

var file_adapter_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_adapter_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_adapter_proto_goTypes = []interface{}{
	(AdapterLanguage)(0),           // 0: erda.msp.apm.adapter.AdapterLanguage
	(AdapterStrategyKey)(0),        // 1: erda.msp.apm.adapter.AdapterStrategyKey
	(*GetAdaptersRequest)(nil),     // 2: erda.msp.apm.adapter.GetAdaptersRequest
	(*GetAdaptersResponse)(nil),    // 3: erda.msp.apm.adapter.GetAdaptersResponse
	(*GetAdapterDocsRequest)(nil),  // 4: erda.msp.apm.adapter.GetAdapterDocsRequest
	(*GetAdapterDocsResponse)(nil), // 5: erda.msp.apm.adapter.GetAdapterDocsResponse
	(*AdapterStrategy)(nil),        // 6: erda.msp.apm.adapter.AdapterStrategy
	(*Adapters)(nil),               // 7: erda.msp.apm.adapter.Adapters
	(*AdapterDocs)(nil),            // 8: erda.msp.apm.adapter.AdapterDocs
}
var file_adapter_proto_depIdxs = []int32{
	7, // 0: erda.msp.apm.adapter.GetAdaptersResponse.data:type_name -> erda.msp.apm.adapter.Adapters
	7, // 1: erda.msp.apm.adapter.GetAdapterDocsResponse.data:type_name -> erda.msp.apm.adapter.Adapters
	6, // 2: erda.msp.apm.adapter.Adapters.strategies:type_name -> erda.msp.apm.adapter.AdapterStrategy
	2, // 3: erda.msp.apm.adapter.AdapterService.GetAdapters:input_type -> erda.msp.apm.adapter.GetAdaptersRequest
	4, // 4: erda.msp.apm.adapter.AdapterService.GetAdapterDocs:input_type -> erda.msp.apm.adapter.GetAdapterDocsRequest
	3, // 5: erda.msp.apm.adapter.AdapterService.GetAdapters:output_type -> erda.msp.apm.adapter.GetAdaptersResponse
	5, // 6: erda.msp.apm.adapter.AdapterService.GetAdapterDocs:output_type -> erda.msp.apm.adapter.GetAdapterDocsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_adapter_proto_init() }
func file_adapter_proto_init() {
	if File_adapter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adapter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdaptersRequest); i {
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
		file_adapter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdaptersResponse); i {
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
		file_adapter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdapterDocsRequest); i {
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
		file_adapter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdapterDocsResponse); i {
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
		file_adapter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdapterStrategy); i {
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
		file_adapter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adapters); i {
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
		file_adapter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdapterDocs); i {
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
			RawDescriptor: file_adapter_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adapter_proto_goTypes,
		DependencyIndexes: file_adapter_proto_depIdxs,
		EnumInfos:         file_adapter_proto_enumTypes,
		MessageInfos:      file_adapter_proto_msgTypes,
	}.Build()
	File_adapter_proto = out.File
	file_adapter_proto_rawDesc = nil
	file_adapter_proto_goTypes = nil
	file_adapter_proto_depIdxs = nil
}