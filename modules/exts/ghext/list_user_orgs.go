package ghext

import (
	"fmt"
	"strings"

	"github.com/cli/go-gh"
)

/*
List organizations for the authenticated user

https://docs.github.com/en/rest/orgs/orgs?apiVersion=2022-11-28#list-organizations-for-the-authenticated-user

	curl -L \
		-H "Accept: application/vnd.github+json" \
		-H "Authorization: Bearer <YOUR-TOKEN>"\
		-H "X-GitHub-Api-Version: 2022-11-28" \
		https://api.github.com/user/orgs
*/

type Organizations []struct {
	Login string `json:"login"`
	// ID               int    `json:"id"`
	// NodeID           string `json:"node_id"`
	// URL              string `json:"url"`
	// ReposURL         string `json:"repos_url"`
	// EventsURL        string `json:"events_url"`
	// HooksURL         string `json:"hooks_url"`
	// IssuesURL        string `json:"issues_url"`
	// MembersURL       string `json:"members_url"`
	// PublicMembersURL string `json:"public_members_url"`
	// AvatarURL        string `json:"avatar_url"`
	// Description      string `json:"description"`
}

func ListUserOrganizations() (string, error) {
	client, err := gh.RESTClient(nil)
	if err != nil {
		return "", err
	}
	response := Organizations{}
	err = client.Get("user/orgs", &response)
	if err != nil {
		return "", err
	}
	res := make([]string, len(response))
	for i, org := range response {
		res[i] = fmt.Sprintf("[+] %-30s https://github.com/%s", org.Login, org.Login)
	}
	return strings.Join(res, "\r\n"), nil
}
