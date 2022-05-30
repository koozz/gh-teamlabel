package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
)

func main() {
	org := flag.String("org", "", "Organization to use.")
	flag.Parse()

	usageOneLiner := "Usage: gh teamlabel -org=<org> <team_slug:label> [<team_slug:label>]"
	if *org == "" {
		fmt.Printf("Error: Need to set an organization\n%s\n", usageOneLiner)
		os.Exit(1)
	}

	teamLabels, err := parseTeamLabels(flag.Args())
	if err != nil {
		fmt.Printf("%s\n%s\n", err, usageOneLiner)
		os.Exit(1)
	}

	if len(teamLabels) == 0 {
		fmt.Printf("Need at least 1 team label pair, i.e.: -labels=team_slug:label\n%s\n", usageOneLiner)
		os.Exit(1)
	}

	author, err := getAuthor()
	if err != nil {
		log.Fatalf("ERROR: No pull request author found: %s\n", err)
	}

	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatalf("ERROR: Failed to initialize REST client: %s\n", err)
	}

	labelsToAdd := map[string]struct{}{}
	for team, label := range teamLabels {
		if authorInTeam(client, *org, author, team) {
			log.Printf("Author found in team: %s\n", team)
			labelsToAdd[label] = struct{}{}
		} else {
			log.Printf("Author not found in team: %s\n", team)
		}
	}
	for label, _ := range labelsToAdd {
		addTeamLabel(label)
	}
}

func parseTeamLabels(labels []string) (map[string]string, error) {
	teamLabels := make(map[string]string)
	for _, label := range labels {
		teamLabel := strings.SplitN(label, ":", 2)
		if len(teamLabel) < 2 {
			return nil, fmt.Errorf("Label configuration 'team_slug:label' incorrect for: %s", label)
		}
		teamLabels[teamLabel[0]] = teamLabel[1]
	}
	return teamLabels, nil
}

func getAuthor() (string, error) {
	if author, ok := os.LookupEnv("GITHUB_ACTOR"); ok {
		return author, nil
	}
	args := []string{"pr", "view", "--json", "author", "--jq", ".author.login"}
	stdOut, _, err := gh.Exec(args...)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(stdOut.String()), nil
}

func authorInTeam(client api.RESTClient, org, author, team string) bool {
	response := struct{ State string }{}
	uri := fmt.Sprintf("orgs/%s/teams/%s/memberships/%s", org, team, author)
	err := client.Get(uri, &response)
	return err == nil
}

func addTeamLabel(label string) {
	args := []string{"pr", "edit", "--add-label", fmt.Sprintf("%q", label)}
	stdOut, _, err := gh.Exec(args...)
	if err != nil {
		log.Printf("ERROR: Check if repository has label: %s\n", err)
	} else {
		log.Printf("Label '%s' added to: %s\n", label, strings.TrimSpace(stdOut.String()))
	}
}
