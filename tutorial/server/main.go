package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
)

//Date represent RFC3399 date
type Date time.Time

//MarshalJSON override marshalJSON
func (t *Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*t).Format(`"` + time.RFC3339 + `"`)), nil
}

//MarshalText override marshalText
func (t *Date) MarshalText() ([]byte, error) {
	return []byte(time.Time(*t).Format(`"` + time.RFC3339 + `"`)), nil
}

//UnmarshalJSON override unmarshalJSON
func (t *Date) UnmarshalJSON(b []byte) error {
	ts, err := time.Parse(`"`+time.RFC3339+`"`, string(b))
	if err != nil {
		return err
	}

	*t = Date(ts)
	return nil
}

//UnmarshalText override unmarshalText
func (t *Date) UnmarshalText(b []byte) error {
	ts, err := time.Parse(`"`+time.RFC3339+`"`, string(b))
	if err != nil {
		return err
	}

	*t = Date(ts)
	return nil
}

func (t *Date) String() string {
	return time.Time(*t).String()
}

func main() {
	// input validator
	validator.SetValidationFunc("multipleOf", multipleOf)

	r := mux.NewRouter()

	JobsInterfaceRoutes(r, JobsAPI{})

	ProjectsInterfaceRoutes(r, ProjectsAPI{})

	ResourcesInterfaceRoutes(r, ResourcesAPI{})

	log.Println("starting server")
	http.ListenAndServe(":8080", r)
}
