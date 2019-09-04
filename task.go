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
	Type           *string `json:"type,omitempty"`
	Duration       *int    `json:"duration,omitempty"`
	Start          *string `json:"start,omitempty"`
	Due            *string `json:"due,omitempty"`
	WorkOnWeekends *bool   `json:"workOnWeekends,omitempty"`
}

type TaskMetadata struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type TaskCustomField struct {
	ID    *string `json:"id,omitempty"`
	Value *string `json:"value,omitempty"`
}

type TaskEffort struct {
	Mode            *string `json:"mode,omitempty"`
	TotalEffort     *int    `json:"totalEffort,omitempty"`
	AllocatedEffort *int    `json:"allocatedEffort,omitempty"`
}

type CreateTaskRequest struct {
	FolderID *string
	Payload  *CreateTaskPayload
}

type CreateTaskResponse struct {
	Kind *string `json:"kind"`
	Data []Task  `json:"data"`
}

type CreateTaskPayload struct {
	Title            *string           `json:"title,omitempty"`
	Description      *string           `json:"description,omitempty"`
	Status           *string           `json:"status,omitempty"`
	Importance       *string           `json:"importance,omitempty"`
	Dates            *TaskDates        `json:"dates,omitempty"`
	Shareds          []string          `json:"shareds,omitempty"`
	Parents          []string          `json:"parents,omitempty"`
	Responsibles     []string          `json:"responsibles,omitempty"`
	Followers        []string          `json:"followers,omitempty"`
	Follow           *bool             `json:"follow,omitempty"`
	PriorityBefore   *string           `json:"priorityBefore,omitempty"`
	PriorityAfter    *string           `json:"priorityAfter,omitempty"`
	SuperTasks       []string          `json:"superTasks,omitempty"`
	Metadata         []TaskMetadata    `json:"metadata,omitempty"`
	CustomFields     []TaskCustomField `json:"customFields,omitempty"`
	CustomStatus     *string           `json:"customStatus,omitempty"`
	EffortAllocation *TaskEffort       `json:"effortAllocation,omitempty"`
}

type UpdateTaskRequest struct {
	TaskID  *string
	Payload *UpdateTaskPayload
}

type UpdateTaskResponse struct {
	Kind *string `json:"kind"`
	Data []Task  `json:"data"`
}

type UpdateTaskPayload struct {
	Title              *string           `json:"title,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Status             *string           `json:"status,omitempty"`
	Importance         *string           `json:"importance,omitempty"`
	Dates              *TaskDates        `json:"dates,omitempty"`
	AddParents         []string          `json:"addParents,omitempty"`
	RemoveParents      []string          `json:"removeParents,omitempty"`
	AddShareds         []string          `json:"addShareds,omitempty"`
	RemoveShareds      []string          `json:"removeShareds,omitempty"`
	AddResponsibles    []string          `json:"addResponsibles,omitempty"`
	RemoveResponsibles []string          `json:"removeResponsibles,omitempty"`
	AddFollowers       []string          `json:"addFollowers,omitempty"`
	RemoveFollowers    []string          `json:"removeFollowers,omitempty"`
	Follow             *bool             `json:"follow,omitempty"`
	PriorityBefore     *string           `json:"priorityBefore,omitempty"`
	PriorityAfter      *string           `json:"priorityAfter,omitempty"`
	AddSuperTasks      []string          `json:"addSuperTasks,omitempty"`
	RemoveSuperTasks   []string          `json:"removeSuperTasks,omitempty"`
	Metadata           []TaskMetadata    `json:"metadata,omitempty"`
	CustomFields       []TaskCustomField `json:"customFields,omitempty"`
	CustomStatus       *string           `json:"customStatus,omitempty"`
	Restore            *bool             `json:"restore,omitempty"`
	EffortAllocation   *TaskEffort       `json:"effortAllocation,omitempty"`
}

type DeleteTaskRequest struct {
	TaskID *string
}

type DeleteTaskResponse struct {
	Kind *string `json:"kind"`
	Data []Task  `json:"data"`
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

func (client *Client) CreateTask(req *CreateTaskRequest) (*CreateTaskResponse, error) {
	var res CreateTaskResponse

	if req.FolderID == nil {
		return nil, fmt.Errorf("FolderID is required to request the CreateTask API")
	}
	folderID := StringValue(req.FolderID)

	if req.Payload == nil {
		return nil, fmt.Errorf("Payload is required to request the CreateTask API")
	}
	payload := *req.Payload

	path := fmt.Sprintf("/folders/%s/tasks",
		url.PathEscape(folderID),
	)

	err := client.requestPost(path, &payload, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (client *Client) UpdateTask(req *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	var res UpdateTaskResponse

	if req.TaskID == nil {
		return nil, fmt.Errorf("TaskID is required to request the UpdateTask API")
	}
	taskID := StringValue(req.TaskID)

	if req.Payload == nil {
		return nil, fmt.Errorf("Payload is required to request the UpdateTask API")
	}
	payload := *req.Payload

	path := fmt.Sprintf("/tasks/%s",
		url.PathEscape(taskID),
	)

	err := client.requestPut(path, &payload, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (client *Client) DeleteTask(req *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	var res DeleteTaskResponse

	if req.TaskID == nil {
		return nil, fmt.Errorf("TaskID is required to request the DeleteTask API")
	}
	taskID := StringValue(req.TaskID)

	path := fmt.Sprintf("/tasks/%s",
		url.PathEscape(taskID),
	)

	err := client.requestDelete(path, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
