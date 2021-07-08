// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meta.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListMetricNamesRequest) Validate() error {
	if this.Scope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must not be an empty string`, this.Scope))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *ListMetricNamesResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *ListMetricMetaRequest) Validate() error {
	if this.Scope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must not be an empty string`, this.Scope))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *ListMetricMetaResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *ListMetricGroupsRequest) Validate() error {
	if this.Scope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must not be an empty string`, this.Scope))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *ListMetricGroupsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetMetricGroupRequest) Validate() error {
	if this.Scope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must not be an empty string`, this.Scope))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *GetMetricGroupResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *MetricGroup) Validate() error {
	if this.Meta != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Meta); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Meta", err)
		}
	}
	for _, item := range this.Metrics {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Metrics", err)
			}
		}
	}
	return nil
}
func (this *GroupMetricMeta) Validate() error {
	for _, item := range this.Filters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Filters", err)
			}
		}
	}
	for _, item := range this.Fields {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Fields", err)
			}
		}
	}
	for _, item := range this.Tags {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Tags", err)
			}
		}
	}
	return nil
}
func (this *MetricMeta) Validate() error {
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *NameDefine) Validate() error {
	return nil
}
func (this *TagDefine) Validate() error {
	for _, item := range this.Values {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Values", err)
			}
		}
	}
	return nil
}
func (this *FieldDefine) Validate() error {
	for _, item := range this.Values {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Values", err)
			}
		}
	}
	return nil
}
func (this *ValueDefine) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *Group) Validate() error {
	for _, item := range this.Children {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Children", err)
			}
		}
	}
	return nil
}
func (this *MetaMode) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Filters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Filters", err)
			}
		}
	}
	return nil
}
func (this *TypeDefine) Validate() error {
	for _, item := range this.Aggregations {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Aggregations", err)
			}
		}
	}
	for _, item := range this.Operations {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Operations", err)
			}
		}
	}
	for _, item := range this.Filters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Filters", err)
			}
		}
	}
	return nil
}
func (this *Aggregation) Validate() error {
	return nil
}
func (this *Operation) Validate() error {
	return nil
}
func (this *TagFilter) Validate() error {
	return nil
}
