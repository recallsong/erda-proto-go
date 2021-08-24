// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: credential.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*CreateAccessKeyRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateAccessKeyResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*AccessKey)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteAccessKeyRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteAccessKeyResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetAccessKeyRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetAccessKeyResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DownloadAccessKeyFileRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DownloadAccessKeyFileResponse)(nil)

// CreateAccessKeyRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateAccessKeyRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "subjectType":
				m.SubjectType = vals[0]
			case "subject":
				m.Subject = vals[0]
			case "isSystem":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsSystem = val
			case "description":
				m.Description = vals[0]
			}
		}
	}
	return nil
}

// CreateAccessKeyResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateAccessKeyResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
			case "data.uuid":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.Uuid = vals[0]
			case "data.accessKeyId":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.AccessKeyId = vals[0]
			case "data.secretKey":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.SecretKey = vals[0]
			case "data.isSystem":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.IsSystem = val
			case "data.status":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.Status = vals[0]
			case "data.subjectType":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.SubjectType = vals[0]
			case "data.subject":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.Subject = vals[0]
			case "data.description":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.Description = vals[0]
			case "data.createdAt":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
			case "data.createdAt.seconds":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Seconds = val
			case "data.createdAt.nanos":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Nanos = int32(val)
			case "data.updatedAt":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "data.updatedAt.seconds":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.UpdatedAt.Seconds = val
			case "data.updatedAt.nanos":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.UpdatedAt.Nanos = int32(val)
			}
		}
	}
	return nil
}

// AccessKey implement urlenc.URLValuesUnmarshaler.
func (m *AccessKey) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "uuid":
				m.Uuid = vals[0]
			case "accessKeyId":
				m.AccessKeyId = vals[0]
			case "secretKey":
				m.SecretKey = vals[0]
			case "isSystem":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsSystem = val
			case "status":
				m.Status = vals[0]
			case "subjectType":
				m.SubjectType = vals[0]
			case "subject":
				m.Subject = vals[0]
			case "description":
				m.Description = vals[0]
			case "createdAt":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
			case "createdAt.seconds":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreatedAt.Seconds = val
			case "createdAt.nanos":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.CreatedAt.Nanos = int32(val)
			case "updatedAt":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "updatedAt.seconds":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.UpdatedAt.Seconds = val
			case "updatedAt.nanos":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.UpdatedAt.Nanos = int32(val)
			}
		}
	}
	return nil
}

// DeleteAccessKeyRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteAccessKeyRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "accessKeyId":
				m.AccessKeyId = vals[0]
			}
		}
	}
	return nil
}

// DeleteAccessKeyResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteAccessKeyResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetAccessKeyRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetAccessKeyRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "accessKeyId":
				m.AccessKeyId = vals[0]
			}
		}
	}
	return nil
}

// GetAccessKeyResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetAccessKeyResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
			case "data.uuid":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.Uuid = vals[0]
			case "data.accessKeyId":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.AccessKeyId = vals[0]
			case "data.secretKey":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.SecretKey = vals[0]
			case "data.isSystem":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.IsSystem = val
			case "data.status":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.Status = vals[0]
			case "data.subjectType":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.SubjectType = vals[0]
			case "data.subject":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.Subject = vals[0]
			case "data.description":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				m.Data.Description = vals[0]
			case "data.createdAt":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
			case "data.createdAt.seconds":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Seconds = val
			case "data.createdAt.nanos":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Nanos = int32(val)
			case "data.updatedAt":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "data.updatedAt.seconds":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.UpdatedAt.Seconds = val
			case "data.updatedAt.nanos":
				if m.Data == nil {
					m.Data = &AccessKey{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.UpdatedAt.Nanos = int32(val)
			}
		}
	}
	return nil
}

// DownloadAccessKeyFileRequest implement urlenc.URLValuesUnmarshaler.
func (m *DownloadAccessKeyFileRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "accessKeyId":
				m.AccessKeyId = vals[0]
			case "path":
				m.Path = vals[0]
			}
		}
	}
	return nil
}

// DownloadAccessKeyFileResponse implement urlenc.URLValuesUnmarshaler.
func (m *DownloadAccessKeyFileResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
