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
)

// Push Struct
type PushEventPayload struct {
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

// PullRequest Struct
type PullRequestEventPayload struct {
	Action      string `json:"action"`
	Number      int    `json:"number"`
	PullRequest struct {
		URL      string `json:"url"`
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		HTMLURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
		IssueURL string `json:"issue_url"`
		Number   int    `json:"number"`
		State    string `json:"state"`
		Locked   bool   `json:"locked"`
		Title    string `json:"title"`
		User     struct {
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
		} `json:"user"`
		Body               interface{}   `json:"body"`
		CreatedAt          time.Time     `json:"created_at"`
		UpdatedAt          time.Time     `json:"updated_at"`
		ClosedAt           interface{}   `json:"closed_at"`
		MergedAt           interface{}   `json:"merged_at"`
		MergeCommitSha     interface{}   `json:"merge_commit_sha"`
		Assignee           interface{}   `json:"assignee"`
		Assignees          []interface{} `json:"assignees"`
		RequestedReviewers []interface{} `json:"requested_reviewers"`
		RequestedTeams     []interface{} `json:"requested_teams"`
		Labels             []interface{} `json:"labels"`
		Milestone          interface{}   `json:"milestone"`
		Draft              bool          `json:"draft"`
		CommitsURL         string        `json:"commits_url"`
		ReviewCommentsURL  string        `json:"review_comments_url"`
		ReviewCommentURL   string        `json:"review_comment_url"`
		CommentsURL        string        `json:"comments_url"`
		StatusesURL        string        `json:"statuses_url"`
		Head               struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
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
			} `json:"user"`
			Repo struct {
				ID       int    `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Private  bool   `json:"private"`
				Owner    struct {
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
				HTMLURL                   string        `json:"html_url"`
				Description               string        `json:"description"`
				Fork                      bool          `json:"fork"`
				URL                       string        `json:"url"`
				ForksURL                  string        `json:"forks_url"`
				KeysURL                   string        `json:"keys_url"`
				CollaboratorsURL          string        `json:"collaborators_url"`
				TeamsURL                  string        `json:"teams_url"`
				HooksURL                  string        `json:"hooks_url"`
				IssueEventsURL            string        `json:"issue_events_url"`
				EventsURL                 string        `json:"events_url"`
				AssigneesURL              string        `json:"assignees_url"`
				BranchesURL               string        `json:"branches_url"`
				TagsURL                   string        `json:"tags_url"`
				BlobsURL                  string        `json:"blobs_url"`
				GitTagsURL                string        `json:"git_tags_url"`
				GitRefsURL                string        `json:"git_refs_url"`
				TreesURL                  string        `json:"trees_url"`
				StatusesURL               string        `json:"statuses_url"`
				LanguagesURL              string        `json:"languages_url"`
				StargazersURL             string        `json:"stargazers_url"`
				ContributorsURL           string        `json:"contributors_url"`
				SubscribersURL            string        `json:"subscribers_url"`
				SubscriptionURL           string        `json:"subscription_url"`
				CommitsURL                string        `json:"commits_url"`
				GitCommitsURL             string        `json:"git_commits_url"`
				CommentsURL               string        `json:"comments_url"`
				IssueCommentURL           string        `json:"issue_comment_url"`
				ContentsURL               string        `json:"contents_url"`
				CompareURL                string        `json:"compare_url"`
				MergesURL                 string        `json:"merges_url"`
				ArchiveURL                string        `json:"archive_url"`
				DownloadsURL              string        `json:"downloads_url"`
				IssuesURL                 string        `json:"issues_url"`
				PullsURL                  string        `json:"pulls_url"`
				MilestonesURL             string        `json:"milestones_url"`
				NotificationsURL          string        `json:"notifications_url"`
				LabelsURL                 string        `json:"labels_url"`
				ReleasesURL               string        `json:"releases_url"`
				DeploymentsURL            string        `json:"deployments_url"`
				CreatedAt                 time.Time     `json:"created_at"`
				UpdatedAt                 time.Time     `json:"updated_at"`
				PushedAt                  time.Time     `json:"pushed_at"`
				GitURL                    string        `json:"git_url"`
				SSHURL                    string        `json:"ssh_url"`
				CloneURL                  string        `json:"clone_url"`
				SvnURL                    string        `json:"svn_url"`
				Homepage                  string        `json:"homepage"`
				Size                      int           `json:"size"`
				StargazersCount           int           `json:"stargazers_count"`
				WatchersCount             int           `json:"watchers_count"`
				Language                  string        `json:"language"`
				HasIssues                 bool          `json:"has_issues"`
				HasProjects               bool          `json:"has_projects"`
				HasDownloads              bool          `json:"has_downloads"`
				HasWiki                   bool          `json:"has_wiki"`
				HasPages                  bool          `json:"has_pages"`
				ForksCount                int           `json:"forks_count"`
				MirrorURL                 interface{}   `json:"mirror_url"`
				Archived                  bool          `json:"archived"`
				Disabled                  bool          `json:"disabled"`
				OpenIssuesCount           int           `json:"open_issues_count"`
				License                   interface{}   `json:"license"`
				AllowForking              bool          `json:"allow_forking"`
				IsTemplate                bool          `json:"is_template"`
				Topics                    []interface{} `json:"topics"`
				Visibility                string        `json:"visibility"`
				Forks                     int           `json:"forks"`
				OpenIssues                int           `json:"open_issues"`
				Watchers                  int           `json:"watchers"`
				DefaultBranch             string        `json:"default_branch"`
				AllowSquashMerge          bool          `json:"allow_squash_merge"`
				AllowMergeCommit          bool          `json:"allow_merge_commit"`
				AllowRebaseMerge          bool          `json:"allow_rebase_merge"`
				AllowAutoMerge            bool          `json:"allow_auto_merge"`
				DeleteBranchOnMerge       bool          `json:"delete_branch_on_merge"`
				AllowUpdateBranch         bool          `json:"allow_update_branch"`
				UseSquashPrTitleAsDefault bool          `json:"use_squash_pr_title_as_default"`
			} `json:"repo"`
		} `json:"head"`
		Base struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
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
			} `json:"user"`
			Repo struct {
				ID       int    `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Private  bool   `json:"private"`
				Owner    struct {
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
				HTMLURL                   string        `json:"html_url"`
				Description               string        `json:"description"`
				Fork                      bool          `json:"fork"`
				URL                       string        `json:"url"`
				ForksURL                  string        `json:"forks_url"`
				KeysURL                   string        `json:"keys_url"`
				CollaboratorsURL          string        `json:"collaborators_url"`
				TeamsURL                  string        `json:"teams_url"`
				HooksURL                  string        `json:"hooks_url"`
				IssueEventsURL            string        `json:"issue_events_url"`
				EventsURL                 string        `json:"events_url"`
				AssigneesURL              string        `json:"assignees_url"`
				BranchesURL               string        `json:"branches_url"`
				TagsURL                   string        `json:"tags_url"`
				BlobsURL                  string        `json:"blobs_url"`
				GitTagsURL                string        `json:"git_tags_url"`
				GitRefsURL                string        `json:"git_refs_url"`
				TreesURL                  string        `json:"trees_url"`
				StatusesURL               string        `json:"statuses_url"`
				LanguagesURL              string        `json:"languages_url"`
				StargazersURL             string        `json:"stargazers_url"`
				ContributorsURL           string        `json:"contributors_url"`
				SubscribersURL            string        `json:"subscribers_url"`
				SubscriptionURL           string        `json:"subscription_url"`
				CommitsURL                string        `json:"commits_url"`
				GitCommitsURL             string        `json:"git_commits_url"`
				CommentsURL               string        `json:"comments_url"`
				IssueCommentURL           string        `json:"issue_comment_url"`
				ContentsURL               string        `json:"contents_url"`
				CompareURL                string        `json:"compare_url"`
				MergesURL                 string        `json:"merges_url"`
				ArchiveURL                string        `json:"archive_url"`
				DownloadsURL              string        `json:"downloads_url"`
				IssuesURL                 string        `json:"issues_url"`
				PullsURL                  string        `json:"pulls_url"`
				MilestonesURL             string        `json:"milestones_url"`
				NotificationsURL          string        `json:"notifications_url"`
				LabelsURL                 string        `json:"labels_url"`
				ReleasesURL               string        `json:"releases_url"`
				DeploymentsURL            string        `json:"deployments_url"`
				CreatedAt                 time.Time     `json:"created_at"`
				UpdatedAt                 time.Time     `json:"updated_at"`
				PushedAt                  time.Time     `json:"pushed_at"`
				GitURL                    string        `json:"git_url"`
				SSHURL                    string        `json:"ssh_url"`
				CloneURL                  string        `json:"clone_url"`
				SvnURL                    string        `json:"svn_url"`
				Homepage                  string        `json:"homepage"`
				Size                      int           `json:"size"`
				StargazersCount           int           `json:"stargazers_count"`
				WatchersCount             int           `json:"watchers_count"`
				Language                  string        `json:"language"`
				HasIssues                 bool          `json:"has_issues"`
				HasProjects               bool          `json:"has_projects"`
				HasDownloads              bool          `json:"has_downloads"`
				HasWiki                   bool          `json:"has_wiki"`
				HasPages                  bool          `json:"has_pages"`
				ForksCount                int           `json:"forks_count"`
				MirrorURL                 interface{}   `json:"mirror_url"`
				Archived                  bool          `json:"archived"`
				Disabled                  bool          `json:"disabled"`
				OpenIssuesCount           int           `json:"open_issues_count"`
				License                   interface{}   `json:"license"`
				AllowForking              bool          `json:"allow_forking"`
				IsTemplate                bool          `json:"is_template"`
				Topics                    []interface{} `json:"topics"`
				Visibility                string        `json:"visibility"`
				Forks                     int           `json:"forks"`
				OpenIssues                int           `json:"open_issues"`
				Watchers                  int           `json:"watchers"`
				DefaultBranch             string        `json:"default_branch"`
				AllowSquashMerge          bool          `json:"allow_squash_merge"`
				AllowMergeCommit          bool          `json:"allow_merge_commit"`
				AllowRebaseMerge          bool          `json:"allow_rebase_merge"`
				AllowAutoMerge            bool          `json:"allow_auto_merge"`
				DeleteBranchOnMerge       bool          `json:"delete_branch_on_merge"`
				AllowUpdateBranch         bool          `json:"allow_update_branch"`
				UseSquashPrTitleAsDefault bool          `json:"use_squash_pr_title_as_default"`
			} `json:"repo"`
		} `json:"base"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Issue struct {
				Href string `json:"href"`
			} `json:"issue"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			ReviewComments struct {
				Href string `json:"href"`
			} `json:"review_comments"`
			ReviewComment struct {
				Href string `json:"href"`
			} `json:"review_comment"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
		} `json:"_links"`
		AuthorAssociation   string      `json:"author_association"`
		AutoMerge           interface{} `json:"auto_merge"`
		ActiveLockReason    interface{} `json:"active_lock_reason"`
		Merged              bool        `json:"merged"`
		Mergeable           interface{} `json:"mergeable"`
		Rebaseable          interface{} `json:"rebaseable"`
		MergeableState      string      `json:"mergeable_state"`
		MergedBy            interface{} `json:"merged_by"`
		Comments            int         `json:"comments"`
		ReviewComments      int         `json:"review_comments"`
		MaintainerCanModify bool        `json:"maintainer_can_modify"`
		Commits             int         `json:"commits"`
		Additions           int         `json:"additions"`
		Deletions           int         `json:"deletions"`
		ChangedFiles        int         `json:"changed_files"`
	} `json:"pull_request"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
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
		CreatedAt        time.Time     `json:"created_at"`
		UpdatedAt        time.Time     `json:"updated_at"`
		PushedAt         time.Time     `json:"pushed_at"`
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
	} `json:"repository"`
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
}

func Handler(request events.LambdaFunctionURLRequest) {
	// httpリクエストの作成
	slackRequest, err := http.NewRequest("POST", os.Getenv("SLACK_URL"), nil)
	if err != nil {
		panic("Error: Can't make NewHttpRequest.")
	}
	slackRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	slackRequest.Header.Set("Authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))
	githubEventType := request.Headers["x-github-event"]
	githubPayloadBytes := ([]byte)(request.Body)
	slackParams := slackRequest.URL.Query()
	slackParams.Add("channel", os.Getenv("CHANNEL_ID"))
	if githubEventType == "push" {
		pushData := new(PushEventPayload)
		if err := json.Unmarshal(githubPayloadBytes, pushData); err != nil {
			panic("Error: Can't unmarshal push json data.")
		}
		slackParams.Add("text", "New commit was pushed to `"+strings.Split(pushData.Ref, "/")[2]+"` by `"+pushData.Pusher.Name+"`\n"+pushData.Commits[0].URL)
	} else {
		pullRequestData := new(PullRequestEventPayload)
		if err := json.Unmarshal(githubPayloadBytes, pullRequestData); err != nil {
			panic("Error: Can't unmarshal pullRequest json data.")
		}
		slackParams.Add("text", "This action was happened: `"+pullRequestData.Action+"` by `"+pullRequestData.Sender.Login+"`\n"+pullRequestData.PullRequest.HTMLURL)
	}
	slackRequest.URL.RawQuery = slackParams.Encode()

	// httpリクエストの実行
	httpClient := new(http.Client)
	slackResponse, err := httpClient.Do(slackRequest)
	if err != nil {
		panic("Error: httpRequest stopped by Error.")
	}
	defer slackResponse.Body.Close()

	// httpリクエストの結果
	responseArray, err := ioutil.ReadAll(slackResponse.Body)
	if err != nil {
		panic("Error: Can't get response.")
	}
	fmt.Printf("%v", string(responseArray))
}

func main() {
	lambda.Start(Handler)
}
