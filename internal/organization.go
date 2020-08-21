package internal

import (
	"context"
	"github.com/hashicorp/go-tfe"
)

func GetOrganizations(client *tfe.Client) ([]*tfe.Organization, error) {
	organizations := make([]*tfe.Organization, 0)

	currentPage := 0
	totalPages := 1
	pageSize := 10

	for currentPage < totalPages {
		orgPage, err := getOrganizationPage(client, tfe.OrganizationListOptions{
			ListOptions: tfe.ListOptions{
				PageNumber: currentPage,
				PageSize:   pageSize,
			},
		})

		if err != nil { return nil, err }

		organizations = append(organizations, orgPage.Items...)

		totalPages = orgPage.TotalPages
		currentPage++
	}

	return organizations, nil
}

func getOrganizationPage(client *tfe.Client, options tfe.OrganizationListOptions) (*tfe.OrganizationList, error) {
	organizations, err := client.Organizations.List(context.Background(), options)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}
