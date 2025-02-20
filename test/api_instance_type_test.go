/*
FastAPI

Testing InstanceTypeAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package perian

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Perian-io/perian-go-client"
)

func Test_perian_InstanceTypeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InstanceTypeAPIService GetAvailableInstanceZonesInstanceTypeZonesAvailableGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InstanceTypeAPI.GetAvailableInstanceZonesInstanceTypeZonesAvailableGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceTypeAPIService GetInstanceTypeById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceTypeId string

		resp, httpRes, err := apiClient.InstanceTypeAPI.GetInstanceTypeById(context.Background(), instanceTypeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceTypeAPIService GetInstanceTypeByRequirements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InstanceTypeAPI.GetInstanceTypeByRequirements(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
