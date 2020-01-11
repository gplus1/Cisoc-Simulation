package user

import (
	"github.com/betsegawlemma/restaurant-rest/entity"
	"github.com/HouseRentalSystem/Back_End/entity"
)

"github.com/HouseRentalSystem/Back_End/entity"
"github.com/HouseRentalSystem/Back_End/entity/menu"

// Manager_Service specifies application user related services
type Manager_Service interface {
	Users() ([]Manager.Manager, []error)
	User(id uint) (*Manager.Manager, []error)
	UpdateUser(user *Manager.Manager) (*Manager.Manager, []error)
	DeleteUser(id uint) (*Manager.Manager, []error)
	StoreUser(user *Manager.Manager) (*Manager.Manager, []error)
}


type Admin_Service interface {
	Users() ([]Admin.Admin, []error)
	User(id uint) (*Admin.Admin, []error)
	UpdateUser(user *Admin.Admin) (*Admin.Admin, []error)
	DeleteUser(id uint) (*Admin.Admin, []error)
	StoreUser(user *Admin.Admin) (*Admin.Admin, []error)
}
type House_Holder interface {
	Users() ([]LandLord.LandLord, []error)
	User(id uint) (*LandLord.LandLord, []error)
	UpdateUser(user *LandLord.LandLord) (*LandLord.LandLord, []error)
	DeleteUser(id uint) (*LandLord.LandLord, []error)
	StoreUser(user *LandLord.LandLord) (*LandLord.LandLord, []error)
}
type Tourist interface {
	Users() ([]Tourist.Tourist, []error)
	User(id uint) (*Tourist.Tourist, []error)
	UpdateUser(user *Tourist.Tourist) (*Tourist.Tourist, []error)
	DeleteUser(id uint) (*Tourist.Tourist, []error)
	StoreUser(user *Tourist.Tourist) (*Tourist.Tourist, []error)
}

// RoleService speifies application user role related services
type RoleService interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}


