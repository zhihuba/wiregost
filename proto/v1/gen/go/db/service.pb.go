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
// source: db/service.proto

package dbpb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Service - A service running behind a port
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// General
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// @inject_tag: gorm:"not null"
	PortID uint32 `protobuf:"varint,2,opt,name=PortID,proto3" json:"PortID,omitempty" gorm:"not null"`
	// @inject_tag: xml:"proto,attr"
	Protocol string `protobuf:"bytes,3,opt,name=Protocol,proto3" json:"Protocol,omitempty" xml:"proto,attr"`
	// @inject_tag: xml:"name,attr"
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty" xml:"name,attr"`
	// @inject_tag: xml:"extrainfo,attr"
	ExtraInfo string `protobuf:"bytes,5,opt,name=ExtraInfo,proto3" json:"ExtraInfo,omitempty" xml:"extrainfo,attr"`
	// Nmap
	// @inject_tag: xml:"devicetype,attr"
	DeviceType string `protobuf:"bytes,10,opt,name=DeviceType,proto3" json:"DeviceType,omitempty" xml:"devicetype,attr"`
	// @inject_tag: xml:"hostname,attr"
	Hostname string `protobuf:"bytes,12,opt,name=Hostname,proto3" json:"Hostname,omitempty" xml:"hostname,attr"`
	// @inject_tag: xml:"method,attr"
	Method string `protobuf:"bytes,14,opt,name=Method,proto3" json:"Method,omitempty" xml:"method,attr"`
	// @inject_tag: xml:"ostype,attr"
	OSType string `protobuf:"bytes,15,opt,name=OSType,proto3" json:"OSType,omitempty" xml:"ostype,attr"`
	// @inject_tag: xml:"product,attr"
	Product string `protobuf:"bytes,16,opt,name=Product,proto3" json:"Product,omitempty" xml:"product,attr"`
	// @inject_tag: xml:"rpcnum,attr"
	RPCNum string `protobuf:"bytes,17,opt,name=RPCNum,proto3" json:"RPCNum,omitempty" xml:"rpcnum,attr"`
	// @inject_tag: xml:"servicefp,attr"
	ServiceFP string `protobuf:"bytes,18,opt,name=ServiceFP,proto3" json:"ServiceFP,omitempty" xml:"servicefp,attr"`
	// @inject_tag: xml:"tunnel,attr"
	Tunnel string `protobuf:"bytes,19,opt,name=Tunnel,proto3" json:"Tunnel,omitempty" xml:"tunnel,attr"`
	// @inject_tag: xml:"lowver,attr"
	LowVersion string `protobuf:"bytes,13,opt,name=LowVersion,proto3" json:"LowVersion,omitempty" xml:"lowver,attr"`
	// @inject_tag: xml:"highver,attr"
	HighVersion string `protobuf:"bytes,11,opt,name=HighVersion,proto3" json:"HighVersion,omitempty" xml:"highver,attr"`
	// @inject_tag: xml:"version,attr"
	Version string `protobuf:"bytes,20,opt,name=Version,proto3" json:"Version,omitempty" xml:"version,attr"`
	// @inject_tag: xml:"conf,attr"
	Configuration int32 `protobuf:"varint,21,opt,name=Configuration,proto3" json:"Configuration,omitempty" xml:"conf,attr"`
	// @inject_tag: xml:"cpe"
	CPEs []string `protobuf:"bytes,6,rep,name=CPEs,proto3" json:"CPEs,omitempty" xml:"cpe"` // "Common Platform Enumeration" is standardized way to name software applications, OSs and Hardware platforms
	// Timestamp
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,38,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,39,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_db_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_db_service_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Service) GetPortID() uint32 {
	if x != nil {
		return x.PortID
	}
	return 0
}

func (x *Service) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetExtraInfo() string {
	if x != nil {
		return x.ExtraInfo
	}
	return ""
}

func (x *Service) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *Service) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Service) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Service) GetOSType() string {
	if x != nil {
		return x.OSType
	}
	return ""
}

func (x *Service) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *Service) GetRPCNum() string {
	if x != nil {
		return x.RPCNum
	}
	return ""
}

func (x *Service) GetServiceFP() string {
	if x != nil {
		return x.ServiceFP
	}
	return ""
}

func (x *Service) GetTunnel() string {
	if x != nil {
		return x.Tunnel
	}
	return ""
}

func (x *Service) GetLowVersion() string {
	if x != nil {
		return x.LowVersion
	}
	return ""
}

func (x *Service) GetHighVersion() string {
	if x != nil {
		return x.HighVersion
	}
	return ""
}

func (x *Service) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Service) GetConfiguration() int32 {
	if x != nil {
		return x.Configuration
	}
	return 0
}

func (x *Service) GetCPEs() []string {
	if x != nil {
		return x.CPEs
	}
	return nil
}

func (x *Service) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Service) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_db_service_proto protoreflect.FileDescriptor

var file_db_service_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x64, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x04, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x53, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f,
	0x53, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x50, 0x43, 0x4e, 0x75, 0x6d, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x52, 0x50, 0x43, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x46, 0x50, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x46, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x4c, 0x6f, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x4c, 0x6f, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x48, 0x69, 0x67, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x48, 0x69, 0x67, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x50, 0x45, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x50, 0x45, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x26, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x27, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x64, 0x62, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_service_proto_rawDescOnce sync.Once
	file_db_service_proto_rawDescData = file_db_service_proto_rawDesc
)

func file_db_service_proto_rawDescGZIP() []byte {
	file_db_service_proto_rawDescOnce.Do(func() {
		file_db_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_service_proto_rawDescData)
	})
	return file_db_service_proto_rawDescData
}

var file_db_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_db_service_proto_goTypes = []interface{}{
	(*Service)(nil),             // 0: db.Service
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_db_service_proto_depIdxs = []int32{
	1, // 0: db.Service.CreatedAt:type_name -> google.protobuf.Timestamp
	1, // 1: db.Service.UpdatedAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_db_service_proto_init() }
func file_db_service_proto_init() {
	if File_db_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
			RawDescriptor: file_db_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_service_proto_goTypes,
		DependencyIndexes: file_db_service_proto_depIdxs,
		MessageInfos:      file_db_service_proto_msgTypes,
	}.Build()
	File_db_service_proto = out.File
	file_db_service_proto_rawDesc = nil
	file_db_service_proto_goTypes = nil
	file_db_service_proto_depIdxs = nil
}