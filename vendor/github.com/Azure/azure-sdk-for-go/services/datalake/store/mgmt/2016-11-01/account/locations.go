package account

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

// LocationsClient is the creates an Azure Data Lake Store account management client.
type LocationsClient struct {
	BaseClient
}

// NewLocationsClient creates an instance of the LocationsClient client.
func NewLocationsClient(subscriptionID string) LocationsClient {
	return NewLocationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationsClientWithBaseURI creates an instance of the LocationsClient client.
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return LocationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetCapability gets subscription-level properties and limits for Data Lake Store specified by resource location.
//
// location is the resource location without whitespace.
func (client LocationsClient) GetCapability(ctx context.Context, location string) (result CapabilityInformation, err error) {
	req, err := client.GetCapabilityPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.LocationsClient", "GetCapability", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCapabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "account.LocationsClient", "GetCapability", resp, "Failure sending request")
		return
	}

	result, err = client.GetCapabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.LocationsClient", "GetCapability", resp, "Failure responding to request")
	}

	return
}

// GetCapabilityPreparer prepares the GetCapability request.
func (client LocationsClient) GetCapabilityPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/capability", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetCapabilitySender sends the GetCapability request. The method will close the
// http.Response Body if it receives an error.
func (client LocationsClient) GetCapabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetCapabilityResponder handles the response to the GetCapability request. The method always
// closes the http.Response Body.
func (client LocationsClient) GetCapabilityResponder(resp *http.Response) (result CapabilityInformation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
