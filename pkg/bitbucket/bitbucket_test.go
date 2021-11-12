package bitbucket

import (
	"fmt"
	"log"
	"testing"

	config "github.com/bradmccoydev/terraform-plan-validator/util"
)

func TestBitbucketAPI(t *testing.T) {
	testCases := []struct {
		baseURL   string
		workspace string
		repoSlug  string
		id        string
		err       error
	}{
		{"https://api.bitbucket.org/2.0/repositories", "moula", "infrastructure-as-code", "73", nil},
		{"https://api.bitbucket.org/2.0/repositories", "moula", "infrastructure-as-code", "73", nil},
	}

	for _, tc := range testCases {
		config, err := config.LoadConfig("./../../")
		if err != nil {
			log.Fatal("cannot load config:", err)
		}

		result := CommentOnPR(tc.baseURL, tc.workspace, tc.repoSlug, tc.id, *config)
		fmt.Println(tc.id, result)
	}
}
