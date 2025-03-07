/*
ArvanCloud CDN Services

Testing PlanApiService

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

func Test_arvancloud_PlanApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PlanApiService DomainsPlans", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.PlanApi.DomainsPlans(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlanApiService DomainsPlansUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.PlanApi.DomainsPlansUpdate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlanApiService DomainsPlansUsages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.PlanApi.DomainsPlansUsages(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlanApiService DomainsPlansViolations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.PlanApi.DomainsPlansViolations(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlanApiService PlansIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PlanApi.PlansIndex(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
