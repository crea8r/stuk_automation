// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: TaskCreating.sql

package database

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

const createTask = `-- name: CreateTask :one
WITH new_task AS (
  INSERT INTO tasks (
    object_id,
    status,
    input
  )
  SELECT 
    $1,
    'pending',
    $2
  WHERE NOT EXISTS (
    SELECT 1 
    FROM tasks 
    WHERE object_id = $1 
    AND status = 'pending'
  )
  RETURNING id, object_id, status, input, output, error, created_at, started_at, completed_at
)
SELECT id, object_id, status, input, output, error, created_at, started_at, completed_at FROM new_task
`

type CreateTaskParams struct {
	ObjectID *uuid.UUID      `json:"object_id"`
	Input    json.RawMessage `json:"input"`
}

type CreateTaskRow struct {
	ID          *uuid.UUID            `json:"id"`
	ObjectID    *uuid.UUID            `json:"object_id"`
	Status      string                `json:"status"`
	Input       json.RawMessage       `json:"input"`
	Output      pqtype.NullRawMessage `json:"output"`
	Error       sql.NullString        `json:"error"`
	CreatedAt   sql.NullTime          `json:"created_at"`
	StartedAt   sql.NullTime          `json:"started_at"`
	CompletedAt sql.NullTime          `json:"completed_at"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (CreateTaskRow, error) {
	row := q.queryRow(ctx, q.createTaskStmt, createTask, arg.ObjectID, arg.Input)
	var i CreateTaskRow
	err := row.Scan(
		&i.ID,
		&i.ObjectID,
		&i.Status,
		&i.Input,
		&i.Output,
		&i.Error,
		&i.CreatedAt,
		&i.StartedAt,
		&i.CompletedAt,
	)
	return i, err
}