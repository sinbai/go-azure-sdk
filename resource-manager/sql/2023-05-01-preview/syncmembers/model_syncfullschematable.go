package syncmembers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncFullSchemaTable struct {
	Columns    *[]SyncFullSchemaTableColumn `json:"columns,omitempty"`
	ErrorId    *string                      `json:"errorId,omitempty"`
	HasError   *bool                        `json:"hasError,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	QuotedName *string                      `json:"quotedName,omitempty"`
}
