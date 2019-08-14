// Copyright 2019 Intel Corp.
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

package usrspcni

import (
	v1 "k8s.io/api/core/v1"

	"github.com/containernetworking/cni/pkg/skel"
	_ "github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"

	"github.com/intel/userspace-cni-network-plugin/usrsptypes"
	"github.com/intel/userspace-cni-network-plugin/k8sclient"
)

//
// Exported Types
//
type UsrSpCni interface {
	AddOnHost(conf *usrsptypes.NetConf,
			  args *skel.CmdArgs,
			  kubeClient k8sclient.KubeClient,
			  sharedDir string,
			  ipResult *current.Result) error
	AddOnContainer(conf *usrsptypes.NetConf,
				   args *skel.CmdArgs,
				   kubeClient k8sclient.KubeClient,
				   sharedDir string,
				   pod *v1.Pod,
				   ipResult *current.Result) (*v1.Pod, error)
	DelFromHost(conf *usrsptypes.NetConf,
				args *skel.CmdArgs,
				sharedDir string) error
	DelFromContainer(conf *usrsptypes.NetConf,
		             args *skel.CmdArgs,
		             sharedDir string,
		             pod *v1.Pod) error
}
