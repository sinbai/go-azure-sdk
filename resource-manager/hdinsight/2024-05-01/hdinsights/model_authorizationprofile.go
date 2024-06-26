package hdinsights

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationProfile struct {
	GroupIds *[]string `json:"groupIds,omitempty"`
	UserIds  *[]string `json:"userIds,omitempty"`
}
