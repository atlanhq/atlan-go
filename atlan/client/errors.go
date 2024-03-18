package client

import "errors"

var (
	ErrNotFound                       = errors.New("not found")
	ErrDuplicateCustomAttributes      = errors.New("duplicate custom attributes")
	ErrMissingCustomMetadataName      = errors.New("missing custom metadata name")
	ErrMissingCustomMetadataID        = errors.New("missing custom metadata ID")
	ErrCustomMetadataNotFoundByID     = errors.New("custom metadata not found by ID")
	ErrCustomMetadataNotFoundByName   = errors.New("custom metadata not found by name")
	ErrMissingCustomMetadataAttrID    = errors.New("missing custom metadata attribute ID")
	ErrCustomMetadataAttrNotFoundByID = errors.New("custom metadata attribute not found by ID")
	ErrExpiredAPIToken                = errors.New("API token has expired")
	ErrMissingCMName                  = errors.New("missing custom metadata name")
	ErrCMNotFoundByName               = errors.New("custom metadata not found by name")
	ErrMissingCMID                    = errors.New("missing custom metadata ID")
	ErrCMNotFoundByID                 = errors.New("custom metadata not found by ID")
)
