package main

import (
	"encoding/json"
	"exam/data"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type M map[string]interface{}

/* заявки */
type Route struct {
	GuideId      string `json:"guide_id"`
	RouteId      string `json:"route_id"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	Duration     string `json:"duration"`
	Persons      string `json:"persons"`
	Price        string `json:"price"`
	OptionFirst  string `json:"option_first"`
	OptionSecond string `json:"option_second"`
	StudentId    string `json:"student_id"`
}

/* гиды */
type Guides struct {
	Name         string `json:"name"`
	Lang         string `json:"lang"`
	PricePerHour string `json:"price_per_hour"`
	RouteId      string `json:"route_route`
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	u := data.Routes{}
	u.GetAll()
	t, _ := template.ParseFiles("LK/main.html")
	t.Execute(w, u)
}

func about(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("LK/about.html")
	t.Execute(w, nil)
}

func maps(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("LK/map.html")
	t.Execute(w, nil)
}

func route(w http.ResponseWriter, r *http.Request) {
	coordinations := strings.Split(r.URL.Query().Get("data"), "|")
	u := data.Location{
		Start: coordinations[0],
		Stop:  coordinations[1],
	}

	t, _ := template.ParseFiles("LK/route.html")
	t.Execute(w, u)
}

func orders(w http.ResponseWriter, r *http.Request) {
	u := data.Orders{Arr: [][]string{}}
	for _, c := range r.Cookies() {
		c.Value, _ = url.PathUnescape(c.Value)
		if strings.Contains(c.Name, "order") && len(strings.Split(c.Value, "|")) == 11 {
			u.Arr = append(u.Arr, getCookiesList(r, strings.Replace(c.Name, "order", "", -1)))
		}
	}

	t, _ := template.ParseFiles("LK/order.html")
	t.Execute(w, u)
}

func getCookiesList(r *http.Request, id string) []string {
	cookie, _ := r.Cookie("order" + id)
	cookie.Value, _ = url.PathUnescape(cookie.Value)
	cookie_list := strings.Split((cookie.Value), "|")

	return cookie_list
}

// api func
func getRoutests(w http.ResponseWriter, r *http.Request) {
	response := []M{}
	for _, c := range r.Cookies() {
		c.Value, _ = url.PathUnescape(c.Value)
		if strings.Contains(c.Name, "route") && len(strings.Split(c.Value, "|")) == 5 {
			tmp_list := strings.Split(c.Value, "|")
			response = append(response, M{
				"id":          strings.Replace(c.Name, "route", "", -1),
				"name":        tmp_list[0],
				"description": tmp_list[1],
				"mainObject":  tmp_list[2],
				"start":       tmp_list[3],
				"stop":        tmp_list[4],
			})
		}
	}

	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func getGuides(w http.ResponseWriter, r *http.Request) {
	response := []M{}
	for _, c := range r.Cookies() {
		c.Value, _ = url.PathUnescape(c.Value)
		if strings.Contains(c.Name, "guid") && len(strings.Split(c.Value, "|")) == 4 {
			tmp_list := strings.Split(c.Value, "|")
			response = append(response, M{
				"id":             strings.Replace(c.Name, "guid", "", -1),
				"name":           tmp_list[0],
				"workExperience": tmp_list[1],
				"language":       tmp_list[2],
				"pricePerHour":   tmp_list[3],
			})
		}
	}

	if len(response) == 0 {
		response = append(response, M{
			"name":           "Игнатов Марат Игоревич",
			"workExperience": "7 лет",
			"language":       "Русский, Английский",
			"pricePerHour":   "1000",
			"route":          "Арбат - душа Москвы",
		})
		response = append(response, M{
			"name":           "Бурунов Сергей Маратович",
			"workExperience": "10 лет",
			"language":       "Русский",
			"pricePerHour":   "1500",
			"route":          "Исторический маршрут для всей семьи",
		})
		response = append(response, M{
			"name":           "Осин Кирилл Андреевич",
			"workExperience": "12 лет",
			"language":       "Испанский",
			"pricePerHour":   "2500",
			"route":          "Ахматовские места в Москве",
		})
		response = append(response, M{
			"name":           "Репчанский Николай Евгеньевич",
			"workExperience": "7 лет",
			"language":       "Русский",
			"pricePerHour":   "1250",
			"route":          "Нескучный сад",
		})
		response = append(response, M{
			"name":           "Сазонов Иван Максимович",
			"workExperience": "17 лет",
			"language":       "Немецкий, Арабский, Русский",
			"pricePerHour":   "3200",
			"route":          "Авторский маршрут Королькова А.Ю. «Правда о Соколе»",
		})
	}

	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	response := []M{}
	for _, c := range r.Cookies() {
		c.Value, _ = url.PathUnescape(c.Value)
		if strings.Contains(c.Name, "order") && len(strings.Split(c.Value, "|")) == 11 {
			tmp_list := strings.Split(c.Value, "|")
			response = append(response, M{
				"id":           strings.Replace(c.Name, "order", "", -1),
				"guide_id":     tmp_list[0],
				"route_id":     tmp_list[1],
				"date":         tmp_list[2],
				"time":         tmp_list[3],
				"duration":     tmp_list[4],
				"persons":      tmp_list[5],
				"price":        tmp_list[6],
				"optionFirst":  tmp_list[7],
				"optionSecond": tmp_list[8],
				"student_id":   tmp_list[9],
			})
		}
	}

	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func addOrders(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rou Route
	if err := decoder.Decode(&rou); err != nil {
		panic(err)
	}

	uuidString := uuid.New().String()
	response := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s",
		uuidString,
		strings.Replace(rou.GuideId, "+", " ", -1),
		rou.RouteId,
		rou.Date,
		rou.Time,
		rou.Duration,
		rou.Persons,
		rou.Price,
		rou.OptionFirst,
		rou.OptionSecond,
		rou.StudentId)

	fmt.Println(response)

	//set cookie
	cookie := &http.Cookie{
		Name:     "order" + uuidString,
		Value:    url.QueryEscape(response),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	fmt.Println("cookies is setting")
}

func putOrders(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rou Route
	if err := decoder.Decode(&rou); err != nil {
		panic(err)
	}

	response := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s",
		mux.Vars(r)["id"],
		strings.Replace(rou.GuideId, "+", " ", -1),
		rou.RouteId,
		rou.Date,
		rou.Time,
		rou.Duration,
		rou.Persons,
		rou.Price,
		rou.OptionFirst,
		rou.OptionSecond,
		rou.StudentId)

	fmt.Println(response)

	//set cookie
	cookie := &http.Cookie{
		Name:     "order" + mux.Vars(r)["id"],
		Value:    url.QueryEscape(response),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	fmt.Println("cookies is setting")
}

func deleteOrders(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	cookie := &http.Cookie{
		Name:     "order" + id,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}

func getOrderInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response := []M{}
	for _, c := range r.Cookies() {
		c.Value, _ = url.PathUnescape(c.Value)
		if c.Name == ("order"+id) && len(strings.Split(c.Value, "|")) == 11 {
			tmp_list := strings.Split(c.Value, "|")
			response = append(response, M{
				"id":           strings.Replace(c.Name, "order", "", -1),
				"guide_id":     tmp_list[0],
				"route_id":     tmp_list[1],
				"date":         tmp_list[2],
				"time":         tmp_list[3],
				"duration":     tmp_list[4],
				"persons":      tmp_list[5],
				"price":        tmp_list[6],
				"optionFirst":  tmp_list[7],
				"optionSecond": tmp_list[8],
				"student_id":   tmp_list[9],
			})
			break
		}
	}

	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func getGuideInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response := []M{}
	for _, c := range r.Cookies() {
		c.Value, _ = url.PathUnescape(c.Value)
		if c.Name == ("guid"+id) && len(strings.Split(c.Value, "|")) == 4 {
			tmp_list := strings.Split(c.Value, "|")
			response = append(response, M{
				"id":             strings.Replace(c.Name, "guid", "", -1),
				"name":           tmp_list[0],
				"workExperience": tmp_list[1],
				"language":       tmp_list[2],
				"pricePerHour":   tmp_list[3],
			})
		}
	}

	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func run() {
	http.Handle("/style/",
		http.StripPrefix("/style/",
			http.FileServer(http.Dir("./style/"))))
	http.Handle("/js/",
		http.StripPrefix("/js/",
			http.FileServer(http.Dir("./js/"))))
	http.Handle("/LK/",
		http.StripPrefix("/LK/",
			http.FileServer(http.Dir("./LK/"))))

	rout := mux.NewRouter()
	rout.HandleFunc("/", homePage).Methods("GET")
	rout.HandleFunc("/about/", about).Methods("GET")
	rout.HandleFunc("/map/", maps).Methods("GET")
	rout.HandleFunc("/route/", route).Methods("GET")
	rout.HandleFunc("/order/", orders).Methods("GET")

	// api
	rout.HandleFunc("/api/routes/", getRoutests).Methods("GET")
	rout.HandleFunc("/api/guides/", getGuides).Methods("GET")
	rout.HandleFunc("/api/orders/", getOrders).Methods("GET")
	rout.HandleFunc("/api/orders/", addOrders).Methods("POST")
	rout.HandleFunc("/api/orders/{id:[-a-z0-9]+}", putOrders).Methods("PUT")
	rout.HandleFunc("/api/orders/{id:[-a-z0-9]+}", deleteOrders).Methods("DELETE")
	rout.HandleFunc("/api/orders/{id:[-a-z0-9]+}", getOrderInfo).Methods("GET")
	rout.HandleFunc("/api/guides/{id:[-a-z0-9]+}/", getGuideInfo).Methods("GET")

	http.Handle("/", rout) // перенаправление на роутер
	http.ListenAndServe(":8080", nil)
}

func main() {
	log.Println("SERVER START")
	run()
}
