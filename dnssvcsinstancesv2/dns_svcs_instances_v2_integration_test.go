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

package dnssvcsinstancesv2_test

import (
	"os"

	"github.com/IBM/dns-svcs-go-sdk/dnssvcsinstancesv2"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`DnsSvcsInstancesV2`, func() {
	err := godotenv.Load("../.env")
	It(`Successfully loading .env file`, func() {
		Expect(err).To(BeNil())
	})

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
	}
	options := &dnssvcsinstancesv2.DnsSvcsInstancesV2Options{
		ServiceName:   "DnsSvcsInstancesV2_Mokcing",
		Authenticator: authenticator,
	}
	service, err := dnssvcsinstancesv2.NewDnsSvcsInstancesV2UsingExternalConfig(options)
	It(`Successfully created DnsSvcsInstancesV2 service instance`, func() {
		Expect(err).To(BeNil())
	})

	Describe(`ListResourceInstances(listResourceInstancesOptions *ListResourceInstancesOptions)`, func() {
		Context(`Successfully list resource by resourceID and type`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			resourceID := os.Getenv("RESOURCE_ID")
			resourceType := dnssvcsinstancesv2.ListResourceInstancesOptions_Type_ServiceInstance
			listResourceInstancesOptions := service.NewListResourceInstancesOptions(resourceID, resourceType).
				SetLimit("2").
				SetHeaders(header)

			It(`Successfully list all resources`, func() {
				result, detailedResponse, err := service.ListResourceInstances(listResourceInstancesOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				firstResource := result.Resources[0]
				Expect(*firstResource.ResourceID).To(Equal(resourceID))
				Expect(*firstResource.Name).ToNot(BeNil())

				secondResource := result.Resources[1]
				Expect(*secondResource.ResourceID).To(Equal(resourceID))
				Expect(*secondResource.Name).ToNot(BeNil())
			})
		})

		Context(`Failed to list resource by resourceID and type`, func() {
			resourceID := os.Getenv("RESOURCE_ID")
			resourceType := dnssvcsinstancesv2.ListResourceInstancesOptions_Type_ServiceInstance
			listResourceInstancesOptions := &dnssvcsinstancesv2.ListResourceInstancesOptions{}
			listResourceInstancesOptions.SetResourceID(resourceID)
			listResourceInstancesOptions.SetType(resourceType)
			listResourceInstancesOptions.SetName("testString")
			listResourceInstancesOptions.SetGuid("testString")
			listResourceInstancesOptions.SetResourceGroupID("testString")
			listResourceInstancesOptions.SetResourcePlanID("testString")
			listResourceInstancesOptions.SetUpdatedFrom("testString")
			listResourceInstancesOptions.SetUpdatedTo("testString")
			It(`Failed to list all resouces`, func() {
				_, _, err := service.ListResourceInstances(listResourceInstancesOptions)
				Expect(err).Should(HaveOccurred())
			})
		})

	})

	Describe(`CreateResourceInstance(createResourceInstanceOptions *CreateResourceInstanceOptions)`, func() {
		Context(`Successfully create resource instance`, func() {
			name := "To Kill a Mockingbird"
			target := os.Getenv("TARGET")
			resourceGroup := os.Getenv("RESOURCE_GROUP")
			resourcePlanID := os.Getenv("RESOURCE_PLAN_ID")
			createResourceInstanceOptions := service.NewCreateResourceInstanceOptions(name, target, resourceGroup, resourcePlanID)

			It(`Successfully create new resource`, func() {
				result, detailedResponse, err := service.CreateResourceInstance(createResourceInstanceOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(201))
				Expect(*result.Name).To(Equal("To Kill a Mockingbird"))
				Expect(*result.ResourceGroupID).To(Equal(resourceGroup))
			})
		})

		Context(`Fail to create resource instance`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			createResourceInstanceOptions := &dnssvcsinstancesv2.CreateResourceInstanceOptions{}
			createResourceInstanceOptions.SetAllowCleanup(false)
			createResourceInstanceOptions.SetName("testString")
			createResourceInstanceOptions.SetResourceGroup("testString")
			createResourceInstanceOptions.SetResourcePlanID("testString")
			createResourceInstanceOptions.SetTarget("testString")
			createResourceInstanceOptions.SetHeaders(header)
			It(`Fail to create new resource`, func() {
				result, detailedResponse, err := service.CreateResourceInstance(createResourceInstanceOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(err).Should(HaveOccurred())
			})
		})

	})

	Describe(`UpdateResourceInstance(updateResourceInstanceOptions *UpdateResourceInstanceOptions)`, func() {
		Context(`Successfully update resource by instanceID`, func() {
			instanceID := os.Getenv("INSTANCE_ID")
			instanceName := "To Update a Mockingbird"
			updateResourceInstanceOptions := service.NewUpdateResourceInstanceOptions(instanceID).
				SetName(instanceName).
				SetAllowCleanup(true)

			It(`Successfully update resource by instanceID`, func() {
				result, detailedResponse, err := service.UpdateResourceInstance(updateResourceInstanceOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.Guid).To(Equal(instanceID))
			})
		})

		Context(`Failed to update resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			instanceName := "To Update a Mockingbird"
			updateResourceInstanceOptions := &dnssvcsinstancesv2.UpdateResourceInstanceOptions{}
			updateResourceInstanceOptions.SetID(badinstanceID)
			updateResourceInstanceOptions.SetName(instanceName)
			updateResourceInstanceOptions.SetResourcePlanID("testString")
			updateResourceInstanceOptions.SetHeaders(header)
			It(`Failed to update resource by instanceID`, func() {
				result, detailedResponse, err := service.UpdateResourceInstance(updateResourceInstanceOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`GetResourceInstance(getResourceInstanceOptions *GetResourceInstanceOptions)`, func() {
		Context(`Successfully get resource by instanceID`, func() {
			instanceID := os.Getenv("INSTANCE_ID")
			getResourceInstanceOptions := service.NewGetResourceInstanceOptions(instanceID)
			It(`Successfully get resource by instanceID`, func() {
				result, detailedResponse, err := service.GetResourceInstance(getResourceInstanceOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.Guid).To(Equal(instanceID))
			})
		})

		Context(`Failed to get resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			getResourceInstanceOptions := &dnssvcsinstancesv2.GetResourceInstanceOptions{}
			getResourceInstanceOptions.SetID(badinstanceID)
			getResourceInstanceOptions.SetHeaders(header)
			It(`Failed to get resource by instanceID`, func() {
				result, detailedResponse, err := service.GetResourceInstance(getResourceInstanceOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`DeleteResourceInstance(deleteResourceInstanceOptions *DeleteResourceInstanceOptions)`, func() {
		Context(`Successfully delete resource by instanceID`, func() {
			instanceID := os.Getenv("INSTANCE_ID")
			deleteResourceInstanceOptions := service.NewDeleteResourceInstanceOptions(instanceID)
			It(`Successfully delete resource by instanceID`, func() {
				detailedResponse, err := service.DeleteResourceInstance(deleteResourceInstanceOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})

		Context(`Failed to delete resource by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "111"
			deleteResourceInstanceOptions := &dnssvcsinstancesv2.DeleteResourceInstanceOptions{}
			deleteResourceInstanceOptions.SetID(badinstanceID)
			deleteResourceInstanceOptions.SetHeaders(header)
			It(`Failed to delete resource by instanceID`, func() {
				detailedResponse, err := service.DeleteResourceInstance(deleteResourceInstanceOptions)
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
