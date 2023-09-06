package tracker

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	HomePage   string = "./web/html/index.html"
	ArtistPage string = "./web/html/artist.html"
	ErrorPage  string = "./web/html/errors.html"
	SearchPage string = "./web/html/search.html"
	FilterPage string = "./web/html/filter.html"
	NotFound   string = "./web/html/not-found.html"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles(HomePage)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = Unmarshal()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Artists)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func ArtistHandl(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(url[2])
	if err != nil || len(url) > 3 || id > len(Artists) || id < 1 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles(ArtistPage)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	if len(Artists) == 0 {
		err = Unmarshal()
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	}
	err = UnmarshalRelation()
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Artists[id-1])
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func SearhHandel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	search, ok := r.Form["text"]
	if !ok {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	SearchArtist(search[0])
	tmpl, err := template.ParseFiles(SearchPage)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	if FoundArtist == nil {
		tmpl, err = template.ParseFiles(NotFound)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	}
	data := Alldata{Artists, FoundArtist}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	location, locationFlag := r.Form["inputLocation"]
	minCreate, minCreateFlag := r.Form["minCreate"]
	maxCreate, maxCreateFlag := r.Form["maxCreate"]
	minFirst, minFirstFlag := r.Form["minFirst"]
	maxFirst, maxFirstFlag := r.Form["maxFirst"]
	var membersNum []string
	membersNum = append(membersNum, r.Form["membersNum"]...)

	if !minCreateFlag || !maxCreateFlag || !maxFirstFlag || !minFirstFlag || !locationFlag {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	FilterData(minCreate[0], maxCreate[0], minFirst[0], maxFirst[0], location[0], membersNum)
	tmpl, err := template.ParseFiles(FilterPage)
	if FiltredArtists == nil {
		tmpl, err = template.ParseFiles(NotFound)
	}
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	err = tmpl.Execute(w, FiltredArtists)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, code int) {
	tmpl, err := template.ParseFiles(ErrorPage)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.WriteHeader(code)
	errInf := Er{code, http.StatusText(code)}
	err = tmpl.Execute(w, errInf)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
