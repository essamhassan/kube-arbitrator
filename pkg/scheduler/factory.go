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

package scheduler

import (
	"github.com/kubernetes-incubator/kube-arbitrator/pkg/scheduler/actions/allocate"
	"github.com/kubernetes-incubator/kube-arbitrator/pkg/scheduler/actions/preempt"
	"github.com/kubernetes-incubator/kube-arbitrator/pkg/scheduler/actions/reclaim"

	"github.com/kubernetes-incubator/kube-arbitrator/pkg/scheduler/plugins/drf"
	"github.com/kubernetes-incubator/kube-arbitrator/pkg/scheduler/plugins/gang"
	"github.com/kubernetes-incubator/kube-arbitrator/pkg/scheduler/plugins/hostport"
	"github.com/kubernetes-incubator/kube-arbitrator/pkg/scheduler/plugins/nodeaffinity"
	"github.com/kubernetes-incubator/kube-arbitrator/pkg/scheduler/plugins/priority"
	"github.com/kubernetes-incubator/kube-arbitrator/pkg/scheduler/plugins/proportion"

	"github.com/kubernetes-incubator/kube-arbitrator/pkg/scheduler/framework"
)

func init() {
	// Plugins for Jobs
	framework.RegisterPluginBuilder("drf", drf.New)
	framework.RegisterPluginBuilder("gang", gang.New)
	framework.RegisterPluginBuilder("hostport", hostport.New)
	framework.RegisterPluginBuilder("nodeaffinity", nodeaffinity.New)
	framework.RegisterPluginBuilder("priority", priority.New)

	// Plugins for Queues
	framework.RegisterPluginBuilder("proportion", proportion.New)

	// Actions
	framework.RegisterAction(reclaim.New())
	framework.RegisterAction(allocate.New())
	framework.RegisterAction(preempt.New())
}
