package assets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/atlanhq/atlan-go/atlan/model"
)

// A client for operating on Atlan's tenant object storage.
type FileClient struct {
	*AtlanClient
}

// NewFileClient creates a new instance of FileClient.
func NewFileClient(client *AtlanClient) *FileClient {
	return &FileClient{client}
}

func handleFileUpload(filePath string, fileBuffer *bytes.Buffer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(fileBuffer, file)
	if err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}
	return nil
}

// Generates a presigned URL based on Atlan's tenant object store.
func (client *FileClient) GeneratePresignedURL(request *model.PresignedURLRequest) (string, error) {
	rawJSON, err := client.CallAPI(&PRESIGNED_URL, nil, request)
	if err != nil {
		return "", AtlanError{
			ErrorCode: errorCodes[CONNECTION_ERROR],
			Args:      []interface{}{"IOException"},
		}
	}
	// Now unmarshal `rawJSON` to the `PresignedURLResponse`
	var response model.PresignedURLResponse
	err = json.Unmarshal(rawJSON, &response)
	if err != nil {
		return "", fmt.Errorf("Error while unmarshaling PresignedURLResponse JSON: %v", err)
	}
	return response.URL, nil
}

// Uploads a file to Atlan's object storage.
func (client *FileClient) UploadFile(presignedUrl string, filePath string) error {
	var PRESIGNED_URL_UPLOAD = API{
		Path:     presignedUrl,
		Method:   http.MethodPut,
		Status:   http.StatusOK,
		Endpoint: HeraclesEndpoint,
	}
	var fileBuffer bytes.Buffer

	err := handleFileUpload(filePath, &fileBuffer)
	if err != nil {
		return err
	}

	// Currently supported upload methods for different cloud storage providers
	switch {
	case strings.Contains(presignedUrl, string(model.S3)):
		err = client.s3PresignedUrlFileUpload(&PRESIGNED_URL_UPLOAD, fileBuffer)
	default:
		return InvalidRequestError{AtlanError{ErrorCode: errorCodes[UNSUPPORTED_PRESIGNED_URL]}}
	}

	if err != nil {
		return err
	}
	return nil
}

// Downloads a file from Atlan's tenant object storage.
func (client *FileClient) DownloadFile(presignedUrl string, filePath string) error {
	var PRESIGNED_URL_DOWNLOAD = API{
		Path:     presignedUrl,
		Method:   http.MethodGet,
		Status:   http.StatusOK,
		Endpoint: HeraclesEndpoint,
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create download file: %v", err)
	}
	defer file.Close()

	err = client.s3PresignedUrlFileDownload(&PRESIGNED_URL_DOWNLOAD, file)
	if err != nil {
		return err
	}
	return nil
}
