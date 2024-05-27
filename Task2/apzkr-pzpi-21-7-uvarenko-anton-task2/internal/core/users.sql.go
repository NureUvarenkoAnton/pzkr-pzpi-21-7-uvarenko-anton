// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package core

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (name, email, password, user_type) VALUES(?, ?, ?, ?)
`

type CreateUserParams struct {
	Name     sql.NullString
	Email    sql.NullString
	Password sql.NullString
	UserType NullUsersUserType
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.UserType,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE users
SET is_deleted = true
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, name, email, password, user_type, is_banned, is_deleted FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.UserType,
			&i.IsBanned,
			&i.IsDeleted,
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

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, user_type, is_banned, is_deleted FROM users
WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.UserType,
		&i.IsBanned,
		&i.IsDeleted,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, email, password, user_type, is_banned, is_deleted FROM users
WHERE id = ?
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.UserType,
		&i.IsBanned,
		&i.IsDeleted,
	)
	return i, err
}

const getUserByUserType = `-- name: GetUserByUserType :many
SELECT id, name, email, password, user_type, is_banned, is_deleted FROM users
WHERE user_type = ?
`

func (q *Queries) GetUserByUserType(ctx context.Context, userType NullUsersUserType) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUserByUserType, userType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.UserType,
			&i.IsBanned,
			&i.IsDeleted,
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

const getUsers = `-- name: GetUsers :many
SELECT id, name, email, password, user_type, is_banned, is_deleted FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.UserType,
			&i.IsBanned,
			&i.IsDeleted,
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

const setBanState = `-- name: SetBanState :exec
UPDATE users
SET is_banned = ?
WHERE id = ?
`

type SetBanStateParams struct {
	IsBanned sql.NullBool
	ID       int64
}

func (q *Queries) SetBanState(ctx context.Context, arg SetBanStateParams) error {
	_, err := q.db.ExecContext(ctx, setBanState, arg.IsBanned, arg.ID)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET 
  name = ?,
  email = ?
WHERE
  id = ?
`

type UpdateUserParams struct {
	Name  sql.NullString
	Email sql.NullString
	ID    int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.Name, arg.Email, arg.ID)
	return err
}
