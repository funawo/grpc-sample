package main

// import "net"
import "net/http"
import "fmt"
import "bytes"
import "time"
import "encoding/json"
// import "example.com/grpc-sample/funawo"
type Member struct {
	name string
	old int32
}

func (m Member) ServeHTTP(r http.ResponseWriter, request *http.Request) {
	fmt.Fprint(r, m)

}

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	DueDate time.Time `json:"due_date"`
}

var tasks = []Task{{
	ID:      1,
	Title:   "A",
	Content: "Aタスク",
	DueDate: time.Now(),
}, {
	ID:      2,
	Title:   "B",
	Content: "Bタスク",
	DueDate: time.Now(),
}, {
	ID:      3,
	Title:   "C",
	Content: "Cタスク",
	DueDate: time.Now(),
}}

type User struct {
	Name string
	Age int32
}
type UserMethod interface {
	getName() string
	getAge() int32
}

func (u *User) getName() string {
	return u.Name
}
func (u User) getAge() int32 {
	return u.Age
}
func method(u UserMethod) {
	fmt.Println(u.getAge())
	fmt.Println(u.getName())
}
func main() {
	// http.Handle("/", new(Member))
	// err := http.ListenAndServe("localhost:54101", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(tasks)
	fmt.Println(buf.String())
	var user UserMethod
	user = &User{
		Name: "funawo",
		Age: 30,
	}
	fmt.Println(user.getAge())
	fmt.Println(user.getName())
	type T struct {
		s string
	}
	u := User {
		Name: "funasshi",
		Age: 31,
	}
	method(&u)
}


