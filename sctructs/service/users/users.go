package usersdata

import "github.com/supperdoggy/score/sctructs"

// CreateUserRequest - struct for creating new user
type CreateUserRequest struct {
	User sctructs.User `json:"user"`
}

// CreateUserResponse - struct to return answer to request
type CreateUserResponse struct {
	// error if anything gone wrong
	Error error `json:"error"`
	// user we created goes here
	User sctructs.User `json:"user"`
}

// GetAllUsersResponse - struct to return all users to request
type GetAllUsersResponse struct {
	Users []sctructs.User `json:"users"`
	Error error           `json:"error"`
}

// DeleteRequest - struct for deleting user by input data
type DeleteRequest struct {
	// by ID
	ID int `json:"id"`
	// or by Username
	Username string `json:"username"`
	// or by Email
	Email string `json:"email"`
}

// DeleteResponse - response struct to delete user request
type DeleteResponse struct {
	Error error `json:"error"`
	// deleted User
	User sctructs.User `json:"user"`
}

// FindRequest - struct for finding specific user
type FindRequest struct {
	// by ID
	ID int `json:"id"`
	// or by Username
	Username string `json:"username"`
	// or by Email
	Email string `json:"email"`
}

// FindResponse - return struct to finding user request
type FindResponse struct {
	Error error         `json:"error"`
	Users sctructs.User `json:"users"`
}

// FindWithPasswordRequest - struct for finding user && checking his pass
type FindWithPasswordRequest struct {
	// firstly looking for user with the given Username
	Username string `json:"username"`
	// then by given Email
	Email    string `json:"email"`
	// checking Password of found user and if they are similar -> return user
	Password string `json:"password"`
}

// FindWithPasswordResponse - response struct for returning user and password
type FindWithPasswordResponse struct {
	Error error         `json:"error"`
	User  sctructs.User `json:"user"`
}
