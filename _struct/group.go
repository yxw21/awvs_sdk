package _struct

type Group struct {
	GroupId     string    `json:"group_id"`
	Name        string    `json:"name"`
	TargetCount int       `json:"target_count"`
	Description string    `json:"description"`
	VulnCount   VulnCount `json:"vuln_count"`
}

type GroupId string

func (groupId GroupId) String() string{
	return string(groupId)
}

type AddGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddGroupResponse struct {
	GroupId     GroupId `json:"group_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteGroupRequest struct {
	GroupId GroupId `json:"group_id"`
}

type DeleteGroupsRequest struct {
	GroupIdList []GroupId `json:"group_id_list"`
}

type UpdateGroupRequest struct {
	GroupId     string `json:"group_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type QueryGroupRequest struct {
	GroupId string `json:"group_id"`
}

type QueryGroupResponse Group

type QueryGroupsRequest struct {
	L int `json:"l"`
}

type QueryGroupsResponse struct {
	Groups     []Group    `json:"groups"`
	Pagination Pagination `json:"pagination"`
}

type SetTargetsToGroupRequest struct {
	GroupId      string   `json:"group_id"`
	TargetIdList []TargetId `json:"target_id_list"`
}
