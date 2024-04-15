// main.go
package main

import (
	"fmt"
	"github.com/atlanhq/atlan-go/atlan/client"
	"log"
)

func main() {

	client.LoggingEnabled = false
	client.NewContext()

	client.GetAll()
	client.GetAtlanTagCache().RefreshCache()
	client.GetAtlanTagCache()
	a, _ := client.GetAtlanTagIDForName("PII")
	fmt.Printf("ID for name is : %s\n", a)

	tagName := "Hourly"
	tagID, err := client.GetAtlanTagIDForName(tagName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Atlan tag ID for %s: %s\n", tagName, tagID)

	tagID = "a3el9UemzJZqZAUFcsDjy4"
	tagName, err = client.GetAtlanTagNameForID(tagID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Atlan tag name for %s: %s\n", tagID, tagName)

	/*
		//ctx, _ := client.Context("eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICIwNC1UVWs2Z3RCdmtBdmsxd19XVTZqd0pVXzdNX1pyV0JXWkJEMjZ2WHl3In0.eyJleHAiOjE3MTMwMzI5OTksImlhdCI6MTcxMDI3MjM4OCwianRpIjoiZTQ1ZTA5ODUtOTc0MS00YTI5LWFlYjgtYzlhMWE3NDEwNDI5IiwiaXNzIjoiaHR0cHM6Ly9kZXZ4OC5hdGxhbi5jb20vYXV0aC9yZWFsbXMvZGVmYXVsdCIsImF1ZCI6WyJyZWFsbS1tYW5hZ2VtZW50IiwiYWNjb3VudCJdLCJzdWIiOiI4OWNiMWY2Zi1lOWUyLTRkZDYtOTVhMC0xYWQ0Yjk5YTdmNmYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJhcGlrZXktM2U1N2NmMjYtZjhiNi00MWNkLThhNzMtOTAzYmZhZTY3YjVkIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIiRndWVzdCIsIm9mZmxpbmVfYWNjZXNzIiwiZGVmYXVsdC1yb2xlcy1kZWZhdWx0IiwiJGFwaS10b2tlbi1kZWZhdWx0LWFjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsicmVhbG0tbWFuYWdlbWVudCI6eyJyb2xlcyI6WyJ2aWV3LXJlYWxtIiwidmlldy1pZGVudGl0eS1wcm92aWRlcnMiLCJtYW5hZ2UtaWRlbnRpdHktcHJvdmlkZXJzIiwiaW1wZXJzb25hdGlvbiIsInJlYWxtLWFkbWluIiwiY3JlYXRlLWNsaWVudCIsIm1hbmFnZS11c2VycyIsInF1ZXJ5LXJlYWxtcyIsInZpZXctYXV0aG9yaXphdGlvbiIsInF1ZXJ5LWNsaWVudHMiLCJxdWVyeS11c2VycyIsIm1hbmFnZS1ldmVudHMiLCJtYW5hZ2UtcmVhbG0iLCJ2aWV3LWV2ZW50cyIsInZpZXctdXNlcnMiLCJ2aWV3LWNsaWVudHMiLCJtYW5hZ2UtYXV0aG9yaXphdGlvbiIsIm1hbmFnZS1jbGllbnRzIiwicXVlcnktZ3JvdXBzIl19LCJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUgb2ZmbGluZV9hY2Nlc3MiLCJjcmVhdGVkQXQiOiIxNzEwMjcyMzg4MjM5IiwiY2xpZW50SWQiOiJhcGlrZXktM2U1N2NmMjYtZjhiNi00MWNkLThhNzMtOTAzYmZhZTY3YjVkIiwiY2xpZW50SG9zdCI6IjEwLjE1MC40My4yMDciLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImdyb3VwcyI6W10sInJlYWxtIjoiZGVmYXVsdCIsInByZWZlcnJlZF91c2VybmFtZSI6InNlcnZpY2UtYWNjb3VudC1hcGlrZXktM2U1N2NmMjYtZjhiNi00MWNkLThhNzMtOTAzYmZhZTY3YjVkIiwidXNlcklkIjoiODljYjFmNmYtZTllMi00ZGQ2LTk1YTAtMWFkNGI5OWE3ZjZmIiwiY2xpZW50QWRkcmVzcyI6IjEwLjE1MC40My4yMDciLCJ1c2VybmFtZSI6InNlcnZpY2UtYWNjb3VudC1hcGlrZXktM2U1N2NmMjYtZjhiNi00MWNkLThhNzMtOTAzYmZhZTY3YjVkIn0.Z2ItcUjf_RN11rbjknIljSfhxBLXyJBtkEqYX4EAVN_4oDOUaCZNBGrLxbCw3YrowPgfaLaOAxTH787hYXUJj9xYbQ-fTvu56Ia3KlJuFW-DDogKGR1aH0CahEyQh9ZrTWqQJKDo9UJf0lD4vVEdUh_bFcZUYNYYO8TcfdCi7SCLiCDu1Syll2qhU1b_DDYglfpn7axKkNKCKhvm_gDO8UQ-SWGV9KFapwz6W8y2bqaPTEoroyvnzYq31j9nCPDlrlWr-yrV0bvAsWGZA8cLPQDX5F90SVtI_2aeDINTRHLsN6syVAIxu36RK4Q_1JEvhx4gnNetdw0ZKg1JZNv7Fw", "https://devx8.atlan.com")
		ctx := client.NewContext()
		assetQualifiedName := "default/mssql/1711817247/WideWorldImporters/Purchasing/SupplierCategories_Archive"
		columnSearchResponse, err := client.NewFluentSearch().
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

		searchResult, err := client.NewFluentSearch().
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

		columnResult, err := client.NewFluentSearch().
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
			resp, err := client.GetAll()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
		*
		/*
			response, err := client.GetGlossaryByGuid("fc36342b-ddb5-44ba-b774-4c90cc66d5a2")

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				println("Response:", *response.TypeName)
			}
	*/
	//client.GetAtlanTagCache().RefreshCache()
	//id, _ := client.GetAtlanTagCache().GetIDForName("Hourly")
	//fmt.Println("Print Response:", id)

	/*
		client.GetCustomMetadataCache().RefreshCache()
		id, err := client.GetCustomMetadataCache().GetIDForName("go-test")
		if err != nil {
			fmt.Println("Error:", err)
		}
		name, _ := client.GetCustomMetadataCache().GetNameForID("Wd2QonGuCz6tj4uryxEoBs")
		attrID, _ := client.GetCustomMetadataCache().GetAttrIDForName("go-test", "go-test")
		attrName, _ := client.GetCustomMetadataCache().GetAttrNameForID("Wd2QonGuCz6tj4uryxEoBs", "fmfsP9P9IqWzqbXXCyMibP")
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
	//client.Init()
	/*
		response, err := client.GetGlossaryByGuid("f273e814-f80e-4699-83f3-9462a153fb14")
		if err != nil {
			println("Error:", err)
		}
		print("Response:", response.Name)
	*/
	//response, err := client.FindCategoryByName("Oak Ridge", "DDwycTZ007zZYxRajRVDK")

	/*
		response, err := client.FindGlossaryByName("go-sdk-test")
		fmt.Println("Response:", response)
		if err != nil {
			fmt.Println("Error:", err)
		}
	*/
	/*
		// IndexSearch
		//boolQuery, _ := client.WithActiveGlossary("go-sdk-test")
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

		response1, _ := client.Search(request)

		fmt.Println("Guid:", *response1.Entities[0].Guid)
		fmt.Println("Total Results", response1.ApproximateCount)
		fmt.Println("Typename:", *response1.Entities[0].TypeName)
	*/
	//if err != nil {
	//	fmt.Printf("Error fetching model: %v\n", err)
	//}

	//fmt.Printf("Response: %+v\n", response)

	//ctx := client.NewContext()

	// client.Init()
	/*
		g := &client.AtlasGlossary{} // create a new Glossary instance

		g.Creator("TestGlossary14", atlan.AtlanIconAirplaneInFlight)

		response, err := client.Save(g) // save the Glossary
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
		g := &client.AtlasGlossary{}
		DisplayName := "Testing5"
		g.Updater("TestGlossary8", "CBtveYe0Avp5iwU8q3M7Y1", "78e820ef-21a2-4f1d-a1dc-8c5a648cb1e3")
		g.DisplayName = &DisplayName
		response, err := client.Save(g)
		if err != nil {
		} else {
			for _, entity := range response.MutatedEntities.UPDATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}
		}
	*/
	// Deleting an asset
	//client.DeleteByGuid([]string{"024f11b6-a9fa-4f45-84f5-f734c47c4743", "b280b09b-5c28-45c4-a899-d8535fb651eb", "8679e70a-513e-4e2e-9861-4f5559206f36"})
	//client.DeleteByGuid([]string{"dbe090bd-1549-4cce-98dd-6542138963f1"})
	/*
		resp, _ := client.PurgeByGuid([]string{"1d9f74c6-faa9-4840-ac9e-21723b4c63ca"})
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
			Field: client.NAME,
			Value: "Concepts",
		}

		searchResult, err := client.NewFluentSearch().
			PageSizes(10).
			ActiveAssets().
			AssetType("AtlasGlossary").
			Where(&model.TermQuery{
				Field: ctx.Table.TYPENAME.GetElasticFieldName(),
				Value: "Table",
			}).
			Where(ctx.Glossary.NAME.Eq("Metrics", nil)).
			Where(ctx.Glossary.NAME.StartsWith("M", nil)).
			Sort(client.NAME, atlan.SortOrderAscending).
			//Sort(string(client.GUID), client.Ascending).
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

		//t, err := client.GetAll()
		//client.DefaultAtlanTagCache.RefreshCache()
		//if err != nil {
		//	fmt.Printf("Error fetching model: %v\n", err)
		//	return
		//}

		//client.GetCache().RefreshCache()
		//client.GetCache()
		//a, _ := client.GetCache().GetIDForName("PII")
		//fmt.Printf("ID for name is : %s\n", a)
		/*
			tagName := "Hourly"
			tagID, err := client.GetIDForName(tagName)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Atlan tag ID for %s: %s\n", tagName, tagID)
	*/
	/*
		tagID = "a3el9UemzJZqZAUFcsDjy4"
		tagName, err = client.DefaultAtlanTagCache.GetNameForID(tagID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Atlan tag name for %s: %s\n", tagID, tagName)

		/*
			// Use the GlossaryClient to get a model by its GUID
			g, err := client.GetGlossaryByGuid(glossaryGuid)
			if err != nil {
				fmt.Printf("Error fetching model: %v\n", err)
				return
			}

			gt, err := client.GetGlossaryTermByGuid(glossaryGuidterm)
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
