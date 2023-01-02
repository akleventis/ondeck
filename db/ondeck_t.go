package db

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone int    `json:"phone"`
}

type Drink struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Order struct {
	Person      Person  `json:"person"`
	Drinks      []Drink `json:"drinks"`
	OrderNumber int     `json:"order_number"`
}

type QueueObject struct {
	Person Person  `json:"person"`
	Drinks []Drink `json:"drinks"`
}
type QueueResponse struct {
	Queue map[int]QueueObject `json:"queue"`
}

// Person
func (db *DB) CreatePerson(p *Person) (*Person, error) {
	query := `INSERT INTO persons_t (name, phone) VALUES ($1, $2, $3);`
	if _, err := db.Exec(query, p.Name, p.Phone); err != nil {
		return nil, err
	}
	return p, nil
}

func (db *DB) RetrievePerson(p *Person) (*Person, error) {
	return nil, nil
}

func (db *DB) UpdatePerson(p *Person) (*Person, error) {
	return nil, nil
}

func (db *DB) RemovePerson(p *Person) error {
	return nil
}

// Drinks
func (db *DB) RetrieveDrinks() ([]*Drink, error) {
	return nil, nil
}

// Drink
func (db *DB) CreateDrink(drink *Drink) (*Drink, error) {
	return nil, nil
}

func (db *DB) RetrieveDrink(name string) (*Drink, error) {
	return nil, nil
}

func (db *DB) UpdateDrink(drink *Drink) (*Drink, error) {
	return nil, nil
}

func (db *DB) RemoveDrink(name string) error {
	return nil
}

// Order
func (db *DB) CreateOrder(order *Order) (*Order, error) {
	return nil, nil
}

func (db *DB) RemoveOrder(orderNumber int) error {
	return nil
}

// Orders
func (db *DB) RetrieveOrders(personID int) ([]*Order, error) {
	return nil, nil
}

// Queue
func (db *DB) RetrieveQueue() (*QueueResponse, error) {
	return nil, nil
}
