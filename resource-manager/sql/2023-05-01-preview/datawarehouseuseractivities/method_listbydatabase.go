package datawarehouseuseractivities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDatabaseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DataWarehouseUserActivities
}

type ListByDatabaseCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DataWarehouseUserActivities
}

// ListByDatabase ...
func (c DataWarehouseUserActivitiesClient) ListByDatabase(ctx context.Context, id commonids.SqlDatabaseId) (result ListByDatabaseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/dataWarehouseUserActivities", id.ID()),
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
		Values *[]DataWarehouseUserActivities `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDatabaseComplete retrieves all the results into a single object
func (c DataWarehouseUserActivitiesClient) ListByDatabaseComplete(ctx context.Context, id commonids.SqlDatabaseId) (ListByDatabaseCompleteResult, error) {
	return c.ListByDatabaseCompleteMatchingPredicate(ctx, id, DataWarehouseUserActivitiesOperationPredicate{})
}

// ListByDatabaseCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataWarehouseUserActivitiesClient) ListByDatabaseCompleteMatchingPredicate(ctx context.Context, id commonids.SqlDatabaseId, predicate DataWarehouseUserActivitiesOperationPredicate) (result ListByDatabaseCompleteResult, err error) {
	items := make([]DataWarehouseUserActivities, 0)

	resp, err := c.ListByDatabase(ctx, id)
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
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
