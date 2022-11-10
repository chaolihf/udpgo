// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.0
// source: PrometheusExport.proto

package prometheus

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

type MetricData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//指标名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//指标值
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	//属性参数
	Labels []*LabelNameData `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *MetricData) Reset() {
	*x = MetricData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PrometheusExport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricData) ProtoMessage() {}

func (x *MetricData) ProtoReflect() protoreflect.Message {
	mi := &file_PrometheusExport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricData.ProtoReflect.Descriptor instead.
func (*MetricData) Descriptor() ([]byte, []int) {
	return file_PrometheusExport_proto_rawDescGZIP(), []int{0}
}

func (x *MetricData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *MetricData) GetLabels() []*LabelNameData {
	if x != nil {
		return x.Labels
	}
	return nil
}

type LabelNameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//属性名称
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	//属性值
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LabelNameData) Reset() {
	*x = LabelNameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PrometheusExport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelNameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelNameData) ProtoMessage() {}

func (x *LabelNameData) ProtoReflect() protoreflect.Message {
	mi := &file_PrometheusExport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelNameData.ProtoReflect.Descriptor instead.
func (*LabelNameData) Descriptor() ([]byte, []int) {
	return file_PrometheusExport_proto_rawDescGZIP(), []int{1}
}

func (x *LabelNameData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LabelNameData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MetricsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//采集时间
	ReceiveTime int64 `protobuf:"varint,1,opt,name=receiveTime,proto3" json:"receiveTime,omitempty"`
	//监控对象的唯一标志
	MonitorObject string `protobuf:"bytes,2,opt,name=monitorObject,proto3" json:"monitorObject,omitempty"`
	//监控对象的类型
	MonitorType string `protobuf:"bytes,3,opt,name=monitorType,proto3" json:"monitorType,omitempty"`
	//收集器的唯一标志
	Collector string        `protobuf:"bytes,4,opt,name=collector,proto3" json:"collector,omitempty"`
	Datas     []*MetricData `protobuf:"bytes,5,rep,name=datas,proto3" json:"datas,omitempty"`
}

func (x *MetricsData) Reset() {
	*x = MetricsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PrometheusExport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsData) ProtoMessage() {}

func (x *MetricsData) ProtoReflect() protoreflect.Message {
	mi := &file_PrometheusExport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsData.ProtoReflect.Descriptor instead.
func (*MetricsData) Descriptor() ([]byte, []int) {
	return file_PrometheusExport_proto_rawDescGZIP(), []int{2}
}

func (x *MetricsData) GetReceiveTime() int64 {
	if x != nil {
		return x.ReceiveTime
	}
	return 0
}

func (x *MetricsData) GetMonitorObject() string {
	if x != nil {
		return x.MonitorObject
	}
	return ""
}

func (x *MetricsData) GetMonitorType() string {
	if x != nil {
		return x.MonitorType
	}
	return ""
}

func (x *MetricsData) GetCollector() string {
	if x != nil {
		return x.Collector
	}
	return ""
}

func (x *MetricsData) GetDatas() []*MetricData {
	if x != nil {
		return x.Datas
	}
	return nil
}

var File_PrometheusExport_proto protoreflect.FileDescriptor

var file_PrometheusExport_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x63,
	0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x22, 0x35, 0x0a, 0x0d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x0b, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x42, 0x5c, 0x0a, 0x2b, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x68, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x6c, 0x65, 0x63, 0x6f, 0x6d, 0x2e, 0x6f,
	0x6e, 0x65, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x50, 0x01, 0x5a, 0x2b, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x6c, 0x65, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e,
	0x65, 0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PrometheusExport_proto_rawDescOnce sync.Once
	file_PrometheusExport_proto_rawDescData = file_PrometheusExport_proto_rawDesc
)

func file_PrometheusExport_proto_rawDescGZIP() []byte {
	file_PrometheusExport_proto_rawDescOnce.Do(func() {
		file_PrometheusExport_proto_rawDescData = protoimpl.X.CompressGZIP(file_PrometheusExport_proto_rawDescData)
	})
	return file_PrometheusExport_proto_rawDescData
}

var file_PrometheusExport_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_PrometheusExport_proto_goTypes = []interface{}{
	(*MetricData)(nil),    // 0: main.MetricData
	(*LabelNameData)(nil), // 1: main.LabelNameData
	(*MetricsData)(nil),   // 2: main.MetricsData
}
var file_PrometheusExport_proto_depIdxs = []int32{
	1, // 0: main.MetricData.labels:type_name -> main.LabelNameData
	0, // 1: main.MetricsData.datas:type_name -> main.MetricData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PrometheusExport_proto_init() }
func file_PrometheusExport_proto_init() {
	if File_PrometheusExport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PrometheusExport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricData); i {
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
		file_PrometheusExport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelNameData); i {
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
		file_PrometheusExport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsData); i {
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
			RawDescriptor: file_PrometheusExport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PrometheusExport_proto_goTypes,
		DependencyIndexes: file_PrometheusExport_proto_depIdxs,
		MessageInfos:      file_PrometheusExport_proto_msgTypes,
	}.Build()
	File_PrometheusExport_proto = out.File
	file_PrometheusExport_proto_rawDesc = nil
	file_PrometheusExport_proto_goTypes = nil
	file_PrometheusExport_proto_depIdxs = nil
}