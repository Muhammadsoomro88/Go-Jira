package utils

type Jira struct {
	Key    string `json:"key"`
	Fields struct {
		Issuetype struct {
			Name string `json:"name"`
		} `json:"issuetype"`
		Assignee struct {
			DisplayName string `json:"displayName"`
		} `json:"assignee"`
		Updated string `json:"updated"`
		Status  struct {
			Name string `json:"name"`
		} `json:"status"`
		Summary           string  `json:"summary"`
		Customfield_10105 float32 `json:"customfield_10105"`
		Creator           struct {
			DisplayName string `json:"displayName"`
		} `json:"creator"`
		Subtasks []struct {
			Key    string `json:"key"`
			Fields struct {
				Summary string `json:"summary"`
				Status  struct {
					Name string `json:"name"`
				} `json:"status"`
			} `json:"fields"`
		} `json:"subtasks"`
		Parent struct {
			Key    string `json:"key"`
			Fields struct {
				Summary string `json:"summary"`
				Status  struct {
					Name string `json:"name"`
				} `json:"status"`
			} `json:"fields"`
		} `json:"parent"`
		Reporter struct {
			DisplayName string `json:"displayName"`
		} `json:"reporter"`
		Comment struct {
			Total int `json:"total"`
		} `json:"comment"`
		Worklog struct {
			Worklogs []struct {
				TimeSpent string `json:"timeSpent"`
			} `json:"worklogs"`
		} `json:"worklog"`
	} `json:"fields"`
}
