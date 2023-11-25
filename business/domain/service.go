package domain

//abstraksi untuk content nya
// disini isinya ada repository dan service

//ingoing port : handle teknologi yang ada di sarver
type Repository interface {
	FindContentbyID(id int) (content *Content, err error)
	FindContent() (contents []Content, err error)
	InsertContent(content Content) (err error)
	UpdateContent(content Content, currentVersion int) (err error)
}

//outgoing port :
type Service interface {
	GetContentbyID(id int) (content *Content, err error)
	GetContent() (contents []Content, err error)
	CreateContent(content Content) (err error)
}
