package yundun_dbaudit

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

// DescribeInstanceAttribute invokes the yundun_dbaudit.DescribeInstanceAttribute API synchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describeinstanceattribute.html
func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (response *DescribeInstanceAttributeResponse, err error) {
	response = CreateDescribeInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceAttributeWithChan invokes the yundun_dbaudit.DescribeInstanceAttribute API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describeinstanceattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceAttributeWithChan(request *DescribeInstanceAttributeRequest) (<-chan *DescribeInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceAttribute(request)
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

// DescribeInstanceAttributeWithCallback invokes the yundun_dbaudit.DescribeInstanceAttribute API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describeinstanceattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceAttributeWithCallback(request *DescribeInstanceAttributeRequest, callback func(response *DescribeInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceAttribute(request)
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

// DescribeInstanceAttributeRequest is the request struct for api DescribeInstanceAttribute
type DescribeInstanceAttributeRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// DescribeInstanceAttributeResponse is the response struct for api DescribeInstanceAttribute
type DescribeInstanceAttributeResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	InstanceAttribute InstanceAttribute `json:"InstanceAttribute" xml:"InstanceAttribute"`
}

// CreateDescribeInstanceAttributeRequest creates a request to invoke DescribeInstanceAttribute API
func CreateDescribeInstanceAttributeRequest() (request *DescribeInstanceAttributeRequest) {
	request = &DescribeInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-dbaudit", "2018-10-29", "DescribeInstanceAttribute", "dbaudit", "openAPI")
	return
}

// CreateDescribeInstanceAttributeResponse creates a response to parse from DescribeInstanceAttribute response
func CreateDescribeInstanceAttributeResponse() (response *DescribeInstanceAttributeResponse) {
	response = &DescribeInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}