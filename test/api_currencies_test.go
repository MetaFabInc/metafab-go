/*
MetaFab API

Testing CurrenciesApiService

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

func Test_metafab_CurrenciesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CurrenciesApiService BatchTransferCurrency", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var currencyId string

        resp, httpRes, err := apiClient.CurrenciesApi.BatchTransferCurrency(context.Background(), currencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService BurnCurrency", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var currencyId string

        resp, httpRes, err := apiClient.CurrenciesApi.BurnCurrency(context.Background(), currencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService CreateCurrency", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CurrenciesApi.CreateCurrency(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService GetCurrencies", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CurrenciesApi.GetCurrencies(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService GetCurrencyBalance", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var currencyId string

        resp, httpRes, err := apiClient.CurrenciesApi.GetCurrencyBalance(context.Background(), currencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService GetCurrencyFees", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var currencyId string

        resp, httpRes, err := apiClient.CurrenciesApi.GetCurrencyFees(context.Background(), currencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService GetCurrencyRole", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var currencyId string

        resp, httpRes, err := apiClient.CurrenciesApi.GetCurrencyRole(context.Background(), currencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService GrantCurrencyRole", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var currencyId string

        resp, httpRes, err := apiClient.CurrenciesApi.GrantCurrencyRole(context.Background(), currencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService MintCurrency", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var currencyId string

        resp, httpRes, err := apiClient.CurrenciesApi.MintCurrency(context.Background(), currencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService RevokeCurrencyRole", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var currencyId string

        resp, httpRes, err := apiClient.CurrenciesApi.RevokeCurrencyRole(context.Background(), currencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService SetCurrencyFees", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var currencyId string

        resp, httpRes, err := apiClient.CurrenciesApi.SetCurrencyFees(context.Background(), currencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CurrenciesApiService TransferCurrency", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var currencyId string

        resp, httpRes, err := apiClient.CurrenciesApi.TransferCurrency(context.Background(), currencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
