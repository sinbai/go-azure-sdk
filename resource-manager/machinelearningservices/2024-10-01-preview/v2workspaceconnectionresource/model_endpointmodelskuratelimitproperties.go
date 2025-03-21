package v2workspaceconnectionresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointModelSkuRateLimitProperties struct {
	Count         *float64                                   `json:"count,omitempty"`
	RenewalPeriod *float64                                   `json:"renewalPeriod,omitempty"`
	Rules         *[]EndpointModelSkuRateLimitRuleProperties `json:"rules,omitempty"`
}
