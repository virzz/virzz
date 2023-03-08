package ghext

import (
	"bytes"
	"fmt"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/repository"
	"github.com/goccy/go-json"
	"github.com/virzz/logger"
)

/*
Transfer a repository

https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#transfer-a-repository

	curl -L \
		-X POST \
		-H "Accept: application/vnd.github+json" \
		-H "Authorization: Bearer <YOUR-TOKEN>"\
		-H "X-GitHub-Api-Version: 2022-11-28" \
		https://api.github.com/repos/OWNER/REPO/transfer \
		-d '{"new_owner":"github","team_ids":[12,345],"new_name":"octorepo"}'
*/
type Transfer struct {
	NewOwner string `json:"new_owner"`
}

func TransferRepository(newOwner, owner, repoName string) (res string, err error) {
	var currRepo repository.Repository
	if owner == "" {
		currRepo, err = gh.CurrentRepository()
		if err != nil {
			logger.Warn(err)
		} else {
			owner = currRepo.Owner()
		}
	}
	if repoName == "" && currRepo != nil {
		repoName = currRepo.Name()
	}
	if repoName == "" || owner == "" || repoName == "" {
		err = fmt.Errorf("repo name or owner or new owner is empty")
		return
	}

	client, err := gh.RESTClient(nil)
	if err != nil {
		return
	}
	transfer := Transfer{NewOwner: newOwner}
	transferBytes, err := json.Marshal(transfer)
	if err != nil {
		return
	}
	body := bytes.NewReader(transferBytes)
	err = client.Post(fmt.Sprintf("repos/%s/%s/transfer", owner, repoName), body, nil)
	if err != nil {
		return
	}
	res = fmt.Sprintf("Transfer [%s/%s] to [%s/%s] success\n", owner, repoName, newOwner, repoName)
	return
}
