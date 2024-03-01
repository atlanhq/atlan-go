package client

import (
	"atlan-go/atlan/model"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type CustomMetadataCache struct {
	AtlanClient     *AtlanClient
	CacheByID       map[string]model.CustomMetadataDef
	AttrCacheByID   map[string]model.AttributeDef
	MapIDToName     map[string]string
	MapNameToID     map[string]string
	MapAttrIDToName map[string]map[string]string
	MapAttrNameToID map[string]map[string]string
	archivedAttrIds map[string]struct{}
	mutex           sync.RWMutex
}

// NewCustomMetadataCache creates a new CustomMetadataCache instance.
func NewCustomMetadataCache(atlanClient *AtlanClient) *CustomMetadataCache {
	return &CustomMetadataCache{
		AtlanClient:     DefaultAtlanClient,
		CacheByID:       make(map[string]model.CustomMetadataDef),
		AttrCacheByID:   make(map[string]model.AttributeDef),
		MapIDToName:     make(map[string]string),
		MapNameToID:     make(map[string]string),
		MapAttrIDToName: make(map[string]map[string]string),
		MapAttrNameToID: make(map[string]map[string]string),
		archivedAttrIds: make(map[string]struct{}),
	}
}

var (
	customMetadataCaches = make(map[string]*CustomMetadataCache)
)

func RefreshCustomMetadataCache() {
	GetCustomMetadataCache().RefreshCache()
}

// GetCustomMetadataCache returns the CustomMetadataCache for the default AtlanClient.
func GetCustomMetadataCache() *CustomMetadataCache {
	client := DefaultAtlanClient
	cacheKey := client.ApiKey

	mu.Lock()
	defer mu.Unlock()

	if _, ok := customMetadataCaches[cacheKey]; !ok {
		customMetadataCaches[cacheKey] = &CustomMetadataCache{
			AtlanClient:     client,
			CacheByID:       make(map[string]model.CustomMetadataDef),
			AttrCacheByID:   make(map[string]model.AttributeDef),
			MapIDToName:     make(map[string]string),
			MapNameToID:     make(map[string]string),
			MapAttrIDToName: make(map[string]map[string]string),
			MapAttrNameToID: make(map[string]map[string]string),
			archivedAttrIds: make(map[string]struct{}),
		}
	}
	return customMetadataCaches[cacheKey]
}

// RefreshCache refreshes the cache of custom metadata structures.
func (c *CustomMetadataCache) RefreshCache() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	api := &GET_ALL_TYPE_DEFS

	response, err := DefaultAtlanClient.CallAPI(api, nil, nil)
	if err != nil {
		return err
	}

	// Parse the response and populate the cacheByID, mapIDToName, mapNameToID accordingly
	var customMetadataDefs []model.CustomMetadataDef
	err = json.Unmarshal(response, &customMetadataDefs)
	if err != nil {
		return err
	}

	// Clear existing cache data
	c.CacheByID = make(map[string]model.CustomMetadataDef)
	c.MapIDToName = make(map[string]string)
	c.MapNameToID = make(map[string]string)
	c.MapAttrIDToName = make(map[string]map[string]string)
	c.MapAttrNameToID = make(map[string]map[string]string)
	c.archivedAttrIds = make(map[string]struct{})

	// Populate the cache with the fetched custom metadata structures and their attributes
	for _, cmDef := range customMetadataDefs {
		typeID := cmDef.Name
		typeName := cmDef.DisplayName

		c.CacheByID[typeID] = cmDef
		c.MapIDToName[typeID] = typeName
		c.MapNameToID[typeName] = typeID
		c.MapAttrIDToName[typeID] = make(map[string]string)
		c.MapAttrNameToID[typeID] = make(map[string]string)

		for _, attr := range cmDef.AttributeDefs {
			attrID := *attr.Name
			attrName := *attr.DisplayName

			if _, exists := c.MapAttrNameToID[typeID][attrName]; exists {
				return err
			}

			c.MapAttrIDToName[typeID][attrID] = attrName
			c.MapAttrNameToID[typeID][attrName] = attrID

			if *attr.Options.IsArchived {
				c.archivedAttrIds[attrID] = struct{}{}
			}
		}
	}
	return nil
}

func (c *CustomMetadataCache) GetIDForName(name string) (string, error) {
	if name = strings.TrimSpace(name); name == "" {
		//return "", ErrMissingCMName
	}

	c.mutex.RLock()
	cmID, exists := c.MapNameToID[name]
	c.mutex.RUnlock()

	if !exists {
		if err := c.RefreshCache(); err != nil {
			return "", err
		}

		c.mutex.RLock()
		cmID, exists = c.MapNameToID[name]
		c.mutex.RUnlock()

		if !exists {
			//return "", errorWithParameters(ErrCMNotFoundByName, name)
		}
	}

	return cmID, nil
}

func (c *CustomMetadataCache) GetNameForID(idstr string) (string, error) {
	if idstr = strings.TrimSpace(idstr); idstr == "" {
		return "", ErrMissingCMID
	}

	c.mutex.RLock()
	cmName, exists := c.MapIDToName[idstr]
	c.mutex.RUnlock()

	if !exists {
		if err := c.RefreshCache(); err != nil {
			return "", err
		}

		c.mutex.RLock()
		cmName, exists = c.MapIDToName[idstr]
		c.mutex.RUnlock()

		if !exists {
			//return "", errorWithParameters(ErrCMNotFoundByID, idstr)
		}
	}

	return cmName, nil
}

func (c *CustomMetadataCache) GetAllCustomAttributes(includeDeleted, forceRefresh bool) (map[string][]model.AttributeDef, error) {
	if len(c.CacheByID) == 0 || forceRefresh {
		c.RefreshCache()
	}
	m := make(map[string][]model.AttributeDef)
	for typeID, cm := range c.CacheByID {
		typeName, err := c.GetNameForID(typeID)
		if err != nil {
			return nil, fmt.Errorf("custom metadata not found by name")
			//return nil, ErrorCode.CM_NOT_FOUND_BY_ID
		}
		var toInclude []model.AttributeDef
		if includeDeleted {
			toInclude = cm.AttributeDefs
		} else {
			for _, attr := range cm.AttributeDefs {
				if attr.Options == nil || !*attr.Options.IsArchived {
					toInclude = append(toInclude, attr)
				}
			}
		}
		m[typeName] = toInclude
	}
	return m, nil
}

func (c *CustomMetadataCache) GetAttrIDForName(setName, attrName string) (string, error) {
	setID, err := c.GetIDForName(setName)
	if err != nil {
		return "", err
	}
	if subMap, ok := c.MapAttrNameToID[setID]; ok {
		if attrID, ok := subMap[attrName]; ok {
			return attrID, nil
		}
	}
	// Refresh cache and try again
	c.RefreshCache()
	if subMap, ok := c.MapAttrNameToID[setID]; ok {
		if attrID, ok := subMap[attrName]; ok {
			return attrID, nil
		}
		return "", fmt.Errorf("attribute %s not found in set %s", attrName, setName)
	}
	return "", fmt.Errorf("set %s not found", setName)
}

func (c *CustomMetadataCache) GetAttrNameForID(setID, attrID string) (string, error) {
	if subMap, ok := c.MapAttrIDToName[setID]; ok {
		if attrName, ok := subMap[attrID]; ok {
			return attrName, nil
		}
	}
	// Refresh cache and try again
	c.RefreshCache()
	if subMap, ok := c.MapAttrIDToName[setID]; ok {
		if attrName, ok := subMap[attrID]; ok {
			return attrName, nil
		}
	}
	//return "", ErrorCode.CM_ATTR_NOT_FOUND_BY_ID
	return "", fmt.Errorf("attribute not found by id")
}

func (c *CustomMetadataCache) GetAttributesForSearchResults(setID string) []string {
	if subMap, ok := c.MapAttrNameToID[setID]; ok {
		var dotNames []string
		for idstr := range subMap {
			dotNames = append(dotNames, fmt.Sprintf("%s.%s", setID, idstr))
		}
		return dotNames
	}
	return nil
}

func (c *CustomMetadataCache) GetAttributeForSearchResults(setID, attrName string) string {
	if subMap, ok := c.MapAttrNameToID[setID]; ok {
		if attrID, ok := subMap[attrName]; ok {
			return attrID
		}
	}
	return ""
}

func (c *CustomMetadataCache) GetAttributesForSearchResultsByName(setName string) []string {
	setID, _ := c.GetIDForName(setName) // Assuming this method is defined
	if setID == "" {
		return nil
	}
	return c.GetAttributesForSearchResults(setID)
}

func (c *CustomMetadataCache) GetCustomMetadataDef(name string) (model.CustomMetadataDef, error) {
	baID, _ := c.GetIDForName(name) // Assuming this method is defined
	if baID == "" {
		//return CustomMetadataDef{}, ErrorCode.CM_NOT_FOUND_BY_NAME
		return model.CustomMetadataDef{}, fmt.Errorf("missing custom metadata attribute id")

	}
	if typedef, ok := c.CacheByID[baID]; ok {
		return typedef, nil
	}
	//return CustomMetadataDef{}, ErrorCode.CM_NOT_FOUND_BY_NAME
	return model.CustomMetadataDef{}, fmt.Errorf("missing custom metadata attribute id")

}

func (c *CustomMetadataCache) GetAttributeDef(attrID string) (model.AttributeDef, error) {
	if attrID == "" {
		//return AttributeDef{}, ErrorCode.MISSING_CM_ATTR_ID
		return model.AttributeDef{}, fmt.Errorf("missing custom metadata attribute id")
	}
	if c.AttrCacheByID == nil {
		// Assuming _refresh_cache() is defined
		c.RefreshCache()
	}
	if attrDef, ok := c.AttrCacheByID[attrID]; ok {
		return attrDef, nil
	}
	//return AttributeDef{}, ErrorCode.CM_ATTR_NOT_FOUND_BY_ID
	return model.AttributeDef{}, fmt.Errorf("missing custom metadata attribute id")

}
