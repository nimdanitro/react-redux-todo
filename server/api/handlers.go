package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	//"ssomiddleware"
)

// Index simple hello page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Awesome Hackathon API!\n")
}

// ProjectIndex shows
func ProjectIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Repo.GetAll()); err != nil {
		panic(err)
	}
}

// ProjectShow is the handler function for showing a project by its ID
func ProjectShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var projectID uint64
	var err error
	if projectID, err = strconv.ParseUint(vars["projectId"], 10, 64); err != nil {
		panic(err)
	}
	project := Repo.Get(projectID)
	if project.ID > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(project); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

// ProjectDelete is the handler function for showing a project by its ID
func ProjectDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var projectID uint64
	var err error
	if projectID, err = strconv.ParseUint(vars["projectId"], 10, 64); err != nil {
		panic(err)
	}

	_, err = getUserFromContext(r)
	if err == nil {
		project := Repo.Delete(projectID)
		if project.ID > 0 {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(project); err != nil {
				panic(err)
			}
			return
		}

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	// Unauthorized
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnauthorized, Text: "You are not logged in"}); err != nil {
		panic(err)
	}

}

/*
ProjectCreate is the handler function for creating a new Project
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Project","description":"My fancy new project","creator":"foouser"}' http://localhost:8080/projects

*/
func ProjectCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	user, err := getUserFromContext(r)
	if err == nil {
		var project Project
		err = json.Unmarshal(body, &project)
		project.CreatedAt = time.Now().Unix()
		project.Description = template.HTMLEscapeString(project.Description)

		project.Creator = user
		p := Repo.Add(project)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(p); err != nil {
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
		return
	}

	// Unauthorized
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnauthorized, Text: "You are not logged in"}); err != nil {
		panic(err)
	}

}

/*
ProjectEdit is the handler function for editing a Project
Test with this curl command:

curl -H "Content-Type: application/json" -X PUT -d '{"name":"New Project","description":"My fancy new project"}' http://localhost:8080/projects

*/
func ProjectEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var projectID uint64
	var err error

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if projectID, err = strconv.ParseUint(vars["projectId"], 10, 64); err != nil {
		panic(err)
	}

	user, err := getUserFromContext(r)
	if err == nil {
		project := Repo.Get(projectID)
		if project.ID > 0 {
			if user != project.Creator {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusUnauthorized)
				if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnauthorized, Text: "You are not the creator of this project."}); err != nil {
					panic(err)
				}
				return
			}

			var updatedProject Project
			err = json.Unmarshal(body, &updatedProject)

			project.Name = updatedProject.Name
			project.Description =
			project.Description = template.HTMLEscapeString(updatedProject.Description)


			p := Repo.Save(project)

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(p); err != nil {
				panic(err)
			}
			return

		}
		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnauthorized, Text: "Not found"}); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnauthorized, Text: "You are not logged in"}); err != nil {
		panic(err)
	}
}

// ProjectJoin is the handler function for joining a project
func ProjectJoin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var projectID uint64
	var err error

	if projectID, err = strconv.ParseUint(vars["projectId"], 10, 64); err != nil {
		panic(err)
	}

	user, err := getUserFromContext(r)
	if err == nil {
		project := Repo.Get(projectID)
		if project.ID > 0 {
			project.Members = appendUserIfMissing(project.Members, user)
			p := Repo.Save(project)

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(p); err != nil {
				panic(err)
			}
			return
		}

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnauthorized, Text: "Not found"}); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnauthorized, Text: "You are not logged in"}); err != nil {
		panic(err)
	}
}

// ProjectLike is the handler function for liking a project
func ProjectLike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var projectID uint64
	var err error

	if projectID, err = strconv.ParseUint(vars["projectId"], 10, 64); err != nil {
		panic(err)
	}

	user, err := getUserFromContext(r)
	if err == nil {
		project := Repo.Get(projectID)
		if project.ID > 0 {
			project.LikeUsers = appendUserIfMissing(project.LikeUsers, user)
			project.Likes = len(project.LikeUsers)
			p := Repo.Save(project)

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(p); err != nil {
				panic(err)
			}
			return
		}

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnauthorized, Text: "Not found"}); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnauthorized, Text: "You are not logged in"}); err != nil {
		panic(err)
	}
}

func appendUserIfMissing(members []User, user User) []User {
	for _, ele := range members {
		if ele == user {
			return members
		}
	}
	return append(members, user)
}

// UserShow returns the current stored user in the context
func UserShow(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromContext(r)
	if err != nil {
		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnauthorized, Text: "You are not logged in"}); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
}

func getUserFromContext(r *http.Request) (User, error) {
	user := context.Get(r, "user")
	if user != nil {
		currentUser, _ := user.(map[string]string)
		return User{Name: currentUser["Name"], Email: currentUser["Email"], Username: currentUser["Username"]}, nil
	}

	return User{}, errors.New("not authenticated")

}
