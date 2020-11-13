package _struct

type StatsResponse struct {
	MostVulnerableTargets []struct {
		Address       string   `json:"address"`
		Criticality   int      `json:"criticality"`
		HighVulnCount int      `json:"high_vuln_count"`
		MedVulnCount  int      `json:"med_vuln_count"`
		TargetId      TargetId `json:"target_id"`
	} `json:"most_vulnerable_targets"`
	ScansConductedCount int `json:"scans_conducted_count"`
	ScansRunningCount   int `json:"scans_running_count"`
	ScansWaitingCount   int `json:"scans_waiting_count"`
	TargetsCount        int `json:"targets_count"`
	TopVulnerabilities  []struct {
		Count    int    `json:"count"`
		Name     string `json:"name"`
		Severity int    `json:"severity"`
		VtId     string `json:"vt_id"`
	} `json:"top_vulnerabilities"`
	VulnCount struct {
		High int `json:"high"`
		Med  int `json:"med"`
		Low  int `json:"low"`
	} `json:"vuln_count"`
	VulnCountByCriticality struct {
		Critical int `json:"critical"`
		High     int `json:"high"`
		Low      int `json:"low"`
		Normal   struct {
			High int `json:"high"`
			Med  int `json:"med"`
			Low  int `json:"low"`
		} `json:"normal"`
	} `json:"vuln_count_by_criticality"`
	VulnerabilitiesOpenCount int `json:"vulnerabilities_open_count"`
}
