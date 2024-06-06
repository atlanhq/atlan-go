package main

import (
	"fmt"
	"github.com/atlanhq/atlan-go/atlan/assets"
)

func main() {

	ctx := assets.NewContext()

	//ctx, _ := assets.Context("tenet.atlan.com", "xyz")

	ctx.SetLogger(true, "debug")

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
}
