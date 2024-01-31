package sensitivitylabels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabel struct {
	Id         *string                     `json:"id,omitempty"`
	ManagedBy  *string                     `json:"managedBy,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Properties *SensitivityLabelProperties `json:"properties,omitempty"`
	Type       *string                     `json:"type,omitempty"`
}
