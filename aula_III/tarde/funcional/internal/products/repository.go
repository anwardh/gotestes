package products

// Adicionando a Estrutura Product e seus campos rotulados
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Count    int     `json:"count"`
	Price    float64 `json:"price"`
}

/*
Criação da variável para guardar os produtos

	Corresponde a nossa Camada de Persistência de Dados
*/

// Criação da Interface e Declaração dos Métodos
type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, category string, count int, price float64) (Product, error)
	LastID() (int, error)
	// Declaração do Método Update - que cuidará de atualizar um dado
	Update(id int, name, productType string, count int, price float64) (Product, error)

	// Declaração do Método UpdateName
	UpdateName(id int, name string) (Product, error)

	// Declaração do Método Delete
	Delete(id int) error
}
