package listkeys

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeysListByAutomationAccountOperationResponse struct {
	HttpResponse *http.Response
	Model        *KeyListResult
}

// KeysListByAutomationAccount ...
func (c ListKeysClient) KeysListByAutomationAccount(ctx context.Context, id AutomationAccountId) (result KeysListByAutomationAccountOperationResponse, err error) {
	req, err := c.preparerForKeysListByAutomationAccount(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "listkeys.ListKeysClient", "KeysListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "listkeys.ListKeysClient", "KeysListByAutomationAccount", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForKeysListByAutomationAccount(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "listkeys.ListKeysClient", "KeysListByAutomationAccount", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForKeysListByAutomationAccount prepares the KeysListByAutomationAccount request.
func (c ListKeysClient) preparerForKeysListByAutomationAccount(ctx context.Context, id AutomationAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listKeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForKeysListByAutomationAccount handles the response to the KeysListByAutomationAccount request. The method always
// closes the http.Response Body.
func (c ListKeysClient) responderForKeysListByAutomationAccount(resp *http.Response) (result KeysListByAutomationAccountOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
