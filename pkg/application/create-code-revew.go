package application

import (
	"math/rand"
)

var comments = []string{"shark", "jellyfish", "squid", "octopus", "dolphin"}

func GetRandomComments() string {
	i := rand.Intn(len(comments))
	return comments[i]
}
