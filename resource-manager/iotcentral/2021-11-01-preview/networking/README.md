
## `github.com/hashicorp/go-azure-sdk/resource-manager/iotcentral/2021-11-01-preview/networking` Documentation

The `networking` SDK allows for interaction with the Azure Resource Manager Service `iotcentral` (API Version `2021-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/iotcentral/2021-11-01-preview/networking"
```


### Client Initialization

```go
client := networking.NewNetworkingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `NetworkingClient.PrivateEndpointConnectionsCreate`

```go
ctx := context.TODO()
id := networking.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "privateEndpointConnectionValue")

payload := networking.PrivateEndpointConnection{
	// ...
}

future, err := client.PrivateEndpointConnectionsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkingClient.PrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := networking.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "privateEndpointConnectionValue")
future, err := client.PrivateEndpointConnectionsDelete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkingClient.PrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := networking.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "privateEndpointConnectionValue")
read, err := client.PrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkingClient.PrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := networking.NewIotAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")
read, err := client.PrivateEndpointConnectionsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkingClient.PrivateLinksGet`

```go
ctx := context.TODO()
id := networking.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "groupIdValue")
read, err := client.PrivateLinksGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkingClient.PrivateLinksList`

```go
ctx := context.TODO()
id := networking.NewIotAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")
read, err := client.PrivateLinksList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```