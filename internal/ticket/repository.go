package ticket

/* type ITicketRepository interface {
	CreateTicket(title string, description string, due_date time.Time) (db.Ticket, error)
	UpdateTicket(title string, description string, due_date time.Time) (db.Ticket, error)
	GetTickets(offset int, page_size int) ([]db.Ticket, error)
	DeleteTicket(ID uint) (db.Ticket, error)
}

type GORMTicketRepository struct {
	db *gorm.DB
}

func CreateNewGORMTicketRepository(db *gorm.DB) ITicketRepository {
	return &GORMTicketRepository{
		db: db,
	}
}

func (r *GORMTicketRepository) CreateTicket(title string, description string, due_date time.Time) (db.Ticket, error) {
	ticket := db.Ticket{
		Title:       title,
		Description: description,
		DueDate:     due_date,
	}

	result := r.db.Create(&ticket)
	if result.Error != nil {
		return db.Ticket{}, result.Error
	}

	return ticket, nil
}

func (r *GORMTicketRepository) UpdateTicket(ID uint, title string, description string, assigned_to uint, due_date time.Time) (db.Ticket, error) {

	ticket := db.Ticket{
		ID:          ID,
		Title:       title,
		Description: description,
		DueDate:     due_date,
	}

	result := r.db.Model(&ticket).Updates(&ticket)
	if result.Error != nil {
		return db.Ticket{}, result.Error
	}

	return ticket, nil
}

func (r *GORMTicketRepository) GetTickets(offset int, page_size int) ([]db.Ticket, error) {

}

func (r *GORMTicketRepository) DeleteTicket(ID uint) (db.Ticket, error)                   {}
*/
