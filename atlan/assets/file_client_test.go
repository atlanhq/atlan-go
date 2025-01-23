package assets

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"testing"

	"github.com/atlanhq/atlan-go/atlan/model"
	"github.com/stretchr/testify/assert"
)

const TestDataDirectoy = "test_data"

const (
	UrlExpiry               = "10s"
	ImageFileName           = "go-sdk.png"
	TextFileName            = "go-sdk.txt"
	TextDownloadFileName    = "go-sdk-download.txt"
	ImageDownloadFileName   = "go-sdk-download.png"
	TenantS3BucketDirectory = "presigned-url-sdk-integration-tests"
	ExpectedTextContent     = "test data 12345.\n"
)

var (
	ImageS3UploadFilePath = fmt.Sprintf("%s/%s", TenantS3BucketDirectory, ImageFileName)
	TextS3UploadFilePath  = fmt.Sprintf("%s/%s", TenantS3BucketDirectory, TextFileName)
	UnsupportedURL        = "https://unsupported.storage.com/upload"
)

func TestIntegrationFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	ctx := NewContext()
	fileClient := NewFileClient(ctx)

	testCases := []struct {
		fileName       string
		downloadName   string
		s3UploadPath   string
		expectedFormat string // Expected file format, e.g: "png" or "txt"
	}{
		{ImageFileName, ImageDownloadFileName, ImageS3UploadFilePath, "png"},
		{TextFileName, TextDownloadFileName, TextS3UploadFilePath, "txt"},
	}

	for _, tc := range testCases {
		// Test S3 upload
		requestPUT := model.PresignedURLRequest{
			Key:    tc.s3UploadPath,
			Expiry: UrlExpiry,
			Method: model.PUT,
		}
		presignedURL := testGeneratePresignedURL(t, fileClient, requestPUT)

		fileToUpload := fmt.Sprintf("%s/%s", TestDataDirectoy, tc.fileName)
		testUploadFile(t, fileClient, presignedURL, fileToUpload)

		// Test S3 download
		requestGET := model.PresignedURLRequest{
			Key:    tc.s3UploadPath,
			Expiry: UrlExpiry,
			Method: model.GET,
		}
		presignedURL = testGeneratePresignedURL(t, fileClient, requestGET)
		fileDownloadPath := fmt.Sprintf("%s/%s", TestDataDirectoy, tc.downloadName)
		testDownloadFile(t, fileClient, presignedURL, fileDownloadPath, tc.expectedFormat)
	}

	// Test unsupported URL upload
	testUploadUnsupportedURL(t, fileClient, UnsupportedURL, fmt.Sprintf("%s/%s", TestDataDirectoy, TextFileName))
}

func testGeneratePresignedURL(t *testing.T, client *FileClient, request model.PresignedURLRequest) string {
	url, err := client.GeneratePresignedURL(&request)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, url, "Generated presigned url cannot be nil")
	return url
}

func testUploadUnsupportedURL(t *testing.T, client *FileClient, presignedURL string, filePath string) {
	err := client.UploadFile(presignedURL, filePath)
	assert.Error(t, err, "Expected an error for unsupported presigned URL")
	assert.Contains(t, err.Error(), "Provided presigned URL's cloud provider storage "+
		"is currently not supported for file uploads.", "Error message should indicate unsupported URL")
}

func testUploadFile(t *testing.T, client *FileClient, presignedURL string, filePath string) {
	err := client.UploadFile(presignedURL, filePath)
	if err != nil {
		t.Errorf("Encountered an error while uploading: %v", err)
	}
}

func testDownloadFile(t *testing.T, client *FileClient, presignedURL string, filePath string, expectedFormat string) {
	err := client.DownloadFile(presignedURL, filePath)
	if err != nil {
		t.Errorf("Encountered an error while downloading: %v", err)
	}

	// Check if the file exists
	_, err = os.Stat(filePath)
	assert.NoError(t, err, "The file does not exist")

	if expectedFormat == "png" {
		// Open the file
		file, err := os.Open(filePath)
		assert.NoError(t, err, "Failed to open the file")
		defer file.Close()

		// Check if it is a PNG image
		_, format, err := image.DecodeConfig(file)
		assert.NoError(t, err, "Failed to decode the image")
		assert.Equal(t, "png", format, "The file is not a PNG image")
	} else if expectedFormat == "txt" {
		// Read and check the file content
		fileContent, err := os.ReadFile(filePath)
		assert.NoError(t, err, "Failed to read the text file")
		assert.Equal(t, ExpectedTextContent, string(fileContent), "The file content does not match the expected content")
	}

	// Remove the file
	err = os.Remove(filePath)
	assert.NoError(t, err, "Failed to remove the file")
}
