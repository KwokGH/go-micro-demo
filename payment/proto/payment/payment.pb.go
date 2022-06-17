// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/payment/payment.proto

package payment

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

type PaymentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PaymentName   string `protobuf:"bytes,2,opt,name=payment_name,json=paymentName,proto3" json:"payment_name,omitempty"`
	PaymentSid    string `protobuf:"bytes,3,opt,name=payment_sid,json=paymentSid,proto3" json:"payment_sid,omitempty"`
	PaymentStatus bool   `protobuf:"varint,4,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty"`
	PaymentImage  string `protobuf:"bytes,5,opt,name=Payment_image,json=PaymentImage,proto3" json:"Payment_image,omitempty"`
}

func (x *PaymentInfo) Reset() {
	*x = PaymentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentInfo) ProtoMessage() {}

func (x *PaymentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentInfo.ProtoReflect.Descriptor instead.
func (*PaymentInfo) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentInfo) GetPaymentName() string {
	if x != nil {
		return x.PaymentName
	}
	return ""
}

func (x *PaymentInfo) GetPaymentSid() string {
	if x != nil {
		return x.PaymentSid
	}
	return ""
}

func (x *PaymentInfo) GetPaymentStatus() bool {
	if x != nil {
		return x.PaymentStatus
	}
	return false
}

func (x *PaymentInfo) GetPaymentImage() string {
	if x != nil {
		return x.PaymentImage
	}
	return ""
}

type PaymentID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId int64 `protobuf:"varint,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
}

func (x *PaymentID) Reset() {
	*x = PaymentID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentID) ProtoMessage() {}

func (x *PaymentID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentID.ProtoReflect.Descriptor instead.
func (*PaymentID) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentID) GetPaymentId() int64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type All struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *All) Reset() {
	*x = All{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *All) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*All) ProtoMessage() {}

func (x *All) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use All.ProtoReflect.Descriptor instead.
func (*All) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{3}
}

type PaymentAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentInfo []*PaymentInfo `protobuf:"bytes,1,rep,name=payment_info,json=paymentInfo,proto3" json:"payment_info,omitempty"`
}

func (x *PaymentAll) Reset() {
	*x = PaymentAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentAll) ProtoMessage() {}

func (x *PaymentAll) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentAll.ProtoReflect.Descriptor instead.
func (*PaymentAll) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{4}
}

func (x *PaymentAll) GetPaymentInfo() []*PaymentInfo {
	if x != nil {
		return x.PaymentInfo
	}
	return nil
}

var File_proto_payment_payment_proto protoreflect.FileDescriptor

var file_proto_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xad, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x05, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x22, 0x45, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x37, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xb3,
	0x02, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x12, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x12, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x0e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0c, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x6c, 0x6c, 0x1a, 0x13, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x6c, 0x6c, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_payment_payment_proto_rawDescOnce sync.Once
	file_proto_payment_payment_proto_rawDescData = file_proto_payment_payment_proto_rawDesc
)

func file_proto_payment_payment_proto_rawDescGZIP() []byte {
	file_proto_payment_payment_proto_rawDescOnce.Do(func() {
		file_proto_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_payment_payment_proto_rawDescData)
	})
	return file_proto_payment_payment_proto_rawDescData
}

var file_proto_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_payment_payment_proto_goTypes = []interface{}{
	(*PaymentInfo)(nil), // 0: payment.PaymentInfo
	(*PaymentID)(nil),   // 1: payment.PaymentID
	(*Response)(nil),    // 2: payment.Response
	(*All)(nil),         // 3: payment.All
	(*PaymentAll)(nil),  // 4: payment.PaymentAll
}
var file_proto_payment_payment_proto_depIdxs = []int32{
	0, // 0: payment.PaymentAll.payment_info:type_name -> payment.PaymentInfo
	0, // 1: payment.Payment.AddPayment:input_type -> payment.PaymentInfo
	0, // 2: payment.Payment.UpdatePayment:input_type -> payment.PaymentInfo
	1, // 3: payment.Payment.DeletePaymentByID:input_type -> payment.PaymentID
	1, // 4: payment.Payment.FindPaymentByID:input_type -> payment.PaymentID
	3, // 5: payment.Payment.FindAllPayment:input_type -> payment.All
	1, // 6: payment.Payment.AddPayment:output_type -> payment.PaymentID
	2, // 7: payment.Payment.UpdatePayment:output_type -> payment.Response
	2, // 8: payment.Payment.DeletePaymentByID:output_type -> payment.Response
	0, // 9: payment.Payment.FindPaymentByID:output_type -> payment.PaymentInfo
	4, // 10: payment.Payment.FindAllPayment:output_type -> payment.PaymentAll
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_payment_payment_proto_init() }
func file_proto_payment_payment_proto_init() {
	if File_proto_payment_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_payment_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentInfo); i {
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
		file_proto_payment_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentID); i {
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
		file_proto_payment_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_payment_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*All); i {
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
		file_proto_payment_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentAll); i {
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
			RawDescriptor: file_proto_payment_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_payment_payment_proto_goTypes,
		DependencyIndexes: file_proto_payment_payment_proto_depIdxs,
		MessageInfos:      file_proto_payment_payment_proto_msgTypes,
	}.Build()
	File_proto_payment_payment_proto = out.File
	file_proto_payment_payment_proto_rawDesc = nil
	file_proto_payment_payment_proto_goTypes = nil
	file_proto_payment_payment_proto_depIdxs = nil
}
