package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

type SingerSong struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Singer struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Slug      string       `json:"slug"`
	SongCount int          `json:"songCount"`
	Songs     []SingerSong `json:"songs"`

}

type SingersResource struct {
	buffalo.Resource
}

func (v SingersResource) List(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	singers := []Singer{}

	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	q = q.RawQuery("SELECT singer.id, singer.name, singer.slug from singers as singer")

	if err := q.All(singers); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.JSON(&singers))
}

// the path GET /singers/{singer_id}
func (v SingersResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	singer := &Singer{}
	songs := []SingerSong{}

	q := tx.RawQuery("SELECT singer.id, singer.name, singer.slug FROM singers as singer  WHERE singer.slug = ? ", c.Param("slug"))

	if err := q.Find(singer, c.Param("slug")); err != nil {
		return c.Error(404, err)
	}

	//qSong := tx.PaginateFromParams(c.Params())
	qSong := tx.RawQuery("SELECT song.id as id, song.name as name, song.slug as slug FROM songs as song  WHERE song.singer_id = ?  ", singer.Id)
	// To find the Singer the parameter singer_id is used.
	if err := qSong.All(&songs); err != nil {
		return c.Error(404, err)
	}

	singer.Songs = songs

	q = q.RawQuery("SELECT COUNT(id) FROM songs where singer_id = ?", singer.Id)
	if err := q.Find(&singer.SongCount, singer.Id); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(&singer))
}
