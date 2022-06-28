
## `github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/deletedworkspaces` Documentation

The `deletedworkspaces` SDK allows for interaction with the Azure Resource Manager Service `operationalinsights` (API Version `2020-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/deletedworkspaces"
```


### Client Initialization

```go
client := deletedworkspaces.NewDeletedWorkspacesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `DeletedWorkspacesClient.List`

```go
ctx := context.TODO()
id := deletedworkspaces.NewSubscriptionID()
read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedWorkspacesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := deletedworkspaces.NewResourceGroupID()
read, err := client.ListByResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```