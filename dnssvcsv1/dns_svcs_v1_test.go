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

package dnssvcsv1_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`DnsSvcsV1`, func() {
	Describe(`ListDnszones(listDnszonesOptions *ListDnszonesOptions)`, func() {
		listDnszonesPath := "/instances/testString/dnszones"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listDnszonesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"count": 1, "dnszones": [{"created_on": "2019-01-01T05:20:00.12345Z", "description": "The DNS zone is used for VPCs in us-east region", "id": "example.com:2d0f862b-67cc-41f3-b6a2-59860d0aa90e", "instance_id": "1407a753-a93f-4bb0-9784-bcfc269ee1b3", "label": "us-east", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "example.com", "state": "pending_network_add"}], "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones?page=1&per_page=20"}, "limit": 20, "next": {"href": "https://api.pdns.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones?page=2&per_page=20"}, "offset": 1, "total_count": 200}`)
			}))
			It(`Invoke ListDnszones successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListDnszones(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDnszonesOptions model
				listDnszonesOptionsModel := new(dnssvcsv1.ListDnszonesOptions)
				listDnszonesOptionsModel.InstanceID = core.StringPtr("testString")
				listDnszonesOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListDnszones(listDnszonesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateDnszone(createDnszoneOptions *CreateDnszoneOptions)`, func() {
		createDnszonePath := "/instances/testString/dnszones"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createDnszonePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_on": "2019-01-01T05:20:00.12345Z", "description": "The DNS zone is used for VPCs in us-east region", "id": "example.com:2d0f862b-67cc-41f3-b6a2-59860d0aa90e", "instance_id": "1407a753-a93f-4bb0-9784-bcfc269ee1b3", "label": "us-east", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "example.com", "state": "pending_network_add"}`)
			}))
			It(`Invoke CreateDnszone successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateDnszone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDnszoneOptions model
				createDnszoneOptionsModel := new(dnssvcsv1.CreateDnszoneOptions)
				createDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				createDnszoneOptionsModel.Description = core.StringPtr("The DNS zone is used for VPCs in us-east region")
				createDnszoneOptionsModel.Label = core.StringPtr("us-east")
				createDnszoneOptionsModel.Name = core.StringPtr("example.com")
				createDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateDnszone(createDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteDnszone(deleteDnszoneOptions *DeleteDnszoneOptions)`, func() {
		deleteDnszonePath := "/instances/testString/dnszones/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteDnszonePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.WriteHeader(204)
			}))
			It(`Invoke DeleteDnszone successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteDnszone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDnszoneOptions model
				deleteDnszoneOptionsModel := new(dnssvcsv1.DeleteDnszoneOptions)
				deleteDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				deleteDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteDnszone(deleteDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetDnszone(getDnszoneOptions *GetDnszoneOptions)`, func() {
		getDnszonePath := "/instances/testString/dnszones/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getDnszonePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_on": "2019-01-01T05:20:00.12345Z", "description": "The DNS zone is used for VPCs in us-east region", "id": "example.com:2d0f862b-67cc-41f3-b6a2-59860d0aa90e", "instance_id": "1407a753-a93f-4bb0-9784-bcfc269ee1b3", "label": "us-east", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "example.com", "state": "pending_network_add"}`)
			}))
			It(`Invoke GetDnszone successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetDnszone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDnszoneOptions model
				getDnszoneOptionsModel := new(dnssvcsv1.GetDnszoneOptions)
				getDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				getDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				getDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetDnszone(getDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateDnszone(updateDnszoneOptions *UpdateDnszoneOptions)`, func() {
		updateDnszonePath := "/instances/testString/dnszones/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateDnszonePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_on": "2019-01-01T05:20:00.12345Z", "description": "The DNS zone is used for VPCs in us-east region", "id": "example.com:2d0f862b-67cc-41f3-b6a2-59860d0aa90e", "instance_id": "1407a753-a93f-4bb0-9784-bcfc269ee1b3", "label": "us-east", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "example.com", "state": "pending_network_add"}`)
			}))
			It(`Invoke UpdateDnszone successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateDnszone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateDnszoneOptions model
				updateDnszoneOptionsModel := new(dnssvcsv1.UpdateDnszoneOptions)
				updateDnszoneOptionsModel.InstanceID = core.StringPtr("testString")
				updateDnszoneOptionsModel.DnszoneID = core.StringPtr("testString")
				updateDnszoneOptionsModel.Description = core.StringPtr("The DNS zone is used for VPCs in us-east region")
				updateDnszoneOptionsModel.Label = core.StringPtr("us-east")
				updateDnszoneOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateDnszone(updateDnszoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListPermittedNetworks(listPermittedNetworksOptions *ListPermittedNetworksOptions)`, func() {
		listPermittedNetworksPath := "/instances/testString/dnszones/testString/permitted_networks"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listPermittedNetworksPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"permitted_networks": [{"permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "created_on": "2019-01-01T05:20:00.12345Z", "id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "state": "ACTIVE", "modified_on": "2019-01-01T05:20:00.12345Z", "type": "vpc"}], "count": 1, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones?page=1&per_page=20"}, "limit": 20, "next": {"href": "https://api.pdns.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones?page=2&per_page=20"}, "offset": 1, "total_count": 200}`)
			}))
			It(`Invoke ListPermittedNetworks successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListPermittedNetworks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPermittedNetworksOptions model
				listPermittedNetworksOptionsModel := new(dnssvcsv1.ListPermittedNetworksOptions)
				listPermittedNetworksOptionsModel.InstanceID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.DnszoneID = core.StringPtr("testString")
				listPermittedNetworksOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListPermittedNetworks(listPermittedNetworksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreatePermittedNetwork(createPermittedNetworkOptions *CreatePermittedNetworkOptions)`, func() {
		createPermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createPermittedNetworkPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "created_on": "2019-01-01T05:20:00.12345Z", "id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "state": "ACTIVE", "modified_on": "2019-01-01T05:20:00.12345Z", "type": "vpc"}`)
			}))
			It(`Invoke CreatePermittedNetwork successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreatePermittedNetwork(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PermittedNetworkVpc model
				permittedNetworkVpcModel := new(dnssvcsv1.PermittedNetworkVpc)
				permittedNetworkVpcModel.VpcCrn = core.StringPtr("crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6")

				// Construct an instance of the CreatePermittedNetworkOptions model
				createPermittedNetworkOptionsModel := new(dnssvcsv1.CreatePermittedNetworkOptions)
				createPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				createPermittedNetworkOptionsModel.PermittedNetwork = &dnssvcsv1.PermittedNetworkVpc{
					VpcCrn: permittedNetworkVpcModel.VpcCrn,
				}
				createPermittedNetworkOptionsModel.Type = core.StringPtr("vpc")
				createPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreatePermittedNetwork(createPermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeletePermittedNetwork(deletePermittedNetworkOptions *DeletePermittedNetworkOptions)`, func() {
		deletePermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deletePermittedNetworkPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{"permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "created_on": "2019-01-01T05:20:00.12345Z", "id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "state": "ACTIVE", "modified_on": "2019-01-01T05:20:00.12345Z", "type": "vpc"}`)
			}))
			It(`Invoke DeletePermittedNetwork successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeletePermittedNetwork(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeletePermittedNetworkOptions model
				deletePermittedNetworkOptionsModel := new(dnssvcsv1.DeletePermittedNetworkOptions)
				deletePermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				deletePermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeletePermittedNetwork(deletePermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetPermittedNetwork(getPermittedNetworkOptions *GetPermittedNetworkOptions)`, func() {
		getPermittedNetworkPath := "/instances/testString/dnszones/testString/permitted_networks/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getPermittedNetworkPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"permitted_network": {"vpc_crn": "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"}, "created_on": "2019-01-01T05:20:00.12345Z", "id": "fecd0173-3919-456b-b202-3029dfa1b0f7", "state": "ACTIVE", "modified_on": "2019-01-01T05:20:00.12345Z", "type": "vpc"}`)
			}))
			It(`Invoke GetPermittedNetwork successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetPermittedNetwork(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPermittedNetworkOptions model
				getPermittedNetworkOptionsModel := new(dnssvcsv1.GetPermittedNetworkOptions)
				getPermittedNetworkOptionsModel.InstanceID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.DnszoneID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.PermittedNetworkID = core.StringPtr("testString")
				getPermittedNetworkOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetPermittedNetwork(getPermittedNetworkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListResourceRecords(listResourceRecordsOptions *ListResourceRecordsOptions)`, func() {
		listResourceRecordsPath := "/instances/testString/dnszones/testString/resource_records"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listResourceRecordsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"count": 1, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones?page=1&per_page=20"}, "limit": 20, "next": {"href": "https://api.pdns.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones?page=2&per_page=20"}, "offset": 1, "resource_records": [{"created_on": "2019-01-01T05:20:00.12345Z", "id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "protocol": "udp", "service": "_sip", "ttl": 120, "type": "SRV"}], "total_count": 200}`)
			}))
			It(`Invoke ListResourceRecords successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListResourceRecords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceRecordsOptions model
				listResourceRecordsOptionsModel := new(dnssvcsv1.ListResourceRecordsOptions)
				listResourceRecordsOptionsModel.InstanceID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.DnszoneID = core.StringPtr("testString")
				listResourceRecordsOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceRecords(listResourceRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateResourceRecord(createResourceRecordOptions *CreateResourceRecordOptions)`, func() {
		createResourceRecordPath := "/instances/testString/dnszones/testString/resource_records"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createResourceRecordPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_on": "2019-01-01T05:20:00.12345Z", "id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "protocol": "udp", "service": "_sip", "ttl": 120, "type": "SRV"}`)
			}))
			It(`Invoke CreateResourceRecord successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceRecordInputRdataRdataARecord model
				resourceRecordInputRdataModel := new(dnssvcsv1.ResourceRecordInputRdataRdataARecord)
				resourceRecordInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the CreateResourceRecordOptions model
				createResourceRecordOptionsModel := new(dnssvcsv1.CreateResourceRecordOptions)
				createResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				createResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				createResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				createResourceRecordOptionsModel.Protocol = core.StringPtr("udp")
				createResourceRecordOptionsModel.Rdata = resourceRecordInputRdataModel
				createResourceRecordOptionsModel.Service = core.StringPtr("_sip")
				createResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				createResourceRecordOptionsModel.Type = core.StringPtr("SRV")
				createResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceRecord(createResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteResourceRecord(deleteResourceRecordOptions *DeleteResourceRecordOptions)`, func() {
		deleteResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteResourceRecordPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.WriteHeader(204)
			}))
			It(`Invoke DeleteResourceRecord successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceRecordOptions model
				deleteResourceRecordOptionsModel := new(dnssvcsv1.DeleteResourceRecordOptions)
				deleteResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				deleteResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceRecord(deleteResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResourceRecord(getResourceRecordOptions *GetResourceRecordOptions)`, func() {
		getResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourceRecordPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_on": "2019-01-01T05:20:00.12345Z", "id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "protocol": "udp", "service": "_sip", "ttl": 120, "type": "SRV"}`)
			}))
			It(`Invoke GetResourceRecord successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceRecordOptions model
				getResourceRecordOptionsModel := new(dnssvcsv1.GetResourceRecordOptions)
				getResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				getResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				getResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				getResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceRecord(getResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateResourceRecord(updateResourceRecordOptions *UpdateResourceRecordOptions)`, func() {
		updateResourceRecordPath := "/instances/testString/dnszones/testString/resource_records/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateResourceRecordPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
				Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"created_on": "2019-01-01T05:20:00.12345Z", "id": "SRV:5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "modified_on": "2019-01-01T05:20:00.12345Z", "name": "_sip._udp.test.example.com", "protocol": "udp", "service": "_sip", "ttl": 120, "type": "SRV"}`)
			}))
			It(`Invoke UpdateResourceRecord successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateResourceRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceRecordUpdateInputRdataRdataARecord model
				resourceRecordUpdateInputRdataModel := new(dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord)
				resourceRecordUpdateInputRdataModel.Ip = core.StringPtr("10.110.201.214")

				// Construct an instance of the UpdateResourceRecordOptions model
				updateResourceRecordOptionsModel := new(dnssvcsv1.UpdateResourceRecordOptions)
				updateResourceRecordOptionsModel.InstanceID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.DnszoneID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.RecordID = core.StringPtr("testString")
				updateResourceRecordOptionsModel.Name = core.StringPtr("test.example.com")
				updateResourceRecordOptionsModel.Rdata = resourceRecordUpdateInputRdataModel
				updateResourceRecordOptionsModel.TTL = core.Int64Ptr(int64(120))
				updateResourceRecordOptionsModel.XCorrelationID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceRecord(updateResourceRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
		Describe(`ListLoadBalancers(listLoadBalancersOptions *ListLoadBalancersOptions)`, func() {
			listLoadBalancersPath := "/instances/testString/dnszones/testString/load_balancers"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listLoadBalancersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"load_balancers": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "name": "glb.example.com", "enabled": true, "ttl": 120, "health": true, "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": {"us-south-1": ["UsSouth1"], "us-south-2": ["UsSouth2"], "us-south-3": ["UsSouth3"], "us-east-1": ["UsEast1"], "us-east-2": ["UsEast2"], "us-east-3": ["UsEast3"], "eu-gb-1": ["EuGb1"], "eu-gb-2": ["EuGb2"], "eu-gb-3": ["EuGb3"], "eu-de-1": ["EuDe1"], "eu-de-2": ["EuDe2"], "eu-de-3": ["EuDe3"], "au-syd-1": ["AuSyd1"], "au-syd-2": ["AuSyd2"], "au-syd-3": ["AuSyd3"], "jp-tok-1": ["JpTok1"], "jp-tok-2": ["JpTok2"], "jp-tok-3": ["JpTok3"]}}], "offset": 1, "limit": 20, "count": 1, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/load_balancers?page=1&per_page=20"}, "next": {"href": "https://api.pdns.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/load_balancers?page=2&per_page=20"}}`)
				}))
				It(`Invoke ListLoadBalancers successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.ListLoadBalancers(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the ListLoadBalancersOptions model
					listLoadBalancersOptionsModel := new(dnssvcsv1.ListLoadBalancersOptions)
					listLoadBalancersOptionsModel.InstanceID = core.StringPtr("testString")
					listLoadBalancersOptionsModel.DnszoneID = core.StringPtr("testString")
					listLoadBalancersOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.ListLoadBalancers(listLoadBalancersOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions)`, func() {
			createLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createLoadBalancerPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "name": "glb.example.com", "enabled": true, "ttl": 120, "health": true, "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": {"us-south-1": ["UsSouth1"], "us-south-2": ["UsSouth2"], "us-south-3": ["UsSouth3"], "us-east-1": ["UsEast1"], "us-east-2": ["UsEast2"], "us-east-3": ["UsEast3"], "eu-gb-1": ["EuGb1"], "eu-gb-2": ["EuGb2"], "eu-gb-3": ["EuGb3"], "eu-de-1": ["EuDe1"], "eu-de-2": ["EuDe2"], "eu-de-3": ["EuDe3"], "au-syd-1": ["AuSyd1"], "au-syd-2": ["AuSyd2"], "au-syd-3": ["AuSyd3"], "jp-tok-1": ["JpTok1"], "jp-tok-2": ["JpTok2"], "jp-tok-3": ["JpTok3"]}}`)
				}))
				It(`Invoke CreateLoadBalancer successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.CreateLoadBalancer(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the AzPools model
					azPoolsModel := new(dnssvcsv1.AzPools)
					azPoolsModel.UsSouth1 = []string{"testString"}
					azPoolsModel.UsSouth2 = []string{"testString"}
					azPoolsModel.UsSouth3 = []string{"testString"}
					azPoolsModel.UsEast1 = []string{"testString"}
					azPoolsModel.UsEast2 = []string{"testString"}
					azPoolsModel.UsEast3 = []string{"testString"}
					azPoolsModel.EuGb1 = []string{"testString"}
					azPoolsModel.EuGb2 = []string{"testString"}
					azPoolsModel.EuGb3 = []string{"testString"}
					azPoolsModel.EuDe1 = []string{"testString"}
					azPoolsModel.EuDe2 = []string{"testString"}
					azPoolsModel.EuDe3 = []string{"testString"}
					azPoolsModel.AuSyd1 = []string{"testString"}
					azPoolsModel.AuSyd2 = []string{"testString"}
					azPoolsModel.AuSyd3 = []string{"testString"}
					azPoolsModel.JpTok1 = []string{"testString"}
					azPoolsModel.JpTok2 = []string{"testString"}
					azPoolsModel.JpTok3 = []string{"testString"}

					// Construct an instance of the CreateLoadBalancerOptions model
					createLoadBalancerOptionsModel := new(dnssvcsv1.CreateLoadBalancerOptions)
					createLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
					createLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
					createLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
					createLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
					createLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
					createLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
					createLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
					createLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
					createLoadBalancerOptionsModel.AzPools = azPoolsModel
					createLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.CreateLoadBalancer(createLoadBalancerOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`DeleteLoadBalancer(deleteLoadBalancerOptions *DeleteLoadBalancerOptions)`, func() {
			deleteLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteLoadBalancerPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
				It(`Invoke DeleteLoadBalancer successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					response, operationErr := testService.DeleteLoadBalancer(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())

					// Construct an instance of the DeleteLoadBalancerOptions model
					deleteLoadBalancerOptionsModel := new(dnssvcsv1.DeleteLoadBalancerOptions)
					deleteLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
					deleteLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
					deleteLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
					deleteLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					response, operationErr = testService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
				})
			})
		})
		Describe(`GetLoadBalancer(getLoadBalancerOptions *GetLoadBalancerOptions)`, func() {
			getLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getLoadBalancerPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "name": "glb.example.com", "enabled": true, "ttl": 120, "health": true, "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": {"us-south-1": ["UsSouth1"], "us-south-2": ["UsSouth2"], "us-south-3": ["UsSouth3"], "us-east-1": ["UsEast1"], "us-east-2": ["UsEast2"], "us-east-3": ["UsEast3"], "eu-gb-1": ["EuGb1"], "eu-gb-2": ["EuGb2"], "eu-gb-3": ["EuGb3"], "eu-de-1": ["EuDe1"], "eu-de-2": ["EuDe2"], "eu-de-3": ["EuDe3"], "au-syd-1": ["AuSyd1"], "au-syd-2": ["AuSyd2"], "au-syd-3": ["AuSyd3"], "jp-tok-1": ["JpTok1"], "jp-tok-2": ["JpTok2"], "jp-tok-3": ["JpTok3"]}}`)
				}))
				It(`Invoke GetLoadBalancer successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.GetLoadBalancer(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the GetLoadBalancerOptions model
					getLoadBalancerOptionsModel := new(dnssvcsv1.GetLoadBalancerOptions)
					getLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
					getLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
					getLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
					getLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.GetLoadBalancer(getLoadBalancerOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`UpdateLoadBalancer(updateLoadBalancerOptions *UpdateLoadBalancerOptions)`, func() {
			updateLoadBalancerPath := "/instances/testString/dnszones/testString/load_balancers/testString"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateLoadBalancerPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "name": "glb.example.com", "enabled": true, "ttl": 120, "health": true, "fallback_pool": "24ccf79a-4ae0-4769-b4c8-17f8f230072e", "default_pools": ["DefaultPools"], "az_pools": {"us-south-1": ["UsSouth1"], "us-south-2": ["UsSouth2"], "us-south-3": ["UsSouth3"], "us-east-1": ["UsEast1"], "us-east-2": ["UsEast2"], "us-east-3": ["UsEast3"], "eu-gb-1": ["EuGb1"], "eu-gb-2": ["EuGb2"], "eu-gb-3": ["EuGb3"], "eu-de-1": ["EuDe1"], "eu-de-2": ["EuDe2"], "eu-de-3": ["EuDe3"], "au-syd-1": ["AuSyd1"], "au-syd-2": ["AuSyd2"], "au-syd-3": ["AuSyd3"], "jp-tok-1": ["JpTok1"], "jp-tok-2": ["JpTok2"], "jp-tok-3": ["JpTok3"]}}`)
				}))
				It(`Invoke UpdateLoadBalancer successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.UpdateLoadBalancer(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the AzPools model
					azPoolsModel := new(dnssvcsv1.AzPools)
					azPoolsModel.UsSouth1 = []string{"testString"}
					azPoolsModel.UsSouth2 = []string{"testString"}
					azPoolsModel.UsSouth3 = []string{"testString"}
					azPoolsModel.UsEast1 = []string{"testString"}
					azPoolsModel.UsEast2 = []string{"testString"}
					azPoolsModel.UsEast3 = []string{"testString"}
					azPoolsModel.EuGb1 = []string{"testString"}
					azPoolsModel.EuGb2 = []string{"testString"}
					azPoolsModel.EuGb3 = []string{"testString"}
					azPoolsModel.EuDe1 = []string{"testString"}
					azPoolsModel.EuDe2 = []string{"testString"}
					azPoolsModel.EuDe3 = []string{"testString"}
					azPoolsModel.AuSyd1 = []string{"testString"}
					azPoolsModel.AuSyd2 = []string{"testString"}
					azPoolsModel.AuSyd3 = []string{"testString"}
					azPoolsModel.JpTok1 = []string{"testString"}
					azPoolsModel.JpTok2 = []string{"testString"}
					azPoolsModel.JpTok3 = []string{"testString"}

					// Construct an instance of the UpdateLoadBalancerOptions model
					updateLoadBalancerOptionsModel := new(dnssvcsv1.UpdateLoadBalancerOptions)
					updateLoadBalancerOptionsModel.InstanceID = core.StringPtr("testString")
					updateLoadBalancerOptionsModel.DnszoneID = core.StringPtr("testString")
					updateLoadBalancerOptionsModel.LbID = core.StringPtr("testString")
					updateLoadBalancerOptionsModel.Name = core.StringPtr("glb.example.com")
					updateLoadBalancerOptionsModel.Description = core.StringPtr("Load balancer for glb.example.com.")
					updateLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
					updateLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(120))
					updateLoadBalancerOptionsModel.FallbackPool = core.StringPtr("24ccf79a-4ae0-4769-b4c8-17f8f230072e")
					updateLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
					updateLoadBalancerOptionsModel.AzPools = azPoolsModel
					updateLoadBalancerOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.UpdateLoadBalancer(updateLoadBalancerOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`Service constructor tests`, func() {
			It(`Instantiate service client`, func() {
				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
			})
			It(`Instantiate service client with error: Invalid URL`, func() {
				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL: "{{BAD_URL_STRING",
				})
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
			})
			It(`Instantiate service client with error: Invalid Auth`, func() {
				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://dnssvcsv1/api",
					Authenticator: &core.BasicAuthenticator{
						Username: "",
						Password: "",
					},
				})
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
			})
		})
		Describe(`Service constructor tests using external config`, func() {
			Context(`Using external config, construct service client instances`, func() {
				// Map containing environment variables used in testing.
				var testEnvironment = map[string]string{
					"DNS_SVCS_URL":       "https://dnssvcsv1/api",
					"DNS_SVCS_AUTH_TYPE": "noauth",
				}

				It(`Create service client using external config successfully`, func() {
					SetTestEnvironment(testEnvironment)
					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{})
					Expect(testService).ToNot(BeNil())
					Expect(testServiceErr).To(BeNil())
					ClearTestEnvironment(testEnvironment)
				})
				It(`Create service client using external config and set url from constructor successfully`, func() {
					SetTestEnvironment(testEnvironment)
					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
						URL: "https://testService/api",
					})
					Expect(testService).ToNot(BeNil())
					Expect(testServiceErr).To(BeNil())
					Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
					ClearTestEnvironment(testEnvironment)
				})
				It(`Create service client using external config and set url programatically successfully`, func() {
					SetTestEnvironment(testEnvironment)
					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{})
					err := testService.SetServiceURL("https://testService/api")
					Expect(err).To(BeNil())
					Expect(testService).ToNot(BeNil())
					Expect(testServiceErr).To(BeNil())
					Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
					ClearTestEnvironment(testEnvironment)
				})
			})
			Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
				// Map containing environment variables used in testing.
				var testEnvironment = map[string]string{
					"DNS_SVCS_URL":       "https://dnssvcsv1/api",
					"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
				}

				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{})

				It(`Instantiate service client with error`, func() {
					Expect(testService).To(BeNil())
					Expect(testServiceErr).ToNot(BeNil())
					ClearTestEnvironment(testEnvironment)
				})
			})
			Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
				// Map containing environment variables used in testing.
				var testEnvironment = map[string]string{
					"DNS_SVCS_AUTH_TYPE": "NOAuth",
				}

				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "{{BAD_URL_STRING",
				})

				It(`Instantiate service client with error`, func() {
					Expect(testService).To(BeNil())
					Expect(testServiceErr).ToNot(BeNil())
					ClearTestEnvironment(testEnvironment)
				})
			})
		})
		Describe(`ListPools(listPoolsOptions *ListPoolsOptions)`, func() {
			listPoolsPath := "/instances/testString/pools"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listPoolsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"pools": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "minimum_origins": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "weight": 1}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_type": "email", "notification_channel": "xxx@mail.example.com"}], "offset": 1, "limit": 20, "count": 1, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/load_balancers?page=1&per_page=20"}, "next": {"href": "https://api.pdns.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/load_balancers?page=2&per_page=20"}}`)
				}))
				It(`Invoke ListPools successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.ListPools(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the ListPoolsOptions model
					listPoolsOptionsModel := new(dnssvcsv1.ListPoolsOptions)
					listPoolsOptionsModel.InstanceID = core.StringPtr("testString")
					listPoolsOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.ListPools(listPoolsOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`CreatePool(createPoolOptions *CreatePoolOptions)`, func() {
			createPoolPath := "/instances/testString/pools"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createPoolPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "minimum_origins": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "weight": 1}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_type": "email", "notification_channel": "xxx@mail.example.com"}`)
				}))
				It(`Invoke CreatePool successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.CreatePool(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the Origin model
					originModel := new(dnssvcsv1.Origin)
					originModel.Name = core.StringPtr("app-server-1")
					originModel.Description = core.StringPtr("description of the origin server")
					originModel.Address = core.StringPtr("10.10.16.8")
					originModel.Enabled = core.BoolPtr(true)
					originModel.Weight = core.Int64Ptr(int64(1))

					// Construct an instance of the CreatePoolOptions model
					createPoolOptionsModel := new(dnssvcsv1.CreatePoolOptions)
					createPoolOptionsModel.InstanceID = core.StringPtr("testString")
					createPoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
					createPoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
					createPoolOptionsModel.Enabled = core.BoolPtr(true)
					createPoolOptionsModel.MinimumOrigins = core.Int64Ptr(int64(1))
					createPoolOptionsModel.Origins = []dnssvcsv1.Origin{*originModel}
					createPoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
					createPoolOptionsModel.NotificationType = core.StringPtr("email")
					createPoolOptionsModel.NotificationChannel = core.StringPtr("xxx@mail.example.com")
					createPoolOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.CreatePool(createPoolOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`DeletePool(deletePoolOptions *DeletePoolOptions)`, func() {
			deletePoolPath := "/instances/testString/pools/testString"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deletePoolPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
				It(`Invoke DeletePool successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					response, operationErr := testService.DeletePool(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())

					// Construct an instance of the DeletePoolOptions model
					deletePoolOptionsModel := new(dnssvcsv1.DeletePoolOptions)
					deletePoolOptionsModel.InstanceID = core.StringPtr("testString")
					deletePoolOptionsModel.PoolID = core.StringPtr("testString")
					deletePoolOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					response, operationErr = testService.DeletePool(deletePoolOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
				})
			})
		})
		Describe(`GetPool(getPoolOptions *GetPoolOptions)`, func() {
			getPoolPath := "/instances/testString/pools/testString"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPoolPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "minimum_origins": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "weight": 1}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_type": "email", "notification_channel": "xxx@mail.example.com"}`)
				}))
				It(`Invoke GetPool successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.GetPool(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the GetPoolOptions model
					getPoolOptionsModel := new(dnssvcsv1.GetPoolOptions)
					getPoolOptionsModel.InstanceID = core.StringPtr("testString")
					getPoolOptionsModel.PoolID = core.StringPtr("testString")
					getPoolOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.GetPool(getPoolOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`UpdatePool(updatePoolOptions *UpdatePoolOptions)`, func() {
			updatePoolPath := "/instances/testString/pools/testString"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updatePoolPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "name": "dal10-az-pool", "description": "Load balancer pool for dal10 availability zone.", "enabled": true, "minimum_origins": 1, "origins": [{"name": "app-server-1", "description": "description of the origin server", "address": "10.10.16.8", "enabled": true, "weight": 1}], "monitor": "7dd6841c-264e-11ea-88df-062967242a6a", "notification_type": "email", "notification_channel": "xxx@mail.example.com"}`)
				}))
				It(`Invoke UpdatePool successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.UpdatePool(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the Origin model
					originModel := new(dnssvcsv1.Origin)
					originModel.Name = core.StringPtr("app-server-1")
					originModel.Description = core.StringPtr("description of the origin server")
					originModel.Address = core.StringPtr("10.10.16.8")
					originModel.Enabled = core.BoolPtr(true)
					originModel.Weight = core.Int64Ptr(int64(1))

					// Construct an instance of the UpdatePoolOptions model
					updatePoolOptionsModel := new(dnssvcsv1.UpdatePoolOptions)
					updatePoolOptionsModel.InstanceID = core.StringPtr("testString")
					updatePoolOptionsModel.PoolID = core.StringPtr("testString")
					updatePoolOptionsModel.Name = core.StringPtr("dal10-az-pool")
					updatePoolOptionsModel.Description = core.StringPtr("Load balancer pool for dal10 availability zone.")
					updatePoolOptionsModel.Enabled = core.BoolPtr(true)
					updatePoolOptionsModel.MinimumOrigins = core.Int64Ptr(int64(1))
					updatePoolOptionsModel.Origins = []dnssvcsv1.Origin{*originModel}
					updatePoolOptionsModel.Monitor = core.StringPtr("7dd6841c-264e-11ea-88df-062967242a6a")
					updatePoolOptionsModel.NotificationType = core.StringPtr("email")
					updatePoolOptionsModel.NotificationChannel = core.StringPtr("xxx@mail.example.com")
					updatePoolOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.UpdatePool(updatePoolOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`Service constructor tests`, func() {
			It(`Instantiate service client`, func() {
				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
			})
			It(`Instantiate service client with error: Invalid URL`, func() {
				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL: "{{BAD_URL_STRING",
				})
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
			})
			It(`Instantiate service client with error: Invalid Auth`, func() {
				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
					URL: "https://dnssvcsv1/api",
					Authenticator: &core.BasicAuthenticator{
						Username: "",
						Password: "",
					},
				})
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
			})
		})
		Describe(`Service constructor tests using external config`, func() {
			Context(`Using external config, construct service client instances`, func() {
				// Map containing environment variables used in testing.
				var testEnvironment = map[string]string{
					"DNS_SVCS_URL":       "https://dnssvcsv1/api",
					"DNS_SVCS_AUTH_TYPE": "noauth",
				}

				It(`Create service client using external config successfully`, func() {
					SetTestEnvironment(testEnvironment)
					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{})
					Expect(testService).ToNot(BeNil())
					Expect(testServiceErr).To(BeNil())
					ClearTestEnvironment(testEnvironment)
				})
				It(`Create service client using external config and set url from constructor successfully`, func() {
					SetTestEnvironment(testEnvironment)
					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
						URL: "https://testService/api",
					})
					Expect(testService).ToNot(BeNil())
					Expect(testServiceErr).To(BeNil())
					Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
					ClearTestEnvironment(testEnvironment)
				})
				It(`Create service client using external config and set url programatically successfully`, func() {
					SetTestEnvironment(testEnvironment)
					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{})
					err := testService.SetServiceURL("https://testService/api")
					Expect(err).To(BeNil())
					Expect(testService).ToNot(BeNil())
					Expect(testServiceErr).To(BeNil())
					Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
					ClearTestEnvironment(testEnvironment)
				})
			})
			Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
				// Map containing environment variables used in testing.
				var testEnvironment = map[string]string{
					"DNS_SVCS_URL":       "https://dnssvcsv1/api",
					"DNS_SVCS_AUTH_TYPE": "someOtherAuth",
				}

				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{})

				It(`Instantiate service client with error`, func() {
					Expect(testService).To(BeNil())
					Expect(testServiceErr).ToNot(BeNil())
					ClearTestEnvironment(testEnvironment)
				})
			})
			Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
				// Map containing environment variables used in testing.
				var testEnvironment = map[string]string{
					"DNS_SVCS_AUTH_TYPE": "NOAuth",
				}

				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(&dnssvcsv1.DnsSvcsV1Options{
					URL: "{{BAD_URL_STRING",
				})

				It(`Instantiate service client with error`, func() {
					Expect(testService).To(BeNil())
					Expect(testServiceErr).ToNot(BeNil())
					ClearTestEnvironment(testEnvironment)
				})
			})
		})
		Describe(`ListMonitors(listMonitorsOptions *ListMonitorsOptions)`, func() {
			listMonitorsPath := "/instances/testString/monitors"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listMonitorsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"monitors": [{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "header": {"anyKey": "anyValue"}, "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "follow_redirects": false}], "offset": 1, "limit": 20, "count": 1, "total_count": 200, "first": {"href": "https://api.dns-svcs.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/load_balancers?page=1&per_page=20"}, "next": {"href": "https://api.pdns.cloud.ibm.com/v1/instances/434f6c3e-6014-4124-a61d-2e910bca19b1/zones/example.com:d04d3a7a-7f6d-47d4-b811-08c5478fa1a4/load_balancers?page=2&per_page=20"}}`)
				}))
				It(`Invoke ListMonitors successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.ListMonitors(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the ListMonitorsOptions model
					listMonitorsOptionsModel := new(dnssvcsv1.ListMonitorsOptions)
					listMonitorsOptionsModel.InstanceID = core.StringPtr("testString")
					listMonitorsOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.ListMonitors(listMonitorsOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`CreateMonitor(createMonitorOptions *CreateMonitorOptions)`, func() {
			createMonitorPath := "/instances/testString/monitors"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createMonitorPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "header": {"anyKey": "anyValue"}, "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "follow_redirects": false}`)
				}))
				It(`Invoke CreateMonitor successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.CreateMonitor(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the CreateMonitorOptions model
					createMonitorOptionsModel := new(dnssvcsv1.CreateMonitorOptions)
					createMonitorOptionsModel.InstanceID = core.StringPtr("testString")
					createMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
					createMonitorOptionsModel.Type = core.StringPtr("HTTPS")
					createMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
					createMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
					createMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
					createMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
					createMonitorOptionsModel.Method = core.StringPtr("GET")
					createMonitorOptionsModel.Path = core.StringPtr("/health")
					createMonitorOptionsModel.Header = CreateMockMap()
					createMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
					createMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
					createMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
					createMonitorOptionsModel.FollowRedirects = core.BoolPtr(false)
					createMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.CreateMonitor(createMonitorOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`DeleteMonitor(deleteMonitorOptions *DeleteMonitorOptions)`, func() {
			deleteMonitorPath := "/instances/testString/monitors/testString"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteMonitorPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
				It(`Invoke DeleteMonitor successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					response, operationErr := testService.DeleteMonitor(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())

					// Construct an instance of the DeleteMonitorOptions model
					deleteMonitorOptionsModel := new(dnssvcsv1.DeleteMonitorOptions)
					deleteMonitorOptionsModel.InstanceID = core.StringPtr("testString")
					deleteMonitorOptionsModel.MonitorID = core.StringPtr("testString")
					deleteMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					response, operationErr = testService.DeleteMonitor(deleteMonitorOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
				})
			})
		})
		Describe(`GetMonitor(getMonitorOptions *GetMonitorOptions)`, func() {
			getMonitorPath := "/instances/testString/monitors/testString"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getMonitorPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "header": {"anyKey": "anyValue"}, "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "follow_redirects": false}`)
				}))
				It(`Invoke GetMonitor successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.GetMonitor(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the GetMonitorOptions model
					getMonitorOptionsModel := new(dnssvcsv1.GetMonitorOptions)
					getMonitorOptionsModel.InstanceID = core.StringPtr("testString")
					getMonitorOptionsModel.MonitorID = core.StringPtr("testString")
					getMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.GetMonitor(getMonitorOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
		Describe(`UpdateMonitor(updateMonitorOptions *UpdateMonitorOptions)`, func() {
			updateMonitorPath := "/instances/testString/monitors/testString"
			Context(`Using mock server endpoint`, func() {
				testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateMonitorPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "5365b73c-ce6f-4d6f-ad9f-d9c131b26370", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "description": "Load balancer monitor for glb.example.com.", "type": "HTTPS", "port": 8080, "interval": 60, "retries": 2, "timeout": 5, "method": "GET", "path": "/health", "header": {"anyKey": "anyValue"}, "allow_insecure": false, "expected_codes": "200", "expected_body": "alive", "follow_redirects": false}`)
				}))
				It(`Invoke UpdateMonitor successfully`, func() {
					defer testServer.Close()

					testService, testServiceErr := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
						URL:           testServer.URL,
						Authenticator: &core.NoAuthAuthenticator{},
					})
					Expect(testServiceErr).To(BeNil())
					Expect(testService).ToNot(BeNil())

					// Invoke operation with nil options model (negative test)
					result, response, operationErr := testService.UpdateMonitor(nil)
					Expect(operationErr).NotTo(BeNil())
					Expect(response).To(BeNil())
					Expect(result).To(BeNil())

					// Construct an instance of the UpdateMonitorOptions model
					updateMonitorOptionsModel := new(dnssvcsv1.UpdateMonitorOptions)
					updateMonitorOptionsModel.InstanceID = core.StringPtr("testString")
					updateMonitorOptionsModel.MonitorID = core.StringPtr("testString")
					updateMonitorOptionsModel.Description = core.StringPtr("Load balancer monitor for glb.example.com.")
					updateMonitorOptionsModel.Type = core.StringPtr("HTTPS")
					updateMonitorOptionsModel.Port = core.Int64Ptr(int64(8080))
					updateMonitorOptionsModel.Interval = core.Int64Ptr(int64(60))
					updateMonitorOptionsModel.Retries = core.Int64Ptr(int64(2))
					updateMonitorOptionsModel.Timeout = core.Int64Ptr(int64(5))
					updateMonitorOptionsModel.Method = core.StringPtr("GET")
					updateMonitorOptionsModel.Path = core.StringPtr("/health")
					updateMonitorOptionsModel.Header = CreateMockMap()
					updateMonitorOptionsModel.AllowInsecure = core.BoolPtr(false)
					updateMonitorOptionsModel.ExpectedCodes = core.StringPtr("200")
					updateMonitorOptionsModel.ExpectedBody = core.StringPtr("alive")
					updateMonitorOptionsModel.FollowRedirects = core.BoolPtr(false)
					updateMonitorOptionsModel.XCorrelationID = core.StringPtr("testString")

					// Invoke operation with valid options model (positive test)
					result, response, operationErr = testService.UpdateMonitor(updateMonitorOptionsModel)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				})
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a sample service client instance`, func() {
			testService, _ := dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
				URL:           "http://dnssvcsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewPermittedNetworkVpc successfully`, func() {
				vpcCrn := "crn:v1:bluemix:public:is:eu-de:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:6e6cc326-04d1-4c99-a289-efb3ae4193d6"
				model, err := testService.NewPermittedNetworkVpc(vpcCrn)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataARecord successfully`, func() {
				ip := "10.110.201.214"
				model, err := testService.NewResourceRecordInputRdataRdataARecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataAaaaRecord successfully`, func() {
				ip := "2019::2019"
				model, err := testService.NewResourceRecordInputRdataRdataAaaaRecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataCnameRecord successfully`, func() {
				cname := "www.example.com"
				model, err := testService.NewResourceRecordInputRdataRdataCnameRecord(cname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataMxRecord successfully`, func() {
				exchange := "mail.example.com"
				preference := int64(10)
				model, err := testService.NewResourceRecordInputRdataRdataMxRecord(exchange, preference)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataPtrRecord successfully`, func() {
				ptrdname := "www.example.com"
				model, err := testService.NewResourceRecordInputRdataRdataPtrRecord(ptrdname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataSrvRecord successfully`, func() {
				port := int64(80)
				priority := int64(10)
				target := "www.example.com"
				weight := int64(10)
				model, err := testService.NewResourceRecordInputRdataRdataSrvRecord(port, priority, target, weight)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordInputRdataRdataTxtRecord successfully`, func() {
				txtdata := "This is a text record"
				model, err := testService.NewResourceRecordInputRdataRdataTxtRecord(txtdata)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataARecord successfully`, func() {
				ip := "10.110.201.214"
				model, err := testService.NewResourceRecordUpdateInputRdataRdataARecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataAaaaRecord successfully`, func() {
				ip := "2019::2019"
				model, err := testService.NewResourceRecordUpdateInputRdataRdataAaaaRecord(ip)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataCnameRecord successfully`, func() {
				cname := "www.example.com"
				model, err := testService.NewResourceRecordUpdateInputRdataRdataCnameRecord(cname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataMxRecord successfully`, func() {
				exchange := "mail.example.com"
				preference := int64(10)
				model, err := testService.NewResourceRecordUpdateInputRdataRdataMxRecord(exchange, preference)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataPtrRecord successfully`, func() {
				ptrdname := "www.example.com"
				model, err := testService.NewResourceRecordUpdateInputRdataRdataPtrRecord(ptrdname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataSrvRecord successfully`, func() {
				port := int64(80)
				priority := int64(10)
				target := "www.example.com"
				weight := int64(10)
				model, err := testService.NewResourceRecordUpdateInputRdataRdataSrvRecord(port, priority, target, weight)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceRecordUpdateInputRdataRdataTxtRecord successfully`, func() {
				txtdata := "This is a text record"
				model, err := testService.NewResourceRecordUpdateInputRdataRdataTxtRecord(txtdata)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
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

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
