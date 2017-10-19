package main

import (
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

type Scorecard struct {
	Unitid string
	Opeid string
	Opeid6 string
	Instnm string
	City string
	Stabbr string
	Insturl string
}



func listScorecard(sct_column_search string, txt_search string) []Scorecard {
	db, err := sql.Open("postgres", "user=postgres dbname=transparency password=99253164 sslmode=disable")
	if err != nil {
		log.Fatal("Error: parametros invalidos para a conexão com o database")
	}
	query := fmt.Sprintf("select unitid, opeid, opeid6, instnm, city, stabbr, insturl from scorecard where %s = '%s' order by stabbr, city, instnm asc limit 100", sct_column_search, txt_search)
	rows, err := db.Query(query)
	if err != nil {
	  log.Fatal(err)
	}
	defer rows.Close()

	rows_scorecard := []

	for rows.Next() {
		s := Scorecard{}
		err = rows.Scan(&s.Unitid, &s.Opeid, &s.Opeid6, &s.Instnm, &s.City, &s.Stabbr, &s.Insturl)
		
		
		
		if err != nil {
			log.Fatal(err)
		}

		log.Println(s)

		rows_scorecard = append(rows_scorecard, s)
	}

	//fmt.Println(rows_scorecard)

	return rows_scorecard
}

func scorecardSearchHandler(w http.ResponseWriter, r *http.Request) {
	sct_column_search := r.FormValue("sct_column_search")
	txt_search := r.FormValue("txt_search")
	list := listScorecard(sct_column_search, txt_search)
	t, _ := template.ParseFiles("template/table.html")
	//fmt.Println(list)
    t.Execute(w, list)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
    t.Execute(w, nil)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/scorecard_search", scorecardSearchHandler) 
	fmt.Println("start on localhost:8000")
	http.ListenAndServe(":8000", nil)
}
