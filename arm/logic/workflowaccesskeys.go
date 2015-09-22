package logic

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// WorkflowAccessKeys Client
type WorkflowAccessKeysClient struct {
	LogicManagementClient
}

func NewWorkflowAccessKeysClient(subscriptionId string) WorkflowAccessKeysClient {
	return NewWorkflowAccessKeysClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewWorkflowAccessKeysClientWithBaseUri(baseUri string, subscriptionId string) WorkflowAccessKeysClient {
	return WorkflowAccessKeysClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// List gets a list of workflow access keys.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. top is the number of items to be included in the result.
func (client WorkflowAccessKeysClient) List(resourceGroupName string, workflowName string, top int) (result WorkflowAccessKeyListResult, ae autorest.Error) {
	req, err := client.NewListRequest(resourceGroupName, workflowName, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "List", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client WorkflowAccessKeysClient) NewListRequest(resourceGroupName string, workflowName string, top int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"$top":        top,
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client WorkflowAccessKeysClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys"))
}

// Get gets a workflow access key.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. accessKeyName is the workflow access key name.
func (client WorkflowAccessKeysClient) Get(resourceGroupName string, workflowName string, accessKeyName string) (result WorkflowAccessKey, ae autorest.Error) {
	req, err := client.NewGetRequest(resourceGroupName, workflowName, accessKeyName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Get", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Get", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Get", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Get", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Get request.
func (client WorkflowAccessKeysClient) NewGetRequest(resourceGroupName string, workflowName string, accessKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     url.QueryEscape(accessKeyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.GetRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Get request.
func (client WorkflowAccessKeysClient) GetRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}"))
}

// CreateOrUpdate creates or updates a workflow access key.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. accessKeyName is the workflow access key name. workflowAccesskey is
// the workflow access key.
func (client WorkflowAccessKeysClient) CreateOrUpdate(resourceGroupName string, workflowName string, accessKeyName string, workflowAccesskey WorkflowAccessKey) (result WorkflowAccessKey, ae autorest.Error) {
	req, err := client.NewCreateOrUpdateRequest(resourceGroupName, workflowName, accessKeyName, workflowAccesskey)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "CreateOrUpdate", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "CreateOrUpdate", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "CreateOrUpdate", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the CreateOrUpdate request.
func (client WorkflowAccessKeysClient) NewCreateOrUpdateRequest(resourceGroupName string, workflowName string, accessKeyName string, workflowAccesskey WorkflowAccessKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     url.QueryEscape(accessKeyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CreateOrUpdateRequestPreparer(),
		autorest.WithJSON(workflowAccesskey),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CreateOrUpdate request.
func (client WorkflowAccessKeysClient) CreateOrUpdateRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}"))
}

// Delete deletes a workflow access key.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. accessKeyName is the workflow access key name.
func (client WorkflowAccessKeysClient) Delete(resourceGroupName string, workflowName string, accessKeyName string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewDeleteRequest(resourceGroupName, workflowName, accessKeyName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Delete", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Delete", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Delete", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Delete", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the Delete request.
func (client WorkflowAccessKeysClient) NewDeleteRequest(resourceGroupName string, workflowName string, accessKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     url.QueryEscape(accessKeyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.DeleteRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Delete request.
func (client WorkflowAccessKeysClient) DeleteRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}"))
}

// ListSecretKeys lists secret keys.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. accessKeyName is the workflow access key name.
func (client WorkflowAccessKeysClient) ListSecretKeys(resourceGroupName string, workflowName string, accessKeyName string) (result WorkflowSecretKeys, ae autorest.Error) {
	req, err := client.NewListSecretKeysRequest(resourceGroupName, workflowName, accessKeyName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "ListSecretKeys", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "ListSecretKeys", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "ListSecretKeys", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "ListSecretKeys", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the ListSecretKeys request.
func (client WorkflowAccessKeysClient) NewListSecretKeysRequest(resourceGroupName string, workflowName string, accessKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     url.QueryEscape(accessKeyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListSecretKeysRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the ListSecretKeys request.
func (client WorkflowAccessKeysClient) ListSecretKeysRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}/list"))
}

// RegenerateSecretKey regenerates secret key.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. accessKeyName is the workflow access key name. parameters is the
// parameters.
func (client WorkflowAccessKeysClient) RegenerateSecretKey(resourceGroupName string, workflowName string, accessKeyName string, parameters RegenerateSecretKeyParameters) (result WorkflowSecretKeys, ae autorest.Error) {
	req, err := client.NewRegenerateSecretKeyRequest(resourceGroupName, workflowName, accessKeyName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "RegenerateSecretKey", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "RegenerateSecretKey", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "RegenerateSecretKey", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "RegenerateSecretKey", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the RegenerateSecretKey request.
func (client WorkflowAccessKeysClient) NewRegenerateSecretKeyRequest(resourceGroupName string, workflowName string, accessKeyName string, parameters RegenerateSecretKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     url.QueryEscape(accessKeyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.RegenerateSecretKeyRequestPreparer(),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the RegenerateSecretKey request.
func (client WorkflowAccessKeysClient) RegenerateSecretKeyRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}/regenerate"))
}