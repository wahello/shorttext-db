// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package task

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskPayload struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	//较大二级制流，单独封装传递给GRPC
	BigPayload           []byte   `protobuf:"bytes,2,opt,name=bigPayload,proto3" json:"bigPayload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskPayload) Reset()         { *m = TaskPayload{} }
func (m *TaskPayload) String() string { return proto.CompactTextString(m) }
func (*TaskPayload) ProtoMessage()    {}
func (*TaskPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *TaskPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskPayload.Unmarshal(m, b)
}
func (m *TaskPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskPayload.Marshal(b, m, deterministic)
}
func (m *TaskPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskPayload.Merge(m, src)
}
func (m *TaskPayload) XXX_Size() int {
	return xxx_messageInfo_TaskPayload.Size(m)
}
func (m *TaskPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskPayload.DiscardUnknown(m)
}

var xxx_messageInfo_TaskPayload proto.InternalMessageInfo

func (m *TaskPayload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TaskPayload) GetBigPayload() []byte {
	if m != nil {
		return m.BigPayload
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskPayload)(nil), "task.TaskPayload")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xdc, 0xb9, 0xb8, 0x43, 0x12,
	0x8b, 0xb3, 0x03, 0x12, 0x2b, 0x73, 0xf2, 0x13, 0x53, 0x84, 0x24, 0xb8, 0xd8, 0x0b, 0x20, 0x4c,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x18, 0x57, 0x48, 0x8e, 0x8b, 0x2b, 0x29, 0x33, 0x1d,
	0xaa, 0x4e, 0x82, 0x09, 0x2c, 0x89, 0x24, 0xe2, 0xa4, 0x1d, 0xa5, 0x99, 0x9c, 0x9f, 0xab, 0x97,
	0x97, 0x9a, 0x5a, 0xa0, 0x9f, 0x9e, 0x5f, 0x90, 0x93, 0x58, 0x92, 0x96, 0x5f, 0x94, 0xab, 0x9f,
	0x9a, 0x58, 0x5c, 0x99, 0x5b, 0xa4, 0x9f, 0x58, 0x54, 0x92, 0x99, 0x96, 0x98, 0x5c, 0x52, 0xac,
	0x0f, 0xb2, 0x35, 0x89, 0x0d, 0xec, 0x04, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0x17,
	0xee, 0x53, 0x90, 0x00, 0x00, 0x00,
}