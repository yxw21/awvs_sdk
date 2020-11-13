package _struct

type Target struct {
	Address     string `json:"address"`
	Description string `json:"description"`
}

type TargetId string

func (targetId TargetId) String() string {
	return string(targetId)
}

type AddTargetsRequest struct {
	Targets []Target  `json:"targets"`
	Groups  []GroupId `json:"groups"`
}

type AddTargetsResponse struct {
	Targets []struct {
		Address              string   `json:"address"`
		CanonicalAddress     string   `json:"canonical_address"`
		CanonicalAddressHash string   `json:"canonical_address_hash"`
		Criticality          int      `json:"criticality"`
		Description          string   `json:"description"`
		Domain               string   `json:"domain"`
		TargetId             TargetId `json:"target_id"`
		TargetType           string   `json:"target_type"`
		Type                 string   `json:"type"`
	} `json:"targets"`
}

type DeleteTargetsRequest struct {
	TargetIdList []TargetId `json:"target_id_list"`
}

type QueryTargetRequest struct {
	TargetId TargetId `json:"target_id"`
}

type QueryTargetResponse struct {
	Address               string    `json:"address"`
	ContinuousMode        bool      `json:"continuous_mode"`
	Criticality           int       `json:"criticality"`
	DeletedAt             string    `json:"deleted_at"`
	Description           string    `json:"description"`
	Fqdn                  string    `json:"fqdn"`
	FqdnHash              string    `json:"fqdn_hash"`
	LastScanDate          string    `json:"last_scan_date"`
	LastScanId            ScanId    `json:"last_scan_id"`
	LastScanSessionId     string    `json:"last_scan_session_id"`
	LastScanSessionStatus string    `json:"last_scan_session_status"`
	Locked                bool      `json:"locked"`
	ManualIntervention    bool      `json:"manual_intervention"`
	SeverityCounts        VulnCount `json:"severity_counts"`
	TargetId              TargetId  `json:"target_id"`
	Threat                int       `json:"threat"`
	Type                  string    `json:"type"`
	Verification          string    `json:"verification"`
}
