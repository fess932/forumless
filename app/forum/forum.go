package forum

type Repo interface {
	GetName() string
}

type Forum struct {
	Host string
	Name string
	Port string

	repo Repo
}

func New() {

}

func (f Forum) Run() {

}

func (f Forum) GetName() {
	f.repo.GetName()
}
