package longtermretentionmanagedinstancebackups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDatabaseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedInstanceLongTermRetentionBackup
}

type ListByDatabaseCompleteResult struct {
	Items []ManagedInstanceLongTermRetentionBackup
}

type ListByDatabaseOperationOptions struct {
	DatabaseState         *DatabaseState
	OnlyLatestPerDatabase *bool
}

func DefaultListByDatabaseOperationOptions() ListByDatabaseOperationOptions {
	return ListByDatabaseOperationOptions{}
}

func (o ListByDatabaseOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByDatabaseOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByDatabaseOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DatabaseState != nil {
		out.Append("databaseState", fmt.Sprintf("%v", *o.DatabaseState))
	}
	if o.OnlyLatestPerDatabase != nil {
		out.Append("onlyLatestPerDatabase", fmt.Sprintf("%v", *o.OnlyLatestPerDatabase))
	}
	return &out
}

// ListByDatabase ...
func (c LongTermRetentionManagedInstanceBackupsClient) ListByDatabase(ctx context.Context, id LongTermRetentionDatabaseId, options ListByDatabaseOperationOptions) (result ListByDatabaseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/longTermRetentionManagedInstanceBackups", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]ManagedInstanceLongTermRetentionBackup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDatabaseComplete retrieves all the results into a single object
func (c LongTermRetentionManagedInstanceBackupsClient) ListByDatabaseComplete(ctx context.Context, id LongTermRetentionDatabaseId, options ListByDatabaseOperationOptions) (ListByDatabaseCompleteResult, error) {
	return c.ListByDatabaseCompleteMatchingPredicate(ctx, id, options, ManagedInstanceLongTermRetentionBackupOperationPredicate{})
}

// ListByDatabaseCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LongTermRetentionManagedInstanceBackupsClient) ListByDatabaseCompleteMatchingPredicate(ctx context.Context, id LongTermRetentionDatabaseId, options ListByDatabaseOperationOptions, predicate ManagedInstanceLongTermRetentionBackupOperationPredicate) (result ListByDatabaseCompleteResult, err error) {
	items := make([]ManagedInstanceLongTermRetentionBackup, 0)

	resp, err := c.ListByDatabase(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByDatabaseCompleteResult{
		Items: items,
	}
	return
}
