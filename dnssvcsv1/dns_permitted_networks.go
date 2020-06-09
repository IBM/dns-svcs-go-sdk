/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package dnssvcsv1 : Operations and models for the DnsSvcsV1 service
package dnssvcsv1

import (
	"encoding/json"
	"fmt"
	"reflect"

	common "github.com/IBM/dns-svcs-go-sdk/common"
	"github.com/IBM/go-sdk-core/v4/core"
)

// ListPermittedNetworks : List permitted networks
// List the permitted networks for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) ListPermittedNetworks(listPermittedNetworksOptions *ListPermittedNetworksOptions) (result *ListPermittedNetworks, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPermittedNetworksOptions, "listPermittedNetworksOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPermittedNetworksOptions, "listPermittedNetworksOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*listPermittedNetworksOptions.InstanceID, *listPermittedNetworksOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPermittedNetworksOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListPermittedNetworks")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPermittedNetworksOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listPermittedNetworksOptions.XCorrelationID))
	}

	if listPermittedNetworksOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listPermittedNetworksOptions.Offset))
	}
	if listPermittedNetworksOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listPermittedNetworksOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListPermittedNetworks)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreatePermittedNetwork : Create a permitted network
// Create a permitted network for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) CreatePermittedNetwork(createPermittedNetworkOptions *CreatePermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPermittedNetworkOptions, "createPermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPermittedNetworkOptions, "createPermittedNetworkOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*createPermittedNetworkOptions.InstanceID, *createPermittedNetworkOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreatePermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createPermittedNetworkOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createPermittedNetworkOptions.Type != nil {
		body["type"] = createPermittedNetworkOptions.Type
	}
	if createPermittedNetworkOptions.PermittedNetwork != nil {
		body["permitted_network"] = createPermittedNetworkOptions.PermittedNetwork
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPermittedNetwork)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeletePermittedNetwork : Remove a permitted network
// Remove a permitted network.
func (dnsSvcs *DnsSvcsV1) DeletePermittedNetwork(deletePermittedNetworkOptions *DeletePermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePermittedNetworkOptions, "deletePermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePermittedNetworkOptions, "deletePermittedNetworkOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*deletePermittedNetworkOptions.InstanceID, *deletePermittedNetworkOptions.DnszoneID, *deletePermittedNetworkOptions.PermittedNetworkID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeletePermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deletePermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deletePermittedNetworkOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPermittedNetwork)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetPermittedNetwork : Get a permitted network
// Get details of a permitted network.
func (dnsSvcs *DnsSvcsV1) GetPermittedNetwork(getPermittedNetworkOptions *GetPermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPermittedNetworkOptions, "getPermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPermittedNetworkOptions, "getPermittedNetworkOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*getPermittedNetworkOptions.InstanceID, *getPermittedNetworkOptions.DnszoneID, *getPermittedNetworkOptions.PermittedNetworkID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetPermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getPermittedNetworkOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPermittedNetwork)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreatePermittedNetworkOptions : The CreatePermittedNetwork options.
type CreatePermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The type of a permitted network.
	Type *string `json:"type,omitempty"`

	// Permitted network data for VPC.
	PermittedNetwork *PermittedNetworkVpc `json:"permitted_network,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreatePermittedNetworkOptions.Type property.
// The type of a permitted network.
const (
	CreatePermittedNetworkOptions_Type_Vpc = "vpc"
)

// NewCreatePermittedNetworkOptions : Instantiate CreatePermittedNetworkOptions
func (*DnsSvcsV1) NewCreatePermittedNetworkOptions(instanceID string, dnszoneID string) *CreatePermittedNetworkOptions {
	return &CreatePermittedNetworkOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreatePermittedNetworkOptions) SetInstanceID(instanceID string) *CreatePermittedNetworkOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *CreatePermittedNetworkOptions) SetDnszoneID(dnszoneID string) *CreatePermittedNetworkOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetPermittedNetwork : Allow user to set PermittedNetwork
func (options *CreatePermittedNetworkOptions) SetPermittedNetwork(permittedNetwork *PermittedNetworkVpc) *CreatePermittedNetworkOptions {
	options.PermittedNetwork = permittedNetwork
	return options
}

// SetType : Allow user to set Type
func (options *CreatePermittedNetworkOptions) SetType(typeVar string) *CreatePermittedNetworkOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreatePermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *CreatePermittedNetworkOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePermittedNetworkOptions) SetHeaders(param map[string]string) *CreatePermittedNetworkOptions {
	options.Headers = param
	return options
}

// DeletePermittedNetworkOptions : The DeletePermittedNetwork options.
type DeletePermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a permitted network.
	PermittedNetworkID *string `json:"permitted_network_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePermittedNetworkOptions : Instantiate DeletePermittedNetworkOptions
func (*DnsSvcsV1) NewDeletePermittedNetworkOptions(instanceID string, dnszoneID string, permittedNetworkID string) *DeletePermittedNetworkOptions {
	return &DeletePermittedNetworkOptions{
		InstanceID:         core.StringPtr(instanceID),
		DnszoneID:          core.StringPtr(dnszoneID),
		PermittedNetworkID: core.StringPtr(permittedNetworkID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeletePermittedNetworkOptions) SetInstanceID(instanceID string) *DeletePermittedNetworkOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *DeletePermittedNetworkOptions) SetDnszoneID(dnszoneID string) *DeletePermittedNetworkOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetPermittedNetworkID : Allow user to set PermittedNetworkID
func (options *DeletePermittedNetworkOptions) SetPermittedNetworkID(permittedNetworkID string) *DeletePermittedNetworkOptions {
	options.PermittedNetworkID = core.StringPtr(permittedNetworkID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeletePermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *DeletePermittedNetworkOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePermittedNetworkOptions) SetHeaders(param map[string]string) *DeletePermittedNetworkOptions {
	options.Headers = param
	return options
}

// GetPermittedNetworkOptions : The GetPermittedNetwork options.
type GetPermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a permitted network.
	PermittedNetworkID *string `json:"permitted_network_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPermittedNetworkOptions : Instantiate GetPermittedNetworkOptions
func (*DnsSvcsV1) NewGetPermittedNetworkOptions(instanceID string, dnszoneID string, permittedNetworkID string) *GetPermittedNetworkOptions {
	return &GetPermittedNetworkOptions{
		InstanceID:         core.StringPtr(instanceID),
		DnszoneID:          core.StringPtr(dnszoneID),
		PermittedNetworkID: core.StringPtr(permittedNetworkID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetPermittedNetworkOptions) SetInstanceID(instanceID string) *GetPermittedNetworkOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *GetPermittedNetworkOptions) SetDnszoneID(dnszoneID string) *GetPermittedNetworkOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetPermittedNetworkID : Allow user to set PermittedNetworkID
func (options *GetPermittedNetworkOptions) SetPermittedNetworkID(permittedNetworkID string) *GetPermittedNetworkOptions {
	options.PermittedNetworkID = core.StringPtr(permittedNetworkID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetPermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *GetPermittedNetworkOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPermittedNetworkOptions) SetHeaders(param map[string]string) *GetPermittedNetworkOptions {
	options.Headers = param
	return options
}

// ListPermittedNetworksOptions : The ListPermittedNetworks options.
type ListPermittedNetworksOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Specify how many permitted networks to skip over, the default value is 0.
	Offset *int64 `json:"offset,omitempty"`

	// Specify how many permitted networks are returned, the default value is 10.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPermittedNetworksOptions : Instantiate ListPermittedNetworksOptions
func (*DnsSvcsV1) NewListPermittedNetworksOptions(instanceID string, dnszoneID string) *ListPermittedNetworksOptions {
	return &ListPermittedNetworksOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListPermittedNetworksOptions) SetInstanceID(instanceID string) *ListPermittedNetworksOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *ListPermittedNetworksOptions) SetDnszoneID(dnszoneID string) *ListPermittedNetworksOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListPermittedNetworksOptions) SetXCorrelationID(xCorrelationID string) *ListPermittedNetworksOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListPermittedNetworksOptions) SetOffset(offset int64) *ListPermittedNetworksOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListPermittedNetworksOptions) SetLimit(limit int64) *ListPermittedNetworksOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListPermittedNetworksOptions) SetHeaders(param map[string]string) *ListPermittedNetworksOptions {
	options.Headers = param
	return options
}

// ListPermittedNetworks : List permitted networks response.
type ListPermittedNetworks struct {
	// An array of permitted networks.
	PermittedNetworks []PermittedNetwork `json:"permitted_networks" validate:"required"`

	// Specify how many permitted networks to skip over, the default value is 0.
	Offset *int64 `json:"offset" validate:"required"`

	// Specify how many permitted networks are returned, the default value is 10.
	Limit *int64 `json:"limit" validate:"required"`

	// Total number of permitted networks.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next,omitempty"`
}

// UnmarshalListPermittedNetworks unmarshals an instance of ListPermittedNetworks from the specified map of raw messages.
func UnmarshalListPermittedNetworks(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListPermittedNetworks)
	err = core.UnmarshalModel(m, "permitted_networks", &obj.PermittedNetworks, UnmarshalPermittedNetwork)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PermittedNetwork : Permitted network details.
type PermittedNetwork struct {
	// Unique identifier of a permitted network.
	ID *string `json:"id,omitempty"`

	// The time when a permitted network is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// The recent time when a permitted network is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// Permitted network data for VPC.
	PermittedNetwork *PermittedNetworkVpc `json:"permitted_network,omitempty"`

	// The type of a permitted network.
	Type *string `json:"type,omitempty"`

	// The state of a permitted network.
	State *string `json:"state,omitempty"`
}

// Constants associated with the PermittedNetwork.Type property.
// The type of a permitted network.
const (
	PermittedNetwork_Type_Vpc = "vpc"
)

// Constants associated with the PermittedNetwork.State property.
// The state of a permitted network.
const (
	PermittedNetwork_State_Active            = "ACTIVE"
	PermittedNetwork_State_RemovalInProgress = "REMOVAL_IN_PROGRESS"
)

// UnmarshalPermittedNetwork unmarshals an instance of PermittedNetwork from the specified map of raw messages.
func UnmarshalPermittedNetwork(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PermittedNetwork)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "permitted_network", &obj.PermittedNetwork, UnmarshalPermittedNetworkVpc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PermittedNetworkVpc : Permitted network data for VPC.
type PermittedNetworkVpc struct {
	// CRN string uniquely identifies a VPC.
	VpcCrn *string `json:"vpc_crn" validate:"required"`
}

// NewPermittedNetworkVpc : Instantiate PermittedNetworkVpc (Generic Model Constructor)
func (*DnsSvcsV1) NewPermittedNetworkVpc(vpcCrn string) (model *PermittedNetworkVpc, err error) {
	model = &PermittedNetworkVpc{
		VpcCrn: core.StringPtr(vpcCrn),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalPermittedNetworkVpc unmarshals an instance of PermittedNetworkVpc from the specified map of raw messages.
func UnmarshalPermittedNetworkVpc(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PermittedNetworkVpc)
	err = core.UnmarshalPrimitive(m, "vpc_crn", &obj.VpcCrn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
