package wrike

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestGetTasksResponseJsonUnmarshal(t *testing.T) {
	body := []byte(`
	{
		"kind": "tasks",
		"data":
		[
				{
						"id": "IEABRDPJKQAQOJ3P",
						"accountId": "IEABRDPJ",
						"title": "Test task",
						"description": "",
						"briefDescription": "",
						"parentIds":
						[
								"IEABRDPJI4AQOJ3J"
						],
						"superParentIds":
						[
						],
						"sharedIds":
						[
								"KUGEYQ5H"
						],
						"responsibleIds":
						[
						],
						"status": "Active",
						"importance": "Normal",
						"createdDate": "2019-07-31T15:40:45Z",
						"updatedDate": "2019-07-31T15:40:54Z",
						"dates":
						{
								"type": "Planned",
								"duration": 2880,
								"start": "2019-07-29T09:00:00",
								"due": "2019-08-05T17:00:00"
						},
						"scope": "WsTask",
						"authorIds":
						[
								"KUGEYQ5H"
						],
						"customStatusId": "IEABRDPJJMAAAAAA",
						"hasAttachments": false,
						"permalink": "https://www.wrike.com/open.htm?id=17246063",
						"priority": "029378008000000000005c00",
						"followedByMe": true,
						"followerIds":
						[
								"KUGEYQ5H"
						],
						"superTaskIds":
						[
						],
						"subTaskIds":
						[
								"IEABRDPJKQAQOJ3R"
						],
						"dependencyIds":
						[
								"IEABRDPJIUAQOJ3PKMAQOJ3R",
								"IEABRDPJIUAQOJ3TKMAQOJ3P"
						],
						"metadata":
						[
								{
										"key": "testMetaKey",
										"value": "testMetaValue"
								}
						],
						"customFields":
						[
								{
										"id": "IEABRDPJJUAAMVQM",
										"value": "testValue"
								},
								{
										"id": "IEABRDPJJUAAMVQN",
										"value": "testValue"
								}
						]
				},
				{
						"id": "IEABRDPJKQAQOJ3T",
						"accountId": "IEABRDPJ",
						"title": "New title",
						"description": "New description",
						"briefDescription": "New description",
						"parentIds":
						[
								"IEABRDPJI4AQOJ3I",
								"IEABRDPJI4AQOJ3J"
						],
						"superParentIds":
						[
								"IEABRDPJI4AQOJ3K"
						],
						"sharedIds":
						[
								"KUGEYQ5H"
						],
						"responsibleIds":
						[
								"KUGEYQ5H"
						],
						"status": "Deferred",
						"importance": "Low",
						"createdDate": "2019-07-31T15:40:46Z",
						"updatedDate": "2019-07-31T15:40:54Z",
						"dates":
						{
								"type": "Planned",
								"duration": 1920,
								"start": "2019-07-31T09:00:00",
								"due": "2019-08-05T17:00:00"
						},
						"scope": "WsTask",
						"authorIds":
						[
								"KUGEYQ5H"
						],
						"customStatusId": "IEABRDPJJMAAAAAC",
						"hasAttachments": true,
						"permalink": "https://www.wrike.com/open.htm?id=17246067",
						"priority": "029378008000000000005a00",
						"followedByMe": true,
						"followerIds":
						[
								"KUGEYQ5H"
						],
						"superTaskIds":
						[
								"IEABRDPJKQAQOJ3R"
						],
						"subTaskIds":
						[
						],
						"dependencyIds":
						[
								"IEABRDPJIUAQOJ3TKMAQOJ3P"
						],
						"metadata":
						[
								{
										"key": "testMetaKey",
										"value": "testMetaValue"
								}
						],
						"customFields":
						[
								{
										"id": "IEABRDPJJUAAMVQN",
										"value": "testValue"
								},
								{
										"id": "IEABRDPJJUAAMVQM",
										"value": "testValue"
								}
						]
				}
		]
 }
	`)

	var res GetTasksResponse
	err := json.Unmarshal(body, &res)
	if assert.NoError(t, err) {
		assert.Equal(t, "tasks", *res.Kind)

		assert.Equal(t, "IEABRDPJKQAQOJ3P", *res.Data[0].ID)
		assert.Equal(t, "IEABRDPJ", *res.Data[0].AccountID)
		assert.Equal(t, "Test task", *res.Data[0].Title)
		assert.Equal(t, "", *res.Data[0].Description)
		assert.Equal(t, []string{"IEABRDPJI4AQOJ3J"}, res.Data[0].ParentIDs)
		assert.Equal(t, []string{}, res.Data[0].SuperParentIDs)
		assert.Equal(t, []string{"KUGEYQ5H"}, res.Data[0].SharedIDs)
		assert.Equal(t, []string{}, res.Data[0].ResponsibleIDs)
		assert.Equal(t, "Active", *res.Data[0].Status)
		assert.Equal(t, "Normal", *res.Data[0].Importance)
		assert.Equal(t, "2019-07-31T15:40:45Z", *res.Data[0].CreatedDate)
		assert.Equal(t, "2019-07-31T15:40:54Z", *res.Data[0].UpdatedDate)
		assert.Equal(t, TaskDates{
			Type:     String("Planned"),
			Duration: Int(2880),
			Start:    String("2019-07-29T09:00:00"),
			Due:      String("2019-08-05T17:00:00"),
		}, *res.Data[0].Dates)
		assert.Equal(t, "WsTask", *res.Data[0].Scope)
		assert.Equal(t, []string{"KUGEYQ5H"}, res.Data[0].AuthorIDs)
		assert.Equal(t, "IEABRDPJJMAAAAAA", *res.Data[0].CustomStatusID)
		assert.Equal(t, false, *res.Data[0].HasAttachments)
		assert.Equal(t, "https://www.wrike.com/open.htm?id=17246063", *res.Data[0].Permalink)
		assert.Equal(t, "029378008000000000005c00", *res.Data[0].Priority)
		assert.Equal(t, true, *res.Data[0].FollowedByMe)
		assert.Equal(t, []string{"KUGEYQ5H"}, res.Data[0].FollowerIDs)
		assert.Equal(t, []string{}, res.Data[0].SuperTaskIDs)
		assert.Equal(t, []string{"IEABRDPJKQAQOJ3R"}, res.Data[0].SubTaskIDs)
		assert.Equal(t, []string{"IEABRDPJIUAQOJ3PKMAQOJ3R", "IEABRDPJIUAQOJ3TKMAQOJ3P"}, res.Data[0].DependencyIDs)
		assert.Equal(t, []TaskMetadata{
			TaskMetadata{
				Key:   String("testMetaKey"),
				Value: String("testMetaValue"),
			},
		}, res.Data[0].Metadata)
		assert.Equal(t, []TaskCustomField{
			TaskCustomField{
				ID:    String("IEABRDPJJUAAMVQM"),
				Value: String("testValue"),
			},
			TaskCustomField{
				ID:    String("IEABRDPJJUAAMVQN"),
				Value: String("testValue"),
			},
		}, res.Data[0].CustomFields)

		assert.Equal(t, "IEABRDPJKQAQOJ3T", *res.Data[1].ID)
		assert.Equal(t, "IEABRDPJ", *res.Data[1].AccountID)
		assert.Equal(t, "New title", *res.Data[1].Title)
		assert.Equal(t, "New description", *res.Data[1].Description)
		assert.Equal(t, "New description", *res.Data[1].BriefDescription)
		assert.Equal(t, []string{"IEABRDPJI4AQOJ3I", "IEABRDPJI4AQOJ3J"}, res.Data[1].ParentIDs)
		assert.Equal(t, []string{"IEABRDPJI4AQOJ3K"}, res.Data[1].SuperParentIDs)
		assert.Equal(t, []string{"KUGEYQ5H"}, res.Data[1].SharedIDs)
		assert.Equal(t, []string{"KUGEYQ5H"}, res.Data[1].ResponsibleIDs)
		assert.Equal(t, "Deferred", *res.Data[1].Status)
		assert.Equal(t, "Low", *res.Data[1].Importance)
		assert.Equal(t, "2019-07-31T15:40:46Z", *res.Data[1].CreatedDate)
		assert.Equal(t, "2019-07-31T15:40:54Z", *res.Data[1].UpdatedDate)
		assert.Equal(t, TaskDates{
			Type:     String("Planned"),
			Duration: Int(1920),
			Start:    String("2019-07-31T09:00:00"),
			Due:      String("2019-08-05T17:00:00"),
		}, *res.Data[1].Dates)
		assert.Equal(t, "WsTask", *res.Data[1].Scope)
		assert.Equal(t, []string{"KUGEYQ5H"}, res.Data[1].AuthorIDs)
		assert.Equal(t, "IEABRDPJJMAAAAAC", *res.Data[1].CustomStatusID)
		assert.Equal(t, true, *res.Data[1].HasAttachments)
		assert.Equal(t, "https://www.wrike.com/open.htm?id=17246067", *res.Data[1].Permalink)
		assert.Equal(t, "029378008000000000005a00", *res.Data[1].Priority)
		assert.Equal(t, true, *res.Data[1].FollowedByMe)
		assert.Equal(t, []string{"KUGEYQ5H"}, res.Data[1].FollowerIDs)
		assert.Equal(t, []string{"IEABRDPJKQAQOJ3R"}, res.Data[1].SuperTaskIDs)
		assert.Equal(t, []string{}, res.Data[1].SubTaskIDs)
		assert.Equal(t, []string{"IEABRDPJIUAQOJ3TKMAQOJ3P"}, res.Data[1].DependencyIDs)
		assert.Equal(t, []TaskMetadata{
			TaskMetadata{
				Key:   String("testMetaKey"),
				Value: String("testMetaValue"),
			},
		}, res.Data[1].Metadata)
		assert.Equal(t, []TaskCustomField{
			TaskCustomField{
				ID:    String("IEABRDPJJUAAMVQN"),
				Value: String("testValue"),
			},
			TaskCustomField{
				ID:    String("IEABRDPJJUAAMVQM"),
				Value: String("testValue"),
			},
		}, res.Data[1].CustomFields)
	}
}
