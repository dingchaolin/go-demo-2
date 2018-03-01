// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

/*
Package myproto is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	PhoneNumber
	Person
	AddressBook
	AddPersonRequest
	AddPersonResponse
*/
package myproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type PhoneType int32

const (
	PhoneType_MOBILE PhoneType = 0
	PhoneType_HOME   PhoneType = 1
	PhoneType_WORK   PhoneType = 2
)

var PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}
func (PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PhoneNumber struct {
	Number string    `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   PhoneType `protobuf:"varint,2,opt,name=type,enum=myproto.PhoneType" json:"type,omitempty"`
}

func (m *PhoneNumber) Reset()                    { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()               {}
func (*PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *PhoneNumber) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_MOBILE
}

type Person struct {
	Id     int32          `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name   string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email  string         `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

type AddPersonRequest struct {
	Person *Person `protobuf:"bytes,1,opt,name=person" json:"person,omitempty"`
}

func (m *AddPersonRequest) Reset()                    { *m = AddPersonRequest{} }
func (m *AddPersonRequest) String() string            { return proto.CompactTextString(m) }
func (*AddPersonRequest) ProtoMessage()               {}
func (*AddPersonRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddPersonRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type AddPersonResponse struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *AddPersonResponse) Reset()                    { *m = AddPersonResponse{} }
func (m *AddPersonResponse) String() string            { return proto.CompactTextString(m) }
func (*AddPersonResponse) ProtoMessage()               {}
func (*AddPersonResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddPersonResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*PhoneNumber)(nil), "myproto.PhoneNumber")
	proto.RegisterType((*Person)(nil), "myproto.Person")
	proto.RegisterType((*AddressBook)(nil), "myproto.AddressBook")
	proto.RegisterType((*AddPersonRequest)(nil), "myproto.AddPersonRequest")
	proto.RegisterType((*AddPersonResponse)(nil), "myproto.AddPersonResponse")
	proto.RegisterEnum("myproto.PhoneType", PhoneType_name, PhoneType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AddressBookStore service

type AddressBookStoreClient interface {
	AddPerson(ctx context.Context, in *AddPersonRequest, opts ...grpc.CallOption) (*AddPersonResponse, error)
}

type addressBookStoreClient struct {
	cc *grpc.ClientConn
}

func NewAddressBookStoreClient(cc *grpc.ClientConn) AddressBookStoreClient {
	return &addressBookStoreClient{cc}
}

func (c *addressBookStoreClient) AddPerson(ctx context.Context, in *AddPersonRequest, opts ...grpc.CallOption) (*AddPersonResponse, error) {
	out := new(AddPersonResponse)
	err := grpc.Invoke(ctx, "/myproto.AddressBookStore/AddPerson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AddressBookStore service

type AddressBookStoreServer interface {
	AddPerson(context.Context, *AddPersonRequest) (*AddPersonResponse, error)
}

func RegisterAddressBookStoreServer(s *grpc.Server, srv AddressBookStoreServer) {
	s.RegisterService(&_AddressBookStore_serviceDesc, srv)
}

func _AddressBookStore_AddPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookStoreServer).AddPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproto.AddressBookStore/AddPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookStoreServer).AddPerson(ctx, req.(*AddPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddressBookStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myproto.AddressBookStore",
	HandlerType: (*AddressBookStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPerson",
			Handler:    _AddressBookStore_AddPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addressbook.proto",
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4f, 0xf2, 0x40,
	0x10, 0xc5, 0xbf, 0x96, 0xb2, 0x9f, 0x1d, 0x12, 0x2c, 0x13, 0x62, 0x2a, 0x27, 0x52, 0x13, 0x25,
	0x6a, 0x38, 0x60, 0xe2, 0xc5, 0x13, 0x24, 0x24, 0x1a, 0x45, 0xc8, 0x6a, 0xf4, 0x0c, 0xd9, 0x49,
	0x24, 0xd0, 0xee, 0xda, 0x2d, 0x07, 0xfe, 0x7b, 0xd3, 0xd9, 0x8a, 0x4a, 0x38, 0xf5, 0xcd, 0xbc,
	0xbe, 0x99, 0xdf, 0x2c, 0xb4, 0xe6, 0x4a, 0xe5, 0x64, 0xed, 0x42, 0xeb, 0x55, 0xdf, 0xe4, 0xba,
	0xd0, 0xf8, 0x3f, 0xdd, 0xb2, 0x48, 0x26, 0xd0, 0x98, 0x7d, 0xe8, 0x8c, 0x9e, 0x37, 0xe9, 0x82,
	0x72, 0x3c, 0x01, 0x91, 0xb1, 0x8a, 0xbd, 0xae, 0xd7, 0x0b, 0x65, 0x55, 0xe1, 0x39, 0x04, 0xc5,
	0xd6, 0x50, 0xec, 0x77, 0xbd, 0x5e, 0x73, 0x80, 0xfd, 0x2a, 0xde, 0xe7, 0xec, 0xeb, 0xd6, 0x90,
	0x64, 0x3f, 0x31, 0x20, 0x66, 0x94, 0x5b, 0x9d, 0x61, 0x13, 0xfc, 0xa5, 0xe2, 0x29, 0x75, 0xe9,
	0x2f, 0x15, 0x22, 0x04, 0xd9, 0x3c, 0x75, 0x13, 0x42, 0xc9, 0x1a, 0xdb, 0x50, 0xa7, 0x74, 0xbe,
	0x5c, 0xc7, 0x35, 0x6e, 0xba, 0x02, 0xaf, 0x41, 0x98, 0x72, 0xac, 0x8d, 0x83, 0x6e, 0xad, 0xd7,
	0x18, 0xb4, 0xff, 0x6e, 0x73, 0xa4, 0xb2, 0xfa, 0x27, 0xb9, 0x85, 0xc6, 0xd0, 0x9d, 0x37, 0xd2,
	0x7a, 0x85, 0x17, 0x20, 0x0c, 0x69, 0xb3, 0xa6, 0xd8, 0xe3, 0xf0, 0xf1, 0x4f, 0x98, 0xb9, 0x64,
	0x65, 0x27, 0x77, 0x10, 0x0d, 0x95, 0xaa, 0x9a, 0xf4, 0xb9, 0x21, 0x5b, 0xb8, 0x70, 0xd9, 0x60,
	0xee, 0xc3, 0xe1, 0xf2, 0x9b, 0x9c, 0x41, 0xeb, 0x57, 0xd8, 0x1a, 0x9d, 0x59, 0xda, 0xbf, 0xf8,
	0xf2, 0x0a, 0xc2, 0xdd, 0xf3, 0x20, 0x80, 0x98, 0x4c, 0x47, 0x0f, 0x4f, 0xe3, 0xe8, 0x1f, 0x1e,
	0x41, 0x70, 0x3f, 0x9d, 0x8c, 0x23, 0xaf, 0x54, 0xef, 0x53, 0xf9, 0x18, 0xf9, 0x83, 0x37, 0xc6,
	0xf9, 0x3e, 0xe3, 0xa5, 0xd0, 0x39, 0xe1, 0x08, 0xc2, 0xdd, 0x16, 0x3c, 0xdd, 0xb1, 0xec, 0x63,
	0x77, 0x3a, 0x87, 0x2c, 0x07, 0xb5, 0x10, 0x6c, 0xdc, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x87, 0xce, 0xe5, 0x04, 0x02, 0x00, 0x00,
}
