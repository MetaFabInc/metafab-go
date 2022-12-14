/*
MetaFab API

Testing ShopsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package metafab

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_metafab_ShopsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ShopsApiService CreateShop", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ShopsApi.CreateShop(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ShopsApiService GetShopOffer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var shopId string
        var shopOfferId string

        resp, httpRes, err := apiClient.ShopsApi.GetShopOffer(context.Background(), shopId, shopOfferId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ShopsApiService GetShopOffers", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var shopId string

        resp, httpRes, err := apiClient.ShopsApi.GetShopOffers(context.Background(), shopId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ShopsApiService GetShops", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ShopsApi.GetShops(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ShopsApiService RemoveShopOffer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var shopId string
        var shopOfferId string

        resp, httpRes, err := apiClient.ShopsApi.RemoveShopOffer(context.Background(), shopId, shopOfferId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ShopsApiService SetShopOffer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var shopId string

        resp, httpRes, err := apiClient.ShopsApi.SetShopOffer(context.Background(), shopId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ShopsApiService UseShopOffer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var shopId string
        var shopOfferId string

        resp, httpRes, err := apiClient.ShopsApi.UseShopOffer(context.Background(), shopId, shopOfferId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ShopsApiService WithdrawFromShop", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var shopId string

        resp, httpRes, err := apiClient.ShopsApi.WithdrawFromShop(context.Background(), shopId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
