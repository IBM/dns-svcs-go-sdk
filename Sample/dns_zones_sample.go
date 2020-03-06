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
	"log"
	"net/http"
	"os"

	"github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/joho/godotenv"
)

// DNS Service Zones
var dnsSvc *dnssvcsv1.DnsSvcsV1
var dnsSvcErr error
var instanceID string

func init() {

	// Initialization

	// Load environment variables (should be stored in .env file)
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/*
		// Create the IAM token authenticator
		token := os.Getenv("TOKEN")

		authenticator := &core.BearerTokenAuthenticator{
			BearerToken: token,
		}
	*/
	instanceID = os.Getenv("INSTANCE_ID")
	// Create the IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
	}

	if err == nil {
		dnsSvc, dnsSvcErr = dnssvcsv1.NewDnsSvcsV1(
			&dnssvcsv1.DnsSvcsV1Options{
				Authenticator: authenticator, // Get APIKey from your .env file
			})

		if dnsSvcErr == nil {
			customHeaders := http.Header{}
			customHeaders.Add("Content-type", "application/json")
			dnsSvc.Service.SetDefaultHeaders(customHeaders)
		}
	}
}

/************************************** DNS Service Zone Operations ****************************/

// listDnsZones - List zones based on DNS Instance ID
func listDnsZones() {
	listZonesOptions := dnsSvc.NewListDnszonesOptions(instanceID)
	_, detailedResponse, reqErr := dnsSvc.ListDnszones(listZonesOptions)
	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// createDnsZone - Create a Zone on DNS service Instance
func createDnsZone(zoneName string) {
	createZoneOptions := dnsSvc.NewCreateDnszoneOptions(instanceID, zoneName)
	createZoneOptions.SetDescription("zone description")
	createZoneOptions.SetLabel("zone label")
	_, createZoneResponse, reqErr := dnsSvc.CreateDnszone(createZoneOptions)
	if reqErr == nil {
		fmt.Println(createZoneResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// getDnsZone - Get a Zone on DNS service Instance
func getDnsZone(zoneID string) {
	getZoneOptions := dnsSvc.NewGetDnszoneOptions(instanceID, zoneID)
	_, getZoneResponse, reqErr := dnsSvc.GetDnszone(getZoneOptions)
	if reqErr == nil {
		fmt.Println(getZoneResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// updateDnsZone - Update a Zone on DNS service Instance
func updateDnsZone(zoneID string) {
	updateZoneOptions := dnsSvc.NewUpdateDnszoneOptions(instanceID, zoneID)
	updateZoneOptions.SetDescription("update description")
	updateZoneOptions.SetLabel("update label")

	_, updateZoneResponse, reqErr := dnsSvc.UpdateDnszone(updateZoneOptions)
	if reqErr == nil {
		fmt.Println(updateZoneResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// deleteDnsZone - Delete a Zone on DNS service Instance
func deleteDnsZone(zoneID string) {
	deleteZoneOptions := dnsSvc.NewDeleteDnszoneOptions(instanceID, zoneID)
	deleteZoneResponse, reqErr := dnsSvc.DeleteDnszone(deleteZoneOptions)
	if reqErr == nil {
		fmt.Println(deleteZoneResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}
