package sumcomp

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	buzzwords  = []string{"Disruption", "Leadership", "Blockchain", "Team", "Win", "Sales", "Excellence", "Greatness"}
	firstNames = []string{"Alice", "Bob", "Charlene", "Dan", "Eve", "Fin", "Gianna", "Henry"}
	lastNames  = []string{"Allison", "Bobson", "Charleton", "Dannings", "Evans", "Forry", "Geldof", "Henneman"}
)

type Summary struct {
	DataId int
	Title  string
	Author string
}

func RandomSummary() Summary {
	dataId := rand.Intn(9000) + 1000
	title := strings.Join(Pick(2, buzzwords), " of ")
	author := strings.Join([]string{Pick(1, firstNames)[0], Pick(1, lastNames)[0]}, " ")
	return Summary{DataId: dataId, Title: title, Author: author}
}

func (s Summary) String() string {
	return fmt.Sprintf("%d: «%s» by %s", s.DataId, s.Title, s.Author)
}
