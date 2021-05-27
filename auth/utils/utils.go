package utils

import (
	"fmt"
	"github.com/supperdoggy/score/auth/hiddenConst"
	"github.com/supperdoggy/score/sctructs"
	authdata "github.com/supperdoggy/score/sctructs/service/auth"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomString - generates random string with given length
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	if length < 0 {
		return ""
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CheckCredentials(req authdata.RegisterReq) error {
	var err error
	if req.Username == "" || req.Email == "" || req.Password == "" || req.Name == "" {
		err = fmt.Errorf("fill all of the fields")
	}
	// check for @ in username and email validation
	if strings.Contains(req.Username, "@") {
		err = fmt.Errorf("username cant contain @")
	}
	if match, _ := regexp.MatchString(sctructs.RegexpEmail, req.Email); !match {
		err = fmt.Errorf("looks like given email is invalid")
	}
	if len(req.Username) < 4 {
		err = fmt.Errorf("username should be more than 4 chars")
	}
	if len(req.Password) < 8 {
		err = fmt.Errorf("password should be at least 8 chars")
	}
	return err
}

// HashAndSalt - returns hashed password
func HashAndSalt(pwd string) (string, error) {
	pwdByte := []byte(pwd + hiddenConst.Salt)
	hash, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
