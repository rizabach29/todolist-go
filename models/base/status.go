package base

type Status struct {
	TableName struct{} `pg:"status"`
	Id        int      `pg:"id,pk"`
	Name      string   `pg:"name"`
}