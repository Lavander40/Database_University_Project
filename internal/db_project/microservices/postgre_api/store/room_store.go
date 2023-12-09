package store

import "db_project/internal/db_project/microservices/postgre_api/model"

type RoomStore struct {
	store *Store
}

func (s *RoomStore) Get(roomId int) (model.Room, error) {
	var room model.Room

	row := s.store.db.QueryRow("SELECT * FROM rooms WHERE id = $1", roomId)
	if err := row.Scan(&room.Id, &room.Number); err != nil {
		return room, err
	}

	return room, nil
}

func (s *RoomStore) GetAll() ([]model.Room, error) {
	var rooms []model.Room

	rows, err := s.store.db.Query("SELECT * FROM rooms")
	if err != nil {
		return rooms, err
	}
	defer rows.Close()

	for rows.Next() {
		room := model.Room{}
		err := rows.Scan(&room.Id, &room.Number)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}