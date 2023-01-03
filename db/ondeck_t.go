package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"
)

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone int    `json:"phone"`
}

type Drink struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type DrinkOrder struct {
	DrinkID int    `json:"drink_id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Comment string `json:"comment"`
}

type Drinks []DrinkOrder

type Order struct {
	Order Drinks `json:"order"`
}

type FullOrder struct {
	Person      Person `json:"person"`
	Drinks      Drinks `json:"drinks"`
	OrderNumber int    `json:"order_number"`
}

type QueueObject struct {
	Person      Person `json:"person"`
	Drinks      Drinks `json:"drinks"`
	OrderNumber int    `json:"order_number"`
}
type QueueResponse struct {
	Queue map[int]QueueObject `json:"queue"`
}

// Person
func (db *DB) CreatePerson(p *Person) (*Person, error) {
	query := `INSERT INTO persons_t (name, phone) VALUES ($1, $2) RETURNING id;`

	err := db.QueryRow(query, p.Name, p.Phone).Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (db *DB) RetrievePerson(id string) (*Person, error) {
	var p Person
	query := `SELECT id, name, phone FROM persons_t WHERE id=$1 LIMIT 1;`
	if err := db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Phone); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (db *DB) UpdatePerson(p *Person) (*Person, error) {
	query := `UPDATE persons_t SET name=$1, phone=$2 WHERE id=$3;`
	if _, err := db.Exec(query, p.Name, p.Phone, p.ID); err != nil {
		return nil, err
	}
	return p, nil
}

func (db *DB) RemovePerson(id string) error {
	query := `DELETE FROM persons_t WHERE id=$1;`
	if _, err := db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

func (db *DB) RetrievePersons() ([]Person, error) {
	persons := make([]Person, 0)
	query := `SELECT id, name, phone FROM persons_t;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		person := Person{}
		if err := rows.Scan(&person.ID, &person.Name, &person.Phone); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return persons, nil
}

// Drinks
func (db *DB) RetrieveDrinks() ([]Drink, error) {
	drinks := make([]Drink, 0)
	query := `SELECT id, name, price FROM drinks_t;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		drink := Drink{}
		if err := rows.Scan(&drink.ID, &drink.Name, &drink.Price); err != nil {
			return nil, err
		}
		drinks = append(drinks, drink)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return drinks, nil
}

// Drink
func (db *DB) CreateDrink(d *Drink) (*Drink, error) {
	query := `INSERT INTO drinks_t (name, price) VALUES ($1, $2) RETURNING id;`
	err := db.QueryRow(query, d.Name, d.Price).Scan(&d.ID)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (db *DB) RetrieveDrink(id string) (*Drink, error) {
	var d Drink
	query := `SELECT id, name, price FROM drinks_t WHERE id=$1 LIMIT 1;`
	if err := db.QueryRow(query, id).Scan(&d.ID, &d.Name, &d.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

func (db *DB) UpdateDrink(d *Drink) (*Drink, error) {
	query := `UPDATE drinks_t SET name=$1, price=$2 WHERE id=$3`
	if _, err := db.Exec(query, d.Name, d.Price, d.ID); err != nil {
		return nil, err
	}
	return d, nil
}

func (db *DB) RemoveDrink(id string) error {
	query := `DELETE FROM drinks_t WHERE id=$1`
	if _, err := db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

// Order
func (db *DB) CreateOrder(order *FullOrder) (*FullOrder, error) {
	query := `INSERT INTO orders_t (person_id, drinks) VALUES ($1, $2) RETURNING order_number;`

	jsonDrinks, err := json.Marshal(order.Drinks)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(query, order.Person.ID, jsonDrinks).Scan(&order.OrderNumber)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (db *DB) RemoveOrder(orderNumber string) error {
	query := `DELETE FROM orders_t WHERE order_number=$1`
	if _, err := db.Exec(query, orderNumber); err != nil {
		return err
	}
	return nil
}

// Orders
func (db *DB) RetrieveOrders(personID string) ([]FullOrder, error) {
	orders := make([]FullOrder, 0)
	query := `SELECT order_number, person_id, drinks from orders_t WHERE person_id=$1;`
	rows, err := db.Query(query, personID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		order := FullOrder{}
		if err := rows.Scan(&order.OrderNumber, &order.Person.ID, &order.Drinks); err != nil {
			return nil, err
		}

		p, err := db.RetrievePerson(personID)
		if err != nil {
			return nil, err
		}
		order.Person.Name = p.Name
		order.Person.Phone = p.Phone

		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

// Queue
func (db *DB) RetrieveQueue() (*QueueResponse, error) {
	var queueResp QueueResponse
	queue := make(map[int]QueueObject, 0)
	query := `SELECT order_number, person_id, drinks from orders_t;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		q := QueueObject{}
		var orderNumber int
		if err := rows.Scan(&orderNumber, &q.Person.ID, &q.Drinks); err != nil {
			return nil, err
		}
		personID := strconv.Itoa(q.Person.ID)

		p, err := db.RetrievePerson(personID)
		if err != nil {
			return nil, err
		}
		q.Person.Name = p.Name
		q.Person.Phone = p.Phone
		q.OrderNumber = orderNumber

		queue[orderNumber] = q
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	queueResp.Queue = queue
	return &queueResp, nil
}

func (d *Drinks) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, d)
	case string:
		return json.Unmarshal([]byte(v), d)
	}
	return errors.New("type assertion failed")
}
