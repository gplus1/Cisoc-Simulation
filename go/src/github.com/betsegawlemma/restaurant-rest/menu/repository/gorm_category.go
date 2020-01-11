package repository

import (
	"github.com/betsegawlemma/restaurant-rest/entity"
	"github.com/betsegawlemma/restaurant-rest/menu"
	"github.com/jinzhu/gorm"
)

// HCategoryGormRepo implements the menu.CategoryRepository interface
type HCategoryGormRepo struct {
	conn *gorm.DB
}

// NewHCategoryGormRepo will create a new object of HCategoryGormRepo
func NewHCategoryGormRepo(db *gorm.DB) menu.CategoryRepository {
	return &HCategoryGormRepo{conn: db}
}

// Categories returns all categories stored in the database
func (cRepo *HCategoryGormRepo) Categories() ([]entity.Category, []error) {
	ctgs := []entity.Category{}
	errs := cRepo.conn.Find(&ctgs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ctgs, errs
}

// Category retrieve a category from the database by its id
func (cRepo *HCategoryGormRepo) Category(id uint) (*entity.Category, []error) {
	ctg := entity.Category{}
	errs := cRepo.conn.First(&ctg, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ctg, errs
}

// UpdateCategory updates a given category in the database
func (cRepo *HCategoryGormRepo) UpdateCategory(category *entity.Category) (*entity.Category, []error) {
	cat := category
	errs := cRepo.conn.Save(cat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

// DeleteCategory deletes a given category from the database
func (cRepo *HCategoryGormRepo) DeleteCategory(id uint) (*entity.Category, []error) {
	cat, errs := cRepo.Category(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = cRepo.conn.Delete(cat, cat.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

// StoreCategory stores a given category in the database
func (cRepo *HCategoryGormRepo) StoreCategory(category *entity.Category) (*entity.Category, []error) {
	cat := category
	errs := cRepo.conn.Create(cat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

// ItemsInCategory retrive from a database a list of food item menus from a given category
func (cRepo *HCategoryGormRepo) ItemsInCategory(category *entity.Category) ([]entity.Item, []error) {
	items := []entity.Item{}
	cat, errs := cRepo.Category(category.ID)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cRepo.conn.Model(cat).Related(&items, "Items").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return items, errs
}
