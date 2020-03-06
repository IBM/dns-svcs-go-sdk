package sample

/**
 * Copyright 2019 IBM All Rights Reserved.
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

import (
	"fmt"
	"os"

	"github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
)

// DNS Service Resource Records

/************************************** DNS Service Resouce Record Operations ****************************/

// listDnsResourceRecords - List resource records in DNS zone
func listDnsResourceRecords() {
	zoneID = os.Getenv("ZONE_ID")
	listResourceRecordOptions := dnsSvc.NewListResourceRecordsOptions(instanceID, zoneID)
	_, listResourceRecordResponse, reqErr := dnsSvc.ListResourceRecords(listResourceRecordOptions)
	if reqErr == nil {
		fmt.Println(listResourceRecordResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// createDnsResourceRecord -  Create a resource record in DNS zone
func createDnsResourceRecord() {
	zoneID = os.Getenv("ZONE_ID")
	createResourceRecordOptions := dnsSvc.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName("exmaple")
	createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_A)
	resourceRecordAData, _ := dnsSvc.NewResourceRecordInputRdataRdataARecord("1.1.1.1")
	createResourceRecordOptions.SetRdata(resourceRecordAData)
	_, createResourceRecordResponse, reqErr := dnsSvc.CreateResourceRecord(createResourceRecordOptions)
	if reqErr == nil {
		fmt.Println(createResourceRecordResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// getDnsResourceRecord - Get a resource record in DNS zone
func getDnsResourceRecord(resourceRecordID string) {
	zoneID = os.Getenv("ZONE_ID")
	getResourceRecordOptions := dnsSvc.NewGetResourceRecordOptions(instanceID, zoneID, resourceRecordID)
	_, getResourceRecordResponse, reqErr := dnsSvc.GetResourceRecord(getResourceRecordOptions)
	if reqErr == nil {
		fmt.Println(getResourceRecordResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// updateDnsResourceRecord - Update a resource record in DNS zone
func updateDnsResourceRecord(resourceRecordID string) {
	zoneID = os.Getenv("ZONE_ID")
	updateResourceRecordOptions := dnsSvc.NewUpdateResourceRecordOptions(instanceID, zoneID, resourceRecordID)
	updateresourceRecordAData, _ := dnsSvc.NewResourceRecordUpdateInputRdataRdataARecord("1.1.1.2")
	updateResourceRecordOptions.SetName("update-example")
	updateResourceRecordOptions.SetRdata(updateresourceRecordAData)
	_, updateResourceRecordResponse, reqErr := dnsSvc.UpdateResourceRecord(updateResourceRecordOptions)
	if reqErr == nil {
		fmt.Println(updateResourceRecordResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// deleteDnsResourceRecord - Delete a resource record in DNS zone
func deleteDnsResourceRecord(resourceRecordID string) {
	zoneID = os.Getenv("ZONE_ID")
	deleteResourceRecordOptions := dnsSvc.NewDeleteResourceRecordOptions(instanceID, zoneID, resourceRecordID)
	deleteResouceRecordResult, reqErr := dnsSvc.DeleteResourceRecord(deleteResourceRecordOptions)
	if reqErr == nil {
		fmt.Println(deleteResouceRecordResult.String())
	} else {
		fmt.Println(reqErr)
	}
}
