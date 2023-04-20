package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/go-github/v51/github"
	"github.com/hashicorp/go-version"
	"github.com/lonegunmanb/terraform-tls-schema/v4/schema"
	"golang.org/x/oauth2"
)

func main() {
	refreshFlag := flag.Bool("refresh", false, "Refresh tls schema")
	commitFlag := flag.String("commit", "", "Commit schema with specified version")

	// Parse flags
	flag.Parse()

	if *refreshFlag {
		v := refreshSchema()
		fmt.Println(v.String())
	} else if *commitFlag != "" {
		commitSchema(*commitFlag)
	} else {
		log.Fatal("No valid flag provided. Use -refresh or -commit=<schema_version>")
	}
}

func commitSchema(schemaVersion string) {
	pat := os.Getenv("GITHUB_TOKEN")
	if pat == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}
	var auth transport.AuthMethod = &githttp.BasicAuth{
		Username: "github-actions[bot]",
		Password: pat,
	}
	tagExists, err := checkGitHubTag(schemaVersion)
	if err != nil {
		log.Fatalf("Failed to check GitHub tag: %v", err)
	}

	if !tagExists {
		commitMsg := fmt.Sprintf("update schema to version %s", schemaVersion)

		repo, err := git.PlainOpen(".")
		if err != nil {
			log.Fatalf("Failed to open the git repository: %v", err)
		}

		w, err := repo.Worktree()
		if err != nil {
			log.Fatalf("Failed to get the worktree: %v", err)
		}

		_, err = w.Add(".")
		if err != nil {
			log.Fatalf("Failed to add changes to the staging area: %v", err)
		}

		commit, err := w.Commit(commitMsg, &git.CommitOptions{
			Author: &object.Signature{
				Name:  "github-actions[bot]",
				Email: "github-actions[bot]@users.noreply.github.com",
				When:  time.Now(),
			},
		})
		if err != nil {
			log.Fatalf("Failed to commit changes: %v", err)
		}
		obj, err := repo.CommitObject(commit)
		if err != nil {
			log.Fatalf("Failed to get the commit object: %v", err)
		}
		remoteURL, err := convertToHttpsUrl(repo)
		if err != nil {
			log.Fatalf("Failed to convert remote URL to HTTPS: %v", err)
		}
		err = repo.Push(&git.PushOptions{
			RemoteName: "origin",
			Auth:       auth,
			Progress:   os.Stdout,
			RemoteURL:  remoteURL,
		})
		if err != nil {
			log.Fatalf("Failed to push changes: %v", err)
		}
		tag := fmt.Sprintf("v%s", schemaVersion)
		// Check if the local tag exists and delete it
		_, err = repo.Tag(tag)
		if err == nil {
			err = deleteLocalTag(repo, tag)
			if err != nil {
				log.Fatalf("Failed to delete existing local tag: %v", err)
			}
		}

		_, err = repo.CreateTag(tag, obj.Hash, &git.CreateTagOptions{
			Tagger:  &obj.Author,
			Message: commitMsg,
		})
		if err != nil {
			log.Fatalf("Failed to create a new tag: %v", err)
		}

		if err != nil {
			log.Fatalf("Failed to convert remote URL to HTTPS: %v", err)
		}
		tagRef := plumbing.ReferenceName("refs/tags/" + tag)
		pushOptions := &git.PushOptions{
			RemoteName: "origin",
			Auth:       auth,
			RefSpecs:   []config.RefSpec{config.RefSpec(tagRef + ":" + tagRef)},
			RemoteURL:  remoteURL,
			Progress:   os.Stdout,
		}
		err = repo.Push(pushOptions)
		if err != nil {
			log.Fatalf("Failed to push tag: %v", err)
		}
	}
}

func refreshSchema() *version.Version {
	err := os.RemoveAll("./generated")
	if err != nil {
		log.Fatalf("Failed to remove 'generated' folder: %v", err)
	}
	v, err := schema.RefreshSchema("generated")
	if err != nil {
		log.Fatalf("Failed to refresh provider schema: %v", err)
	}
	return v
}

func convertToHttpsUrl(repo *git.Repository) (string, error) {
	remote, err := repo.Remote("origin")
	if err != nil {
		return "", err
	}
	remoteURL := remote.Config().URLs[0]

	if strings.HasPrefix(remoteURL, "https://") {
		return remoteURL, nil
	}
	if strings.HasPrefix(remoteURL, "git@") {
		httpsURL := strings.Replace(remoteURL, ":", "/", 1)
		httpsURL = strings.Replace(httpsURL, "git@", "https://", 1)
		return httpsURL, nil
	}
	return "", fmt.Errorf("remote URL is not in the expected format")
}

func deleteLocalTag(repo *git.Repository, tagName string) error {
	if err := repo.DeleteTag(tagName); err != nil {
		return fmt.Errorf("failed to delete local tag: %w", err)
	}
	return nil
}

func checkGitHubTag(schemaVersion string) (bool, error) {
	var tc *http.Client
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken != "" {
		// Replace "your-access-token" with your personal access token.
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubToken},
		)
		tc = oauth2.NewClient(context.Background(), ts)
	}
	client := github.NewClient(tc)
	options := &github.ListOptions{PerPage: 100}

	for {
		tags, resp, err := client.Repositories.ListTags(context.Background(), "lonegunmanb", "terraform-tls-schema", options)
		if err != nil {
			return false, err
		}
		for _, tag := range tags {
			_, err := version.NewVersion(*tag.Name)
			if err != nil {
				continue
			}
			if *tag.Name == fmt.Sprintf("v%s", schemaVersion) || *tag.Name == schemaVersion {
				return true, nil
			}
		}
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	return false, nil
}
