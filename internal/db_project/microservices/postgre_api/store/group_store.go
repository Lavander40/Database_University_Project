package store

import "db_project/internal/db_project/microservices/postgre_api/model"

type GroupStore struct {
	store *Store
}

func (s *GroupStore) Get(groupId int) (model.Group, error) {
	var group model.Group

	row := s.store.db.QueryRow("SELECT * FROM groups WHERE id = $1", groupId)
	if err := row.Scan(&group.Id, &group.DepartId, &group.SpecId); err != nil {
		return group, err
	}

	return group, nil
}

func (s *GroupStore) GetAll() ([]model.Group, error) {
	var groups []model.Group

	rows, err := s.store.db.Query("SELECT * FROM groups")
	if err != nil {
		return groups, err
	}
	defer rows.Close()

	for rows.Next() {
		group := model.Group{}
		err := rows.Scan(&group.Id, &group.DepartId, &group.SpecId)
		if err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}