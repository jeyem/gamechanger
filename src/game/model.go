package game

import (
	"company/bab/application"
	"strconv"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Owner struct {
	UserID    bson.ObjectId `bson:"user"`
	OwnedDate time.Time     `bson:"owned_date"`
}

type Game struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Type        string        `bson:"type"`
	Name        string        `bson:"name"`
	Description string        `bson: "description"`
	ReleaseDate time.Time     `bson: "release_date"`
	Images      []string      `bson:"images"`
	SourceURL   string        `bson:"source_url"`
	Owners      []Owner       `bson:"owners"`
	Tags        []string      `bson:"tags"`
	Keywords    []string      `bson:"keywords"`
}

func (Model) Meta() []mgo.Index {
	return []mgo.Index{
		{
			Key: []string{"keywords"},
		},
	}
}

func (g *Model) setKeywords() {
	g.Keywords = g.Tags
	year := strconv.Itoa(g.ReleaseDate.Year())
	g.Keywords = append(g.Keywords, g.Name, g.Type, year)
}

func (g *Model) Save() error {
	if g.ID.Valid() {
		return g.Update()
	}
	g.CreatedAt = time.Now()
	g.setKeywords()
	return application.DB.Create(g)
}

func (g *Model) Update() error {
	return db.Update(g)
}

func (g *Model) AppendOwner(id bson.ObjectId) {
	for _, owner := range g.Owners {
		if owner.UserID == id {
			return
		}
	}
	g.Owners = append(g.Owners, Owner{id, time.Now()})
	g.Save()
}

func LoadModels(limit, page int, filter ...string) (res []Model) {
	db.Where(bson.M{"keywors": bson.M{
		"$in": filter,
	}}).Find(&res)
	return res
}

func TotalModelPages(limit int) int {
	count := db.Collection(&Model{}).Count()
	return int(float32(count)/float32(limit)) + 1
}
