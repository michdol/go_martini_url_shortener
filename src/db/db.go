package db


type Website struct {
	Id int
	Url string
}

type DB interface {
	Get(id int) *Website
	GetAll() map[int]*Website
	Add(a *Website) (int, error)
}

type Database struct {
	m map[int]*Website
}

func (db *Database) Get(id int) *Website {
	if website, exists := db.m[id]; exists {
		return website
	}
	return &Website{}
}

func (db *Database) GetAll() map[int]*Website {
	return db.m
}

func (db *Database) Add(a *Website) (int, error) {
	id := len(db.m) + 1
	a.Id = id
	db.m[id] = a
	return a.Id, nil
}

var AppDatabase DB

func init() {
	AppDatabase = &Database{
		m: make(map[int]*Website),
	}
}