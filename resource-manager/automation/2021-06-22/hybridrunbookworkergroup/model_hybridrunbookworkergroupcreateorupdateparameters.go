package hybridrunbookworkergroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridRunbookWorkerGroupCreateOrUpdateParameters struct {
	Credential *RunAsCredentialAssociationProperty `json:"credential,omitempty"`
}
