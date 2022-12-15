// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feder.proto

package federpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/milvus-io/milvus-proto/go-api/commonpb"
	schemapb "github.com/milvus-io/milvus-proto/go-api/schemapb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FederSearchRequest struct {
	FederSegments        []int64  `protobuf:"varint,1,rep,packed,name=feder_segments,json=federSegments,proto3" json:"feder_segments,omitempty"`
	SearchTotal          bool     `protobuf:"varint,2,opt,name=search_total,json=searchTotal,proto3" json:"search_total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FederSearchRequest) Reset()         { *m = FederSearchRequest{} }
func (m *FederSearchRequest) String() string { return proto.CompactTextString(m) }
func (*FederSearchRequest) ProtoMessage()    {}
func (*FederSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{0}
}

func (m *FederSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederSearchRequest.Unmarshal(m, b)
}
func (m *FederSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederSearchRequest.Marshal(b, m, deterministic)
}
func (m *FederSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederSearchRequest.Merge(m, src)
}
func (m *FederSearchRequest) XXX_Size() int {
	return xxx_messageInfo_FederSearchRequest.Size(m)
}
func (m *FederSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FederSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FederSearchRequest proto.InternalMessageInfo

func (m *FederSearchRequest) GetFederSegments() []int64 {
	if m != nil {
		return m.FederSegments
	}
	return nil
}

func (m *FederSearchRequest) GetSearchTotal() bool {
	if m != nil {
		return m.SearchTotal
	}
	return false
}

type FederSegmentSearchResult struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	SegmentID            int64            `protobuf:"varint,2,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	VisitInfo            string           `protobuf:"bytes,3,opt,name=visit_info,json=visitInfo,proto3" json:"visit_info,omitempty"`
	Offsets              []int64          `protobuf:"varint,4,rep,packed,name=offsets,proto3" json:"offsets,omitempty"`
	PrimaryKeys          *schemapb.IDs    `protobuf:"bytes,5,opt,name=primary_keys,json=primaryKeys,proto3" json:"primary_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FederSegmentSearchResult) Reset()         { *m = FederSegmentSearchResult{} }
func (m *FederSegmentSearchResult) String() string { return proto.CompactTextString(m) }
func (*FederSegmentSearchResult) ProtoMessage()    {}
func (*FederSegmentSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{1}
}

func (m *FederSegmentSearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederSegmentSearchResult.Unmarshal(m, b)
}
func (m *FederSegmentSearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederSegmentSearchResult.Marshal(b, m, deterministic)
}
func (m *FederSegmentSearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederSegmentSearchResult.Merge(m, src)
}
func (m *FederSegmentSearchResult) XXX_Size() int {
	return xxx_messageInfo_FederSegmentSearchResult.Size(m)
}
func (m *FederSegmentSearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FederSegmentSearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_FederSegmentSearchResult proto.InternalMessageInfo

func (m *FederSegmentSearchResult) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *FederSegmentSearchResult) GetSegmentID() int64 {
	if m != nil {
		return m.SegmentID
	}
	return 0
}

func (m *FederSegmentSearchResult) GetVisitInfo() string {
	if m != nil {
		return m.VisitInfo
	}
	return ""
}

func (m *FederSegmentSearchResult) GetOffsets() []int64 {
	if m != nil {
		return m.Offsets
	}
	return nil
}

func (m *FederSegmentSearchResult) GetPrimaryKeys() *schemapb.IDs {
	if m != nil {
		return m.PrimaryKeys
	}
	return nil
}

type FederSearchResult struct {
	Status *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	//segmentID -> FederSegmentSearchResult
	SearchViews          map[int64]*FederSegmentSearchResult `protobuf:"bytes,2,rep,name=search_views,json=searchViews,proto3" json:"search_views,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *FederSearchResult) Reset()         { *m = FederSearchResult{} }
func (m *FederSearchResult) String() string { return proto.CompactTextString(m) }
func (*FederSearchResult) ProtoMessage()    {}
func (*FederSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{2}
}

func (m *FederSearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederSearchResult.Unmarshal(m, b)
}
func (m *FederSearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederSearchResult.Marshal(b, m, deterministic)
}
func (m *FederSearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederSearchResult.Merge(m, src)
}
func (m *FederSearchResult) XXX_Size() int {
	return xxx_messageInfo_FederSearchResult.Size(m)
}
func (m *FederSearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FederSearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_FederSearchResult proto.InternalMessageInfo

func (m *FederSearchResult) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *FederSearchResult) GetSearchViews() map[int64]*FederSegmentSearchResult {
	if m != nil {
		return m.SearchViews
	}
	return nil
}

type SegmentIndexData struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	SegmentID            int64            `protobuf:"varint,2,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	Rows                 int64            `protobuf:"varint,3,opt,name=rows,proto3" json:"rows,omitempty"`
	CreateTime           uint64           `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	IndexData            string           `protobuf:"bytes,5,opt,name=index_data,json=indexData,proto3" json:"index_data,omitempty"`
	Offsets              []int64          `protobuf:"varint,6,rep,packed,name=offsets,proto3" json:"offsets,omitempty"`
	PrimaryKeys          *schemapb.IDs    `protobuf:"bytes,7,opt,name=primary_keys,json=primaryKeys,proto3" json:"primary_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SegmentIndexData) Reset()         { *m = SegmentIndexData{} }
func (m *SegmentIndexData) String() string { return proto.CompactTextString(m) }
func (*SegmentIndexData) ProtoMessage()    {}
func (*SegmentIndexData) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{3}
}

func (m *SegmentIndexData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentIndexData.Unmarshal(m, b)
}
func (m *SegmentIndexData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentIndexData.Marshal(b, m, deterministic)
}
func (m *SegmentIndexData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentIndexData.Merge(m, src)
}
func (m *SegmentIndexData) XXX_Size() int {
	return xxx_messageInfo_SegmentIndexData.Size(m)
}
func (m *SegmentIndexData) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentIndexData.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentIndexData proto.InternalMessageInfo

func (m *SegmentIndexData) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SegmentIndexData) GetSegmentID() int64 {
	if m != nil {
		return m.SegmentID
	}
	return 0
}

func (m *SegmentIndexData) GetRows() int64 {
	if m != nil {
		return m.Rows
	}
	return 0
}

func (m *SegmentIndexData) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *SegmentIndexData) GetIndexData() string {
	if m != nil {
		return m.IndexData
	}
	return ""
}

func (m *SegmentIndexData) GetOffsets() []int64 {
	if m != nil {
		return m.Offsets
	}
	return nil
}

func (m *SegmentIndexData) GetPrimaryKeys() *schemapb.IDs {
	if m != nil {
		return m.PrimaryKeys
	}
	return nil
}

type DescribeSegmentIndexDataRequest struct {
	CollectionName       string   `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	SegmentsIDs          []int64  `protobuf:"varint,2,rep,packed,name=segmentsIDs,proto3" json:"segmentsIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeSegmentIndexDataRequest) Reset()         { *m = DescribeSegmentIndexDataRequest{} }
func (m *DescribeSegmentIndexDataRequest) String() string { return proto.CompactTextString(m) }
func (*DescribeSegmentIndexDataRequest) ProtoMessage()    {}
func (*DescribeSegmentIndexDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{4}
}

func (m *DescribeSegmentIndexDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeSegmentIndexDataRequest.Unmarshal(m, b)
}
func (m *DescribeSegmentIndexDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeSegmentIndexDataRequest.Marshal(b, m, deterministic)
}
func (m *DescribeSegmentIndexDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeSegmentIndexDataRequest.Merge(m, src)
}
func (m *DescribeSegmentIndexDataRequest) XXX_Size() int {
	return xxx_messageInfo_DescribeSegmentIndexDataRequest.Size(m)
}
func (m *DescribeSegmentIndexDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeSegmentIndexDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeSegmentIndexDataRequest proto.InternalMessageInfo

func (m *DescribeSegmentIndexDataRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *DescribeSegmentIndexDataRequest) GetSegmentsIDs() []int64 {
	if m != nil {
		return m.SegmentsIDs
	}
	return nil
}

type DescribeSegmentIndexDataResponse struct {
	Status               *commonpb.Status    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IndexDatas           []*SegmentIndexData `protobuf:"bytes,2,rep,name=index_datas,json=indexDatas,proto3" json:"index_datas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DescribeSegmentIndexDataResponse) Reset()         { *m = DescribeSegmentIndexDataResponse{} }
func (m *DescribeSegmentIndexDataResponse) String() string { return proto.CompactTextString(m) }
func (*DescribeSegmentIndexDataResponse) ProtoMessage()    {}
func (*DescribeSegmentIndexDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{5}
}

func (m *DescribeSegmentIndexDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeSegmentIndexDataResponse.Unmarshal(m, b)
}
func (m *DescribeSegmentIndexDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeSegmentIndexDataResponse.Marshal(b, m, deterministic)
}
func (m *DescribeSegmentIndexDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeSegmentIndexDataResponse.Merge(m, src)
}
func (m *DescribeSegmentIndexDataResponse) XXX_Size() int {
	return xxx_messageInfo_DescribeSegmentIndexDataResponse.Size(m)
}
func (m *DescribeSegmentIndexDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeSegmentIndexDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeSegmentIndexDataResponse proto.InternalMessageInfo

func (m *DescribeSegmentIndexDataResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeSegmentIndexDataResponse) GetIndexDatas() []*SegmentIndexData {
	if m != nil {
		return m.IndexDatas
	}
	return nil
}

func init() {
	proto.RegisterType((*FederSearchRequest)(nil), "milvus.proto.feder.FederSearchRequest")
	proto.RegisterType((*FederSegmentSearchResult)(nil), "milvus.proto.feder.FederSegmentSearchResult")
	proto.RegisterType((*FederSearchResult)(nil), "milvus.proto.feder.FederSearchResult")
	proto.RegisterMapType((map[int64]*FederSegmentSearchResult)(nil), "milvus.proto.feder.FederSearchResult.SearchViewsEntry")
	proto.RegisterType((*SegmentIndexData)(nil), "milvus.proto.feder.SegmentIndexData")
	proto.RegisterType((*DescribeSegmentIndexDataRequest)(nil), "milvus.proto.feder.DescribeSegmentIndexDataRequest")
	proto.RegisterType((*DescribeSegmentIndexDataResponse)(nil), "milvus.proto.feder.DescribeSegmentIndexDataResponse")
}

func init() { proto.RegisterFile("feder.proto", fileDescriptor_84d670fd126c7825) }

var fileDescriptor_84d670fd126c7825 = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xdf, 0x4e, 0xd4, 0x4e,
	0x18, 0xfd, 0x75, 0xbb, 0xc0, 0xaf, 0x5f, 0x57, 0xc4, 0xb9, 0x6a, 0x56, 0x09, 0xb5, 0xd1, 0xd8,
	0x0b, 0x29, 0x66, 0x31, 0xc6, 0xe8, 0x1d, 0x59, 0x48, 0x36, 0x26, 0x5e, 0x0c, 0xc4, 0x44, 0x2f,
	0x6c, 0x66, 0xcb, 0x57, 0x98, 0xd0, 0x76, 0xd6, 0xce, 0x74, 0x71, 0xdf, 0xc1, 0x0b, 0x7d, 0x00,
	0x7d, 0x39, 0x5f, 0xc4, 0x74, 0xa6, 0x05, 0x96, 0x05, 0x0c, 0x24, 0xde, 0xcd, 0x9c, 0x7e, 0x7f,
	0x4e, 0xcf, 0x39, 0x03, 0x6e, 0x8a, 0x87, 0x58, 0x46, 0x93, 0x52, 0x28, 0x41, 0x48, 0xce, 0xb3,
	0x69, 0x25, 0xcd, 0x2d, 0xd2, 0x5f, 0xfa, 0xbd, 0x44, 0xe4, 0xb9, 0x28, 0x0c, 0xd6, 0xef, 0xc9,
	0xe4, 0x18, 0x73, 0x66, 0x6e, 0xc1, 0x67, 0x20, 0x7b, 0x75, 0xd1, 0x3e, 0xb2, 0x32, 0x39, 0xa6,
	0xf8, 0xa5, 0x42, 0xa9, 0xc8, 0x53, 0x58, 0xd5, 0xad, 0xb1, 0xc4, 0xa3, 0x1c, 0x0b, 0x25, 0x3d,
	0xcb, 0xb7, 0x43, 0x9b, 0xde, 0x4b, 0x4d, 0xad, 0x01, 0xc9, 0x63, 0xe8, 0x49, 0xdd, 0x17, 0x2b,
	0xa1, 0x58, 0xe6, 0x75, 0x7c, 0x2b, 0xfc, 0x9f, 0xba, 0x06, 0x3b, 0xa8, 0xa1, 0xe0, 0xb7, 0x05,
	0xde, 0xde, 0x85, 0xa6, 0x76, 0x8f, 0xac, 0x32, 0x45, 0xb6, 0x61, 0x59, 0x2a, 0xa6, 0xaa, 0x7a,
	0xbc, 0x15, 0xba, 0x83, 0x87, 0xd1, 0x1c, 0xfb, 0x86, 0xf6, 0xbe, 0x2e, 0xa1, 0x4d, 0x29, 0x79,
	0x04, 0x4e, 0xc3, 0x6a, 0x34, 0xd4, 0x1b, 0x6d, 0x7a, 0x0e, 0x90, 0x75, 0x80, 0x29, 0x97, 0x5c,
	0xc5, 0xbc, 0x48, 0x85, 0x67, 0xfb, 0x56, 0xe8, 0x50, 0x47, 0x23, 0xa3, 0x22, 0x15, 0xc4, 0x83,
	0x15, 0x91, 0xa6, 0x12, 0x95, 0xf4, 0xba, 0xfa, 0x8f, 0xda, 0x2b, 0x79, 0x0b, 0xbd, 0x49, 0xc9,
	0x73, 0x56, 0xce, 0xe2, 0x13, 0x9c, 0x49, 0x6f, 0x49, 0x33, 0xf2, 0xe6, 0x19, 0x35, 0xd2, 0x8d,
	0x86, 0x92, 0xba, 0x4d, 0xf5, 0x3b, 0x9c, 0xc9, 0xe0, 0x47, 0x07, 0x1e, 0xcc, 0xc9, 0x78, 0xf7,
	0xdf, 0xfb, 0x78, 0xa6, 0xe9, 0x94, 0xe3, 0xa9, 0xf4, 0x3a, 0xbe, 0x1d, 0xba, 0x83, 0x57, 0xd1,
	0xa2, 0xaf, 0xd1, 0xc2, 0xc6, 0xc8, 0x5c, 0x3e, 0xd4, 0x8d, 0xbb, 0x85, 0x2a, 0x67, 0xad, 0x17,
	0x1a, 0xe9, 0x67, 0xb0, 0x76, 0xb9, 0x80, 0xac, 0x81, 0x7d, 0x82, 0x33, 0x4d, 0xd0, 0xa6, 0xf5,
	0x91, 0xec, 0xc0, 0xd2, 0x94, 0x65, 0x15, 0x6a, 0x6d, 0xdd, 0xc1, 0xf3, 0x1b, 0x36, 0x2f, 0x38,
	0x4a, 0x4d, 0xeb, 0x9b, 0xce, 0x6b, 0x2b, 0xf8, 0xde, 0xa9, 0xd7, 0x19, 0x5f, 0x8a, 0x43, 0xfc,
	0x3a, 0x64, 0x8a, 0xfd, 0x0b, 0xc7, 0x09, 0x74, 0x4b, 0x71, 0x2a, 0xb5, 0xd7, 0x36, 0xd5, 0x67,
	0xb2, 0x01, 0x6e, 0x52, 0x22, 0x53, 0x18, 0x2b, 0x9e, 0xa3, 0xd7, 0xf5, 0xad, 0xb0, 0x4b, 0xc1,
	0x40, 0x07, 0x3c, 0xc7, 0x3a, 0x26, 0xbc, 0x26, 0x15, 0x1f, 0x32, 0xc5, 0xb4, 0xd7, 0x0e, 0x75,
	0xf8, 0x19, 0xcd, 0x0b, 0x31, 0x59, 0xbe, 0x39, 0x26, 0x2b, 0xb7, 0x89, 0x49, 0x06, 0x1b, 0x43,
	0x94, 0x49, 0xc9, 0xc7, 0x78, 0x59, 0x99, 0xf6, 0xe5, 0x3d, 0x83, 0xfb, 0x89, 0xc8, 0x32, 0x4c,
	0x14, 0x17, 0x45, 0x5c, 0xb0, 0x1c, 0xb5, 0x52, 0x0e, 0x5d, 0x3d, 0x87, 0xdf, 0xb3, 0x1c, 0x89,
	0x0f, 0x6e, 0xfb, 0x38, 0x47, 0x43, 0x13, 0x13, 0x9b, 0x5e, 0x84, 0x82, 0x5f, 0x16, 0xf8, 0xd7,
	0xaf, 0x93, 0x13, 0x51, 0x48, 0xbc, 0x9b, 0x21, 0xbb, 0xe0, 0x9e, 0xab, 0xd7, 0x46, 0xf4, 0xc9,
	0x55, 0x41, 0x59, 0xd8, 0x0b, 0x67, 0x22, 0xcb, 0xc1, 0x4f, 0x0b, 0x7a, 0x4d, 0x92, 0xca, 0x29,
	0x4f, 0x90, 0x7c, 0xb3, 0x60, 0x5d, 0x03, 0xd7, 0xd1, 0x26, 0xdb, 0x57, 0x2d, 0xf9, 0x8b, 0xa6,
	0xfd, 0x97, 0xb7, 0x6b, 0x32, 0xca, 0x04, 0xff, 0xed, 0x0c, 0x3e, 0xbd, 0x38, 0xe2, 0xea, 0xb8,
	0x1a, 0xd7, 0x32, 0x6c, 0x99, 0x19, 0x9b, 0x5c, 0xb4, 0x27, 0x3d, 0x6d, 0xeb, 0x48, 0x6c, 0xb2,
	0x09, 0xdf, 0xd2, 0x43, 0x27, 0xe3, 0xf1, 0xb2, 0x46, 0xb7, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x32, 0xcf, 0x02, 0x53, 0x95, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FederServiceClient is the client API for FederService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FederServiceClient interface {
	FederDescribeSegmentIndexData(ctx context.Context, in *DescribeSegmentIndexDataRequest, opts ...grpc.CallOption) (*DescribeSegmentIndexDataResponse, error)
}

type federServiceClient struct {
	cc *grpc.ClientConn
}

func NewFederServiceClient(cc *grpc.ClientConn) FederServiceClient {
	return &federServiceClient{cc}
}

func (c *federServiceClient) FederDescribeSegmentIndexData(ctx context.Context, in *DescribeSegmentIndexDataRequest, opts ...grpc.CallOption) (*DescribeSegmentIndexDataResponse, error) {
	out := new(DescribeSegmentIndexDataResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.feder.FederService/FederDescribeSegmentIndexData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederServiceServer is the server API for FederService service.
type FederServiceServer interface {
	FederDescribeSegmentIndexData(context.Context, *DescribeSegmentIndexDataRequest) (*DescribeSegmentIndexDataResponse, error)
}

// UnimplementedFederServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFederServiceServer struct {
}

func (*UnimplementedFederServiceServer) FederDescribeSegmentIndexData(ctx context.Context, req *DescribeSegmentIndexDataRequest) (*DescribeSegmentIndexDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FederDescribeSegmentIndexData not implemented")
}

func RegisterFederServiceServer(s *grpc.Server, srv FederServiceServer) {
	s.RegisterService(&_FederService_serviceDesc, srv)
}

func _FederService_FederDescribeSegmentIndexData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSegmentIndexDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederServiceServer).FederDescribeSegmentIndexData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.feder.FederService/FederDescribeSegmentIndexData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederServiceServer).FederDescribeSegmentIndexData(ctx, req.(*DescribeSegmentIndexDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FederService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.feder.FederService",
	HandlerType: (*FederServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FederDescribeSegmentIndexData",
			Handler:    _FederService_FederDescribeSegmentIndexData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feder.proto",
}
