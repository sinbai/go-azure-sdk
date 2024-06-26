package replicationprotectionclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type A2ASharedDiskIRErrorDetails struct {
	ErrorCode         *string `json:"errorCode,omitempty"`
	ErrorCodeEnum     *string `json:"errorCodeEnum,omitempty"`
	ErrorMessage      *string `json:"errorMessage,omitempty"`
	PossibleCauses    *string `json:"possibleCauses,omitempty"`
	RecommendedAction *string `json:"recommendedAction,omitempty"`
}
