package data

var d *Data

func init() {
	db, err := New("../test_data/data.db")
	if err != nil {
		panic(err)
	}
	d = db
}
