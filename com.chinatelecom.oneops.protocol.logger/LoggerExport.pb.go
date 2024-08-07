// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.0
// source: LoggerExport.proto

package logger

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

// 日志数据，一次发送多个表数据
type LoggerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 采集时间
	ReceiveTime int64 `protobuf:"varint,1,opt,name=receiveTime,proto3" json:"receiveTime,omitempty"`
	// 监控对象的唯一标志
	MonitorObject string `protobuf:"bytes,2,opt,name=monitorObject,proto3" json:"monitorObject,omitempty"`
	// 监控对象的类型
	MonitorType string `protobuf:"bytes,3,opt,name=monitorType,proto3" json:"monitorType,omitempty"`
	// 收集器的唯一标志
	Collector string       `protobuf:"bytes,4,opt,name=collector,proto3" json:"collector,omitempty"`
	TableData []*TableData `protobuf:"bytes,5,rep,name=tableData,proto3" json:"tableData,omitempty"`
}

func (x *LoggerData) Reset() {
	*x = LoggerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LoggerExport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggerData) ProtoMessage() {}

func (x *LoggerData) ProtoReflect() protoreflect.Message {
	mi := &file_LoggerExport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggerData.ProtoReflect.Descriptor instead.
func (*LoggerData) Descriptor() ([]byte, []int) {
	return file_LoggerExport_proto_rawDescGZIP(), []int{0}
}

func (x *LoggerData) GetReceiveTime() int64 {
	if x != nil {
		return x.ReceiveTime
	}
	return 0
}

func (x *LoggerData) GetMonitorObject() string {
	if x != nil {
		return x.MonitorObject
	}
	return ""
}

func (x *LoggerData) GetMonitorType() string {
	if x != nil {
		return x.MonitorType
	}
	return ""
}

func (x *LoggerData) GetCollector() string {
	if x != nil {
		return x.Collector
	}
	return ""
}

func (x *LoggerData) GetTableData() []*TableData {
	if x != nil {
		return x.TableData
	}
	return nil
}

// 表数据，一个可以插入多条记录
type TableData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 表名称
	TableName string `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	// 列名称
	Columns []*ColumnData `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	// 字段值
	Rows []*RowValue `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *TableData) Reset() {
	*x = TableData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LoggerExport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableData) ProtoMessage() {}

func (x *TableData) ProtoReflect() protoreflect.Message {
	mi := &file_LoggerExport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableData.ProtoReflect.Descriptor instead.
func (*TableData) Descriptor() ([]byte, []int) {
	return file_LoggerExport_proto_rawDescGZIP(), []int{1}
}

func (x *TableData) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *TableData) GetColumns() []*ColumnData {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *TableData) GetRows() []*RowValue {
	if x != nil {
		return x.Rows
	}
	return nil
}

// 列属性
type ColumnData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 字段名称
	ColumnName string `protobuf:"bytes,1,opt,name=columnName,proto3" json:"columnName,omitempty"`
	// 字段类型
	ColumnType int32 `protobuf:"varint,2,opt,name=columnType,proto3" json:"columnType,omitempty"`
}

func (x *ColumnData) Reset() {
	*x = ColumnData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LoggerExport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnData) ProtoMessage() {}

func (x *ColumnData) ProtoReflect() protoreflect.Message {
	mi := &file_LoggerExport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnData.ProtoReflect.Descriptor instead.
func (*ColumnData) Descriptor() ([]byte, []int) {
	return file_LoggerExport_proto_rawDescGZIP(), []int{2}
}

func (x *ColumnData) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

func (x *ColumnData) GetColumnType() int32 {
	if x != nil {
		return x.ColumnType
	}
	return 0
}

// 对应每列的数据
type RowValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldValue []*FieldValue `protobuf:"bytes,1,rep,name=fieldValue,proto3" json:"fieldValue,omitempty"`
}

func (x *RowValue) Reset() {
	*x = RowValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LoggerExport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RowValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RowValue) ProtoMessage() {}

func (x *RowValue) ProtoReflect() protoreflect.Message {
	mi := &file_LoggerExport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RowValue.ProtoReflect.Descriptor instead.
func (*RowValue) Descriptor() ([]byte, []int) {
	return file_LoggerExport_proto_rawDescGZIP(), []int{3}
}

func (x *RowValue) GetFieldValue() []*FieldValue {
	if x != nil {
		return x.FieldValue
	}
	return nil
}

// 每列的具体值
type FieldValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*FieldValue_I
	//	*FieldValue_L
	//	*FieldValue_D
	//	*FieldValue_B
	//	*FieldValue_S
	Data isFieldValue_Data `protobuf_oneof:"data"`
}

func (x *FieldValue) Reset() {
	*x = FieldValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LoggerExport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldValue) ProtoMessage() {}

func (x *FieldValue) ProtoReflect() protoreflect.Message {
	mi := &file_LoggerExport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldValue.ProtoReflect.Descriptor instead.
func (*FieldValue) Descriptor() ([]byte, []int) {
	return file_LoggerExport_proto_rawDescGZIP(), []int{4}
}

func (m *FieldValue) GetData() isFieldValue_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *FieldValue) GetI() int32 {
	if x, ok := x.GetData().(*FieldValue_I); ok {
		return x.I
	}
	return 0
}

func (x *FieldValue) GetL() int64 {
	if x, ok := x.GetData().(*FieldValue_L); ok {
		return x.L
	}
	return 0
}

func (x *FieldValue) GetD() float64 {
	if x, ok := x.GetData().(*FieldValue_D); ok {
		return x.D
	}
	return 0
}

func (x *FieldValue) GetB() bool {
	if x, ok := x.GetData().(*FieldValue_B); ok {
		return x.B
	}
	return false
}

func (x *FieldValue) GetS() string {
	if x, ok := x.GetData().(*FieldValue_S); ok {
		return x.S
	}
	return ""
}

type isFieldValue_Data interface {
	isFieldValue_Data()
}

type FieldValue_I struct {
	I int32 `protobuf:"varint,1,opt,name=i,proto3,oneof"`
}

type FieldValue_L struct {
	L int64 `protobuf:"varint,2,opt,name=l,proto3,oneof"`
}

type FieldValue_D struct {
	D float64 `protobuf:"fixed64,3,opt,name=d,proto3,oneof"`
}

type FieldValue_B struct {
	B bool `protobuf:"varint,4,opt,name=b,proto3,oneof"`
}

type FieldValue_S struct {
	S string `protobuf:"bytes,5,opt,name=s,proto3,oneof"`
}

func (*FieldValue_I) isFieldValue_Data() {}

func (*FieldValue_L) isFieldValue_Data() {}

func (*FieldValue_D) isFieldValue_Data() {}

func (*FieldValue_B) isFieldValue_Data() {}

func (*FieldValue_S) isFieldValue_Data() {}

var File_LoggerExport_proto protoreflect.FileDescriptor

var file_LoggerExport_proto_rawDesc = []byte{
	0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xc3, 0x01, 0x0a, 0x0a, 0x4c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x2d, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x79, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x6f, 0x77,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x4c, 0x0a, 0x0a, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3c, 0x0a, 0x08, 0x52, 0x6f, 0x77,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x64, 0x0a, 0x0a, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x01, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x01, 0x69, 0x12, 0x0e, 0x0a, 0x01, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x01, 0x6c, 0x12, 0x0e, 0x0a, 0x01, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x48, 0x00, 0x52, 0x01, 0x64, 0x12, 0x0e, 0x0a, 0x01, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x01, 0x62, 0x12, 0x0e, 0x0a, 0x01, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x01, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x54, 0x0a,
	0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x6c, 0x65, 0x63, 0x6f,
	0x6d, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x27, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x6c, 0x65, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x65,
	0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LoggerExport_proto_rawDescOnce sync.Once
	file_LoggerExport_proto_rawDescData = file_LoggerExport_proto_rawDesc
)

func file_LoggerExport_proto_rawDescGZIP() []byte {
	file_LoggerExport_proto_rawDescOnce.Do(func() {
		file_LoggerExport_proto_rawDescData = protoimpl.X.CompressGZIP(file_LoggerExport_proto_rawDescData)
	})
	return file_LoggerExport_proto_rawDescData
}

var file_LoggerExport_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_LoggerExport_proto_goTypes = []interface{}{
	(*LoggerData)(nil), // 0: main.LoggerData
	(*TableData)(nil),  // 1: main.TableData
	(*ColumnData)(nil), // 2: main.ColumnData
	(*RowValue)(nil),   // 3: main.RowValue
	(*FieldValue)(nil), // 4: main.FieldValue
}
var file_LoggerExport_proto_depIdxs = []int32{
	1, // 0: main.LoggerData.tableData:type_name -> main.TableData
	2, // 1: main.TableData.columns:type_name -> main.ColumnData
	3, // 2: main.TableData.rows:type_name -> main.RowValue
	4, // 3: main.RowValue.fieldValue:type_name -> main.FieldValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_LoggerExport_proto_init() }
func file_LoggerExport_proto_init() {
	if File_LoggerExport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LoggerExport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggerData); i {
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
		file_LoggerExport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableData); i {
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
		file_LoggerExport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnData); i {
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
		file_LoggerExport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RowValue); i {
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
		file_LoggerExport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldValue); i {
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
	file_LoggerExport_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*FieldValue_I)(nil),
		(*FieldValue_L)(nil),
		(*FieldValue_D)(nil),
		(*FieldValue_B)(nil),
		(*FieldValue_S)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_LoggerExport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LoggerExport_proto_goTypes,
		DependencyIndexes: file_LoggerExport_proto_depIdxs,
		MessageInfos:      file_LoggerExport_proto_msgTypes,
	}.Build()
	File_LoggerExport_proto = out.File
	file_LoggerExport_proto_rawDesc = nil
	file_LoggerExport_proto_goTypes = nil
	file_LoggerExport_proto_depIdxs = nil
}
