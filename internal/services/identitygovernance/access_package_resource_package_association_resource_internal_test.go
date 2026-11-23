// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

func TestFindAccessPackageResourceRole(t *testing.T) {
	resource := &beta.AccessPackageResource{
		AccessPackageResourceRoles: &[]beta.AccessPackageResourceRole{
			{
				Id:           pointer.To("role-default"),
				DisplayName:  nullable.Value("Default Access"),
				OriginId:     nullable.Value("00000000-0000-0000-0000-000000000000"),
				OriginSystem: nullable.Value("AadApplication"),
			},
			{
				Id:           pointer.To("role-reader"),
				DisplayName:  nullable.Value("Reader"),
				OriginId:     nullable.Value("00000000-0000-0000-0000-000000000001"),
				OriginSystem: nullable.Value("AadApplication"),
			},
		},
	}

	role, err := findAccessPackageResourceRole(resource, "00000000-0000-0000-0000-000000000001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pointer.From(role.Id) != "role-reader" || role.DisplayName.GetOrZero() != "Reader" || role.OriginSystem.GetOrZero() != "AadApplication" {
		t.Fatalf("unexpected role: %+v", role)
	}

	if _, err := findAccessPackageResourceRole(resource, "11111111-1111-1111-1111-111111111111"); err == nil {
		t.Fatal("expected an error for an unknown origin ID")
	}

	if _, err := findAccessPackageResourceRole(&beta.AccessPackageResource{}, "00000000-0000-0000-0000-000000000000"); err == nil {
		t.Fatal("expected an error when the catalog returns no roles")
	}
}
