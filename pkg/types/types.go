// Copyright 2017 Intel Corp.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
)

// Exported Types
type MemifConf struct {
	Role string `json:"role,omitempty"` // Role of memif: master|slave
	Mode string `json:"mode,omitempty"` // Mode of memif: ip|ethernet|inject-punt

	// Autogenerated as memif-<ContainerID[:12]>-<IfName>.sock i.e. memif-0958c8871b32-net1.sock
	// Filename only, no path. Will use if populated, but used to passed filename to container.
	Socketfile string `json:"socketfile,omitempty"`
}

type VhostConf struct {
	Mode  string `json:"mode,omitempty"`  // vhost-user mode: client|server
	Group string `json:"group,omitempty"` // vhost-user socket file group ownership

	// Autogenerated as <ContainerID[:12]>-<IfName> i.e. 0958c8871b32-net1
	// Filename only, no path. Will use if populated, but used to passed filename to container.
	Socketfile string `json:"socketfile,omitempty"`
}

type BridgeConf struct {
	// ovs-dpdk specific note:
	//   ovs-dpdk requires a bridge to create an interfaces. So if 'NetType' is set
	//   to something other than 'bridge', a bridge is still need and this field will
	//   be inspected. For ovs-dpdk, if bridge data is not populated, it will default
	//   to 'br-0'.
	BridgeName string `json:"bridgeName,omitempty"` // Bridge Name
	BridgeId   int    `json:"bridgeId,omitempty"`   // Bridge Id - Deprecated in favor of BridgeName
	VlanId     int    `json:"vlanId,omitempty"`     // Optional VLAN Id
}

type UserSpaceConf struct {
	// The Container Instance will default to the Host Instance value if a given attribute
	// is not provided. However, they are not required to be the same and a Container
	// attribute can be provided to override. All values are listed as 'omitempty' to
	// allow the Container struct to be empty where desired.
	Engine     string     `json:"engine,omitempty"`  // CNI Implementation {vpp|ovs-dpdk}
	IfType     string     `json:"iftype,omitempty"`  // Type of interface {memif|vhostuser}
	NetType    string     `json:"netType,omitempty"` // Interface network type {none|bridge|interface}
	MemifConf  MemifConf  `json:"memif,omitempty"`
	VhostConf  VhostConf  `json:"vhost,omitempty"`
	BridgeConf BridgeConf `json:"bridge,omitempty"`
}

type NetConf struct {
	types.NetConf

	/*
		// Support chaining
		RawPrevResult *map[string]interface{} `json:"prevResult"`
		PrevResult    *current.Result         `json:"-"`
	*/

	// One of the following two must be provided: KubeConfig or SharedDir
	//
	// KubeConfig:
	//  Example: "kubeconfig": "/etc/cni/net.d/multus.d/multus.kubeconfig",
	//  Provides credentials for Userspace CNI to call KubeAPI to:
	//  - Read Volume Mounts:
	//    - "shared-dir": Directory on host socketfiles are created in
	//  - Write annotations:
	//    - "userspace/configuration-data": Configuration data passed
	//      to containe in JSON format.
	//    - "userspace/mapped-dir": Directory in container socketfiles
	//      are created in. Scraped from Volume Mounts above.
	//
	// SharedDir:
	//  Example: "sharedDir": "/usr/local/var/run/openvswitch/",
	//  Since credentials are not provided, Userspace CNI cannot call KubeAPI
	//  to read the Volume Mounts, so this is the same directory used in the
	//  "hostPath". Difference from the "kubeConfig" is that with the "sharedDir"
	//  method, the directory is not unique per POD. That is because the Network
	//  Attachment Definition (where this is defined) is used by multiple PODs.
	//  So this is the base directory and the CNI creates a sub-directory with
	//  the ContainerId as the sub-directory name.
	//
	//  Along the same lines, no annotations are written by Userspace CNI.
	//   1) Configuration data will be written to a file in the same
	//      directory as the socketfiles instead of to an annotation.
	//   2) The "userspace/mapped-dir" annotation must be added to the
	//      pod spec manually (not done by CNI) so container know where to
	//      retrieve data.
	//      Example: userspace/mappedDir: /var/lib/cni/usrspcni/
	KubeConfig string `json:"kubeconfig,omitempty"`
	SharedDir  string `json:"sharedDir,omitempty"`

	LogFile  string `json:"logFile,omitempty"`
	LogLevel string `json:"logLevel,omitempty"`

	Name          string        `json:"name"`
	HostConf      UserSpaceConf `json:"host,omitempty"`
	ContainerConf UserSpaceConf `json:"container,omitempty"`
}

// Defines the JSON data written to container. It is either written to:
//  1. Annotation - "userspace/configuration-data"
//     -- OR --
//  2. a file in the directory designated by NetConf.SharedDir.
type ConfigurationData struct {
	ContainerId string         `json:"containerId"` // From args.ContainerId, used locally. Used in several place, namely in the socket filenames.
	IfName      string         `json:"ifName"`      // From args.IfName, used locally. Used in several place, namely in the socket filenames.
	Name        string         `json:"name"`        // From NetConf.Name
	Config      UserSpaceConf  `json:"config"`      // From NetConf.ContainerConf
	IPResult    current.Result `json:"ipResult"`    // Network Status also has IP, but wrong format
}

const DefaultSwIfIndex = 4294967295 // vpp default interface id, used when querying bridges
