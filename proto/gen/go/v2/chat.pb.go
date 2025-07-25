package v2
import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)
const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)
type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NickName      string                 `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *User) Reset() {
	*x = User{}
	mi := &file_proto_v2_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*User) ProtoMessage() {}
func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v2_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_v2_chat_proto_rawDescGZIP(), []int{0}
}
func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}
func (x *User) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}
func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}
func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}
func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}
type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MessageId     string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	SenderId      string                 `protobuf:"bytes,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	GroupId       string                 `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Text          string                 `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *Message) Reset() {
	*x = Message{}
	mi := &file_proto_v2_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*Message) ProtoMessage() {}
func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v2_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_v2_chat_proto_rawDescGZIP(), []int{1}
}
func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}
func (x *Message) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}
func (x *Message) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}
func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}
func (x *Message) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}
type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NickName      string                 `protobuf:"bytes,1,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_proto_v2_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*CreateUserRequest) ProtoMessage() {}
func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v2_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_v2_chat_proto_rawDescGZIP(), []int{2}
}
func (x *CreateUserRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}
func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}
func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}
type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_proto_v2_chat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*GetUserRequest) ProtoMessage() {}
func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v2_chat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_v2_chat_proto_rawDescGZIP(), []int{3}
}
func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}
type CreateMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SenderId      string                 `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	GroupId       string                 `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Text          string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *CreateMessageRequest) Reset() {
	*x = CreateMessageRequest{}
	mi := &file_proto_v2_chat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *CreateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*CreateMessageRequest) ProtoMessage() {}
func (x *CreateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v2_chat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_v2_chat_proto_rawDescGZIP(), []int{4}
}
func (x *CreateMessageRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}
func (x *CreateMessageRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}
func (x *CreateMessageRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}
type GetMessagesByGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GroupId       string                 `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *GetMessagesByGroupRequest) Reset() {
	*x = GetMessagesByGroupRequest{}
	mi := &file_proto_v2_chat_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *GetMessagesByGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*GetMessagesByGroupRequest) ProtoMessage() {}
func (x *GetMessagesByGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v2_chat_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*GetMessagesByGroupRequest) Descriptor() ([]byte, []int) {
	return file_proto_v2_chat_proto_rawDescGZIP(), []int{5}
}
func (x *GetMessagesByGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}
func (x *GetMessagesByGroupRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}
func (x *GetMessagesByGroupRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}
var File_proto_v2_chat_proto protoreflect.FileDescriptor
const file_proto_v2_chat_proto_rawDesc = "" +
	"\n" +
	"\x13proto/v2/chat.proto\x12\x02v2\x1a\x1fgoogle/protobuf/timestamp.proto\"\xbf\x01\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tnick_name\x18\x02 \x01(\tR\bnickName\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"\xaf\x01\n" +
	"\aMessage\x12\x1d\n" +
	"\n" +
	"message_id\x18\x01 \x01(\tR\tmessageId\x12\x1b\n" +
	"\tsender_id\x18\x02 \x01(\tR\bsenderId\x12\x19\n" +
	"\bgroup_id\x18\x03 \x01(\tR\agroupId\x12\x12\n" +
	"\x04text\x18\x04 \x01(\tR\x04text\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"b\n" +
	"\x11CreateUserRequest\x12\x1b\n" +
	"\tnick_name\x18\x01 \x01(\tR\bnickName\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\" \n" +
	"\x0eGetUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"b\n" +
	"\x14CreateMessageRequest\x12\x1b\n" +
	"\tsender_id\x18\x01 \x01(\tR\bsenderId\x12\x19\n" +
	"\bgroup_id\x18\x02 \x01(\tR\agroupId\x12\x12\n" +
	"\x04text\x18\x03 \x01(\tR\x04text\"d\n" +
	"\x19GetMessagesByGroupRequest\x12\x19\n" +
	"\bgroup_id\x18\x01 \x01(\tR\agroupId\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\x03 \x01(\x05R\x06offset2i\n" +
	"\vUserService\x12/\n" +
	"\n" +
	"CreateUser\x12\x15.v2.CreateUserRequest\x1a\b.v2.User\"\x00\x12)\n" +
	"\aGetUser\x12\x12.v2.GetUserRequest\x1a\b.v2.User\"\x002\x90\x01\n" +
	"\x0eMessageService\x128\n" +
	"\rCreateMessage\x12\x18.v2.CreateMessageRequest\x1a\v.v2.Message\"\x00\x12D\n" +
	"\x12GetMessagesByGroup\x12\x1d.v2.GetMessagesByGroupRequest\x1a\v.v2.Message\"\x000\x01B#Z!MuchUp/backend/proto/gen/go/v2;v2b\x06proto3"
var (
	file_proto_v2_chat_proto_rawDescOnce sync.Once
	file_proto_v2_chat_proto_rawDescData []byte
)
func file_proto_v2_chat_proto_rawDescGZIP() []byte {
	file_proto_v2_chat_proto_rawDescOnce.Do(func() {
		file_proto_v2_chat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_v2_chat_proto_rawDesc), len(file_proto_v2_chat_proto_rawDesc)))
	})
	return file_proto_v2_chat_proto_rawDescData
}
var file_proto_v2_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_v2_chat_proto_goTypes = []any{
	(*User)(nil),
	(*Message)(nil),
	(*CreateUserRequest)(nil),
	(*GetUserRequest)(nil),
	(*CreateMessageRequest)(nil),
	(*GetMessagesByGroupRequest)(nil),
	(*timestamppb.Timestamp)(nil),
}
var file_proto_v2_chat_proto_depIdxs = []int32{
	6,
	6,
	6,
	2,
	3,
	4,
	5,
	0,
	0,
	1,
	1,
	7,
	3,
	3,
	3,
	0,
}
func init() { file_proto_v2_chat_proto_init() }
func file_proto_v2_chat_proto_init() {
	if File_proto_v2_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_v2_chat_proto_rawDesc), len(file_proto_v2_chat_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_v2_chat_proto_goTypes,
		DependencyIndexes: file_proto_v2_chat_proto_depIdxs,
		MessageInfos:      file_proto_v2_chat_proto_msgTypes,
	}.Build()
	File_proto_v2_chat_proto = out.File
	file_proto_v2_chat_proto_goTypes = nil
	file_proto_v2_chat_proto_depIdxs = nil
}
