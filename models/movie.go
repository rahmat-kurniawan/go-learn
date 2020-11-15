package models

import (
	"github.com/Massad/gin-boilerplate/db"
)

type Movie struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Title     string `db:"title" json:"title"`
	Rating    int32  `db:"content" json:"content"`
	Year      int32  `db:"content" json:"content"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}

type MovieModel struct{}

func (m MovieModel) All(userID int64) (articles []DataList, err error) {
	_, err = db.GetDB().Select(&articles, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data, (SELECT row_to_json(n) FROM ( SELECT count(a.id) AS total FROM public.article AS a WHERE a.user_id=$1 LIMIT 1 ) n ) AS meta FROM ( SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.article a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC) d", userID)
	return articles, err
}
