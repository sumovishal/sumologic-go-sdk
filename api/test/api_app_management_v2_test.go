/*
Sumo Logic API

Testing AppManagementV2APIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sumologic

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_sumologic_AppManagementV2APIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppManagementV2APIService AsyncInstallApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.AppManagementV2API.AsyncInstallApp(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppManagementV2APIService AsyncUninstallApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.AppManagementV2API.AsyncUninstallApp(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppManagementV2APIService AsyncUpgradeApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.AppManagementV2API.AsyncUpgradeApp(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppManagementV2APIService GetAsyncInstallAppStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.AppManagementV2API.GetAsyncInstallAppStatus(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppManagementV2APIService GetAsyncUninstallAppStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.AppManagementV2API.GetAsyncUninstallAppStatus(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppManagementV2APIService GetAsyncUpgradeAppStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.AppManagementV2API.GetAsyncUpgradeAppStatus(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
