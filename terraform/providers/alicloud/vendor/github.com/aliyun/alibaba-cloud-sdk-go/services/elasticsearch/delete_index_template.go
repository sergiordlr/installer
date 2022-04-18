package elasticsearch

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

// DeleteIndexTemplate invokes the elasticsearch.DeleteIndexTemplate API synchronously
func (client *Client) DeleteIndexTemplate(request *DeleteIndexTemplateRequest) (response *DeleteIndexTemplateResponse, err error) {
	response = CreateDeleteIndexTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteIndexTemplateWithChan invokes the elasticsearch.DeleteIndexTemplate API asynchronously
func (client *Client) DeleteIndexTemplateWithChan(request *DeleteIndexTemplateRequest) (<-chan *DeleteIndexTemplateResponse, <-chan error) {
	responseChan := make(chan *DeleteIndexTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteIndexTemplate(request)
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

// DeleteIndexTemplateWithCallback invokes the elasticsearch.DeleteIndexTemplate API asynchronously
func (client *Client) DeleteIndexTemplateWithCallback(request *DeleteIndexTemplateRequest, callback func(response *DeleteIndexTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteIndexTemplateResponse
		var err error
		defer close(result)
		response, err = client.DeleteIndexTemplate(request)
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

// DeleteIndexTemplateRequest is the request struct for api DeleteIndexTemplate
type DeleteIndexTemplateRequest struct {
	*requests.RoaRequest
	InstanceId    string `position:"Path" name:"InstanceId"`
	IndexTemplate string `position:"Path" name:"IndexTemplate"`
}

// DeleteIndexTemplateResponse is the response struct for api DeleteIndexTemplate
type DeleteIndexTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateDeleteIndexTemplateRequest creates a request to invoke DeleteIndexTemplate API
func CreateDeleteIndexTemplateRequest() (request *DeleteIndexTemplateRequest) {
	request = &DeleteIndexTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DeleteIndexTemplate", "/openapi/instances/[InstanceId]/index-templates/[IndexTemplate]", "elasticsearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteIndexTemplateResponse creates a response to parse from DeleteIndexTemplate response
func CreateDeleteIndexTemplateResponse() (response *DeleteIndexTemplateResponse) {
	response = &DeleteIndexTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}