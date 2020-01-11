package repository

import (
	"github.com/HouseRentalSystem/Back_End/entity"
	"github.com/HouseRentalSystem/Back_End/entity/menu"
	"github.com/jinzhu/gorm"
)

// HouseTypeGormRepo implements the menu.HouseTypeRepository interface
type HouseTypeGormRepo struct {
	conn *gorm.DB
}

// NewHouseTypeGormRepo will create a new object of HouseTypeGormRepo
func NewHouseTypeGormRepo(db *gorm.DB) menu.HouseTypeRepository {
	return &HouseTypeGormRepo{conn: db}
}

// HouseTypes returns all food menus stored in the database
func (HouseTypeRepo *HouseTypeGormRepo) HouseTypes() ([]entity.HouseType, []error) {
	HouseTypes := []entity.HouseType{}
	errs := HouseTypeRepo.conn.Find(&HouseTypes).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return HouseTypes, errs
}

// HouseType retrieves a food menu by its id from the database
func (HouseTypeRepo *HouseTypeGormRepo) HouseType(id uint) (*entity.HouseType, []error) {
	HouseType := entity.HouseType{}
	errs := HouseTypeRepo.conn.First(&HouseType, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &HouseType, errs
}

// UpdateHouseType updates a given food menu HouseType in the database
func (HouseTypeRepo *HouseTypeGormRepo) UpdateHouseType(HouseType *entity.HouseType) (*entity.HouseType, []error) {
	itm := HouseType
	errs := HouseTypeRepo.conn.Save(itm).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// DeleteHouseType deletes a given food menu HouseType from the database
func (HouseTypeRepo *HouseTypeGormRepo) DeleteHouseType(id uint) (*entity.HouseType, []error) {
	itm, errs := HouseTypeRepo.HouseType(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = HouseTypeRepo.conn.Delete(itm, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// StoreHouseType stores a given food menu HouseType in the database
func (HouseTypeRepo *HouseTypeGormRepo) StoreHouseType(HouseType *entity.HouseType) (*entity.HouseType, []error) {
	itm := HouseType
	errs := HouseTypeRepo.conn.Create(itm).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}
