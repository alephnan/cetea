package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	verifier "github.com/alephnan/google-auth-id-token-verifier"
	"github.com/dgrijalva/jwt-go"
	xsrf "golang.org/x/net/xsrftoken"
	"golang.org/x/oauth2"
)

type ContainerClaims struct {
	ContainedJwt string `json:"jwt"`
	jwt.StandardClaims
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

type AuthorizationStruct struct {
	Code     string
	Id_Token string
}

// TODO: populate more fields
// TODO: Support orgs and folders.
type AuthorizationResponse struct {
	Projects []string `json:"projects"`
}

const (
	// TODO: Revaluate case sensitive
	// Client browser identifiers for tokens.
	COOKIE_SESSION_NAME = "token"
	COOKIE_XSRF_NAME    = "XSRF-TOKEN"
	HEADER_XSRF_NAME    = "X-XSRF-TOKEN"

	SESSION_EXPIRATION_MINUTES        = 5
	SESSION_REFRESH_THRESHOLD_MINUTES = 1
	XSRF_KEY                          = "my_secret_key"
	XSRF_ACTION_ID                    = "global"
)

var (
	googleIdTokenVerifier = verifier.Verifier{}
	// TODO: might make sense to be global
	UNIX_EPOCH = time.Unix(0, 0)

	CONTAINER_JWT_KEY = []byte("my_secret_key_2")
	JWT_KEY           = []byte("my_secret_key")
)

func auth_Login(w http.ResponseWriter, r *http.Request) {
	auth_sign(w, "sudo")
}

func auth_Authenticate(w http.ResponseWriter, r *http.Request) {
	// https://stackoverflow.com/questions/17478731/whats-the-point-of-the-x-requested-with-header
	if xRequestedWithHeader := r.Header.Get("X-Requested-With"); xRequestedWithHeader != "XMLHttpRequest" {
		http.Error(w, "Untrusted request", http.StatusForbidden)
		return
	}
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var auth AuthorizationStruct
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	idToken, err := auth_verifyIdToken(auth.Id_Token)
	if err != nil {
		http.Error(w, "Cannot verify id_token JWT", http.StatusForbidden)
		return
	}

	token, err := googleAuth.Exchange(oauth2.NoContext, auth.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	if token == nil {
		http.Error(w, "No token response received", http.StatusForbidden)
	}

	log.Printf("Creating session for user %s", idToken.Email)
	auth_sign(w, idToken.Email)

	response, err := project_list(token)
	if err != nil {
		// TODO: don't fail this hard
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	io.WriteString(w, string(*response))
}

func auth_Logout(w http.ResponseWriter, r *http.Request) {
	auth_unsetCookieSession(w)
	auth_unsetCookieXsrf(w)
}

func auth_Refresh(w http.ResponseWriter, r *http.Request) {
	claims, _ := auth_verify(w, r)
	if claims == nil {
		return
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// expiry threshold. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > time.Duration(SESSION_REFRESH_THRESHOLD_MINUTES)*time.Minute {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	auth_sign(w, claims.Username)
}

func auth_AuthTest(w http.ResponseWriter, r *http.Request) {
	claims, containedJwt := auth_verify(w, r)
	if claims == nil {
		return
	}
	xXsrfTokenHeader := r.Header.Get(HEADER_XSRF_NAME)
	if xXsrfTokenHeader == "" {
		http.Error(w, "Missing XSRF", http.StatusForbidden)
		return
	}

	if isValidXsrf := xsrf.Valid(xXsrfTokenHeader, XSRF_KEY, *containedJwt, XSRF_ACTION_ID); !isValidXsrf {
		http.Error(w, "Invalid XSRF", http.StatusForbidden)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}

func auth_verifyIdToken(idToken string) (*verifier.ClaimSet, error) {
	logger.Printf("Verifying id_token: " + idToken)
	if err := googleIdTokenVerifier.VerifyIDToken(idToken, idTokenAudience); err != nil {
		logger.Printf("Error verifying id_token.")
		return nil, err
	}
	claims, err := verifier.Decode(idToken)
	if err != nil {
		logger.Print("Error decoding id_token.")
	}
	return claims, err
}

func auth_signWithClaims(w http.ResponseWriter, key []byte, claims jwt.Claims) *string {
	// Create the JWT string
	jwtStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}
	return &jwtStr
}

func auth_sign(w http.ResponseWriter, username string) {
	expirationTime := time.Now().Add(time.Duration(SESSION_EXPIRATION_MINUTES) * time.Minute)
	// In JWT, the expiry time is expressed as unix milliseconds
	standardClaims := jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	// Sign inner JWT
	containedJwt := auth_signWithClaims(w, JWT_KEY, &Claims{
		Username:       username,
		StandardClaims: standardClaims,
	})
	if containedJwt == nil {
		return
	}

	// Sign outer JWT.
	containerJwt := auth_signWithClaims(w, CONTAINER_JWT_KEY, &ContainerClaims{
		ContainedJwt:   *containedJwt,
		StandardClaims: standardClaims,
	})
	if containerJwt == nil {
		return
	}

	// Set an expiry time which is the same as the token itself.
	auth_setCookieSession(w, *containerJwt, expirationTime)

	// By generating the XSRF token using the JWT, the xsrf token is valid
	// only if the JWT is valid, sidestepping limitation of net/xsrftoken library
	// having 24 hour expiration, and pose risk where if the XSRF token cookie
	// is leaked or stolen, it can only be used with the corresponding JWT and
	// none other.
	xsrfToken := xsrf.Generate(XSRF_KEY, *containedJwt, XSRF_ACTION_ID)
	// Since some time has elapsed after the time xsrfToken issued, we want the
	// cookie to expire shortly before the token does. This doesn't matter too
	// much as the xsrf-token lifespan bounded by JWT's lifespan, as long as JWT
	// is verified first, and expiration shortcircuits request.
	xsrfCookieExpiration := time.Now().Add(xsrf.Timeout).Add(time.Duration(-1 * time.Minute))
	auth_setCookieXsrf(w, xsrfToken, xsrfCookieExpiration)
}

func auth_extractClaims(w http.ResponseWriter, jwtStr string, key []byte, claims jwt.Claims) bool {
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(jwtStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	return true
}

func auth_verify(w http.ResponseWriter, r *http.Request) (*Claims, *string) {
	// Extract the session cookie.
	c, err := r.Cookie(COOKIE_SESSION_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return nil, nil
		}
		w.WriteHeader(http.StatusBadRequest)
		return nil, nil
	}

	containerClaims := &ContainerClaims{}
	if success := auth_extractClaims(w, c.Value, CONTAINER_JWT_KEY, containerClaims); !success {
		return nil, nil
	}
	containedJwt := containerClaims.ContainedJwt

	claims := &Claims{}
	if success := auth_extractClaims(w, containedJwt, JWT_KEY, claims); !success {
		return nil, nil
	}

	return claims, &containedJwt
}

func auth_unsetCookieSession(w http.ResponseWriter) {
	auth_setCookieSession(w, "", UNIX_EPOCH)
}

func auth_unsetCookieXsrf(w http.ResponseWriter) {
	auth_setCookieXsrf(w, "", UNIX_EPOCH)
}

func auth_setCookieSession(w http.ResponseWriter, value string, t time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:    COOKIE_SESSION_NAME,
		Path:    "/",
		Value:   value,
		Expires: t,
		// prevents cookie from being read by JavaScript. Cookie will still
		// be automatically attached to http requests. This has
		// nothing to do with https vs http
		HttpOnly: true,
	})
}

func auth_setCookieXsrf(w http.ResponseWriter, value string, t time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name: COOKIE_XSRF_NAME,
		// TODO: evaluate scoping this to authorized pages
		Path:    "/",
		Value:   value,
		Expires: t,
		// Allows cookie to be read by JavaScript
		HttpOnly: false,
	})
}
