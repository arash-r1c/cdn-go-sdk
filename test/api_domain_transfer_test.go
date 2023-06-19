/*
ArvanCloud CDN Services

Testing DomainTransferApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package r1cdn

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/arvancloud/cdn-go"
)

func Test_r1cdn_DomainTransferApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DomainTransferApiService DomainsTransferIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DomainTransferApi.DomainsTransferIndex(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainTransferApiService DomainsTransferStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DomainTransferApi.DomainsTransferStore(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainTransferApiService DomainsTransferUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DomainTransferApi.DomainsTransferUpdate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
