// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v5.26.1
// source: service_rapid_loyalty.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_rapid_loyalty_proto protoreflect.FileDescriptor

var file_service_rapid_loyalty_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x5f,
	0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x72, 0x70,
	0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x62,
	0x79, 0x5f, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xcd, 0x04, 0x0a, 0x0c, 0x52, 0x61, 0x70, 0x69, 0x64, 0x4c, 0x6f, 0x79, 0x61,
	0x6c, 0x74, 0x79, 0x12, 0xa0, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x92, 0x41, 0x40, 0x12, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x1a, 0x26, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x6e, 0x65, 0x77, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0xc5, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x42, 0x79, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x42, 0x79, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6a, 0x92, 0x41, 0x48, 0x12, 0x1a, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x62, 0x61, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x1a, 0x2a, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x7d, 0x12, 0xd1,
	0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x42, 0x79, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x42, 0x79, 0x42, 0x61,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x42,
	0x79, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x70, 0x92, 0x41, 0x4c, 0x12, 0x1c, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x20, 0x62, 0x79, 0x20, 0x62, 0x61, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x1a, 0x2c, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x20, 0x62, 0x79, 0x20, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x2f, 0x7b, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x7d, 0x42, 0x85, 0x01, 0x92, 0x41, 0x5a, 0x12, 0x58, 0x0a, 0x11, 0x52, 0x61, 0x70, 0x69,
	0x64, 0x20, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x20, 0x41, 0x50, 0x49, 0x22, 0x3e, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x67, 0x20, 0x4c, 0x65, 0x12, 0x1d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x74, 0x63,
	0x6f, 0x6e, 0x67, 0x31, 0x34, 0x31, 0x31, 0x1a, 0x14, 0x6c, 0x74, 0x63, 0x6f, 0x6e, 0x67, 0x31,
	0x34, 0x31, 0x31, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31,
	0x2e, 0x30, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x74, 0x63, 0x6f, 0x6e, 0x67, 0x31, 0x34, 0x31, 0x31, 0x2f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2d,
	0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_service_rapid_loyalty_proto_goTypes = []interface{}{
	(*CreateCustomerRequest)(nil),          // 0: pb.CreateCustomerRequest
	(*GetCustomerByBarcodeRequest)(nil),    // 1: pb.GetCustomerByBarcodeRequest
	(*GetMembershipByBarcodeRequest)(nil),  // 2: pb.GetMembershipByBarcodeRequest
	(*CreateCustomerResponse)(nil),         // 3: pb.CreateCustomerResponse
	(*GetCustomerByBarcodeResponse)(nil),   // 4: pb.GetCustomerByBarcodeResponse
	(*GetMembershipByBarcodeResponse)(nil), // 5: pb.GetMembershipByBarcodeResponse
}
var file_service_rapid_loyalty_proto_depIdxs = []int32{
	0, // 0: pb.RapidLoyalty.CreateUser:input_type -> pb.CreateCustomerRequest
	1, // 1: pb.RapidLoyalty.GetCustomerByBarcode:input_type -> pb.GetCustomerByBarcodeRequest
	2, // 2: pb.RapidLoyalty.GetMembershipByBarcode:input_type -> pb.GetMembershipByBarcodeRequest
	3, // 3: pb.RapidLoyalty.CreateUser:output_type -> pb.CreateCustomerResponse
	4, // 4: pb.RapidLoyalty.GetCustomerByBarcode:output_type -> pb.GetCustomerByBarcodeResponse
	5, // 5: pb.RapidLoyalty.GetMembershipByBarcode:output_type -> pb.GetMembershipByBarcodeResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_rapid_loyalty_proto_init() }
func file_service_rapid_loyalty_proto_init() {
	if File_service_rapid_loyalty_proto != nil {
		return
	}
	file_rpc_create_customer_proto_init()
	file_rpc_get_customer_by_barcode_proto_init()
	file_rpc_get_membership_by_barcode_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_rapid_loyalty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_rapid_loyalty_proto_goTypes,
		DependencyIndexes: file_service_rapid_loyalty_proto_depIdxs,
	}.Build()
	File_service_rapid_loyalty_proto = out.File
	file_service_rapid_loyalty_proto_rawDesc = nil
	file_service_rapid_loyalty_proto_goTypes = nil
	file_service_rapid_loyalty_proto_depIdxs = nil
}
