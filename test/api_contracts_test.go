/*
MetaFab API

Testing ContractsApiService

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

func Test_metafab_ContractsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ContractsApiService CreateContract", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ContractsApi.CreateContract(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ContractsApiService GetContracts", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ContractsApi.GetContracts(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ContractsApiService ReadContract", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var contractId string

        resp, httpRes, err := apiClient.ContractsApi.ReadContract(context.Background(), contractId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ContractsApiService WriteContract", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var contractId string

        resp, httpRes, err := apiClient.ContractsApi.WriteContract(context.Background(), contractId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
