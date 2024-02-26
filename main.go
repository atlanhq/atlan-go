// main.go
package main

import (
	"atlan-go/atlan/client"
)

func main() {

	client.LoggingEnabled = true

	client.Init()

	response, err := client.GetGlossaryByGuid("f273e814-f80e-4699-83f3-9462a153fb14")
	if err != nil {
		println("Error:", err)
	}
	print("Response:", response.AssetIcon)

	//response, err := client.FindGlossaryByName("Manhattan Project")
	//response, err := client.FindCategoryByName("Oak Ridge", "DDwycTZ007zZYxRajRVDK")

	//if err != nil {
	//	fmt.Printf("Error fetching model: %v\n", err)
	//}

	//fmt.Printf("Response: %+v\n", response)
	//excludeCondition := &client.TermQuery{
	//	Field: string(client.Name),
	//	Value: "Concepts",
	//}

	// ctx := client.NewContext()
	/*
		client.Init()

		g := &client.AtlasGlossary{} // create a new Glossary instance

		g.Creator("TestGlossary7", "") // initialize the Glossary
		response, err := client.Save(g)
		fmt.Println("Resp1:", response)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				fmt.Println("Response:", entity)
				fmt.Printf("Entity ID: %s, Display Text: %s\n", entity.Guid, entity.DisplayText)
			}
		}

	*/

	// Modify an existing Glossary
	/*
		g.Updater("TestGlossary7", "CBtveYe0Avp5iwU8q3M7Y", "e63cf857-c788-4197-a60e-397b24e749ee")
		g.Entities[0].Attributes.DisplayName = "Testing"
		response, err := g.Save()
		if err != nil {
			println("Error:", err)
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
	// query := ctx.Glossary.TYPENAME.Eq("AtlasGlossary", nil)

	/*
		searchResult, err := client.NewFluentSearch().
			PageSizes(10).
			ActiveAssets().
			AssetType("AtlasGlossary").
			//Where(&client.TermQuery{
			//	Field: string(client.TypeName),
			//	Value: "AtlasGlossary",
			//}).
			Where(ctx.Glossary.Name.Eq("Metrics", nil)).
			//Where(ctx.Glossary.Name.StartsWith("M", nil)).
			Sort(string(client.Name), client.Ascending).
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
	*/
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
