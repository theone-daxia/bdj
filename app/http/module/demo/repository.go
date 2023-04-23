package demo

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetUserIDs() []int {
	return []int{1, 2}
}

func (r *Repository) GetUserByIDs(ids []int) []UserModel {
	return []UserModel{
		{
			ID:   1,
			Name: "foo",
			Age:  18,
		},
		{
			ID:   2,
			Name: "bar",
			Age:  18,
		},
	}
}