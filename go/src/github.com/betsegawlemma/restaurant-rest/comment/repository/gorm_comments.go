package repository

import (
	"github.com/betsegawlemma/restaurant-rest/FeedBack"
	"github.com/betsegawlemma/restaurant-rest/entity"
	"github.com/jinzhu/gorm"
)

// FeedBackGormRepo implements menu.FeedBackRepository interface
type FeedBackGormRepo struct {
	conn *gorm.DB
}

// NewFeedBackGormRepo returns new object of FeedBackGormRepo
func NewFeedBackGormRepo(db *gorm.DB) FeedBack.FeedBackRepository {
	return &FeedBackGormRepo{conn: db}
}

// FeedBacks returns all customer FeedBacks stored in the database
func (cmntRepo *FeedBackGormRepo) FeedBacks() ([]entity.FeedBack, []error) {
	cmnts := []entity.FeedBack{}
	errs := cmntRepo.conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// FeedBack retrieves a customer FeedBack from the database by its id
func (cmntRepo *FeedBackGormRepo) FeedBack(id uint) (*entity.FeedBack, []error) {
	cmnt := entity.FeedBack{}
	errs := cmntRepo.conn.First(&cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cmnt, errs
}

// UpdateFeedBack updates a given customer FeedBack in the database
func (cmntRepo *FeedBackGormRepo) UpdateFeedBack(FeedBack *entity.FeedBack) (*entity.FeedBack, []error) {
	cmnt := FeedBack
	errs := cmntRepo.conn.Save(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// DeleteFeedBack deletes a given customer FeedBack from the database
func (cmntRepo *FeedBackGormRepo) DeleteFeedBack(id uint) (*entity.FeedBack, []error) {
	cmnt, errs := cmntRepo.FeedBack(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cmntRepo.conn.Delete(cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// StoreFeedBack stores a given customer FeedBack in the database
func (cmntRepo *FeedBackGormRepo) StoreFeedBack(FeedBack *entity.FeedBack) (*entity.FeedBack, []error) {
	cmnt := FeedBack
	errs := cmntRepo.conn.Create(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
