// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v5.26.1
// source: rpc_get_customer_by_barcode.proto

package pb

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

type GetCustomerByBarcodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barcode string `protobuf:"bytes,1,opt,name=barcode,proto3" json:"barcode,omitempty"`
}

func (x *GetCustomerByBarcodeRequest) Reset() {
	*x = GetCustomerByBarcodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_customer_by_barcode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByBarcodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByBarcodeRequest) ProtoMessage() {}

func (x *GetCustomerByBarcodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_customer_by_barcode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByBarcodeRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerByBarcodeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_customer_by_barcode_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerByBarcodeRequest) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

type GetCustomerByBarcodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *GetCustomerByBarcodeResponse) Reset() {
	*x = GetCustomerByBarcodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_customer_by_barcode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByBarcodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByBarcodeResponse) ProtoMessage() {}

func (x *GetCustomerByBarcodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_customer_by_barcode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByBarcodeResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerByBarcodeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_customer_by_barcode_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomerByBarcodeResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

var File_rpc_get_customer_by_barcode_proto protoreflect.FileDescriptor

var file_rpc_get_customer_by_barcode_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x48, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42,
	0x79, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x74, 0x63, 0x6f, 0x6e, 0x67, 0x31,
	0x34, 0x31, 0x31, 0x2f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2d, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74,
	0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_get_customer_by_barcode_proto_rawDescOnce sync.Once
	file_rpc_get_customer_by_barcode_proto_rawDescData = file_rpc_get_customer_by_barcode_proto_rawDesc
)

func file_rpc_get_customer_by_barcode_proto_rawDescGZIP() []byte {
	file_rpc_get_customer_by_barcode_proto_rawDescOnce.Do(func() {
		file_rpc_get_customer_by_barcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_get_customer_by_barcode_proto_rawDescData)
	})
	return file_rpc_get_customer_by_barcode_proto_rawDescData
}

var file_rpc_get_customer_by_barcode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_get_customer_by_barcode_proto_goTypes = []interface{}{
	(*GetCustomerByBarcodeRequest)(nil),  // 0: pb.GetCustomerByBarcodeRequest
	(*GetCustomerByBarcodeResponse)(nil), // 1: pb.GetCustomerByBarcodeResponse
	(*Customer)(nil),                     // 2: pb.Customer
}
var file_rpc_get_customer_by_barcode_proto_depIdxs = []int32{
	2, // 0: pb.GetCustomerByBarcodeResponse.customer:type_name -> pb.Customer
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_get_customer_by_barcode_proto_init() }
func file_rpc_get_customer_by_barcode_proto_init() {
	if File_rpc_get_customer_by_barcode_proto != nil {
		return
	}
	file_customer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_get_customer_by_barcode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerByBarcodeRequest); i {
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
		file_rpc_get_customer_by_barcode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerByBarcodeResponse); i {
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
			RawDescriptor: file_rpc_get_customer_by_barcode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_get_customer_by_barcode_proto_goTypes,
		DependencyIndexes: file_rpc_get_customer_by_barcode_proto_depIdxs,
		MessageInfos:      file_rpc_get_customer_by_barcode_proto_msgTypes,
	}.Build()
	File_rpc_get_customer_by_barcode_proto = out.File
	file_rpc_get_customer_by_barcode_proto_rawDesc = nil
	file_rpc_get_customer_by_barcode_proto_goTypes = nil
	file_rpc_get_customer_by_barcode_proto_depIdxs = nil
}
