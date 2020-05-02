// Wiregost - Golang Exploitation Framework
// Copyright © 2020 Para
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.8.0
// source: db/address.proto

package dbpb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Address - A network address
type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	HostID uint32 `protobuf:"varint,2,opt,name=HostID,proto3" json:"HostID,omitempty"`
	// @inject_tag: xml:"addr,attr"
	IP string `protobuf:"bytes,3,opt,name=IP,proto3" json:"IP,omitempty"`
	// @inject_tag: xml:"addrtype,attr"
	Type string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	// @inject_tag: xml:"vendor,attr"
	Vendor string `protobuf:"bytes,5,opt,name=Vendor,proto3" json:"Vendor,omitempty"`
	// We might have two subnets 192.168.1.1/24. How to know, when adding a host,
	// to which subnet it belongs ? We need to check a few things:
	// - Gateway for each address
	Gateway string `protobuf:"bytes,10,opt,name=Gateway,proto3" json:"Gateway,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_db_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_db_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Address) GetHostID() uint32 {
	if x != nil {
		return x.HostID
	}
	return 0
}

func (x *Address) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *Address) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Address) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *Address) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

var File_db_address_proto protoreflect.FileDescriptor

var file_db_address_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x62, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x64, 0x62, 0x22, 0x87, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x42, 0x06, 0x5a, 0x04, 0x64, 0x62, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_address_proto_rawDescOnce sync.Once
	file_db_address_proto_rawDescData = file_db_address_proto_rawDesc
)

func file_db_address_proto_rawDescGZIP() []byte {
	file_db_address_proto_rawDescOnce.Do(func() {
		file_db_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_address_proto_rawDescData)
	})
	return file_db_address_proto_rawDescData
}

var file_db_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_db_address_proto_goTypes = []interface{}{
	(*Address)(nil), // 0: db.Address
}
var file_db_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_db_address_proto_init() }
func file_db_address_proto_init() {
	if File_db_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_db_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_address_proto_goTypes,
		DependencyIndexes: file_db_address_proto_depIdxs,
		MessageInfos:      file_db_address_proto_msgTypes,
	}.Build()
	File_db_address_proto = out.File
	file_db_address_proto_rawDesc = nil
	file_db_address_proto_goTypes = nil
	file_db_address_proto_depIdxs = nil
}
