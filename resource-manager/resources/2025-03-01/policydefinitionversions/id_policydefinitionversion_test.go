package policydefinitionversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyDefinitionVersionId{}

func TestNewPolicyDefinitionVersionID(t *testing.T) {
	id := NewPolicyDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policyDefinitionName", "versionName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.PolicyDefinitionName != "policyDefinitionName" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyDefinitionName'", id.PolicyDefinitionName, "policyDefinitionName")
	}

	if id.VersionName != "versionName" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "versionName")
	}
}

func TestFormatPolicyDefinitionVersionID(t *testing.T) {
	actual := NewPolicyDefinitionVersionID("12345678-1234-9876-4563-123456789012", "policyDefinitionName", "versionName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions/versionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyDefinitionVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyDefinitionVersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions/versionName",
			Expected: &PolicyDefinitionVersionId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				PolicyDefinitionName: "policyDefinitionName",
				VersionName:          "versionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions/versionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyDefinitionVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.PolicyDefinitionName != v.Expected.PolicyDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicyDefinitionName", v.Expected.PolicyDefinitionName, actual.PolicyDefinitionName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestParsePolicyDefinitionVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyDefinitionVersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNnAmE/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions/versionName",
			Expected: &PolicyDefinitionVersionId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				PolicyDefinitionName: "policyDefinitionName",
				VersionName:          "versionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions/versionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNnAmE/vErSiOnS/vErSiOnNaMe",
			Expected: &PolicyDefinitionVersionId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				PolicyDefinitionName: "pOlIcYdEfInItIoNnAmE",
				VersionName:          "vErSiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNnAmE/vErSiOnS/vErSiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyDefinitionVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.PolicyDefinitionName != v.Expected.PolicyDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicyDefinitionName", v.Expected.PolicyDefinitionName, actual.PolicyDefinitionName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestSegmentsForPolicyDefinitionVersionId(t *testing.T) {
	segments := PolicyDefinitionVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyDefinitionVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
