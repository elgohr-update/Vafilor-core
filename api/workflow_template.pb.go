// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workflow_template.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateWorkflowTemplateRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WorkflowTemplate     *WorkflowTemplate `protobuf:"bytes,2,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateWorkflowTemplateRequest) Reset()         { *m = CreateWorkflowTemplateRequest{} }
func (m *CreateWorkflowTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWorkflowTemplateRequest) ProtoMessage()    {}
func (*CreateWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{0}
}

func (m *CreateWorkflowTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkflowTemplateRequest.Unmarshal(m, b)
}
func (m *CreateWorkflowTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkflowTemplateRequest.Marshal(b, m, deterministic)
}
func (m *CreateWorkflowTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkflowTemplateRequest.Merge(m, src)
}
func (m *CreateWorkflowTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWorkflowTemplateRequest.Size(m)
}
func (m *CreateWorkflowTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkflowTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkflowTemplateRequest proto.InternalMessageInfo

func (m *CreateWorkflowTemplateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateWorkflowTemplateRequest) GetWorkflowTemplate() *WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplate
	}
	return nil
}

type GetWorkflowTemplateRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Version              int32    `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWorkflowTemplateRequest) Reset()         { *m = GetWorkflowTemplateRequest{} }
func (m *GetWorkflowTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*GetWorkflowTemplateRequest) ProtoMessage()    {}
func (*GetWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{1}
}

func (m *GetWorkflowTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkflowTemplateRequest.Unmarshal(m, b)
}
func (m *GetWorkflowTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkflowTemplateRequest.Marshal(b, m, deterministic)
}
func (m *GetWorkflowTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkflowTemplateRequest.Merge(m, src)
}
func (m *GetWorkflowTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_GetWorkflowTemplateRequest.Size(m)
}
func (m *GetWorkflowTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkflowTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkflowTemplateRequest proto.InternalMessageInfo

func (m *GetWorkflowTemplateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetWorkflowTemplateRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GetWorkflowTemplateRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type ListWorkflowTemplateVersionsRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWorkflowTemplateVersionsRequest) Reset()         { *m = ListWorkflowTemplateVersionsRequest{} }
func (m *ListWorkflowTemplateVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowTemplateVersionsRequest) ProtoMessage()    {}
func (*ListWorkflowTemplateVersionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{2}
}

func (m *ListWorkflowTemplateVersionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowTemplateVersionsRequest.Unmarshal(m, b)
}
func (m *ListWorkflowTemplateVersionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowTemplateVersionsRequest.Marshal(b, m, deterministic)
}
func (m *ListWorkflowTemplateVersionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowTemplateVersionsRequest.Merge(m, src)
}
func (m *ListWorkflowTemplateVersionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowTemplateVersionsRequest.Size(m)
}
func (m *ListWorkflowTemplateVersionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowTemplateVersionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowTemplateVersionsRequest proto.InternalMessageInfo

func (m *ListWorkflowTemplateVersionsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListWorkflowTemplateVersionsRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type ListWorkflowTemplateVersionsResponse struct {
	Count                int32               `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	WorkflowTemplates    []*WorkflowTemplate `protobuf:"bytes,2,rep,name=workflowTemplates,proto3" json:"workflowTemplates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListWorkflowTemplateVersionsResponse) Reset()         { *m = ListWorkflowTemplateVersionsResponse{} }
func (m *ListWorkflowTemplateVersionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowTemplateVersionsResponse) ProtoMessage()    {}
func (*ListWorkflowTemplateVersionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{3}
}

func (m *ListWorkflowTemplateVersionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowTemplateVersionsResponse.Unmarshal(m, b)
}
func (m *ListWorkflowTemplateVersionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowTemplateVersionsResponse.Marshal(b, m, deterministic)
}
func (m *ListWorkflowTemplateVersionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowTemplateVersionsResponse.Merge(m, src)
}
func (m *ListWorkflowTemplateVersionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowTemplateVersionsResponse.Size(m)
}
func (m *ListWorkflowTemplateVersionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowTemplateVersionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowTemplateVersionsResponse proto.InternalMessageInfo

func (m *ListWorkflowTemplateVersionsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListWorkflowTemplateVersionsResponse) GetWorkflowTemplates() []*WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplates
	}
	return nil
}

type ListWorkflowTemplatesRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWorkflowTemplatesRequest) Reset()         { *m = ListWorkflowTemplatesRequest{} }
func (m *ListWorkflowTemplatesRequest) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowTemplatesRequest) ProtoMessage()    {}
func (*ListWorkflowTemplatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{4}
}

func (m *ListWorkflowTemplatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowTemplatesRequest.Unmarshal(m, b)
}
func (m *ListWorkflowTemplatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowTemplatesRequest.Marshal(b, m, deterministic)
}
func (m *ListWorkflowTemplatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowTemplatesRequest.Merge(m, src)
}
func (m *ListWorkflowTemplatesRequest) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowTemplatesRequest.Size(m)
}
func (m *ListWorkflowTemplatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowTemplatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowTemplatesRequest proto.InternalMessageInfo

func (m *ListWorkflowTemplatesRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ListWorkflowTemplatesResponse struct {
	Count                int32               `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	WorkflowTemplates    []*WorkflowTemplate `protobuf:"bytes,2,rep,name=workflowTemplates,proto3" json:"workflowTemplates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListWorkflowTemplatesResponse) Reset()         { *m = ListWorkflowTemplatesResponse{} }
func (m *ListWorkflowTemplatesResponse) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowTemplatesResponse) ProtoMessage()    {}
func (*ListWorkflowTemplatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{5}
}

func (m *ListWorkflowTemplatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowTemplatesResponse.Unmarshal(m, b)
}
func (m *ListWorkflowTemplatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowTemplatesResponse.Marshal(b, m, deterministic)
}
func (m *ListWorkflowTemplatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowTemplatesResponse.Merge(m, src)
}
func (m *ListWorkflowTemplatesResponse) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowTemplatesResponse.Size(m)
}
func (m *ListWorkflowTemplatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowTemplatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowTemplatesResponse proto.InternalMessageInfo

func (m *ListWorkflowTemplatesResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListWorkflowTemplatesResponse) GetWorkflowTemplates() []*WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplates
	}
	return nil
}

type WorkflowTemplate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              int32    `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Manifest             string   `protobuf:"bytes,4,opt,name=manifest,proto3" json:"manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowTemplate) Reset()         { *m = WorkflowTemplate{} }
func (m *WorkflowTemplate) String() string { return proto.CompactTextString(m) }
func (*WorkflowTemplate) ProtoMessage()    {}
func (*WorkflowTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{6}
}

func (m *WorkflowTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTemplate.Unmarshal(m, b)
}
func (m *WorkflowTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTemplate.Marshal(b, m, deterministic)
}
func (m *WorkflowTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTemplate.Merge(m, src)
}
func (m *WorkflowTemplate) XXX_Size() int {
	return xxx_messageInfo_WorkflowTemplate.Size(m)
}
func (m *WorkflowTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTemplate proto.InternalMessageInfo

func (m *WorkflowTemplate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *WorkflowTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowTemplate) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *WorkflowTemplate) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateWorkflowTemplateRequest)(nil), "api.CreateWorkflowTemplateRequest")
	proto.RegisterType((*GetWorkflowTemplateRequest)(nil), "api.GetWorkflowTemplateRequest")
	proto.RegisterType((*ListWorkflowTemplateVersionsRequest)(nil), "api.ListWorkflowTemplateVersionsRequest")
	proto.RegisterType((*ListWorkflowTemplateVersionsResponse)(nil), "api.ListWorkflowTemplateVersionsResponse")
	proto.RegisterType((*ListWorkflowTemplatesRequest)(nil), "api.ListWorkflowTemplatesRequest")
	proto.RegisterType((*ListWorkflowTemplatesResponse)(nil), "api.ListWorkflowTemplatesResponse")
	proto.RegisterType((*WorkflowTemplate)(nil), "api.WorkflowTemplate")
}

func init() { proto.RegisterFile("workflow_template.proto", fileDescriptor_b9a07547748a96e8) }

var fileDescriptor_b9a07547748a96e8 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xbf, 0x4f, 0xf3, 0x30,
	0x10, 0x95, 0x9b, 0xe6, 0xfb, 0xc8, 0xb1, 0x04, 0x0b, 0x84, 0x55, 0xb5, 0x52, 0x64, 0x18, 0x32,
	0x65, 0x80, 0x95, 0x05, 0x75, 0x60, 0x61, 0x8a, 0xf8, 0x31, 0x22, 0x53, 0x2e, 0x92, 0x45, 0x63,
	0x9b, 0xd8, 0x21, 0x12, 0x13, 0xfc, 0xe7, 0xa8, 0x6e, 0x4a, 0xa5, 0x84, 0x20, 0x54, 0x89, 0xed,
	0xce, 0x7e, 0xf7, 0xde, 0xf3, 0x3b, 0xc3, 0x71, 0xa3, 0xab, 0xe7, 0x62, 0xa9, 0x9b, 0x07, 0x87,
	0xa5, 0x59, 0x0a, 0x87, 0x99, 0xa9, 0xb4, 0xd3, 0x34, 0x10, 0x46, 0xf2, 0x77, 0x02, 0xb3, 0x79,
	0x85, 0xc2, 0xe1, 0x7d, 0x0b, 0xbb, 0x69, 0x51, 0x39, 0xbe, 0xd4, 0x68, 0x1d, 0x9d, 0x42, 0xa4,
	0x44, 0x89, 0xd6, 0x88, 0x05, 0x32, 0x92, 0x90, 0x34, 0xca, 0xb7, 0x07, 0xf4, 0x12, 0xe2, 0xa6,
	0x33, 0xc8, 0x46, 0x09, 0x49, 0xf7, 0xcf, 0x8e, 0x32, 0x61, 0x64, 0xd6, 0x63, 0xed, 0xc1, 0x79,
	0x01, 0x93, 0x2b, 0x74, 0xbb, 0xc9, 0xc7, 0x10, 0xd4, 0xf2, 0xc9, 0x2b, 0x46, 0xf9, 0xaa, 0xa4,
	0x0c, 0xfe, 0xbf, 0x62, 0x65, 0xa5, 0x56, 0x2c, 0x48, 0x48, 0x1a, 0xe6, 0x9b, 0x96, 0xdf, 0xc2,
	0xc9, 0xb5, 0xb4, 0x3d, 0xa1, 0xbb, 0xf5, 0xb5, 0xdd, 0x51, 0x90, 0x7f, 0x10, 0x38, 0xfd, 0x99,
	0xd7, 0x1a, 0xad, 0x2c, 0xd2, 0x43, 0x08, 0x17, 0xba, 0x56, 0xce, 0x93, 0x86, 0xf9, 0xba, 0xa1,
	0x73, 0x38, 0xe8, 0x26, 0x62, 0xd9, 0x28, 0x09, 0x86, 0x13, 0xec, 0xe3, 0xf9, 0x05, 0x4c, 0xbf,
	0xb3, 0xf0, 0xbb, 0x37, 0xf1, 0x37, 0x98, 0x0d, 0x4c, 0xff, 0xbd, 0x73, 0x05, 0x71, 0x17, 0xb6,
	0xc9, 0x98, 0x6c, 0x97, 0x4a, 0x61, 0xbc, 0xb2, 0xdb, 0xc6, 0xee, 0xeb, 0xe1, 0x45, 0xd3, 0x09,
	0xec, 0x95, 0x42, 0xc9, 0x02, 0xad, 0x63, 0x63, 0x3f, 0xf1, 0xd5, 0x3f, 0xfe, 0xf3, 0x7f, 0xff,
	0xfc, 0x33, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x8f, 0xf0, 0x3d, 0x16, 0x03, 0x00, 0x00,
}
