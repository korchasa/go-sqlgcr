package sqlgcr

type Result struct {

}

func (r *Result) LastInsertId() (int64, error) {
	panic("implement LastInsertId")
}

func (r *Result) RowsAffected() (int64, error) {
	panic("implement RowsAffected")
}

