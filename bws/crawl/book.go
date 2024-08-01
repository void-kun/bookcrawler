package crawl

type BookState int

const (
	CONTINUE BookState = 0
	PENDING  BookState = 1
	COMPLETE BookState = 2
)

type Author struct {
	Name     string `json:"name"`
	NameSlug string `json:"name_slug"`
}

func (c *Author) New() {}

type BookAttr struct {
	NORead     int32 `json:"no_read"`
	NOVote     int32 `json:"no_vote"`
	NOBookmark int32 `json:"no_bookmark"`
}

func (c *BookAttr) New() {}

type Comment struct {
	Owner         string   `json:"owner"`
	Content       string   `json:"content"`
	NOLike        int32    `json:"no_like"`
	NestedComment *Comment `json:"nested_comment"`
}

func (c *Comment) New() {}

type Book struct {
	Name       string     `json:"name"`
	URL        string     `json:"url"`
	State      BookState  `json:"state"`
	Categories []Category `json:"categories"`
	AvartarURL string     `json:"avartar_url"`
	Author     Author     `json:"author"`

	// status info
	Summary  string   `json:"summary"`
	BookAttr BookAttr `json:"book_attr"`

	// comments
	Comments []Comment `json:"comments"`
}

func (b *Book) New() {}
