// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: configcenter.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*GetGroupRequest)(nil)
var _ json.Unmarshaler = (*GetGroupRequest)(nil)
var _ json.Marshaler = (*GetGroupResponse)(nil)
var _ json.Unmarshaler = (*GetGroupResponse)(nil)
var _ json.Marshaler = (*Groups)(nil)
var _ json.Unmarshaler = (*Groups)(nil)
var _ json.Marshaler = (*GetGroupPropertiesRequest)(nil)
var _ json.Unmarshaler = (*GetGroupPropertiesRequest)(nil)
var _ json.Marshaler = (*GetGroupPropertiesResponse)(nil)
var _ json.Unmarshaler = (*GetGroupPropertiesResponse)(nil)
var _ json.Marshaler = (*SaveGroupPropertiesRequest)(nil)
var _ json.Unmarshaler = (*SaveGroupPropertiesRequest)(nil)
var _ json.Marshaler = (*SaveGroupPropertiesResponse)(nil)
var _ json.Unmarshaler = (*SaveGroupPropertiesResponse)(nil)
var _ json.Marshaler = (*GroupProperties)(nil)
var _ json.Unmarshaler = (*GroupProperties)(nil)
var _ json.Marshaler = (*Property)(nil)
var _ json.Unmarshaler = (*Property)(nil)

// GetGroupRequest implement json.Marshaler.
func (m *GetGroupRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetGroupRequest implement json.Marshaler.
func (m *GetGroupRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetGroupResponse implement json.Marshaler.
func (m *GetGroupResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetGroupResponse implement json.Marshaler.
func (m *GetGroupResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Groups implement json.Marshaler.
func (m *Groups) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Groups implement json.Marshaler.
func (m *Groups) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetGroupPropertiesRequest implement json.Marshaler.
func (m *GetGroupPropertiesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetGroupPropertiesRequest implement json.Marshaler.
func (m *GetGroupPropertiesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetGroupPropertiesResponse implement json.Marshaler.
func (m *GetGroupPropertiesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetGroupPropertiesResponse implement json.Marshaler.
func (m *GetGroupPropertiesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SaveGroupPropertiesRequest implement json.Marshaler.
func (m *SaveGroupPropertiesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SaveGroupPropertiesRequest implement json.Marshaler.
func (m *SaveGroupPropertiesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SaveGroupPropertiesResponse implement json.Marshaler.
func (m *SaveGroupPropertiesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SaveGroupPropertiesResponse implement json.Marshaler.
func (m *SaveGroupPropertiesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GroupProperties implement json.Marshaler.
func (m *GroupProperties) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GroupProperties implement json.Marshaler.
func (m *GroupProperties) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Property implement json.Marshaler.
func (m *Property) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Property implement json.Marshaler.
func (m *Property) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
