package integration

import (
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func TestSetDefaultHttpClientIdleConnTimeout(t *testing.T) {
	secretId := os.Getenv("TENCENTCLOUD_SECRET_ID")
	secretKey := os.Getenv("TENCENTCLOUD_SECRET_KEY")

	// Save original state
	originalDefaultClient := common.DefaultHttpClient
	defer func() {
		common.DefaultHttpClient = originalDefaultClient
	}()

	// Test setting custom timeout using the new function
	customTimeout := 40 * time.Second
	common.SetDefaultHttpClientIdleConnTimeout(customTimeout)

	// Create client which should use the configured DefaultHttpClient
	credential := common.NewCredential(secretId, secretKey)
	client, err := cvm.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())
	if err != nil {
		t.Fatalf("fail to init client: %v", err)
	}

	// Use reflection to access the private httpClient field
	clientValue := reflect.ValueOf(client).Elem()
	commonClientValue := clientValue.FieldByName("Client")
	if !commonClientValue.IsValid() {
		t.Fatal("could not access common.Client field")
	}

	httpClientField := commonClientValue.FieldByName("httpClient")
	if !httpClientField.IsValid() {
		t.Fatal("could not access httpClient field")
	}

	// Make the field accessible
	httpClientField = reflect.NewAt(httpClientField.Type(), unsafe.Pointer(httpClientField.UnsafeAddr())).Elem()
	httpClient := httpClientField.Interface().(*http.Client)

	// Verify the client uses the configured DefaultHttpClient
	if httpClient != common.DefaultHttpClient {
		t.Error("client should use DefaultHttpClient when it's set")
	}

	transport, ok := httpClient.Transport.(*http.Transport)
	if !ok {
		t.Fatal("client transport should be *http.Transport")
	}

	if transport.IdleConnTimeout != customTimeout {
		t.Errorf("IdleConnTimeout expected %v, got %v", customTimeout, transport.IdleConnTimeout)
	}

	request := cvm.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)
	if err != nil {
		t.Fatalf("fail to invoke api: %v", err)
	}
	if *response.Response.TotalCount != (int64)(len(response.Response.InstanceSet)) {
		t.Fatalf("response item count inconsisitent!")
	}
}
