/**
 * @auhtor Lucas Burg <lucasburgmota@gmail.com>
 */

package main

import (
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

/**
 * Modelo da tabela do banco de dados
 * Essa estrutura será hidratada com os registros do banco de dados 
 */
type Scorecard struct {
	Unitid string
	Opeid string
	Opeid6 string
	Instnm string
	City string
	Stabbr string
	Insturl string
}

/**
 * Conecta ao banco e consulta os registro conforme os critérios da documentação.
 *
 * @var sct_column_search string Nome da coluna da tabela do banco de dados.
 * @var txt_search string texto que será procurado na coluna informada.
 * @return []Scorecard Coleção de registros basedo na estrutura modelo.
 */
func listScorecard(sct_column_search string, txt_search string) ([]Scorecard) {
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

	rows_scorecard := []Scorecard{}

	for rows.Next() {
		var unitid string
		var opeid string
		var opeid6 string
		var instnm string
		var city string
		var stabbr string
		var insturl string
		
		err = rows.Scan(&unitid, &opeid, &opeid6, &instnm, &city, &stabbr, &insturl)
		
		if err != nil {
			log.Fatal(err)
		}
		
		current := Scorecard{Unitid: unitid, Opeid: opeid, Opeid6: opeid, Instnm: instnm, City: city, Stabbr: stabbr, Insturl: insturl}
		rows_scorecard = append(rows_scorecard, current)
	}

	return rows_scorecard
}

/**
 * Essa função deve receber os inputs e retornar a view do template
 */
func scorecardSearchHandler(w http.ResponseWriter, r *http.Request) {
	sct_column_search := r.FormValue("sct_column_search")
	txt_search := r.FormValue("txt_search")
	list := listScorecard(sct_column_search, txt_search)
	t, _ := template.ParseFiles("template/table.html")
	//fmt.Println(list)
    t.Execute(w, list)
}

/**
 * Retorna a view inicial
 */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
    t.Execute(w, nil)
}

/**
 * Inicia a aplicação, setando as funções de controller e diretorio publico para arquivos staticos
 */
func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/scorecard_search", scorecardSearchHandler) 
	fmt.Println("start on localhost:8000")
	http.ListenAndServe(":8000", nil)
}
