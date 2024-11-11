package generator

// Legacy Code

/*


// AssetInfo is the struct to hold all asset information

	type AssetInfo struct {
		Name                 string
		EntityDef            *model.EntityDef
		ModuleInfo           *ModuleInfo
		RequiredAssetInfos   map[string]*AssetInfo
		CircularDependencies map[string]*AssetInfo
		Order                int
		ModuleName           string
		SuperClass           string
		SuperType            *AssetInfo
		IsCoreAsset          bool
	}

	type ModuleInfo struct {
		Assets map[string]*AssetInfo
	}

// AssetInfoByName stores asset information by name
var AssetInfoByName = make(map[string]*AssetInfo)

// SuperTypeNamesToIgnore holds the set of super types to ignore
var SuperTypeNamesToIgnore = make(map[string]struct{})

// EntityDefsByName holds entity definitions by name
var EntityDefsByName = make(map[string]*model.EntityDef)

// SubTypeNamesToIgnore holds subtypes to ignore
var SubTypeNamesToIgnore = make(map[string]struct{})

var Referenceable = "Referenceable"

var HierarchyGraph = simple.NewDirectedGraph()

// CoreAssets is the set of core assets

	var CoreAssets = map[string]struct{}{
		"Referenceable":    {},
		"Asset":            {},
		"AuthPolicy":       {},
		"DataModel":        {},
		"DataModeling":     {},
		"MatillionGroup":   {},
		"Stakeholder":      {},
		"StakeholderTitle": {},
	}

// NewAssetInfo initializes a new AssetInfo

	func NewAssetInfo(name string, entityDef *model.EntityDef) *AssetInfo {
		assetInfo := &AssetInfo{
			Name:                 name,
			EntityDef:            entityDef,
			RequiredAssetInfos:   make(map[string]*AssetInfo),
			CircularDependencies: make(map[string]*AssetInfo),
			ModuleName:           name,
		}
		assetInfo.UpdateAttributeDefs()
		AssetInfoByName[name] = assetInfo
		return assetInfo
	}

// SetEntityDefs sets the entity definitions and updates the hierarchy graph

	func SetEntityDefs(entityDefs []*model.EntityDef) {
		// Step 1: Populate EntityDefsByName map
		for _, entityDef := range entityDefs {
			EntityDefsByName[entityDef.Name] = entityDef
		}

		// Step 2: Sort entity definitions by super types
		sort.Slice(entityDefs, func(i, j int) bool {
			superTypesI := joinStringPointers(entityDefs[i].SuperTypes)
			superTypesJ := joinStringPointers(entityDefs[j].SuperTypes)
			return superTypesI < superTypesJ
		})

		// Step 3: Iterate through sorted entity definitions
		for _, entityDef := range entityDefs {
			name := entityDef.Name

			// Handle super types
			if (len(entityDef.SuperTypes) == 0 && name != "Referenceable") || hasSuperTypeToIgnore(entityDef.SuperTypes) {
				SuperTypeNamesToIgnore[name] = struct{}{}
				continue
			}

			// Add edges to hierarchy graph based on subtypes
			for _, assetName := range entityDef.SubTypes {
				if _, ignore := SubTypeNamesToIgnore[*assetName]; !ignore {
					// Create nodes for both the parent and child
					parentNode := HierarchyGraph.NewNode()
					childNode := HierarchyGraph.NewNode()

					// Add directed edge from parent to child
					HierarchyGraph.AddDirectedEdge(parentNode.ID(), childNode.ID()) // Adjust as needed
				}
			}

			// Create AssetInfo and store in AssetInfoByName
			assetInfo := &AssetInfo{
				Name:      name,
				EntityDef: entityDef,
			}
			AssetInfoByName[name] = assetInfo
		}

		// Step 4: Update required asset names for all AssetInfo instances
		for _, assetInfo := range AssetInfoByName {
			assetInfo.UpdateRequiredAssetNames()
		}
	}

// Helper function to join string pointers into a single string

	func joinStringPointers(strPtrs []*string) string {
		var strs []string
		for _, strPtr := range strPtrs {
			if strPtr != nil {
				strs = append(strs, *strPtr)
			}
		}
		return strings.Join(strs, ",")
	}

// Helper function to check if any super type should be ignored

	func hasSuperTypeToIgnore(superTypes []*string) bool {
		for _, superType := range superTypes {
			if superType != nil {
				if _, exists := SuperTypeNamesToIgnore[*superType]; exists {
					return true
				}
			}
		}
		return false
	}

	func (a *AssetInfo) UpdateAttributeDefs() {
		if len(a.EntityDef.SuperTypes) > 1 {
			a.EntityDef.AttributeDefs = mergeAttributes(*a.EntityDef)
		}
		names := make(map[string]struct{})
		for _, attributeDef := range a.EntityDef.AttributeDefs {
			names[attributeDef["name"].(string)] = struct{}{}
		}
		superTypeRelationshipDefs := getAncestorRelationshipDefs(*a.EntityDef.SuperTypes[0], map[string]struct{}{})
		a.EntityDef.RelationshipAttributeDefs = filterRelationshipDefs(a.EntityDef.RelationshipAttributeDefs, names, superTypeRelationshipDefs)
	}

	func getAncestorRelationshipDefs(ancestorName string, ancestorRelationshipDefs map[string]struct{}) map[string]struct{} {
		ancestorEntityDef, exists := EntityDefsByName[ancestorName]
		if !exists || len(ancestorEntityDef.SuperTypes) == 0 {
			return ancestorRelationshipDefs
		}

		for _, relationshipDef := range ancestorEntityDef.RelationshipAttributeDefs {
			ancestorRelationshipDefs[relationshipDef["name"].(string)] = struct{}{}
		}
		return getAncestorRelationshipDefs(*ancestorEntityDef.SuperTypes[0], ancestorRelationshipDefs)
	}

	func filterRelationshipDefs(relationshipDefs []map[string]interface{}, names map[string]struct{}, superTypeRelationshipDefs map[string]struct{}) []map[string]interface{} {
		var filtered []map[string]interface{}
		for _, relationshipDef := range relationshipDefs {
			name := relationshipDef["name"].(string)
			if _, exists := names[name]; !exists && superTypeRelationshipDefs[name] == struct{}{} {
				filtered = append(filtered, relationshipDef)
			}
		}
		return filtered
	}

	func mergeAttributes(entityDef model.EntityDef) []map[string]interface{} {
		attributes := map[string]map[string]interface{}{}
		for _, attribute := range entityDef.AttributeDefs {
			attributes[attribute["name"].(string)] = attribute
		}

		for _, superType := range entityDef.SuperTypes {
			mergeFromSuperType(*superType, attributes)
		}

		var mergedAttributes []map[string]interface{}
		for _, attribute := range attributes {
			mergedAttributes = append(mergedAttributes, attribute)
		}
		return mergedAttributes
	}

	func mergeFromSuperType(superType string, attributes map[string]map[string]interface{}) {
		entity, exists := EntityDefsByName[superType]
		if !exists {
			return
		}

		for _, attribute := range entity.AttributeDefs {
			if _, exists := attributes[attribute["name"].(string)]; !exists {
				attributes[attribute["name"].(string)] = attribute
			}
		}
		for _, sType := range entity.SuperTypes {
			mergeFromSuperType(*sType, attributes)
		}
	}

	func (a *AssetInfo) UpdateCircularDependencies() {
		for _, requiredAsset := range a.RequiredAssetInfos {
			if _, exists := requiredAsset.RequiredAssetInfos[a.Name]; exists {
				a.CircularDependencies[requiredAsset.Name] = requiredAsset
			}
		}
		if len(a.EntityDef.SuperTypes) > 0 {
			superType := AssetInfoByName[*a.EntityDef.SuperTypes[0]]
			if _, exists := superType.RequiredAssetInfos[a.Name]; exists {
				a.CircularDependencies[superType.Name] = superType
			}
		}
	}

	func (a *AssetInfo) UpdateRequiredAssetNames() {
		attributesToRemove := map[string]struct{}{}
		for _, attribute := range append(a.EntityDef.AttributeDefs, a.EntityDef.RelationshipAttributeDefs...) {
			typeName := strings.ReplaceAll(strings.ReplaceAll(attribute["typeName"].(string), "array<", ""), ">", "")
			if typeName == a.Name {
				continue
			}
			if _, exists := SuperTypeNamesToIgnore[typeName]; exists {
				attributesToRemove[attribute["name"].(string)] = struct{}{}
			} else if assetInfo, exists := AssetInfoByName[typeName]; exists {
				a.RequiredAssetInfos[assetInfo.Name] = assetInfo
			}
		}

		// Update attribute defs by filtering out attributes to remove
		a.EntityDef.AttributeDefs = filterAttributes(a.EntityDef.AttributeDefs, attributesToRemove)
		a.EntityDef.RelationshipAttributeDefs = filterAttributes(a.EntityDef.RelationshipAttributeDefs, attributesToRemove)
	}

	func filterAttributes(attributes []map[string]interface{}, attributesToRemove map[string]struct{}) []map[string]interface{} {
		var filtered []map[string]interface{}
		for _, attribute := range attributes {
			if _, exists := attributesToRemove[attribute["name"].(string)]; !exists {
				filtered = append(filtered, attribute)
			}
		}
		return filtered
	}

	func UpdateSubTypeNamesToIgnore(customEntityDefNames []string) {
		for _, name := range customEntityDefNames {
			SubTypeNamesToIgnore[name] = struct{}{}
		}
	}

	type Generator struct {
		templates *ModuleInfo
	}

// NewGenerator initializes the generator and loads templates

	func NewGenerator() (*Generator, error) {
		_, err := template.ParseGlob("templates/*.tmpl")
		if err != nil {
			return nil, fmt.Errorf("error loading templates: %v", err)
		}

		return &Generator{templates: &ModuleInfo{Assets: make(map[string]*AssetInfo)}}, nil
	}

// Helper function to retrieve the asset name from a node

	func getAssetNameFromNode(n graph.Node) string {
		assetInfo, exists := AssetInfoByName[fmt.Sprint(n.ID())]
		if !exists {
			return ""
		}
		return assetInfo.Name
	}

// Method to create modules using BFS traversal

	func createModules() {
		order := 0

		// Perform BFS traversal on the hierarchy graph starting from "Referenceable"
		bfs := traverse.BreadthFirst{
			Visit: func(n graph.Node, depth int) bool {
				parentAssetName := getAssetNameFromNode(n)

				// If parent asset name is empty, continue
				if parentAssetName == "" {
					return false
				}

				parentAssetInfo := AssetInfoByName[parentAssetName]

				// Set the order for the asset and increment the order counter
				parentAssetInfo.Order = order
				order++

				// Check if the asset is a core asset
				if _, isCore := CoreAssets[parentAssetName]; isCore {
					// If the super class isn't "AtlanObject", mark the super asset as core
					if parentAssetInfo.SuperClass != "AtlanObject" {
						superAssetInfo := AssetInfoByName[parentAssetInfo.SuperClass]
						superAssetInfo.IsCoreAsset = true
						CoreAssets[parentAssetInfo.SuperClass] = struct{}{}
					}

					// Mark related assets as core assets if applicable
					for _, relatedAsset := range parentAssetInfo.RequiredAssetInfos {
						if !relatedAsset.IsCoreAsset {
							relatedAsset.IsCoreAsset = true
							CoreAssets[relatedAsset.Name] = struct{}{}
						}
					}
				}

				// Handle the case where the super class is in the core assets
				if _, isCore := CoreAssets[parentAssetInfo.SuperClass]; isCore {
					for _, relatedAsset := range parentAssetInfo.RequiredAssetInfos {
						if relatedAsset.IsCoreAsset {
							parentAssetInfo.IsCoreAsset = true
							CoreAssets[parentAssetName] = struct{}{}
							continue
						}

						// Mark the super class of the related asset as core
						superAssetInfo := AssetInfoByName[relatedAsset.SuperClass]
						superAssetInfo.IsCoreAsset = true
						CoreAssets[relatedAsset.SuperClass] = struct{}{}
					}
				}
				return false
			},
		}

		// Start BFS traversal from the "Referenceable" node
		startNode := findNodeByAssetName("Referenceable")
		bfs.Walk(HierarchyGraph, startNode, nil)
	}

// Helper function to find a node by asset name

	func findNodeByAssetName(name string) graph.Node {
		for nodes := HierarchyGraph.Nodes(); nodes.Next(); {
			node := nodes.Node()
			if AssetInfoByName[fmt.Sprint(node.ID())].Name == name {
				return node
			}
		}
		return nil
	}
*/
