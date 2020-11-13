package Server

import (
	"AITUBank/pkg/middleware/logger"
	"AITUBank/pkg/middleware/recoverer"
	"AITUBank/pkg/models"
	"AITUBank/pkg/regexpmux"
	"AITUBank/pkg/service"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
)
const doc = `
<!DOCTYPE html>
<html>
    <head>
        <title>AITU bank</title>
    </head>
    <body>
        <h3>Hi, {{.Name}} We are glad that you use our bank! </h3>
		<ui>
		<li>Here you can find information regarding cards here: <a href="/getCards?uid=2">Cards</a></li>
		<li>Here you can find information regarding transactions here: <a href="/getTransactions?cid=3">Transactions</a></li>
		<li>Stats of your account: <a href="/stats?cid=3">Stats</a></li>
		</ui>
    </body>
</html>
`
type Server struct {
	service service.ServiceInterface
	Mux *remux.ReMUX
}
func NewServer(service service.ServiceInterface, mux *remux.ReMUX) *Server {
	return &Server{service: service, Mux: mux}
}

func (s *Server) Init() error{
	myLogger := logger.Logger
	myRecoverer := recoverer.Recoverer
	if err := s.Mux.NewPlain(remux.GET,  "/getCards", http.HandlerFunc(s.getCards),myLogger, myRecoverer); err != nil{
		return err
	}
	if err := s.Mux.NewPlain(remux.GET,  "/getTransactions", http.HandlerFunc(s.getTransactions),myLogger, myRecoverer); err != nil{
		return err
	}
	if err := s.Mux.NewPlain(remux.GET,  "/getMostvisited", http.HandlerFunc(s.getMostVisited),myLogger, myRecoverer); err != nil{
		return err
	}
	if err := s.Mux.NewPlain(remux.GET,  "/getMostspent", http.HandlerFunc(s.getMostSpent),myLogger, myRecoverer); err != nil{
		return err
	}
	if err := s.Mux.NewPlain(remux.GET,  "/stats", http.HandlerFunc(s.getStats),myLogger, myRecoverer); err != nil{
		return err
	}
	if err := s.Mux.NewPlain(remux.GET,  "/login", http.HandlerFunc(s.login),myLogger, myRecoverer); err != nil{
		return err
	}
	if err := s.Mux.NewPlain(remux.POST,  "/login", http.HandlerFunc(s.login),myLogger, myRecoverer); err != nil{
		return err
	}

	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}
func(s*Server) getCards(w http.ResponseWriter, r *http.Request){
	suid := r.URL.Query().Get("uid")
	if suid != ""{
		uid, err := strconv.ParseInt(suid, 10, 64)
		if err != nil{
			response := models.Result{Result: "Error", ErrorDescription: "Bad uid"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		cards, err := s.service.GetCards(uid)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
		}
		if len(cards) == 0{
			response := models.Result{Result: "Error", ErrorDescription: "No results"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		respBody, err := json.Marshal(cards)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)

	} else {
		response := models.Result{Result: "Error", ErrorDescription: "No uid"}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
	}
}

func (s*Server)  getTransactions(w http.ResponseWriter, r *http.Request){
	scid := r.URL.Query().Get("cid")
	if scid !=""{
		cid, err := strconv.ParseInt(scid, 10, 64)
		if err != nil{
			response := models.Result{Result: "Error", ErrorDescription: "Bad cid"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		transactions, err := s.service.GetTransactions(cid)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(transactions) == 0{
			response := models.Result{Result: "Error", ErrorDescription: "No results"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		respBody, err := json.Marshal(transactions)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
	} else {
		response := models.Result{Result: "Error", ErrorDescription: "No cid"}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
		return
	}
}
func(s*Server) getMostSpent(w http.ResponseWriter, r*http.Request){
	scid := r.URL.Query().Get("cid")
	if scid !=""{
		cid, err := strconv.ParseInt(scid, 10, 64)
		if err != nil{
			response := models.Result{Result: "Error", ErrorDescription: "Bad cid"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		mcc, value, err := s.service.GetMostSpent(cid)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if mcc == ""{
			response := models.Result{Result: "Error", ErrorDescription: "No such card"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		response := models.MostSpentDTO{Mcc: mcc , Value: value}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
	} else {
		response := models.Result{Result: "Error", ErrorDescription: "No cid"}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
		return
	}
}

func(s*Server) getMostVisited(w http.ResponseWriter, r*http.Request){
	scid := r.URL.Query().Get("cid")
	if scid !=""{
		cid, err := strconv.ParseInt(scid, 10, 64)
		if err != nil{
			response := models.Result{Result: "Error", ErrorDescription: "Bad cid"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		mcc, value, err := s.service.GetMostVisited(cid)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if mcc == ""{
			response := models.Result{Result: "Error", ErrorDescription: "No such card"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		response := models.MostSpentDTO{Mcc: mcc , Value: value}
		respBody, err := json.Marshal(response)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
	} else {
		response := models.Result{Result: "Error", ErrorDescription: "No cid"}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
		return
	}
}
func(s*Server) getStats(w http.ResponseWriter, r*http.Request){
	scid := r.URL.Query().Get("cid")
	if scid !=""{
		cid, err := strconv.ParseInt(scid, 10, 64)
		if err != nil{
			response := models.Result{Result: "Error", ErrorDescription: "Bad cid"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		mccVisited, valueVisited, err := s.service.GetMostVisited(cid)
		mccSpent, valueSpent, err := s.service.GetMostSpent(cid)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if mccVisited == "" || mccSpent == ""{
			response := models.Result{Result: "Error", ErrorDescription: "No such card"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		response := models.StatsDTO{MccSpent: mccSpent,MccVisited: mccVisited, ValueSpent: valueSpent, ValueVisited: valueVisited}
		respBody, err := json.Marshal(response)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
	} else {
		response := models.Result{Result: "Error", ErrorDescription: "No cid"}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
		return
	}
}
func (s*Server) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("username")); !m {
			result := models.Result{"Error", "Wrong format"}
			respbody, _ := json.Marshal(result)
			makeResponse(respbody, w, r)
			return
		}
		name, ok, err := s.service.Login(r.Form.Get("username"), r.Form.Get("password"))
		if !ok || err !=nil{
			result := models.Result{"Error", "Not found"}
			respbody, _ := json.Marshal(result)
			makeResponse(respbody, w, r)
			return
		} else {
			w.Header().Add("Content Type", "text/html")
			// The template name "template" does not matter here
			templates := template.New("template")
			// "doc" is the constant that holds the HTML content
			templates.New("doc").Parse(doc)
			context := models.SimpleDTO{Name:name}
			templates.Lookup("doc").Execute(w, context)
			return
		}
	}
}
func makeResponse(respBody []byte, w http.ResponseWriter, r*http.Request) {
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(respBody)
	if err != nil {
		log.Println(err)
	}
}