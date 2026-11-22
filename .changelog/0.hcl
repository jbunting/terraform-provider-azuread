change "resource-fix" {
  body = "`azuread_access_package` - send `catalogId` when creating, so a caller with only a catalog-scoped role (Catalog owner, Access package manager) is no longer refused with a 403"
}
