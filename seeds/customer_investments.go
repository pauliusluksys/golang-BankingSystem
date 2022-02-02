package seeds

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"time"
)

func (s Seed) CustomerInvestmentSeed() {

	var customerIds []int
	var investmentIds []int
	err := s.db.Select(&customerIds, "select customer_id from customers")
	if err != nil {
		panic("error whilst obtaining customer ids: " + err.Error())
	}
	err = s.db.Select(&investmentIds, "select id from investments")
	if err != nil {
		panic("error whilst obtaining investment ids: " + err.Error())
	}
	var amountRandom int
	customerSliceLength := len(customerIds)
	investmentSliceLength := len(investmentIds)
	for i := 0; i < 10000; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO customer_investment(customer_id,investment_id,invested_amount,withdrawn_state,created_at) VALUES (?,?,?,?,?)`)
		// execute query
		rand.Seed(time.Now().UnixNano())
		min := 1000
		max := 1000000
		amountRandom = rand.Intn(max-min+1) + min
		customerIdRand := customerIds[rand.Intn(customerSliceLength-1)]
		rand.Seed(time.Now().UnixNano())
		investmentIdRand := investmentIds[rand.Intn(investmentSliceLength-1)]
		//runes := []rune(faker.Name())
		fmt.Println("amount number: ", amountRandom, "    customer id: ", customerIdRand, "     investment id:", investmentIdRand)
		_, err := stmt.Exec(customerIdRand, investmentIdRand, amountRandom, "false", faker.Date())
		if err != nil {
			panic(err)
		}
	}
}
