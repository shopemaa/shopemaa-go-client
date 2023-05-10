<img src="https://shopemaa.com/assets/homev2/images/shopemaa/shopemaa.png" alt="" height="150" /> 

# Shopemaa Go Client

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/shopemaa/shopemaa-go-client?label=current-version)

Golang client to use shopemaa in your Go based projects.

### Installation

* Add dependency

```text
go get github.com/shopemaa/shopemaa-go-client
```

### Example

* Get Store Info

```go
shopemaaClient := client.NewClient(
		http.DefaultClient,
		"https://api.shopemaa.com/query",
		nil,
		func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
			// Setup store credentials
			req.Header.Set("store-key", os.Getenv("STORE_KEY"))
			req.Header.Set("store-secret", os.Getenv("STORE_SECRET"))
			return next(ctx, req, gqlInfo, res)
		},
	)

store := shopemaaClient.StoreBySecret(context.Background())
store.GetStoreBySecret().GetName()
```

# Contact

Shopemaa Support: support@shopemaa.com

# License

Licensed under the [MIT License](./LICENSE)