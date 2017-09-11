package main

// Project represents a project definition
type Project struct {
	ID          uint64 	`json:"id"`
	Name        string 	`json:"name"`
	Description string 	`json:"description"`
	Creator     User	`json:"creator"`
	CreatedAt   int64 	`json:"createdAt"`
	Members		[]User	`json:"members"`
	LikeUsers	[]User	`json:"likeUsers"`
	Likes		int		`json:"likes"`
}

// User represents the user structure
type User struct {
	Name  		string `json:"name"`
	Username   	string `json:"username"`
	Email 		string `json:"email",omitempty`
}

// Projects is a list of Projects.. simple as that
type Projects []Project
