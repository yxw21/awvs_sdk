package _struct

type TemplateId string

func (templateId TemplateId) String() string {
	return string(templateId)
}

type ListType string

func (listType ListType) String() string {
	return string(listType)
}

type Source struct {
	ListType    ListType `json:"list_type"`
	IdList      []ScanId `json:"id_list"`
	Description string   `json:"description"`
}

type Report struct {
	Download       []string   `json:"download"`
	GenerationDate string     `json:"generation_date"`
	ReportId       string     `json:"report_id"`
	Source         Source     `json:"source"`
	Status         string     `json:"status"`
	TemplateId     TemplateId `json:"template_id"`
	TemplateName   string     `json:"template_name"`
	TemplateType   int        `json:"template_type"`
}

type AddReportRequest struct {
	Source     Source     `json:"source"`
	TemplateId TemplateId `json:"template_id"`
}

type QueryReportRequest struct {
	ReportId string `json:"report_id"`
}

type QueryReportResponse struct {
	ReportId       string   `json:"report_id"`
	Source         Source   `json:"source"`
	TemplateId     string   `json:"template_id"`
	TemplateName   string   `json:"template_name"`
	TemplateType   int      `json:"template_type"`
	GenerationDate string   `json:"generation_date"`
	Status         string   `json:"status"`
	Download       []string `json:"download"`
}

type QueryReportsRequest struct {
	L int `json:"l"`
}

type QueryReportsResponse struct {
	Pagination Pagination `json:"pagination"`
	Reports    []Report   `json:"reports"`
}
