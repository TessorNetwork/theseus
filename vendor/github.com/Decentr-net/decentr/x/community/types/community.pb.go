// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: community/community.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Category int32

const (
	Category_CATEGORY_UNDEFINED              Category = 0
	Category_CATEGORY_WORLD_NEWS             Category = 1
	Category_CATEGORY_TRAVEL_AND_TOURISM     Category = 2
	Category_CATEGORY_SCIENCE_AND_TECHNOLOGY Category = 3
	Category_CATEGORY_STRANGE_WORLD          Category = 4
	Category_CATEGORY_ARTS_AND_ENTERTAINMENT Category = 5
	Category_CATEGORY_WRITERS_AND_WRITING    Category = 6
	Category_CATEGORY_HEALTH_AND_FITNESS     Category = 7
	Category_CATEGORY_CRYPTO_AND_BLOCKCHAIN  Category = 8
	Category_CATEGORY_SPORTS                 Category = 9
)

var Category_name = map[int32]string{
	0: "CATEGORY_UNDEFINED",
	1: "CATEGORY_WORLD_NEWS",
	2: "CATEGORY_TRAVEL_AND_TOURISM",
	3: "CATEGORY_SCIENCE_AND_TECHNOLOGY",
	4: "CATEGORY_STRANGE_WORLD",
	5: "CATEGORY_ARTS_AND_ENTERTAINMENT",
	6: "CATEGORY_WRITERS_AND_WRITING",
	7: "CATEGORY_HEALTH_AND_FITNESS",
	8: "CATEGORY_CRYPTO_AND_BLOCKCHAIN",
	9: "CATEGORY_SPORTS",
}

var Category_value = map[string]int32{
	"CATEGORY_UNDEFINED":              0,
	"CATEGORY_WORLD_NEWS":             1,
	"CATEGORY_TRAVEL_AND_TOURISM":     2,
	"CATEGORY_SCIENCE_AND_TECHNOLOGY": 3,
	"CATEGORY_STRANGE_WORLD":          4,
	"CATEGORY_ARTS_AND_ENTERTAINMENT": 5,
	"CATEGORY_WRITERS_AND_WRITING":    6,
	"CATEGORY_HEALTH_AND_FITNESS":     7,
	"CATEGORY_CRYPTO_AND_BLOCKCHAIN":  8,
	"CATEGORY_SPORTS":                 9,
}

func (x Category) String() string {
	return proto.EnumName(Category_name, int32(x))
}

func (Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_663aacd61135a87b, []int{0}
}

type LikeWeight int32

const (
	LikeWeight_LIKE_WEIGHT_ZERO LikeWeight = 0
	LikeWeight_LIKE_WEIGHT_UP   LikeWeight = 1
	LikeWeight_LIKE_WEIGHT_DOWN LikeWeight = -1
)

var LikeWeight_name = map[int32]string{
	0:  "LIKE_WEIGHT_ZERO",
	1:  "LIKE_WEIGHT_UP",
	-1: "LIKE_WEIGHT_DOWN",
}

var LikeWeight_value = map[string]int32{
	"LIKE_WEIGHT_ZERO": 0,
	"LIKE_WEIGHT_UP":   1,
	"LIKE_WEIGHT_DOWN": -1,
}

func (x LikeWeight) String() string {
	return proto.EnumName(LikeWeight_name, int32(x))
}

func (LikeWeight) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_663aacd61135a87b, []int{1}
}

type Params struct {
	Moderators []string       `protobuf:"bytes,1,rep,name=moderators,proto3" json:"moderators,omitempty"`
	FixedGas   FixedGasParams `protobuf:"bytes,2,opt,name=fixed_gas,json=fixedGas,proto3" json:"fixed_gas"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_663aacd61135a87b, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetModerators() []string {
	if m != nil {
		return m.Moderators
	}
	return nil
}

func (m *Params) GetFixedGas() FixedGasParams {
	if m != nil {
		return m.FixedGas
	}
	return FixedGasParams{}
}

type FixedGasParams struct {
	CreatePost uint64 `protobuf:"varint,1,opt,name=create_post,json=createPost,proto3" json:"create_post,omitempty"`
	DeletePost uint64 `protobuf:"varint,2,opt,name=delete_post,json=deletePost,proto3" json:"delete_post,omitempty"`
	SetLike    uint64 `protobuf:"varint,3,opt,name=set_like,json=setLike,proto3" json:"set_like,omitempty"`
	Follow     uint64 `protobuf:"varint,4,opt,name=follow,proto3" json:"follow,omitempty"`
	Unfollow   uint64 `protobuf:"varint,5,opt,name=unfollow,proto3" json:"unfollow,omitempty"`
}

func (m *FixedGasParams) Reset()         { *m = FixedGasParams{} }
func (m *FixedGasParams) String() string { return proto.CompactTextString(m) }
func (*FixedGasParams) ProtoMessage()    {}
func (*FixedGasParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_663aacd61135a87b, []int{1}
}
func (m *FixedGasParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FixedGasParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FixedGasParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FixedGasParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedGasParams.Merge(m, src)
}
func (m *FixedGasParams) XXX_Size() int {
	return m.Size()
}
func (m *FixedGasParams) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedGasParams.DiscardUnknown(m)
}

var xxx_messageInfo_FixedGasParams proto.InternalMessageInfo

func (m *FixedGasParams) GetCreatePost() uint64 {
	if m != nil {
		return m.CreatePost
	}
	return 0
}

func (m *FixedGasParams) GetDeletePost() uint64 {
	if m != nil {
		return m.DeletePost
	}
	return 0
}

func (m *FixedGasParams) GetSetLike() uint64 {
	if m != nil {
		return m.SetLike
	}
	return 0
}

func (m *FixedGasParams) GetFollow() uint64 {
	if m != nil {
		return m.Follow
	}
	return 0
}

func (m *FixedGasParams) GetUnfollow() uint64 {
	if m != nil {
		return m.Unfollow
	}
	return 0
}

type Post struct {
	Owner        string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Uuid         string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title        string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	PreviewImage string   `protobuf:"bytes,4,opt,name=preview_image,json=previewImage,proto3" json:"preview_image,omitempty"`
	Category     Category `protobuf:"varint,5,opt,name=category,proto3,enum=community.Category" json:"category,omitempty"`
	Text         string   `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_663aacd61135a87b, []int{2}
}
func (m *Post) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Post.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return m.Size()
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Post) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetPreviewImage() string {
	if m != nil {
		return m.PreviewImage
	}
	return ""
}

func (m *Post) GetCategory() Category {
	if m != nil {
		return m.Category
	}
	return Category_CATEGORY_UNDEFINED
}

func (m *Post) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Like struct {
	Owner     string     `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	PostOwner string     `protobuf:"bytes,2,opt,name=post_owner,json=postOwner,proto3" json:"post_owner,omitempty"`
	PostUuid  string     `protobuf:"bytes,3,opt,name=post_uuid,json=postUuid,proto3" json:"post_uuid,omitempty"`
	Weight    LikeWeight `protobuf:"varint,4,opt,name=weight,proto3,enum=community.LikeWeight" json:"weight,omitempty"`
}

func (m *Like) Reset()         { *m = Like{} }
func (m *Like) String() string { return proto.CompactTextString(m) }
func (*Like) ProtoMessage()    {}
func (*Like) Descriptor() ([]byte, []int) {
	return fileDescriptor_663aacd61135a87b, []int{3}
}
func (m *Like) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Like) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Like.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Like) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Like.Merge(m, src)
}
func (m *Like) XXX_Size() int {
	return m.Size()
}
func (m *Like) XXX_DiscardUnknown() {
	xxx_messageInfo_Like.DiscardUnknown(m)
}

var xxx_messageInfo_Like proto.InternalMessageInfo

func (m *Like) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Like) GetPostOwner() string {
	if m != nil {
		return m.PostOwner
	}
	return ""
}

func (m *Like) GetPostUuid() string {
	if m != nil {
		return m.PostUuid
	}
	return ""
}

func (m *Like) GetWeight() LikeWeight {
	if m != nil {
		return m.Weight
	}
	return LikeWeight_LIKE_WEIGHT_ZERO
}

func init() {
	proto.RegisterEnum("community.Category", Category_name, Category_value)
	proto.RegisterEnum("community.LikeWeight", LikeWeight_name, LikeWeight_value)
	proto.RegisterType((*Params)(nil), "community.Params")
	proto.RegisterType((*FixedGasParams)(nil), "community.FixedGasParams")
	proto.RegisterType((*Post)(nil), "community.Post")
	proto.RegisterType((*Like)(nil), "community.Like")
}

func init() { proto.RegisterFile("community/community.proto", fileDescriptor_663aacd61135a87b) }

var fileDescriptor_663aacd61135a87b = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xc1, 0x6e, 0xda, 0x4c,
	0x14, 0x85, 0x31, 0x21, 0x04, 0xdf, 0xfc, 0x3f, 0xb5, 0x26, 0x69, 0x4a, 0x92, 0xc6, 0x20, 0xb2,
	0x89, 0x22, 0x25, 0x54, 0xe9, 0xb6, 0x1b, 0xc7, 0x4c, 0xc0, 0x0d, 0xb1, 0xd1, 0x60, 0x8a, 0x92,
	0x8d, 0xe5, 0xc0, 0xe0, 0x58, 0x01, 0x8c, 0xec, 0xa1, 0x24, 0x4f, 0xd0, 0x6d, 0x5f, 0xa1, 0x6f,
	0xd0, 0x27, 0xe8, 0x3a, 0xcb, 0x2c, 0xbb, 0xaa, 0xaa, 0xe4, 0x41, 0x5a, 0x79, 0x86, 0x18, 0x90,
	0xca, 0xea, 0xde, 0x73, 0x3e, 0xe6, 0x9e, 0x3b, 0x96, 0x06, 0xb6, 0xbb, 0xc1, 0x70, 0x38, 0x19,
	0xf9, 0xec, 0xbe, 0x92, 0x54, 0xc7, 0xe3, 0x30, 0x60, 0x01, 0x92, 0x13, 0x61, 0x67, 0xd3, 0x0b,
	0xbc, 0x80, 0xab, 0x95, 0xb8, 0x12, 0x40, 0xb9, 0x0f, 0xd9, 0xa6, 0x1b, 0xba, 0xc3, 0x08, 0xa9,
	0x00, 0xc3, 0xa0, 0x47, 0x43, 0x97, 0x05, 0x61, 0x54, 0x90, 0x4a, 0x2b, 0x07, 0x32, 0x59, 0x50,
	0xd0, 0x07, 0x90, 0xfb, 0xfe, 0x1d, 0xed, 0x39, 0x9e, 0x1b, 0x15, 0xd2, 0x25, 0xe9, 0x60, 0xfd,
	0x64, 0xfb, 0x78, 0x3e, 0xef, 0x2c, 0xf6, 0x6a, 0x6e, 0x24, 0x4e, 0x3b, 0xcd, 0x3c, 0xfc, 0x2a,
	0xa6, 0x48, 0xae, 0x3f, 0x53, 0xcb, 0xdf, 0x24, 0xc8, 0x2f, 0x23, 0xa8, 0x08, 0xeb, 0xdd, 0x90,
	0xba, 0x8c, 0x3a, 0xe3, 0x20, 0x62, 0x05, 0xa9, 0x24, 0x1d, 0x64, 0x08, 0x08, 0xa9, 0x19, 0x44,
	0x2c, 0x06, 0x7a, 0x74, 0x40, 0x5f, 0x80, 0xb4, 0x00, 0x84, 0xc4, 0x81, 0x6d, 0xc8, 0x45, 0x94,
	0x39, 0x03, 0xff, 0x96, 0x16, 0x56, 0xb8, 0xbb, 0x16, 0x51, 0xd6, 0xf0, 0x6f, 0x29, 0xda, 0x82,
	0x6c, 0x3f, 0x18, 0x0c, 0x82, 0x69, 0x21, 0xc3, 0x8d, 0x59, 0x87, 0x76, 0x20, 0x37, 0x19, 0xcd,
	0x9c, 0x55, 0xee, 0x24, 0x7d, 0xf9, 0xbb, 0x04, 0x19, 0x7e, 0xee, 0x26, 0xac, 0x06, 0xd3, 0x11,
	0x0d, 0x79, 0x26, 0x99, 0x88, 0x06, 0x21, 0xc8, 0x4c, 0x26, 0x7e, 0x8f, 0xe7, 0x90, 0x09, 0xaf,
	0x63, 0x92, 0xf9, 0x6c, 0x20, 0xc6, 0xcb, 0x44, 0x34, 0x68, 0x1f, 0xfe, 0x1f, 0x87, 0xf4, 0xb3,
	0x4f, 0xa7, 0x8e, 0x3f, 0x74, 0x3d, 0xca, 0x33, 0xc8, 0xe4, 0xbf, 0x99, 0x68, 0xc4, 0x1a, 0xaa,
	0x40, 0xae, 0xeb, 0x32, 0xea, 0x05, 0xe1, 0x3d, 0x4f, 0x92, 0x3f, 0xd9, 0x58, 0xb8, 0x4e, 0x7d,
	0x66, 0x91, 0x04, 0x8a, 0xe7, 0x33, 0x7a, 0xc7, 0x0a, 0x59, 0x31, 0x3f, 0xae, 0xcb, 0x5f, 0x24,
	0xc8, 0xf0, 0x7d, 0xff, 0x1d, 0x79, 0x0f, 0x20, 0xbe, 0x3a, 0x47, 0x58, 0x22, 0xb8, 0x1c, 0x2b,
	0x16, 0xb7, 0x77, 0x81, 0x37, 0x0e, 0x5f, 0x4b, 0x6c, 0x90, 0x8b, 0x85, 0x76, 0xbc, 0xda, 0x11,
	0x64, 0xa7, 0xd4, 0xf7, 0x6e, 0x18, 0x4f, 0x9f, 0x3f, 0x79, 0xbd, 0x90, 0x2e, 0x1e, 0xd9, 0xe1,
	0x26, 0x99, 0x41, 0x87, 0x3f, 0xd2, 0x90, 0x7b, 0x09, 0x8d, 0xb6, 0x00, 0xe9, 0x9a, 0x8d, 0x6b,
	0x16, 0xb9, 0x74, 0xda, 0x66, 0x15, 0x9f, 0x19, 0x26, 0xae, 0x2a, 0x29, 0xf4, 0x06, 0x36, 0x12,
	0xbd, 0x63, 0x91, 0x46, 0xd5, 0x31, 0x71, 0xa7, 0xa5, 0x48, 0xa8, 0x08, 0xbb, 0x89, 0x61, 0x13,
	0xed, 0x13, 0x6e, 0x38, 0x9a, 0x59, 0x75, 0x6c, 0xab, 0x4d, 0x8c, 0xd6, 0x85, 0x92, 0x46, 0xfb,
	0x50, 0x4c, 0x80, 0x96, 0x6e, 0x60, 0x53, 0xc7, 0x82, 0xc0, 0x7a, 0xdd, 0xb4, 0x1a, 0x56, 0xed,
	0x52, 0x59, 0x41, 0x3b, 0xb0, 0x35, 0x87, 0x6c, 0xa2, 0x99, 0x35, 0x2c, 0xc6, 0x28, 0x99, 0xa5,
	0x03, 0x34, 0x62, 0xb7, 0xf8, 0xbf, 0xb1, 0x69, 0x63, 0x62, 0x6b, 0x86, 0x79, 0x81, 0x4d, 0x5b,
	0x59, 0x45, 0x25, 0x78, 0x3b, 0xcf, 0x47, 0x0c, 0x1b, 0x13, 0xc1, 0xc5, 0xb5, 0x61, 0xd6, 0x94,
	0xec, 0x52, 0xd0, 0x3a, 0xd6, 0x1a, 0x76, 0x9d, 0x03, 0x67, 0x86, 0x6d, 0xe2, 0x56, 0x4b, 0x59,
	0x43, 0x65, 0x50, 0x13, 0x40, 0x27, 0x97, 0x4d, 0xdb, 0xe2, 0xc0, 0x69, 0xc3, 0xd2, 0xcf, 0xf5,
	0xba, 0x66, 0x98, 0x4a, 0x0e, 0x6d, 0xc0, 0xab, 0x79, 0xce, 0xa6, 0x45, 0xec, 0x96, 0x22, 0x1f,
	0xb6, 0x01, 0xe6, 0xd7, 0x8a, 0x36, 0x41, 0x69, 0x18, 0xe7, 0xd8, 0xe9, 0x60, 0xa3, 0x56, 0xb7,
	0x9d, 0x2b, 0x4c, 0x2c, 0x25, 0x85, 0x10, 0xe4, 0x17, 0xd5, 0x76, 0x53, 0x91, 0xd0, 0xde, 0x32,
	0x59, 0xb5, 0x3a, 0xa6, 0xf2, 0xe7, 0xe5, 0x27, 0x9d, 0x7e, 0x7c, 0x78, 0x52, 0xa5, 0xc7, 0x27,
	0x55, 0xfa, 0xfd, 0xa4, 0x4a, 0x5f, 0x9f, 0xd5, 0xd4, 0xe3, 0xb3, 0x9a, 0xfa, 0xf9, 0xac, 0xa6,
	0xae, 0xde, 0x79, 0x3e, 0xbb, 0x99, 0x5c, 0xc7, 0x9f, 0xb5, 0x52, 0xa5, 0x5d, 0x3a, 0x62, 0xe1,
	0xd1, 0x88, 0xb2, 0x4a, 0x4f, 0xd4, 0x95, 0xbb, 0xf9, 0x6b, 0x52, 0x61, 0xf7, 0x63, 0x1a, 0x5d,
	0x67, 0xf9, 0x9b, 0xf1, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0xed, 0xf8, 0xa7, 0x71,
	0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.FixedGas.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCommunity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Moderators) > 0 {
		for iNdEx := len(m.Moderators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Moderators[iNdEx])
			copy(dAtA[i:], m.Moderators[iNdEx])
			i = encodeVarintCommunity(dAtA, i, uint64(len(m.Moderators[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FixedGasParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FixedGasParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FixedGasParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Unfollow != 0 {
		i = encodeVarintCommunity(dAtA, i, uint64(m.Unfollow))
		i--
		dAtA[i] = 0x28
	}
	if m.Follow != 0 {
		i = encodeVarintCommunity(dAtA, i, uint64(m.Follow))
		i--
		dAtA[i] = 0x20
	}
	if m.SetLike != 0 {
		i = encodeVarintCommunity(dAtA, i, uint64(m.SetLike))
		i--
		dAtA[i] = 0x18
	}
	if m.DeletePost != 0 {
		i = encodeVarintCommunity(dAtA, i, uint64(m.DeletePost))
		i--
		dAtA[i] = 0x10
	}
	if m.CreatePost != 0 {
		i = encodeVarintCommunity(dAtA, i, uint64(m.CreatePost))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Post) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Post) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Post) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintCommunity(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x32
	}
	if m.Category != 0 {
		i = encodeVarintCommunity(dAtA, i, uint64(m.Category))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PreviewImage) > 0 {
		i -= len(m.PreviewImage)
		copy(dAtA[i:], m.PreviewImage)
		i = encodeVarintCommunity(dAtA, i, uint64(len(m.PreviewImage)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintCommunity(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintCommunity(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCommunity(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Like) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Like) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Like) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Weight != 0 {
		i = encodeVarintCommunity(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PostUuid) > 0 {
		i -= len(m.PostUuid)
		copy(dAtA[i:], m.PostUuid)
		i = encodeVarintCommunity(dAtA, i, uint64(len(m.PostUuid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PostOwner) > 0 {
		i -= len(m.PostOwner)
		copy(dAtA[i:], m.PostOwner)
		i = encodeVarintCommunity(dAtA, i, uint64(len(m.PostOwner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCommunity(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommunity(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommunity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Moderators) > 0 {
		for _, s := range m.Moderators {
			l = len(s)
			n += 1 + l + sovCommunity(uint64(l))
		}
	}
	l = m.FixedGas.Size()
	n += 1 + l + sovCommunity(uint64(l))
	return n
}

func (m *FixedGasParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreatePost != 0 {
		n += 1 + sovCommunity(uint64(m.CreatePost))
	}
	if m.DeletePost != 0 {
		n += 1 + sovCommunity(uint64(m.DeletePost))
	}
	if m.SetLike != 0 {
		n += 1 + sovCommunity(uint64(m.SetLike))
	}
	if m.Follow != 0 {
		n += 1 + sovCommunity(uint64(m.Follow))
	}
	if m.Unfollow != 0 {
		n += 1 + sovCommunity(uint64(m.Unfollow))
	}
	return n
}

func (m *Post) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCommunity(uint64(l))
	}
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovCommunity(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovCommunity(uint64(l))
	}
	l = len(m.PreviewImage)
	if l > 0 {
		n += 1 + l + sovCommunity(uint64(l))
	}
	if m.Category != 0 {
		n += 1 + sovCommunity(uint64(m.Category))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovCommunity(uint64(l))
	}
	return n
}

func (m *Like) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCommunity(uint64(l))
	}
	l = len(m.PostOwner)
	if l > 0 {
		n += 1 + l + sovCommunity(uint64(l))
	}
	l = len(m.PostUuid)
	if l > 0 {
		n += 1 + l + sovCommunity(uint64(l))
	}
	if m.Weight != 0 {
		n += 1 + sovCommunity(uint64(m.Weight))
	}
	return n
}

func sovCommunity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommunity(x uint64) (n int) {
	return sovCommunity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommunity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moderators", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommunity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommunity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moderators = append(m.Moderators, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedGas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommunity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommunity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FixedGas.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommunity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommunity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FixedGasParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommunity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FixedGasParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FixedGasParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatePost", wireType)
			}
			m.CreatePost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatePost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletePost", wireType)
			}
			m.DeletePost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeletePost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetLike", wireType)
			}
			m.SetLike = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SetLike |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Follow", wireType)
			}
			m.Follow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Follow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unfollow", wireType)
			}
			m.Unfollow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Unfollow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommunity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommunity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Post) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommunity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Post: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Post: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommunity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommunity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommunity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommunity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommunity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommunity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviewImage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommunity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommunity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviewImage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			m.Category = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Category |= Category(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommunity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommunity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommunity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommunity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Like) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommunity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Like: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Like: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommunity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommunity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommunity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommunity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostUuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommunity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommunity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostUuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= LikeWeight(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommunity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommunity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommunity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommunity
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommunity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCommunity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommunity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommunity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommunity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommunity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommunity = fmt.Errorf("proto: unexpected end of group")
)
