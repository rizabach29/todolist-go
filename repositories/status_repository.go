package repositories

import (
	"github.com/go-pg/pg"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/models/base"
)

type IStatusRepository interface {
	Create(id int, m models.CreateStatusModel) error
	Update(id int, m models.UpdateStatusModel) error
	Delete(id int) error
	GetAll() ([]base.Status, error)
	GetById(id int) (base.Status, error)
}

type StatusRepository struct {
	db *pg.DB
}

func NewStatusRepository(db *pg.DB) IStatusRepository {
	return &StatusRepository{db}
}

func (r *StatusRepository) Create(id int, m models.CreateStatusModel) error {
	status := &base.Status{
		Name: m.Name,
	}

	_, err := r.db.Model(&status).Insert()
	return err
}

func (r *StatusRepository) Update(id int, m models.UpdateStatusModel) error {
	status := new(base.Status)
	_, err := r.db.Model(&status).Set("name = ?", m.Name).Where("id = ?", id).Update()
	return err
}

func (r *StatusRepository) Delete(id int) error {
	model := new(base.Status)
	_, err := r.db.Model(model).Where("id = ?", id).Delete()
	return err
}

func (r *StatusRepository) GetAll() ([]base.Status, error) {
	var status []base.Status

	err := r.db.Model(&status).Select()
	return status, err
}

func (r *StatusRepository) GetById(id int) (base.Status, error) {
	var status base.Status
	
	err := r.db.Model(&status).Where("id = ?", id).Select()
	return status, err
}

