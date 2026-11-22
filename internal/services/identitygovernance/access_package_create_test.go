// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

func TestAccessPackageCreateRequestIncludesCatalogId(t *testing.T) {
	body, err := json.Marshal(accessPackageCreateRequest{AccessPackage: beta.AccessPackage{
		DisplayName: nullable.Value("test"),
		CatalogId:   nullable.Value("00000000-0000-0000-0000-000000000001"),
	}})
	if err != nil {
		t.Fatalf("marshaling: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshaling: %+v", err)
	}

	if got := decoded["catalogId"]; got != "00000000-0000-0000-0000-000000000001" {
		t.Fatalf("expected catalogId in body, got %v (body %s)", got, body)
	}
	if got := decoded["displayName"]; got != "test" {
		t.Fatalf("expected displayName in body, got %v", got)
	}
	if got := decoded["@odata.type"]; got != "#microsoft.graph.accessPackage" {
		t.Fatalf("expected @odata.type in body, got %v", got)
	}
}
