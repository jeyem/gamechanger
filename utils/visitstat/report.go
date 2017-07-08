package visitstat

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jeyem/mogo"
)

type Report struct {
	db    *mogo.DB
	query bson.M
}

func NewReport(db *mogo.DB) *Report {
	r := new(Report)
	r.db = db

	return r
}

func (r Report) Query(q bson.M) Report {
	r.query = q
	return r
}

func (r Report) Domain(domain string) Report {
	r.query["domain"] = domain
	return r
}

func (r Report) Path(path string) Report {
	r.query["path"] = path
	return r
}

func (r Report) Keywords(keys ...interface{}) Report {
	r.query["keywords"] = keys
	return r
}

func (r Report) Total() int {
	return r.count()
}

func (r Report) RangeReport(start, end time.Time) [][]int {
	var response [][]int
	for start.Equal(end) {
		r.query["created_at"] = start
		response = append(response, []int{int(start.Unix()), r.count()})
		start = start.AddDate(0, 0, 1)
	}
	return response
}

func (r Report) count() int {
	c, _ := r.db.Where(r.query).Count(&visit{})
	return c
}
