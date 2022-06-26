package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

// push
type EventPayload struct {
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Name              string `json:"name"`
			Email             string `json:"email"`
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		HTMLURL          string        `json:"html_url"`
		Description      string        `json:"description"`
		Fork             bool          `json:"fork"`
		URL              string        `json:"url"`
		ForksURL         string        `json:"forks_url"`
		KeysURL          string        `json:"keys_url"`
		CollaboratorsURL string        `json:"collaborators_url"`
		TeamsURL         string        `json:"teams_url"`
		HooksURL         string        `json:"hooks_url"`
		IssueEventsURL   string        `json:"issue_events_url"`
		EventsURL        string        `json:"events_url"`
		AssigneesURL     string        `json:"assignees_url"`
		BranchesURL      string        `json:"branches_url"`
		TagsURL          string        `json:"tags_url"`
		BlobsURL         string        `json:"blobs_url"`
		GitTagsURL       string        `json:"git_tags_url"`
		GitRefsURL       string        `json:"git_refs_url"`
		TreesURL         string        `json:"trees_url"`
		StatusesURL      string        `json:"statuses_url"`
		LanguagesURL     string        `json:"languages_url"`
		StargazersURL    string        `json:"stargazers_url"`
		ContributorsURL  string        `json:"contributors_url"`
		SubscribersURL   string        `json:"subscribers_url"`
		SubscriptionURL  string        `json:"subscription_url"`
		CommitsURL       string        `json:"commits_url"`
		GitCommitsURL    string        `json:"git_commits_url"`
		CommentsURL      string        `json:"comments_url"`
		IssueCommentURL  string        `json:"issue_comment_url"`
		ContentsURL      string        `json:"contents_url"`
		CompareURL       string        `json:"compare_url"`
		MergesURL        string        `json:"merges_url"`
		ArchiveURL       string        `json:"archive_url"`
		DownloadsURL     string        `json:"downloads_url"`
		IssuesURL        string        `json:"issues_url"`
		PullsURL         string        `json:"pulls_url"`
		MilestonesURL    string        `json:"milestones_url"`
		NotificationsURL string        `json:"notifications_url"`
		LabelsURL        string        `json:"labels_url"`
		ReleasesURL      string        `json:"releases_url"`
		DeploymentsURL   string        `json:"deployments_url"`
		CreatedAt        int           `json:"created_at"`
		UpdatedAt        time.Time     `json:"updated_at"`
		PushedAt         int           `json:"pushed_at"`
		GitURL           string        `json:"git_url"`
		SSHURL           string        `json:"ssh_url"`
		CloneURL         string        `json:"clone_url"`
		SvnURL           string        `json:"svn_url"`
		Homepage         string        `json:"homepage"`
		Size             int           `json:"size"`
		StargazersCount  int           `json:"stargazers_count"`
		WatchersCount    int           `json:"watchers_count"`
		Language         string        `json:"language"`
		HasIssues        bool          `json:"has_issues"`
		HasProjects      bool          `json:"has_projects"`
		HasDownloads     bool          `json:"has_downloads"`
		HasWiki          bool          `json:"has_wiki"`
		HasPages         bool          `json:"has_pages"`
		ForksCount       int           `json:"forks_count"`
		MirrorURL        interface{}   `json:"mirror_url"`
		Archived         bool          `json:"archived"`
		Disabled         bool          `json:"disabled"`
		OpenIssuesCount  int           `json:"open_issues_count"`
		License          interface{}   `json:"license"`
		AllowForking     bool          `json:"allow_forking"`
		IsTemplate       bool          `json:"is_template"`
		Topics           []interface{} `json:"topics"`
		Visibility       string        `json:"visibility"`
		Forks            int           `json:"forks"`
		OpenIssues       int           `json:"open_issues"`
		Watchers         int           `json:"watchers"`
		DefaultBranch    string        `json:"default_branch"`
		Stargazers       int           `json:"stargazers"`
		MasterBranch     string        `json:"master_branch"`
	} `json:"repository"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
	Sender struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
	Created bool        `json:"created"`
	Deleted bool        `json:"deleted"`
	Forced  bool        `json:"forced"`
	BaseRef interface{} `json:"base_ref"`
	Compare string      `json:"compare"`
	Commits []struct {
		ID        string    `json:"id"`
		TreeID    string    `json:"tree_id"`
		Distinct  bool      `json:"distinct"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Added    []interface{} `json:"added"`
		Removed  []interface{} `json:"removed"`
		Modified []string      `json:"modified"`
	} `json:"commits"`
	HeadCommit struct {
		ID        string    `json:"id"`
		TreeID    string    `json:"tree_id"`
		Distinct  bool      `json:"distinct"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Added    []interface{} `json:"added"`
		Removed  []interface{} `json:"removed"`
		Modified []string      `json:"modified"`
	} `json:"head_commit"`
}

// TODO: pullrequest

func HandleLambdaEvent(request events.LambdaFunctionURLRequest) {
	godotenv.Load()
	// requestBody := &RequestBody{
	// 	channel: os.Getenv("CHANNEL_ID"),
	// 	text:    "Hello world",
	// }
	// jsonString, err := json.Marshal(requestBody)
	// if err != nil {
	// 	panic("Error")
	// }
	req, err := http.NewRequest("POST", os.Getenv("SLACK_URL"), nil)
	if err != nil {
		panic(("Error"))
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))

	dataBytes := ([]byte)(request.Body)
	data := new(EventPayload)

	fmt.Printf("request.Body -> %v\n", request.Body)
	fmt.Printf("dataBytes -> %v\n", dataBytes)
	fmt.Printf("data -> %v\n", data)

	if err := json.Unmarshal(dataBytes, data); err != nil {
		panic("Error!")
	}

	params := req.URL.Query()
	params.Add("channel", os.Getenv("CHANNEL_ID"))
	params.Add("text", "New commit was pushed to "+strings.Split(data.Ref, "/")[2]+" by "+data.Pusher.Name+"\n"+data.Commits[0].URL)
	req.URL.RawQuery = params.Encode()

	fmt.Printf("request -> %v\n", req)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		panic("Error")
	}
	defer res.Body.Close()

	resArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Error")
	}

	fmt.Printf("%#v", string(resArray))
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
