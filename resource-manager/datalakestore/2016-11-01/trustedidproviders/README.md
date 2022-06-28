
## `github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/trustedidproviders` Documentation

The `trustedidproviders` SDK allows for interaction with the Azure Resource Manager Service `datalakestore` (API Version `2016-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/trustedidproviders"
```


### Client Initialization

```go
client := trustedidproviders.NewTrustedIdProvidersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `TrustedIdProvidersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := trustedidproviders.NewTrustedIdProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "trustedIdProviderValue")

payload := trustedidproviders.CreateOrUpdateTrustedIdProviderParameters{
	// ...
}

read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TrustedIdProvidersClient.Delete`

```go
ctx := context.TODO()
id := trustedidproviders.NewTrustedIdProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "trustedIdProviderValue")
read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TrustedIdProvidersClient.Get`

```go
ctx := context.TODO()
id := trustedidproviders.NewTrustedIdProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "trustedIdProviderValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TrustedIdProvidersClient.ListByAccount`

```go
ctx := context.TODO()
id := trustedidproviders.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")
// alternatively `client.ListByAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TrustedIdProvidersClient.Update`

```go
ctx := context.TODO()
id := trustedidproviders.NewTrustedIdProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "trustedIdProviderValue")

payload := trustedidproviders.UpdateTrustedIdProviderParameters{
	// ...
}

read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```