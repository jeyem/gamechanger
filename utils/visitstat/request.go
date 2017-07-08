package visitstat

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jeyem/mogo"
)

type visit struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Domain    string        `bson:"domain"`
	Path      string        `bson:"path"`
	IP        string        `bson:"ip"`
	Keywords  []interface{} `bson:"keywords"`
	CreatedAt time.Time     `bson:"created_at"`
	db        *mogo.DB      `bson:"-"`
}

func Request(r *http.Request, db *mogo.DB, keywords ...interface{}) {
	visit := new(visit)
	visit.Keywords = keywords
	visit.db = db
	go visit.parse(r)
}

func (v *visit) parse(request *http.Request) {
	v.IP = strings.Split(request.RemoteAddr, ":")[0]
	host := request.Host
	if host == "" {
		host = request.URL.Host
	}
	v.Domain = strings.Split(host, ":")[0]
	v.Path = request.URL.Path
	v.save()
}

func (v *visit) save() {
	if v.isStatic() {
		return
	}
	v.CreatedAt = time.Now()
	lastVisit := new(visit)
	if err := v.db.Where(bson.M{
		"ip":     v.IP,
		"domain": v.Domain,
		"path":   v.Path,
	}).Find(lastVisit); err == nil {
		duration := v.CreatedAt.Sub(lastVisit.CreatedAt)
		if duration < (time.Minute * 15) {
			return
		}
	}
	v.db.Create(v)
}

func (v *visit) isStatic() bool {
	regx := regexp.MustCompile(`static`)
	if regx.MatchString(v.Path) {
		return true
	}
	splited := strings.Split(v.Path, ".")
	return len(splited) > 1
}

func (v *visit) Meta() []mgo.Index {
	return []mgo.Index{
		{Key: []string{"keywords"}},
		{Key: []string{"ip", "keywords"}},
		{Key: []string{"ip", "create_at"}},
	}
}
