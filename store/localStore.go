package store

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Muhammadsoomro88/go-jira/utils"
)

func StoreJiraData(story string, d utils.Jira) {
	output := fmt.Sprintf("Story number: %s\nIssue Type: %s\nAssignee: %s\nStatus: %s\nDescription: %s\nStory Points: %.2f\nCreator: %s",
		d.Key, d.Fields.Issuetype.Name, d.Fields.Assignee, d.Fields.Status.Name,
		d.Fields.Summary, d.Fields.Customfield_10105, d.Fields.Creator.DisplayName)

	fileName := story + ".txt"

	// adding data into file
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := ioutil.WriteFile(fileName, []byte(output), 0644); err != nil {
		log.Fatal(err)
	}
}
