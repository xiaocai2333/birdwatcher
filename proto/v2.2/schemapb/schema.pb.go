// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package schemapb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/milvus-io/birdwatcher/proto/v2.2/commonpb"
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

//*
// @brief Field data type
type DataType int32

const (
	DataType_None         DataType = 0
	DataType_Bool         DataType = 1
	DataType_Int8         DataType = 2
	DataType_Int16        DataType = 3
	DataType_Int32        DataType = 4
	DataType_Int64        DataType = 5
	DataType_Float        DataType = 10
	DataType_Double       DataType = 11
	DataType_String       DataType = 20
	DataType_VarChar      DataType = 21
	DataType_BinaryVector DataType = 100
	DataType_FloatVector  DataType = 101
)

var DataType_name = map[int32]string{
	0:   "None",
	1:   "Bool",
	2:   "Int8",
	3:   "Int16",
	4:   "Int32",
	5:   "Int64",
	10:  "Float",
	11:  "Double",
	20:  "String",
	21:  "VarChar",
	100: "BinaryVector",
	101: "FloatVector",
}

var DataType_value = map[string]int32{
	"None":         0,
	"Bool":         1,
	"Int8":         2,
	"Int16":        3,
	"Int32":        4,
	"Int64":        5,
	"Float":        10,
	"Double":       11,
	"String":       20,
	"VarChar":      21,
	"BinaryVector": 100,
	"FloatVector":  101,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}

func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

type FieldState int32

const (
	FieldState_FieldCreated  FieldState = 0
	FieldState_FieldCreating FieldState = 1
	FieldState_FieldDropping FieldState = 2
	FieldState_FieldDropped  FieldState = 3
)

var FieldState_name = map[int32]string{
	0: "FieldCreated",
	1: "FieldCreating",
	2: "FieldDropping",
	3: "FieldDropped",
}

var FieldState_value = map[string]int32{
	"FieldCreated":  0,
	"FieldCreating": 1,
	"FieldDropping": 2,
	"FieldDropped":  3,
}

func (x FieldState) String() string {
	return proto.EnumName(FieldState_name, int32(x))
}

func (FieldState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}

//*
// @brief Field schema
type FieldSchema struct {
	FieldID              int64                    `protobuf:"varint,1,opt,name=fieldID,proto3" json:"fieldID,omitempty"`
	Name                 string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsPrimaryKey         bool                     `protobuf:"varint,3,opt,name=is_primary_key,json=isPrimaryKey,proto3" json:"is_primary_key,omitempty"`
	Description          string                   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DataType             DataType                 `protobuf:"varint,5,opt,name=data_type,json=dataType,proto3,enum=milvus.protov2.schema.DataType" json:"data_type,omitempty"`
	TypeParams           []*commonpb.KeyValuePair `protobuf:"bytes,6,rep,name=type_params,json=typeParams,proto3" json:"type_params,omitempty"`
	IndexParams          []*commonpb.KeyValuePair `protobuf:"bytes,7,rep,name=index_params,json=indexParams,proto3" json:"index_params,omitempty"`
	AutoID               bool                     `protobuf:"varint,8,opt,name=autoID,proto3" json:"autoID,omitempty"`
	State                FieldState               `protobuf:"varint,9,opt,name=state,proto3,enum=milvus.protov2.schema.FieldState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FieldSchema) Reset()         { *m = FieldSchema{} }
func (m *FieldSchema) String() string { return proto.CompactTextString(m) }
func (*FieldSchema) ProtoMessage()    {}
func (*FieldSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *FieldSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldSchema.Unmarshal(m, b)
}
func (m *FieldSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldSchema.Marshal(b, m, deterministic)
}
func (m *FieldSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldSchema.Merge(m, src)
}
func (m *FieldSchema) XXX_Size() int {
	return xxx_messageInfo_FieldSchema.Size(m)
}
func (m *FieldSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldSchema.DiscardUnknown(m)
}

var xxx_messageInfo_FieldSchema proto.InternalMessageInfo

func (m *FieldSchema) GetFieldID() int64 {
	if m != nil {
		return m.FieldID
	}
	return 0
}

func (m *FieldSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldSchema) GetIsPrimaryKey() bool {
	if m != nil {
		return m.IsPrimaryKey
	}
	return false
}

func (m *FieldSchema) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FieldSchema) GetDataType() DataType {
	if m != nil {
		return m.DataType
	}
	return DataType_None
}

func (m *FieldSchema) GetTypeParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.TypeParams
	}
	return nil
}

func (m *FieldSchema) GetIndexParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.IndexParams
	}
	return nil
}

func (m *FieldSchema) GetAutoID() bool {
	if m != nil {
		return m.AutoID
	}
	return false
}

func (m *FieldSchema) GetState() FieldState {
	if m != nil {
		return m.State
	}
	return FieldState_FieldCreated
}

//*
// @brief Collection schema
type CollectionSchema struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	AutoID               bool           `protobuf:"varint,3,opt,name=autoID,proto3" json:"autoID,omitempty"`
	Fields               []*FieldSchema `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CollectionSchema) Reset()         { *m = CollectionSchema{} }
func (m *CollectionSchema) String() string { return proto.CompactTextString(m) }
func (*CollectionSchema) ProtoMessage()    {}
func (*CollectionSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}

func (m *CollectionSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionSchema.Unmarshal(m, b)
}
func (m *CollectionSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionSchema.Marshal(b, m, deterministic)
}
func (m *CollectionSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionSchema.Merge(m, src)
}
func (m *CollectionSchema) XXX_Size() int {
	return xxx_messageInfo_CollectionSchema.Size(m)
}
func (m *CollectionSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionSchema.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionSchema proto.InternalMessageInfo

func (m *CollectionSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CollectionSchema) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CollectionSchema) GetAutoID() bool {
	if m != nil {
		return m.AutoID
	}
	return false
}

func (m *CollectionSchema) GetFields() []*FieldSchema {
	if m != nil {
		return m.Fields
	}
	return nil
}

type BoolArray struct {
	Data                 []bool   `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolArray) Reset()         { *m = BoolArray{} }
func (m *BoolArray) String() string { return proto.CompactTextString(m) }
func (*BoolArray) ProtoMessage()    {}
func (*BoolArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{2}
}

func (m *BoolArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolArray.Unmarshal(m, b)
}
func (m *BoolArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolArray.Marshal(b, m, deterministic)
}
func (m *BoolArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolArray.Merge(m, src)
}
func (m *BoolArray) XXX_Size() int {
	return xxx_messageInfo_BoolArray.Size(m)
}
func (m *BoolArray) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolArray.DiscardUnknown(m)
}

var xxx_messageInfo_BoolArray proto.InternalMessageInfo

func (m *BoolArray) GetData() []bool {
	if m != nil {
		return m.Data
	}
	return nil
}

type IntArray struct {
	Data                 []int32  `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntArray) Reset()         { *m = IntArray{} }
func (m *IntArray) String() string { return proto.CompactTextString(m) }
func (*IntArray) ProtoMessage()    {}
func (*IntArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{3}
}

func (m *IntArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntArray.Unmarshal(m, b)
}
func (m *IntArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntArray.Marshal(b, m, deterministic)
}
func (m *IntArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntArray.Merge(m, src)
}
func (m *IntArray) XXX_Size() int {
	return xxx_messageInfo_IntArray.Size(m)
}
func (m *IntArray) XXX_DiscardUnknown() {
	xxx_messageInfo_IntArray.DiscardUnknown(m)
}

var xxx_messageInfo_IntArray proto.InternalMessageInfo

func (m *IntArray) GetData() []int32 {
	if m != nil {
		return m.Data
	}
	return nil
}

type LongArray struct {
	Data                 []int64  `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongArray) Reset()         { *m = LongArray{} }
func (m *LongArray) String() string { return proto.CompactTextString(m) }
func (*LongArray) ProtoMessage()    {}
func (*LongArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{4}
}

func (m *LongArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongArray.Unmarshal(m, b)
}
func (m *LongArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongArray.Marshal(b, m, deterministic)
}
func (m *LongArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongArray.Merge(m, src)
}
func (m *LongArray) XXX_Size() int {
	return xxx_messageInfo_LongArray.Size(m)
}
func (m *LongArray) XXX_DiscardUnknown() {
	xxx_messageInfo_LongArray.DiscardUnknown(m)
}

var xxx_messageInfo_LongArray proto.InternalMessageInfo

func (m *LongArray) GetData() []int64 {
	if m != nil {
		return m.Data
	}
	return nil
}

type FloatArray struct {
	Data                 []float32 `protobuf:"fixed32,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FloatArray) Reset()         { *m = FloatArray{} }
func (m *FloatArray) String() string { return proto.CompactTextString(m) }
func (*FloatArray) ProtoMessage()    {}
func (*FloatArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{5}
}

func (m *FloatArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatArray.Unmarshal(m, b)
}
func (m *FloatArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatArray.Marshal(b, m, deterministic)
}
func (m *FloatArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatArray.Merge(m, src)
}
func (m *FloatArray) XXX_Size() int {
	return xxx_messageInfo_FloatArray.Size(m)
}
func (m *FloatArray) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatArray.DiscardUnknown(m)
}

var xxx_messageInfo_FloatArray proto.InternalMessageInfo

func (m *FloatArray) GetData() []float32 {
	if m != nil {
		return m.Data
	}
	return nil
}

type DoubleArray struct {
	Data                 []float64 `protobuf:"fixed64,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DoubleArray) Reset()         { *m = DoubleArray{} }
func (m *DoubleArray) String() string { return proto.CompactTextString(m) }
func (*DoubleArray) ProtoMessage()    {}
func (*DoubleArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{6}
}

func (m *DoubleArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleArray.Unmarshal(m, b)
}
func (m *DoubleArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleArray.Marshal(b, m, deterministic)
}
func (m *DoubleArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleArray.Merge(m, src)
}
func (m *DoubleArray) XXX_Size() int {
	return xxx_messageInfo_DoubleArray.Size(m)
}
func (m *DoubleArray) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleArray.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleArray proto.InternalMessageInfo

func (m *DoubleArray) GetData() []float64 {
	if m != nil {
		return m.Data
	}
	return nil
}

// For special fields such as bigdecimal, array...
type BytesArray struct {
	Data                 [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesArray) Reset()         { *m = BytesArray{} }
func (m *BytesArray) String() string { return proto.CompactTextString(m) }
func (*BytesArray) ProtoMessage()    {}
func (*BytesArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{7}
}

func (m *BytesArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesArray.Unmarshal(m, b)
}
func (m *BytesArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesArray.Marshal(b, m, deterministic)
}
func (m *BytesArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesArray.Merge(m, src)
}
func (m *BytesArray) XXX_Size() int {
	return xxx_messageInfo_BytesArray.Size(m)
}
func (m *BytesArray) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesArray.DiscardUnknown(m)
}

var xxx_messageInfo_BytesArray proto.InternalMessageInfo

func (m *BytesArray) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type StringArray struct {
	Data                 []string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringArray) Reset()         { *m = StringArray{} }
func (m *StringArray) String() string { return proto.CompactTextString(m) }
func (*StringArray) ProtoMessage()    {}
func (*StringArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{8}
}

func (m *StringArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringArray.Unmarshal(m, b)
}
func (m *StringArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringArray.Marshal(b, m, deterministic)
}
func (m *StringArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringArray.Merge(m, src)
}
func (m *StringArray) XXX_Size() int {
	return xxx_messageInfo_StringArray.Size(m)
}
func (m *StringArray) XXX_DiscardUnknown() {
	xxx_messageInfo_StringArray.DiscardUnknown(m)
}

var xxx_messageInfo_StringArray proto.InternalMessageInfo

func (m *StringArray) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

type ScalarField struct {
	// Types that are valid to be assigned to Data:
	//	*ScalarField_BoolData
	//	*ScalarField_IntData
	//	*ScalarField_LongData
	//	*ScalarField_FloatData
	//	*ScalarField_DoubleData
	//	*ScalarField_StringData
	//	*ScalarField_BytesData
	Data                 isScalarField_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ScalarField) Reset()         { *m = ScalarField{} }
func (m *ScalarField) String() string { return proto.CompactTextString(m) }
func (*ScalarField) ProtoMessage()    {}
func (*ScalarField) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{9}
}

func (m *ScalarField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScalarField.Unmarshal(m, b)
}
func (m *ScalarField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScalarField.Marshal(b, m, deterministic)
}
func (m *ScalarField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScalarField.Merge(m, src)
}
func (m *ScalarField) XXX_Size() int {
	return xxx_messageInfo_ScalarField.Size(m)
}
func (m *ScalarField) XXX_DiscardUnknown() {
	xxx_messageInfo_ScalarField.DiscardUnknown(m)
}

var xxx_messageInfo_ScalarField proto.InternalMessageInfo

type isScalarField_Data interface {
	isScalarField_Data()
}

type ScalarField_BoolData struct {
	BoolData *BoolArray `protobuf:"bytes,1,opt,name=bool_data,json=boolData,proto3,oneof"`
}

type ScalarField_IntData struct {
	IntData *IntArray `protobuf:"bytes,2,opt,name=int_data,json=intData,proto3,oneof"`
}

type ScalarField_LongData struct {
	LongData *LongArray `protobuf:"bytes,3,opt,name=long_data,json=longData,proto3,oneof"`
}

type ScalarField_FloatData struct {
	FloatData *FloatArray `protobuf:"bytes,4,opt,name=float_data,json=floatData,proto3,oneof"`
}

type ScalarField_DoubleData struct {
	DoubleData *DoubleArray `protobuf:"bytes,5,opt,name=double_data,json=doubleData,proto3,oneof"`
}

type ScalarField_StringData struct {
	StringData *StringArray `protobuf:"bytes,6,opt,name=string_data,json=stringData,proto3,oneof"`
}

type ScalarField_BytesData struct {
	BytesData *BytesArray `protobuf:"bytes,7,opt,name=bytes_data,json=bytesData,proto3,oneof"`
}

func (*ScalarField_BoolData) isScalarField_Data() {}

func (*ScalarField_IntData) isScalarField_Data() {}

func (*ScalarField_LongData) isScalarField_Data() {}

func (*ScalarField_FloatData) isScalarField_Data() {}

func (*ScalarField_DoubleData) isScalarField_Data() {}

func (*ScalarField_StringData) isScalarField_Data() {}

func (*ScalarField_BytesData) isScalarField_Data() {}

func (m *ScalarField) GetData() isScalarField_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ScalarField) GetBoolData() *BoolArray {
	if x, ok := m.GetData().(*ScalarField_BoolData); ok {
		return x.BoolData
	}
	return nil
}

func (m *ScalarField) GetIntData() *IntArray {
	if x, ok := m.GetData().(*ScalarField_IntData); ok {
		return x.IntData
	}
	return nil
}

func (m *ScalarField) GetLongData() *LongArray {
	if x, ok := m.GetData().(*ScalarField_LongData); ok {
		return x.LongData
	}
	return nil
}

func (m *ScalarField) GetFloatData() *FloatArray {
	if x, ok := m.GetData().(*ScalarField_FloatData); ok {
		return x.FloatData
	}
	return nil
}

func (m *ScalarField) GetDoubleData() *DoubleArray {
	if x, ok := m.GetData().(*ScalarField_DoubleData); ok {
		return x.DoubleData
	}
	return nil
}

func (m *ScalarField) GetStringData() *StringArray {
	if x, ok := m.GetData().(*ScalarField_StringData); ok {
		return x.StringData
	}
	return nil
}

func (m *ScalarField) GetBytesData() *BytesArray {
	if x, ok := m.GetData().(*ScalarField_BytesData); ok {
		return x.BytesData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ScalarField) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ScalarField_BoolData)(nil),
		(*ScalarField_IntData)(nil),
		(*ScalarField_LongData)(nil),
		(*ScalarField_FloatData)(nil),
		(*ScalarField_DoubleData)(nil),
		(*ScalarField_StringData)(nil),
		(*ScalarField_BytesData)(nil),
	}
}

type VectorField struct {
	Dim int64 `protobuf:"varint,1,opt,name=dim,proto3" json:"dim,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*VectorField_FloatVector
	//	*VectorField_BinaryVector
	Data                 isVectorField_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *VectorField) Reset()         { *m = VectorField{} }
func (m *VectorField) String() string { return proto.CompactTextString(m) }
func (*VectorField) ProtoMessage()    {}
func (*VectorField) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{10}
}

func (m *VectorField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VectorField.Unmarshal(m, b)
}
func (m *VectorField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VectorField.Marshal(b, m, deterministic)
}
func (m *VectorField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VectorField.Merge(m, src)
}
func (m *VectorField) XXX_Size() int {
	return xxx_messageInfo_VectorField.Size(m)
}
func (m *VectorField) XXX_DiscardUnknown() {
	xxx_messageInfo_VectorField.DiscardUnknown(m)
}

var xxx_messageInfo_VectorField proto.InternalMessageInfo

func (m *VectorField) GetDim() int64 {
	if m != nil {
		return m.Dim
	}
	return 0
}

type isVectorField_Data interface {
	isVectorField_Data()
}

type VectorField_FloatVector struct {
	FloatVector *FloatArray `protobuf:"bytes,2,opt,name=float_vector,json=floatVector,proto3,oneof"`
}

type VectorField_BinaryVector struct {
	BinaryVector []byte `protobuf:"bytes,3,opt,name=binary_vector,json=binaryVector,proto3,oneof"`
}

func (*VectorField_FloatVector) isVectorField_Data() {}

func (*VectorField_BinaryVector) isVectorField_Data() {}

func (m *VectorField) GetData() isVectorField_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *VectorField) GetFloatVector() *FloatArray {
	if x, ok := m.GetData().(*VectorField_FloatVector); ok {
		return x.FloatVector
	}
	return nil
}

func (m *VectorField) GetBinaryVector() []byte {
	if x, ok := m.GetData().(*VectorField_BinaryVector); ok {
		return x.BinaryVector
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*VectorField) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*VectorField_FloatVector)(nil),
		(*VectorField_BinaryVector)(nil),
	}
}

type FieldData struct {
	Type      DataType `protobuf:"varint,1,opt,name=type,proto3,enum=milvus.protov2.schema.DataType" json:"type,omitempty"`
	FieldName string   `protobuf:"bytes,2,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	// Types that are valid to be assigned to Field:
	//	*FieldData_Scalars
	//	*FieldData_Vectors
	Field                isFieldData_Field `protobuf_oneof:"field"`
	FieldId              int64             `protobuf:"varint,5,opt,name=field_id,json=fieldId,proto3" json:"field_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FieldData) Reset()         { *m = FieldData{} }
func (m *FieldData) String() string { return proto.CompactTextString(m) }
func (*FieldData) ProtoMessage()    {}
func (*FieldData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{11}
}

func (m *FieldData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldData.Unmarshal(m, b)
}
func (m *FieldData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldData.Marshal(b, m, deterministic)
}
func (m *FieldData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldData.Merge(m, src)
}
func (m *FieldData) XXX_Size() int {
	return xxx_messageInfo_FieldData.Size(m)
}
func (m *FieldData) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldData.DiscardUnknown(m)
}

var xxx_messageInfo_FieldData proto.InternalMessageInfo

func (m *FieldData) GetType() DataType {
	if m != nil {
		return m.Type
	}
	return DataType_None
}

func (m *FieldData) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

type isFieldData_Field interface {
	isFieldData_Field()
}

type FieldData_Scalars struct {
	Scalars *ScalarField `protobuf:"bytes,3,opt,name=scalars,proto3,oneof"`
}

type FieldData_Vectors struct {
	Vectors *VectorField `protobuf:"bytes,4,opt,name=vectors,proto3,oneof"`
}

func (*FieldData_Scalars) isFieldData_Field() {}

func (*FieldData_Vectors) isFieldData_Field() {}

func (m *FieldData) GetField() isFieldData_Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *FieldData) GetScalars() *ScalarField {
	if x, ok := m.GetField().(*FieldData_Scalars); ok {
		return x.Scalars
	}
	return nil
}

func (m *FieldData) GetVectors() *VectorField {
	if x, ok := m.GetField().(*FieldData_Vectors); ok {
		return x.Vectors
	}
	return nil
}

func (m *FieldData) GetFieldId() int64 {
	if m != nil {
		return m.FieldId
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FieldData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FieldData_Scalars)(nil),
		(*FieldData_Vectors)(nil),
	}
}

type IDs struct {
	// Types that are valid to be assigned to IdField:
	//	*IDs_IntId
	//	*IDs_StrId
	IdField              isIDs_IdField `protobuf_oneof:"id_field"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IDs) Reset()         { *m = IDs{} }
func (m *IDs) String() string { return proto.CompactTextString(m) }
func (*IDs) ProtoMessage()    {}
func (*IDs) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{12}
}

func (m *IDs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDs.Unmarshal(m, b)
}
func (m *IDs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDs.Marshal(b, m, deterministic)
}
func (m *IDs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDs.Merge(m, src)
}
func (m *IDs) XXX_Size() int {
	return xxx_messageInfo_IDs.Size(m)
}
func (m *IDs) XXX_DiscardUnknown() {
	xxx_messageInfo_IDs.DiscardUnknown(m)
}

var xxx_messageInfo_IDs proto.InternalMessageInfo

type isIDs_IdField interface {
	isIDs_IdField()
}

type IDs_IntId struct {
	IntId *LongArray `protobuf:"bytes,1,opt,name=int_id,json=intId,proto3,oneof"`
}

type IDs_StrId struct {
	StrId *StringArray `protobuf:"bytes,2,opt,name=str_id,json=strId,proto3,oneof"`
}

func (*IDs_IntId) isIDs_IdField() {}

func (*IDs_StrId) isIDs_IdField() {}

func (m *IDs) GetIdField() isIDs_IdField {
	if m != nil {
		return m.IdField
	}
	return nil
}

func (m *IDs) GetIntId() *LongArray {
	if x, ok := m.GetIdField().(*IDs_IntId); ok {
		return x.IntId
	}
	return nil
}

func (m *IDs) GetStrId() *StringArray {
	if x, ok := m.GetIdField().(*IDs_StrId); ok {
		return x.StrId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*IDs) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*IDs_IntId)(nil),
		(*IDs_StrId)(nil),
	}
}

type SearchResultData struct {
	NumQueries           int64        `protobuf:"varint,1,opt,name=num_queries,json=numQueries,proto3" json:"num_queries,omitempty"`
	TopK                 int64        `protobuf:"varint,2,opt,name=top_k,json=topK,proto3" json:"top_k,omitempty"`
	FieldsData           []*FieldData `protobuf:"bytes,3,rep,name=fields_data,json=fieldsData,proto3" json:"fields_data,omitempty"`
	Scores               []float32    `protobuf:"fixed32,4,rep,packed,name=scores,proto3" json:"scores,omitempty"`
	Ids                  *IDs         `protobuf:"bytes,5,opt,name=ids,proto3" json:"ids,omitempty"`
	Topks                []int64      `protobuf:"varint,6,rep,packed,name=topks,proto3" json:"topks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SearchResultData) Reset()         { *m = SearchResultData{} }
func (m *SearchResultData) String() string { return proto.CompactTextString(m) }
func (*SearchResultData) ProtoMessage()    {}
func (*SearchResultData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{13}
}

func (m *SearchResultData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResultData.Unmarshal(m, b)
}
func (m *SearchResultData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResultData.Marshal(b, m, deterministic)
}
func (m *SearchResultData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResultData.Merge(m, src)
}
func (m *SearchResultData) XXX_Size() int {
	return xxx_messageInfo_SearchResultData.Size(m)
}
func (m *SearchResultData) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResultData.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResultData proto.InternalMessageInfo

func (m *SearchResultData) GetNumQueries() int64 {
	if m != nil {
		return m.NumQueries
	}
	return 0
}

func (m *SearchResultData) GetTopK() int64 {
	if m != nil {
		return m.TopK
	}
	return 0
}

func (m *SearchResultData) GetFieldsData() []*FieldData {
	if m != nil {
		return m.FieldsData
	}
	return nil
}

func (m *SearchResultData) GetScores() []float32 {
	if m != nil {
		return m.Scores
	}
	return nil
}

func (m *SearchResultData) GetIds() *IDs {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *SearchResultData) GetTopks() []int64 {
	if m != nil {
		return m.Topks
	}
	return nil
}

func init() {
	proto.RegisterEnum("milvus.protov2.schema.DataType", DataType_name, DataType_value)
	proto.RegisterEnum("milvus.protov2.schema.FieldState", FieldState_name, FieldState_value)
	proto.RegisterType((*FieldSchema)(nil), "milvus.protov2.schema.FieldSchema")
	proto.RegisterType((*CollectionSchema)(nil), "milvus.protov2.schema.CollectionSchema")
	proto.RegisterType((*BoolArray)(nil), "milvus.protov2.schema.BoolArray")
	proto.RegisterType((*IntArray)(nil), "milvus.protov2.schema.IntArray")
	proto.RegisterType((*LongArray)(nil), "milvus.protov2.schema.LongArray")
	proto.RegisterType((*FloatArray)(nil), "milvus.protov2.schema.FloatArray")
	proto.RegisterType((*DoubleArray)(nil), "milvus.protov2.schema.DoubleArray")
	proto.RegisterType((*BytesArray)(nil), "milvus.protov2.schema.BytesArray")
	proto.RegisterType((*StringArray)(nil), "milvus.protov2.schema.StringArray")
	proto.RegisterType((*ScalarField)(nil), "milvus.protov2.schema.ScalarField")
	proto.RegisterType((*VectorField)(nil), "milvus.protov2.schema.VectorField")
	proto.RegisterType((*FieldData)(nil), "milvus.protov2.schema.FieldData")
	proto.RegisterType((*IDs)(nil), "milvus.protov2.schema.IDs")
	proto.RegisterType((*SearchResultData)(nil), "milvus.protov2.schema.SearchResultData")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 1053 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x5d, 0x6f, 0x1b, 0x45,
	0x17, 0xf6, 0x7a, 0xfd, 0xb1, 0x7b, 0xd6, 0xcd, 0xbb, 0xef, 0xd0, 0xa2, 0xa5, 0x12, 0xc4, 0x35,
	0x20, 0x59, 0x15, 0x4d, 0xd4, 0x04, 0x95, 0xaf, 0x8a, 0xaa, 0x8e, 0x89, 0x62, 0x05, 0x55, 0x61,
	0x03, 0xb9, 0xe0, 0xc6, 0x1a, 0x7b, 0xa7, 0xc9, 0x28, 0xeb, 0x9d, 0x65, 0x66, 0x1c, 0xe1, 0x6b,
	0xae, 0xf8, 0x03, 0x5c, 0xf5, 0xa2, 0xff, 0x8e, 0x0b, 0xfe, 0x04, 0x9a, 0x33, 0xb3, 0xb6, 0x69,
	0xe2, 0x10, 0xee, 0xce, 0xcc, 0x9c, 0xe7, 0xf1, 0x39, 0xcf, 0xf9, 0x58, 0x43, 0x47, 0x4d, 0x2f,
	0xd8, 0x8c, 0xee, 0x94, 0x52, 0x68, 0x41, 0x1e, 0xcc, 0x78, 0x7e, 0x35, 0x57, 0xf6, 0x74, 0xb5,
	0xb7, 0x63, 0x1f, 0x1f, 0x76, 0xa6, 0x62, 0x36, 0x13, 0x85, 0xbd, 0xee, 0xbd, 0xf5, 0x21, 0x3a,
	0xe4, 0x2c, 0xcf, 0x4e, 0xf1, 0x95, 0x24, 0xd0, 0x7e, 0x6d, 0x8e, 0xa3, 0x61, 0xe2, 0x75, 0xbd,
	0xbe, 0x9f, 0x56, 0x47, 0x42, 0xa0, 0x51, 0xd0, 0x19, 0x4b, 0xea, 0x5d, 0xaf, 0x1f, 0xa6, 0x68,
	0x93, 0x4f, 0x60, 0x8b, 0xab, 0x71, 0x29, 0xf9, 0x8c, 0xca, 0xc5, 0xf8, 0x92, 0x2d, 0x12, 0xbf,
	0xeb, 0xf5, 0x83, 0xb4, 0xc3, 0xd5, 0x89, 0xbd, 0x3c, 0x66, 0x0b, 0xd2, 0x85, 0x28, 0x63, 0x6a,
	0x2a, 0x79, 0xa9, 0xb9, 0x28, 0x92, 0x06, 0x12, 0xac, 0x5f, 0x91, 0xe7, 0x10, 0x66, 0x54, 0xd3,
	0xb1, 0x5e, 0x94, 0x2c, 0x69, 0x76, 0xbd, 0xfe, 0xd6, 0xde, 0xf6, 0xce, 0x8d, 0xe1, 0xef, 0x0c,
	0xa9, 0xa6, 0x3f, 0x2e, 0x4a, 0x96, 0x06, 0x99, 0xb3, 0xc8, 0x10, 0x22, 0x03, 0x1c, 0x97, 0x54,
	0xd2, 0x99, 0x4a, 0x5a, 0x5d, 0xbf, 0x1f, 0xed, 0x7d, 0xfc, 0x2e, 0xde, 0xa5, 0x7d, 0xcc, 0x16,
	0x67, 0x34, 0x9f, 0xb3, 0x13, 0xca, 0x65, 0x0a, 0x06, 0x77, 0x82, 0x30, 0x72, 0x08, 0x1d, 0x5e,
	0x64, 0xec, 0xd7, 0x8a, 0xa6, 0x7d, 0x77, 0x9a, 0x08, 0x81, 0x8e, 0xe7, 0x7d, 0x68, 0xd1, 0xb9,
	0x16, 0xa3, 0x61, 0x12, 0xa0, 0x16, 0xee, 0x44, 0xbe, 0x80, 0xa6, 0xd2, 0x54, 0xb3, 0x24, 0xc4,
	0xfc, 0x1e, 0x6d, 0xc8, 0xcf, 0x16, 0xc3, 0x38, 0xa6, 0xd6, 0xbf, 0xf7, 0xc6, 0x83, 0xf8, 0x40,
	0xe4, 0x39, 0x9b, 0x1a, 0xad, 0x5c, 0x9d, 0xaa, 0x6a, 0x78, 0x6b, 0xd5, 0x78, 0x47, 0xe7, 0xfa,
	0x75, 0x9d, 0x57, 0xb1, 0xf9, 0xff, 0x88, 0xed, 0x6b, 0x68, 0x61, 0x99, 0x55, 0xd2, 0xc0, 0xac,
	0x7b, 0xb7, 0x06, 0x87, 0x76, 0xea, 0x10, 0xbd, 0x6d, 0x08, 0x07, 0x42, 0xe4, 0x2f, 0xa5, 0xa4,
	0x0b, 0x13, 0x96, 0x29, 0x4b, 0xe2, 0x75, 0xfd, 0x7e, 0x90, 0xa2, 0xdd, 0xfb, 0x08, 0x82, 0x51,
	0xa1, 0xaf, 0xbf, 0x37, 0xdd, 0xfb, 0x36, 0x84, 0xdf, 0x8b, 0xe2, 0xfc, 0xba, 0x83, 0xef, 0x1c,
	0xba, 0x00, 0x87, 0xb9, 0xa0, 0x37, 0x50, 0xd4, 0x9d, 0xc7, 0x23, 0x88, 0x86, 0x62, 0x3e, 0xc9,
	0xd9, 0x75, 0x17, 0x6f, 0x45, 0x32, 0x58, 0x68, 0xa6, 0xae, 0x7b, 0x74, 0x56, 0x24, 0xa7, 0x5a,
	0xf2, 0x9b, 0x22, 0x09, 0x9d, 0xcb, 0x5f, 0x3e, 0x44, 0xa7, 0x53, 0x9a, 0x53, 0x89, 0x4a, 0x90,
	0x17, 0x10, 0x4e, 0x84, 0xc8, 0xc7, 0xce, 0xd1, 0xeb, 0x47, 0x7b, 0xdd, 0x0d, 0xd2, 0x2d, 0x35,
	0x3a, 0xaa, 0xa5, 0x81, 0x01, 0x99, 0x46, 0x26, 0xcf, 0x21, 0xe0, 0x85, 0xb6, 0xf8, 0x3a, 0xe2,
	0x37, 0xf5, 0x7d, 0x25, 0xe1, 0x51, 0x2d, 0x6d, 0xf3, 0x42, 0x23, 0xfa, 0x05, 0x84, 0xb9, 0x28,
	0xce, 0x2d, 0xdc, 0xbf, 0xf5, 0xe7, 0x97, 0x0a, 0x9b, 0x9f, 0x37, 0x20, 0x24, 0x18, 0x00, 0xbc,
	0x36, 0xca, 0x5a, 0x86, 0x06, 0x32, 0x6c, 0x6c, 0xcc, 0x65, 0x09, 0x8e, 0x6a, 0x69, 0x88, 0x30,
	0xe4, 0xf8, 0x0e, 0xa2, 0x0c, 0xb5, 0xb7, 0x24, 0x4d, 0x24, 0xd9, 0xd4, 0x40, 0x6b, 0x55, 0x3a,
	0xaa, 0xa5, 0x60, 0x81, 0x15, 0x8d, 0x42, 0xf5, 0x2d, 0x4d, 0xeb, 0x56, 0x9a, 0xb5, 0x3a, 0x19,
	0x1a, 0x0b, 0xac, 0x32, 0x9a, 0x98, 0x32, 0x5b, 0x96, 0xf6, 0xad, 0x19, 0xad, 0xfa, 0xc1, 0x64,
	0x84, 0x30, 0xc3, 0x31, 0x68, 0xd9, 0xca, 0xf7, 0xfe, 0xf0, 0x20, 0x3a, 0x63, 0x53, 0x2d, 0x5c,
	0xb5, 0x63, 0xf0, 0x33, 0x3e, 0x73, 0x7b, 0xd1, 0x98, 0x66, 0x67, 0x58, 0xfd, 0xae, 0xd0, 0xcd,
	0x95, 0xf0, 0x4e, 0x0a, 0x46, 0x08, 0xb4, 0xf4, 0xe4, 0x53, 0xb8, 0x37, 0xe1, 0x85, 0xd9, 0xa1,
	0x8e, 0xc8, 0x14, 0xb3, 0x73, 0x54, 0x4b, 0x3b, 0xf6, 0xda, 0xba, 0x2d, 0x03, 0xfb, 0xad, 0x0e,
	0x21, 0x86, 0x84, 0x29, 0xef, 0x43, 0x03, 0xf7, 0xa6, 0x77, 0xb7, 0xbd, 0x89, 0xce, 0xe4, 0x43,
	0x00, 0x9c, 0xdf, 0xf1, 0xda, 0x4e, 0x0f, 0xf1, 0xe6, 0x95, 0x59, 0x25, 0xdf, 0x42, 0x5b, 0x61,
	0x9f, 0x2b, 0xd7, 0x57, 0x1b, 0x2b, 0xb1, 0x9a, 0x06, 0xd3, 0x99, 0x0e, 0x64, 0xf0, 0x36, 0x13,
	0xe5, 0xba, 0x6a, 0x13, 0x7e, 0x4d, 0x5f, 0x83, 0x77, 0x20, 0xf2, 0x01, 0x04, 0x36, 0x3c, 0x9e,
	0x61, 0x47, 0x2d, 0xbf, 0x43, 0xd9, 0xa0, 0x0d, 0x4d, 0x34, 0x7b, 0xbf, 0x7b, 0xe0, 0x8f, 0x86,
	0x8a, 0x7c, 0x05, 0x2d, 0x33, 0x43, 0x3c, 0xfb, 0x97, 0x09, 0x5c, 0x1f, 0x81, 0x26, 0x2f, 0xf4,
	0x28, 0x23, 0xdf, 0x40, 0x4b, 0x69, 0x69, 0xa0, 0xf5, 0xff, 0xd0, 0x6f, 0x4d, 0xa5, 0xe5, 0x28,
	0x1b, 0x00, 0x04, 0x3c, 0x1b, 0xdb, 0x58, 0xfe, 0xf4, 0x20, 0x3e, 0x65, 0x54, 0x4e, 0x2f, 0x52,
	0xa6, 0xe6, 0xb9, 0x9d, 0x8c, 0x6d, 0x88, 0x8a, 0xf9, 0x6c, 0xfc, 0xcb, 0x9c, 0x49, 0xce, 0x94,
	0xeb, 0x1b, 0x28, 0xe6, 0xb3, 0x1f, 0xec, 0x0d, 0x79, 0x0f, 0x9a, 0x5a, 0x94, 0xe3, 0x4b, 0xfc,
	0x75, 0x3f, 0x6d, 0x68, 0x51, 0x1e, 0x93, 0x97, 0x10, 0xd9, 0xcd, 0x5a, 0x8d, 0xb5, 0x7f, 0x4b,
	0x4e, 0xcb, 0x2e, 0x48, 0x6d, 0x39, 0xb1, 0x81, 0xcd, 0x9a, 0x57, 0x53, 0x21, 0x99, 0x5d, 0xe7,
	0xf5, 0xd4, 0x9d, 0xc8, 0x67, 0xe0, 0xf3, 0x4c, 0xb9, 0x11, 0x7d, 0xb8, 0x69, 0xd1, 0x0c, 0x55,
	0x6a, 0xdc, 0xc8, 0x7d, 0x8c, 0xee, 0xd2, 0x7e, 0x50, 0xfd, 0xd4, 0x1e, 0x1e, 0xbf, 0xf1, 0x20,
	0xa8, 0x7a, 0x89, 0x04, 0xd0, 0x78, 0x25, 0x0a, 0x16, 0xd7, 0x8c, 0x65, 0x36, 0x5c, 0xec, 0x19,
	0x6b, 0x54, 0xe8, 0x2f, 0xe3, 0x3a, 0x09, 0xa1, 0x39, 0x2a, 0xf4, 0xd3, 0x67, 0xb1, 0xef, 0xcc,
	0xfd, 0xbd, 0xb8, 0xe1, 0xcc, 0x67, 0x9f, 0xc7, 0x4d, 0x63, 0xe2, 0x4c, 0xc4, 0x40, 0x00, 0x5a,
	0x76, 0x37, 0xc4, 0x91, 0xb1, 0xad, 0xe0, 0xf1, 0x7d, 0x12, 0x41, 0xfb, 0x8c, 0xca, 0x83, 0x0b,
	0x2a, 0xe3, 0x07, 0x24, 0x86, 0xce, 0x60, 0x6d, 0x1e, 0xe2, 0x8c, 0xfc, 0x0f, 0xa2, 0xc3, 0xd5,
	0x1c, 0xc5, 0xec, 0xf1, 0x19, 0xc0, 0xea, 0x0b, 0x6a, 0x00, 0x78, 0x3a, 0x90, 0x8c, 0x6a, 0x96,
	0xc5, 0x35, 0xf2, 0x7f, 0xb8, 0xb7, 0xba, 0x31, 0x3f, 0xe1, 0x2d, 0xaf, 0x86, 0x52, 0x94, 0xa5,
	0xb9, 0xaa, 0x2f, 0x71, 0x78, 0xc5, 0xb2, 0xd8, 0x1f, 0xfc, 0x04, 0x5b, 0x5c, 0x54, 0x8a, 0x9d,
	0xcb, 0x72, 0x3a, 0x88, 0xec, 0x77, 0xf0, 0xc4, 0xa8, 0x77, 0xe2, 0xfd, 0xfc, 0xf4, 0x9c, 0xeb,
	0x8b, 0xf9, 0xc4, 0xfc, 0x3f, 0xd8, 0xb5, 0x6e, 0x4f, 0xb8, 0xa8, 0x2c, 0x94, 0x78, 0xf7, 0x5c,
	0x3c, 0xa1, 0x25, 0xdf, 0xb5, 0x3a, 0x97, 0x93, 0xb7, 0x9e, 0x37, 0x69, 0xe1, 0xcb, 0xfe, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xd1, 0x4e, 0xb7, 0xba, 0x09, 0x00, 0x00,
}
