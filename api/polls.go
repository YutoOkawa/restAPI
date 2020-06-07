package main

import (
	"errors"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type poll struct {
	ID      bson.ObjectId  `bson:"_id" json:"id"`
	Title   string         `json:"title"`
	Options []string       `json:"options"`
	Results map[string]int `json:"results,omitempty"`
}

func handlerPolls(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlerPollsGet(w, r)
		return
	case "POST":
		handlerPollsPost(w, r)
		return
	case "DELETE":
		handlerPollsDelete(w, r)
		return
	}
	respondHTTPErr(w, r, http.StatusNotFound)
}

func handlerPollsGet(w http.ResponseWriter, r *http.Request) {
	//respondErr(w, r, http.StatusInternalServerError, errors.New("未実装です"))
	db := GetVar(r, "db").(*mgo.Database)
	c := db.C("polls")
	var q *mgo.Query
	p := NewPath(r.URL.Path)
	if p.HasID() {
		// 特定の調査項目の詳細
		q = c.FindId(bson.ObjectIdHex(p.ID))
	} else {
		// 全ての調査項目のリスト
		q = c.Find(nil)
	}
	var result []*poll
	if err := q.All(&result); err != nil {
		respondErr(w, r, http.StatusInternalServerError, err)
		return
	}
	respond(w, r, http.StatusOK, &result)
}

func handlerPollsPost(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("未実装です"))
}

func handlerPollsDelete(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("未実装です"))
}
