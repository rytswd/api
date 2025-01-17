// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: telemetry/v1alpha1/telemetry.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/api/type/v1beta1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Telemetry
func (this *Telemetry) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Telemetry
func (this *Telemetry) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing
func (this *Tracing) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing
func (this *Tracing) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_CustomTag
func (this *Tracing_CustomTag) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_CustomTag
func (this *Tracing_CustomTag) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_Literal
func (this *Tracing_Literal) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_Literal
func (this *Tracing_Literal) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_Environment
func (this *Tracing_Environment) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_Environment
func (this *Tracing_Environment) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_RequestHeader
func (this *Tracing_RequestHeader) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_RequestHeader
func (this *Tracing_RequestHeader) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ProviderRef
func (this *ProviderRef) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ProviderRef
func (this *ProviderRef) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	TelemetryMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	TelemetryUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
