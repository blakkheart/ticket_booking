package repository

type BaseRepository struct {
	//DB *sql.DB
}

func (repo *BaseRepository) Create(user Account) (Account, error) {
	q := New(DB.Conn)
	account, err := q.CreateAccount(DB.Ctx,
		CreateAccountParams{
			Name:  user.Name,
			Email: user.Email,
		})

	return account, err
}

func (repo *BaseRepository) Delete(id int) (int, error) {
	var count int = 2
	return count, nil
}
func (repo *BaseRepository) Get(id int) (int, error) {
	var count int = 3
	return count, nil
}
func (repo *BaseRepository) GetMany(filter any) {

}

func (repo *BaseRepository) CreateEvent(event_params CreateEventParams) (Event, error) {
	q := New(DB.Conn)
	event, err := q.CreateEvent(
		DB.Ctx,
		event_params,
	)
	return event, err
}

func (repo *BaseRepository) GetEvent(id int64) (Event, error) {
	q := New(DB.Conn)
	event, err := q.GetEvent(
		DB.Ctx,
		id,
	)
	return event, err
}
