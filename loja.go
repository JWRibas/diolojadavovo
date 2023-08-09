package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Usuario struct {
	Nome     string
	Tipo     string
	Idade    string
	Telefone string
}

var users struct {
	Usuarios []Usuario
}

func validateReqMethod(w http.ResponseWriter, req *http.Request) bool {
	if req.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func createUser(req *http.Request) (Usuario, error) {
	err := req.ParseForm()
	if err != nil {
		return Usuario{}, err
	}
	return Usuario{
		Nome:     req.FormValue("nome"),
		Tipo:     req.FormValue("produto"),
		Idade:    req.FormValue("idade"),
		Telefone: req.FormValue("telefone"),
	}, nil
}

func cadastrar(w http.ResponseWriter, req *http.Request) {
	if !validateReqMethod(w, req) {
		return
	}

	user, err := createUser(req)
	if err != nil {
		http.Error(w, "Erro ao ler os dados do formulário", http.StatusInternalServerError)
		return
	}

	users.Usuarios = append(users.Usuarios, user)

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Erro ao serializar os dados dos usuários", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("./api/cadastro.json", jsonData, 0644)
	if err != nil {
		http.Error(w, "Erro ao salvar o arquivo JSON", http.StatusInternalServerError)
		return
	}
}

func consultar(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	nome := req.URL.Query().Get("nome")

	var user Usuario
	for _, u := range users.Usuarios {
		if u.Nome == nome {
			user = u
			break
		}
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Erro ao serializar os dados do usuário", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/cadastrar", cadastrar)
	http.HandleFunc("/consultar", consultar)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
