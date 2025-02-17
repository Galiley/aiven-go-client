// Package aiven provides a client for using the Aiven API.
package aiven

type (
	// OrganizationsHandler is the client which interacts with the Organizations API on Aiven.
	OrganizationsHandler struct {
		// client is the API client to use.
		client *Client
	}

	// OrganizationInfo is a response from Aiven for a single organization.
	OrganizationInfo struct {
		APIResponse

		// ID is the unique identifier of the organization.
		ID string `json:"organization_id"`
		// Name is the name of the organization.
		Name string `json:"organization_name"`
		// AccountID is the unique identifier of the account.
		AccountID string `json:"account_id"`
		// CreateTime is the time when the organization was created.
		CreateTime string `json:"create_time"`
		// UpdateTime is the time when the organization was last updated.
		UpdateTime string `json:"update_time"`
	}
)

// Get returns information about the specified organization.
func (h *OrganizationsHandler) Get(id string) (*OrganizationInfo, error) {
	path := buildPath("organization", id)

	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r OrganizationInfo

	if err = checkAPIResponse(bts, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
