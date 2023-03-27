package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type DB struct {
	*sql.DB
}

func Open() (*DB, error) {
	var (
		dbUser = os.Getenv("DB_USER")
		dbName = os.Getenv("DB_NAME")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	logrus.Printf("%s database connected", dbName)

	odDB := &DB{
		db,
	}

	if err := odDB.createPersonsTable(); err != nil {
		return nil, err
	}
	if err := odDB.createDrinksTable(); err != nil {
		return nil, err
	}
	if err := odDB.createOrdersTable(); err != nil {
		return nil, err
	}

	return odDB, nil
}

// ondeck=# \d persons_t;
// Table "public.persons_t"
//  Column |         Type          | Collation | Nullable |                Default
// --------+-----------------------+-----------+----------+---------------------------------------
//  id     | integer               |           | not null | nextval('persons_t_id_seq'::regclass)
//  phone  | character varying(10) |           |          |
//  name   | character varying(50) |           |          |
func (db *DB) createPersonsTable() error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS persons_t (
		id serial primary key,
		phone varchar(10),
		name varchar(50)
	  )`); err != nil {
		return err
	}
	return nil
}

// ondeck=# \d drinks_t;
// Table "public.drinks_t"
// Column |         Type          | Collation | Nullable |               Default
// --------+-----------------------+-----------+----------+--------------------------------------
//  id     | integer               |           | not null | nextval('drinks_t_id_seq'::regclass)
//  name   | character varying(50) |           |          |
//  price  | integer               |           |          |
func (db *DB) createDrinksTable() error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS drinks_t (
		id serial primary key,
		name varchar(50),
		price integer
	  )`); err != nil {
		return err
	}
	return nil
}

// ondeck=# \d orders_t
// Table "public.orders_t"
// Column    |  Type   | Collation | Nullable |                    Default
// --------------+---------+-----------+----------+------------------------------------------------
//  order_number | integer |           | not null | nextval('orders_t_order_number_seq'::regclass)
//  person_id    | integer |           |          |
//  drinks       | json    |           |          |
//  done         | boolean |           |          |
// Indexes:
// "orders_t_pkey" PRIMARY KEY, btree (order_number)
// Foreign-key constraints:
// "orders_t_person_id_fkey" FOREIGN KEY (person_id) REFERENCES persons_t(id)
func (db *DB) createOrdersTable() error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS orders_t (
		order_number serial primary key,
		person_id integer references persons_t(id),
		drinks json,
		done boolean
	  )`); err != nil {
		return err
	}
	return nil
}
