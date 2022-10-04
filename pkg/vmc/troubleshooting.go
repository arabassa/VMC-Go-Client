package vmc

type ConnectivityValidationGroup struct {
	Groups []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		SubGroups []struct {
			Help   string `json:"help"`
			ID     string `json:"id"`
			Inputs []struct {
				ID    string `json:"id"`
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"inputs"`
			Label string `json:"label"`
			Tests []struct {
				Path    string   `json:"path"`
				Pktsize string   `json:"pktsize"`
				Ports   []string `json:"ports"`
				Source  string   `json:"source"`
				Type    string   `json:"type"`
			} `json:"tests"`
		} `json:"sub_groups"`
	} `json:"groups"`
}
