package bitbucket

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	config "github.com/bradmccoydev/terraform-plan-validator/util"
)

func CommentOnPR(baseURL string, workspace string, repoSlug string, id string, cfg config.Config) string {
	url := fmt.Sprintf("%v/%v/%v/pullrequests/%v/", baseURL, workspace, repoSlug, id)
	fmt.Println(url)
	encoded := base64.URLEncoding.EncodeToString([]byte("user:"))
	credentials := fmt.Sprintf("Basic %v", encoded)
	message := "test"
	var jsonStr = []byte(fmt.Sprintf(`{"content": {"raw": "%v"}}`, message))
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-ExperimentalApi", "opt-in")
	req.Header.Set("Authorization", credentials)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
