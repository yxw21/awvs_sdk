package _struct

type VulnCount struct {
	Hign   int `json:"hign"`
	Medium int `json:"medium"`
	Low    int `json:"low"`
	Info   int `json:"info"`
}

type Pagination struct {
	Count      int      `json:"count"`
	Cursors    []string `json:"cursors"`
	CursorHash string   `json:"cursor_hash"`
	Sort       string   `json:"sort"`
}

type Error struct {
	Code    int      `json:"code"`
	Reason  string   `json:"reason"`
	Details []string `json:"details"`
}

const (
	Version = "13.0.0"

	HighRiskVulnerabilities           ProfileId = "11111111-1111-1111-1111-111111111112"
	WeakPasswords                     ProfileId = "11111111-1111-1111-1111-111111111115"
	CrawlOnly                         ProfileId = "11111111-1111-1111-1111-111111111117"
	CrossSiteScriptingVulnerabilities ProfileId = "11111111-1111-1111-1111-111111111116"
	SQLInjectionVulnerabilities       ProfileId = "11111111-1111-1111-1111-111111111113"
	FullScan                          ProfileId = "11111111-1111-1111-1111-111111111111"

	Developer                TemplateId = "11111111-1111-1111-1111-111111111111"
	Quick                    TemplateId = "11111111-1111-1111-1111-111111111112"
	ExecutiveSummary         TemplateId = "11111111-1111-1111-1111-111111111113"
	HIPAA                    TemplateId = "11111111-1111-1111-1111-111111111114"
	AffectedItems            TemplateId = "11111111-1111-1111-1111-111111111115"
	ScanComparison           TemplateId = "11111111-1111-1111-1111-111111111124"
	CWE2011                  TemplateId = "11111111-1111-1111-1111-111111111116"
	ISO27001                 TemplateId = "11111111-1111-1111-1111-111111111117"
	NISTSP80053              TemplateId = "11111111-1111-1111-1111-111111111118"
	OWASPTop102013           TemplateId = "11111111-1111-1111-1111-111111111119"
	OWASPTop102017           TemplateId = "11111111-1111-1111-1111-111111111125"
	PCIDSS32                 TemplateId = "11111111-1111-1111-1111-111111111120"
	SarbanesOxley            TemplateId = "11111111-1111-1111-1111-111111111121"
	STIGDISA                 TemplateId = "11111111-1111-1111-1111-111111111122"
	WASCThreatClassification TemplateId = "11111111-1111-1111-1111-111111111123"

	AllVulnerabilities  ListType = "all_vulnerabilities"
	Targets             ListType = "targets"
	Groups              ListType = "groups"
	Scans               ListType = "scans"
	ScanResult          ListType = "scan_result"
	Vulnerabilities     ListType = "vulnerabilities"
	ScanVulnerabilities ListType = "scan_vulnerabilities"
	ScanPair            ListType = "scan_pair"
	ScanResultPair      ListType = "scan_result_pair"

	Scheduled  Status = "scheduled"
	Queued     Status = "queued"
	Starting   Status = "starting"
	Processing Status = "processing"
	Aborting   Status = "aborting"
	Aborted    Status = "aborted"
	Pausing    Status = "pausing"
	Paused     Status = "paused"
	Completed  Status = "completed"
	Failed     Status = "failed"
)
