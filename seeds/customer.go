package seeds

import (
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

func (s Seed) CustomerSeed() {
	for i := 0; i < 1000; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO customers(name,date_of_birth,city,zipcode,status) VALUES (?,?,?,?,?)`)
		// execute query
		runes := []rune(faker.Name())
		zipcode := rand.Intn(10000)
		_, err := stmt.Exec(string(runes[1:19]), faker.BaseDateFormat, faker.Word(), zipcode, 1)
		if err != nil {
			panic(err)
		}
	}
}
