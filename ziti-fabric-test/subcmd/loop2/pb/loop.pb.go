/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: loop.proto

package loop2_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Test struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TxRequests           int32    `protobuf:"varint,2,opt,name=txRequests,proto3" json:"txRequests,omitempty"`
	TxPacing             int32    `protobuf:"varint,3,opt,name=txPacing,proto3" json:"txPacing,omitempty"`
	TxMaxJitter          int32    `protobuf:"varint,4,opt,name=txMaxJitter,proto3" json:"txMaxJitter,omitempty"`
	RxRequests           int32    `protobuf:"varint,5,opt,name=rxRequests,proto3" json:"rxRequests,omitempty"`
	RxTimeout            int32    `protobuf:"varint,6,opt,name=rxTimeout,proto3" json:"rxTimeout,omitempty"`
	PayloadMinBytes      int32    `protobuf:"varint,7,opt,name=payloadMinBytes,proto3" json:"payloadMinBytes,omitempty"`
	PayloadMaxBytes      int32    `protobuf:"varint,8,opt,name=payloadMaxBytes,proto3" json:"payloadMaxBytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_loop_1b5ea8ed7cfa23ba, []int{0}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (dst *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(dst, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Test) GetTxRequests() int32 {
	if m != nil {
		return m.TxRequests
	}
	return 0
}

func (m *Test) GetTxPacing() int32 {
	if m != nil {
		return m.TxPacing
	}
	return 0
}

func (m *Test) GetTxMaxJitter() int32 {
	if m != nil {
		return m.TxMaxJitter
	}
	return 0
}

func (m *Test) GetRxRequests() int32 {
	if m != nil {
		return m.RxRequests
	}
	return 0
}

func (m *Test) GetRxTimeout() int32 {
	if m != nil {
		return m.RxTimeout
	}
	return 0
}

func (m *Test) GetPayloadMinBytes() int32 {
	if m != nil {
		return m.PayloadMinBytes
	}
	return 0
}

func (m *Test) GetPayloadMaxBytes() int32 {
	if m != nil {
		return m.PayloadMaxBytes
	}
	return 0
}

type Block struct {
	Sequence             int32    `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_loop_1b5ea8ed7cfa23ba, []int{1}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Block) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Result struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_loop_1b5ea8ed7cfa23ba, []int{2}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "loop2.pb.Test")
	proto.RegisterType((*Block)(nil), "loop2.pb.Block")
	proto.RegisterType((*Result)(nil), "loop2.pb.Result")
}

func init() { proto.RegisterFile("loop.proto", fileDescriptor_loop_1b5ea8ed7cfa23ba) }

var fileDescriptor_loop_1b5ea8ed7cfa23ba = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x6b, 0x83, 0x40,
	0x10, 0x85, 0x31, 0x55, 0xa3, 0xd3, 0x40, 0x61, 0x4f, 0x4b, 0x29, 0x45, 0x3c, 0x79, 0xca, 0xa1,
	0xbd, 0xf6, 0x94, 0x63, 0x4b, 0xa0, 0x2c, 0xf9, 0x03, 0x1b, 0xb3, 0x24, 0x52, 0x75, 0xad, 0x33,
	0xc2, 0xe6, 0x2f, 0xf4, 0x57, 0x97, 0x1d, 0x9b, 0x68, 0x73, 0x9b, 0xf7, 0xbd, 0xc7, 0x43, 0xdf,
	0x02, 0xd4, 0xd6, 0x76, 0xeb, 0xae, 0xb7, 0x64, 0x45, 0xe2, 0xef, 0x97, 0x75, 0xb7, 0xcf, 0x7f,
	0x16, 0x10, 0xee, 0x0c, 0x92, 0x10, 0x10, 0xb6, 0xba, 0x31, 0x32, 0xc8, 0x82, 0x22, 0x55, 0x7c,
	0x8b, 0x67, 0x00, 0x72, 0xca, 0x7c, 0x0f, 0x06, 0x09, 0xe5, 0x22, 0x0b, 0x8a, 0x48, 0xcd, 0x88,
	0x78, 0x84, 0x84, 0xdc, 0xa7, 0x2e, 0xab, 0xf6, 0x28, 0xef, 0xd8, 0xbd, 0x6a, 0x91, 0xc1, 0x3d,
	0xb9, 0xad, 0x76, 0xef, 0x15, 0x91, 0xe9, 0x65, 0xc8, 0xf6, 0x1c, 0xf9, 0xf6, 0x7e, 0x6a, 0x8f,
	0xc6, 0xf6, 0x89, 0x88, 0x27, 0x48, 0x7b, 0xb7, 0xab, 0x1a, 0x63, 0x07, 0x92, 0x31, 0xdb, 0x13,
	0x10, 0x05, 0x3c, 0x74, 0xfa, 0x5c, 0x5b, 0x7d, 0xd8, 0x56, 0xed, 0xe6, 0x4c, 0x06, 0xe5, 0x92,
	0x33, 0xb7, 0x78, 0x9e, 0xd4, 0x6e, 0x4c, 0x26, 0xff, 0x93, 0x7f, 0x38, 0xff, 0x80, 0x68, 0x53,
	0xdb, 0xf2, 0xcb, 0xff, 0x18, 0xfa, 0xcf, 0x68, 0xcb, 0x71, 0x90, 0x48, 0x5d, 0xb5, 0x1f, 0xea,
	0xa0, 0x49, 0xf3, 0x1c, 0x2b, 0xc5, 0xb7, 0x67, 0x27, 0x8d, 0x27, 0x1e, 0x61, 0xa5, 0xf8, 0xce,
	0xdf, 0x20, 0x56, 0x06, 0x87, 0x9a, 0x84, 0x84, 0x25, 0x0e, 0x65, 0x69, 0x10, 0xb9, 0x2c, 0x51,
	0x17, 0xe9, 0x9d, 0xc6, 0x20, 0xea, 0xa3, 0xe1, 0xba, 0x54, 0x5d, 0xe4, 0x3e, 0xe6, 0x87, 0x7a,
	0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x55, 0xc7, 0x03, 0x63, 0xb6, 0x01, 0x00, 0x00,
}
