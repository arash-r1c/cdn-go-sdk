/*
ArvanCloud CDN Services

Testing CachingApiService

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

func Test_arvancloud_CachingApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CachingApiService CachingDeprecatedPurge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.CachingApi.CachingDeprecatedPurge(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CachingApiService CachingIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.CachingApi.CachingIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CachingApiService CachingPurge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.CachingApi.CachingPurge(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CachingApiService CachingUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.CachingApi.CachingUpdate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CachingApiService PurgeTagsDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.CachingApi.PurgeTagsDestroy(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CachingApiService PurgeTagsIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.CachingApi.PurgeTagsIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
