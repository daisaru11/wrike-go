package wrike

import (
	"fmt"
	"net/url"

	"strings"
)

type GetTasksRequest struct {
	IDs []string
}

type GetTasksResponse struct {
	Kind *string `json:"kind"`
	Data []Task  `json:"data"`
}

type Task struct {
	ID               *string           `json:"id"`
	AccountID        *string           `json:"accountId"`
	Title            *string           `json:"title"`
	Description      *string           `json:"description"`
	BriefDescription *string           `json:"briefDescription"`
	ParentIDs        []string          `json:"parentIds"`
	SuperParentIDs   []string          `json:"superParentIds"`
	SharedIDs        []string          `json:"sharedIds"`
	ResponsibleIDs   []string          `json:"responsibleIds"`
	Status           *string           `json:"status"`
	Importance       *string           `json:"importance"`
	CreatedDate      *string           `json:"createdDate"`
	UpdatedDate      *string           `json:"updatedDate"`
	CompletedDate    *string           `json:"completedDate"`
	Dates            *TaskDates        `json:"dates"`
	Scope            *string           `json:"scope"`
	AuthorIDs        []string          `json:"authorIds"`
	CustomStatusID   *string           `json:"customStatusId"`
	HasAttachments   *bool             `json:"hasAttachments"`
	AttachmentCount  *int              `json:"attachmentCount"`
	Permalink        *string           `json:"permalink"`
	Priority         *string           `json:"priority"`
	FollowedByMe     *bool             `json:"followedByMe"`
	FollowerIDs      []string          `json:"followerIds"`
	Recurrent        *bool             `json:"recurrent"`
	SuperTaskIDs     []string          `json:"superTaskIds"`
	SubTaskIDs       []string          `json:"subTaskIds"`
	DependencyIDs    []string          `json:"dependencyIds"`
	Metadata         []TaskMetadata    `json:"metadata"`
	CustomFields     []TaskCustomField `json:"customFields"`
	EffortAllocation *TaskEffort       `json:"effortAllocation"`
}

type TaskDates struct {
	Type           *string `json:"type"`
	Duration       *int    `json:"duration"`
	Start          *string `json:"start"`
	Due            *string `json:"due"`
	WorkOnWeekends *bool   `json:"workOnWeekends"`
}

type TaskMetadata struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type TaskCustomField struct {
	ID    *string `json:"id"`
	Value *string `json:"value"`
}

type TaskEffort struct {
	Mode            *string `json:"mode"`
	TotalEffort     *int    `json:"totalEffort"`
	AllocatedEffort *int    `json:"allocatedEffort"`
}

func (client *Client) GetTasks(req *GetTasksRequest) (*GetTasksResponse, error) {
	var res GetTasksResponse

	path := fmt.Sprintf("/tasks/%s",
		url.PathEscape(strings.Join(req.IDs, ",")),
	)

	err := client.requestGet(path, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
