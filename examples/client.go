package examples

import (
	"context"
	"github.com/Yamashou/gqlgenc/clientv2"
	client "github.com/shopemaa/shopemaa-go-client"
	"net/http"
	"os"
)

func CreateClient() client.ShopemaaGraphQLClient {
	return client.NewClient(
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
}
