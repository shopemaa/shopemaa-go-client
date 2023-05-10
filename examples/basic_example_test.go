package examples

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStore(t *testing.T) {
	store, err := GetStore()
	if err != nil {
		t.Error(err)
		return
	}

	assert.NotNil(t, store.GetStoreBySecret())
	assert.NotEmpty(t, store.GetStoreBySecret().GetName())
	assert.True(t, store.GetStoreBySecret().GetCurrency().IsValid())

	t.Log(store.GetStoreBySecret().GetName())
}
