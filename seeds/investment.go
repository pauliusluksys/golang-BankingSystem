package seeds

import (
	"github.com/bxcodec/faker/v3"
)

func (s Seed) InvestmentSeed() {
	for i := 0; i < 1000; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO investments(created_at,title,category_investment_id,company_investment_id,risk_level_investment_id) VALUES (?,?,?,?,?)`)
		// execute query
		runes := []rune(faker.Name())
		_, err := stmt.Exec(faker.Date(), string(runes[1:19]), 1, 1, 1)
		if err != nil {
			panic(err)
		}
	}
}
