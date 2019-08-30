package ros

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

// ListStackEvents invokes the ros.ListStackEvents API synchronously
// api document: https://help.aliyun.com/api/ros/liststackevents.html
func (client *Client) ListStackEvents(request *ListStackEventsRequest) (response *ListStackEventsResponse, err error) {
	response = CreateListStackEventsResponse()
	err = client.DoAction(request, response)
	return
}

// ListStackEventsWithChan invokes the ros.ListStackEvents API asynchronously
// api document: https://help.aliyun.com/api/ros/liststackevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListStackEventsWithChan(request *ListStackEventsRequest) (<-chan *ListStackEventsResponse, <-chan error) {
	responseChan := make(chan *ListStackEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListStackEvents(request)
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

// ListStackEventsWithCallback invokes the ros.ListStackEvents API asynchronously
// api document: https://help.aliyun.com/api/ros/liststackevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListStackEventsWithCallback(request *ListStackEventsRequest, callback func(response *ListStackEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListStackEventsResponse
		var err error
		defer close(result)
		response, err = client.ListStackEvents(request)
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

// ListStackEventsRequest is the request struct for api ListStackEvents
type ListStackEventsRequest struct {
	*requests.RpcRequest
	StackId           string           `position:"Query" name:"StackId"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	LogicalResourceId *[]string        `position:"Query" name:"LogicalResourceId"  type:"Repeated"`
	ResourceType      *[]string        `position:"Query" name:"ResourceType"  type:"Repeated"`
	PageNumber        requests.Integer `position:"Query" name:"PageNumber"`
	Status            *[]string        `position:"Query" name:"Status"  type:"Repeated"`
}

// ListStackEventsResponse is the response struct for api ListStackEvents
type ListStackEventsResponse struct {
	*responses.BaseResponse
	PageNumber int     `json:"PageNumber" xml:"PageNumber"`
	PageSize   int     `json:"PageSize" xml:"PageSize"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	TotalCount int     `json:"TotalCount" xml:"TotalCount"`
	Events     []Event `json:"Events" xml:"Events"`
}

// CreateListStackEventsRequest creates a request to invoke ListStackEvents API
func CreateListStackEventsRequest() (request *ListStackEventsRequest) {
	request = &ListStackEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ROS", "2019-09-10", "ListStackEvents", "ROS", "openAPI")
	return
}

// CreateListStackEventsResponse creates a response to parse from ListStackEvents response
func CreateListStackEventsResponse() (response *ListStackEventsResponse) {
	response = &ListStackEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}