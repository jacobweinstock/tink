//
// A template is a set of jobs and actions that can be performed to some
// hardware.
// It uses YAML to describe order and relation of those actions. Each action is
// a Docker container.
// When a template is stored in Tinkerbell it can be used to trigger a workflow
// targeting a specific hardwarwe.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: protos/template/template.proto

package template

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Empty represents an empty response
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_template_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_template_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_template_template_proto_rawDescGZIP(), []int{0}
}

// WorkflowTemplate describes the template itself.
type WorkflowTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The template identifier
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the template. You can see it as a friendly way to remember a
	// template
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// When a template got created
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The last time a template was modified
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// When a template got deleted. This is the value used to identify a deleted
	// template as well. If empty the template can be used to generate workflows.
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	// The content of the template in its YAML representation
	Data string `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WorkflowTemplate) Reset() {
	*x = WorkflowTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_template_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowTemplate) ProtoMessage() {}

func (x *WorkflowTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_template_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowTemplate.ProtoReflect.Descriptor instead.
func (*WorkflowTemplate) Descriptor() ([]byte, []int) {
	return file_protos_template_template_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowTemplate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkflowTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkflowTemplate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *WorkflowTemplate) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *WorkflowTemplate) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *WorkflowTemplate) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// CreateResponse returns the ID of the created template
type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_template_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_template_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_protos_template_template_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetRequest enables you to filter templates by name or id.
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to GetBy:
	//
	//	*GetRequest_Id
	//	*GetRequest_Name
	GetBy isGetRequest_GetBy `protobuf_oneof:"get_by"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_template_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_template_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_protos_template_template_proto_rawDescGZIP(), []int{3}
}

func (m *GetRequest) GetGetBy() isGetRequest_GetBy {
	if m != nil {
		return m.GetBy
	}
	return nil
}

func (x *GetRequest) GetId() string {
	if x, ok := x.GetGetBy().(*GetRequest_Id); ok {
		return x.Id
	}
	return ""
}

func (x *GetRequest) GetName() string {
	if x, ok := x.GetGetBy().(*GetRequest_Name); ok {
		return x.Name
	}
	return ""
}

type isGetRequest_GetBy interface {
	isGetRequest_GetBy()
}

type GetRequest_Id struct {
	// Filter by the template id
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type GetRequest_Name struct {
	// Filter by the template name
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*GetRequest_Id) isGetRequest_GetBy() {}

func (*GetRequest_Name) isGetRequest_GetBy() {}

// The request to use during a List template request
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FilterBy:
	//
	//	*ListRequest_Name
	FilterBy isListRequest_FilterBy `protobuf_oneof:"filter_by"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_template_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_template_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_protos_template_template_proto_rawDescGZIP(), []int{4}
}

func (m *ListRequest) GetFilterBy() isListRequest_FilterBy {
	if m != nil {
		return m.FilterBy
	}
	return nil
}

func (x *ListRequest) GetName() string {
	if x, ok := x.GetFilterBy().(*ListRequest_Name); ok {
		return x.Name
	}
	return ""
}

type isListRequest_FilterBy interface {
	isListRequest_FilterBy()
}

type ListRequest_Name struct {
	// Filter by the template name
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

func (*ListRequest_Name) isListRequest_FilterBy() {}

var File_protos_template_template_proto protoreflect.FileDescriptor

var file_protos_template_template_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x81, 0x02, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x08, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x22, 0x30, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x32, 0x9a, 0x06, 0x0a,
	0x0f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xa4, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x3c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x1a, 0x3a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x9f, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e,
	0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x97, 0x01, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72,
	0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x9f, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x37, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x30, 0x01, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x31, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65,
	0x6c, 0x6c, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_template_template_proto_rawDescOnce sync.Once
	file_protos_template_template_proto_rawDescData = file_protos_template_template_proto_rawDesc
)

func file_protos_template_template_proto_rawDescGZIP() []byte {
	file_protos_template_template_proto_rawDescOnce.Do(func() {
		file_protos_template_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_template_template_proto_rawDescData)
	})
	return file_protos_template_template_proto_rawDescData
}

var (
	file_protos_template_template_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
	file_protos_template_template_proto_goTypes  = []interface{}{
		(*Empty)(nil),                 // 0: github.com.tinkerbell.tink.protos.template.Empty
		(*WorkflowTemplate)(nil),      // 1: github.com.tinkerbell.tink.protos.template.WorkflowTemplate
		(*CreateResponse)(nil),        // 2: github.com.tinkerbell.tink.protos.template.CreateResponse
		(*GetRequest)(nil),            // 3: github.com.tinkerbell.tink.protos.template.GetRequest
		(*ListRequest)(nil),           // 4: github.com.tinkerbell.tink.protos.template.ListRequest
		(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	}
)

var file_protos_template_template_proto_depIdxs = []int32{
	5, // 0: github.com.tinkerbell.tink.protos.template.WorkflowTemplate.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: github.com.tinkerbell.tink.protos.template.WorkflowTemplate.updated_at:type_name -> google.protobuf.Timestamp
	5, // 2: github.com.tinkerbell.tink.protos.template.WorkflowTemplate.deleted_at:type_name -> google.protobuf.Timestamp
	1, // 3: github.com.tinkerbell.tink.protos.template.TemplateService.CreateTemplate:input_type -> github.com.tinkerbell.tink.protos.template.WorkflowTemplate
	3, // 4: github.com.tinkerbell.tink.protos.template.TemplateService.GetTemplate:input_type -> github.com.tinkerbell.tink.protos.template.GetRequest
	3, // 5: github.com.tinkerbell.tink.protos.template.TemplateService.DeleteTemplate:input_type -> github.com.tinkerbell.tink.protos.template.GetRequest
	4, // 6: github.com.tinkerbell.tink.protos.template.TemplateService.ListTemplates:input_type -> github.com.tinkerbell.tink.protos.template.ListRequest
	1, // 7: github.com.tinkerbell.tink.protos.template.TemplateService.UpdateTemplate:input_type -> github.com.tinkerbell.tink.protos.template.WorkflowTemplate
	2, // 8: github.com.tinkerbell.tink.protos.template.TemplateService.CreateTemplate:output_type -> github.com.tinkerbell.tink.protos.template.CreateResponse
	1, // 9: github.com.tinkerbell.tink.protos.template.TemplateService.GetTemplate:output_type -> github.com.tinkerbell.tink.protos.template.WorkflowTemplate
	0, // 10: github.com.tinkerbell.tink.protos.template.TemplateService.DeleteTemplate:output_type -> github.com.tinkerbell.tink.protos.template.Empty
	1, // 11: github.com.tinkerbell.tink.protos.template.TemplateService.ListTemplates:output_type -> github.com.tinkerbell.tink.protos.template.WorkflowTemplate
	0, // 12: github.com.tinkerbell.tink.protos.template.TemplateService.UpdateTemplate:output_type -> github.com.tinkerbell.tink.protos.template.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_template_template_proto_init() }
func file_protos_template_template_proto_init() {
	if File_protos_template_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_template_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protos_template_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowTemplate); i {
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
		file_protos_template_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_protos_template_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_protos_template_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
	file_protos_template_template_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*GetRequest_Id)(nil),
		(*GetRequest_Name)(nil),
	}
	file_protos_template_template_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ListRequest_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_template_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_template_template_proto_goTypes,
		DependencyIndexes: file_protos_template_template_proto_depIdxs,
		MessageInfos:      file_protos_template_template_proto_msgTypes,
	}.Build()
	File_protos_template_template_proto = out.File
	file_protos_template_template_proto_rawDesc = nil
	file_protos_template_template_proto_goTypes = nil
	file_protos_template_template_proto_depIdxs = nil
}
