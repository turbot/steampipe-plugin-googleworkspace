package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"google.golang.org/api/people/v1"
)

//// TABLE DEFINITION

func tableGoogleWorkspacePeopleDirectoryPeople(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_people_directory_people",
		Description: "Domain contacts in the authenticated user's domain directory.",
		List: &plugin.ListConfig{
			Hydrate:           listPeopleDirecoryPeople,
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		Columns: peopleContacts(),
	}
}

//// LIST FUNCTION

func listPeopleDirecoryPeople(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := PeopleService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Define fields the API should return
	personFields := "addresses,biographies,birthdays,calendarUrls,clientData,coverPhotos,emailAddresses,events,externalIds,genders,interests,locations,memberships,metadata,miscKeywords,names,nicknames,occupations,organizations,phoneNumbers,photos,relations,sipAddresses,skills,urls,userDefined"

	// By default, API can return maximum 1000 records in a single page
	maxResult := int64(1000)

	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < maxResult {
			maxResult = *limit
		}
	}

	resp := service.People.ListDirectoryPeople().ReadMask(personFields).Sources("DIRECTORY_SOURCE_TYPE_DOMAIN_PROFILE").PageSize(maxResult)
	if err := resp.Pages(ctx, func(page *people.ListDirectoryPeopleResponse) error {
		for _, people := range page.People {
			// Since, 'names', 'birthdays', 'genders' and 'biographies' are singleton fields
			var conn contacts
			if people.Names != nil {
				conn.Name = *people.Names[0]
			}
			if people.Birthdays != nil {
				conn.Birthday = *people.Birthdays[0]
			}
			if people.Genders != nil {
				conn.Gender = *people.Genders[0]
			}
			if people.Biographies != nil {
				conn.Biography = *people.Biographies[0]
			}
			d.StreamListItem(
				ctx,
				contacts{
					conn.Name,
					conn.Birthday,
					conn.Gender,
					conn.Biography,
					*people,
				})

			// Check if the context is cancelled for query
			// Break for loop if requested no of results achieved
			if plugin.IsCancelled(ctx) {
				page.NextPageToken = ""
				break
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, nil
}
