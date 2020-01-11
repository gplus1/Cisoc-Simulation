package FeedBack

import "github.com/betsegawlemma/restaurant-rest/entity"

// FeedBackervice specifies customer FeedBack related service
type FeedBackervice interface {
	FeedBack() ([]entity.FeedBack, []error)
	FeedBack(id uint) (*entity.FeedBack, []error)
	UpdateFeedBack(FeedBack *entity.FeedBack) (*entity.FeedBack, []error)
	DeleteFeedBack(id uint) (*entity.FeedBack, []error)
	StoreFeedBack(FeedBack *entity.FeedBack) (*entity.FeedBack, []error)
}
