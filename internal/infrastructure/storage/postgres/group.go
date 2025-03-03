package postgres

import (
	"database/sql"
	"errors"

	"github.com/sdimitrenco/grammurrr/internal/domains"
)

type GroupRepositoryPostgres struct {
	db *sql.DB
}

func NewGroupRepositoryPostgres(db *sql.DB) *GroupRepositoryPostgres {
	return &GroupRepositoryPostgres{db: db}
}

func (r *GroupRepositoryPostgres) FindAll() ([]domains.WordGroup, error) {
	rows, err := r.db.Query("SELECT id, name FROM word_groups")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []domains.WordGroup
	for rows.Next() {
		var group domains.WordGroup
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (r *GroupRepositoryPostgres) FindById(group domains.WordGroup) (domains.WordGroup, error) {
	var result domains.WordGroup
	err := r.db.QueryRow("SELECT id, name FROM word_groups WHERE id = $1", group.ID).Scan(&result.ID, &result.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, nil
		}
		return result, err
	}
	return result, nil
}

func (r *GroupRepositoryPostgres) FindByGroupName(group domains.WordGroup) (domains.WordGroup, error) {
	var result domains.WordGroup
	err := r.db.QueryRow("SELECT id, name FROM word_groups WHERE name = $1", group.Name).Scan(&result.ID, &result.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, nil
		}
		return result, err
	}
	return result, nil
}

func (r *GroupRepositoryPostgres) Create(group domains.WordGroup) (domains.WordGroup, error) {
	err := r.db.QueryRow("INSERT INTO word_groups (name) VALUES ($1) RETURNING id", group.Name).Scan(&group.ID)
	if err != nil {
		return group, err
	}
	return group, nil
}

func (r *GroupRepositoryPostgres) Update(group domains.WordGroup) (domains.WordGroup, error) {
	_, err := r.db.Exec("UPDATE word_groups SET name = $1 WHERE id = $2", group.Name, group.ID)
	if err != nil {
		return group, err
	}
	return group, nil
}

func (r *GroupRepositoryPostgres) Delete(group domains.WordGroup) error {
	_, err := r.db.Exec("DELETE FROM word_groups WHERE id = $1", group.ID)
	return err
}