// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: device/proto.proto

package device

import (
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

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid      string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Mac       string `protobuf:"bytes,3,opt,name=mac,proto3" json:"mac,omitempty"`
	Firmware  string `protobuf:"bytes,4,opt,name=firmware,proto3" json:"firmware,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_device_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Device) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Device) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *Device) GetFirmware() string {
	if x != nil {
		return x.Firmware
	}
	return ""
}

func (x *Device) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Device) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac      string `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	Firmware string `protobuf:"bytes,2,opt,name=firmware,proto3" json:"firmware,omitempty"`
}

func (x *CreateDeviceRequest) Reset() {
	*x = CreateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceRequest) ProtoMessage() {}

func (x *CreateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceRequest.ProtoReflect.Descriptor instead.
func (*CreateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_device_proto_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeviceRequest) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *CreateDeviceRequest) GetFirmware() string {
	if x != nil {
		return x.Firmware
	}
	return ""
}

type DeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeviceRequest) Reset() {
	*x = DeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRequest) ProtoMessage() {}

func (x *DeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRequest.ProtoReflect.Descriptor instead.
func (*DeviceRequest) Descriptor() ([]byte, []int) {
	return file_device_proto_proto_rawDescGZIP(), []int{2}
}

type DeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *DeviceResponse) Reset() {
	*x = DeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceResponse) ProtoMessage() {}

func (x *DeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceResponse.ProtoReflect.Descriptor instead.
func (*DeviceResponse) Descriptor() ([]byte, []int) {
	return file_device_proto_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

var File_device_proto_proto protoreflect.FileDescriptor

var file_device_proto_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x98, 0x01, 0x0a,
	0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x61, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x43, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x22, 0x0f, 0x0a, 0x0d,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a,
	0x0e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0x85, 0x01, 0x0a, 0x05, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x15, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x70, 0x75, 0x74, 0x72, 0x61, 0x2f, 0x74, 0x75, 0x74, 0x6f, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x32, 0x34, 0x30,
	0x2f, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x61, 0x73, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_proto_proto_rawDescOnce sync.Once
	file_device_proto_proto_rawDescData = file_device_proto_proto_rawDesc
)

func file_device_proto_proto_rawDescGZIP() []byte {
	file_device_proto_proto_rawDescOnce.Do(func() {
		file_device_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_proto_proto_rawDescData)
	})
	return file_device_proto_proto_rawDescData
}

var file_device_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_device_proto_proto_goTypes = []interface{}{
	(*Device)(nil),              // 0: device.Device
	(*CreateDeviceRequest)(nil), // 1: device.CreateDeviceRequest
	(*DeviceRequest)(nil),       // 2: device.DeviceRequest
	(*DeviceResponse)(nil),      // 3: device.DeviceResponse
}
var file_device_proto_proto_depIdxs = []int32{
	0, // 0: device.DeviceResponse.devices:type_name -> device.Device
	1, // 1: device.Cloud.CreateDevice:input_type -> device.CreateDeviceRequest
	2, // 2: device.Cloud.GetDevices:input_type -> device.DeviceRequest
	0, // 3: device.Cloud.CreateDevice:output_type -> device.Device
	3, // 4: device.Cloud.GetDevices:output_type -> device.DeviceResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_device_proto_proto_init() }
func file_device_proto_proto_init() {
	if File_device_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_device_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceRequest); i {
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
		file_device_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRequest); i {
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
		file_device_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceResponse); i {
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
			RawDescriptor: file_device_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_device_proto_proto_goTypes,
		DependencyIndexes: file_device_proto_proto_depIdxs,
		MessageInfos:      file_device_proto_proto_msgTypes,
	}.Build()
	File_device_proto_proto = out.File
	file_device_proto_proto_rawDesc = nil
	file_device_proto_proto_goTypes = nil
	file_device_proto_proto_depIdxs = nil
}
