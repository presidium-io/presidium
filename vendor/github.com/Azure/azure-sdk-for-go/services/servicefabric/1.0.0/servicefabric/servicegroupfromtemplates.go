package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ServiceGroupFromTemplatesClient is the client for the ServiceGroupFromTemplates methods of the Servicefabric
// service.
type ServiceGroupFromTemplatesClient struct {
	BaseClient
}

// NewServiceGroupFromTemplatesClient creates an instance of the ServiceGroupFromTemplatesClient client.
func NewServiceGroupFromTemplatesClient(timeout *int32) ServiceGroupFromTemplatesClient {
	return NewServiceGroupFromTemplatesClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewServiceGroupFromTemplatesClientWithBaseURI creates an instance of the ServiceGroupFromTemplatesClient client.
func NewServiceGroupFromTemplatesClientWithBaseURI(baseURI string, timeout *int32) ServiceGroupFromTemplatesClient {
	return ServiceGroupFromTemplatesClient{NewWithBaseURI(baseURI, timeout)}
}

// Create create service group from templates
//
// applicationName is the name of the application serviceDescriptionTemplate is the template of the service
// description
func (client ServiceGroupFromTemplatesClient) Create(ctx context.Context, applicationName string, serviceDescriptionTemplate ServiceDescriptionTemplate) (result String, err error) {
	req, err := client.CreatePreparer(ctx, applicationName, serviceDescriptionTemplate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupFromTemplatesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupFromTemplatesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupFromTemplatesClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ServiceGroupFromTemplatesClient) CreatePreparer(ctx context.Context, applicationName string, serviceDescriptionTemplate ServiceDescriptionTemplate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}/$/GetServiceGroups/$/CreateServiceGroupFromTemplate", pathParameters),
		autorest.WithJSON(serviceDescriptionTemplate),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceGroupFromTemplatesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ServiceGroupFromTemplatesClient) CreateResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
