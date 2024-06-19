package assets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/atlanhq/atlan-go/atlan/model"
)

// A client for operating on Atlan's tenant object storage.
type FileClient struct {
	Client *AtlanClient
}

// NewFileClient creates a new instance of FileClient.
func NewFileClient(client *AtlanClient) *FileClient {
	return &FileClient{Client: client}
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
	rawJSON, err := DefaultAtlanClient.CallAPI(&PRESIGNED_URL, nil, request)
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
func (client *FileClient) UploadFile(presignedUrl string, filePath string) (string, error) {
	var PRESIGNED_URL_UPLOAD = API{
		Path:     presignedUrl,
		Method:   http.MethodPut,
		Status:   http.StatusOK,
		Endpoint: HeraclesEndpoint,
	}
	var fileBuffer bytes.Buffer

	err := handleFileUpload(filePath, &fileBuffer)
	if err != nil {
		return "", err
	}

	response, err := DefaultAtlanClient.s3PresignedUrlFileUpload(&PRESIGNED_URL_UPLOAD, fileBuffer)
	if err != nil {
		return "", err
	}
	return string(response), nil
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

	_, err = DefaultAtlanClient.s3PresignedUrlFileDownload(&PRESIGNED_URL_DOWNLOAD, file)
	if err != nil {
		return err
	}
	return nil
}
