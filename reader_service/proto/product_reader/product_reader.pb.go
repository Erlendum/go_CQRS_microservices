// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: product_reader.proto

package readerService

import (
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

var File_product_reader_proto protoreflect.FileDescriptor

var file_product_reader_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xaf, 0x03, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x72, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x20,
	0x2e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x20, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x12, 0x43, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x12, 0x5d, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x23, 0x2e, 0x72,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x23, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x3b, 0x72, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_product_reader_proto_goTypes = []interface{}{
	(*CreateProductReq)(nil),     // 0: readerService.CreateProductReq
	(*UpdateProductReq)(nil),     // 1: readerService.UpdateProductReq
	(*GetProductByIdReq)(nil),    // 2: readerService.GetProductByIdReq
	(*SearchReq)(nil),            // 3: readerService.SearchReq
	(*DeleteProductByIdReq)(nil), // 4: readerService.DeleteProductByIdReq
	(*CreateProductRes)(nil),     // 5: readerService.CreateProductRes
	(*UpdateProductRes)(nil),     // 6: readerService.UpdateProductRes
	(*GetProductByIdRes)(nil),    // 7: readerService.GetProductByIdRes
	(*SearchRes)(nil),            // 8: readerService.SearchRes
	(*DeleteProductByIdRes)(nil), // 9: readerService.DeleteProductByIdRes
}
var file_product_reader_proto_depIdxs = []int32{
	0, // 0: readerService.readerService.CreateProduct:input_type -> readerService.CreateProductReq
	1, // 1: readerService.readerService.UpdateProduct:input_type -> readerService.UpdateProductReq
	2, // 2: readerService.readerService.GetProductById:input_type -> readerService.GetProductByIdReq
	3, // 3: readerService.readerService.SearchProduct:input_type -> readerService.SearchReq
	4, // 4: readerService.readerService.DeleteProductByID:input_type -> readerService.DeleteProductByIdReq
	5, // 5: readerService.readerService.CreateProduct:output_type -> readerService.CreateProductRes
	6, // 6: readerService.readerService.UpdateProduct:output_type -> readerService.UpdateProductRes
	7, // 7: readerService.readerService.GetProductById:output_type -> readerService.GetProductByIdRes
	8, // 8: readerService.readerService.SearchProduct:output_type -> readerService.SearchRes
	9, // 9: readerService.readerService.DeleteProductByID:output_type -> readerService.DeleteProductByIdRes
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_product_reader_proto_init() }
func file_product_reader_proto_init() {
	if File_product_reader_proto != nil {
		return
	}
	file_product_reader_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_reader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_reader_proto_goTypes,
		DependencyIndexes: file_product_reader_proto_depIdxs,
	}.Build()
	File_product_reader_proto = out.File
	file_product_reader_proto_rawDesc = nil
	file_product_reader_proto_goTypes = nil
	file_product_reader_proto_depIdxs = nil
}