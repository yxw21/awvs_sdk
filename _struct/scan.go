package _struct

type ScanId string

func (scanId ScanId) String() string {
	return string(scanId)
}

type ProfileId string

func (profileId ProfileId) String() string {
	return string(profileId)
}

type Status string

func (status Status) String() string {
	return string(status)
}

type Schedule struct {
	Disable       bool    `json:"disable"`
	StartDate     *string `json:"start_date"`
	TimeSensitive bool    `json:"time_sensitive"`
	Triggerable   bool    `json:"triggerable"`
}

type AddScanRequest struct {
	Incremental bool      `json:"incremental"`
	ProfileId   ProfileId `json:"profile_id"`
	Schedule    Schedule  `json:"schedule"`
	TargetId    TargetId  `json:"target_id"`
}

type AddScanResponse struct {
	Incremental bool      `json:"incremental"`
	MaxScanTime int       `json:"max_scan_time"`
	ProfileId   ProfileId `json:"profile_id"`
	Schedule    Schedule  `json:"schedule"`
	TargetId    TargetId  `json:"target_id"`
}

type AbortScanRequest struct {
	ScanId ScanId `json:"scan_id"`
}

type ResumeScanRequest struct {
	ScanId ScanId `json:"scan_id"`
}

type QueryScanRequest struct {
	ScanId ScanId `json:"scan_id"`
}

type QueryScanResponse struct {
	Criticality    int `json:"criticality"`
	CurrentSession struct {
		EventLevel     int       `json:"event_level"`
		Progress       int       `json:"progress"`
		ScanSessionId  string    `json:"scan_session_id"`
		SeverityCounts VulnCount `json:"severity_counts"`
		StartDate      string    `json:"start_date"`
		Status         string    `json:"status"`
		Threat         int       `json:"threat"`
	} `json:"current_session"`
	Incremental        bool      `json:"incremental"`
	ManualIntervention bool      `json:"manual_intervention"`
	MaxScanTime        int       `json:"max_scan_time"`
	NextRun            string    `json:"next_run"`
	ProfileId          ProfileId `json:"profile_id"`
	ProfileName        string    `json:"profile_name"`
	ReportTemplateId   string    `json:"report_template_id"`
	ScanId             ScanId    `json:"scan_id"`
	Schedule           struct {
		Disable       bool   `json:"disable"`
		HistoryLimit  int    `json:"history_limit"`
		Recurrence    string `json:"recurrence"`
		StartDate     string `json:"start_date"`
		TimeSensitive bool   `json:"time_sensitive"`
		Triggerable   bool   `json:"triggerable"`
	} `json:"schedule"`
	Target struct {
		Address     string `json:"address"`
		Criticality int    `json:"criticality"`
		Description string `json:"description"`
		Type        string `json:"type"`
	} `json:"target"`
	TargetId TargetId `json:"target_id"`
}
