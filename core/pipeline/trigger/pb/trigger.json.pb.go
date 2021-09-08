// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: trigger.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*PipelineTriggerRequest)(nil)
var _ json.Unmarshaler = (*PipelineTriggerRequest)(nil)
var _ json.Marshaler = (*PipelineTriggerResponse)(nil)
var _ json.Unmarshaler = (*PipelineTriggerResponse)(nil)

// PipelineTriggerRequest implement json.Marshaler.
func (m *PipelineTriggerRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PipelineTriggerRequest implement json.Marshaler.
func (m *PipelineTriggerRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// PipelineTriggerResponse implement json.Marshaler.
func (m *PipelineTriggerResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PipelineTriggerResponse implement json.Marshaler.
func (m *PipelineTriggerResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}