package game

import (
	"company/bab/application"
	"strconv"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Game struct {
	ID          bson.ObjectId   `bson:"_id,omitempty"`
	Type        string          `bson:"type"`
	Name        string          `bson:"name"`
	Description string          `bson: "description"`
	ReleaseDate time.Time       `bson: "release_date"`
	Images      []string        `bson:"images"`
	SourceURL   string          `bson:"source_url"`
	Owners      []bson.ObjectId `bson:"owners"`
	Tags        []string        `bson:"tags"`
	Keywords    []string        `bson:"keywords"`
}

func (Game) Meta() []mgo.Index {
	return []mgo.Index{
		{
			Key: []string{"keywords"},
		},
	}
}

func (g *Game) setKeywords() {
	g.Keywords = g.Tags
	year := strconv.Itoa(g.ReleaseDate.Year())
	g.Keywords = append(g.Keywords, g.Name, g.Type, year)
}

func (g *Game) Save() error {
	if g.ID.Valid() {
		return g.Update()
	}
	g.CreatedAt = time.Now()
	g.setKeywords()
	return application.DB.Create(g)
}

func (g *Game) Update() error {
	return db.Update(g)
}

func LoadGames(limit, page int, filter ...string) (res []Game) {
	db.Where(bson.M{"keywors": bson.M{
		"$in": filter,
	}}).Find(&res)
	return res
}

func TotalGamePages(limit int) int {
	count := db.Collection(&Game{}).Count()
	return int(float32(count)/float32(limit)) + 1
}
