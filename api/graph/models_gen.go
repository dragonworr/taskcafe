// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"time"

	"github.com/google/uuid"
	"github.com/jordanknott/project-citadel/api/pg"
)

type AddTaskLabelInput struct {
	TaskID         uuid.UUID `json:"taskID"`
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
}

type AssignTaskInput struct {
	TaskID uuid.UUID `json:"taskID"`
	UserID uuid.UUID `json:"userID"`
}

type ChecklistBadge struct {
	Complete int `json:"complete"`
	Total    int `json:"total"`
}

type CreateTaskChecklist struct {
	TaskID   uuid.UUID `json:"taskID"`
	Name     string    `json:"name"`
	Position float64   `json:"position"`
}

type CreateTaskChecklistItem struct {
	TaskChecklistID uuid.UUID `json:"taskChecklistID"`
	Name            string    `json:"name"`
	Position        float64   `json:"position"`
}

type DeleteProjectLabel struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
}

type DeleteTaskChecklistItem struct {
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
}

type DeleteTaskChecklistItemPayload struct {
	Ok                bool                  `json:"ok"`
	TaskChecklistItem *pg.TaskChecklistItem `json:"taskChecklistItem"`
}

type DeleteTaskGroupInput struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
}

type DeleteTaskGroupPayload struct {
	Ok           bool          `json:"ok"`
	AffectedRows int           `json:"affectedRows"`
	TaskGroup    *pg.TaskGroup `json:"taskGroup"`
}

type DeleteTaskInput struct {
	TaskID string `json:"taskID"`
}

type DeleteTaskPayload struct {
	TaskID string `json:"taskID"`
}

type FindProject struct {
	ProjectID string `json:"projectId"`
}

type FindTask struct {
	TaskID uuid.UUID `json:"taskID"`
}

type FindUser struct {
	UserID string `json:"userId"`
}

type LogoutUser struct {
	UserID string `json:"userID"`
}

type NewProject struct {
	UserID uuid.UUID `json:"userID"`
	TeamID uuid.UUID `json:"teamID"`
	Name   string    `json:"name"`
}

type NewProjectLabel struct {
	ProjectID    uuid.UUID `json:"projectID"`
	LabelColorID uuid.UUID `json:"labelColorID"`
	Name         *string   `json:"name"`
}

type NewRefreshToken struct {
	UserID string `json:"userId"`
}

type NewTask struct {
	TaskGroupID string  `json:"taskGroupID"`
	Name        string  `json:"name"`
	Position    float64 `json:"position"`
}

type NewTaskGroup struct {
	ProjectID string  `json:"projectID"`
	Name      string  `json:"name"`
	Position  float64 `json:"position"`
}

type NewTaskGroupLocation struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Position    float64   `json:"position"`
}

type NewTaskLocation struct {
	TaskID      uuid.UUID `json:"taskID"`
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Position    float64   `json:"position"`
}

type NewTeam struct {
	Name           string `json:"name"`
	OrganizationID string `json:"organizationID"`
}

type NewUserAccount struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Initials string `json:"initials"`
	Password string `json:"password"`
}

type ProfileIcon struct {
	URL      *string `json:"url"`
	Initials *string `json:"initials"`
	BgColor  *string `json:"bgColor"`
}

type ProjectMember struct {
	ID          uuid.UUID    `json:"id"`
	FullName    string       `json:"fullName"`
	ProfileIcon *ProfileIcon `json:"profileIcon"`
}

type ProjectsFilter struct {
	TeamID *string `json:"teamID"`
}

type RemoveTaskLabelInput struct {
	TaskLabelID uuid.UUID `json:"taskLabelID"`
}

type SetTaskChecklistItemComplete struct {
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
	Complete            bool      `json:"complete"`
}

type SetTaskComplete struct {
	TaskID   uuid.UUID `json:"taskID"`
	Complete bool      `json:"complete"`
}

type TaskBadges struct {
	Checklist *ChecklistBadge `json:"checklist"`
}

type ToggleTaskLabelInput struct {
	TaskID         uuid.UUID `json:"taskID"`
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
}

type ToggleTaskLabelPayload struct {
	Active bool     `json:"active"`
	Task   *pg.Task `json:"task"`
}

type UnassignTaskInput struct {
	TaskID uuid.UUID `json:"taskID"`
	UserID uuid.UUID `json:"userID"`
}

type UpdateProjectLabel struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
	LabelColorID   uuid.UUID `json:"labelColorID"`
	Name           string    `json:"name"`
}

type UpdateProjectLabelColor struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
	LabelColorID   uuid.UUID `json:"labelColorID"`
}

type UpdateProjectLabelName struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
	Name           string    `json:"name"`
}

type UpdateProjectName struct {
	ProjectID uuid.UUID `json:"projectID"`
	Name      string    `json:"name"`
}

type UpdateTaskChecklistItemName struct {
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
	Name                string    `json:"name"`
}

type UpdateTaskDescriptionInput struct {
	TaskID      uuid.UUID `json:"taskID"`
	Description string    `json:"description"`
}

type UpdateTaskDueDate struct {
	TaskID  uuid.UUID  `json:"taskID"`
	DueDate *time.Time `json:"dueDate"`
}

type UpdateTaskGroupName struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Name        string    `json:"name"`
}

type UpdateTaskLocationPayload struct {
	PreviousTaskGroupID uuid.UUID `json:"previousTaskGroupID"`
	Task                *pg.Task  `json:"task"`
}

type UpdateTaskName struct {
	TaskID string `json:"taskID"`
	Name   string `json:"name"`
}
