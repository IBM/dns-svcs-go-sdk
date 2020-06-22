// +build integration

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
	"log"
	"os"
	"strings"
	"testing"

	"github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var service *dnssvcsv1.DnsSvcsV1
var serviceErr error
var instanceID string
var zoneID string
var zoneName string

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create the authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
	}

	service, serviceErr = dnssvcsv1.NewDnsSvcsV1(
		&dnssvcsv1.DnsSvcsV1Options{
			Authenticator: authenticator,
		})

	// Retrieive dns service instance ID and zone ID
	instanceID = os.Getenv("INSTANCE_ID")
	if len(instanceID) == 0 {
		log.Fatal("No instanceID set")
	}
	zoneID = os.Getenv("ZONE_ID")
	if len(instanceID) == 0 {
		log.Fatal("No zoneID set")
	}
	parts := strings.SplitN(zoneID, ":", 2)
	zoneName = parts[0]
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}

func TestDnsZonesOperation(t *testing.T) {
	shouldSkipTest(t)

	header := map[string]string{
		"test": "teststring",
	}
	// Test List DNS Zone
	_, _, returnValueErr := service.ListDnszones(nil)
	assert.NotNil(t, returnValueErr)

	listDnszonesOptions := service.NewListDnszonesOptions(instanceID)
	listDnszonesOptions.SetXCorrelationID("abc123")
	listDnszonesOptions.SetHeaders(header)
	listDnszonesOptions.SetOffset(int64(0))
	listDnszonesOptions.SetLimit(int64(10))
	results, response, reqErr := service.ListDnszones(listDnszonesOptions)
	assert.NotNil(t, results)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Nil(t, reqErr)
	firstResource := results.Dnszones[0]
	assert.Equal(t, instanceID, *firstResource.InstanceID)

	// Test Create DNS Zone
	zoneName := "test.com"
	createDnszoneOptions := service.NewCreateDnszoneOptions(instanceID, zoneName)
	createDnszoneOptions.SetDescription("testString")
	createDnszoneOptions.SetLabel("testString")
	createDnszoneOptions.SetXCorrelationID("abc123")
	createDnszoneOptions.SetHeaders(header)
	result, response, reqErr := service.CreateDnszone(createDnszoneOptions)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "testString", *result.Description)
	assert.Nil(t, reqErr)

	testzoneID := result.ID
	// Test Get DNS Zone
	getDnszoneOptions := service.NewGetDnszoneOptions(instanceID, *testzoneID)
	getDnszoneOptions.SetXCorrelationID("abc123")
	getDnszoneOptions.SetHeaders(header)
	result, response, reqErr = service.GetDnszone(getDnszoneOptions)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "testString", *result.Description)
	assert.Nil(t, reqErr)

	// Test Update DNS Zone
	updateDnszoneOptions := service.NewUpdateDnszoneOptions(instanceID, *testzoneID)
	updateDnszoneOptions.SetDescription("testUpdate")
	updateDnszoneOptions.SetLabel("testUpdate")
	updateDnszoneOptions.SetXCorrelationID("abc123")
	updateDnszoneOptions.SetHeaders(header)
	result, response, reqErr = service.UpdateDnszone(updateDnszoneOptions)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "testUpdate", *result.Description)
	assert.Nil(t, reqErr)

	// Test Delete DNS Zone
	deleteDnszoneOptions := service.NewDeleteDnszoneOptions(instanceID, *testzoneID)
	deleteDnszoneOptions.SetXCorrelationID("abc123")
	deleteDnszoneOptions.SetHeaders(header)
	response, reqErr = service.DeleteDnszone(deleteDnszoneOptions)
	assert.Nil(t, reqErr)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())

	// Test Delete DNS Zone fail
	fdeleteDnszoneOptions := new(dnssvcsv1.DeleteDnszoneOptions)
	fdeleteDnszoneOptions.SetInstanceID(instanceID)
	fdeleteDnszoneOptions.SetDnszoneID("invalid_id")
	_, reqErr = service.DeleteDnszone(fdeleteDnszoneOptions)
	assert.NotNil(t, reqErr)
}

func TestDnsPermittedNetworkOperation(t *testing.T) {
	shouldSkipTest(t)

	vpcCrn := os.Getenv("VPC_CRN")
	assert.NotEmpty(t, vpcCrn)

	header := map[string]string{
		"test": "teststring",
	}
	// Test Add Permitted Network
	createPermittedNetworkOptions := service.NewCreatePermittedNetworkOptions(instanceID, zoneID)
	permittedNetworkCrn := &dnssvcsv1.PermittedNetworkVpc{
		VpcCrn: &vpcCrn,
	}
	createPermittedNetworkOptions.SetPermittedNetwork(permittedNetworkCrn)
	createPermittedNetworkOptions.SetType(dnssvcsv1.CreatePermittedNetworkOptions_Type_Vpc)
	createPermittedNetworkOptions.SetXCorrelationID("abc123")
	createPermittedNetworkOptions.SetHeaders(header)
	result, response, reqErr := service.CreatePermittedNetwork(createPermittedNetworkOptions)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "ACTIVE", *result.State)
	assert.Nil(t, reqErr)

	permittednetworkID := result.ID

	// Test List Permitted Networks
	listPermittedNetworksOptions := service.NewListPermittedNetworksOptions(instanceID, zoneID)
	listPermittedNetworksOptions.SetXCorrelationID("abc123")
	listPermittedNetworksOptions.SetHeaders(header)
	results, response, reqErr := service.ListPermittedNetworks(listPermittedNetworksOptions)
	assert.NotNil(t, results)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Nil(t, reqErr)
	firstResource := results.PermittedNetworks[0]
	assert.NotNil(t, *firstResource.ID)

	// Test Get Permitted Network
	getPermittedNetworkOptions := service.NewGetPermittedNetworkOptions(instanceID, zoneID, *permittednetworkID)
	getPermittedNetworkOptions.SetXCorrelationID("abc123")
	getPermittedNetworkOptions.SetHeaders(header)
	result, response, reqErr = service.GetPermittedNetwork(getPermittedNetworkOptions)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "ACTIVE", *result.State)
	assert.Nil(t, reqErr)

	// Test Get Permitted Network Fail
	fgetPermittedNetworkOptions := new(dnssvcsv1.GetPermittedNetworkOptions)
	fgetPermittedNetworkOptions.SetInstanceID(instanceID)
	fgetPermittedNetworkOptions.SetDnszoneID(zoneID)
	fgetPermittedNetworkOptions.SetPermittedNetworkID("invalid_id")
	_, _, reqErr = service.GetPermittedNetwork(fgetPermittedNetworkOptions)
	assert.NotNil(t, reqErr)

	// Test Remove Permitted Network
	deletePermittedNetworkOptions := service.NewDeletePermittedNetworkOptions(instanceID, zoneID, *permittednetworkID)
	deletePermittedNetworkOptions.SetXCorrelationID("abc123")
	deletePermittedNetworkOptions.SetHeaders(header)
	result, response, reqErr = service.DeletePermittedNetwork(deletePermittedNetworkOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 202, response.GetStatusCode())
	assert.Equal(t, "REMOVAL_IN_PROGRESS", *result.State)
	assert.Nil(t, reqErr)

	// Test Rmove Permitted Network Fail
	fdeletePermittedNetworkOptions := new(dnssvcsv1.DeletePermittedNetworkOptions)
	fdeletePermittedNetworkOptions.SetInstanceID(instanceID)
	fdeletePermittedNetworkOptions.SetDnszoneID(zoneID)
	fdeletePermittedNetworkOptions.SetPermittedNetworkID("invalid_id")
	_, _, reqErr = service.DeletePermittedNetwork(fdeletePermittedNetworkOptions)
	assert.NotNil(t, reqErr)
}

func TestDnsResourceRecordsOperation(t *testing.T) {
	shouldSkipTest(t)

	header := map[string]string{
		"test": "teststring",
	}
	// Test List Resource Records
	listResourceRecordsOptions := service.NewListResourceRecordsOptions(instanceID, zoneID)
	listResourceRecordsOptions.SetXCorrelationID("abc123")
	listResourceRecordsOptions.SetHeaders(header)
	results, response, reqErr := service.ListResourceRecords(listResourceRecordsOptions)
	assert.NotNil(t, results)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Nil(t, reqErr)
	firstResource := results.ResourceRecords[0]
	assert.NotNil(t, *firstResource.ID)

	// Test List Resource Records Fail
	flistResourceRecordsOptions := new(dnssvcsv1.ListResourceRecordsOptions)
	flistResourceRecordsOptions.SetInstanceID(instanceID)
	flistResourceRecordsOptions.SetDnszoneID("invaid_id")
	flistResourceRecordsOptions.SetOffset(int64(1))
	flistResourceRecordsOptions.SetLimit(int64(20))
	_, _, reqErr = service.ListResourceRecords(flistResourceRecordsOptions)
	assert.NotNil(t, reqErr)
}

func TestDnsResourceRecordAandPTROperation(t *testing.T) {
	shouldSkipTest(t)

	header := map[string]string{
		"test": "teststring",
	}
	// Test Create Resource Record A
	createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName("testa")
	createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_A)
	createResourceRecordOptions.SetTTL(120)
	rdataARecord, err := service.NewResourceRecordInputRdataRdataARecord("1.1.1.1")
	assert.Nil(t, err)
	createResourceRecordOptions.SetRdata(rdataARecord)
	createResourceRecordOptions.SetXCorrelationID("abc123")
	createResourceRecordOptions.SetHeaders(header)
	result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "A", *result.Type)
	assert.Nil(t, reqErr)

	// Test Create Resource Record A PTR
	createResourceRecordPtrOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordPtrOptions.SetName("1.1.1.1")
	createResourceRecordPtrOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Ptr)
	createResourceRecordPtrOptions.SetTTL(120)
	rdataPtrRecord, err := service.NewResourceRecordInputRdataRdataPtrRecord("testa." + zoneName)
	assert.Nil(t, err)
	createResourceRecordPtrOptions.SetRdata(rdataPtrRecord)
	createResourceRecordPtrOptions.SetXCorrelationID("abc123")
	createResourceRecordPtrOptions.SetHeaders(header)
	ptrresult, response, reqErr := service.CreateResourceRecord(createResourceRecordPtrOptions)
	assert.NotNil(t, ptrresult)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "PTR", *ptrresult.Type)
	assert.Nil(t, reqErr)

	aRecordID := result.ID
	ptrRecordID := ptrresult.ID
	// Test Get Resource Record A
	getResourceRecordOptions := service.NewGetResourceRecordOptions(instanceID, zoneID, *aRecordID)
	getResourceRecordOptions.SetXCorrelationID("abc123")
	getResourceRecordOptions.SetHeaders(header)
	result, response, reqErr = service.GetResourceRecord(getResourceRecordOptions)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "A", *result.Type)
	assert.Nil(t, reqErr)

	// Test Get Resource Record PTR
	getResourceRecordPtrOptions := new(dnssvcsv1.GetResourceRecordOptions)
	getResourceRecordPtrOptions.SetInstanceID(instanceID)
	getResourceRecordPtrOptions.SetDnszoneID(zoneID)
	getResourceRecordPtrOptions.SetRecordID(*ptrRecordID)
	ptrresult, response, reqErr = service.GetResourceRecord(getResourceRecordPtrOptions)
	assert.NotNil(t, ptrresult)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "PTR", *ptrresult.Type)
	assert.Nil(t, reqErr)

	// Test Update Resource Record A
	updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, *aRecordID)
	updateResourceRecordOptions.SetName("updatea")
	updateResourceRecordOptions.SetTTL(300)
	updateResourceRecordOptions.SetXCorrelationID("abc123")
	updateResourceRecordOptions.SetHeaders(header)
	updaterdataARecord, err := service.NewResourceRecordUpdateInputRdataRdataARecord("1.1.1.2")
	assert.Nil(t, err)
	updateResourceRecordOptions.SetRdata(updaterdataARecord)
	result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "A", *result.Type)
	assert.Nil(t, reqErr)

	// Test Update Resource Record PTR
	updateResourceRecordPtrOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, *ptrRecordID)
	updateResourceRecordPtrOptions.SetTTL(300)
	updateResourceRecordPtrOptions.SetXCorrelationID("abc123")
	updateResourceRecordPtrOptions.SetHeaders(header)
	ptrresult, response, reqErr = service.UpdateResourceRecord(updateResourceRecordPtrOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "PTR", *ptrresult.Type)
	assert.Nil(t, reqErr)

	// Test Delete Resource Record PTR
	deleteResourceRecordOptions := &dnssvcsv1.DeleteResourceRecordOptions{}
	deleteResourceRecordOptions.SetInstanceID(instanceID)
	deleteResourceRecordOptions.SetDnszoneID(zoneID)
	deleteResourceRecordOptions.SetRecordID(*ptrRecordID)
	deleteResourceRecordOptions.SetXCorrelationID("abc123")
	deleteResourceRecordOptions.SetHeaders(header)
	response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
	assert.Nil(t, reqErr)

	// Test Delete Resource Record A
	deleteResourceRecordOptions = service.NewDeleteResourceRecordOptions(instanceID, zoneID, *aRecordID)
	deleteResourceRecordOptions.SetXCorrelationID("abc123")
	deleteResourceRecordOptions.SetHeaders(header)
	response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
	assert.Nil(t, reqErr)
}

func TestDnsResourceRecordAAAAOperation(t *testing.T) {
	shouldSkipTest(t)

	header := map[string]string{
		"test": "teststring",
	}
	// Test Create Resource Record AAAA
	createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName("testaaaa")
	createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Aaaa)
	createResourceRecordOptions.SetTTL(120)
	rdataAaaaRecord, err := service.NewResourceRecordInputRdataRdataAaaaRecord("2001::8888")
	assert.Nil(t, err)
	createResourceRecordOptions.SetRdata(rdataAaaaRecord)
	createResourceRecordOptions.SetXCorrelationID("abc123")
	createResourceRecordOptions.SetHeaders(header)
	result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "AAAA", *result.Type)
	assert.Equal(t, "testaaaa."+zoneName, *result.Name)
	assert.Nil(t, reqErr)

	aaaaRecordID := result.ID
	// Test Update Resource Record AAAA
	updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, *aaaaRecordID)
	updateResourceRecordOptions.SetName("updateaaaa")
	updateResourceRecordOptions.SetTTL(300)
	updaterdataAaaaRecord, err := service.NewResourceRecordUpdateInputRdataRdataAaaaRecord("2001::8889")
	assert.Nil(t, err)
	updateResourceRecordOptions.SetRdata(updaterdataAaaaRecord)
	updateResourceRecordOptions.SetXCorrelationID("abc123")
	updateResourceRecordOptions.SetHeaders(header)
	result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "AAAA", *result.Type)
	assert.Equal(t, "updateaaaa."+zoneName, *result.Name)
	assert.Nil(t, reqErr)

	// Test Delete Resource Record AAAA
	deleteResourceRecordOptions := service.NewDeleteResourceRecordOptions(instanceID, zoneID, *aaaaRecordID)
	deleteResourceRecordOptions.SetXCorrelationID("abc123")
	deleteResourceRecordOptions.SetHeaders(header)
	response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
	assert.Nil(t, reqErr)
}

func TestDnsResourceRecordCNAMEOperation(t *testing.T) {
	shouldSkipTest(t)

	header := map[string]string{
		"test": "teststring",
	}
	// Test Create Resource Record CNAME
	createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName("testcname")
	createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Cname)
	createResourceRecordOptions.SetTTL(120)
	rdataCnameRecord, err := service.NewResourceRecordInputRdataRdataCnameRecord("testcnamedata.com")
	assert.Nil(t, err)
	createResourceRecordOptions.SetRdata(rdataCnameRecord)
	createResourceRecordOptions.SetXCorrelationID("abc123")
	createResourceRecordOptions.SetHeaders(header)
	result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "CNAME", *result.Type)
	assert.Equal(t, "testcname."+zoneName, *result.Name)
	assert.Nil(t, reqErr)

	cnameRecordID := result.ID
	// Test Update Resource Record CNAME
	updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, *cnameRecordID)
	updateResourceRecordOptions.SetName("updatecname")
	updateResourceRecordOptions.SetTTL(300)
	updaterdataCnameRecord, err := service.NewResourceRecordUpdateInputRdataRdataCnameRecord("updatecnamedata.com")
	assert.Nil(t, err)
	updateResourceRecordOptions.SetRdata(updaterdataCnameRecord)
	updateResourceRecordOptions.SetXCorrelationID("abc123")
	updateResourceRecordOptions.SetHeaders(header)
	result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "CNAME", *result.Type)
	assert.Equal(t, "updatecname."+zoneName, *result.Name)
	assert.Nil(t, reqErr)

	// Test Delete Resource Record CNAME
	deleteResourceRecordOptions := service.NewDeleteResourceRecordOptions(instanceID, zoneID, *cnameRecordID)
	deleteResourceRecordOptions.SetXCorrelationID("abc123")
	deleteResourceRecordOptions.SetHeaders(header)
	response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
	assert.Nil(t, reqErr)
}

func TestDnsResourceRecordMXOperation(t *testing.T) {
	shouldSkipTest(t)

	header := map[string]string{
		"test": "teststring",
	}
	// Test Create Resource Record MX
	createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName("testmx")
	createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Mx)
	createResourceRecordOptions.SetTTL(120)
	rdataMxRecord, err := service.NewResourceRecordInputRdataRdataMxRecord("mail.testmx.com", 1)
	assert.Nil(t, err)
	createResourceRecordOptions.SetRdata(rdataMxRecord)
	createResourceRecordOptions.SetXCorrelationID("abc123")
	createResourceRecordOptions.SetHeaders(header)
	result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "MX", *result.Type)
	assert.Equal(t, "testmx."+zoneName, *result.Name)
	assert.Nil(t, reqErr)

	mxRecordID := result.ID
	// Test Update Resource Record MX
	updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, *mxRecordID)
	updateResourceRecordOptions.SetName("testupdatemx")
	updateResourceRecordOptions.SetTTL(300)
	updaterdataMxRecord, err := service.NewResourceRecordUpdateInputRdataRdataMxRecord("mail1.testmx.com", 2)
	assert.Nil(t, err)
	updateResourceRecordOptions.SetRdata(updaterdataMxRecord)
	updateResourceRecordOptions.SetXCorrelationID("abc123")
	updateResourceRecordOptions.SetHeaders(header)
	result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "MX", *result.Type)
	assert.Equal(t, "testupdatemx."+zoneName, *result.Name)
	assert.Nil(t, reqErr)

	// Test Delete Resource Record MX
	deleteResourceRecordOptions := service.NewDeleteResourceRecordOptions(instanceID, zoneID, *mxRecordID)
	deleteResourceRecordOptions.SetXCorrelationID("abc123")
	deleteResourceRecordOptions.SetHeaders(header)
	response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
	assert.Nil(t, reqErr)
}

func TestDnsResourceRecordSRVOperation(t *testing.T) {
	shouldSkipTest(t)

	header := map[string]string{
		"test": "teststring",
	}
	// Test Create Resource Record SRV
	createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName("testsrv")
	createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Srv)
	createResourceRecordOptions.SetTTL(120)
	createResourceRecordOptions.SetService("_sip")
	createResourceRecordOptions.SetProtocol("udp")
	rdataSrvRecord, err := service.NewResourceRecordInputRdataRdataSrvRecord(1, 1, "siphost.com", 1)
	assert.Nil(t, err)
	createResourceRecordOptions.SetRdata(rdataSrvRecord)
	createResourceRecordOptions.SetXCorrelationID("abc123")
	createResourceRecordOptions.SetHeaders(header)
	result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "SRV", *result.Type)
	assert.Equal(t, "udp", *result.Protocol)
	assert.Nil(t, reqErr)

	srvRecordID := result.ID
	// Test Update Resource Record SRV
	updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, *srvRecordID)
	updateResourceRecordOptions.SetName("updatesrv")
	updateResourceRecordOptions.SetTTL(300)
	updateResourceRecordOptions.SetService("_sip")
	updateResourceRecordOptions.SetProtocol("udp")
	updaterdataSrvRecord, err := service.NewResourceRecordUpdateInputRdataRdataSrvRecord(2, 2, "updatesiphost.com", 2)
	assert.Nil(t, err)
	updateResourceRecordOptions.SetRdata(updaterdataSrvRecord)
	updateResourceRecordOptions.SetXCorrelationID("abc123")
	updateResourceRecordOptions.SetHeaders(header)
	result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "SRV", *result.Type)
	assert.Nil(t, reqErr)

	// Test Delete Resource Record SRV
	deleteResourceRecordOptions := service.NewDeleteResourceRecordOptions(instanceID, zoneID, *srvRecordID)
	deleteResourceRecordOptions.SetXCorrelationID("abc123")
	deleteResourceRecordOptions.SetHeaders(header)
	response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
	assert.Nil(t, reqErr)
}

func TestDnsResourceRecordTXTOperation(t *testing.T) {
	shouldSkipTest(t)

	header := map[string]string{
		"test": "teststring",
	}
	// Test Create Resource Record TXT
	createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName("testtxt")
	createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Txt)
	createResourceRecordOptions.SetTTL(120)
	rdataTxtRecord, err := service.NewResourceRecordInputRdataRdataTxtRecord("txtdata string")
	assert.Nil(t, err)
	createResourceRecordOptions.SetRdata(rdataTxtRecord)
	createResourceRecordOptions.SetXCorrelationID("abc123")
	createResourceRecordOptions.SetHeaders(header)
	result, response, reqErr := service.CreateResourceRecord(createResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "TXT", *result.Type)
	assert.Equal(t, "testtxt."+zoneName, *result.Name)
	assert.Nil(t, reqErr)

	txtRecordID := result.ID
	// Test Update Resource Record TXT
	updateResourceRecordOptions := &dnssvcsv1.UpdateResourceRecordOptions{}
	updateResourceRecordOptions.SetInstanceID(instanceID)
	updateResourceRecordOptions.SetDnszoneID(zoneID)
	updateResourceRecordOptions.SetRecordID(*txtRecordID)
	updateResourceRecordOptions.SetTTL(300)
	updateResourceRecordOptions.SetName("updatetxt")
	updaterdataTxtRecord, err := service.NewResourceRecordUpdateInputRdataRdataTxtRecord("update txtdata string")
	assert.Nil(t, err)
	updateResourceRecordOptions.SetRdata(updaterdataTxtRecord)
	updateResourceRecordOptions.SetXCorrelationID("abc123")
	updateResourceRecordOptions.SetHeaders(header)
	result, response, reqErr = service.UpdateResourceRecord(updateResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "TXT", *result.Type)
	assert.Equal(t, "updatetxt."+zoneName, *result.Name)
	assert.Nil(t, reqErr)

	// Test Delete Resource Record TXT
	deleteResourceRecordOptions := service.NewDeleteResourceRecordOptions(instanceID, zoneID, *txtRecordID)
	deleteResourceRecordOptions.SetXCorrelationID("abc123")
	deleteResourceRecordOptions.SetHeaders(header)
	response, reqErr = service.DeleteResourceRecord(deleteResourceRecordOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
	assert.Nil(t, reqErr)
}

func TestDnsLoadBalancerOperation(t *testing.T) {
	shouldSkipTest(t)

	headers := map[string]string{
		"test": "teststring",
	}
	healthcheckHeader1 := dnssvcsv1.HealthcheckHeader{
		Name:  core.StringPtr("Host"),
		Value: []string{"testexample.com"},
	}
	healthcheckHeader2 := dnssvcsv1.HealthcheckHeader{
		Name:  core.StringPtr("AppId"),
		Value: []string{"abc123"},
	}
	var healthcheckHeader []dnssvcsv1.HealthcheckHeader
	healthcheckHeader = append(healthcheckHeader, healthcheckHeader1, healthcheckHeader2)
	// Test Create GLB Monitor
	createDnsGlbMonitorOptions := service.NewCreateMonitorOptions(instanceID)
	createDnsGlbMonitorOptions.SetName("glbMonitor")
	createDnsGlbMonitorOptions.SetDescription("Load balancer monitor for example.com")
	createDnsGlbMonitorOptions.SetType("HTTPS")
	createDnsGlbMonitorOptions.SetPort(int64(8080))
	createDnsGlbMonitorOptions.SetInterval(int64(60))
	createDnsGlbMonitorOptions.SetRetries(int64(2))
	createDnsGlbMonitorOptions.SetTimeout(int64(5))
	createDnsGlbMonitorOptions.SetMethod("GET")
	createDnsGlbMonitorOptions.SetPath("/health")
	createDnsGlbMonitorOptions.SetAllowInsecure(false)
	createDnsGlbMonitorOptions.SetExpectedCodes("200")
	createDnsGlbMonitorOptions.SetExpectedBody("alive")
	createDnsGlbMonitorOptions.SetHeadersVar(healthcheckHeader)
	createDnsGlbMonitorOptions.SetFollowRedirects(false)
	createDnsGlbMonitorOptions.SetXCorrelationID("abc123")
	createDnsGlbMonitorOptions.SetHeaders(headers)
	monitorResult, response, reqErr := service.CreateMonitor(createDnsGlbMonitorOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "HTTPS", *monitorResult.Type)
	assert.Equal(t, "GET", *monitorResult.Method)
	assert.Nil(t, reqErr)

	monitorID := monitorResult.ID
	// Test List GLB Monitor
	listDnsGlbMonitorOptions := service.NewListMonitorsOptions(instanceID)
	listDnsGlbMonitorOptions.SetXCorrelationID("abc123")
	listDnsGlbMonitorOptions.SetHeaders(headers)
	monitorResults, response, reqErr := service.ListMonitors(listDnsGlbMonitorOptions)
	assert.NotNil(t, monitorResults)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Nil(t, reqErr)
	firstMonitor := monitorResults.Monitors[0]
	assert.NotNil(t, *firstMonitor.ID)

	// Test Get GLB Monitor
	getDnsGlbMonitorOptions := service.NewGetMonitorOptions(instanceID, *monitorID)
	getDnsGlbMonitorOptions.SetXCorrelationID("abc123")
	getDnsGlbMonitorOptions.SetHeaders(headers)
	monitorResult, response, reqErr = service.GetMonitor(getDnsGlbMonitorOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "HTTPS", *monitorResult.Type)
	assert.Equal(t, "GET", *monitorResult.Method)
	assert.Nil(t, reqErr)

	// Test Update GLB Monitor
	updateDnsGlbMonitorOptions := service.NewUpdateMonitorOptions(instanceID, *monitorID)
	updateDnsGlbMonitorOptions.SetName("glbMonitorUpdate")
	updateDnsGlbMonitorOptions.SetDescription("Update Load balancer monitor for example.com")
	updateDnsGlbMonitorOptions.SetType("HTTP")
	updateDnsGlbMonitorOptions.SetPort(int64(8080))
	updateDnsGlbMonitorOptions.SetInterval(int64(60))
	updateDnsGlbMonitorOptions.SetRetries(int64(2))
	updateDnsGlbMonitorOptions.SetTimeout(int64(5))
	updateDnsGlbMonitorOptions.SetMethod("GET")
	updateDnsGlbMonitorOptions.SetPath("/health")
	updateDnsGlbMonitorOptions.SetAllowInsecure(false)
	updateDnsGlbMonitorOptions.SetExpectedCodes("200")
	updateDnsGlbMonitorOptions.SetExpectedBody("alive")
	updateDnsGlbMonitorOptions.SetFollowRedirects(false)
	updateDnsGlbMonitorOptions.SetXCorrelationID("abc123")
	updateDnsGlbMonitorOptions.SetHeaders(headers)
	monitorResult, response, reqErr = service.UpdateMonitor(updateDnsGlbMonitorOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "HTTP", *monitorResult.Type)
	assert.Equal(t, "GET", *monitorResult.Method)
	assert.Nil(t, reqErr)

	// Test Create GLB Pool
	createDnsGlbPoolOptions := service.NewCreatePoolOptions(instanceID)
	createDnsGlbPoolOptions.SetName("dal-pool")
	createDnsGlbPoolOptions.SetDescription("dallas pool for example.com")
	createDnsGlbPoolOptions.SetEnabled(true)
	createDnsGlbPoolOptions.SetHealthyOriginsThreshold(int64(1))
	origin1 := new(dnssvcsv1.Origin)
	origin1.Name = core.StringPtr("dal-origin01")
	origin1.Description = core.StringPtr("description of the origin server")
	origin1.Address = core.StringPtr("10.10.16.8")
	origin1.Enabled = core.BoolPtr(true)
	origin1.Weight = core.Int64Ptr(int64(1))
	createDnsGlbPoolOptions.SetOrigins([]dnssvcsv1.Origin{*origin1})
	createDnsGlbPoolOptions.SetMonitor(*monitorID)
	createDnsGlbPoolOptions.SetNotificationType(dnssvcsv1.CreatePoolOptions_NotificationType_Webhook)
	createDnsGlbPoolOptions.SetNotificationChannel("https://mywebsite.com/dns/webhook")
	createDnsGlbPoolOptions.SetXCorrelationID("abc123")
	createDnsGlbPoolOptions.SetHeaders(headers)
	poolResult, response, reqErr := service.CreatePool(createDnsGlbPoolOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "dal-pool", *poolResult.Name)
	assert.NotNil(t, *poolResult.ID)
	assert.Nil(t, reqErr)

	poolID := poolResult.ID
	// Test List GLB Pool
	listDnsGlbPoolOptions := service.NewListPoolsOptions(instanceID)
	listDnsGlbPoolOptions.SetXCorrelationID("abc123")
	listDnsGlbPoolOptions.SetHeaders(headers)
	poolResults, response, reqErr := service.ListPools(listDnsGlbPoolOptions)
	assert.NotNil(t, poolResults)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Nil(t, reqErr)
	firstPool := poolResults.Pools[0]
	assert.NotNil(t, *firstPool.ID)

	// Test Get GLB Pool
	getDnsGlbPoolOptions := service.NewGetPoolOptions(instanceID, *poolID)
	getDnsGlbPoolOptions.SetXCorrelationID("abc123")
	getDnsGlbPoolOptions.SetHeaders(headers)
	poolResult, response, reqErr = service.GetPool(getDnsGlbPoolOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "dal-pool", *poolResult.Name)
	assert.NotNil(t, *poolResult.ID)
	assert.Nil(t, reqErr)

	// Test Update GLB Pool
	updateDnsGlbPoolOptions := service.NewUpdatePoolOptions(instanceID, *poolID)
	updateDnsGlbPoolOptions.SetName("dal-pool-update")
	updateDnsGlbPoolOptions.SetDescription("dallas pool update for example.com")
	updateDnsGlbPoolOptions.SetEnabled(true)
	updateDnsGlbPoolOptions.SetHealthyOriginsThreshold(int64(1))
	origin2 := new(dnssvcsv1.Origin)
	origin2.Name = core.StringPtr("dal-origin02")
	origin2.Description = core.StringPtr("description of the origin server")
	origin2.Address = core.StringPtr("10.10.16.9")
	origin2.Enabled = core.BoolPtr(true)
	origin2.Weight = core.Int64Ptr(int64(1))
	updateDnsGlbPoolOptions.SetOrigins([]dnssvcsv1.Origin{*origin2})
	updateDnsGlbPoolOptions.SetMonitor(*monitorID)
	updateDnsGlbPoolOptions.SetNotificationType(dnssvcsv1.CreatePoolOptions_NotificationType_Webhook)
	updateDnsGlbPoolOptions.SetNotificationChannel("https://mywebsite.com/dns/webhookupdate")
	updateDnsGlbPoolOptions.SetXCorrelationID("abc123")
	updateDnsGlbPoolOptions.SetHeaders(headers)
	poolResult, response, reqErr = service.UpdatePool(updateDnsGlbPoolOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "dal-pool-update", *poolResult.Name)
	assert.NotNil(t, *poolResult.ID)
	assert.Nil(t, reqErr)

	// Test Create GLB
	createDnsGlbOptions := service.NewCreateLoadBalancerOptions(instanceID, zoneID)
	createDnsGlbOptions.SetName("glbtest")
	createDnsGlbOptions.SetDescription("Global load balancer 01")
	createDnsGlbOptions.SetEnabled(true)
	createDnsGlbOptions.SetTTL(int64(300))
	createDnsGlbOptions.SetDefaultPools([]string{*poolID})
	createDnsGlbOptions.SetFallbackPool(*poolID)
	azPoolsItem := new(dnssvcsv1.LoadBalancerAzPoolsItem)
	azPoolsItem.AvailabilityZone = core.StringPtr("us-south-1")
	azPoolsItem.Pools = []string{*poolID}

	createDnsGlbOptions.SetAzPools([]dnssvcsv1.LoadBalancerAzPoolsItem{*azPoolsItem})
	createDnsGlbOptions.SetXCorrelationID("abc123")
	createDnsGlbOptions.SetHeaders(headers)
	glbResult, response, reqErr := service.CreateLoadBalancer(createDnsGlbOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "glbtest."+zoneName, *glbResult.Name)
	assert.NotNil(t, *glbResult.ID)
	assert.Nil(t, reqErr)

	glbID := glbResult.ID
	// Test List GLB
	listDnsGlbOptions := service.NewListLoadBalancersOptions(instanceID, zoneID)
	listDnsGlbOptions.SetXCorrelationID("abc123")
	listDnsGlbOptions.SetHeaders(headers)
	glbResults, response, reqErr := service.ListLoadBalancers(listDnsGlbOptions)
	assert.NotNil(t, glbResults)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Nil(t, reqErr)
	firstGlb := glbResults.LoadBalancers[0]
	assert.NotNil(t, *firstGlb.ID)

	// Test Get GLB
	getDnsGlbOptions := service.NewGetLoadBalancerOptions(instanceID, zoneID, *glbID)
	getDnsGlbOptions.SetXCorrelationID("abc123")
	getDnsGlbOptions.SetHeaders(headers)
	glbResult, response, reqErr = service.GetLoadBalancer(getDnsGlbOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "glbtest."+zoneName, *glbResult.Name)
	assert.NotNil(t, *glbResult.ID)
	assert.Nil(t, reqErr)

	// Test Update GLB
	updateDnsGlbOptions := service.NewUpdateLoadBalancerOptions(instanceID, zoneID, *glbID)
	updateDnsGlbOptions.SetName("updateglbtest")
	updateDnsGlbOptions.SetDescription("Update Global load balancer 01")
	updateDnsGlbOptions.SetEnabled(true)
	updateDnsGlbOptions.SetTTL(int64(300))
	updateDnsGlbOptions.SetDefaultPools([]string{*poolID})
	updateDnsGlbOptions.SetFallbackPool(*poolID)
	updateazPoolsItem := new(dnssvcsv1.LoadBalancerAzPoolsItem)
	updateazPoolsItem.AvailabilityZone = core.StringPtr("us-south-2")
	updateazPoolsItem.Pools = []string{*poolID}

	updateDnsGlbOptions.SetAzPools([]dnssvcsv1.LoadBalancerAzPoolsItem{*updateazPoolsItem})
	updateDnsGlbOptions.SetXCorrelationID("abc123")
	updateDnsGlbOptions.SetHeaders(headers)
	glbResult, response, reqErr = service.UpdateLoadBalancer(updateDnsGlbOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.GetStatusCode())
	assert.Equal(t, "updateglbtest."+zoneName, *glbResult.Name)
	assert.NotNil(t, *glbResult.ID)
	assert.Nil(t, reqErr)

	// Test Delete GLB
	deleteDnsGlbOptions := service.NewDeleteLoadBalancerOptions(instanceID, zoneID, *glbID)
	deleteDnsGlbOptions.SetXCorrelationID("abc123")
	deleteDnsGlbOptions.SetHeaders(headers)
	response, reqErr = service.DeleteLoadBalancer(deleteDnsGlbOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
	assert.Nil(t, reqErr)

	// Test Delete GLB Fail
	fdeleteDnsGlbOptions := new(dnssvcsv1.DeleteLoadBalancerOptions)
	fdeleteDnsGlbOptions.SetInstanceID(instanceID)
	fdeleteDnsGlbOptions.SetDnszoneID(zoneID)
	fdeleteDnsGlbOptions.SetLbID("invalid")
	_, reqErr = service.DeleteLoadBalancer(fdeleteDnsGlbOptions)
	assert.NotNil(t, reqErr)

	// Test Delete GLB Pool
	deleteDnsGlbPoolOptions := service.NewDeletePoolOptions(instanceID, *poolID)
	deleteDnsGlbPoolOptions.SetXCorrelationID("abc123")
	deleteDnsGlbPoolOptions.SetHeaders(headers)
	response, reqErr = service.DeletePool(deleteDnsGlbPoolOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
	assert.Nil(t, reqErr)

	// Test Delete GLB Monitor
	deleteDnsGlbMonitorOptions := service.NewDeleteMonitorOptions(instanceID, *monitorID)
	deleteDnsGlbMonitorOptions.SetXCorrelationID("abc123")
	deleteDnsGlbMonitorOptions.SetHeaders(headers)
	response, reqErr = service.DeleteMonitor(deleteDnsGlbMonitorOptions)
	assert.NotNil(t, response)
	assert.Equal(t, 204, response.GetStatusCode())
	assert.Nil(t, reqErr)
}
