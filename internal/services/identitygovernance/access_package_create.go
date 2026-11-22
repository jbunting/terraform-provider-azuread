// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
)

// accessPackageCreateRequest is the body for creating an access package. The SDK's beta.AccessPackage marshaler drops
// catalogId as read-only, but entitlement management authorizes a catalog-scoped caller (Catalog owner, Access package
// manager) from that property, so without it such a caller is refused with a 403. This puts catalogId back.
type accessPackageCreateRequest struct {
	beta.AccessPackage
}

func (r accessPackageCreateRequest) MarshalJSON() ([]byte, error) {
	encoded, err := json.Marshal(r.AccessPackage)
	if err != nil {
		return nil, err
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackage: %+v", err)
	}

	if catalogId := r.CatalogId.GetOrZero(); catalogId != "" {
		decoded["catalogId"] = catalogId
	}

	return json.Marshal(decoded)
}

// createAccessPackage does what EntitlementManagementAccessPackageClient.CreateEntitlementManagementAccessPackage
// does, sending accessPackageCreateRequest as the body.
func createAccessPackage(ctx context.Context, c *entitlementmanagementaccesspackage.EntitlementManagementAccessPackageClient, input beta.AccessPackage) (*beta.AccessPackage, error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: entitlementmanagementaccesspackage.DefaultCreateEntitlementManagementAccessPackageOperationOptions(),
		Path:          "/identityGovernance/entitlementManagement/accessPackages",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err = req.Marshal(accessPackageCreateRequest{AccessPackage: input}); err != nil {
		return nil, err
	}

	resp, err := req.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var model beta.AccessPackage
	if err = resp.Unmarshal(&model); err != nil {
		return nil, err
	}

	return &model, nil
}
