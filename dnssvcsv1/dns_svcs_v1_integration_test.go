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
	"os"
	"strings"

	"github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`DnsSvcsV1`, func() {

	err := godotenv.Load("../.env")
	It(`Successfully loading .env file`, func() {
		Expect(err).To(BeNil())
	})

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("IAMAPIKEY"),
	}
	options := &dnssvcsv1.DnsSvcsV1Options{
		ServiceName:   "DnsSvcsV1_Mokcing",
		Authenticator: authenticator,
	}
	service, err := dnssvcsv1.NewDnsSvcsV1UsingExternalConfig(options)
	It(`Successfully created DnsSvcsV1 service instance`, func() {
		Expect(err).To(BeNil())
	})
	err = service.SetServiceURL(dnssvcsv1.DefaultServiceURL)
	It(`Successfully set DnsSvcsV1 service URL`, func() {
		Expect(err).To(BeNil())
	})
	instanceID := os.Getenv("INSTANCE_ID")
	It(`Successfully Get instance ID`, func() {
		Expect(instanceID).NotTo(BeEmpty())
	})
	zoneID := os.Getenv("ZONE_ID")
	It(`Successfully Get dns zone ID`, func() {
		Expect(instanceID).NotTo(BeEmpty())
	})

	Describe(`List Dns Zones`, func() {
		Context(`Successfully list dns zones by instanceID`, func() {
			listDnszonesOptions := service.NewListDnszonesOptions(instanceID)
			It(`Successfully list all dns zones`, func() {
				result, detailedResponse, err := service.ListDnszones(listDnszonesOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				firstResource := result.Dnszones[0]
				Expect(*firstResource.InstanceID).To(Equal(instanceID))
			})
		})

		Context(`Failed to list dns zones by instanceID`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badinstanceID := "bad_bird"
			listDnszonesOptions := &dnssvcsv1.ListDnszonesOptions{}
			listDnszonesOptions.SetInstanceID(badinstanceID)
			listDnszonesOptions.SetHeaders(header)
			It(`Failed to list all dns zones`, func() {
				_, _, err := service.ListDnszones(listDnszonesOptions)
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`Create/Update/Get/Delete a DNS zone`, func() {
		Context(`Successfully create/get/update/delete dns zone`, func() {
			var (
				zoneID string
			)
			zoneName := "test.com"
			createDnszoneOptions := service.NewCreateDnszoneOptions(instanceID, zoneName)
			zoneresult, detailedResponse, err := service.CreateDnszone(createDnszoneOptions)
			It(`Successfully create new dns zone`, func() {
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*zoneresult.Name).To(Equal(zoneName))
				Expect(*zoneresult.InstanceID).To(Equal(instanceID))
				Expect(*zoneresult.ID).NotTo(BeNil())
			})

			BeforeEach(func() {
				zoneID = *zoneresult.ID
			})

			It(`Successfully update dns zone`, func() {
				updateDnszoneOptions := service.NewUpdateDnszoneOptions(instanceID, zoneID).
					SetDescription("testUpdate")
				result, detailedResponse, err := service.UpdateDnszone(updateDnszoneOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.InstanceID).To(Equal(instanceID))
				Expect(*result.ID).To(Equal(zoneID))
			})

			It(`Successfully get dns zone`, func() {
				getDnszoneOptions := service.NewGetDnszoneOptions(instanceID, zoneID)
				result, detailedResponse, err := service.GetDnszone(getDnszoneOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(zoneID))
				Expect(*result.Description).To(Equal("testUpdate"))
			})

			It(`Successfully delete dns zone`, func() {
				deleteDnszoneOptions := service.NewDeleteDnszoneOptions(instanceID, zoneID)
				detailedResponse, err := service.DeleteDnszone(deleteDnszoneOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			})
		})

		Context(`Fail to create dns zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			createDnszoneOptions := &dnssvcsv1.CreateDnszoneOptions{}
			createDnszoneOptions.SetInstanceID(instanceID)
			createDnszoneOptions.SetName("ibm.com")
			createDnszoneOptions.SetDescription("testString")
			createDnszoneOptions.SetLabel("testString")
			createDnszoneOptions.SetXCorrelationID("testString")
			createDnszoneOptions.SetHeaders(header)
			It(`Fail to create dns zone`, func() {
				result, detailedResponse, err := service.CreateDnszone(createDnszoneOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(err).Should(HaveOccurred())
			})
		})

		Context(`Failed to update dns zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badZoneID := "111"
			updateDnszoneOptions := &dnssvcsv1.UpdateDnszoneOptions{}
			updateDnszoneOptions.SetInstanceID(instanceID)
			updateDnszoneOptions.SetDnszoneID(badZoneID)
			updateDnszoneOptions.SetDescription("testString")
			updateDnszoneOptions.SetLabel("testString")
			updateDnszoneOptions.SetXCorrelationID("testString")
			updateDnszoneOptions.SetHeaders(header)
			It(`Failed to update dns zone`, func() {
				result, detailedResponse, err := service.UpdateDnszone(updateDnszoneOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(err).Should(HaveOccurred())
			})
		})

		Context(`Failed to get dns zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badZoneID := "111"
			getDnszoneOptions := &dnssvcsv1.GetDnszoneOptions{}
			getDnszoneOptions.SetInstanceID(instanceID)
			getDnszoneOptions.SetDnszoneID(badZoneID)
			getDnszoneOptions.SetXCorrelationID("testString")
			getDnszoneOptions.SetHeaders(header)
			It(`Failed to get dns zone`, func() {
				result, detailedResponse, err := service.GetDnszone(getDnszoneOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(err).Should(HaveOccurred())
			})
		})

		Context(`Failed to delete dns zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badZoneID := "111"
			deleteDnszoneOptions := &dnssvcsv1.DeleteDnszoneOptions{}
			deleteDnszoneOptions.SetInstanceID(instanceID)
			deleteDnszoneOptions.SetDnszoneID(badZoneID)
			deleteDnszoneOptions.SetXCorrelationID("testString")
			deleteDnszoneOptions.SetHeaders(header)
			It(`Failed to delete dns zone`, func() {
				detailedResponse, err := service.DeleteDnszone(deleteDnszoneOptions)
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`Add a permitted network`, func() {
		Context(`Successfully create permitted network on zone`, func() {
			zoneID := os.Getenv("ZONE_ID")
			vpcCrn := os.Getenv("VPC_CRN")
			permittednetworkID := os.Getenv("VPC_ID")
			createPermittedNetworkOptions := service.NewCreatePermittedNetworkOptions(instanceID, zoneID)
			permittedNetworkCrn := &dnssvcsv1.PermittedNetworkVpc{
				VpcCrn: &vpcCrn,
			}
			createPermittedNetworkOptions.SetPermittedNetwork(permittedNetworkCrn)
			createPermittedNetworkOptions.SetType(dnssvcsv1.CreatePermittedNetworkOptions_Type_Vpc)
			It(`Successfully create new dns zone`, func() {
				result, detailedResponse, err := service.CreatePermittedNetwork(createPermittedNetworkOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.Type).To(Equal(dnssvcsv1.CreatePermittedNetworkOptions_Type_Vpc))
				Expect(*result.State).To(Equal("ACTIVE"))
				Expect(*result.ID).To(Equal(permittednetworkID))
			})
		})

		Context(`Fail to create permitted network on zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			zoneID := os.Getenv("ZONE_ID")
			vpcCrn := "bad_crn"
			createPermittedNetworkOptions := &dnssvcsv1.CreatePermittedNetworkOptions{}
			createPermittedNetworkOptions.SetInstanceID(instanceID)
			createPermittedNetworkOptions.SetDnszoneID(zoneID)
			permittedNetworkCrn := &dnssvcsv1.PermittedNetworkVpc{
				VpcCrn: &vpcCrn,
			}
			createPermittedNetworkOptions.SetPermittedNetwork(permittedNetworkCrn)
			createPermittedNetworkOptions.SetType(dnssvcsv1.CreatePermittedNetworkOptions_Type_Vpc)
			createPermittedNetworkOptions.SetXCorrelationID("testString")
			createPermittedNetworkOptions.SetHeaders(header)
			It(`Fail to create new resource`, func() {
				result, detailedResponse, err := service.CreatePermittedNetwork(createPermittedNetworkOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`List permitted networks`, func() {
		Context(`Successfully list permitted network on zone`, func() {
			zoneID := os.Getenv("ZONE_ID")
			listPermittedNetworksOptions := service.NewListPermittedNetworksOptions(instanceID, zoneID)
			It(`Successfully list all resources`, func() {
				result, detailedResponse, err := service.ListPermittedNetworks(listPermittedNetworksOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				firstResource := result.PermittedNetworks[0]
				Expect(*firstResource.ID).NotTo(BeNil())
			})
		})

		Context(`Failed to list permitted network on zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			badzoneID := "bad_bird"
			listPermittedNetworksOptions := &dnssvcsv1.ListPermittedNetworksOptions{}
			listPermittedNetworksOptions.SetInstanceID(instanceID)
			listPermittedNetworksOptions.SetDnszoneID(badzoneID)
			listPermittedNetworksOptions.SetXCorrelationID("testString")
			listPermittedNetworksOptions.SetHeaders(header)
			It(`Failed to list all resouces`, func() {
				_, _, err := service.ListPermittedNetworks(listPermittedNetworksOptions)
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`Get a permitted network`, func() {
		Context(`Successfully get permitted network on zone`, func() {
			zoneID := os.Getenv("ZONE_ID")
			permittednetworkID := os.Getenv("VPC_ID")
			getPermittedNetworkOptions := service.NewGetPermittedNetworkOptions(instanceID, zoneID, permittednetworkID)
			It(`Successfully get permitted network on zone`, func() {
				result, detailedResponse, err := service.GetPermittedNetwork(getPermittedNetworkOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))

				Expect(*result.ID).To(Equal(permittednetworkID))
				Expect(*result.State).To(Equal("ACTIVE"))
			})
		})

		Context(`Failed to get permitted network on zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			zoneID := os.Getenv("ZONE_ID")
			permittednetworkID := "bad_bird"
			getPermittedNetworkOptions := &dnssvcsv1.GetPermittedNetworkOptions{}
			getPermittedNetworkOptions.SetInstanceID(instanceID)
			getPermittedNetworkOptions.SetDnszoneID(zoneID)
			getPermittedNetworkOptions.SetPermittedNetworkID(permittednetworkID)
			getPermittedNetworkOptions.SetXCorrelationID("testString")
			getPermittedNetworkOptions.SetHeaders(header)
			It(`Failed to get dns zone`, func() {
				result, detailedResponse, err := service.GetPermittedNetwork(getPermittedNetworkOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`Remove a permited network`, func() {
		Context(`Successfully remove permitted network on zone`, func() {
			zoneID := os.Getenv("ZONE_ID")
			permittednetworkID := os.Getenv("VPC_ID")
			deletePermittedNetworkOptions := service.NewDeletePermittedNetworkOptions(instanceID, zoneID, permittednetworkID)
			It(`Successfully remove permitted network on zone`, func() {
				result, detailedResponse, err := service.DeletePermittedNetwork(deletePermittedNetworkOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(202))

				Expect(*result.ID).To(Equal(permittednetworkID))
			})
		})

		Context(`Failed to remove permitted network on zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			zoneID := os.Getenv("ZONE_ID")
			permittednetworkID := "bad_bird"
			deletePermittedNetworkOptions := &dnssvcsv1.DeletePermittedNetworkOptions{}
			deletePermittedNetworkOptions.SetInstanceID(instanceID)
			deletePermittedNetworkOptions.SetDnszoneID(zoneID)
			deletePermittedNetworkOptions.SetPermittedNetworkID(permittednetworkID)
			deletePermittedNetworkOptions.SetXCorrelationID("testString")
			deletePermittedNetworkOptions.SetHeaders(header)
			It(`Failed to get dns zone`, func() {
				result, detailedResponse, err := service.DeletePermittedNetwork(deletePermittedNetworkOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(404))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`Resource Record A`, func() {
		Context(`Successfully create "A" record on zone`, func() {
			createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
			createResourceRecordOptions.SetName("testa")
			createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_A)
			createResourceRecordOptions.SetTTL(120)
			rdataARecord, err := service.NewResourceRecordInputRdataRdataARecord("1.1.1.1")
			It(`Successfully set create "A" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			createResourceRecordOptions.SetRdata(rdataARecord)
			rcaresult, detailedResponse, err := service.CreateResourceRecord(createResourceRecordOptions)
			It(`Successfully set create "A" record on zone`, func() {
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*rcaresult.Type).To(Equal(dnssvcsv1.CreateResourceRecordOptions_Type_A))
				Expect(*rcaresult.ID).NotTo(BeNil())

			})
		})

		Context(`Successfully update "A" record on zone`, func() {
			racID := os.Getenv("ARECORD")
			It(`Successfully Get A record ID`, func() {
				Expect(racID).NotTo(BeEmpty())
			})
			updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, racID)
			updateResourceRecordOptions.SetName("updatea")
			updateResourceRecordOptions.SetTTL(300)
			updaterdataARecord, err := service.NewResourceRecordUpdateInputRdataRdataARecord("1.1.1.2")
			It(`Successfully set update "A" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			updateResourceRecordOptions.SetRdata(updaterdataARecord)

			It(`Successfully update "A" record on zone`, func() {
				result, detailedResponse, err := service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(racID))
			})
		})

		Context(`Successfully get "A" record on zone`, func() {
			racID := os.Getenv("ARECORD")
			It(`Successfully Get A record ID`, func() {
				Expect(racID).NotTo(BeEmpty())
			})
			getResourceRecordOptions := service.NewGetResourceRecordOptions(instanceID, zoneID, racID)
			It(`Successfully get "A" record on zone`, func() {
				result, detailedResponse, err := service.GetResourceRecord(getResourceRecordOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(racID))
			})
		})

		Context(`Fail to create "A" record on zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			createResourceRecordOptions := &dnssvcsv1.CreateResourceRecordOptions{}
			createResourceRecordOptions.SetInstanceID(instanceID)
			createResourceRecordOptions.SetDnszoneID(zoneID)
			createResourceRecordOptions.SetName("testbad")
			createResourceRecordOptions.SetXCorrelationID("teststring")
			createResourceRecordOptions.SetHeaders(header)
			createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_A)
			createResourceRecordOptions.SetTTL(120)
			rdataARecord, err := service.NewResourceRecordInputRdataRdataARecord("1.1.1")
			It(`Set create "A" record Rdata for fail`, func() {
				Expect(err).To(BeNil())
			})
			createResourceRecordOptions.SetRdata(rdataARecord)
			It(`Fail to set create "A" record on zone`, func() {
				rcaresult, detailedResponse, err := service.CreateResourceRecord(createResourceRecordOptions)
				Expect(rcaresult).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(err).Should(HaveOccurred())
			})
		})

		Context(`Fail to get "A" record on zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			racID := "bad_bird"
			getResourceRecordOptions := &dnssvcsv1.GetResourceRecordOptions{}
			getResourceRecordOptions.SetInstanceID(instanceID)
			getResourceRecordOptions.SetDnszoneID(zoneID)
			getResourceRecordOptions.SetRecordID(racID)
			getResourceRecordOptions.SetXCorrelationID("teststring")
			getResourceRecordOptions.SetHeaders(header)
			It(`Fail to get "A" record on zone`, func() {
				result, detailedResponse, err := service.GetResourceRecord(getResourceRecordOptions)
				Expect(result).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`Resource Record AAAA`, func() {
		Context(`Successfully create "AAAA" record on zone`, func() {
			createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
			createResourceRecordOptions.SetName("testaaaa")
			createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Aaaa)
			createResourceRecordOptions.SetTTL(120)
			rdataAaaaRecord, err := service.NewResourceRecordInputRdataRdataAaaaRecord("2001::8888")
			It(`Successfully set create "AAAA" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			createResourceRecordOptions.SetRdata(rdataAaaaRecord)
			rcaaaaresult, detailedResponse, err := service.CreateResourceRecord(createResourceRecordOptions)
			It(`Successfully create "AAAA" record on zone`, func() {
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*rcaaaaresult.Type).To(Equal(dnssvcsv1.CreateResourceRecordOptions_Type_Aaaa))
				Expect(*rcaaaaresult.ID).NotTo(BeNil())
			})
		})

		Context(`Successfully update "AAAA" record on zone`, func() {
			rcaaaaID := os.Getenv("AAAARECORD")
			It(`Successfully Get AAAA record ID`, func() {
				Expect(rcaaaaID).NotTo(BeEmpty())
			})
			updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, rcaaaaID)
			updateResourceRecordOptions.SetName("updateaaaa")
			updateResourceRecordOptions.SetTTL(300)
			updaterdataAaaaRecord, err := service.NewResourceRecordUpdateInputRdataRdataAaaaRecord("2001::8889")
			It(`Successfully set update "AAAA" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			updateResourceRecordOptions.SetRdata(updaterdataAaaaRecord)

			It(`Successfully update "AAAA" record on zone`, func() {
				result, detailedResponse, err := service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(rcaaaaID))
			})
		})

		Context(`Successfully get "AAAA" record on zone`, func() {
			rcaaaaID := os.Getenv("AAAARECORD")
			getResourceRecordOptions := service.NewGetResourceRecordOptions(instanceID, zoneID, rcaaaaID)
			It(`Successfully get "A" record on zone`, func() {
				result, detailedResponse, err := service.GetResourceRecord(getResourceRecordOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(rcaaaaID))
			})
		})

	})

	Describe(`Resource Record CNAME`, func() {
		Context(`Successfully create "CNAME" record on zone`, func() {
			createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
			createResourceRecordOptions.SetName("testcname")
			createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Cname)
			createResourceRecordOptions.SetTTL(120)
			rdataCnameRecord, err := service.NewResourceRecordInputRdataRdataCnameRecord("testcnamedata.com")
			It(`Successfully set create "CNAME" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			createResourceRecordOptions.SetRdata(rdataCnameRecord)
			rccnameresult, detailedResponse, err := service.CreateResourceRecord(createResourceRecordOptions)
			It(`Successfully create "CNAME" record on zone`, func() {
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*rccnameresult.Type).To(Equal(dnssvcsv1.CreateResourceRecordOptions_Type_Cname))
				Expect(*rccnameresult.ID).NotTo(BeNil())
			})
		})

		Context(`Successfully update "CNAME" record on zone`, func() {
			rccnameID := os.Getenv("CNAMERECORD")
			It(`Successfully Get CNAME record ID`, func() {
				Expect(rccnameID).NotTo(BeEmpty())
			})
			updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, rccnameID)
			updateResourceRecordOptions.SetName("updatecname")
			updateResourceRecordOptions.SetTTL(300)
			updaterdataCnameRecord, err := service.NewResourceRecordUpdateInputRdataRdataCnameRecord("updatecnamedata.com")
			It(`Successfully set update "CNAME" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			updateResourceRecordOptions.SetRdata(updaterdataCnameRecord)

			It(`Successfully update "CNAME" record on zone`, func() {
				result, detailedResponse, err := service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(rccnameID))
			})
		})
	})

	Describe(`Resource Record MX`, func() {
		Context(`Successfully create/update/delete "MX" record on zone`, func() {
			createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
			createResourceRecordOptions.SetName("testmx")
			createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Mx)
			createResourceRecordOptions.SetTTL(120)
			rdataMxRecord, err := service.NewResourceRecordInputRdataRdataMxRecord("mail.testmx.com", 1)
			It(`Successfully set create "MX" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			createResourceRecordOptions.SetRdata(rdataMxRecord)
			rcmxresult, detailedResponse, err := service.CreateResourceRecord(createResourceRecordOptions)
			It(`Successfully create "MX" record on zone`, func() {
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*rcmxresult.Type).To(Equal(dnssvcsv1.CreateResourceRecordOptions_Type_Mx))
				Expect(*rcmxresult.ID).NotTo(BeNil())
			})
		})

		Context(`Successfully update "MX" record on zone`, func() {
			rcmxID := os.Getenv("MXRECORD")
			It(`Successfully Get MX record ID`, func() {
				Expect(rcmxID).NotTo(BeEmpty())
			})
			updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, rcmxID)
			updateResourceRecordOptions.SetName("testupdatemx")
			updateResourceRecordOptions.SetTTL(300)
			updaterdataMxRecord, err := service.NewResourceRecordUpdateInputRdataRdataMxRecord("mail1.testmx.com", 2)
			It(`Successfully set update "MX" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			updateResourceRecordOptions.SetRdata(updaterdataMxRecord)
			It(`Successfully update "MX" record on zone`, func() {
				result, detailedResponse, err := service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(rcmxID))
			})
		})
	})

	Describe(`Resource Record PTR`, func() {
		Context(`Successfully create/update/delete "PTR" record on zone`, func() {
			parts := strings.SplitN(zoneID, ":", 2)
			zoneName := parts[0]
			createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
			createResourceRecordOptions.SetName("1.1.1.1")
			createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Ptr)
			createResourceRecordOptions.SetTTL(120)
			rdataPtrRecord, err := service.NewResourceRecordInputRdataRdataPtrRecord("testa." + zoneName)
			It(`Successfully set "PTR" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			createResourceRecordOptions.SetRdata(rdataPtrRecord)
			rcptrresult, detailedResponse, err := service.CreateResourceRecord(createResourceRecordOptions)
			It(`Successfully create "PTR" record on zone`, func() {
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*rcptrresult.Type).To(Equal(dnssvcsv1.CreateResourceRecordOptions_Type_Ptr))
				Expect(*rcptrresult.ID).NotTo(BeNil())
			})
		})

		Context(`Successfully update "PTR" record on zone`, func() {
			rcptrID := os.Getenv("PTRRECORD")
			It(`Successfully Get PTR record ID`, func() {
				Expect(rcptrID).NotTo(BeEmpty())
			})
			updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, rcptrID)
			updateResourceRecordOptions.SetTTL(300)
			It(`Successfully update "PTR" record on zone`, func() {
				result, detailedResponse, err := service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(rcptrID))
			})
		})
	})

	Describe(`Resource Record SRV`, func() {
		Context(`Successfully create "SRV" record on zone`, func() {
			createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
			createResourceRecordOptions.SetName("testsrv")
			createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Srv)
			createResourceRecordOptions.SetTTL(120)
			createResourceRecordOptions.SetService("_sip")
			createResourceRecordOptions.SetProtocol("udp")
			rdataSrvRecord, err := service.NewResourceRecordInputRdataRdataSrvRecord(1, 1, "siphost.com", 1)
			It(`Successfully set "SRV" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			createResourceRecordOptions.SetRdata(rdataSrvRecord)

			rcsrvresult, detailedResponse, err := service.CreateResourceRecord(createResourceRecordOptions)
			It(`Successfully create "SRV" record on zone`, func() {
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*rcsrvresult.Type).To(Equal(dnssvcsv1.CreateResourceRecordOptions_Type_Srv))
				Expect(*rcsrvresult.ID).NotTo(BeNil())
			})
		})

		Context(`Successfully update "SRV" record on zone`, func() {
			rcsrvID := os.Getenv("SRVRECORD")
			It(`Successfully Get SRV record ID`, func() {
				Expect(rcsrvID).NotTo(BeEmpty())
			})
			updateResourceRecordOptions := service.NewUpdateResourceRecordOptions(instanceID, zoneID, rcsrvID)
			updateResourceRecordOptions.SetName("updatesrv")
			updateResourceRecordOptions.SetTTL(300)
			updaterdataSrvRecord, err := service.NewResourceRecordUpdateInputRdataRdataSrvRecord(2, 2, "updatesiphost.com", 2)
			It(`Successfully set update "SRV" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			updateResourceRecordOptions.SetRdata(updaterdataSrvRecord)
			It(`Successfully update "SRV" record on zone`, func() {
				result, detailedResponse, err := service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(rcsrvID))
			})
		})
	})

	Describe(`Resource Record TXT`, func() {
		Context(`Successfully create "TXT" record on zone`, func() {
			createResourceRecordOptions := service.NewCreateResourceRecordOptions(instanceID, zoneID)
			createResourceRecordOptions.SetName("testtxt")
			createResourceRecordOptions.SetType(dnssvcsv1.CreateResourceRecordOptions_Type_Txt)
			createResourceRecordOptions.SetTTL(120)
			rdataTxtRecord, err := service.NewResourceRecordInputRdataRdataTxtRecord("txtdata string")
			It(`Successfully set "TXT" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			createResourceRecordOptions.SetRdata(rdataTxtRecord)
			rctxtresult, detailedResponse, err := service.CreateResourceRecord(createResourceRecordOptions)
			It(`Successfully create "TXT" record on zone`, func() {
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*rctxtresult.Type).To(Equal(dnssvcsv1.CreateResourceRecordOptions_Type_Txt))
				Expect(*rctxtresult.ID).NotTo(BeNil())
			})
		})

		Context(`Successfully update "TXT" record on zone`, func() {
			rctxtID := os.Getenv("TXTRECORD")
			It(`Successfully Get TXT record ID`, func() {
				Expect(rctxtID).NotTo(BeEmpty())
			})
			header := map[string]string{
				"Content-type": "application/json",
			}
			updateResourceRecordOptions := &dnssvcsv1.UpdateResourceRecordOptions{}
			updateResourceRecordOptions.SetInstanceID(instanceID)
			updateResourceRecordOptions.SetDnszoneID(zoneID)
			updateResourceRecordOptions.SetRecordID(rctxtID)
			updateResourceRecordOptions.SetTTL(300)
			updateResourceRecordOptions.SetHeaders(header)
			updateResourceRecordOptions.SetName("updatesttxt")
			updaterdataTxtRecord, err := service.NewResourceRecordUpdateInputRdataRdataTxtRecord("update txtdata string")
			It(`Successfully set update "TXT" record Rdata`, func() {
				Expect(err).To(BeNil())
			})
			updateResourceRecordOptions.SetRdata(updaterdataTxtRecord)
			It(`Successfully set update "TXT" record Rdata`, func() {
				result, detailedResponse, err := service.UpdateResourceRecord(updateResourceRecordOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(*result.ID).To(Equal(rctxtID))
			})
		})

		Context(`Fail to delete "TXT" record on zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			rctxtID := "bad_bird"
			deleteResourceRecordOptions := &dnssvcsv1.DeleteResourceRecordOptions{}
			deleteResourceRecordOptions.SetInstanceID(instanceID)
			deleteResourceRecordOptions.SetDnszoneID(zoneID)
			deleteResourceRecordOptions.SetRecordID(rctxtID)
			deleteResourceRecordOptions.SetHeaders(header)
			deleteResourceRecordOptions.SetXCorrelationID("teststring")
			It(`Fail delete "TXT" record on zone`, func() {
				detailedResponse, err := service.DeleteResourceRecord(deleteResourceRecordOptions)
				Expect(detailedResponse.StatusCode).To(Equal(400))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`List resource records`, func() {
		Context(`Successfully list resource records on zone`, func() {
			zoneID := os.Getenv("ZONE_ID")
			listResourceRecordsOptions := service.NewListResourceRecordsOptions(instanceID, zoneID)
			It(`Successfully list all resources`, func() {
				result, detailedResponse, err := service.ListResourceRecords(listResourceRecordsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				firstResource := result.ResourceRecords[0]
				Expect(*firstResource.ID).NotTo(BeNil())
			})
		})

		Context(`Failed to list resource records on zone`, func() {
			header := map[string]string{
				"Content-type": "application/json",
			}
			zoneID := "bad_bird"
			listResourceRecordsOptions := &dnssvcsv1.ListResourceRecordsOptions{}
			listResourceRecordsOptions.SetInstanceID(instanceID)
			listResourceRecordsOptions.SetDnszoneID(zoneID)
			listResourceRecordsOptions.SetHeaders(header)
			listResourceRecordsOptions.SetXCorrelationID("teststring")
			It(`Failed to list all resouces`, func() {
				_, _, err := service.ListResourceRecords(listResourceRecordsOptions)
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe(`Common Record Unmarshal utility test`, func() {
		var (
			testInputData map[string]interface{}
		)
		BeforeEach(func() {
			testInputData = map[string]interface{}{
				"ip":         "1.1.1.1",
				"cname":      "cname",
				"exchange":   "exchange",
				"preference": float64(1),
				"port":       float64(1),
				"priority":   float64(1),
				"target":     "target",
				"weight":     float64(1),
				"text":       "text",
				"ptrdname":   "ptrdname",
			}
		})

		Context(`Successfully Invoke utility test`, func() {
			It(`Invoke UnmarshalResourceRecordInputRdata() successfully`, func() {
				result, err := dnssvcsv1.UnmarshalResourceRecordInputRdata(testInputData)
				Expect(result).NotTo(BeNil())
				Expect(err).To(BeNil())
			})

			It(`Invoke UnmarshalResourceRecordUpdateInputRdata() successfully`, func() {
				result, err := dnssvcsv1.UnmarshalResourceRecordUpdateInputRdata(testInputData)
				Expect(result).NotTo(BeNil())
				Expect(err).To(BeNil())
			})

		})
	})

	Describe(`Record Rdata Unmarshal utility test`, func() {
		Context(`Successfully Invoke A rdata utility test`, func() {
			var (
				aRdata map[string]interface{}
			)

			BeforeEach(func() {
				aRdata = map[string]interface{}{
					"ip": "1.1.1.1",
				}
			})
			It(`Invoke UnmarshalResourceRecordInputRdataRdataARecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordInputRdataRdataARecord(aRdata)
				Expect(err).To(BeNil())
			})

			It(`Invoke UnmarshalResourceRecordUpdateInputRdataRdataARecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordUpdateInputRdataRdataARecord(aRdata)
				Expect(err).To(BeNil())
			})
		})

		Context(`Successfully Invoke AAAA rdata utility test`, func() {
			var (
				aaaaRdata map[string]interface{}
			)

			BeforeEach(func() {
				aaaaRdata = map[string]interface{}{
					"ip": "2001::1234",
				}
			})

			It(`Invoke UnmarshalResourceRecordInputRdataRdataAaaaRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordInputRdataRdataAaaaRecord(aaaaRdata)
				Expect(err).To(BeNil())
			})

			It(`Invoke UnmarshalResourceRecordUpdateInputRdataRdataAaaaRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordUpdateInputRdataRdataAaaaRecord(aaaaRdata)
				Expect(err).To(BeNil())
			})
		})

		Context(`Successfully Invoke CNAME rdata utility test`, func() {
			var (
				cnameRdata map[string]interface{}
			)

			BeforeEach(func() {
				cnameRdata = map[string]interface{}{
					"cname": "example.com",
				}
			})

			It(`Invoke UnmarshalResourceRecordInputRdataRdataCnameRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordInputRdataRdataCnameRecord(cnameRdata)
				Expect(err).To(BeNil())
			})

			It(`Invoke UnmarshalResourceRecordUpdateInputRdataRdataCnameRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordUpdateInputRdataRdataCnameRecord(cnameRdata)
				Expect(err).To(BeNil())
			})
		})

		Context(`Successfully Invoke MX rdata utility test`, func() {
			var (
				mxRdata map[string]interface{}
			)

			BeforeEach(func() {
				mxRdata = map[string]interface{}{
					"preference": float64(2),
					"exchange":   "mail1.example.com",
				}
			})

			It(`Invoke UnmarshalResourceRecordInputRdataRdataMxRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordInputRdataRdataMxRecord(mxRdata)
				Expect(err).To(BeNil())
			})

			It(`Invoke UnmarshalResourceRecordUpdateInputRdataRdataMxRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordUpdateInputRdataRdataMxRecord(mxRdata)
				Expect(err).To(BeNil())
			})
		})

		Context(`Successfully Invoke PTR rdata utility test`, func() {
			var (
				ptrRdata map[string]interface{}
			)

			BeforeEach(func() {
				ptrRdata = map[string]interface{}{
					"ptrdname": "example.com",
				}
			})

			It(`Invoke UnmarshalResourceRecordInputRdataRdataPtrRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordInputRdataRdataPtrRecord(ptrRdata)
				Expect(err).To(BeNil())
			})

			It(`Invoke UnmarshalResourceRecordUpdateInputRdataRdataPtrRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordUpdateInputRdataRdataPtrRecord(ptrRdata)
				Expect(err).To(BeNil())
			})
		})

		Context(`Successfully Invoke SVR rdata utility test`, func() {
			var (
				srvRdata map[string]interface{}
			)

			BeforeEach(func() {
				srvRdata = map[string]interface{}{
					"priority": float64(2),
					"weight":   float64(2),
					"port":     float64(2),
					"target":   "example.com",
				}
			})

			It(`Invoke UnmarshalResourceRecordInputRdataRdataSrvRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordInputRdataRdataSrvRecord(srvRdata)
				Expect(err).To(BeNil())
			})

			It(`Invoke UnmarshalResourceRecordUpdateInputRdataRdataSrvRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordUpdateInputRdataRdataSrvRecord(srvRdata)
				Expect(err).To(BeNil())
			})
		})

		Context(`Successfully Invoke TXT rdata utility test`, func() {
			var (
				txtRdata map[string]interface{}
			)

			BeforeEach(func() {
				txtRdata = map[string]interface{}{
					"text": "text string",
				}
			})

			It(`Invoke UnmarshalResourceRecordInputRdataRdataTxtRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordInputRdataRdataTxtRecord(txtRdata)
				Expect(err).To(BeNil())
			})

			It(`Invoke UnmarshalResourceRecordUpdateInputRdataRdataTxtRecord() successfully`, func() {
				_, err := dnssvcsv1.UnmarshalResourceRecordUpdateInputRdataRdataTxtRecord(txtRdata)
				Expect(err).To(BeNil())
			})
		})

	})

})
