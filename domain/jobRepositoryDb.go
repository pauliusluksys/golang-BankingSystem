package domain

import (
	"bankingV2/errs"
	"bankingV2/utils"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type JobRepositoryDb struct {
	Client *sqlx.DB
}

func (d JobRepositoryDb) ById(Id string) (*Job, *errs.AppError) {

	var job Job
	findByIdSql := "select DISTINCT aj.id, aj.title, aj.description, aj.apply_until, l.name as location_name,jc.title as category_title from ((available_jobs aj inner join locations l on aj.location_id = l.id)inner join job_categories jc ON aj.job_category_id = jc.id) where aj.id=?;"

	err := d.Client.Get(&job, findByIdSql, Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("job not found")
		} else {
			log.Println("Error while querying job table" + err.Error())
			return nil, errs.NewUnexpectError("Unexpected database error")
		}
	}
	return &job, nil
}
func (d JobRepositoryDb) FindAll(filterMap map[string]string) ([]Job, *errs.AppError) {
	jobs := make([]Job, 0)
	var findAllSql string
	length := len(filterMap)
	var keys []string
	filterMap = filterMapToDbReadableValues(filterMap)
	keys = utils.GetKeys(filterMap)
	var err error
	if length == 0 {
		findAllSql = "select DISTINCT aj.id, aj.title  , aj.apply_until , l.name as location_name,jc.title as category_title from ((available_jobs aj inner join locations l on aj.location_id = l.id)inner join job_categories jc ON aj.job_category_id = jc.id) , banking.available_jobs ;"
		err = d.Client.Select(&jobs, findAllSql)
	} else if length == 1 {
		findAllSql = "select DISTINCT aj.id, aj.title , aj.description , aj.apply_until , l.name as location_name,jc.title as category_title from ((available_jobs aj inner join locations l on aj.location_id = l.id)inner join job_categories jc ON aj.job_category_id = jc.id)  where " + keys[0] + " = ?;"
		fmt.Println(findAllSql)
		fmt.Println(filterMap[keys[0]])
		err = d.Client.Select(&jobs, findAllSql, filterMap[keys[0]])
	} else if length == 2 {
		findAllSql = "select DISTINCT aj.id, aj.title , aj.description , aj.apply_until , l.name as location_name,jc.title as category_title from ((available_jobs aj inner join locations l on aj.location_id = l.id)inner join job_categories jc ON aj.job_category_id = jc.id)  where " + keys[0] + " = ? and " + keys[1] + ";"
		err = d.Client.Select(&jobs, findAllSql, filterMap[keys[0]], filterMap[keys[1]])
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("status not found")
		} else {
			log.Println("Error while querying customer table" + err.Error())
			return nil, errs.NewNotFoundError("Unexpected database error")
		}
	}

	return jobs, nil

}
func NewJobRepositoryDb(dbClient *sqlx.DB) JobRepositoryDb {
	return JobRepositoryDb{dbClient}
}
