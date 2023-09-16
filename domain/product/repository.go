package product

type Repository interface {
	Find(id string) (Product, error)
	FindAll(limit, offset int) ([]Product, error)
	Store(p Product) (Product, error)
	Delete(id string) (Product, error)
	Update(id string, p Product) (Product, error)
}
