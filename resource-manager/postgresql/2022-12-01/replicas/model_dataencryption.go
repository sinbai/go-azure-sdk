package replicas

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataEncryption struct {
	PrimaryKeyURI                 *string           `json:"primaryKeyURI,omitempty"`
	PrimaryUserAssignedIdentityId *string           `json:"primaryUserAssignedIdentityId,omitempty"`
	Type                          *ArmServerKeyType `json:"type,omitempty"`
}
