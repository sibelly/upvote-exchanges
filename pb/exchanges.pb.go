// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: exchanges.proto

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

type Exchange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExchangeId string `protobuf:"bytes,2,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
	Website    string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Upvotes    int32  `protobuf:"varint,5,opt,name=upvotes,proto3" json:"upvotes,omitempty"`
	Downvotes  int32  `protobuf:"varint,6,opt,name=downvotes,proto3" json:"downvotes,omitempty"`
}

func (x *Exchange) Reset() {
	*x = Exchange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchanges_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exchange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exchange) ProtoMessage() {}

func (x *Exchange) ProtoReflect() protoreflect.Message {
	mi := &file_exchanges_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exchange.ProtoReflect.Descriptor instead.
func (*Exchange) Descriptor() ([]byte, []int) {
	return file_exchanges_proto_rawDescGZIP(), []int{0}
}

func (x *Exchange) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Exchange) GetExchangeId() string {
	if x != nil {
		return x.ExchangeId
	}
	return ""
}

func (x *Exchange) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Exchange) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Exchange) GetUpvotes() int32 {
	if x != nil {
		return x.Upvotes
	}
	return 0
}

func (x *Exchange) GetDownvotes() int32 {
	if x != nil {
		return x.Downvotes
	}
	return 0
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeId string `protobuf:"bytes,1,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchanges_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchanges_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_exchanges_proto_rawDescGZIP(), []int{1}
}

func (x *VoteRequest) GetExchangeId() string {
	if x != nil {
		return x.ExchangeId
	}
	return ""
}

type VoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeId string `protobuf:"bytes,1,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Website    string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	Votes      int32  `protobuf:"varint,4,opt,name=votes,proto3" json:"votes,omitempty"`
}

func (x *VoteResponse) Reset() {
	*x = VoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchanges_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResponse) ProtoMessage() {}

func (x *VoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchanges_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteResponse.ProtoReflect.Descriptor instead.
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return file_exchanges_proto_rawDescGZIP(), []int{2}
}

func (x *VoteResponse) GetExchangeId() string {
	if x != nil {
		return x.ExchangeId
	}
	return ""
}

func (x *VoteResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VoteResponse) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *VoteResponse) GetVotes() int32 {
	if x != nil {
		return x.Votes
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchanges_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_exchanges_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_exchanges_proto_rawDescGZIP(), []int{3}
}

type ReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeId string `protobuf:"bytes,1,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
}

func (x *ReadReq) Reset() {
	*x = ReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchanges_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadReq) ProtoMessage() {}

func (x *ReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_exchanges_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadReq.ProtoReflect.Descriptor instead.
func (*ReadReq) Descriptor() ([]byte, []int) {
	return file_exchanges_proto_rawDescGZIP(), []int{4}
}

func (x *ReadReq) GetExchangeId() string {
	if x != nil {
		return x.ExchangeId
	}
	return ""
}

type ReadRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exchange *Exchange `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *ReadRes) Reset() {
	*x = ReadRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchanges_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRes) ProtoMessage() {}

func (x *ReadRes) ProtoReflect() protoreflect.Message {
	mi := &file_exchanges_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRes.ProtoReflect.Descriptor instead.
func (*ReadRes) Descriptor() ([]byte, []int) {
	return file_exchanges_proto_rawDescGZIP(), []int{5}
}

func (x *ReadRes) GetExchange() *Exchange {
	if x != nil {
		return x.Exchange
	}
	return nil
}

var File_exchanges_proto protoreflect.FileDescriptor

var file_exchanges_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x08, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x0b,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x22, 0x73, 0x0a, 0x0c,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2a, 0x0a, 0x07, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x12, 0x2b, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x32, 0xad,
	0x01, 0x0a, 0x10, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x30, 0x0a, 0x0c,
	0x52, 0x65, 0x61, 0x64, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x28,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x62,
	0x65, 0x6c, 0x6c, 0x79, 0x2f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2d, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exchanges_proto_rawDescOnce sync.Once
	file_exchanges_proto_rawDescData = file_exchanges_proto_rawDesc
)

func file_exchanges_proto_rawDescGZIP() []byte {
	file_exchanges_proto_rawDescOnce.Do(func() {
		file_exchanges_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchanges_proto_rawDescData)
	})
	return file_exchanges_proto_rawDescData
}

var file_exchanges_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_exchanges_proto_goTypes = []interface{}{
	(*Exchange)(nil),     // 0: proto.Exchange
	(*VoteRequest)(nil),  // 1: proto.VoteRequest
	(*VoteResponse)(nil), // 2: proto.VoteResponse
	(*Empty)(nil),        // 3: proto.Empty
	(*ReadReq)(nil),      // 4: proto.ReadReq
	(*ReadRes)(nil),      // 5: proto.ReadRes
}
var file_exchanges_proto_depIdxs = []int32{
	0, // 0: proto.ReadRes.exchange:type_name -> proto.Exchange
	1, // 1: proto.ExchangesService.Upvote:input_type -> proto.VoteRequest
	3, // 2: proto.ExchangesService.ListExchanges:input_type -> proto.Empty
	4, // 3: proto.ExchangesService.ReadExchange:input_type -> proto.ReadReq
	2, // 4: proto.ExchangesService.Upvote:output_type -> proto.VoteResponse
	0, // 5: proto.ExchangesService.ListExchanges:output_type -> proto.Exchange
	5, // 6: proto.ExchangesService.ReadExchange:output_type -> proto.ReadRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_exchanges_proto_init() }
func file_exchanges_proto_init() {
	if File_exchanges_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exchanges_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exchange); i {
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
		file_exchanges_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
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
		file_exchanges_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteResponse); i {
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
		file_exchanges_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_exchanges_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadReq); i {
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
		file_exchanges_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRes); i {
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
			RawDescriptor: file_exchanges_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exchanges_proto_goTypes,
		DependencyIndexes: file_exchanges_proto_depIdxs,
		MessageInfos:      file_exchanges_proto_msgTypes,
	}.Build()
	File_exchanges_proto = out.File
	file_exchanges_proto_rawDesc = nil
	file_exchanges_proto_goTypes = nil
	file_exchanges_proto_depIdxs = nil
}
