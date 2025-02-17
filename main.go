package main

import (
	"fmt"

	"github.com/atlanhq/atlan-go/atlan/assets"
	"github.com/atlanhq/atlan-go/atlan/model/structs"
)

func main() {
	ctx := assets.NewContext()
	ctx.EnableLogging("debug")

	// Add a schedule directly on run

	miner := assets.NewSnowflakeMiner("default/snowflake/1739484068").
		Direct(1739491200, "snowflake-database", "ACCOUNT_USAGE").
		ExcludeUsers([]string{"karanjot.singh"}).
		PopularityWindow(30).
		NativeLineage(true).
		CustomConfig(map[string]interface{}{
			"test":    true,
			"feature": 1234,
		}).
		ToWorkflow()

	Schedule := structs.WorkflowSchedule{CronSchedule: "45 5 * * *", Timezone: "Europe/Paris"}

	// Run the workflow
	response, err := ctx.WorkflowClient.Run(miner, &Schedule)
	if err != nil {
		fmt.Println("Error running workflow:", err)
		return
	}
	fmt.Println(response.Spec)

	/*
		// Running Snowflake Miner
		miner := assets.NewSnowflakeMiner("default/snowflake/1739484068").
			Direct(1739491200, "snowflake-database", "ACCOUNT_USAGE").
			ExcludeUsers([]string{"karanjot.singh"}).
			PopularityWindow(30).
			NativeLineage(true).
			CustomConfig(map[string]interface{}{
				"test":    true,
				"feature": 1234,
			}).
			ToWorkflow()

		// Run the workflow
		response, err := ctx.WorkflowClient.Run(miner, nil)
		if err != nil {
			fmt.Println("Error running workflow:", err)
			return
		}

		fmt.Println("Workflow started successfully:", response)
	*/
	/*
		// Update a workflow
		result, _ := ctx.WorkflowClient.FindByID("csa-admin-export-1739443119")

		workflowTask := result.Source.Spec.Templates[0].DAG.Tasks[0]
		workflowParams := workflowTask.Arguments.Parameters

		fmt.Println(workflowTask)
		fmt.Println(workflowParams)

		for _, option := range workflowParams {
			if option.Name == "enable-lineage" {
				option.Value = true
				fmt.Println(option)
			}
		}

		response, err := ctx.WorkflowClient.Update(result.ToWorkflow())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response)
	*/
	/*
		// Delete a Workflow
		_ = ctx.WorkflowClient.Delete("csa-admin-export-1739368706")
	*/
	/*
		// Stop a running workflow
		runs, err := ctx.WorkflowClient.GetRuns("csa-admin-export-1739368706", atlan.AtlanWorkflowPhaseRunning, 0, 100)
		if err != nil {
			fmt.Println(err)
		}
		response, err := ctx.WorkflowClient.Stop(runs[0].ID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response)

	*/
	/*
		// Retrieve runs by their phase:
		result, err := ctx.WorkflowClient.GetRuns("csa-admin-export-1739368706", atlan.AtlanWorkflowPhaseSuccess, 0, 100)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	*/
	/*
		// Retrieve an existing workflow latest run:
		result, err := ctx.WorkflowClient.FindCurrentRun("csa-admin-export-1739368706")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	*/
	/*
		// Retrieve an existing workflow latest run:
		result, err := ctx.WorkflowClient.FindLatestRun("csa-admin-export-1739172254")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*result.Source.Metadata.CreationTimestamp)

	*/
	/*
		// Retrieve an existing workflow run by its ID:
		result, err := ctx.WorkflowClient.FindRunByID("csa-admin-export-1739172254-skdzt")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*result.Source.Metadata.CreationTimestamp)
	*/
	/*
		// Retrieve an existing workflow by its ID:
		result, err := ctx.WorkflowClient.FindByID("csa-admin-export-1739172254\n\n")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)


	*/
	/*
		// Retrieve existing workflows by its type:
		result, err := ctx.WorkflowClient.FindByType(atlan.WorkflowPackageSnowflakeMiner, 5)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)

	*/
	/*
		// Find the GUID of a specific policy in a persona
		PurposeName := "Test-go-sdk-Purpose"
		result, atlanErr := assets.FindPurposesByName(PurposeName)
		if atlanErr != nil {
			log.Fatal(atlanErr)
		}
		fmt.Println(*result.Entities[0].Guid)
		purpose, _ := assets.GetByGuid[*assets.Purpose](*result.Entities[0].Guid)
		for _, policy := range *purpose.Policies {
			println("Policy Found: Guid:", *policy.Guid)
			println("Policy Found: Name:", *policy.DisplayName)
		}
	*/
	/*
		// Find the GUID of a specific policy in a persona
		PersonaName := "Test Persona - Go-sdk"
		result, atlanErr := assets.FindPersonasByName(PersonaName)
		if atlanErr != nil {
			log.Fatal(atlanErr)
		}
		fmt.Println(*result.Entities[0].Guid)
		persona, _ := assets.GetByGuid[*assets.Persona](*result.Entities[0].Guid)
		for _, policy := range *persona.Policies {
			println("Policy Found: Guid:", *policy.Guid)
			println("Policy Found: Name:", *policy.DisplayName)
		}

	*/
	/*
		// Update Policy through Persona/Purpose
		response, atlanErr := assets.NewFluentSearch().
			AssetType("AuthPolicy").
			Where(ctx.AuthPolicy.POLICY_CATEGORY.Eq("persona")).
			Where(ctx.AuthPolicy.GUID.Eq("da6bc10c-c2f0-4945-9592-9b4cc75cbc7b")).
			//Where(ctx.AuthPolicy.POLICY_RESOURCES.StartsWith("entity:default/snowflake/1738006574", nil)).
			IncludeOnResults(ctx.AuthPolicy.NAME.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.ACCESS_CONTROL.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.POLICY_RESOURCES.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.CONNECTION_QUALIFIED_NAME.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.POLICY_TYPE.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.POLICY_SUB_CATEGORY.GetAtlanFieldName()).
			IncludeOnRelations(ctx.AccessControl.IS_ACCESS_CONTROL_ENABLED.GetAtlanFieldName()).
			IncludeOnRelations(ctx.AccessControl.NAME.GetAtlanFieldName()).
			Execute()
		if atlanErr != nil {
			log.Fatal(atlanErr)
		}
		for _, entity := range response[0].Entities {
			if entity.TypeName != nil && *entity.TypeName == "AuthPolicy" {
				fmt.Println("AuthPolicy Found: Name:", *entity.Name, "QualifiedName:", *entity.QualifiedName)
				fmt.Println("AuthPolicy guid:", *entity.Guid)
				fmt.Println("Policy Resources:", *entity.PolicyResources)
				fmt.Println(*entity.AccessControl.TypeName)
				entity.PolicyType = &atlan.AuthPolicyTypeDeny
				_, err := assets.Save(&entity)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	*/
	/*
		// Retrieving policies from a purpose
		response, atlanErr := assets.NewFluentSearch().
			AssetType("AuthPolicy").
			Where(ctx.AuthPolicy.POLICY_CATEGORY.Eq("purpose")).
			//Where(ctx.AuthPolicy.POLICY_RESOURCES.StartsWith("tag:TLwvfawpBhZb6JIQgVc0Fn", nil)).
			IncludeOnResults(ctx.AuthPolicy.NAME.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.ACCESS_CONTROL.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.POLICY_RESOURCES.GetAtlanFieldName()).
			IncludeOnRelations(ctx.AccessControl.IS_ACCESS_CONTROL_ENABLED.GetAtlanFieldName()).
			IncludeOnRelations(ctx.AccessControl.NAME.GetAtlanFieldName()).
			Execute()
		if atlanErr != nil {
			log.Fatal(atlanErr)
		}
		for _, entity := range response[0].Entities {
			if entity.TypeName != nil && *entity.TypeName == "AuthPolicy" {
				println("AuthPolicy Found: Name:", *entity.Name, "QualifiedName:", *entity.QualifiedName)
				println("AuthPolicy guid:", *entity.Guid)
				println("Policy Resources:", *entity.PolicyResources)
			}
		}

	*/
	/*
		// Retrieving policies from a persona
		response, atlanErr := assets.NewFluentSearch().
			AssetType("AuthPolicy").
			Where(ctx.AuthPolicy.POLICY_CATEGORY.Eq("persona")).
			Where(ctx.AuthPolicy.POLICY_RESOURCES.StartsWith("entity:default/snowflake/1738006457", nil)).
			IncludeOnResults(ctx.AuthPolicy.NAME.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.ACCESS_CONTROL.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.POLICY_RESOURCES.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.CONNECTION_QUALIFIED_NAME.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.POLICY_TYPE.GetAtlanFieldName()).
			IncludeOnResults(ctx.AuthPolicy.POLICY_SUB_CATEGORY.GetAtlanFieldName()).
			IncludeOnRelations(ctx.AccessControl.IS_ACCESS_CONTROL_ENABLED.GetAtlanFieldName()).
			IncludeOnRelations(ctx.AccessControl.NAME.GetAtlanFieldName()).
			Execute()
		if atlanErr != nil {
			log.Fatal(atlanErr)
		}
		for _, entity := range response[0].Entities {
			if entity.TypeName != nil && *entity.TypeName == "AuthPolicy" {
				fmt.Println("AuthPolicy Found: Name:", *entity.Name, "QualifiedName:", *entity.QualifiedName)
				fmt.Println("AuthPolicy guid:", *entity.Guid)
				fmt.Println("Policy Resources:", *entity.PolicyResources)
			}
		}

	*/
	/*
		// Delete a Purpose
		assets.PurgeByGuid([]string{"a947beac-ae4d-4d1c-a9f8-efd9c55fe768"})
	*/
	/*
		// List purposes
		response, atlanErr := assets.NewFluentSearch(). //
								PageSizes(20).
								ActiveAssets().
								AssetType("Purpose"). //
								Execute()             //
		if atlanErr != nil {
			fmt.Println("Error:", atlanErr)
		}
		for _, entity := range response[0].Entities { //
			if entity.TypeName != nil && *entity.TypeName == "Purpose" {
				// Do something with the Purpose
				fmt.Println("Purpose Found:", *entity.Name, "QualifiedName:", *entity.QualifiedName)
				fmt.Println("Purpose guid:", *entity.Guid)
			}
		}


	*/

	/*
		// Personalize the purpose
		purpose := &assets.Purpose{}
		purpose.Updater("default/0IFqLjT5JqnZrWOp2U1IUB", "Newly Modified Test Purpose", true)
		purpose.DenyAssetTabs = &[]string{atlan.AssetSidebarTabLineage.Name, atlan.AssetSidebarTabRelations.Name, atlan.AssetSidebarTabQueries.Name}
		response, _ := assets.Save(purpose)
		for _, entity := range response.MutatedEntities.UPDATE {
			println("Response:", entity)
			println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)

		}
	*/
	/*
		// Add a Data Policy
		purpose := &assets.Purpose{}
		policy, _ := purpose.CreateDataPolicy(
			"Test Policy for Masking Data",
			"a947beac-ae4d-4d1c-a9f8-efd9c55fe768",
			atlan.AuthPolicyTypeDatamask,
			nil,
			nil,
			true,
		)
		response, err := assets.Save(policy)
		if err != nil {
			println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				println("Response TypeName:", entity.TypeName)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}

		}


	*/
	/*
		// Add a metadata Policy
		purpose := &assets.Purpose{}
		policy, _ := purpose.CreateMetadataPolicy(
			"Test Policy 3",
			"a947beac-ae4d-4d1c-a9f8-efd9c55fe768",
			atlan.AuthPolicyTypeAllow,
			[]atlan.PurposeMetadataAction{
				atlan.PurposeMetadataActionRead,
			},
			nil,
			nil,
			true,
		)
		response, err := assets.Save(policy)
		if err != nil {
			println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				println("Response TypeName:", entity.TypeName)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}

		}


	*/
	/*
		// Activate or Deactivate a Purpose
		purpose := &assets.Purpose{}
		err := purpose.Updater("default/0IFqLjT5JqnZrWOp2U1IUB", "Newly Modified Test Purpose", false)
		if err != nil {
			fmt.Println(err)
			return
		}
		response, err := assets.Save(purpose)
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
		// Update a Purpose
		purpose := &assets.Purpose{}
		err := purpose.Updater("default/0IFqLjT5JqnZrWOp2U1IUB", "Test Purpose Modified", true)
		if err != nil {
			fmt.Println(err)
			return
		}
		DisplayName := "Newly Modified Test Purpose"
		Description := "This is a modified description"
		purpose.Name = &DisplayName
		purpose.Description = &Description
		response, err := assets.Save(purpose)
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
		// Retrieve Persona by Name
		result, _ := assets.FindPurposesByName("Test Purpose - Go-sdk")
		fmt.Println(*result.Entities[0].Name)

	*/
	/*
		// Create a Purpose
		purpose := &assets.Purpose{}
		purpose.Creator("Test Purpose - Go-sdk", []string{"Confidential", "Issue"})
		response, err := assets.Save(purpose)
		if err != nil {
			println("Error:", err)
		} else {
			for _, entity := range response.MutatedEntities.CREATE {
				println("Response:", entity)
				println("Entity ID:", entity.Guid, "Display Text:", entity.DisplayText)
			}
		}
		// a947beac-ae4d-4d1c-a9f8-efd9c55fe768
	*/

	/*
		// Test User Cache
		UserId, err := assets.GetGroupNameForGroupID("58d547d8-3f4d-4b9e-9666-39980f140661")
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(UserId)
	*/
	/*
		// Delete an API Token
		err := ctx.TokenClient.Purge("a853f1d5-f1f4-4cdb-b86d-c61df3ecade6")
		if err != nil {
			println("Error:", err)
		}
	*/
	/*
		// Update an API Token
		token, _ := ctx.TokenClient.GetByName("Test-Token-Updated-4")
		displayName := "Test-Token-Updated-5"
		update, err := ctx.TokenClient.Update(token.GUID, &displayName, nil, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*update.Attributes.DisplayName)


	*/
	/*
		// Retrieve an API Token
		token, atlanErr := ctx.TokenClient.GetByName("Test-Token")
		if atlanErr != nil {
			fmt.Println(atlanErr)
		}
		fmt.Println("Token Client ID:", *token.Attributes.ClientID)

		// Retrieve by ID
		tokenByID, atlanErr := ctx.TokenClient.GetByID("apikey-286e681a-67b5-4996-b161-5737e890c080")
		if atlanErr != nil {
			fmt.Println(atlanErr)
		}
		fmt.Println("Token Display Name:", *tokenByID.Attributes.DisplayName)


	*/
	/*
		// Create Token
		displayName := "Test-Token"
		token, err := ctx.TokenClient.Create(&displayName, nil, nil, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*token.Attributes.AccessToken)


	*/
	/*
		users := []assets.AtlanUser{
			{
				Email:         "test2@atlan.com",
				WorkspaceRole: "$member",
			},
		}

		createdUsers, _ := ctx.UserClient.CreateUsers(users, true)
		fmt.Println(createdUsers)
		for _, user := range createdUsers {
			fmt.Printf("User: %v\n", *user.Username)
		}

	*/

	/*
		err := ctx.GroupClient.Purge("a99f50bc-46bf-4d08-a987-3411ef5cfc33")
		if err != nil {
			fmt.Println(err)
			return
		}

	*/
	/*
		// Remove users from the group
		err := ctx.GroupClient.RemoveUsers("a99f50bc-46bf-4d08-a987-3411ef5cfc33", []string{"b060a754-4d16-4e13-b5a8-ba42f10aee39"})
		if err != nil {
			fmt.Println(err)
			return
		}

	*/
	/*
		// Change User Role
		roleID, _ := assets.GetRoleIDForRoleName("$admin")
		ctx.UserClient.ChangeUserRole("b060a754-4d16-4e13-b5a8-ba42f10aee39", roleID)

	*/
	/*
			// Add User to Groups
			//user, _ := ctx.UserClient.GetByUsername("karanjot.singh")
			//fmt.Println(user.ID)
			ctx.UserClient.AddUserToGroups("b060a754-4d16-4e13-b5a8-ba42f10aee39", []string{"a99f50bc-46bf-4d08-a987-3411ef5cfc33"})

		/*

			// Update a Group

			group := assets.AtlanGroup{}
			tobeUpdated, _ := group.Updater("a99f50bc-46bf-4d08-a987-3411ef5cfc33", "/test_group_-_go-sdk")

			//Name := "TestUpdatedName3"
			Alias := []string{"TestAliasName4"}
			//description := []string{"This is the updated description 1"}

			//tobeUpdated.Attributes.Description = description
			tobeUpdated.Attributes.Alias = Alias
			ctx.GroupClient.Update(tobeUpdated)


	*/

	/*
		// Retrieve groups in User
		user, _ := ctx.UserClient.GetByUsername("karanjot.singh")
		response, _ := ctx.UserClient.GetGroups(user.ID, nil)
		for _, group := range response {
			fmt.Println(*group.ID)
		}

	*/
	/*
		//group, _ := ctx.GroupClient.GetByName("Admins", 10, 0)
		groupID := "a4610821-66e6-4925-8834-52cbbeefdbeb"
		response, _ := ctx.GroupClient.GetMembers(groupID, nil)
		for _, user := range response {
			fmt.Println(*user.Username)
		}

	*/
	/*
		// Get Multiple Users
		users, _ := ctx.UserClient.GetByEmails([]string{"karanjot.singh@atlan.com", "chris@atlan.com"}, 5, 0)
		for _, user := range users {
			fmt.Println(*user.Username)
		}

	*/
	/*
		// Get User by Email
		user, _ := ctx.UserClient.GetByEmail("karanjot.singh@atlan.com", 10, 0) // Maybe change these to provide nil values here
		fmt.Println(user)

	*/
	/*
		// Get User by Username
		users, _ := ctx.UserClient.GetByUsername("karanjot.singh")
		fmt.Println(*users.Username)

	*/
	/*
		// Retrieve all users
		users, _ := ctx.UserClient.GetAll(20, 0, "")
		for _, user := range users {
			fmt.Println(*user.Username)
		}

	*/
	/*
		// Retrieve group by name
		groups, _ := ctx.GroupClient.GetByName("Admins", 10, 0)
		for _, group := range groups {
			// Do Something with the groups
			fmt.Println(*group.Name)
		}

	*/
	/*
		// Retrieve all groups
		groups, _ := ctx.GroupClient.GetAll(10, 1, "createdAt")
		for _, group := range groups {
			// Do Something with the groups
			fmt.Println(*group.Name)
		}

	*/
	/*
		atlanGroup := assets.AtlanGroup{}
		// Create Group
		tobeCreated, _ := atlanGroup.Create("Test Group - Go-sdk")
		response, err := ctx.GroupClient.Create(tobeCreated, nil)
		if err != nil {
			log.Fatalf("Failed to create group: %v", err)
		}
		fmt.Printf("Group GUID: %s\n", response.Group)

	*/
	/*
		// Create User
		usersToCreate := []assets.AtlanUser{
			{
				Email:         "test.user1@atlan.com",
				WorkspaceRole: "$member",
			},
			{
				Email:         "test.user2@atlan.com",
				WorkspaceRole: "$member",
			},
		}

		atlanUser := &assets.AtlanUser{}

		createdUsers, err := atlanUser.CreateUsers(usersToCreate, true)
		if err != nil {
			log.Fatalf("Error creating users: %v", err)
		}

		for _, user := range createdUsers {
			fmt.Printf("Created user: %s with role: %s\n", user.Email, user.WorkspaceRole)
		}

	*/
	/*
		response, atlanErr := assets.GetAll()
		if atlanErr != nil {
			println("Error:", atlanErr)
		} else {
			fmt.Println(response)
		}


	*/
	/*

		response, atlanErr := assets.FindPersonasByName("Test Persona")
		if atlanErr != nil {
			println("Error:", atlanErr)
		} else {
			for _, entity := range response.Entities {
				if entity.TypeName != nil && *entity.TypeName == "Persona" {
					println("Persona Found: Name:", *entity.Name, "QualifiedName:", *entity.QualifiedName)
				}
			}

		}


	*/
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
	// schemaName := "WIDEWORLDIMPORTERS_PURCHASING"
	// dataBaseName := "RAW"
	// dataBaseQualifiedName := "default/snowflake/1723642516/RAW"
	// connectionQualifiedName := "default/snowflake/1723642516"

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
	// structs.GetAll()

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
	// structs.GetAtlanTagCache().RefreshCache()
	// id, _ := structs.GetAtlanTagCache().GetIDForName("Hourly")
	// fmt.Println("Print Response:", id)

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
	// structs.Init()
	/*
		response, err := structs.GetGlossaryByGuid("f273e814-f80e-4699-83f3-9462a153fb14")
		if err != nil {
			println("Error:", err)
		}
		print("Response:", response.Name)
	*/
	// response, err := structs.FindCategoryByName("Oak Ridge", "DDwycTZ007zZYxRajRVDK")

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
	// if err != nil {
	//	fmt.Printf("Error fetching model: %v\n", err)
	//}

	// fmt.Printf("Response: %+v\n", response)

	// ctx := structs.NewContext()

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
	// Modify an existing Glossary
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
	// structs.DeleteByGuid([]string{"024f11b6-a9fa-4f45-84f5-f734c47c4743", "b280b09b-5c28-45c4-a899-d8535fb651eb", "8679e70a-513e-4e2e-9861-4f5559206f36"})
	// structs.DeleteByGuid([]string{"dbe090bd-1549-4cce-98dd-6542138963f1"})
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

	// query := ctx.Glossary.TYPENAME.Eq("AtlasGlossary", nil)
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
	// structs.GetCache().RefreshCache()
	// structs.GetCache()
	// a, _ := structs.GetCache().GetIDForName("PII")
	// fmt.Printf("ID for name is : %s\n", a)
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
	// fmt.Println("Retrieved Typedef:")
	// fmt.Printf("DisplayName: %s\n", t.AtlanTagDefs[1].DisplayName)
	//
}
