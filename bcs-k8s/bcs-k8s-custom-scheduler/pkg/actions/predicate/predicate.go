/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package predicate

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-custom-scheduler/pkg/actions"
	"github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-custom-scheduler/pkg/ipscheduler"
	"github.com/emicklei/go-restful"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

const (
	PredicatePrefix = "predicate"
)

func init() {
	actions.RegisterAction(actions.Action{Verb: "POST", Path: actions.BcsCustomSchedulerPrefix + "/" + "ipscheduler" + "/" + PredicatePrefix, Params: nil, Handler: handleIpschedulerPredicate})

}

func handleIpschedulerPredicate(req *restful.Request, resp *restful.Response) {

	var extenderArgs schedulerapi.ExtenderArgs
	var extenderFilterResult *schedulerapi.ExtenderFilterResult
	err := req.ReadEntity(&extenderArgs)
	if err != nil {
		blog.Infof("error when read request: %s", err.Error())
		extenderFilterResult = &schedulerapi.ExtenderFilterResult{
			Nodes:       nil,
			FailedNodes: nil,
			Error:       err.Error(),
		}

		resp.WriteEntity(extenderFilterResult)
		return
	}
	extenderFilterResult, err = ipscheduler.HandleIpschedulerPredicate(extenderArgs)
	if err != nil {
		extenderFilterResult = &schedulerapi.ExtenderFilterResult{
			Nodes:       nil,
			FailedNodes: nil,
			Error:       err.Error(),
		}
	}

	resp.WriteEntity(extenderFilterResult)
}
