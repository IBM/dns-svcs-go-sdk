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

// DNS Service Permitted Networks
var zoneID string

/************************************** DNS Service Permitted Network Operations ****************************/

// listDnsPermittedNetworks - List permitted networks in DNS zone
func listDnsPermittedNetworks() {
	zoneID = os.Getenv("ZONE_ID")
	listPermittedNetworkOptions := dnsSvc.NewListPermittedNetworksOptions(instanceID, zoneID)
	_, listPermittedNetworkResponse, reqErr := dnsSvc.ListPermittedNetworks(listPermittedNetworkOptions)
	if reqErr == nil {
		fmt.Println(listPermittedNetworkResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// createDnsPermittedNetwork -- Create a permitted network in DNS zone
func createDnsPermittedNetwork() {
	zoneID = os.Getenv("ZONE_ID")
	vpcCrn := "vpc crn"
	createPermittedNetworkOptions := dnsSvc.NewCreatePermittedNetworkOptions(instanceID, zoneID)
	permittedNetworkCrn, err := dnsSvc.NewPermittedNetworkVpc(vpcCrn)
	if err == nil {
		createPermittedNetworkOptions.SetPermittedNetwork(permittedNetworkCrn)
		createPermittedNetworkOptions.SetType(dnssvcsv1.CreatePermittedNetworkOptions_Type_Vpc)
		_, createPermittedNetworkResponse, reqErr := dnsSvc.CreatePermittedNetwork(createPermittedNetworkOptions)
		if reqErr == nil {
			fmt.Println(createPermittedNetworkResponse.String())
		} else {
			fmt.Println(reqErr)
		}
	}
}

// getDnsPermittedNetwork -- Get a permitted network in DNS zone
func getDnsPermittedNetwork(permittedNetworkID string) {
	zoneID = os.Getenv("ZONE_ID")
	getPermittedNetworkOptions := dnsSvc.NewGetPermittedNetworkOptions(instanceID, zoneID, permittedNetworkID)
	_, getPermittedNetworkResponse, reqErr := dnsSvc.GetPermittedNetwork(getPermittedNetworkOptions)
	if reqErr == nil {
		fmt.Println(getPermittedNetworkResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// deleteDnsPermittedNetwork - Delete a permitted network in DNS zone
func deleteDnsPermittedNetwork(permittedNetworkID string) {
	zoneID = os.Getenv("ZONE_ID")
	deletePermittedNetworkOptions := dnsSvc.NewDeletePermittedNetworkOptions(instanceID, zoneID, permittedNetworkID)
	_, deletePermittedNetworkResp, reqErr := dnsSvc.DeletePermittedNetwork(deletePermittedNetworkOptions)
	if reqErr == nil {
		fmt.Println(deletePermittedNetworkResp.String())
	} else {
		fmt.Println(reqErr)
	}
}
