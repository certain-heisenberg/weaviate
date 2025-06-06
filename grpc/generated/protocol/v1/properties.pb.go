// Code generated by protoc-gen-go. DO NOT EDIT.

package protocol

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Properties struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fields        map[string]*Value      `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Properties) Reset() {
	*x = Properties{}
	mi := &file_v1_properties_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Properties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Properties) ProtoMessage() {}

func (x *Properties) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Properties.ProtoReflect.Descriptor instead.
func (*Properties) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{0}
}

func (x *Properties) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Value struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Kind:
	//
	//	*Value_NumberValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_ObjectValue
	//	*Value_ListValue
	//	*Value_DateValue
	//	*Value_UuidValue
	//	*Value_IntValue
	//	*Value_GeoValue
	//	*Value_BlobValue
	//	*Value_PhoneValue
	//	*Value_NullValue
	//	*Value_TextValue
	Kind          isValue_Kind `protobuf_oneof:"kind"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Value) Reset() {
	*x = Value{}
	mi := &file_v1_properties_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetKind() isValue_Kind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *Value) GetNumberValue() float64 {
	if x != nil {
		if x, ok := x.Kind.(*Value_NumberValue); ok {
			return x.NumberValue
		}
	}
	return 0
}

// Deprecated: Marked as deprecated in v1/properties.proto.
func (x *Value) GetStringValue() string {
	if x != nil {
		if x, ok := x.Kind.(*Value_StringValue); ok {
			return x.StringValue
		}
	}
	return ""
}

func (x *Value) GetBoolValue() bool {
	if x != nil {
		if x, ok := x.Kind.(*Value_BoolValue); ok {
			return x.BoolValue
		}
	}
	return false
}

func (x *Value) GetObjectValue() *Properties {
	if x != nil {
		if x, ok := x.Kind.(*Value_ObjectValue); ok {
			return x.ObjectValue
		}
	}
	return nil
}

func (x *Value) GetListValue() *ListValue {
	if x != nil {
		if x, ok := x.Kind.(*Value_ListValue); ok {
			return x.ListValue
		}
	}
	return nil
}

func (x *Value) GetDateValue() string {
	if x != nil {
		if x, ok := x.Kind.(*Value_DateValue); ok {
			return x.DateValue
		}
	}
	return ""
}

func (x *Value) GetUuidValue() string {
	if x != nil {
		if x, ok := x.Kind.(*Value_UuidValue); ok {
			return x.UuidValue
		}
	}
	return ""
}

func (x *Value) GetIntValue() int64 {
	if x != nil {
		if x, ok := x.Kind.(*Value_IntValue); ok {
			return x.IntValue
		}
	}
	return 0
}

func (x *Value) GetGeoValue() *GeoCoordinate {
	if x != nil {
		if x, ok := x.Kind.(*Value_GeoValue); ok {
			return x.GeoValue
		}
	}
	return nil
}

func (x *Value) GetBlobValue() string {
	if x != nil {
		if x, ok := x.Kind.(*Value_BlobValue); ok {
			return x.BlobValue
		}
	}
	return ""
}

func (x *Value) GetPhoneValue() *PhoneNumber {
	if x != nil {
		if x, ok := x.Kind.(*Value_PhoneValue); ok {
			return x.PhoneValue
		}
	}
	return nil
}

func (x *Value) GetNullValue() structpb.NullValue {
	if x != nil {
		if x, ok := x.Kind.(*Value_NullValue); ok {
			return x.NullValue
		}
	}
	return structpb.NullValue(0)
}

func (x *Value) GetTextValue() string {
	if x != nil {
		if x, ok := x.Kind.(*Value_TextValue); ok {
			return x.TextValue
		}
	}
	return ""
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_NumberValue struct {
	NumberValue float64 `protobuf:"fixed64,1,opt,name=number_value,json=numberValue,proto3,oneof"`
}

type Value_StringValue struct {
	// Deprecated: Marked as deprecated in v1/properties.proto.
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,3,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_ObjectValue struct {
	ObjectValue *Properties `protobuf:"bytes,4,opt,name=object_value,json=objectValue,proto3,oneof"`
}

type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,5,opt,name=list_value,json=listValue,proto3,oneof"`
}

type Value_DateValue struct {
	DateValue string `protobuf:"bytes,6,opt,name=date_value,json=dateValue,proto3,oneof"`
}

type Value_UuidValue struct {
	UuidValue string `protobuf:"bytes,7,opt,name=uuid_value,json=uuidValue,proto3,oneof"`
}

type Value_IntValue struct {
	IntValue int64 `protobuf:"varint,8,opt,name=int_value,json=intValue,proto3,oneof"`
}

type Value_GeoValue struct {
	GeoValue *GeoCoordinate `protobuf:"bytes,9,opt,name=geo_value,json=geoValue,proto3,oneof"`
}

type Value_BlobValue struct {
	BlobValue string `protobuf:"bytes,10,opt,name=blob_value,json=blobValue,proto3,oneof"`
}

type Value_PhoneValue struct {
	PhoneValue *PhoneNumber `protobuf:"bytes,11,opt,name=phone_value,json=phoneValue,proto3,oneof"`
}

type Value_NullValue struct {
	NullValue structpb.NullValue `protobuf:"varint,12,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Value_TextValue struct {
	TextValue string `protobuf:"bytes,13,opt,name=text_value,json=textValue,proto3,oneof"`
}

func (*Value_NumberValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_ObjectValue) isValue_Kind() {}

func (*Value_ListValue) isValue_Kind() {}

func (*Value_DateValue) isValue_Kind() {}

func (*Value_UuidValue) isValue_Kind() {}

func (*Value_IntValue) isValue_Kind() {}

func (*Value_GeoValue) isValue_Kind() {}

func (*Value_BlobValue) isValue_Kind() {}

func (*Value_PhoneValue) isValue_Kind() {}

func (*Value_NullValue) isValue_Kind() {}

func (*Value_TextValue) isValue_Kind() {}

type ListValue struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in v1/properties.proto.
	Values []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// Types that are valid to be assigned to Kind:
	//
	//	*ListValue_NumberValues
	//	*ListValue_BoolValues
	//	*ListValue_ObjectValues
	//	*ListValue_DateValues
	//	*ListValue_UuidValues
	//	*ListValue_IntValues
	//	*ListValue_TextValues
	Kind          isListValue_Kind `protobuf_oneof:"kind"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListValue) Reset() {
	*x = ListValue{}
	mi := &file_v1_properties_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListValue) ProtoMessage() {}

func (x *ListValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListValue.ProtoReflect.Descriptor instead.
func (*ListValue) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{2}
}

// Deprecated: Marked as deprecated in v1/properties.proto.
func (x *ListValue) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *ListValue) GetKind() isListValue_Kind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *ListValue) GetNumberValues() *NumberValues {
	if x != nil {
		if x, ok := x.Kind.(*ListValue_NumberValues); ok {
			return x.NumberValues
		}
	}
	return nil
}

func (x *ListValue) GetBoolValues() *BoolValues {
	if x != nil {
		if x, ok := x.Kind.(*ListValue_BoolValues); ok {
			return x.BoolValues
		}
	}
	return nil
}

func (x *ListValue) GetObjectValues() *ObjectValues {
	if x != nil {
		if x, ok := x.Kind.(*ListValue_ObjectValues); ok {
			return x.ObjectValues
		}
	}
	return nil
}

func (x *ListValue) GetDateValues() *DateValues {
	if x != nil {
		if x, ok := x.Kind.(*ListValue_DateValues); ok {
			return x.DateValues
		}
	}
	return nil
}

func (x *ListValue) GetUuidValues() *UuidValues {
	if x != nil {
		if x, ok := x.Kind.(*ListValue_UuidValues); ok {
			return x.UuidValues
		}
	}
	return nil
}

func (x *ListValue) GetIntValues() *IntValues {
	if x != nil {
		if x, ok := x.Kind.(*ListValue_IntValues); ok {
			return x.IntValues
		}
	}
	return nil
}

func (x *ListValue) GetTextValues() *TextValues {
	if x != nil {
		if x, ok := x.Kind.(*ListValue_TextValues); ok {
			return x.TextValues
		}
	}
	return nil
}

type isListValue_Kind interface {
	isListValue_Kind()
}

type ListValue_NumberValues struct {
	NumberValues *NumberValues `protobuf:"bytes,2,opt,name=number_values,json=numberValues,proto3,oneof"`
}

type ListValue_BoolValues struct {
	BoolValues *BoolValues `protobuf:"bytes,3,opt,name=bool_values,json=boolValues,proto3,oneof"`
}

type ListValue_ObjectValues struct {
	ObjectValues *ObjectValues `protobuf:"bytes,4,opt,name=object_values,json=objectValues,proto3,oneof"`
}

type ListValue_DateValues struct {
	DateValues *DateValues `protobuf:"bytes,5,opt,name=date_values,json=dateValues,proto3,oneof"`
}

type ListValue_UuidValues struct {
	UuidValues *UuidValues `protobuf:"bytes,6,opt,name=uuid_values,json=uuidValues,proto3,oneof"`
}

type ListValue_IntValues struct {
	IntValues *IntValues `protobuf:"bytes,7,opt,name=int_values,json=intValues,proto3,oneof"`
}

type ListValue_TextValues struct {
	TextValues *TextValues `protobuf:"bytes,8,opt,name=text_values,json=textValues,proto3,oneof"`
}

func (*ListValue_NumberValues) isListValue_Kind() {}

func (*ListValue_BoolValues) isListValue_Kind() {}

func (*ListValue_ObjectValues) isListValue_Kind() {}

func (*ListValue_DateValues) isListValue_Kind() {}

func (*ListValue_UuidValues) isListValue_Kind() {}

func (*ListValue_IntValues) isListValue_Kind() {}

func (*ListValue_TextValues) isListValue_Kind() {}

type NumberValues struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// *
	// The values are stored as a byte array, where each 8 bytes represent a single float64 value.
	// The byte array is stored in little-endian order using uint64 encoding.
	Values        []byte `protobuf:"bytes,1,opt,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NumberValues) Reset() {
	*x = NumberValues{}
	mi := &file_v1_properties_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NumberValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberValues) ProtoMessage() {}

func (x *NumberValues) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberValues.ProtoReflect.Descriptor instead.
func (*NumberValues) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{3}
}

func (x *NumberValues) GetValues() []byte {
	if x != nil {
		return x.Values
	}
	return nil
}

type TextValues struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []string               `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TextValues) Reset() {
	*x = TextValues{}
	mi := &file_v1_properties_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TextValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextValues) ProtoMessage() {}

func (x *TextValues) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextValues.ProtoReflect.Descriptor instead.
func (*TextValues) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{4}
}

func (x *TextValues) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type BoolValues struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []bool                 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BoolValues) Reset() {
	*x = BoolValues{}
	mi := &file_v1_properties_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoolValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolValues) ProtoMessage() {}

func (x *BoolValues) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolValues.ProtoReflect.Descriptor instead.
func (*BoolValues) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{5}
}

func (x *BoolValues) GetValues() []bool {
	if x != nil {
		return x.Values
	}
	return nil
}

type ObjectValues struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []*Properties          `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectValues) Reset() {
	*x = ObjectValues{}
	mi := &file_v1_properties_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectValues) ProtoMessage() {}

func (x *ObjectValues) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectValues.ProtoReflect.Descriptor instead.
func (*ObjectValues) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{6}
}

func (x *ObjectValues) GetValues() []*Properties {
	if x != nil {
		return x.Values
	}
	return nil
}

type DateValues struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []string               `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DateValues) Reset() {
	*x = DateValues{}
	mi := &file_v1_properties_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateValues) ProtoMessage() {}

func (x *DateValues) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateValues.ProtoReflect.Descriptor instead.
func (*DateValues) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{7}
}

func (x *DateValues) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type UuidValues struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []string               `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UuidValues) Reset() {
	*x = UuidValues{}
	mi := &file_v1_properties_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UuidValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UuidValues) ProtoMessage() {}

func (x *UuidValues) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UuidValues.ProtoReflect.Descriptor instead.
func (*UuidValues) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{8}
}

func (x *UuidValues) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type IntValues struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// *
	// The values are stored as a byte array, where each 8 bytes represent a single int64 value.
	// The byte array is stored in little-endian order using uint64 encoding.
	Values        []byte `protobuf:"bytes,1,opt,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IntValues) Reset() {
	*x = IntValues{}
	mi := &file_v1_properties_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntValues) ProtoMessage() {}

func (x *IntValues) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntValues.ProtoReflect.Descriptor instead.
func (*IntValues) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{9}
}

func (x *IntValues) GetValues() []byte {
	if x != nil {
		return x.Values
	}
	return nil
}

type GeoCoordinate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Longitude     float32                `protobuf:"fixed32,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude      float32                `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeoCoordinate) Reset() {
	*x = GeoCoordinate{}
	mi := &file_v1_properties_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCoordinate) ProtoMessage() {}

func (x *GeoCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoCoordinate.ProtoReflect.Descriptor instead.
func (*GeoCoordinate) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{10}
}

func (x *GeoCoordinate) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *GeoCoordinate) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type PhoneNumber struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	CountryCode            uint64                 `protobuf:"varint,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	DefaultCountry         string                 `protobuf:"bytes,2,opt,name=default_country,json=defaultCountry,proto3" json:"default_country,omitempty"`
	Input                  string                 `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	InternationalFormatted string                 `protobuf:"bytes,4,opt,name=international_formatted,json=internationalFormatted,proto3" json:"international_formatted,omitempty"`
	National               uint64                 `protobuf:"varint,5,opt,name=national,proto3" json:"national,omitempty"`
	NationalFormatted      string                 `protobuf:"bytes,6,opt,name=national_formatted,json=nationalFormatted,proto3" json:"national_formatted,omitempty"`
	Valid                  bool                   `protobuf:"varint,7,opt,name=valid,proto3" json:"valid,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *PhoneNumber) Reset() {
	*x = PhoneNumber{}
	mi := &file_v1_properties_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneNumber) ProtoMessage() {}

func (x *PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneNumber.ProtoReflect.Descriptor instead.
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{11}
}

func (x *PhoneNumber) GetCountryCode() uint64 {
	if x != nil {
		return x.CountryCode
	}
	return 0
}

func (x *PhoneNumber) GetDefaultCountry() string {
	if x != nil {
		return x.DefaultCountry
	}
	return ""
}

func (x *PhoneNumber) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *PhoneNumber) GetInternationalFormatted() string {
	if x != nil {
		return x.InternationalFormatted
	}
	return ""
}

func (x *PhoneNumber) GetNational() uint64 {
	if x != nil {
		return x.National
	}
	return 0
}

func (x *PhoneNumber) GetNationalFormatted() string {
	if x != nil {
		return x.NationalFormatted
	}
	return ""
}

func (x *PhoneNumber) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_v1_properties_proto protoreflect.FileDescriptor

const file_v1_properties_proto_rawDesc = "" +
	"\n" +
	"\x13v1/properties.proto\x12\vweaviate.v1\x1a\x1cgoogle/protobuf/struct.proto\"\x98\x01\n" +
	"\n" +
	"Properties\x12;\n" +
	"\x06fields\x18\x01 \x03(\v2#.weaviate.v1.Properties.FieldsEntryR\x06fields\x1aM\n" +
	"\vFieldsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12(\n" +
	"\x05value\x18\x02 \x01(\v2\x12.weaviate.v1.ValueR\x05value:\x028\x01\"\xcd\x04\n" +
	"\x05Value\x12#\n" +
	"\fnumber_value\x18\x01 \x01(\x01H\x00R\vnumberValue\x12'\n" +
	"\fstring_value\x18\x02 \x01(\tB\x02\x18\x01H\x00R\vstringValue\x12\x1f\n" +
	"\n" +
	"bool_value\x18\x03 \x01(\bH\x00R\tboolValue\x12<\n" +
	"\fobject_value\x18\x04 \x01(\v2\x17.weaviate.v1.PropertiesH\x00R\vobjectValue\x127\n" +
	"\n" +
	"list_value\x18\x05 \x01(\v2\x16.weaviate.v1.ListValueH\x00R\tlistValue\x12\x1f\n" +
	"\n" +
	"date_value\x18\x06 \x01(\tH\x00R\tdateValue\x12\x1f\n" +
	"\n" +
	"uuid_value\x18\a \x01(\tH\x00R\tuuidValue\x12\x1d\n" +
	"\tint_value\x18\b \x01(\x03H\x00R\bintValue\x129\n" +
	"\tgeo_value\x18\t \x01(\v2\x1a.weaviate.v1.GeoCoordinateH\x00R\bgeoValue\x12\x1f\n" +
	"\n" +
	"blob_value\x18\n" +
	" \x01(\tH\x00R\tblobValue\x12;\n" +
	"\vphone_value\x18\v \x01(\v2\x18.weaviate.v1.PhoneNumberH\x00R\n" +
	"phoneValue\x12;\n" +
	"\n" +
	"null_value\x18\f \x01(\x0e2\x1a.google.protobuf.NullValueH\x00R\tnullValue\x12\x1f\n" +
	"\n" +
	"text_value\x18\r \x01(\tH\x00R\ttextValueB\x06\n" +
	"\x04kind\"\xf0\x03\n" +
	"\tListValue\x12.\n" +
	"\x06values\x18\x01 \x03(\v2\x12.weaviate.v1.ValueB\x02\x18\x01R\x06values\x12@\n" +
	"\rnumber_values\x18\x02 \x01(\v2\x19.weaviate.v1.NumberValuesH\x00R\fnumberValues\x12:\n" +
	"\vbool_values\x18\x03 \x01(\v2\x17.weaviate.v1.BoolValuesH\x00R\n" +
	"boolValues\x12@\n" +
	"\robject_values\x18\x04 \x01(\v2\x19.weaviate.v1.ObjectValuesH\x00R\fobjectValues\x12:\n" +
	"\vdate_values\x18\x05 \x01(\v2\x17.weaviate.v1.DateValuesH\x00R\n" +
	"dateValues\x12:\n" +
	"\vuuid_values\x18\x06 \x01(\v2\x17.weaviate.v1.UuidValuesH\x00R\n" +
	"uuidValues\x127\n" +
	"\n" +
	"int_values\x18\a \x01(\v2\x16.weaviate.v1.IntValuesH\x00R\tintValues\x12:\n" +
	"\vtext_values\x18\b \x01(\v2\x17.weaviate.v1.TextValuesH\x00R\n" +
	"textValuesB\x06\n" +
	"\x04kind\"&\n" +
	"\fNumberValues\x12\x16\n" +
	"\x06values\x18\x01 \x01(\fR\x06values\"$\n" +
	"\n" +
	"TextValues\x12\x16\n" +
	"\x06values\x18\x01 \x03(\tR\x06values\"$\n" +
	"\n" +
	"BoolValues\x12\x16\n" +
	"\x06values\x18\x01 \x03(\bR\x06values\"?\n" +
	"\fObjectValues\x12/\n" +
	"\x06values\x18\x01 \x03(\v2\x17.weaviate.v1.PropertiesR\x06values\"$\n" +
	"\n" +
	"DateValues\x12\x16\n" +
	"\x06values\x18\x01 \x03(\tR\x06values\"$\n" +
	"\n" +
	"UuidValues\x12\x16\n" +
	"\x06values\x18\x01 \x03(\tR\x06values\"#\n" +
	"\tIntValues\x12\x16\n" +
	"\x06values\x18\x01 \x01(\fR\x06values\"I\n" +
	"\rGeoCoordinate\x12\x1c\n" +
	"\tlongitude\x18\x01 \x01(\x02R\tlongitude\x12\x1a\n" +
	"\blatitude\x18\x02 \x01(\x02R\blatitude\"\x89\x02\n" +
	"\vPhoneNumber\x12!\n" +
	"\fcountry_code\x18\x01 \x01(\x04R\vcountryCode\x12'\n" +
	"\x0fdefault_country\x18\x02 \x01(\tR\x0edefaultCountry\x12\x14\n" +
	"\x05input\x18\x03 \x01(\tR\x05input\x127\n" +
	"\x17international_formatted\x18\x04 \x01(\tR\x16internationalFormatted\x12\x1a\n" +
	"\bnational\x18\x05 \x01(\x04R\bnational\x12-\n" +
	"\x12national_formatted\x18\x06 \x01(\tR\x11nationalFormatted\x12\x14\n" +
	"\x05valid\x18\a \x01(\bR\x05validBt\n" +
	"#io.weaviate.client.grpc.protocol.v1B\x17WeaviateProtoPropertiesZ4github.com/weaviate/weaviate/grpc/generated;protocolb\x06proto3"

var (
	file_v1_properties_proto_rawDescOnce sync.Once
	file_v1_properties_proto_rawDescData []byte
)

func file_v1_properties_proto_rawDescGZIP() []byte {
	file_v1_properties_proto_rawDescOnce.Do(func() {
		file_v1_properties_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_properties_proto_rawDesc), len(file_v1_properties_proto_rawDesc)))
	})
	return file_v1_properties_proto_rawDescData
}

var (
	file_v1_properties_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
	file_v1_properties_proto_goTypes  = []any{
		(*Properties)(nil),      // 0: weaviate.v1.Properties
		(*Value)(nil),           // 1: weaviate.v1.Value
		(*ListValue)(nil),       // 2: weaviate.v1.ListValue
		(*NumberValues)(nil),    // 3: weaviate.v1.NumberValues
		(*TextValues)(nil),      // 4: weaviate.v1.TextValues
		(*BoolValues)(nil),      // 5: weaviate.v1.BoolValues
		(*ObjectValues)(nil),    // 6: weaviate.v1.ObjectValues
		(*DateValues)(nil),      // 7: weaviate.v1.DateValues
		(*UuidValues)(nil),      // 8: weaviate.v1.UuidValues
		(*IntValues)(nil),       // 9: weaviate.v1.IntValues
		(*GeoCoordinate)(nil),   // 10: weaviate.v1.GeoCoordinate
		(*PhoneNumber)(nil),     // 11: weaviate.v1.PhoneNumber
		nil,                     // 12: weaviate.v1.Properties.FieldsEntry
		(structpb.NullValue)(0), // 13: google.protobuf.NullValue
	}
)

var file_v1_properties_proto_depIdxs = []int32{
	12, // 0: weaviate.v1.Properties.fields:type_name -> weaviate.v1.Properties.FieldsEntry
	0,  // 1: weaviate.v1.Value.object_value:type_name -> weaviate.v1.Properties
	2,  // 2: weaviate.v1.Value.list_value:type_name -> weaviate.v1.ListValue
	10, // 3: weaviate.v1.Value.geo_value:type_name -> weaviate.v1.GeoCoordinate
	11, // 4: weaviate.v1.Value.phone_value:type_name -> weaviate.v1.PhoneNumber
	13, // 5: weaviate.v1.Value.null_value:type_name -> google.protobuf.NullValue
	1,  // 6: weaviate.v1.ListValue.values:type_name -> weaviate.v1.Value
	3,  // 7: weaviate.v1.ListValue.number_values:type_name -> weaviate.v1.NumberValues
	5,  // 8: weaviate.v1.ListValue.bool_values:type_name -> weaviate.v1.BoolValues
	6,  // 9: weaviate.v1.ListValue.object_values:type_name -> weaviate.v1.ObjectValues
	7,  // 10: weaviate.v1.ListValue.date_values:type_name -> weaviate.v1.DateValues
	8,  // 11: weaviate.v1.ListValue.uuid_values:type_name -> weaviate.v1.UuidValues
	9,  // 12: weaviate.v1.ListValue.int_values:type_name -> weaviate.v1.IntValues
	4,  // 13: weaviate.v1.ListValue.text_values:type_name -> weaviate.v1.TextValues
	0,  // 14: weaviate.v1.ObjectValues.values:type_name -> weaviate.v1.Properties
	1,  // 15: weaviate.v1.Properties.FieldsEntry.value:type_name -> weaviate.v1.Value
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_v1_properties_proto_init() }
func file_v1_properties_proto_init() {
	if File_v1_properties_proto != nil {
		return
	}
	file_v1_properties_proto_msgTypes[1].OneofWrappers = []any{
		(*Value_NumberValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_ObjectValue)(nil),
		(*Value_ListValue)(nil),
		(*Value_DateValue)(nil),
		(*Value_UuidValue)(nil),
		(*Value_IntValue)(nil),
		(*Value_GeoValue)(nil),
		(*Value_BlobValue)(nil),
		(*Value_PhoneValue)(nil),
		(*Value_NullValue)(nil),
		(*Value_TextValue)(nil),
	}
	file_v1_properties_proto_msgTypes[2].OneofWrappers = []any{
		(*ListValue_NumberValues)(nil),
		(*ListValue_BoolValues)(nil),
		(*ListValue_ObjectValues)(nil),
		(*ListValue_DateValues)(nil),
		(*ListValue_UuidValues)(nil),
		(*ListValue_IntValues)(nil),
		(*ListValue_TextValues)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_properties_proto_rawDesc), len(file_v1_properties_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_properties_proto_goTypes,
		DependencyIndexes: file_v1_properties_proto_depIdxs,
		MessageInfos:      file_v1_properties_proto_msgTypes,
	}.Build()
	File_v1_properties_proto = out.File
	file_v1_properties_proto_goTypes = nil
	file_v1_properties_proto_depIdxs = nil
}
