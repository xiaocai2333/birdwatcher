// Code generated by protoc-gen-go. DO NOT EDIT.
// source: etcd_meta.proto

package etcdpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/milvus-io/birdwatcher/proto/v2.0/commonpb"
	schemapb "github.com/milvus-io/birdwatcher/proto/v2.0/schemapb"
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

type ProxyMeta struct {
	ID                   int64             `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Address              *commonpb.Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	ResultChannelIDs     []string          `protobuf:"bytes,3,rep,name=result_channelIDs,json=resultChannelIDs,proto3" json:"result_channelIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProxyMeta) Reset()         { *m = ProxyMeta{} }
func (m *ProxyMeta) String() string { return proto.CompactTextString(m) }
func (*ProxyMeta) ProtoMessage()    {}
func (*ProxyMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{0}
}

func (m *ProxyMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyMeta.Unmarshal(m, b)
}
func (m *ProxyMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyMeta.Marshal(b, m, deterministic)
}
func (m *ProxyMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyMeta.Merge(m, src)
}
func (m *ProxyMeta) XXX_Size() int {
	return xxx_messageInfo_ProxyMeta.Size(m)
}
func (m *ProxyMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyMeta proto.InternalMessageInfo

func (m *ProxyMeta) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ProxyMeta) GetAddress() *commonpb.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ProxyMeta) GetResultChannelIDs() []string {
	if m != nil {
		return m.ResultChannelIDs
	}
	return nil
}

type IndexInfo struct {
	IndexName            string                   `protobuf:"bytes,1,opt,name=index_name,json=indexName,proto3" json:"index_name,omitempty"`
	IndexID              int64                    `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	IndexParams          []*commonpb.KeyValuePair `protobuf:"bytes,3,rep,name=index_params,json=indexParams,proto3" json:"index_params,omitempty"`
	Deleted              bool                     `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
	CreateTime           uint64                   `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IndexInfo) Reset()         { *m = IndexInfo{} }
func (m *IndexInfo) String() string { return proto.CompactTextString(m) }
func (*IndexInfo) ProtoMessage()    {}
func (*IndexInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{1}
}

func (m *IndexInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexInfo.Unmarshal(m, b)
}
func (m *IndexInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexInfo.Marshal(b, m, deterministic)
}
func (m *IndexInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexInfo.Merge(m, src)
}
func (m *IndexInfo) XXX_Size() int {
	return xxx_messageInfo_IndexInfo.Size(m)
}
func (m *IndexInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IndexInfo proto.InternalMessageInfo

func (m *IndexInfo) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *IndexInfo) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *IndexInfo) GetIndexParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.IndexParams
	}
	return nil
}

func (m *IndexInfo) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *IndexInfo) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type FieldIndexInfo struct {
	FiledID              int64    `protobuf:"varint,1,opt,name=filedID,proto3" json:"filedID,omitempty"`
	IndexID              int64    `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldIndexInfo) Reset()         { *m = FieldIndexInfo{} }
func (m *FieldIndexInfo) String() string { return proto.CompactTextString(m) }
func (*FieldIndexInfo) ProtoMessage()    {}
func (*FieldIndexInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{2}
}

func (m *FieldIndexInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldIndexInfo.Unmarshal(m, b)
}
func (m *FieldIndexInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldIndexInfo.Marshal(b, m, deterministic)
}
func (m *FieldIndexInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldIndexInfo.Merge(m, src)
}
func (m *FieldIndexInfo) XXX_Size() int {
	return xxx_messageInfo_FieldIndexInfo.Size(m)
}
func (m *FieldIndexInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldIndexInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FieldIndexInfo proto.InternalMessageInfo

func (m *FieldIndexInfo) GetFiledID() int64 {
	if m != nil {
		return m.FiledID
	}
	return 0
}

func (m *FieldIndexInfo) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

type CollectionInfo struct {
	ID                         int64                      `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Schema                     *schemapb.CollectionSchema `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	CreateTime                 uint64                     `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	PartitionIDs               []int64                    `protobuf:"varint,4,rep,packed,name=partitionIDs,proto3" json:"partitionIDs,omitempty"`
	PartitionNames             []string                   `protobuf:"bytes,5,rep,name=partitionNames,proto3" json:"partitionNames,omitempty"`
	FieldIndexes               []*FieldIndexInfo          `protobuf:"bytes,6,rep,name=field_indexes,json=fieldIndexes,proto3" json:"field_indexes,omitempty"`
	VirtualChannelNames        []string                   `protobuf:"bytes,7,rep,name=virtual_channel_names,json=virtualChannelNames,proto3" json:"virtual_channel_names,omitempty"`
	PhysicalChannelNames       []string                   `protobuf:"bytes,8,rep,name=physical_channel_names,json=physicalChannelNames,proto3" json:"physical_channel_names,omitempty"`
	PartitionCreatedTimestamps []uint64                   `protobuf:"varint,9,rep,packed,name=partition_created_timestamps,json=partitionCreatedTimestamps,proto3" json:"partition_created_timestamps,omitempty"`
	ShardsNum                  int32                      `protobuf:"varint,10,opt,name=shards_num,json=shardsNum,proto3" json:"shards_num,omitempty"`
	StartPositions             []*commonpb.KeyDataPair    `protobuf:"bytes,11,rep,name=start_positions,json=startPositions,proto3" json:"start_positions,omitempty"`
	ConsistencyLevel           commonpb.ConsistencyLevel  `protobuf:"varint,12,opt,name=consistency_level,json=consistencyLevel,proto3,enum=milvus.proto.common.ConsistencyLevel" json:"consistency_level,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                   `json:"-"`
	XXX_unrecognized           []byte                     `json:"-"`
	XXX_sizecache              int32                      `json:"-"`
}

func (m *CollectionInfo) Reset()         { *m = CollectionInfo{} }
func (m *CollectionInfo) String() string { return proto.CompactTextString(m) }
func (*CollectionInfo) ProtoMessage()    {}
func (*CollectionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{3}
}

func (m *CollectionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionInfo.Unmarshal(m, b)
}
func (m *CollectionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionInfo.Marshal(b, m, deterministic)
}
func (m *CollectionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionInfo.Merge(m, src)
}
func (m *CollectionInfo) XXX_Size() int {
	return xxx_messageInfo_CollectionInfo.Size(m)
}
func (m *CollectionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionInfo proto.InternalMessageInfo

func (m *CollectionInfo) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *CollectionInfo) GetSchema() *schemapb.CollectionSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CollectionInfo) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *CollectionInfo) GetPartitionIDs() []int64 {
	if m != nil {
		return m.PartitionIDs
	}
	return nil
}

func (m *CollectionInfo) GetPartitionNames() []string {
	if m != nil {
		return m.PartitionNames
	}
	return nil
}

func (m *CollectionInfo) GetFieldIndexes() []*FieldIndexInfo {
	if m != nil {
		return m.FieldIndexes
	}
	return nil
}

func (m *CollectionInfo) GetVirtualChannelNames() []string {
	if m != nil {
		return m.VirtualChannelNames
	}
	return nil
}

func (m *CollectionInfo) GetPhysicalChannelNames() []string {
	if m != nil {
		return m.PhysicalChannelNames
	}
	return nil
}

func (m *CollectionInfo) GetPartitionCreatedTimestamps() []uint64 {
	if m != nil {
		return m.PartitionCreatedTimestamps
	}
	return nil
}

func (m *CollectionInfo) GetShardsNum() int32 {
	if m != nil {
		return m.ShardsNum
	}
	return 0
}

func (m *CollectionInfo) GetStartPositions() []*commonpb.KeyDataPair {
	if m != nil {
		return m.StartPositions
	}
	return nil
}

func (m *CollectionInfo) GetConsistencyLevel() commonpb.ConsistencyLevel {
	if m != nil {
		return m.ConsistencyLevel
	}
	return commonpb.ConsistencyLevel_Strong
}

type SegmentIndexInfo struct {
	CollectionID         int64    `protobuf:"varint,1,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	PartitionID          int64    `protobuf:"varint,2,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	SegmentID            int64    `protobuf:"varint,3,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	FieldID              int64    `protobuf:"varint,4,opt,name=fieldID,proto3" json:"fieldID,omitempty"`
	IndexID              int64    `protobuf:"varint,5,opt,name=indexID,proto3" json:"indexID,omitempty"`
	BuildID              int64    `protobuf:"varint,6,opt,name=buildID,proto3" json:"buildID,omitempty"`
	EnableIndex          bool     `protobuf:"varint,7,opt,name=enable_index,json=enableIndex,proto3" json:"enable_index,omitempty"`
	CreateTime           uint64   `protobuf:"varint,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentIndexInfo) Reset()         { *m = SegmentIndexInfo{} }
func (m *SegmentIndexInfo) String() string { return proto.CompactTextString(m) }
func (*SegmentIndexInfo) ProtoMessage()    {}
func (*SegmentIndexInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{4}
}

func (m *SegmentIndexInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentIndexInfo.Unmarshal(m, b)
}
func (m *SegmentIndexInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentIndexInfo.Marshal(b, m, deterministic)
}
func (m *SegmentIndexInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentIndexInfo.Merge(m, src)
}
func (m *SegmentIndexInfo) XXX_Size() int {
	return xxx_messageInfo_SegmentIndexInfo.Size(m)
}
func (m *SegmentIndexInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentIndexInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentIndexInfo proto.InternalMessageInfo

func (m *SegmentIndexInfo) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

func (m *SegmentIndexInfo) GetPartitionID() int64 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

func (m *SegmentIndexInfo) GetSegmentID() int64 {
	if m != nil {
		return m.SegmentID
	}
	return 0
}

func (m *SegmentIndexInfo) GetFieldID() int64 {
	if m != nil {
		return m.FieldID
	}
	return 0
}

func (m *SegmentIndexInfo) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *SegmentIndexInfo) GetBuildID() int64 {
	if m != nil {
		return m.BuildID
	}
	return 0
}

func (m *SegmentIndexInfo) GetEnableIndex() bool {
	if m != nil {
		return m.EnableIndex
	}
	return false
}

func (m *SegmentIndexInfo) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type CollectionMeta struct {
	ID                   int64                      `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Schema               *schemapb.CollectionSchema `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	CreateTime           uint64                     `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	SegmentIDs           []int64                    `protobuf:"varint,4,rep,packed,name=segmentIDs,proto3" json:"segmentIDs,omitempty"`
	PartitionTags        []string                   `protobuf:"bytes,5,rep,name=partition_tags,json=partitionTags,proto3" json:"partition_tags,omitempty"`
	PartitionIDs         []int64                    `protobuf:"varint,6,rep,packed,name=partitionIDs,proto3" json:"partitionIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CollectionMeta) Reset()         { *m = CollectionMeta{} }
func (m *CollectionMeta) String() string { return proto.CompactTextString(m) }
func (*CollectionMeta) ProtoMessage()    {}
func (*CollectionMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{5}
}

func (m *CollectionMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionMeta.Unmarshal(m, b)
}
func (m *CollectionMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionMeta.Marshal(b, m, deterministic)
}
func (m *CollectionMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionMeta.Merge(m, src)
}
func (m *CollectionMeta) XXX_Size() int {
	return xxx_messageInfo_CollectionMeta.Size(m)
}
func (m *CollectionMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionMeta.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionMeta proto.InternalMessageInfo

func (m *CollectionMeta) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *CollectionMeta) GetSchema() *schemapb.CollectionSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CollectionMeta) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *CollectionMeta) GetSegmentIDs() []int64 {
	if m != nil {
		return m.SegmentIDs
	}
	return nil
}

func (m *CollectionMeta) GetPartitionTags() []string {
	if m != nil {
		return m.PartitionTags
	}
	return nil
}

func (m *CollectionMeta) GetPartitionIDs() []int64 {
	if m != nil {
		return m.PartitionIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*ProxyMeta)(nil), "milvus.proto.etcd.ProxyMeta")
	proto.RegisterType((*IndexInfo)(nil), "milvus.proto.etcd.IndexInfo")
	proto.RegisterType((*FieldIndexInfo)(nil), "milvus.proto.etcd.FieldIndexInfo")
	proto.RegisterType((*CollectionInfo)(nil), "milvus.proto.etcd.CollectionInfo")
	proto.RegisterType((*SegmentIndexInfo)(nil), "milvus.proto.etcd.SegmentIndexInfo")
	proto.RegisterType((*CollectionMeta)(nil), "milvus.proto.etcd.CollectionMeta")
}

func init() { proto.RegisterFile("etcd_meta.proto", fileDescriptor_975d306d62b73e88) }

var fileDescriptor_975d306d62b73e88 = []byte{
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4b, 0x6b, 0xeb, 0x46,
	0x14, 0x46, 0x96, 0x1f, 0xd1, 0xb1, 0xae, 0x6f, 0x32, 0x7d, 0x30, 0x84, 0xb4, 0xd5, 0x35, 0xdc,
	0x22, 0x28, 0xb5, 0x69, 0x6e, 0xe9, 0xae, 0xd0, 0xd6, 0x22, 0x60, 0xda, 0x06, 0x33, 0x09, 0x5d,
	0x74, 0x23, 0xc6, 0xd2, 0xb1, 0x3d, 0xa0, 0x87, 0xd1, 0x8c, 0x42, 0xbc, 0xeb, 0xb2, 0x7f, 0xac,
	0xbf, 0xa6, 0xff, 0xa1, 0x14, 0xcd, 0x48, 0xf2, 0x2b, 0x59, 0xde, 0x9d, 0xbf, 0xef, 0x3c, 0x34,
	0xe7, 0x9c, 0xef, 0x33, 0xbc, 0x45, 0x15, 0xc5, 0x61, 0x8a, 0x8a, 0x4f, 0xb6, 0x45, 0xae, 0x72,
	0x72, 0x95, 0x8a, 0xe4, 0xa9, 0x94, 0x06, 0x4d, 0xaa, 0xe8, 0xb5, 0x1b, 0xe5, 0x69, 0x9a, 0x67,
	0x86, 0xba, 0x76, 0x65, 0xb4, 0xc1, 0xb4, 0x4e, 0x1f, 0xff, 0x65, 0x81, 0xb3, 0x28, 0xf2, 0xe7,
	0xdd, 0xef, 0xa8, 0x38, 0x19, 0x41, 0x67, 0x1e, 0x50, 0xcb, 0xb3, 0x7c, 0x9b, 0x75, 0xe6, 0x01,
	0xf9, 0x01, 0x06, 0x3c, 0x8e, 0x0b, 0x94, 0x92, 0x76, 0x3c, 0xcb, 0x1f, 0xde, 0xde, 0x4c, 0x8e,
	0xda, 0xd7, 0x8d, 0x7f, 0x36, 0x39, 0xac, 0x49, 0x26, 0xdf, 0xc0, 0x55, 0x81, 0xb2, 0x4c, 0x54,
	0x18, 0x6d, 0x78, 0x96, 0x61, 0x32, 0x0f, 0x24, 0xb5, 0x3d, 0xdb, 0x77, 0xd8, 0xa5, 0x09, 0xcc,
	0x5a, 0x7e, 0xfc, 0x8f, 0x05, 0xce, 0x3c, 0x8b, 0xf1, 0x79, 0x9e, 0xad, 0x72, 0xf2, 0x05, 0x80,
	0xa8, 0x40, 0x98, 0xf1, 0x14, 0xf5, 0x53, 0x1c, 0xe6, 0x68, 0xe6, 0x9e, 0xa7, 0x48, 0x28, 0x0c,
	0x34, 0x98, 0x07, 0xfa, 0x45, 0x36, 0x6b, 0x20, 0x09, 0xc0, 0x35, 0x85, 0x5b, 0x5e, 0xf0, 0xd4,
	0x7c, 0x6e, 0x78, 0xfb, 0xee, 0xc5, 0x07, 0xff, 0x8a, 0xbb, 0x3f, 0x78, 0x52, 0xe2, 0x82, 0x8b,
	0x82, 0x0d, 0x75, 0xd9, 0x42, 0x57, 0x55, 0xfd, 0x63, 0x4c, 0x50, 0x61, 0x4c, 0xbb, 0x9e, 0xe5,
	0x5f, 0xb0, 0x06, 0x92, 0xaf, 0x60, 0x18, 0x15, 0xc8, 0x15, 0x86, 0x4a, 0xa4, 0x48, 0x7b, 0x9e,
	0xe5, 0x77, 0x19, 0x18, 0xea, 0x51, 0xa4, 0x38, 0x0e, 0x60, 0x74, 0x27, 0x30, 0x89, 0xf7, 0xb3,
	0x50, 0x18, 0xac, 0x44, 0x82, 0x71, 0xbb, 0xd3, 0x06, 0xbe, 0x3e, 0xc6, 0xf8, 0xbf, 0x2e, 0x8c,
	0x66, 0x79, 0x92, 0x60, 0xa4, 0x44, 0x9e, 0xe9, 0x36, 0xa7, 0x57, 0xf9, 0x11, 0xfa, 0xe6, 0x86,
	0xf5, 0x51, 0xde, 0x1f, 0xcf, 0x58, 0xdf, 0x77, 0xdf, 0xe4, 0x41, 0x13, 0xac, 0x2e, 0x3a, 0x1d,
	0xc4, 0x3e, 0x1d, 0x84, 0x8c, 0xc1, 0xdd, 0xf2, 0x42, 0x09, 0xfd, 0x80, 0x40, 0xd2, 0xae, 0x67,
	0xfb, 0x36, 0x3b, 0xe2, 0xc8, 0xd7, 0x30, 0x6a, 0x71, 0x75, 0x18, 0x49, 0x7b, 0xfa, 0xbc, 0x27,
	0x2c, 0xb9, 0x83, 0x37, 0xab, 0x6a, 0x29, 0xa1, 0x9e, 0x0f, 0x25, 0xed, 0xbf, 0x74, 0x96, 0x4a,
	0xa6, 0x93, 0xe3, 0xe5, 0x31, 0x77, 0xd5, 0x62, 0x94, 0xe4, 0x16, 0x3e, 0x7b, 0x12, 0x85, 0x2a,
	0x79, 0xd2, 0x48, 0x4a, 0x0b, 0x44, 0xd2, 0x81, 0xfe, 0xec, 0x27, 0x75, 0xb0, 0x96, 0x95, 0xf9,
	0xf6, 0xf7, 0xf0, 0xf9, 0x76, 0xb3, 0x93, 0x22, 0x3a, 0x2b, 0xba, 0xd0, 0x45, 0x9f, 0x36, 0xd1,
	0xa3, 0xaa, 0x9f, 0xe0, 0xa6, 0x9d, 0x21, 0x34, 0x5b, 0x89, 0xf5, 0xa6, 0xa4, 0xe2, 0xe9, 0x56,
	0x52, 0xc7, 0xb3, 0xfd, 0x2e, 0xbb, 0x6e, 0x73, 0x66, 0x26, 0xe5, 0xb1, 0xcd, 0xa8, 0x24, 0x2c,
	0x37, 0xbc, 0x88, 0x65, 0x98, 0x95, 0x29, 0x05, 0xcf, 0xf2, 0x7b, 0xcc, 0x31, 0xcc, 0x7d, 0x99,
	0x92, 0x39, 0xbc, 0x95, 0x8a, 0x17, 0x2a, 0xdc, 0xe6, 0x52, 0x77, 0x90, 0x74, 0xa8, 0x97, 0xe2,
	0xbd, 0xa6, 0xd5, 0x80, 0x2b, 0xae, 0xa5, 0x3a, 0xd2, 0x85, 0x8b, 0xa6, 0x8e, 0x30, 0xb8, 0x8a,
	0xf2, 0x4c, 0x0a, 0xa9, 0x30, 0x8b, 0x76, 0x61, 0x82, 0x4f, 0x98, 0x50, 0xd7, 0xb3, 0xfc, 0xd1,
	0xa9, 0x28, 0xea, 0x66, 0xb3, 0x7d, 0xf6, 0x6f, 0x55, 0x32, 0xbb, 0x8c, 0x4e, 0x98, 0xf1, 0xdf,
	0x1d, 0xb8, 0x7c, 0xc0, 0x75, 0x8a, 0x99, 0xda, 0x2b, 0x79, 0x0c, 0x6e, 0xb4, 0x17, 0x65, 0x23,
	0xc6, 0x23, 0x8e, 0x78, 0x30, 0x3c, 0x90, 0x48, 0xad, 0xeb, 0x43, 0x8a, 0xdc, 0x80, 0x23, 0xeb,
	0xce, 0x81, 0xd6, 0x9d, 0xcd, 0xf6, 0x84, 0x71, 0x4b, 0x75, 0xf2, 0x40, 0x5b, 0x4f, 0xbb, 0x45,
	0xc3, 0x43, 0xb7, 0xf4, 0x8e, 0x4d, 0x4f, 0x61, 0xb0, 0x2c, 0x85, 0xae, 0xe9, 0x9b, 0x48, 0x0d,
	0xc9, 0x3b, 0x70, 0x31, 0xe3, 0xcb, 0x04, 0x8d, 0xf2, 0xe8, 0x40, 0xbb, 0x79, 0x68, 0x38, 0x3d,
	0xd8, 0xa9, 0x11, 0x2e, 0xce, 0x1c, 0xfd, 0xaf, 0x75, 0xe8, 0xc5, 0x17, 0xff, 0x21, 0x3f, 0xb6,
	0x17, 0xbf, 0x04, 0x68, 0x37, 0xd4, 0x38, 0xf1, 0x80, 0x21, 0xef, 0x0f, 0x7c, 0x18, 0x2a, 0xbe,
	0x6e, 0x7c, 0xf8, 0xa6, 0x65, 0x1f, 0xf9, 0x5a, 0x9e, 0x59, 0xba, 0x7f, 0x6e, 0xe9, 0x5f, 0x3e,
	0xfc, 0xf9, 0xdd, 0x5a, 0xa8, 0x4d, 0xb9, 0xac, 0xc4, 0x32, 0x35, 0x63, 0x7c, 0x2b, 0xf2, 0xfa,
	0xd7, 0x54, 0x64, 0x0a, 0x8b, 0x8c, 0x27, 0x53, 0x3d, 0xd9, 0xb4, 0xb2, 0xec, 0x76, 0xb9, 0xec,
	0x6b, 0xf4, 0xe1, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x2d, 0xaa, 0x13, 0x88, 0x06, 0x00,
	0x00,
}
