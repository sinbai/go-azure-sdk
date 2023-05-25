package networksecurityperimeterconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NSPProvisioningIssue struct {
	Name       *string                         `json:"name,omitempty"`
	Properties *NSPProvisioningIssueProperties `json:"properties,omitempty"`
}
