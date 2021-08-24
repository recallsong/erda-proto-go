// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: adapter.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetInstrumentationLibraryRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetInstrumentationLibraryResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetInstrumentationLibraryDocsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetInstrumentationLibraryDocsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Language)(nil)
var _ urlenc.URLValuesUnmarshaler = (*InstrumentationLibrary)(nil)

// GetInstrumentationLibraryRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetInstrumentationLibraryRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetInstrumentationLibraryResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetInstrumentationLibraryResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetInstrumentationLibraryDocsRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetInstrumentationLibraryDocsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			case "serviceName":
				m.ServiceName = vals[0]
			case "type":
				m.Type = vals[0]
			case "language":
				m.Language = vals[0]
			case "strategy":
				m.Strategy = vals[0]
			}
		}
	}
	return nil
}

// GetInstrumentationLibraryDocsResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetInstrumentationLibraryDocsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// Language implement urlenc.URLValuesUnmarshaler.
func (m *Language) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "language":
				m.Language = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			}
		}
	}
	return nil
}

// InstrumentationLibrary implement urlenc.URLValuesUnmarshaler.
func (m *InstrumentationLibrary) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "displayName":
				m.DisplayName = vals[0]
			case "strategy":
				m.Strategy = vals[0]
			}
		}
	}
	return nil
}
