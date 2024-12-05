package main

import (
	_ "github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/assets"
	_ "github.com/atlanhq/atlan-go/atlan/model/structs"
)

func main() {

	ctx := assets.NewContext()
	ctx.EnableLogging("debug")

	/*
		// Get Persona by Guid
	
			response, atlanErr := assets.GetByGuid[*assets.Persona]("6f04ac74-d6b8-4b5e-8c1b-2347f9e55414")
			if atlanErr != nil {
				fmt.Println("Error:", atlanErr)
			} else {
				fmt.Println("RoleID:", *response.RoleId)
				fmt.Println("Users:", *response.PersonaUsers)
				fmt.Println("Groups:", *response.PersonaGroups)
				fmt.Println("DenyAssetFilters:", *response.DenyAssetFilters)
				//firstPolicy := (*response.Policies)[0]
				fmt.Println("Policies", (*response.Policies)[0].DisplayName)
				//		fmt.Println("RoleID:", *response.Policies)

			}


	*/
	/*
		// Personalize Persona

		toUpdate := &assets.Persona{}
		toUpdate.Updater("default/Lnwt6yWFzPfbH95MXauMWR", "Test Persona", true)

		toUpdate.DenyAssetTabs = &[]string{atlan.AssetSidebarTabLineage.Name, atlan.AssetSidebarTabRelations.Name, atlan.AssetSidebarTabQueries.Name}
		toUpdate.DenyAssetTypes = &[]string{"Table", "Column"}
		toUpdate.DenyAssetFilters = &[]string{atlan.AssetFilterGroupTags.Name, atlan.AssetFilterGroupOwners.Name, atlan.AssetFilterGroupCertificate.Name}
		toUpdate.DenyCustomMetadataGuids = &[]string{"default/7b4b4b4b-4b4b-4b4b-4b4b-4b4b4b4b4b4b"}

		response, atlanErr := assets.Save(toUpdate)
		if atlanErr != nil {
			println("Error:", atlanErr)
		} else {
			for _, entity := range response.MutatedEntities.UPDATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}
		}


	*/
	/*
		// List Policies in a Persona
		response, atlanErr := assets.NewFluentSearch().
			PageSizes(20).
			AssetType("Persona").
			Where(ctx.Persona.NAME.Eq("Test Persona")).
			IncludeOnResults("policies").
			IncludeOnRelations("name").
			IncludeOnRelations("policyActions").
			IncludeOnRelations("policyResources").
			IncludeOnRelations("policyType").
			Execute()

		if atlanErr != nil {
			fmt.Println("Error:", atlanErr)
		}

		for _, entity := range response[0].Entities {
			if entity.TypeName != nil && *entity.TypeName == "Persona" {
				fmt.Println("Persona Found: Name:", *entity.Name, "QualifiedName:", *entity.QualifiedName)
				for _, policy := range *entity.Policies {
					fmt.Println("Policy Found: Guid:", *policy.UniqueAttributes.QualifiedName)
				}
			}
		}

	*/

	/*
		// List Personas
		response, atlanErr := assets.NewFluentSearch().
			PageSizes(10).
			ActiveAssets().
			AssetType("Persona").
			Execute()

		if atlanErr != nil {
			fmt.Println("Error:", atlanErr)
		}
		for _, entity := range response[0].Entities {
			if entity.TypeName != nil && *entity.TypeName == "Persona" {
				fmt.Printf("Persona Found: Name: %s, QualifiedName: %s\n", *entity.Name, *entity.QualifiedName)
				// Perform any additional operations with the Persona entity
				revised, err := assets.TrimToRequired(entity)
				if err != nil {
					fmt.Println("Error:", err)
				}
				DisplayName := "Test Persona Modified"
				revised.DisplayName = &DisplayName
				response, err := assets.Save(revised)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					for _, entity := range response.MutatedEntities.UPDATE {
						println("Response:", entity)
						println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
					}
				}
			}
		}

	*/

	/*
		// Allow access to domain
		Persona := &assets.Persona{}
		domain, _ := Persona.CreateDomainPolicy(
			"Allow access to domain",
			"55226625-0b82-4705-8095-4a7d3f0c228d",
			[]atlan.PersonaDomainAction{atlan.PersonaDomainActionRead, atlan.PersonaDomainActionReadSubdomain, atlan.PersonaDomainActionReadProducts},
			[]string{"entity:default/domain/marketing"},
		)
		response, err := assets.Save(domain)
		if err != nil {
			println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}
		}

	*/
	/*
		// Add a glossary policy
		Persona := &assets.Persona{}
		glossary, _ := Persona.CreateGlossaryPolicy(
			"All glossaries",
			"55226625-0b82-4705-8095-4a7d3f0c228d",
			atlan.AuthPolicyTypeAllow,
			[]atlan.PersonaGlossaryAction{atlan.PersonaGlossaryActionCreate, atlan.PersonaGlossaryActionUpdate},
			[]string{"entity:OW0lMXZKyj4VfCsRxK3nr"},
		)
		response, err := assets.Save(glossary)
		if err != nil {
			println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}

		}

	*/
	/*
		// Create data policy
		Persona := &assets.Persona{}
		data, _ := Persona.CreateDataPolicy(
			"Allow access to data",
			"55226625-0b82-4705-8095-4a7d3f0c228d",
			atlan.AuthPolicyTypeAllow,
			"default/app/1732538219",
			[]string{"entity:default/app/1732538219"},
		)
		response, err := assets.Save(data)
		if err != nil {
			println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}

		}

	*/
	/*
		// Create Metadata Policy
		Persona := &assets.Persona{}
		metadata, _ := Persona.CreateMetadataPolicy(
			"Simple read access",
			"55226625-0b82-4705-8095-4a7d3f0c228d",
			atlan.AuthPolicyTypeAllow,
			[]atlan.PersonaMetadataAction{atlan.PersonaMetadataActionRead},
			"default/app/1732538219",
			[]string{"entity:default/app/1732538219"},
		)
		response, err := assets.Save(metadata)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}
		}

	*/
	/*
		// Add subjects to persona

		toUpdate := &assets.Persona{}
		toUpdate.Updater("default/Lnwt6yWFzPfbH95MXauMWR", "Test Persona", true)
		toUpdate.PersonaGroups = &[]string{"group1", "group2"}
		toUpdate.PersonaUsers = &[]string{"jsmith", "jdoe"}
		response, err := assets.Save(toUpdate)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.UPDATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}
		}


	*/
	/*
		// Activate or Deactivate a Persona
		toUpdate := &assets.Persona{}
		toUpdate.Updater("default/Lnwt6yWFzPfbH95MXauMWR", "Test Persona", true)
		response, err := assets.Save(toUpdate)
		if err != nil {
			println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.UPDATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}
		}

	*/
	/*
		// Delete a Persona
		assets.PurgeByGuid([]string{"5ffb7d15-4435-4f5c-8bb1-6e109ea091c2"})


	*/
	/*
		// Updater on Persona
			Persona := &assets.Persona{}

			toUpdate, err := Persona.Updater("default/jtKc6jzvE4UM8yT5uuC3FX", "Test Persona", true)
			if err != nil {
				return
			}
			DisplayName := "Test Persona Modified"
			toUpdate.Name = &DisplayName
			response, err := assets.Save(toUpdate)
			if err != nil {
				println("Error:", err)
			} else {
				for _, entity := range response.MutatedEntities.UPDATE {
					println("Response:", entity)
					println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
				}
			}

	*/
	/*
		// Creator on Persona
		toCreate := &assets.Persona{}

		toCreate.Creator("Test Persona")
		response, err := assets.Save(toCreate)
		if err != nil {
			println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}
		}

	*/

	/*
		ctx := assets.NewContext()

		//ctx, _ := assets.Context("tenet.atlan.com", "xyz")

		ctx.SetLogger(true, "debug")

		generator.RunGenerator()


	*/
	/*
			//t := &assets.Table{} // create a new Table instance

			// Define the Atlan tag details
			qualifiedName := "default/snowflake/1725896074/ANALYTICS/WIDE_WORLD_IMPORTERS/FCT_STOCK_ITEM_HOLDINGS"
			//atlanTagNames := []string{"Daily", "Hourly"} // List of tags to add

			err := assets.RemoveAtlanTag[*assets.Table](qualifiedName, "Confidential")
			/*
				// Set the propagation options
				propagate := true
				removePropagationOnDelete := true
				restrictLineagePropagation := false
				restrictPropagationThroughHierarchy := false


				// Call the AddAtlanTags function
				err := assets.UpdateAtlanTags[*assets.Table](
					qualifiedName,                       // The qualified name of the asset
					atlanTagNames,                       // The list of Atlan tags to add
					propagate,                           // Whether to propagate the tags or not
					removePropagationOnDelete,           // Remove propagation on delete
					restrictLineagePropagation,          // Restrict lineage propagation
					restrictPropagationThroughHierarchy, // Restrict propagation through hierarchy
				)

		if err != nil {
			fmt.Printf("Failed to add Atlan tags: %v\n", err)
		} else {
			fmt.Println("Atlan tags added successfully.")
		}
	*/
	//schemaName := "WIDEWORLDIMPORTERS_PURCHASING"
	//dataBaseName := "RAW"
	//dataBaseQualifiedName := "default/snowflake/1723642516/RAW"
	//connectionQualifiedName := "default/snowflake/1723642516"

	/*
		t.Creator("TestTable6", "default/snowflake/1723642516/RAW/WIDEWORLDIMPORTERS_PURCHASING")
		response, err := assets.Save(t) // save the table
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				//fmt.Println("Response:", entity)
				fmt.Printf("Entity ID: %s, Display Text: %s\n", entity.Guid, entity.DisplayText)
			}
		}

		t1 := &assets.Table{} // create a new Table instance

		t1.Updater("TestTable7", "default/snowflake/1723642516/RAW/WIDEWORLDIMPORTERS_PURCHASING/TestTable4")
		DisplayName := "TestTableModified"
		t1.Name = &DisplayName
		response2, err := assets.Save(t1)
		if err != nil {
		} else {
			for _, entity := range response2.MutatedEntities.UPDATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}
		}

	*/
	/*
		qualifiedname := "default/snowflake/1715371897/RAW/WIDEWORLDIMPORTERS_SALESFORCE/FIVETRAN_API_CALL"

		response, atlanErr := assets.NewFluentSearch().
			PageSizes(10).
			ActiveAssets().
			Where(ctx.Table.SUPERTYPE_NAMES.Eq("SQL")).
			Where(ctx.Table.QUALIFIED_NAME.Eq(qualifiedname)).
			IncludeOnResults("userDescription", "ownerUsers", "ownerGroups", "certificateStatus", "tags").
			Execute()

		if atlanErr != nil {
			fmt.Println(atlanErr)
		}

		fmt.Println(response[0].Entities[0].SearchMeanings[0].Guid)


	*/
	/*
		structs.GetAll()

		// Fetch columns of table from Atlan using Qualified Name
		assetQualifiedName := "default/snowflake/1715371897/RAW/WIDEWORLDIMPORTERS_SALESFORCE/FIVETRAN_API_CALL"

		columnSearchResponse, _ := structs.NewFluentSearch().
			PageSizes(1000).
			ActiveAssets().
			Where(ctx.Column.TYPENAME.Eq("Column")).
			Where(ctx.Column.TABLE_QUALIFIED_NAME.Eq(assetQualifiedName)).
			IncludeOnResults("userDescription", "dataType", "isPrimary").
			SetUtmTags(atlan.PROJECT_SDK_CLI).
			Execute()

		fmt.Println(*columnSearchResponse[0].Entities[0].Name)
		//fmt.Println(*columnSearchResponse[0].Entities[0].DataType)

		qualifiedname := "default/snowflake/1715371897/RAW/WIDEWORLDIMPORTERS_SALESFORCE/FIVETRAN_API_CALL"

		response, atlanErr := structs.NewFluentSearch().
			PageSizes(10).
			ActiveAssets().
			Where(ctx.Table.QUALIFIED_NAME.Eq(qualifiedname)).
			IncludeOnResults("userDescription", "ownerUsers", "ownerGroups", "certificateStatus", "tags").
			Execute()

		if atlanErr != nil {
			fmt.Println(atlanErr)
		}

		fmt.Println(response[0].Entities[0].SearchMeanings[0].Guid)

		/*


	*/

	/*
		// Fetch columns of table from Atlan using Qualified Name
		qualifiedname := "default/snowflake/1714501359/RAW/WIDEWORLDIMPORTERS_SALESFORCE/WAITLIST_WORK_TYPE_HISTORY/ID"

		columnResult, _ := structs.NewFluentSearch().
			PageSizes(50).
			ActiveAssets().
			Where(ctx.Column.TYPENAME.Eq("Column")).
			Where(ctx.Column.QUALIFIED_NAME.Eq(qualifiedname)).
			IncludeOnResults("userDescription", "dataType", "isPrimary", "isNullable").
			Execute()

		fmt.Println(*columnResult[0].Entities[0].IsPrimary)

	*/
	/*
		if err != nil {
			fmt.Printf("Error executing search: %v\n", err)
			return
		}

		fmt.Println("Search results:", *columnResult[0].Entities[0].Name)
		fmt.Println("Search results:", *columnResult[0].Entities[0].QualifiedName)


	*/
	//structs.GetAll()

	/*
		g := &structs.AtlasGlossary{}
		g.Creator("go-sdk-test1", atlan.AtlanIconAirplaneInFlight)
		response, err := structs.Save(g)
		if err != nil {
			fmt.Println(err)
		}
		for _, entity := range response.MutatedEntities.CREATE {
			fmt.Println(entity.DisplayText)
		}


		dc := &structs.DataContract{}
		dc.Creator("DataContractLatestCertified")
		response, err := structs.Save()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("response", response)


	*/
	/*
		structs.GetCustomMetadataCache().RefreshCache()
		id, err := structs.GetCustomMetadataCache().GetIDForName("testcmgsdk")
		if err != nil {
			fmt.Println("Error:", err)
		}
		name, _ := structs.GetCustomMetadataCache().GetNameForID("cvhn5T7YwnsYXiMCKh9PoW")
		attrID, _ := structs.GetCustomMetadataCache().GetAttrIDForName("testcmgsdk", "gsdk")
		attrName, _ := structs.GetCustomMetadataCache().GetAttrNameForID("cvhn5T7YwnsYXiMCKh9PoW", "foYi4v02OVjKt0YzcTCKM3")
		fmt.Println("ID for name is : ", id)
		fmt.Println("Name fo ID is:", name)
		fmt.Printf("\nAttrID for name is: %s\n", attrID)
		fmt.Printf("\nAttrName for ID is: %s\n", attrName)

		fmt.Println(structs.GetCustomMetadataCache().GetAttributesForSearchResultsByName("testcmgsdk"))
		customMetadata := structs.GetCustomMetadataCache().GetAttributesForSearchResultsByName("testcmgsdk")

		customMeta, _ := structs.NewFluentSearch().
			PageSizes(50).
			ActiveAssets().
			Where(ctx.Glossary.QUALIFIED_NAME.Eq("fW6NU2lWKaMy5ZyVlGYes")).
			IncludeOnResults(customMetadata...).
			IncludeOnResults("terms").
			IncludeOnResults("tags").
			Execute()

		for _, entity := range customMeta[0].Entities {
			fmt.Println("Entity:", *entity.AssetDbtJobNextRunHumanized)
		}


	*/
	/*
		ctx := structs.NewContext()
		assetQualifiedName := "default/mssql/1711817247/WideWorldImporters/Purchasing/SupplierCategories_Archive"
		columnSearchResponse, err := structs.NewFluentSearch().
			PageSizes(1000).
			ActiveAssets().
			Where(ctx.Column.TYPENAME.Eq("Column")).
			Where(ctx.Column.TABLE_QUALIFIED_NAME.Eq(assetQualifiedName)).
			Execute()
		if err != nil {
			return
		}

		fmt.Println("Search results:", *columnSearchResponse[0].Entities[2].DisplayName)
		// TEST INTEGRATION ATLAN-CLI WITH GO-SDK

		// Fluent-Search

		query := ctx.Table.QUALIFIED_NAME.Eq("default/mssql/1711817247/WideWorldImporters/Purchasing/SupplierCategories_Archive")
		//query2 := ctx.Column.TYPENAME.Eq("Column", nil)

		searchResult, err := structs.NewFluentSearch().
			PageSizes(50).
			Where(query).
			Execute()

		if err != nil {
			fmt.Printf("Error executing search: %v\n", err)
			return
		}
		fmt.Println("Search results:", *searchResult[0].Entities[0].DisplayName)


	*/
	/*
		// Fetch columns of table from Atlan using Qualified Name
		qualifiedname := "default/snowflake/1711213678/RAW/WIDEWORLDIMPORTERS_SALESFORCE/SELLER_HISTORY/ID"

		columnResult, err := structs.NewFluentSearch().
			PageSizes(50).
			ActiveAssets().
			Where(ctx.Column.TYPENAME.Eq("Column", nil)).
			Where(ctx.Column.QUALIFIED_NAME.Eq(qualifiedname, nil)).
			Execute()

		if err != nil {
			fmt.Printf("Error executing search: %v\n", err)
			return
		}

		fmt.Println("Search results:", *columnResult[0].Entities[0].SearchAttributes.Name)
		fmt.Println("Search results:", *columnResult[0].Entities[0].SearchAttributes.QualifiedName)
		if columnResult[0].Entities[0].Description != nil {
			fmt.Println("Search results:", *columnResult[0].Entities[0].Description)
		}
	*/
	/*
			resp, err := structs.GetAll()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
		*
		/*
			response, err := structs.GetGlossaryByGuid("fc36342b-ddb5-44ba-b774-4c90cc66d5a2")

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				println("Response:", *response.TypeName)
			}
	*/
	//structs.GetAtlanTagCache().RefreshCache()
	//id, _ := structs.GetAtlanTagCache().GetIDForName("Hourly")
	//fmt.Println("Print Response:", id)

	/*
		structs.GetCustomMetadataCache().RefreshCache()
		id, err := structs.GetCustomMetadataCache().GetIDForName("go-test")
		if err != nil {
			fmt.Println("Error:", err)
		}
		name, _ := structs.GetCustomMetadataCache().GetNameForID("Wd2QonGuCz6tj4uryxEoBs")
		attrID, _ := structs.GetCustomMetadataCache().GetAttrIDForName("go-test", "go-test")
		attrName, _ := structs.GetCustomMetadataCache().GetAttrNameForID("Wd2QonGuCz6tj4uryxEoBs", "fmfsP9P9IqWzqbXXCyMibP")
		fmt.Println("ID for name is : ", id)
		fmt.Println("Name fo ID is:", name)
		fmt.Printf("\nAttrID for name is: %s\n", attrID)
		fmt.Printf("\n AttrName for ID is: %s\n", attrName)
	*/
	/*
		atlanTypeCategory := atlan.AtlanTypeCategoryRelationship
		ctx := &model.CustomMetadataDef{
			Category: &atlanTypeCategory,
		}

		fmt.Println("Category:", *ctx.Category)
	*/
	//structs.Init()
	/*
		response, err := structs.GetGlossaryByGuid("f273e814-f80e-4699-83f3-9462a153fb14")
		if err != nil {
			println("Error:", err)
		}
		print("Response:", response.Name)
	*/
	//response, err := structs.FindCategoryByName("Oak Ridge", "DDwycTZ007zZYxRajRVDK")

	/*
		response, err := structs.FindGlossaryByName("go-sdk-test")
		fmt.Println("Response:", response)
		if err != nil {
			fmt.Println("Error:", err)
		}
	*/
	/*
		// IndexSearch
		//boolQuery, _ := structs.WithActiveGlossary("go-sdk-test")
		boolQuery2 := &model.TermQuery{Field: ctx.Column.TYPENAME.GetElasticFieldName(), Value: "Column"}
		//boolQuery3 := &model.PrefixQuery{Field: ctx.Table.NAME.GetElasticFieldName(), Value: "SE"}

		request := model.IndexSearchRequest{
			Dsl: model.Dsl{
				From:           0,
				Size:           30,
				Query:          boolQuery2.ToJSON(),
				TrackTotalHits: true,
			},
			SuppressLogs:           true,
			ShowSearchScore:        false,
			ExcludeMeanings:        false,
			ExcludeClassifications: false,
		}

		response1, _ := structs.Search(request)

		fmt.Println("Guid:", *response1.Entities[0].Guid)
		fmt.Println("Total Results", response1.ApproximateCount)
		fmt.Println("Typename:", *response1.Entities[0].TypeName)
	*/
	//if err != nil {
	//	fmt.Printf("Error fetching model: %v\n", err)
	//}

	//fmt.Printf("Response: %+v\n", response)

	//ctx := structs.NewContext()

	// structs.Init()
	/*
		g := &structs.AtlasGlossary{} // create a new Glossary instance

		g.Creator("TestGlossary14", atlan.AtlanIconAirplaneInFlight)

		response, err := structs.Save(g) // save the Glossary
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				fmt.Println("Response:", entity)
				fmt.Printf("Entity ID: %s, Display Text: %s\n", entity.Guid, entity.DisplayText)
			}
		}
	*/
	//Modify an existing Glossary
	/*
		g := &structs.AtlasGlossary{}
		DisplayName := "Testing5"
		g.Updater("TestGlossary8", "CBtveYe0Avp5iwU8q3M7Y1", "78e820ef-21a2-4f1d-a1dc-8c5a648cb1e3")
		g.DisplayName = &DisplayName
		response, err := structs.Save(g)
		if err != nil {
		} else {
			for _, entity := range response.MutatedEntities.UPDATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}
		}
	*/
	// Deleting an asset
	//structs.DeleteByGuid([]string{"024f11b6-a9fa-4f45-84f5-f734c47c4743", "b280b09b-5c28-45c4-a899-d8535fb651eb", "8679e70a-513e-4e2e-9861-4f5559206f36"})
	//structs.DeleteByGuid([]string{"dbe090bd-1549-4cce-98dd-6542138963f1"})
	/*
		resp, _ := structs.PurgeByGuid([]string{"1d9f74c6-faa9-4840-ac9e-21723b4c63ca"})
		for _, entity := range resp.MutatedEntities.DELETE {
			fmt.Println("Response:", entity)
			fmt.Println("TypeName:", entity.TypeName)
			fmt.Println("Guid:", entity.Guid)
			fmt.Println("Status:", entity.Status)
			fmt.Println("DisplayText:", entity.DisplayText)
			// Add other fields you want to print
		}
	*/

	//query := ctx.Glossary.TYPENAME.Eq("AtlasGlossary", nil)
	/*
		excludeCondition := &model.TermQuery{
			Field: structs.NAME,
			Value: "Concepts",
		}

		searchResult, err := structs.NewFluentSearch().
			PageSizes(10).
			ActiveAssets().
			AssetType("AtlasGlossary").
			Where(&model.TermQuery{
				Field: ctx.Table.TYPENAME.GetElasticFieldName(),
				Value: "Table",
			}).
			Where(ctx.Glossary.NAME.Eq("Metrics", nil)).
			Where(ctx.Glossary.NAME.StartsWith("M", nil)).
			Sort(structs.NAME, atlan.SortOrderAscending).
			//Sort(string(structs.GUID), structs.Ascending).
			WhereNot(excludeCondition).
			IncludeOnResults("guid").
			IncludeOnRelations("terms").
			Execute()

		if err != nil {
			fmt.Printf("Error executing search: %v\n", err)
			return
		}

		println("Search results:")
		fmt.Println(searchResult[0].Entities[0])

		// Process search results
		for _, entity := range searchResult[0].Entities {
			fmt.Printf("Entity ID: %s, Display Text: %s\n", entity.Guid, entity.DisplayText)
		}

		//glossaryGuid := "c1620acb-e89d-4bb2-8bee-3f56be6439b5"
		//glossaryGuidterm := "1ee6a1e5-7afa-4b31-a736-af1f656ae0c3"

		//t, err := structs.GetAll()
		//structs.DefaultAtlanTagCache.RefreshCache()
		//if err != nil {
		//	fmt.Printf("Error fetching model: %v\n", err)
		//	return
		//}

		structs.GetAll()
		structs.GetAtlanTagCache().RefreshCache()
		structs.GetAtlanTagCache()
		a, _ := structs.GetAtlanTagIDForName("PII")
		fmt.Printf("ID for name is : %s\n", a)

		tagName := "Hourly"
		tagID, err := structs.GetAtlanTagIDForName(tagName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Atlan tag ID for %s: %s\n", tagName, tagID)

		tagID = "a3el9UemzJZqZAUFcsDjy4"
		tagName, err = structs.GetAtlanTagNameForID(tagID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Atlan tag name for %s: %s\n", tagID, tagName)

	*/
	//structs.GetCache().RefreshCache()
	//structs.GetCache()
	//a, _ := structs.GetCache().GetIDForName("PII")
	//fmt.Printf("ID for name is : %s\n", a)
	/*
		tagName := "Hourly"
		tagID, err := structs.GetIDForName(tagName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Atlan tag ID for %s: %s\n", tagName, tagID)
	*/
	/*
		tagID = "a3el9UemzJZqZAUFcsDjy4"
		tagName, err = structs.DefaultAtlanTagCache.GetNameForID(tagID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Atlan tag name for %s: %s\n", tagID, tagName)

		/*
			// Use the GlossaryClient to get a model by its GUID
			g, err := structs.GetGlossaryByGuid(glossaryGuid)
			if err != nil {
				fmt.Printf("Error fetching model: %v\n", err)
				return
			}

			gt, err := structs.GetGlossaryTermByGuid(glossaryGuidterm)
			if err != nil {
				fmt.Printf("Error fetching model: %v\n", err)
				return
			}

			fmt.Println("Retrieved Glossary:")
			fmt.Printf("GUID: %s\n", g.Guid)
			fmt.Printf("CreatedBy: %s\n", g.CreatedBy)
			fmt.Printf("TypeName %s\n", g.TypeName)
			fmt.Printf("Term 1 Guid %s\n", g.Terms[0].Guid)
			fmt.Printf("Popularity %f\n", g.Attributes.PopularityScore)

			fmt.Println("Retrieved GlossaryTerm:")
			fmt.Printf("TypeName: %s\n", gt.TypeName)
			fmt.Printf("GUID: %s\n", gt.Guid)
			fmt.Printf("Anchor Guid: %s\n", gt.Anchor.Guid)
			fmt.Printf("Tags Typename %s\n", gt.Tags[0].TypeName)
	*/
	//fmt.Println("Retrieved Typedef:")
	//fmt.Printf("DisplayName: %s\n", t.AtlanTagDefs[1].DisplayName)
	//
}
