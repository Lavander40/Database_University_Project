package store

import "db_project/internal/db_project/microservices/postgre_api/model"

type EquipmentStore struct {
	store *Store
}

func (s *EquipmentStore) Get(equipmentId int) (model.Equipment, error) {
	var equipment model.Equipment

	row := s.store.db.QueryRow("SELECT * FROM equipments WHERE id = $1", equipmentId)
	if err := row.Scan(&equipment.Id, &equipment.Name); err != nil {
		return equipment, err
	}

	return equipment, nil
}

func (s *EquipmentStore) GetAll() ([]model.Equipment, error) {
	var equipments []model.Equipment

	rows, err := s.store.db.Query("SELECT * FROM equipments")
	if err != nil {
		return equipments, err
	}
	defer rows.Close()

	for rows.Next() {
		equipment := model.Equipment{}
		err := rows.Scan(&equipment.Id, &equipment.Name)
		if err != nil {
			return equipments, err
		}
		equipments = append(equipments, equipment)
	}

	return equipments, nil
}