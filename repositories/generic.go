package repositories

import (
	"github.com/template/go-backend-gin-orm/dtos"
	"gorm.io/gorm"
)

type IGenericRepository[T any] interface {
	Create(db *gorm.DB, entity *T) (*T, error)
	GetAll(db *gorm.DB) ([]T, error)
	GetById(db *gorm.DB, id uint) (*T, error)
	Update(db *gorm.DB, entity *T) (*T, error)
	Delete(db *gorm.DB, id uint) (bool, error)
	FindByUniqueField(db *gorm.DB, field string, value interface{}) (*T, error)
	GetPaginated(db *gorm.DB, paginationQuery *dtos.PaginationQuery) (*dtos.ReturningPagination, error)
}

type GenericRepository[T any] struct{}

func NewGenericRepository[T any]() *GenericRepository[T] {
	return &GenericRepository[T]{}
}

func (r *GenericRepository[T]) Create(db *gorm.DB, entity *T) (*T, error) {
	if err := db.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *GenericRepository[T]) GetAll(db *gorm.DB) ([]T, error) {
	var entities []T
	if err := db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *GenericRepository[T]) GetById(db *gorm.DB, id uint) (*T, error) {
	var entity T
	if err := db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GenericRepository[T]) Update(db *gorm.DB, entity *T) (*T, error) {
	if err := db.Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *GenericRepository[T]) Delete(db *gorm.DB, id uint) (bool, error) {
	result := db.Delete(new(T), id)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (r *GenericRepository[T]) FindByUniqueField(db *gorm.DB, field string, value interface{}) (*T, error) {
	var entity T
	if err := db.Where(field+"= ?", value).First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GenericRepository[T]) GetPaginated(db *gorm.DB, paginationQuery *dtos.PaginationQuery) (*dtos.ReturningPagination, error) {
	// Validate page and limit (take)
	page := paginationQuery.Page
	limit := paginationQuery.Take

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	var total int64
	var entities []T

	// Get total count of records
	if err := db.Model(new(T)).Count(&total).Error; err != nil {
		return nil, err
	}

	// Get paginated records
	if err := db.Offset((page - 1) * limit).Limit(limit).Find(&entities).Error; err != nil {
		return nil, err
	}

	// Calculate total pages
	lastPage := int((total + int64(limit) - 1) / int64(limit))

	// Return paginated data
	return &dtos.ReturningPagination{
		Total:       total,
		PerPage:     limit,
		CurrentPage: page,
		TotalPages:  lastPage,
		Data:        entities,
	}, nil
}
