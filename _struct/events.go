package _struct

type ScanSessionId string

func (scanSessionId ScanSessionId) String() string {
	return string(scanSessionId)
}

type QueryEventsRequest struct {
	ResourceId ScanSessionId
}

type QueryEventsResponse struct {
	Notifications []struct {
		Consumed string `json:"consumed"`
		Created  string `json:"created"`
		Data     struct {
			ExtendedStatus string `json:"extended_status"`
			ScanningApp    string `json:"scanning_app"`
			Status         string `json:"status"`
			WorkerId       string `json:"worker_id"`
		} `json:"data"`
		Email          string `json:"email"`
		NotificationId string `json:"notification_id"`
		ResourceId     string `json:"resource_id"`
		ResourceType   int    `json:"resource_type"`
		Severity       int    `json:"severity"`
		TypeId         int    `json:"type_id"`
		UserId         string `json:"user_id"`
	} `json:"notifications"`
	Pagination Pagination `json:"pagination"`
}
