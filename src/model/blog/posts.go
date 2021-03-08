package blog

type Posts struct {
	Id          int64    `bson: "id" json:"blogId"`
	Title       string   `bson: "title" json:"title"`
	Description string   `bson: "description" json:"description"`
	Contents    []string `bson: "contents" json:"contents"`
	Codes       []string `bson: "codes" json:"codes"`
	CreatedOn   string   `bson: "createdOn" json: "createdOn"`
}

func NewDefault() Posts {
	return Posts{}
}

func New(_id int64, title string, description string, contents []string, codes []string, createdOn string) *Posts {
	return &Posts{_id, title, description, contents, codes, createdOn}
}

func (p *Posts) GetTitle() string {
	return p.Title
}

func (p *Posts) GetDescription() string {
	return p.Description
}

func (p *Posts) GetContents() []string {
	return p.Contents
}

func (p *Posts) GetCodes() []string {
	return p.Codes
}

func (p *Posts) GetId() int64 {
	return p.Id
}

func (p *Posts) GetCreatedOn() string {
	return p.CreatedOn
}

func (p *Posts) SetTitle(title string) {
	p.Title = title
}

func (p *Posts) SetDescription(description string) {
	p.Description = description
}

func (p *Posts) SetContents(contents []string) {
	p.Contents = contents
}

func (p *Posts) SetCodes(codes []string) {
	p.Codes = codes
}

func (p *Posts) SetId(id int64) {
	p.Id = id
}

func (p *Posts) SetCreatedOn(createdOn string) {
	p.CreatedOn = createdOn
}
