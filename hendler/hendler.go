package hendler

import (
	"fmt"
	"html/template"
	"log"
	"miniProject_golang/entity"
	"net/http"
	"path"
	"strconv"
)

func HomeHendler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm 02", http.StatusInternalServerError)
			return
		}
		// jsonFile, err := os.Open("/data/takslist.json")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println("Successfully Opened myjson")
		// defer jsonFile.Close()
		// byteValue, _ := ioutil.ReadAll(jsonFile)
		// var tasks Tasks
		// json.Unmarshal(byteValue, &tasks)
		// for i := 0; i < len(tasks.Tasks); i++ {
		// 	fmt.Println("User Type: " + tasks.Tasks[i].Type)
		// 	fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		// 	fmt.Println("User Name: " + users.Users[i].Name)
		// 	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		// }
		data := entity.Tasklist{
			ID:       1,
			Task:     "Task Golang 01",
			Assignee: "Adi",
			Deadline: "31/12/2021",
			Status:   true}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm 03", http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm 01", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm 02", http.StatusInternalServerError)
			return
		}

		data := entity.Tasklist{
			ID:       1,
			Task:     r.Form.Get("taskdetail"),
			Assignee: r.Form.Get("assignee"),
			Deadline: r.Form.Get("date"),
			Status:   false}

		// data := entity.Tasklist{
		// 	ID:       1,
		// 	Task:     "Tugas Golang O1",
		// 	Assignee: "Adi",
		// 	Deadline: "08/10/2022",
		// 	Status:   false}

		// data := []entity.Tasklist{
		// 	{ID: 1,
		// 		Task:     "Tugas Golang O1",
		// 		Assignee: "Adi",
		// 		Deadline: "08/10/2022",
		// 		Status:   false},
		// 	{ID: 2,
		// 		Task:     "Tugas Golang O2",
		// 		Assignee: "Budi",
		// 		Deadline: "08/12/2022",
		// 		Status:   false},
		// 	{ID: 3,
		// 		Task:     "Tugas Golang O3",
		// 		Assignee: "Caca",
		// 		Deadline: "08/15/2022",
		// 		Status:   false},
		// }

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm 03", http.StatusInternalServerError)
			return
		}
	}

	// tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
	// 	return
	// }

	// data := map[string]interface{}{
	//	"title":   "L'm learning Golang web",
	//	"content": "I'm learning Golang web with Agung Setiawan",
	// }

	// data := []entity.Tasklist{
	// 	{ID: 1,
	// 		Task:     "Tugas Golang O1",
	// 		Assignee: "Adi",
	// 		Deadline: "08/10/2022",
	// 		Status:   false},
	// 	{ID: 2,
	// 		Task:     "Tugas Golang O2",
	// 		Assignee: "Budi",
	// 		Deadline: "08/12/2022",
	// 		Status:   false},
	// 	{ID: 3,
	// 		Task:     "Tugas Golang O3",
	// 		Assignee: "Caca",
	// 		Deadline: "08/15/2022",
	// 		Status:   false},
	// }

	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
	// 	return
	// }
}

func CreateTaskHendler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "createTask.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func ProductHendler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Product Page"))

	fmt.Fprintf(w, "Product page : %d", idNumb)
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
	}
}
