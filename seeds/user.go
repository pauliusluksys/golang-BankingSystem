package seeds

import (
	"github.com/bxcodec/faker/v3"
)

func (s Seed) UserSeed() {
	userRole := "user"
	var customerId = 0
	var password = "pass"
	for i := 0; i < 100; i++ {
		customerId++
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO users(username, password,role,customer_id,created_on) VALUES (?,?,?,?,?)`)
		// execute query
		runes := []rune(faker.Name())
		_, err := stmt.Exec(string(runes[1:19]), password, userRole, customerId, faker.BaseDateFormat)
		if err != nil {
			panic(err)
		}
	}
}
