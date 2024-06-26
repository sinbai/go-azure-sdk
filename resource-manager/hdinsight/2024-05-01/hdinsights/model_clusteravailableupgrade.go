package hdinsights

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterAvailableUpgrade struct {
	Id         *string                           `json:"id,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Properties ClusterAvailableUpgradeProperties `json:"properties"`
	SystemData *systemdata.SystemData            `json:"systemData,omitempty"`
	Type       *string                           `json:"type,omitempty"`
}

var _ json.Unmarshaler = &ClusterAvailableUpgrade{}

func (s *ClusterAvailableUpgrade) UnmarshalJSON(bytes []byte) error {
	type alias ClusterAvailableUpgrade
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ClusterAvailableUpgrade: %+v", err)
	}

	s.Id = decoded.Id
	s.Name = decoded.Name
	s.SystemData = decoded.SystemData
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ClusterAvailableUpgrade into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalClusterAvailableUpgradePropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'ClusterAvailableUpgrade': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
