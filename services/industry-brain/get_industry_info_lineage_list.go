package industry_brain

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetIndustryInfoLineageList invokes the industry_brain.GetIndustryInfoLineageList API synchronously
// api document: https://help.aliyun.com/api/industry-brain/getindustryinfolineagelist.html
func (client *Client) GetIndustryInfoLineageList(request *GetIndustryInfoLineageListRequest) (response *GetIndustryInfoLineageListResponse, err error) {
	response = CreateGetIndustryInfoLineageListResponse()
	err = client.DoAction(request, response)
	return
}

// GetIndustryInfoLineageListWithChan invokes the industry_brain.GetIndustryInfoLineageList API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getindustryinfolineagelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetIndustryInfoLineageListWithChan(request *GetIndustryInfoLineageListRequest) (<-chan *GetIndustryInfoLineageListResponse, <-chan error) {
	responseChan := make(chan *GetIndustryInfoLineageListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetIndustryInfoLineageList(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetIndustryInfoLineageListWithCallback invokes the industry_brain.GetIndustryInfoLineageList API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getindustryinfolineagelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetIndustryInfoLineageListWithCallback(request *GetIndustryInfoLineageListRequest, callback func(response *GetIndustryInfoLineageListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetIndustryInfoLineageListResponse
		var err error
		defer close(result)
		response, err = client.GetIndustryInfoLineageList(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetIndustryInfoLineageListRequest is the request struct for api GetIndustryInfoLineageList
type GetIndustryInfoLineageListRequest struct {
	*requests.RpcRequest
	IndustryCode string `position:"Query" name:"IndustryCode"`
}

// GetIndustryInfoLineageListResponse is the response struct for api GetIndustryInfoLineageList
type GetIndustryInfoLineageListResponse struct {
	*responses.BaseResponse
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	IndustryInfoList []IndustryInfoListItem `json:"IndustryInfoList" xml:"IndustryInfoList"`
}

// CreateGetIndustryInfoLineageListRequest creates a request to invoke GetIndustryInfoLineageList API
func CreateGetIndustryInfoLineageListRequest() (request *GetIndustryInfoLineageListRequest) {
	request = &GetIndustryInfoLineageListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("industry-brain", "2018-07-12", "GetIndustryInfoLineageList", "", "")
	return
}

// CreateGetIndustryInfoLineageListResponse creates a response to parse from GetIndustryInfoLineageList response
func CreateGetIndustryInfoLineageListResponse() (response *GetIndustryInfoLineageListResponse) {
	response = &GetIndustryInfoLineageListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
