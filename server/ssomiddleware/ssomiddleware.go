package ssomiddleware

import (
	"github.com/gorilla/context"
	"log"
	"net/http"
)

// A function called whenever an error is encountered
type errorHandler func(w http.ResponseWriter, r *http.Request, err string)

// SSOMiddleware is a gorilla middleware for SSO authentication
type SSOMiddleware struct {
    Options Options
}

// Options for the SSOMiddleware
type Options struct {
	UsernameHeader  string
	GivenNameHeader string
	SurnameHeader   string
	EmailHeader     string
    Debug           bool
    UserProperty    string
}

// User represents the user structure
type User struct {
	Name  		string `json:"name"`
	Username   	string `json:"username"`
	Email 		string `json:"email,omitempty`
}

func OnError(w http.ResponseWriter, r *http.Request, err string) {
	http.Error(w, err, http.StatusUnauthorized)
}

// New constructs a new Secure instance with supplied options.
func New(options ...Options) *SSOMiddleware {

	var opts Options
	if len(options) == 0 {
		opts = Options{}
	} else {
		opts = options[0]
	}

    if opts.UserProperty == "" {
    opts.UserProperty = "user"
}

	if opts.UsernameHeader == "" {
		opts.UsernameHeader = "X-Auth-Username"
	}

	if opts.GivenNameHeader == "" {
		opts.GivenNameHeader = "X-Auth-GivenName"
	}

	if opts.SurnameHeader == "" {
		opts.SurnameHeader = "X-Auth-Surname"
	}

	if opts.EmailHeader == "" {
		opts.EmailHeader = "X-Auth-Email"
	}

	return &SSOMiddleware{
		Options: opts,
	}
}

func (m *SSOMiddleware) logf(format string, args ...interface{}) {
	if m.Options.Debug {
		log.Printf(format, args...)
	}
}

func (m *SSOMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := m.GetUserInformation(w, r)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (m *SSOMiddleware) GetUserInformation(w http.ResponseWriter, r *http.Request) error {

    user := make(map[string]string)
	if (r.Header.Get(m.Options.UsernameHeader) == "") {
		return nil
	}

	user["Username"] = r.Header.Get(m.Options.UsernameHeader)
	user["Name"] = r.Header.Get(m.Options.GivenNameHeader) + " " + r.Header.Get(m.Options.SurnameHeader)
	user["Email"] = r.Header.Get(m.Options.EmailHeader)

    if len(user) > 0 {
        m.logf("User: %v", user)
        // If we get here, everything worked and we can set the
        // user property in context.
        context.Set(r, m.Options.UserProperty, user)
    }


	return nil
}
