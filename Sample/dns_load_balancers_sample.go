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
	"github.com/IBM/go-sdk-core/v3/core"
)

// DNS Service Load Balancers
var monitorID string
var poolID string
var glbID string

// listDnsGlbMonitors - List loadbalancer monitors
func listDnsGlbMonitors() {
	listDnsGlbMonitorOptions := dnsSvc.NewListMonitorsOptions(instanceID)
	_, listDnsGlbMonitorResponse, reqErr := dnsSvc.ListMonitors(listDnsGlbMonitorOptions)
	if reqErr == nil {
		fmt.Println(listDnsGlbMonitorResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// createDnsGlbMonitor - Create loadbalancer monitor
func createDnsGlbMonitor() {
	createDnsGlbMonitorOptions := dnsSvc.NewCreateMonitorOptions(instanceID)
	createDnsGlbMonitorOptions.SetDescription("Load balancer monitor for example.com")
	createDnsGlbMonitorOptions.SetType("HTTPS")
	createDnsGlbMonitorOptions.SetPort(int64(8080))
	createDnsGlbMonitorOptions.SetInterval(int64(60))
	createDnsGlbMonitorOptions.SetRetries(int64(2))
	createDnsGlbMonitorOptions.SetTimeout(int64(5))
	createDnsGlbMonitorOptions.SetMethod("GET")
	createDnsGlbMonitorOptions.SetPath("/health")
	header := map[string]string{"Host": "example.com", "X-App-ID": "abc123"}
	createDnsGlbMonitorOptions.SetHeader(header)
	createDnsGlbMonitorOptions.SetAllowInsecure(false)
	createDnsGlbMonitorOptions.SetExpectedCodes("200")
	createDnsGlbMonitorOptions.SetExpectedBody("alive")
	createDnsGlbMonitorOptions.SetFollowRedirects(false)
	_, createDnsGlbMonitorResponse, reqErr := dnsSvc.CreateMonitor(createDnsGlbMonitorOptions)
	if reqErr == nil {
		fmt.Println(createDnsGlbMonitorResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// getDnsGlbMonitor - Get loadbalancer monitor
func getDnsGlbMonitor() {
	monitorID = os.Getenv("MONITOR_ID")
	getDnsGlbMonitorOptions := dnsSvc.NewGetMonitorOptions(instanceID, monitorID)
	_, getDnsGlbMonitorResponse, reqErr := dnsSvc.GetMonitor(getDnsGlbMonitorOptions)
	if reqErr == nil {
		fmt.Println(getDnsGlbMonitorResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// updateDnsGlbMonitor - Update loadbalancer monitor
func updateDnsGlbMonitor() {
	monitorID = os.Getenv("MONITOR_ID")
	updateDnsGlbMonitorOptions := dnsSvc.NewUpdateMonitorOptions(instanceID, monitorID)
	updateDnsGlbMonitorOptions.SetDescription("Load balancer monitor for example.com")
	updateDnsGlbMonitorOptions.SetType("HTTPS")
	updateDnsGlbMonitorOptions.SetPort(int64(8080))
	updateDnsGlbMonitorOptions.SetInterval(int64(60))
	updateDnsGlbMonitorOptions.SetRetries(int64(2))
	updateDnsGlbMonitorOptions.SetTimeout(int64(5))
	updateDnsGlbMonitorOptions.SetMethod("GET")
	updateDnsGlbMonitorOptions.SetPath("/health")
	header := map[string]string{"Host": "example.com", "X-App-ID": "abc123"}
	updateDnsGlbMonitorOptions.SetHeader(header)
	updateDnsGlbMonitorOptions.SetAllowInsecure(false)
	updateDnsGlbMonitorOptions.SetExpectedCodes("200")
	updateDnsGlbMonitorOptions.SetExpectedBody("alive")
	updateDnsGlbMonitorOptions.SetFollowRedirects(false)
	_, updateDnsGlbMonitorResponse, reqErr := dnsSvc.UpdateMonitor(updateDnsGlbMonitorOptions)
	if reqErr == nil {
		fmt.Println(updateDnsGlbMonitorResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// deleteDnsGlbMonitor - Delete loadbalancer monitor
func deleteDnsGlbMonitor() {
	monitorID = os.Getenv("MONITOR_ID")
	deleteDnsGlbMonitorOptions := dnsSvc.NewDeleteMonitorOptions(instanceID, monitorID)
	deleteDnsGlbMonitorResponse, reqErr := dnsSvc.DeleteMonitor(deleteDnsGlbMonitorOptions)
	if reqErr == nil {
		fmt.Println(deleteDnsGlbMonitorResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// listDnsGlbPools - List Loadbalancer Pools
func listDnsGlbPools() {
	listDnsGlbPoolOptions := dnsSvc.NewListPoolsOptions(instanceID)
	_, listDnsGlbPoolResponse, reqErr := dnsSvc.ListPools(listDnsGlbPoolOptions)
	if reqErr == nil {
		fmt.Println(listDnsGlbPoolResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// createDnsGlbPool - Create Loadbalancer Pool
func createDnsGlbPool() {
	monitorID = os.Getenv("MONITOR_ID")
	createDnsGlbPoolOptions := dnsSvc.NewCreatePoolOptions(instanceID)
	createDnsGlbPoolOptions.SetName("dal-pool")
	createDnsGlbPoolOptions.SetDescription("dallas pool for example.com")
	createDnsGlbPoolOptions.SetEnabled(true)
	createDnsGlbPoolOptions.SetMinimumOrigins(int64(1))
	origin1 := new(dnssvcsv1.Origin)
	origin1.Name = core.StringPtr("dal-origin01")
	origin1.Description = core.StringPtr("description of the origin server")
	origin1.Address = core.StringPtr("10.10.16.8")
	origin1.Enabled = core.BoolPtr(true)
	origin1.Weight = core.Int64Ptr(int64(1))
	createDnsGlbPoolOptions.SetOrigins([]dnssvcsv1.Origin{*origin1})
	createDnsGlbPoolOptions.SetMonitor(monitorID)
	createDnsGlbPoolOptions.SetNotificationType(dnssvcsv1.CreatePoolOptions_NotificationType_Email)
	createDnsGlbPoolOptions.SetNotificationChannel("xxx@email.example.com")
	_, createDnsGlbPoolResponse, reqErr := dnsSvc.CreatePool(createDnsGlbPoolOptions)
	if reqErr == nil {
		fmt.Println(createDnsGlbPoolResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// getDnsGlbPool - Get Loadbalancer Pool
func getDnsGlbPool() {
	poolID = os.Getenv("POOL_ID")
	getDnsGlbPoolOptions := dnsSvc.NewGetPoolOptions(instanceID, poolID)
	_, getDnsGlbPoolResponse, reqErr := dnsSvc.GetPool(getDnsGlbPoolOptions)
	if reqErr == nil {
		fmt.Println(getDnsGlbPoolResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// updateDnsGlbPool - Update Loadbalancer Pool
func updateDnsGlbPool() {
	monitorID = os.Getenv("MONITOR_ID")
	poolID = os.Getenv("POOL_ID")
	updateDnsGlbPoolOptions := dnsSvc.NewUpdatePoolOptions(instanceID, poolID)
	updateDnsGlbPoolOptions.SetName("dal-pool")
	updateDnsGlbPoolOptions.SetDescription("dallas pool for example.com")
	updateDnsGlbPoolOptions.SetEnabled(true)
	updateDnsGlbPoolOptions.SetMinimumOrigins(int64(1))
	origin2 := new(dnssvcsv1.Origin)
	origin2.Name = core.StringPtr("dal-origin02")
	origin2.Description = core.StringPtr("description of the origin server")
	origin2.Address = core.StringPtr("10.10.16.9")
	origin2.Enabled = core.BoolPtr(true)
	origin2.Weight = core.Int64Ptr(int64(1))
	updateDnsGlbPoolOptions.SetOrigins([]dnssvcsv1.Origin{*origin2})
	updateDnsGlbPoolOptions.SetMonitor(monitorID)
	updateDnsGlbPoolOptions.SetNotificationType(dnssvcsv1.CreatePoolOptions_NotificationType_Email)
	updateDnsGlbPoolOptions.SetNotificationChannel("xxx@email.example.com")
	_, updateDnsGlbPoolResponse, reqErr := dnsSvc.UpdatePool(updateDnsGlbPoolOptions)
	if reqErr == nil {
		fmt.Println(updateDnsGlbPoolResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// deleteDnsGlbPool - Delete Loadbalancer Pool
func deleteDnsGlbPool() {
	poolID = os.Getenv("POOL_ID")
	deleteDnsGlbPoolOptions := dnsSvc.NewDeletePoolOptions(instanceID, poolID)
	deleteDnsGlbPoolResponse, reqErr := dnsSvc.DeletePool(deleteDnsGlbPoolOptions)
	if reqErr == nil {
		fmt.Println(deleteDnsGlbPoolResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// listGnsGlb - List Loadbalancer
func listDnsGlb() {
	listDnsGlbOptions := dnsSvc.NewListLoadBalancersOptions(instanceID, zoneID)
	_, listDnsGlbResponse, reqErr := dnsSvc.ListLoadBalancers(listDnsGlbOptions)
	if reqErr == nil {
		fmt.Println(listDnsGlbResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// createDnsGlb - Create Loadbalancer
func createDnsGlb() {
	createDnsGlbOptions := dnsSvc.NewCreateLoadBalancerOptions(instanceID, zoneID)
	createDnsGlbOptions.SetName("glb01.example.com")
	createDnsGlbOptions.SetDescription("Global load balancer 01")
	createDnsGlbOptions.SetEnabled(true)
	createDnsGlbOptions.SetTTL(int64(300))
	createDnsGlbOptions.SetDefaultPools([]string{"pool1 ID", "pool2 ID"})
	createDnsGlbOptions.SetFallbackPool("falback pool ID")
	azPools := new(dnssvcsv1.AzPools)
	azPools.UsSouth1 = []string{"us south1 pool ID"}
	azPools.UsSouth2 = []string{"us south2 pool ID"}
	azPools.UsSouth3 = []string{"us south3 pool ID"}
	createDnsGlbOptions.SetAzPools(azPools)
	_, createDnsGlbResponse, reqErr := dnsSvc.CreateLoadBalancer(createDnsGlbOptions)
	if reqErr == nil {
		fmt.Println(createDnsGlbResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// getDnsGlb - Get Loadbalancer
func getDnsGlb() {
	glbID = os.Getenv("GLB_ID")
	getDnsGlbOptions := dnsSvc.NewGetLoadBalancerOptions(instanceID, zoneID, glbID)
	_, getDnsGlbResponse, reqErr := dnsSvc.GetLoadBalancer(getDnsGlbOptions)
	if reqErr == nil {
		fmt.Println(getDnsGlbResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// updateDnsGlb - Update Loadbalancer
func updateDnsGlb() {
	glbID = os.Getenv("GLB_ID")
	updateDnsGlbOptions := dnsSvc.NewUpdateLoadBalancerOptions(instanceID, zoneID, glbID)
	updateDnsGlbOptions.SetName("glb01.example.com")
	updateDnsGlbOptions.SetDescription("Update Global load balancer 01")
	updateDnsGlbOptions.SetEnabled(true)
	updateDnsGlbOptions.SetTTL(int64(300))
	updateDnsGlbOptions.SetDefaultPools([]string{"pool1 ID", "pool2 ID"})
	updateDnsGlbOptions.SetFallbackPool("falback pool ID")
	updateazPools := new(dnssvcsv1.AzPools)
	updateazPools.UsSouth1 = []string{"us south1 pool ID"}
	updateazPools.UsSouth2 = []string{"us south2 pool ID"}
	updateazPools.UsSouth3 = []string{"us south3 pool ID"}
	updateDnsGlbOptions.SetAzPools(updateazPools)
	_, updateDnsGlbResponse, reqErr := dnsSvc.UpdateLoadBalancer(updateDnsGlbOptions)
	if reqErr == nil {
		fmt.Println(updateDnsGlbResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}

// deleteDnsGlb - Delete Loadbalancer
func deleteDnsGlb() {
	glbID = os.Getenv("GLB_ID")
	deleteDnsGlbOptions := dnsSvc.NewDeleteLoadBalancerOptions(instanceID, zoneID, glbID)
	deleteDnsGlbResponse, reqErr := dnsSvc.DeleteLoadBalancer(deleteDnsGlbOptions)
	if reqErr == nil {
		fmt.Println(deleteDnsGlbResponse.String())
	} else {
		fmt.Println(reqErr)
	}
}
