// main.go
package main

import (
	"atlan-go/atlan/client"
	"fmt"
)

func main() {

	client.LoggingEnabled = false
	err := client.Init()
	if err != nil {
		return
	}

	//response, err := client.FindGlossaryByName("Manhattan Project")
	//response, err := client.FindCategoryByName("Oak Ridge", "DDwycTZ007zZYxRajRVDK")

	//if err != nil {
	//	fmt.Printf("Error fetching model: %v\n", err)
	//}

	//fmt.Printf("Response: %+v\n", response)
	excludeCondition := &client.TermQuery{
		Field: string(client.Name),
		Value: "Retention",
	}

	searchResult, err := client.NewFluentSearch().
		PageSizes(10).
		ActiveAssets().
		AssetType("AtlasGlossaryCategory").
		Where(string(client.TypeName), "AtlasGlossaryCategory").
		Sort(string(client.Name), client.Ascending).
		WhereNot(excludeCondition).
		Execute()

	if err != nil {
		fmt.Printf("Error executing search: %v\n", err)
		return
	}

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
