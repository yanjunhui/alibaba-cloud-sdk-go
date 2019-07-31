package bssopenapi

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

// SubscribeBillToOSS invokes the bssopenapi.SubscribeBillToOSS API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/subscribebilltooss.html
func (client *Client) SubscribeBillToOSS(request *SubscribeBillToOSSRequest) (response *SubscribeBillToOSSResponse, err error) {
	response = CreateSubscribeBillToOSSResponse()
	err = client.DoAction(request, response)
	return
}

// SubscribeBillToOSSWithChan invokes the bssopenapi.SubscribeBillToOSS API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/subscribebilltooss.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubscribeBillToOSSWithChan(request *SubscribeBillToOSSRequest) (<-chan *SubscribeBillToOSSResponse, <-chan error) {
	responseChan := make(chan *SubscribeBillToOSSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubscribeBillToOSS(request)
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

// SubscribeBillToOSSWithCallback invokes the bssopenapi.SubscribeBillToOSS API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/subscribebilltooss.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubscribeBillToOSSWithCallback(request *SubscribeBillToOSSRequest, callback func(response *SubscribeBillToOSSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubscribeBillToOSSResponse
		var err error
		defer close(result)
		response, err = client.SubscribeBillToOSS(request)
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

// SubscribeBillToOSSRequest is the request struct for api SubscribeBillToOSS
type SubscribeBillToOSSRequest struct {
	*requests.RpcRequest
	SubscribeType   string `position:"Query" name:"SubscribeType"`
	SubscribeBucket string `position:"Query" name:"SubscribeBucket"`
}

// SubscribeBillToOSSResponse is the response struct for api SubscribeBillToOSS
type SubscribeBillToOSSResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSubscribeBillToOSSRequest creates a request to invoke SubscribeBillToOSS API
func CreateSubscribeBillToOSSRequest() (request *SubscribeBillToOSSRequest) {
	request = &SubscribeBillToOSSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "SubscribeBillToOSS", "bssopenapi", "openAPI")
	return
}

// CreateSubscribeBillToOSSResponse creates a response to parse from SubscribeBillToOSS response
func CreateSubscribeBillToOSSResponse() (response *SubscribeBillToOSSResponse) {
	response = &SubscribeBillToOSSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
