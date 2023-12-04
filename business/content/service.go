// ini adalah bagian service
package content

//abstraksi untuk content nya
// disini isinya ada repository dan service

//ingoing port : handle teknologi yang ada di sarver
type Repository interface { // handle dari sisi sarver/dalem nya
	FindContentbyID(id int) (content *Content, err error)
	FindAll() (contents []Content, err error)
	InsertContent(content Content) (err error)
	UpdateContent(content Content, currentVersion int) (err error)
}

//outgoing port :
type Service interface { // utk handle dari sisi client
	GetContentbyID(id int) (content *Content, err error)
	GetContent() (contents []Content, err error)
	CreateContent(content Content) (err error)
	UpdateContent(content Content, currentVersion int) (err error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

//fungsi ini untuk handle sisi sarver ke database nanti
func (s *service) GetContentbyID(id int) (content *Content, err error) {
	result, err := s.repository.FindContentbyID(id)
	return result, err
}

func (s *service) GetContent() (contents []Content, err error) {
	contents, err = s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (s *service) CreateContent(content Content) (err error) {
	return
}

func (s *service) UpdateContent(content Content, currentVersion int) (err error) {
	return
}
