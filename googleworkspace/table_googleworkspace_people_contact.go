package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"google.golang.org/api/people/v1"
)

func peopleContacts() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "resource_name",
			Description: "The resource name for the contact group, assigned by the server.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "display_name",
			Description: "The display name formatted according to the locale specified by the viewer's account.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Name.DisplayName").NullIfZero(),
		},
		{
			Name:        "given_name",
			Description: "The given name of the user contact.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Name.GivenName").NullIfZero(),
		},
		{
			Name:        "primary_email_address",
			Description: "The primary email address of the user contact.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.From(extractPrimaryEmailAddress),
		},
		{
			Name:        "gender",
			Description: "The gender for the person.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Gender.Value").NullIfZero(),
		},
		{
			Name:        "birthday",
			Description: "The date of the birthday.",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Birthday.Date").NullIfZero(),
		},
		{
			Name:        "email_addresses",
			Description: "The person's email addresses.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "addresses",
			Description: "The person's street addresses.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "biography",
			Description: "The person's biography.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "calendar_urls",
			Description: "The person's calendar URLs.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "client_data",
			Description: "The person's client data.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "cover_photos",
			Description: "The person's cover photos.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "events",
			Description: "The person's events.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "external_ids",
			Description: "The person's external IDs.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "interests",
			Description: "The person's interests.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "locations",
			Description: "The person's locations.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "memberships",
			Description: "The person's group memberships.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "metadata",
			Description: "Metadata about the person.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "nicknames",
			Description: "The person's nicknames.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "occupations",
			Description: "The person's occupations.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "organizations",
			Description: "The person's past or current organizations.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "phone_numbers",
			Description: "The person's phone numbers.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "photos",
			Description: "The person's photos.",
			Type:        proto.ColumnType_JSON,
		},
	}
}

//// TABLE DEFINITION

func tableGoogleWorkspacePeopleContact(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_people_contact",
		Description: "Contacts owned by the authenticated user.",
		List: &plugin.ListConfig{
			Hydrate:           listPeopleContacts,
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		Columns: peopleContacts(),
	}
}

type contacts = struct {
	Name      people.Name
	Birthday  people.Birthday
	Gender    people.Gender
	Biography people.Biography
	people.Person
}

//// LIST FUNCTION

func listPeopleContacts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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

	resp := service.People.Connections.List("people/me").PersonFields(personFields).PageSize(maxResult)
	if err := resp.Pages(ctx, func(page *people.ListConnectionsResponse) error {
		for _, connection := range page.Connections {
			// Since, 'names', 'birthdays', 'genders' and 'biographies' are singleton fields
			var conn contacts
			if connection.Names != nil {
				conn.Name = *connection.Names[0]
			}
			if connection.Birthdays != nil {
				conn.Birthday = *connection.Birthdays[0]
			}
			if connection.Genders != nil {
				conn.Gender = *connection.Genders[0]
			}
			if connection.Biographies != nil {
				conn.Biography = *connection.Biographies[0]
			}
			d.StreamListItem(
				ctx,
				contacts{
					conn.Name,
					conn.Birthday,
					conn.Gender,
					conn.Biography,
					*connection,
				})

			// Context can be cancelled due to manual cancellation or the limit has been hit
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

//// TRANSFORM FUNCTIONS

func extractPrimaryEmailAddress(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(contacts)

	emailAddresses := data.EmailAddresses

	for _, email := range emailAddresses {
		if email.Metadata != nil && email.Metadata.Primary {
			return email.Value, nil
		}
	}

	return nil, nil
}
