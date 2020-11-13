package model

import (
	"github.com/snail007/gmc"
	"sync"
)

type UserModel struct {
	db    *gmc.MySQL
	table string
	pkey  string
	once  *sync.Once
}

var (
	User = NewUserModel()
)

func NewUserModel() *UserModel {
	u := &UserModel{
		table: "user",
		pkey:  "user_id",
		once:  &sync.Once{},
	}
	return u
}

func (s *UserModel) DB() *gmc.MySQL {
	if s.db == nil {
		s.once.Do(func() {
			s.db = gmc.DB.MySQL()
		})
	}
	return s.db
}

func (s *UserModel) GetByID(id string) (ret gmc.Mss, error error) {
	db := s.DB()
	rs, err := db.Query(db.AR().From(s.table).Where(gmc.M{
		s.pkey: id,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	ret = rs.Row()
	return
}

func (s *UserModel) GetBy(where gmc.M) (ret gmc.Mss, error error) {
	db := s.DB()
	rs, err := db.Query(db.AR().From(s.table).Where(where).Limit(0, 1))
	if err != nil {
		return
	}
	ret = rs.Row()
	return
}

func (s *UserModel) MGetByIDs(ids []string, orderBy ...interface{}) (ret gmc.Mss, error error) {
	db := s.DB()
	ar := db.AR().From(s.table).Where(gmc.M{
		s.pkey: ids,
	})
	if col, by := s.OrderBy(orderBy...); col != "" {
		ar.OrderBy(col, by)
	}
	rs, err := db.Query(ar)
	if err != nil {
		return
	}
	ret = rs.Row()
	return
}

func (s *UserModel) MGetBy(where gmc.M, orderBy ...interface{}) (ret []gmc.Mss, error error) {
	db := s.DB()
	ar := db.AR().From(s.table).Where(where).Limit(0, 1)
	if col, by := s.OrderBy(orderBy...); col != "" {
		ar.OrderBy(col, by)
	}
	rs, err := db.Query(ar)
	if err != nil {
		return
	}
	ret = rs.Rows()
	return
}

func (s *UserModel) DeleteBy(where gmc.M) (cnt int64, err error) {
	db := s.DB()
	rs, err := db.Exec(db.AR().Delete(s.table, where))
	if err != nil {
		return
	}
	cnt = rs.RowsAffected
	return
}

func (s *UserModel) DeleteByIDs(ids []string) (cnt int64, err error) {
	db := s.DB()
	rs, err := db.Exec(db.AR().Delete(s.table, gmc.M{
		s.pkey: ids,
	}))
	if err != nil {
		return
	}
	cnt = rs.RowsAffected
	return
}

func (s *UserModel) Insert(data gmc.M) (cnt int64, err error) {
	db := s.DB()
	rs, err := db.Exec(db.AR().Insert(s.table, data))
	if err != nil {
		return
	}
	cnt = rs.RowsAffected
	return
}

func (s *UserModel) InsertBatch(data []gmc.M) (cnt int64, err error) {
	db := s.DB()
	rs, err := db.Exec(db.AR().InsertBatch(s.table, data))
	if err != nil {
		return
	}
	cnt = rs.RowsAffected
	return
}

func (s *UserModel) UpdateByIDs(ids []string, data gmc.M) (cnt int64, err error) {
	db := s.DB()
	rs, err := db.Exec(db.AR().Update(s.table, data, gmc.M{
		s.pkey: ids,
	}))
	if err != nil {
		return
	}
	cnt = rs.RowsAffected
	return
}

func (s *UserModel) UpdateBy(where, data gmc.M) (cnt int64, err error) {
	db := s.DB()
	rs, err := db.Exec(db.AR().Update(s.table, data, where))
	if err != nil {
		return
	}
	cnt = rs.RowsAffected
	return
}

func (s *UserModel) Page(where gmc.M, offset, length int, orderBy ...interface{}) (ret []gmc.Mss, total int, err error) {
	db := s.DB()
	ar := db.AR().Select("count(*) as total").From(s.table)
	if len(where) > 0 {
		ar.Where(where)
	}
	rs, err := db.Query(ar)
	if err != nil {
		return
	}
	total = gmc.ToInt(rs.Value("total"))

	ar = db.AR().From(s.table).Where(where).Limit(offset, length)
	if len(where) > 0 {
		ar.Where(where)
	}
	if col, by := s.OrderBy(orderBy...); col != "" {
		ar.OrderBy(col, by)
	}
	rs, err = db.Query(ar)
	if err != nil {
		return
	}
	ret = rs.Rows()
	return
}

func (s *UserModel) OrderBy(orderBy ...interface{}) (col, by string) {
	if len(orderBy) > 0 {
		switch val := orderBy[0].(type) {
		case gmc.M:
			for k, v := range val {
				col, by = k, v.(string)
				break
			}
		case gmc.Mss:
			for k, v := range val {
				col, by = k, v
				break
			}
		}
	}
	return
}
