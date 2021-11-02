// Code generated by sqlc. DO NOT EDIT.
// source: task.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createTask = `-- name: CreateTask :one
INSERT INTO task (task_group_id, created_at, name, position)
  VALUES($1, $2, $3, $4) RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time
`

type CreateTaskParams struct {
	TaskGroupID uuid.UUID `json:"task_group_id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Position    float64   `json:"position"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.TaskGroupID,
		arg.CreatedAt,
		arg.Name,
		arg.Position,
	)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
		&i.CompletedAt,
		&i.HasTime,
	)
	return i, err
}

const createTaskAll = `-- name: CreateTaskAll :one
INSERT INTO task (task_group_id, created_at, name, position, description, complete, due_date)
  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time
`

type CreateTaskAllParams struct {
	TaskGroupID uuid.UUID      `json:"task_group_id"`
	CreatedAt   time.Time      `json:"created_at"`
	Name        string         `json:"name"`
	Position    float64        `json:"position"`
	Description sql.NullString `json:"description"`
	Complete    bool           `json:"complete"`
	DueDate     sql.NullTime   `json:"due_date"`
}

func (q *Queries) CreateTaskAll(ctx context.Context, arg CreateTaskAllParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTaskAll,
		arg.TaskGroupID,
		arg.CreatedAt,
		arg.Name,
		arg.Position,
		arg.Description,
		arg.Complete,
		arg.DueDate,
	)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
		&i.CompletedAt,
		&i.HasTime,
	)
	return i, err
}

const createTaskComment = `-- name: CreateTaskComment :one
INSERT INTO task_comment (task_id, message, created_at, created_by)
  VALUES ($1, $2, $3, $4) RETURNING task_comment_id, task_id, created_at, updated_at, created_by, pinned, message
`

type CreateTaskCommentParams struct {
	TaskID    uuid.UUID `json:"task_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`
}

func (q *Queries) CreateTaskComment(ctx context.Context, arg CreateTaskCommentParams) (TaskComment, error) {
	row := q.db.QueryRowContext(ctx, createTaskComment,
		arg.TaskID,
		arg.Message,
		arg.CreatedAt,
		arg.CreatedBy,
	)
	var i TaskComment
	err := row.Scan(
		&i.TaskCommentID,
		&i.TaskID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.Pinned,
		&i.Message,
	)
	return i, err
}

const createTaskWatcher = `-- name: CreateTaskWatcher :one
INSERT INTO task_watcher (user_id, task_id, watched_at) VALUES ($1, $2, $3) RETURNING task_watcher_id, task_id, user_id, watched_at
`

type CreateTaskWatcherParams struct {
	UserID    uuid.UUID `json:"user_id"`
	TaskID    uuid.UUID `json:"task_id"`
	WatchedAt time.Time `json:"watched_at"`
}

func (q *Queries) CreateTaskWatcher(ctx context.Context, arg CreateTaskWatcherParams) (TaskWatcher, error) {
	row := q.db.QueryRowContext(ctx, createTaskWatcher, arg.UserID, arg.TaskID, arg.WatchedAt)
	var i TaskWatcher
	err := row.Scan(
		&i.TaskWatcherID,
		&i.TaskID,
		&i.UserID,
		&i.WatchedAt,
	)
	return i, err
}

const deleteTaskByID = `-- name: DeleteTaskByID :exec
DELETE FROM task WHERE task_id = $1
`

func (q *Queries) DeleteTaskByID(ctx context.Context, taskID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTaskByID, taskID)
	return err
}

const deleteTaskCommentByID = `-- name: DeleteTaskCommentByID :one
DELETE FROM task_comment WHERE task_comment_id = $1 RETURNING task_comment_id, task_id, created_at, updated_at, created_by, pinned, message
`

func (q *Queries) DeleteTaskCommentByID(ctx context.Context, taskCommentID uuid.UUID) (TaskComment, error) {
	row := q.db.QueryRowContext(ctx, deleteTaskCommentByID, taskCommentID)
	var i TaskComment
	err := row.Scan(
		&i.TaskCommentID,
		&i.TaskID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.Pinned,
		&i.Message,
	)
	return i, err
}

const deleteTaskWatcher = `-- name: DeleteTaskWatcher :exec
DELETE FROM task_watcher WHERE user_id = $1 AND task_id = $2
`

type DeleteTaskWatcherParams struct {
	UserID uuid.UUID `json:"user_id"`
	TaskID uuid.UUID `json:"task_id"`
}

func (q *Queries) DeleteTaskWatcher(ctx context.Context, arg DeleteTaskWatcherParams) error {
	_, err := q.db.ExecContext(ctx, deleteTaskWatcher, arg.UserID, arg.TaskID)
	return err
}

const deleteTasksByTaskGroupID = `-- name: DeleteTasksByTaskGroupID :execrows
DELETE FROM task where task_group_id = $1
`

func (q *Queries) DeleteTasksByTaskGroupID(ctx context.Context, taskGroupID uuid.UUID) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteTasksByTaskGroupID, taskGroupID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getAllTasks = `-- name: GetAllTasks :many
SELECT task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time FROM task
`

func (q *Queries) GetAllTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.TaskID,
			&i.TaskGroupID,
			&i.CreatedAt,
			&i.Name,
			&i.Position,
			&i.Description,
			&i.DueDate,
			&i.Complete,
			&i.CompletedAt,
			&i.HasTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAssignedTasksDueDateForUserID = `-- name: GetAssignedTasksDueDateForUserID :many
SELECT task.task_id, task.task_group_id, task.created_at, task.name, task.position, task.description, task.due_date, task.complete, task.completed_at, task.has_time FROM task_assigned
  INNER JOIN task ON task.task_id = task_assigned.task_id
  INNER JOIN task_group ON task_group.task_group_id = task.task_group_id
  WHERE user_id = $1
  AND $4::boolean = true OR (
    $4::boolean = false AND complete = $2 AND (
      $2 = false OR ($2 = true AND completed_at > $3)
    )
  )
  ORDER BY task.due_date DESC, task_group.project_id DESC
`

type GetAssignedTasksDueDateForUserIDParams struct {
	UserID      uuid.UUID    `json:"user_id"`
	Complete    bool         `json:"complete"`
	CompletedAt sql.NullTime `json:"completed_at"`
	Column4     bool         `json:"column_4"`
}

func (q *Queries) GetAssignedTasksDueDateForUserID(ctx context.Context, arg GetAssignedTasksDueDateForUserIDParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getAssignedTasksDueDateForUserID,
		arg.UserID,
		arg.Complete,
		arg.CompletedAt,
		arg.Column4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.TaskID,
			&i.TaskGroupID,
			&i.CreatedAt,
			&i.Name,
			&i.Position,
			&i.Description,
			&i.DueDate,
			&i.Complete,
			&i.CompletedAt,
			&i.HasTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAssignedTasksProjectForUserID = `-- name: GetAssignedTasksProjectForUserID :many
SELECT task.task_id, task.task_group_id, task.created_at, task.name, task.position, task.description, task.due_date, task.complete, task.completed_at, task.has_time FROM task_assigned
  INNER JOIN task ON task.task_id = task_assigned.task_id
  INNER JOIN task_group ON task_group.task_group_id = task.task_group_id
  WHERE user_id = $1
  AND $4::boolean = true OR (
    $4::boolean = false AND complete = $2 AND (
      $2 = false OR ($2 = true AND completed_at > $3)
    )
  )
  ORDER BY task_group.project_id DESC, task_assigned.assigned_date DESC
`

type GetAssignedTasksProjectForUserIDParams struct {
	UserID      uuid.UUID    `json:"user_id"`
	Complete    bool         `json:"complete"`
	CompletedAt sql.NullTime `json:"completed_at"`
	Column4     bool         `json:"column_4"`
}

func (q *Queries) GetAssignedTasksProjectForUserID(ctx context.Context, arg GetAssignedTasksProjectForUserIDParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getAssignedTasksProjectForUserID,
		arg.UserID,
		arg.Complete,
		arg.CompletedAt,
		arg.Column4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.TaskID,
			&i.TaskGroupID,
			&i.CreatedAt,
			&i.Name,
			&i.Position,
			&i.Description,
			&i.DueDate,
			&i.Complete,
			&i.CompletedAt,
			&i.HasTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCommentCountForTask = `-- name: GetCommentCountForTask :one
SELECT COUNT(*) FROM task_comment WHERE task_id = $1
`

func (q *Queries) GetCommentCountForTask(ctx context.Context, taskID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCommentCountForTask, taskID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getCommentsForTaskID = `-- name: GetCommentsForTaskID :many
SELECT task_comment_id, task_id, created_at, updated_at, created_by, pinned, message FROM task_comment WHERE task_id = $1 ORDER BY created_at
`

func (q *Queries) GetCommentsForTaskID(ctx context.Context, taskID uuid.UUID) ([]TaskComment, error) {
	rows, err := q.db.QueryContext(ctx, getCommentsForTaskID, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskComment
	for rows.Next() {
		var i TaskComment
		if err := rows.Scan(
			&i.TaskCommentID,
			&i.TaskID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.Pinned,
			&i.Message,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectIDForTask = `-- name: GetProjectIDForTask :one
SELECT project_id FROM task
  INNER JOIN task_group ON task_group.task_group_id = task.task_group_id
  WHERE task_id = $1
`

func (q *Queries) GetProjectIDForTask(ctx context.Context, taskID uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getProjectIDForTask, taskID)
	var project_id uuid.UUID
	err := row.Scan(&project_id)
	return project_id, err
}

const getProjectIdMappings = `-- name: GetProjectIdMappings :many
SELECT project_id, task_id FROM task
INNER JOIN task_group ON task_group.task_group_id = task.task_group_id
  WHERE task_id = ANY($1::uuid[])
`

type GetProjectIdMappingsRow struct {
	ProjectID uuid.UUID `json:"project_id"`
	TaskID    uuid.UUID `json:"task_id"`
}

func (q *Queries) GetProjectIdMappings(ctx context.Context, dollar_1 []uuid.UUID) ([]GetProjectIdMappingsRow, error) {
	rows, err := q.db.QueryContext(ctx, getProjectIdMappings, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProjectIdMappingsRow
	for rows.Next() {
		var i GetProjectIdMappingsRow
		if err := rows.Scan(&i.ProjectID, &i.TaskID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectInfoForTask = `-- name: GetProjectInfoForTask :one
SELECT project.project_id, project.name FROM task
  INNER JOIN task_group ON task_group.task_group_id = task.task_group_id
  INNER JOIN project ON task_group.project_id = project.project_id
  WHERE task_id = $1
`

type GetProjectInfoForTaskRow struct {
	ProjectID uuid.UUID `json:"project_id"`
	Name      string    `json:"name"`
}

func (q *Queries) GetProjectInfoForTask(ctx context.Context, taskID uuid.UUID) (GetProjectInfoForTaskRow, error) {
	row := q.db.QueryRowContext(ctx, getProjectInfoForTask, taskID)
	var i GetProjectInfoForTaskRow
	err := row.Scan(&i.ProjectID, &i.Name)
	return i, err
}

const getRecentlyAssignedTaskForUserID = `-- name: GetRecentlyAssignedTaskForUserID :many
SELECT task.task_id, task.task_group_id, task.created_at, task.name, task.position, task.description, task.due_date, task.complete, task.completed_at, task.has_time FROM task_assigned INNER JOIN
  task ON task.task_id = task_assigned.task_id WHERE user_id = $1
  AND $4::boolean = true OR (
    $4::boolean = false AND complete = $2 AND (
      $2 = false OR ($2 = true AND completed_at > $3)
    )
  )
  ORDER BY task_assigned.assigned_date DESC
`

type GetRecentlyAssignedTaskForUserIDParams struct {
	UserID      uuid.UUID    `json:"user_id"`
	Complete    bool         `json:"complete"`
	CompletedAt sql.NullTime `json:"completed_at"`
	Column4     bool         `json:"column_4"`
}

func (q *Queries) GetRecentlyAssignedTaskForUserID(ctx context.Context, arg GetRecentlyAssignedTaskForUserIDParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getRecentlyAssignedTaskForUserID,
		arg.UserID,
		arg.Complete,
		arg.CompletedAt,
		arg.Column4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.TaskID,
			&i.TaskGroupID,
			&i.CreatedAt,
			&i.Name,
			&i.Position,
			&i.Description,
			&i.DueDate,
			&i.Complete,
			&i.CompletedAt,
			&i.HasTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTaskByID = `-- name: GetTaskByID :one
SELECT task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time FROM task WHERE task_id = $1
`

func (q *Queries) GetTaskByID(ctx context.Context, taskID uuid.UUID) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, taskID)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
		&i.CompletedAt,
		&i.HasTime,
	)
	return i, err
}

const getTaskWatcher = `-- name: GetTaskWatcher :one
SELECT task_watcher_id, task_id, user_id, watched_at FROM task_watcher WHERE user_id = $1 AND task_id = $2
`

type GetTaskWatcherParams struct {
	UserID uuid.UUID `json:"user_id"`
	TaskID uuid.UUID `json:"task_id"`
}

func (q *Queries) GetTaskWatcher(ctx context.Context, arg GetTaskWatcherParams) (TaskWatcher, error) {
	row := q.db.QueryRowContext(ctx, getTaskWatcher, arg.UserID, arg.TaskID)
	var i TaskWatcher
	err := row.Scan(
		&i.TaskWatcherID,
		&i.TaskID,
		&i.UserID,
		&i.WatchedAt,
	)
	return i, err
}

const getTasksForTaskGroupID = `-- name: GetTasksForTaskGroupID :many
SELECT task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time FROM task WHERE task_group_id = $1
`

func (q *Queries) GetTasksForTaskGroupID(ctx context.Context, taskGroupID uuid.UUID) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTasksForTaskGroupID, taskGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.TaskID,
			&i.TaskGroupID,
			&i.CreatedAt,
			&i.Name,
			&i.Position,
			&i.Description,
			&i.DueDate,
			&i.Complete,
			&i.CompletedAt,
			&i.HasTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setTaskComplete = `-- name: SetTaskComplete :one
UPDATE task SET complete = $2, completed_at = $3 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time
`

type SetTaskCompleteParams struct {
	TaskID      uuid.UUID    `json:"task_id"`
	Complete    bool         `json:"complete"`
	CompletedAt sql.NullTime `json:"completed_at"`
}

func (q *Queries) SetTaskComplete(ctx context.Context, arg SetTaskCompleteParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, setTaskComplete, arg.TaskID, arg.Complete, arg.CompletedAt)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
		&i.CompletedAt,
		&i.HasTime,
	)
	return i, err
}

const updateTaskComment = `-- name: UpdateTaskComment :one
UPDATE task_comment SET message = $2, updated_at = $3 WHERE task_comment_id = $1 RETURNING task_comment_id, task_id, created_at, updated_at, created_by, pinned, message
`

type UpdateTaskCommentParams struct {
	TaskCommentID uuid.UUID    `json:"task_comment_id"`
	Message       string       `json:"message"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
}

func (q *Queries) UpdateTaskComment(ctx context.Context, arg UpdateTaskCommentParams) (TaskComment, error) {
	row := q.db.QueryRowContext(ctx, updateTaskComment, arg.TaskCommentID, arg.Message, arg.UpdatedAt)
	var i TaskComment
	err := row.Scan(
		&i.TaskCommentID,
		&i.TaskID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.Pinned,
		&i.Message,
	)
	return i, err
}

const updateTaskDescription = `-- name: UpdateTaskDescription :one
UPDATE task SET description = $2 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time
`

type UpdateTaskDescriptionParams struct {
	TaskID      uuid.UUID      `json:"task_id"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) UpdateTaskDescription(ctx context.Context, arg UpdateTaskDescriptionParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskDescription, arg.TaskID, arg.Description)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
		&i.CompletedAt,
		&i.HasTime,
	)
	return i, err
}

const updateTaskDueDate = `-- name: UpdateTaskDueDate :one
UPDATE task SET due_date = $2, has_time = $3 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time
`

type UpdateTaskDueDateParams struct {
	TaskID  uuid.UUID    `json:"task_id"`
	DueDate sql.NullTime `json:"due_date"`
	HasTime bool         `json:"has_time"`
}

func (q *Queries) UpdateTaskDueDate(ctx context.Context, arg UpdateTaskDueDateParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskDueDate, arg.TaskID, arg.DueDate, arg.HasTime)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
		&i.CompletedAt,
		&i.HasTime,
	)
	return i, err
}

const updateTaskLocation = `-- name: UpdateTaskLocation :one
UPDATE task SET task_group_id = $2, position = $3 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time
`

type UpdateTaskLocationParams struct {
	TaskID      uuid.UUID `json:"task_id"`
	TaskGroupID uuid.UUID `json:"task_group_id"`
	Position    float64   `json:"position"`
}

func (q *Queries) UpdateTaskLocation(ctx context.Context, arg UpdateTaskLocationParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskLocation, arg.TaskID, arg.TaskGroupID, arg.Position)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
		&i.CompletedAt,
		&i.HasTime,
	)
	return i, err
}

const updateTaskName = `-- name: UpdateTaskName :one
UPDATE task SET name = $2 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time
`

type UpdateTaskNameParams struct {
	TaskID uuid.UUID `json:"task_id"`
	Name   string    `json:"name"`
}

func (q *Queries) UpdateTaskName(ctx context.Context, arg UpdateTaskNameParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskName, arg.TaskID, arg.Name)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
		&i.CompletedAt,
		&i.HasTime,
	)
	return i, err
}

const updateTaskPosition = `-- name: UpdateTaskPosition :one
UPDATE task SET position = $2 WHERE task_id = $1 RETURNING task_id, task_group_id, created_at, name, position, description, due_date, complete, completed_at, has_time
`

type UpdateTaskPositionParams struct {
	TaskID   uuid.UUID `json:"task_id"`
	Position float64   `json:"position"`
}

func (q *Queries) UpdateTaskPosition(ctx context.Context, arg UpdateTaskPositionParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskPosition, arg.TaskID, arg.Position)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.TaskGroupID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
		&i.Description,
		&i.DueDate,
		&i.Complete,
		&i.CompletedAt,
		&i.HasTime,
	)
	return i, err
}
