package vmc

type Task struct {
	UserId                    string `json:"user_id"`
	UserName                  string `json:"user_name"`
	Created                   string `json:"created"`
	Version                   int    `json:"version"`
	Id                        string `json:"id"`
	UpdatedByUserId           string `json:"updated_by_user_id"`
	UpdatedByUserName         string `json:"updated_by_user_name"`
	Updated                   string `json:"updated"`
	CorrelationID             string `json:"correlation_id"`
	CustomerErrorMessage      string `json:"customer_error_message"`
	EndResourceEntityVersion  int    `json:"end_resource_entity_version"`
	EndTime                   string `json:"end_time"`
	ErrorMessage              string `json:"error_message"`
	EstimatedRemainingMinutes int    `json:"estimated_remaining_minutes"`
	LocalizedErrorMessage     string `json:"localized_error_message"`
	OrgID                     string `json:"org_id"`
	OrgType                   string `json:"org_type"`
	Params                    struct {
	} `json:"params"`
	ParentTaskID    string `json:"parent_task_id"`
	PhaseInProgress string `json:"phase_in_progress"`
	ProgressPercent int    `json:"progress_percent"`
	ResourceID      string `json:"resource_id"`
	ResourceType    string `json:"resource_type"`
	ServiceErrors   []struct {
		DefaultMessage           string `json:"default_message"`
		LocalizedMessage         string `json:"localized_message"`
		OriginalService          string `json:"original_service"`
		OriginalServiceErrorCode string `json:"original_service_error_code"`
	} `json:"service_errors"`
	StartResourceEntityVersion int    `json:"start_resource_entity_version"`
	StartTime                  string `json:"start_time"`
	Status                     string `json:"status"`
	SubStatus                  string `json:"sub_status"`
	TaskProgressPhases         []struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		ProgressPercent int    `json:"progress_percent"`
	} `json:"task_progress_phases"`
	TaskType    string `json:"task_type"`
	TaskVersion string `json:"task_version"`
}

type Tasks []Task
