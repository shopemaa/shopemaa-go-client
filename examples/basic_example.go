package examples

import (
	"context"
	client "github.com/shopemaa/shopemaa-go-client"
)

func GetStore() (*client.StoreBySecret, error) {
	shopemaaClient := CreateClient()
	return shopemaaClient.StoreBySecret(context.Background())
}
