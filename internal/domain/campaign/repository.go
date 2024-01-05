package campaign

type Repository interface {
	Save(campaign *Campaign) error
	GetAll() ([]Campaign, error)
	GetBy(id string) (*Campaign, error)
}