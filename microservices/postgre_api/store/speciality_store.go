package store

import "postgre_api/model"

type SpecialityStore struct {
	store *Store
}

func (s *SpecialityStore) Get(specialityId int) (model.Speciality, error) {
	var speciality model.Speciality

	row := s.store.db.QueryRow("SELECT * FROM specialities WHERE id = $1", specialityId)
	if err := row.Scan(&speciality.Id, &speciality.Name, &speciality.Code, &speciality.DepartId); err != nil {
		return speciality, err
	}

	return speciality, nil
}

func (s *SpecialityStore) GetAll() ([]model.Speciality, error) {
	var specialities []model.Speciality

	rows, err := s.store.db.Query("SELECT * FROM specialities")
	if err != nil {
		return specialities, err
	}
	defer rows.Close()

	for rows.Next() {
		speciality := model.Speciality{}
		err := rows.Scan(&speciality.Id, &speciality.Name, &speciality.Code, &speciality.DepartId)
		if err != nil {
			return specialities, err
		}
		specialities = append(specialities, speciality)
	}

	return specialities, nil
}