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

package mesos

import (
	"fmt"
	"io/ioutil"
	"strings"

	"bk-bcs/bcs-common/common"
	"bk-bcs/bcs-common/common/blog"
	bhttp "bk-bcs/bcs-common/common/http"
	"bk-bcs/bcs-common/common/http/httpclient"
	"bk-bcs/bcs-common/common/types"
	"bk-bcs/bcs-services/bcs-api/processor/http/actions"
	"bk-bcs/bcs-services/bcs-api/regdiscv"

	"github.com/emicklei/go-restful"
)

const (
	BcsApiPrefix = "/bcsapi/v4/scheduler/mesos/"
)

func init() {
	actions.RegisterAction(actions.Action{Verb: "POST", Path: "/bcsapi/v4/scheduler/mesos/{uri:*}", Params: nil, Handler: handlerPostActions})
	actions.RegisterAction(actions.Action{Verb: "PUT", Path: "/bcsapi/v4/scheduler/mesos/{uri:*}", Params: nil, Handler: handlerPutActions})
	actions.RegisterAction(actions.Action{Verb: "GET", Path: "/bcsapi/v4/scheduler/mesos/{uri:*}", Params: nil, Handler: handlerGetActions})
	actions.RegisterAction(actions.Action{Verb: "DELETE", Path: "/bcsapi/v4/scheduler/mesos/{uri:*}", Params: nil, Handler: handlerDeleteActions})
}

func request2mesosapi(req *restful.Request, uri, method string) (string, error) {
	data, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		blog.Error("handler url %s read request body failed, error: %s", uri, err.Error())
		err1 := bhttp.InternalError(common.BcsErrCommHttpReadBodyFail, common.BcsErrCommHttpReadBodyFailStr)
		return err1.Error(), nil
	}

	cluster := req.Request.Header.Get("BCS-ClusterID")
	if cluster == "" {
		blog.Error("handler url %s read header BCS-ClusterID is empty", uri)
		err1 := bhttp.InternalError(common.BcsErrCommHttpParametersFailed, "http header BCS-ClusterID can't be empty")
		return err1.Error(), nil
	}

	rd, err := regdiscv.GetRDiscover()
	if err != nil {
		blog.Error("hander url %s get RDiscover error %s", uri, err.Error())
		err1 := bhttp.InternalError(common.BcsErrApiInternalFail, common.BcsErrApiInternalFailStr)
		return err1.Error(), nil
	}

	serv, err := rd.GetModuleServers(fmt.Sprintf("%s/%s", types.BCS_MODULE_MESOSAPISERVER, cluster))
	if err != nil {
		blog.Error("get cluster %s servers %s error %s", cluster, types.BCS_MODULE_MESOSAPISERVER, err.Error())
		err1 := bhttp.InternalError(common.BcsErrApiGetMesosApiFail, fmt.Sprintf("mesos cluster %s not found", cluster))
		return err1.Error(), nil
	}

	ser, ok := serv.(*types.BcsMesosApiserverInfo)
	if !ok {
		blog.Errorf("servers convert to BcsMesosApiserverInfo")
		err1 := bhttp.InternalError(common.BcsErrApiGetMesosApiFail, common.BcsErrApiGetMesosApiFailStr)
		return err1.Error(), nil
	}

	//host := servInfo.Scheme + "://" + servInfo.IP + ":" + strconv.Itoa(int(servInfo.Port))
	host := fmt.Sprintf("%s://%s:%d", ser.Scheme, ser.IP, ser.Port)
	//url := routeHost + "/api/v1/" + uri //a.Conf.BcsRoute
	url := fmt.Sprintf("%s/mesosdriver/v4/%s", host, uri)
	blog.V(3).Infof("do request to url(%s), method(%s)", url, method)

	httpcli := httpclient.NewHttpClient()
	httpcli.SetHeader("Content-Type", "application/json")
	httpcli.SetHeader("Accept", "application/json")
	if strings.ToLower(ser.Scheme) == "https" {
		cliTls, err := rd.GetClientTls()
		if err != nil {
			blog.Errorf("get client tls error %s", err.Error())
		}
		httpcli.SetTlsVerityConfig(cliTls)
	}

	reply, err := httpcli.Request(url, method, req.Request.Header, data)
	if err != nil {
		blog.Error("request url %s error %s", url, err.Error())
		err1 := bhttp.InternalError(common.BcsErrApiRequestMesosApiFail, common.BcsErrApiRequestMesosApiFailStr)
		return err1.Error(), nil
	}

	return string(reply), err
}

func handlerPostActions(req *restful.Request, resp *restful.Response) {
	blog.V(3).Infof("client %s request %s", req.Request.RemoteAddr, req.Request.URL.Path)

	url := strings.Replace(req.Request.URL.Path, BcsApiPrefix, "", 1)

	if req.Request.URL.RawQuery != "" {
		url = fmt.Sprintf("%s?%s", url, req.Request.URL.RawQuery)
	}

	data, _ := request2mesosapi(req, url, "POST")
	resp.Write([]byte(data))
}

func handlerGetActions(req *restful.Request, resp *restful.Response) {
	blog.V(3).Infof("client %s request %s", req.Request.RemoteAddr, req.Request.URL.Path)
	url := strings.Replace(req.Request.URL.Path, BcsApiPrefix, "", 1)

	if req.Request.URL.RawQuery != "" {
		url = fmt.Sprintf("%s?%s", url, req.Request.URL.RawQuery)
	}

	data, _ := request2mesosapi(req, url, "GET")
	resp.Write([]byte(data))
}

func handlerDeleteActions(req *restful.Request, resp *restful.Response) {
	blog.V(3).Infof("client %s request %s", req.Request.RemoteAddr, req.Request.URL.Path)
	url := strings.Replace(req.Request.URL.Path, BcsApiPrefix, "", 1)

	if req.Request.URL.RawQuery != "" {
		url = fmt.Sprintf("%s?%s", url, req.Request.URL.RawQuery)
	}

	data, _ := request2mesosapi(req, url, "DELETE")
	resp.Write([]byte(data))
}

func handlerPutActions(req *restful.Request, resp *restful.Response) {
	blog.V(3).Infof("client %s request %s", req.Request.RemoteAddr, req.Request.URL.Path)
	url := strings.Replace(req.Request.URL.Path, BcsApiPrefix, "", 1)

	if req.Request.URL.RawQuery != "" {
		url = fmt.Sprintf("%s?%s", url, req.Request.URL.RawQuery)
	}

	data, _ := request2mesosapi(req, url, "PUT")
	resp.Write([]byte(data))
}
