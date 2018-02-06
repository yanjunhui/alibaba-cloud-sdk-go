package ecs

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

func (client *Client) DescribeBandwidthLimitation(request *DescribeBandwidthLimitationRequest) (response *DescribeBandwidthLimitationResponse, err error) {
	response = CreateDescribeBandwidthLimitationResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeBandwidthLimitationWithChan(request *DescribeBandwidthLimitationRequest) (<-chan *DescribeBandwidthLimitationResponse, <-chan error) {
	responseChan := make(chan *DescribeBandwidthLimitationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBandwidthLimitation(request)
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

func (client *Client) DescribeBandwidthLimitationWithCallback(request *DescribeBandwidthLimitationRequest, callback func(response *DescribeBandwidthLimitationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBandwidthLimitationResponse
		var err error
		defer close(result)
		response, err = client.DescribeBandwidthLimitation(request)
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

type DescribeBandwidthLimitationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	InstanceChargeType   string           `position:"Query" name:"InstanceChargeType"`
	ResourceId           string           `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OperationType        string           `position:"Query" name:"OperationType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SpotStrategy         string           `position:"Query" name:"SpotStrategy"`
}

type DescribeBandwidthLimitationResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Bandwidths struct {
		Bandwidth []struct {
			InternetChargeType string `json:"InternetChargeType" xml:"InternetChargeType"`
			Min                int    `json:"Min" xml:"Min"`
			Max                int    `json:"Max" xml:"Max"`
			Unit               string `json:"Unit" xml:"Unit"`
		} `json:"Bandwidth" xml:"Bandwidth"`
	} `json:"Bandwidths" xml:"Bandwidths"`
}

func CreateDescribeBandwidthLimitationRequest() (request *DescribeBandwidthLimitationRequest) {
	request = &DescribeBandwidthLimitationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeBandwidthLimitation", "ecs", "openAPI")
	return
}

func CreateDescribeBandwidthLimitationResponse() (response *DescribeBandwidthLimitationResponse) {
	response = &DescribeBandwidthLimitationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
