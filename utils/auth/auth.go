package auth

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func getHexdigits(salt, rawPassword string) string {
	hash := sha1.Sum([]byte(salt + rawPassword))
	return fmt.Sprintf("%x", hash)
}

func MakePassword(rawPassword string) string {
	s1 := strconv.FormatFloat(rand.Float64(), 'E', -1, 64)
	s2 := strconv.FormatFloat(rand.Float64(), 'E', -1, 64)
	salt := getHexdigits(s1, s2)
	hash := getHexdigits(salt, rawPassword)
	return fmt.Sprintf("%s$%s", salt, hash)
}

func CheckPassword(rawPassword, encPassword string) bool {
	encSplited := strings.Split(encPassword, "$")
	salt := encSplited[0]
	hash := encSplited[1]
	return hash == getHexdigits(salt, rawPassword)
}

func EmailFixer(email string) string {
	splited := strings.Split(email, "@")
	username := splited[0]
	domain := splited[1]
	username = strings.Replace(username, ".", "", -1)
	return username + "@" + domain
}
