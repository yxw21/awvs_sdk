package awvs

import (
	"AM/lib/awvs/_struct"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type HttpReturn struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}

type aWvsSDK struct {
	ApiRoot string
	ApiKey  string
}

func (aWvsSDK *aWvsSDK) http(method string, apiPath string, body interface{}) (*HttpReturn, error) {
	var (
		client = http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
		req *http.Request
		err error
	)
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, aWvsSDK.ApiRoot+apiPath, bytes.NewBuffer(jsonBytes))
	} else {
		req, err = http.NewRequest(method, aWvsSDK.ApiRoot+apiPath, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("X-Auth", aWvsSDK.ApiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	httpReturn := &HttpReturn{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
	}
	if err == nil {
		httpReturn.Body = jsonBytes
	}
	return httpReturn, err
}

func (aWvsSDK *aWvsSDK) getLocationIdFromHeader(header http.Header) string {
	location := header.Get("location")
	if location != "" {
		arr := strings.Split(location, "/")
		if arrLen := len(arr); arrLen > 0 {
			return arr[arrLen-1]
		}
	}
	return ""
}

func (aWvsSDK *aWvsSDK) checkResponse(responseBytes []byte, err error, response interface{}) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (aWvsSDK *aWvsSDK) Stats() (*_struct.StatsResponse, error) {
	httpReturn, err := aWvsSDK.http("get", "/api/v1/me/stats", nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.StatsResponse{})
	if resp != nil {
		return resp.(*_struct.StatsResponse), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) AddGroup(request _struct.AddGroupRequest) (*_struct.AddGroupResponse, error) {
	httpReturn, err := aWvsSDK.http("post", "/api/v1/target_groups", request)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.AddGroupResponse{})
	if resp != nil {
		return resp.(*_struct.AddGroupResponse), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) DeleteGroup(request _struct.DeleteGroupRequest) (*_struct.Error, error) {
	httpReturn, err := aWvsSDK.http("delete", "/api/v1/target_groups/"+request.GroupId.String(), nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.Error{})
	if resp != nil {
		return resp.(*_struct.Error), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) DeleteGroups(request _struct.DeleteGroupsRequest) (*_struct.Error, error) {
	httpReturn, err := aWvsSDK.http("post", "/api/v1/target_groups/delete", request)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.Error{})
	if resp != nil {
		return resp.(*_struct.Error), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) UpdateGroup(request _struct.UpdateGroupRequest) (*_struct.Error, error) {
	httpReturn, err := aWvsSDK.http("patch", "/api/v1/target_groups/"+request.GroupId, request)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.Error{})
	if resp != nil {
		return resp.(*_struct.Error), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) QueryGroup(request _struct.QueryGroupRequest) (*_struct.QueryGroupResponse, error) {
	httpReturn, err := aWvsSDK.http("get", "/api/v1/target_groups/"+request.GroupId, nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.QueryGroupResponse{})
	if resp != nil {
		return resp.(*_struct.QueryGroupResponse), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) QueryGroups(request _struct.QueryGroupsRequest) (*_struct.QueryGroupsResponse, error) {
	httpReturn, err := aWvsSDK.http("get", "/api/v1/target_groups?l="+strconv.Itoa(request.L), nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.QueryGroupsResponse{})
	if resp != nil {
		return resp.(*_struct.QueryGroupsResponse), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) AddTargets(request _struct.AddTargetsRequest) (*_struct.AddTargetsResponse, error) {
	httpReturn, err := aWvsSDK.http("post", "/api/v1/targets/add", request)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.AddTargetsResponse{})
	if resp != nil {
		return resp.(*_struct.AddTargetsResponse), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) DeleteTargets(request _struct.DeleteTargetsRequest) (*_struct.Error, error) {
	httpReturn, err := aWvsSDK.http("delete", "/api/v1/targets/delete", request)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.Error{})
	if resp != nil {
		return resp.(*_struct.Error), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) QueryTarget(request _struct.QueryTargetRequest) (*_struct.QueryTargetResponse, error) {
	httpReturn, err := aWvsSDK.http("get", "/api/v1/targets/"+request.TargetId.String(), nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.QueryTargetResponse{})
	if resp != nil {
		return resp.(*_struct.QueryTargetResponse), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) AddScan(request _struct.AddScanRequest) ( /*scanId*/ string, *_struct.AddScanResponse, error) {
	httpReturn, err := aWvsSDK.http("post", "/api/v1/scans", request)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.AddScanResponse{})
	if resp != nil {
		return aWvsSDK.getLocationIdFromHeader(httpReturn.Header), resp.(*_struct.AddScanResponse), err
	}
	return "", nil, err
}

func (aWvsSDK *aWvsSDK) AbortScan(request _struct.AbortScanRequest) (*_struct.Error, error) {
	httpReturn, err := aWvsSDK.http("post", "/api/v1/scans/"+request.ScanId.String()+"/abort", nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.Error{})
	if resp != nil {
		return resp.(*_struct.Error), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) ResumeScan(request _struct.ResumeScanRequest) (*_struct.Error, error) {
	httpReturn, err := aWvsSDK.http("post", "/api/v1/scans/"+request.ScanId.String()+"/resume", nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.Error{})
	if resp != nil {
		return resp.(*_struct.Error), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) QueryScan(request _struct.QueryScanRequest) (*_struct.QueryScanResponse, error) {
	httpReturn, err := aWvsSDK.http("get", "/api/v1/scans/"+request.ScanId.String(), nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.QueryScanResponse{})
	if resp != nil {
		return resp.(*_struct.QueryScanResponse), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) AddReport(request _struct.AddReportRequest) ( /*reportId*/ string, *_struct.Error, error) {
	httpReturn, err := aWvsSDK.http("post", "/api/v1/reports", request)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.Error{})
	if resp != nil {
		return "", resp.(*_struct.Error), err
	}
	return aWvsSDK.getLocationIdFromHeader(httpReturn.Header), nil, err
}

func (aWvsSDK *aWvsSDK) QueryReport(request _struct.QueryReportRequest) (*_struct.QueryReportResponse, error) {
	httpReturn, err := aWvsSDK.http("get", "/api/v1/reports/"+request.ReportId, nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.QueryReportResponse{})
	if resp != nil {
		return resp.(*_struct.QueryReportResponse), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) QueryReports(request _struct.QueryReportsRequest) (*_struct.QueryReportsResponse, error) {
	httpReturn, err := aWvsSDK.http("get", "/api/v1/reports?l="+strconv.Itoa(request.L), nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.QueryReportsResponse{})
	if resp != nil {
		return resp.(*_struct.QueryReportsResponse), err
	}
	return nil, err
}

func (aWvsSDK *aWvsSDK) QueryEvents(request _struct.QueryEventsRequest) (*_struct.QueryEventsResponse, error) {
	httpReturn, err := aWvsSDK.http("get", "/api/v1/events?q=resource_type:6;resource_id:"+request.ResourceId.String()+";", nil)
	resp, err := aWvsSDK.checkResponse(httpReturn.Body, err, &_struct.QueryEventsResponse{})
	if resp != nil {
		return resp.(*_struct.QueryEventsResponse), err
	}
	return nil, err
}

func NewAWvsSDK(apiRoot, apiKey string) (*aWvsSDK, error) {
	if apiRoot == "" || apiKey == "" {
		return nil, errors.New("must provide apiroot and apikey")
	}
	return &aWvsSDK{
		ApiRoot: apiRoot,
		ApiKey:  apiKey,
	}, nil
}
