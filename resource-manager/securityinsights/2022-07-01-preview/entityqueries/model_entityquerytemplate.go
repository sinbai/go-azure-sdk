package entityqueries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityQueryTemplate interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEntityQueryTemplateImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalEntityQueryTemplateImplementation(input []byte) (EntityQueryTemplate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EntityQueryTemplate into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Activity") {
		var out ActivityEntityQueryTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityEntityQueryTemplate: %+v", err)
		}
		return out, nil
	}

	out := RawEntityQueryTemplateImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
