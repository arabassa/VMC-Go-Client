package vmc

type SddcInput struct {
	SddcList []SddcList `json:"inputs"`
}

type SddcList struct {
	Choice      []SddcChoice `json:"choice"`
	Label       string       `json:"label"`
	Name        string       `json:"name"`
	Placeholder string       `json:"placeholder"`
	Required    bool         `json:"required"`
}

type SddcChoice struct {
	DisplayText string `json:"displayText"`
	Value       string `json:"value"`
}
