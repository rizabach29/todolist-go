package repositories

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/models/base"
)

type IAttachmentRepository interface {
	Create(m models.CreateAttachmentModel) error
	Update(id int, m models.UpdateAttachmentModel) error
	Delete(id int) error
	GetAll() ([]base.Attachment, error)
	GetById(id int) (base.Attachment, error)
}

type AttachmentRepository struct {
	db *pg.DB
}

func NewAttachmentRepository(db *pg.DB) IAttachmentRepository {
	return &AttachmentRepository{db}
}

func (r *AttachmentRepository) Create(m models.CreateAttachmentModel) error {
	attach := &base.Attachment{
		IdTodolist: m.IdTodolist,
		Url: m.Url,
		Caption: m.Caption,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := r.db.Model(&attach).Insert()
	return err
}

func (r *AttachmentRepository) Update(id int, m models.UpdateAttachmentModel) error {
	values := map[string]interface{}{
		"url": m.Url,
		"caption": m.Caption,
	}

	_, err := r.db.Model(&values).TableExpr("attachment").Update()
	return err
}

func (r *AttachmentRepository) Delete(id int) error {
	var attach base.Attachment
	_, err := r.db.Model(&attach).Where("id = ?", id).Delete()
	return err
}

func (r *AttachmentRepository) GetAll() ([]base.Attachment, error) {
	var attach []base.Attachment
	err := r.db.Model(&attach).Select()
	return attach, err
}

func (r *AttachmentRepository) GetById(id int) (base.Attachment, error) {
	var attach base.Attachment
	err := r.db.Model(&attach).Where("id = ?", id).Select()
	return attach, err
}



