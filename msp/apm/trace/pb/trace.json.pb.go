// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: trace.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*GetSpansRequest)(nil)
var _ json.Unmarshaler = (*GetSpansRequest)(nil)
var _ json.Marshaler = (*GetTracesRequest)(nil)
var _ json.Unmarshaler = (*GetTracesRequest)(nil)
var _ json.Marshaler = (*GetSpansResponse)(nil)
var _ json.Unmarshaler = (*GetSpansResponse)(nil)
var _ json.Marshaler = (*GetTracesResponse)(nil)
var _ json.Unmarshaler = (*GetTracesResponse)(nil)
var _ json.Marshaler = (*Span)(nil)
var _ json.Unmarshaler = (*Span)(nil)
var _ json.Marshaler = (*Trace)(nil)
var _ json.Unmarshaler = (*Trace)(nil)

// GetSpansRequest implement json.Marshaler.
func (m *GetSpansRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetSpansRequest implement json.Marshaler.
func (m *GetSpansRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetTracesRequest implement json.Marshaler.
func (m *GetTracesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetTracesRequest implement json.Marshaler.
func (m *GetTracesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetSpansResponse implement json.Marshaler.
func (m *GetSpansResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetSpansResponse implement json.Marshaler.
func (m *GetSpansResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetTracesResponse implement json.Marshaler.
func (m *GetTracesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetTracesResponse implement json.Marshaler.
func (m *GetTracesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Span implement json.Marshaler.
func (m *Span) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Span implement json.Marshaler.
func (m *Span) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Trace implement json.Marshaler.
func (m *Trace) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Trace implement json.Marshaler.
func (m *Trace) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
