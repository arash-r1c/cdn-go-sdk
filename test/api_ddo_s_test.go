/*
ArvanCloud CDN Services

Testing DDoSApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package arvancloud

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/arash-r1c/cdn-go-sdk"
)

func Test_arvancloud_DDoSApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DDoSApiService DdosReprioritize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DDoSApi.DdosReprioritize(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DDoSApiService DdosRulesDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.DDoSApi.DdosRulesDestroy(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DDoSApiService DdosRulesIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DDoSApi.DdosRulesIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DDoSApiService DdosRulesShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.DDoSApi.DdosRulesShow(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DDoSApiService DdosRulesStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DDoSApi.DdosRulesStore(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DDoSApiService DdosRulesUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.DDoSApi.DdosRulesUpdate(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DDoSApiService DdosSettingsIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DDoSApi.DdosSettingsIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DDoSApiService DdosSettingsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DDoSApi.DdosSettingsUpdate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
