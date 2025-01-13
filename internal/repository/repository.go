package repository

type PingRepository struct {
}

func NewPing() *PingRepository {
	return &PingRepository{}
}

func (r *PingRepository) Ping() error {

}
