package seeds

import "github.com/bxcodec/faker/v3"

func (s Seed) JobSeed() {
	isPublished := "true"
	for i := 0; i < 100; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO available_jobs(title,description,apply_until,location_id,job_category_id,created_by_id,created_at,is_published) VALUES (?,?,?,?,?,?,?,?)`)
		// execute query

		_, err := stmt.Exec(faker.Word(), faker.Paragraph(), faker.Date(), 1, 1, 1, faker.Date(), isPublished)
		if err != nil {
			panic(err)
		}
	}
}
