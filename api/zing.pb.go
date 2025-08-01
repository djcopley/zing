// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: api/zing.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_api_zing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_zing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_api_zing_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_api_zing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_zing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_api_zing_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_api_zing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_zing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_api_zing_proto_rawDescGZIP(), []int{2}
}

type LogoutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_api_zing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_zing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_api_zing_proto_rawDescGZIP(), []int{3}
}

// GetMessageRequest sends the server the user's id and a token
type ListMessagesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of messages to return. The service may return fewer than
	// this value.
	// If unspecified, at most 50 messages will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListMessages` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListMessages` must match
	// the call that provided the page token.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListMessagesRequest) Reset() {
	*x = ListMessagesRequest{}
	mi := &file_api_zing_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesRequest) ProtoMessage() {}

func (x *ListMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_zing_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesRequest.ProtoReflect.Descriptor instead.
func (*ListMessagesRequest) Descriptor() ([]byte, []int) {
	return file_api_zing_proto_rawDescGZIP(), []int{4}
}

func (x *ListMessagesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListMessagesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListMessagesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The messages for the user.
	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListMessagesResponse) Reset() {
	*x = ListMessagesResponse{}
	mi := &file_api_zing_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesResponse) ProtoMessage() {}

func (x *ListMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_zing_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesResponse.ProtoReflect.Descriptor instead.
func (*ListMessagesResponse) Descriptor() ([]byte, []int) {
	return file_api_zing_proto_rawDescGZIP(), []int{5}
}

func (x *ListMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ListMessagesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type SendMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	To            *User                  `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Message       *Message               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	mi := &file_api_zing_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_zing_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_zing_proto_rawDescGZIP(), []int{6}
}

func (x *SendMessageRequest) GetTo() *User {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SendMessageRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type SendMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	mi := &file_api_zing_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_zing_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_api_zing_proto_rawDescGZIP(), []int{7}
}

var File_api_zing_proto protoreflect.FileDescriptor

const file_api_zing_proto_rawDesc = "" +
	"\n" +
	"\x0eapi/zing.proto\x12\x04zing\x1a\x0eapi/user.proto\x1a\x11api/message.proto\"F\n" +
	"\fLoginRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"%\n" +
	"\rLoginResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"\x0f\n" +
	"\rLogoutRequest\"\x10\n" +
	"\x0eLogoutResponse\"Q\n" +
	"\x13ListMessagesRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tR\tpageToken\"i\n" +
	"\x14ListMessagesResponse\x12)\n" +
	"\bmessages\x18\x01 \x03(\v2\r.zing.MessageR\bmessages\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"Y\n" +
	"\x12SendMessageRequest\x12\x1a\n" +
	"\x02to\x18\x01 \x01(\v2\n" +
	".zing.UserR\x02to\x12'\n" +
	"\amessage\x18\x02 \x01(\v2\r.zing.MessageR\amessage\"\x15\n" +
	"\x13SendMessageResponse2\xf8\x01\n" +
	"\x04Zing\x120\n" +
	"\x05Login\x12\x12.zing.LoginRequest\x1a\x13.zing.LoginResponse\x123\n" +
	"\x06Logout\x12\x13.zing.LogoutRequest\x1a\x14.zing.LogoutResponse\x12E\n" +
	"\fListMessages\x12\x19.zing.ListMessagesRequest\x1a\x1a.zing.ListMessagesResponse\x12B\n" +
	"\vSendMessage\x12\x18.zing.SendMessageRequest\x1a\x19.zing.SendMessageResponseB\x1eZ\x1cgithub.com/djcopley/zing/apib\x06proto3"

var (
	file_api_zing_proto_rawDescOnce sync.Once
	file_api_zing_proto_rawDescData []byte
)

func file_api_zing_proto_rawDescGZIP() []byte {
	file_api_zing_proto_rawDescOnce.Do(func() {
		file_api_zing_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_zing_proto_rawDesc), len(file_api_zing_proto_rawDesc)))
	})
	return file_api_zing_proto_rawDescData
}

var file_api_zing_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_zing_proto_goTypes = []any{
	(*LoginRequest)(nil),         // 0: zing.LoginRequest
	(*LoginResponse)(nil),        // 1: zing.LoginResponse
	(*LogoutRequest)(nil),        // 2: zing.LogoutRequest
	(*LogoutResponse)(nil),       // 3: zing.LogoutResponse
	(*ListMessagesRequest)(nil),  // 4: zing.ListMessagesRequest
	(*ListMessagesResponse)(nil), // 5: zing.ListMessagesResponse
	(*SendMessageRequest)(nil),   // 6: zing.SendMessageRequest
	(*SendMessageResponse)(nil),  // 7: zing.SendMessageResponse
	(*Message)(nil),              // 8: zing.Message
	(*User)(nil),                 // 9: zing.User
}
var file_api_zing_proto_depIdxs = []int32{
	8, // 0: zing.ListMessagesResponse.messages:type_name -> zing.Message
	9, // 1: zing.SendMessageRequest.to:type_name -> zing.User
	8, // 2: zing.SendMessageRequest.message:type_name -> zing.Message
	0, // 3: zing.Zing.Login:input_type -> zing.LoginRequest
	2, // 4: zing.Zing.Logout:input_type -> zing.LogoutRequest
	4, // 5: zing.Zing.ListMessages:input_type -> zing.ListMessagesRequest
	6, // 6: zing.Zing.SendMessage:input_type -> zing.SendMessageRequest
	1, // 7: zing.Zing.Login:output_type -> zing.LoginResponse
	3, // 8: zing.Zing.Logout:output_type -> zing.LogoutResponse
	5, // 9: zing.Zing.ListMessages:output_type -> zing.ListMessagesResponse
	7, // 10: zing.Zing.SendMessage:output_type -> zing.SendMessageResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_zing_proto_init() }
func file_api_zing_proto_init() {
	if File_api_zing_proto != nil {
		return
	}
	file_api_user_proto_init()
	file_api_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_zing_proto_rawDesc), len(file_api_zing_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_zing_proto_goTypes,
		DependencyIndexes: file_api_zing_proto_depIdxs,
		MessageInfos:      file_api_zing_proto_msgTypes,
	}.Build()
	File_api_zing_proto = out.File
	file_api_zing_proto_goTypes = nil
	file_api_zing_proto_depIdxs = nil
}
