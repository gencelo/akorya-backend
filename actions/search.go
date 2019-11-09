package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

type Suggest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Type string `json:"type"`
}

type SingerSearch struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Slug      string       `json:"slug"`
}

type SongSearch struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Slug      string       `json:"slug"`
}

type SearchResource struct {
	buffalo.Resource
}

// ListGetPopularSingers default implementation.
func  (v SearchResource) GetSuggestions(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	suggests := []Suggest{}

	q := tx.RawQuery("SELECT singer.id, singer.name, singer.slug, 'singer' as type "+
		" FROM singers as singer "+
		" WHERE singer.name like (?) " +
		" LIMIT 3" , c.Param("q")+"%")

	if err := q.All(&suggests); err != nil {
		return c.Error(404, err)
	}

	qSong := tx.RawQuery("SELECT song.id, CONCAT(song.name, ' - ', singer.name) as name, song.slug, 'song' as type "+
		" FROM songs as song, singers as singer  "+
		" WHERE song.name like (?) " +
		" AND singer.id = song.singer_id" +
		" LIMIT 7" , c.Param("q")+"%")

	if err := qSong.All(&suggests); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(&suggests))
}