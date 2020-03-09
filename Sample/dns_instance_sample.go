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

	"github.com/IBM/dns-svcs-go-sdk/dnssvcsinstancesv2"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/joho/godotenv"
)

// DNS Service Instance
var dnsInsService *dnssvcsinstancesv2.DnsSvcsInstancesV2
var dnsInsServiceErr error

func main() {

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

	// Create the IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
	}

	if err == nil {
		dnsInsService, dnsInsServiceErr = dnssvcsinstancesv2.NewDnsSvcsInstancesV2(
			&dnssvcsinstancesv2.DnsSvcsInstancesV2Options{
				Authenticator: authenticator, // Get APIKey from your .env file
			})

		if dnsInsServiceErr == nil {
			customHeaders := http.Header{}
			customHeaders.Add("Content-type", "application/json")
			dnsInsService.Service.SetDefaultHeaders(customHeaders)
		}
	}

	// Call any of the functions below to use them.

	// listAllDnsInstance()
	// createNewDnsInstance(os.Getenv("name"), os.Getenv("target"))
	// getDnsInstance()
	// updateDnsInstance(os.Getenv("name"), os.Getenv("target"))
	// deleteDnsInstance()

}

/************************************** DNS Service Instance Operations ****************************/

// listAllDnsInstance - Gets the list of all DNS instances
func listAllDnsInstance() {
	listDnsInstancesOptions := dnsInsService.NewListResourceInstancesOptions("dns service resource ID", dnssvcsinstancesv2.ListResourceInstancesOptions_Type_ServiceInstance)

	// Set neccessary parameters
	listDnsInstancesOptions.SetGuid("")

	listDnsInstancesOptions.SetHeaders(map[string]string{})

	_, detailedResponse, reqErr := dnsInsService.ListResourceInstances(listDnsInstancesOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// createNewDnsInstance - Creates a new DNS instance
func createNewDnsInstance(name string, target string) {
	// ex name := "test-integration-cos-instance"
	// target refers to the depoyment location
	// ex target := "bluemix-global"

	resourceGroup := os.Getenv("RESOURCE_GROUP")
	resourcePlanID := os.Getenv("RESOURCE_PLAN_ID")

	dnsInstanceOptions := dnsInsService.NewCreateResourceInstanceOptions(name, target, resourceGroup, resourcePlanID)
	dnsInstanceOptions.SetTags([]string{"integration-test"})
	dnsInstanceOptions.SetParameters(map[string]interface{}{})
	dnsInstanceOptions.SetHeaders(map[string]string{"Content-Type": "application/json"})

	_, detailedResponse, reqErr := dnsInsService.CreateResourceInstance(dnsInstanceOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// getDnsInstance - Gets DNS instance based on ID
func getDnsInstance() {
	instanceID := os.Getenv("INSTANCE_ID")
	getDnsInstanceByIDOptions := dnsInsService.NewGetResourceInstanceOptions(instanceID)
	getDnsInstanceByIDOptions.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	_, detailedResponse, reqErr := dnsInsService.GetResourceInstance(getDnsInstanceByIDOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// updateDnsInstance - Updates DNS instance based on ID
func updateDnsInstance(name string, target string) {
	// target refers to the depoyment location
	// ex target := "bluemix-global"
	instanceID := os.Getenv("INSTANCE_ID")
	resourcePlanID := os.Getenv("RESOURCE_PLAN_ID")

	updateDnsInstaceOptions := dnsInsService.NewUpdateResourceInstanceOptions(instanceID)
	updateDnsInstaceOptions.SetName(name)
	updateDnsInstaceOptions.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	updateDnsInstaceOptions.SetResourcePlanID(resourcePlanID)
	updateDnsInstaceOptions.SetParameters(map[string]interface{}{})

	_, detailedResponse, reqErr := dnsInsService.UpdateResourceInstance(updateDnsInstaceOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// deleteDnsInstance - delete DNS instance based on ID
func deleteDnsInstance() {
	instanceID := os.Getenv("INSTANCE_ID")
	deleteDnsInstanceOptions := dnsInsService.NewDeleteResourceInstanceOptions(instanceID)

	detailedResponse, reqErr := dnsInsService.DeleteResourceInstance(deleteDnsInstanceOptions)

	if reqErr == nil {
		fmt.Println(detailedResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}
