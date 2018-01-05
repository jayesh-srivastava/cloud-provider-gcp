/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package meta

import (
	"reflect"

	alpha "google.golang.org/api/compute/v0.alpha"
	beta "google.golang.org/api/compute/v0.beta"
	ga "google.golang.org/api/compute/v1"
)

// Version of the API (ga, alpha, beta).
type Version string

const (
	// NoGet prevents the Get() method from being generated.
	NoGet = 1 << iota
	// NoList prevents the List() method from being generated.
	NoList = 1 << iota
	// NoDelete prevents the Delete() method from being generated.
	NoDelete = 1 << iota
	// NoInsert prevents the Insert() method from being generated.
	NoInsert = 1 << iota
	// CustomOps specifies that an empty interface xxxOps will be generated to
	// enable custom method calls to be attached to the generated service
	// interface.
	CustomOps = 1 << iota
	// AggregatedList will generated a method for AggregatedList().
	AggregatedList = 1 << iota

	// ReadOnly specifies that the given resource is read-only and should not
	// have insert() or delete() methods generated for the wrapper.
	ReadOnly = NoDelete | NoInsert

	// VersionGA is the API version in compute.v1.
	VersionGA Version = "ga"
	// VersionAlpha is the API version in computer.v0.alpha.
	VersionAlpha Version = "alpha"
	// VersionBeta is the API version in computer.v0.beta.
	VersionBeta Version = "beta"
)

// AllVersions is a list of all versions of the GCE API.
var AllVersions = []Version{
	VersionGA,
	VersionAlpha,
	VersionBeta,
}

// AllServices are a list of all the services to generate code for. Keep
// this list in lexiographical order by object type.
var AllServices = []*ServiceInfo{
	&ServiceInfo{
		Object:      "Address",
		Service:     "Addresses",
		Resource:    "addresses",
		keyType:     Regional,
		serviceType: reflect.TypeOf(&ga.AddressesService{}),
	},
	&ServiceInfo{
		Object:      "Address",
		Service:     "Addresses",
		Resource:    "addresses",
		version:     VersionAlpha,
		keyType:     Regional,
		serviceType: reflect.TypeOf(&alpha.AddressesService{}),
	},
	&ServiceInfo{
		Object:      "Address",
		Service:     "Addresses",
		Resource:    "addresses",
		version:     VersionBeta,
		keyType:     Regional,
		serviceType: reflect.TypeOf(&beta.AddressesService{}),
	},
	&ServiceInfo{
		Object:      "Address",
		Service:     "GlobalAddresses",
		Resource:    "addresses",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.GlobalAddressesService{}),
	},
	&ServiceInfo{
		Object:      "BackendService",
		Service:     "BackendServices",
		Resource:    "backendServices",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.BackendServicesService{}),
		additionalMethods: []string{
			"GetHealth",
			"Update",
		},
	},
	&ServiceInfo{
		Object:            "BackendService",
		Service:           "BackendServices",
		Resource:    "backendServices",
		version:           VersionAlpha,
		keyType:           Global,
		serviceType:       reflect.TypeOf(&alpha.BackendServicesService{}),
		additionalMethods: []string{"Update"},
	},
	&ServiceInfo{
		Object:      "BackendService",
		Service:     "RegionBackendServices",
		Resource:    "backendServices",
		version:     VersionAlpha,
		keyType:     Regional,
		serviceType: reflect.TypeOf(&alpha.RegionBackendServicesService{}),
		additionalMethods: []string{
			"GetHealth",
			"Update",
		},
	},
	&ServiceInfo{
		Object:      "Disk",
		Service:     "Disks",
		Resource:    "disks",
		keyType:     Zonal,
		serviceType: reflect.TypeOf(&ga.DisksService{}),
	},
	&ServiceInfo{
		Object:      "Disk",
		Service:     "Disks",
		Resource:    "disks",
		version:     VersionAlpha,
		keyType:     Zonal,
		serviceType: reflect.TypeOf(&alpha.DisksService{}),
	},
	&ServiceInfo{
		Object:      "Disk",
		Service:     "RegionDisks",
		Resource:    "disks",
		version:     VersionAlpha,
		keyType:     Regional,
		serviceType: reflect.TypeOf(&alpha.DisksService{}),
	},
	&ServiceInfo{
		Object:      "Firewall",
		Service:     "Firewalls",
		Resource:    "firewalls",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.FirewallsService{}),
		additionalMethods: []string{
			"Update",
		},
	},
	&ServiceInfo{
		Object:      "ForwardingRule",
		Service:     "ForwardingRules",
		Resource:    "forwardingRules",
		keyType:     Regional,
		serviceType: reflect.TypeOf(&ga.ForwardingRulesService{}),
	},
	&ServiceInfo{
		Object:      "ForwardingRule",
		Service:     "ForwardingRules",
		Resource:    "forwardingRules",
		version:     VersionAlpha,
		keyType:     Regional,
		serviceType: reflect.TypeOf(&alpha.ForwardingRulesService{}),
	},
	&ServiceInfo{
		Object:      "ForwardingRule",
		Service:     "GlobalForwardingRules",
		Resource:    "forwardingRules",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.GlobalForwardingRulesService{}),
		additionalMethods: []string{
			"SetTarget",
		},
	},
	&ServiceInfo{
		Object:      "HealthCheck",
		Service:     "HealthChecks",
		Resource:    "healthChecks",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.HealthChecksService{}),
		additionalMethods: []string{
			"Update",
		},
	},
	&ServiceInfo{
		Object:      "HealthCheck",
		Service:     "HealthChecks",
		Resource:    "healthChecks",
		version:     VersionAlpha,
		keyType:     Global,
		serviceType: reflect.TypeOf(&alpha.HealthChecksService{}),
		additionalMethods: []string{
			"Update",
		},
	},
	&ServiceInfo{
		Object:      "HttpHealthCheck",
		Service:     "HttpHealthChecks",
		Resource:    "httpHealthChecks",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.HttpHealthChecksService{}),
		additionalMethods: []string{
			"Update",
		},
	},
	&ServiceInfo{
		Object:      "HttpsHealthCheck",
		Service:     "HttpsHealthChecks",
		Resource:    "httpsHealthChecks",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.HttpsHealthChecksService{}),
		additionalMethods: []string{
			"Update",
		},
	},
	&ServiceInfo{
		Object:      "InstanceGroup",
		Service:     "InstanceGroups",
		Resource:    "instanceGroups",
		keyType:     Zonal,
		serviceType: reflect.TypeOf(&ga.InstanceGroupsService{}),
		additionalMethods: []string{
			"AddInstances",
			"ListInstances",
			"RemoveInstances",
			"SetNamedPorts",
		},
	},
	&ServiceInfo{
		Object:      "Instance",
		Service:     "Instances",
		Resource:    "instances",
		keyType:     Zonal,
		serviceType: reflect.TypeOf(&ga.InstancesService{}),
		additionalMethods: []string{
			"AttachDisk",
			"DetachDisk",
		},
	},
	&ServiceInfo{
		Object:      "Instance",
		Service:     "Instances",
		Resource:    "instances",
		version:     VersionBeta,
		keyType:     Zonal,
		serviceType: reflect.TypeOf(&beta.InstancesService{}),
		additionalMethods: []string{
			"AttachDisk",
			"DetachDisk",
		},
	},
	&ServiceInfo{
		Object:      "Instance",
		Service:     "Instances",
		Resource:    "instances",
		version:     VersionAlpha,
		keyType:     Zonal,
		serviceType: reflect.TypeOf(&alpha.InstancesService{}),
		additionalMethods: []string{
			"AttachDisk",
			"DetachDisk",
			"UpdateNetworkInterface",
		},
	},
	&ServiceInfo{
		Object:      "NetworkEndpointGroup",
		Service:     "NetworkEndpointGroups",
		Resource:    "networkEndpointGroups",
		version:     VersionAlpha,
		keyType:     Zonal,
		serviceType: reflect.TypeOf(&alpha.NetworkEndpointGroupsService{}),
		additionalMethods: []string{
			"AttachNetworkEndpoints",
			"DetachNetworkEndpoints",
		},
		options: AggregatedList,
	},
	&ServiceInfo{
		Object:  "Project",
		Service: "Projects",
		Resource: "projects",
		keyType: Global,
		// Generate only the stub with no methods.
		options:     NoGet | NoList | NoInsert | NoDelete | CustomOps,
		serviceType: reflect.TypeOf(&ga.ProjectsService{}),
	},
	&ServiceInfo{
		Object:      "Region",
		Service:     "Regions",
		Resource:    "regions",
		keyType:     Global,
		options:     ReadOnly,
		serviceType: reflect.TypeOf(&ga.RegionsService{}),
	},
	&ServiceInfo{
		Object:      "Route",
		Service:     "Routes",
		Resource:    "routes",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.RoutesService{}),
	},
	&ServiceInfo{
		Object:      "SslCertificate",
		Service:     "SslCertificates",
		Resource:    "sslCertificates",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.SslCertificatesService{}),
	},
	&ServiceInfo{
		Object:      "TargetHttpProxy",
		Service:     "TargetHttpProxies",
		Resource:    "targetHttpProxies",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.TargetHttpProxiesService{}),
		additionalMethods: []string{
			"SetUrlMap",
		},
	},
	&ServiceInfo{
		Object:      "TargetHttpsProxy",
		Service:     "TargetHttpsProxies",
		Resource:    "targetHttpsProxies",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.TargetHttpsProxiesService{}),
		additionalMethods: []string{
			"SetSslCertificates",
			"SetUrlMap",
		},
	},
	&ServiceInfo{
		Object:      "TargetPool",
		Service:     "TargetPools",
		Resource:    "targetPools",
		keyType:     Regional,
		serviceType: reflect.TypeOf(&ga.TargetPoolsService{}),
		additionalMethods: []string{
			"AddInstance",
			"RemoveInstance",
		},
	},
	&ServiceInfo{
		Object:      "UrlMap",
		Service:     "UrlMaps",
		Resource:    "urlMaps",
		keyType:     Global,
		serviceType: reflect.TypeOf(&ga.UrlMapsService{}),
		additionalMethods: []string{
			"Update",
		},
	},
	&ServiceInfo{
		Object:      "Zone",
		Service:     "Zones",
		Resource:    "zones",
		keyType:     Global,
		options:     ReadOnly,
		serviceType: reflect.TypeOf(&ga.ZonesService{}),
	},
}
