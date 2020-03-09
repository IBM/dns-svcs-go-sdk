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

package dnssvcsinstancesv2_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/IBM/dns-svcs-go-sdk/dnssvcsinstancesv2"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`DnsSvcsInstancesV2`, func() {
	Describe(`ListResourceInstances(listResourceInstancesOptions *ListResourceInstancesOptions)`, func() {
		bearerToken := "0ui9876453"
		listResourceInstancesPath := "/resource_instances"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listResourceInstancesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["type"]).To(Equal([]string{"service_instance"}))

				Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_plan_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["sub_type"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"next_url": "NextURL", "resources": [{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "last_operation": {}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}], "rows_count": 9}`)
			}))
			It(`Invoke ListResourceInstances successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsinstancesv2.NewDnsSvcsInstancesV2(&dnssvcsinstancesv2.DnsSvcsInstancesV2Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListResourceInstances(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := new(dnssvcsinstancesv2.ListResourceInstancesOptions)
				listResourceInstancesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Type = core.StringPtr("service_instance")
				listResourceInstancesOptionsModel.Guid = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Name = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourcePlanID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.SubType = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Limit = core.StringPtr("testString")
				listResourceInstancesOptionsModel.UpdatedFrom = core.StringPtr("testString")
				listResourceInstancesOptionsModel.UpdatedTo = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceInstances(listResourceInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateResourceInstance(createResourceInstanceOptions *CreateResourceInstanceOptions)`, func() {
		bearerToken := "0ui9876453"
		createResourceInstancePath := "/resource_instances"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createResourceInstancePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "last_operation": {}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke CreateResourceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsinstancesv2.NewDnsSvcsInstancesV2(&dnssvcsinstancesv2.DnsSvcsInstancesV2Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsModel := new(dnssvcsinstancesv2.CreateResourceInstanceOptions)
				createResourceInstanceOptionsModel.Name = core.StringPtr("my-instance")
				createResourceInstanceOptionsModel.Target = core.StringPtr("bluemix-us-south")
				createResourceInstanceOptionsModel.ResourceGroup = core.StringPtr("5c49eabc-f5e8-5881-a37e-2d100a33b3df")
				createResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("cloudant-standard")
				createResourceInstanceOptionsModel.Tags = []string{"testString"}
				createResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				createResourceInstanceOptionsModel.Parameters = CreateMockMap()

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceInstance(createResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResourceInstance(getResourceInstanceOptions *GetResourceInstanceOptions)`, func() {
		bearerToken := "0ui9876453"
		getResourceInstancePath := "/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourceInstancePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "last_operation": {}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke GetResourceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsinstancesv2.NewDnsSvcsInstancesV2(&dnssvcsinstancesv2.DnsSvcsInstancesV2Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceInstanceOptions model
				getResourceInstanceOptionsModel := new(dnssvcsinstancesv2.GetResourceInstanceOptions)
				getResourceInstanceOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceInstance(getResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteResourceInstance(deleteResourceInstanceOptions *DeleteResourceInstanceOptions)`, func() {
		bearerToken := "0ui9876453"
		deleteResourceInstancePath := "/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteResourceInstancePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(204)
			}))
			It(`Invoke DeleteResourceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsinstancesv2.NewDnsSvcsInstancesV2(&dnssvcsinstancesv2.DnsSvcsInstancesV2Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceInstanceOptions model
				deleteResourceInstanceOptionsModel := new(dnssvcsinstancesv2.DeleteResourceInstanceOptions)
				deleteResourceInstanceOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceInstance(deleteResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

			})
		})
	})
	Describe(`UpdateResourceInstance(updateResourceInstanceOptions *UpdateResourceInstanceOptions)`, func() {
		bearerToken := "0ui9876453"
		updateResourceInstancePath := "/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateResourceInstancePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "last_operation": {}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke UpdateResourceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsinstancesv2.NewDnsSvcsInstancesV2(&dnssvcsinstancesv2.DnsSvcsInstancesV2Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceInstanceOptions model
				updateResourceInstanceOptionsModel := new(dnssvcsinstancesv2.UpdateResourceInstanceOptions)
				updateResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.Name = core.StringPtr("my-new-instance-name")
				updateResourceInstanceOptionsModel.Parameters = CreateMockMap()
				updateResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("a8dff6d3-d287-4668-a81d-c87c55c2656d")
				updateResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockMap() successfully`, func() {
			mockMap := CreateMockMap()
			Expect(mockMap).ToNot(BeNil())
		})
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, len(mockData))
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Now())
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Now())
	return &d
}
