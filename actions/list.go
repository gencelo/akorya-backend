package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

type ListSong struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	SingerName string `json:"singerName"`
	SingerSlug string `json:"singerSlug"`
}

type ListSinger struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	SongCount int    `json:"-"`
}

// ListGetPopularSongs default implementation.
func ListGetPopularSongs(c buffalo.Context) error {
	ids := []int{4265, 772, 5201, 4280, 4290, 10, 388, 259, 4859, 4569, 1407, 5536, 5106, 4234}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	songs := []ShortSong{}

	q := tx.RawQuery("SELECT song.id, song.name, song.slug, singer.name as singername, singer.slug as singerslug"+
		" FROM songs as song, singers as singer"+
		" WHERE song.singer_id = singer.id AND song.id in (?)" +
		" ORDER by FIELD(song.id, ?)", ids, ids)

	err := q.All(&songs)
	if err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(&songs))
}

// ListGetPopularSingers default implementation.
func ListGetPopularSingers(c buffalo.Context) error {
	ids := []int{21, 43, 123, 47, 52, 15, 92, 153, 53, 112}
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	singers := []ListSinger{}

	q := tx.RawQuery("SELECT singer.id, singer.name, singer.slug "+
		" FROM singers as singer "+
		" WHERE singer.id in (?) " +
		" ORDER by FIELD(singer.id, ?)", ids, ids)

	if err := q.All(&singers); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(&singers))
}

// ListGetRisingSongs default implementation.
func ListGetRisingSongs(c buffalo.Context) error {
	ids := []int{5182, 44, 3645, 5060, 114, 5387, 1126, 5450, 357, 579, 4274, 915, 68, 48, 4997, 5600, 462, 830, 916, 4809, 5386, 1330, 1911, 5387}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	songs := []ShortSong{}

	q := tx.RawQuery("SELECT song.id, song.name, song.slug, singer.name as singername, singer.slug as singerslug"+
		" FROM songs as song, singers as singer"+
		" WHERE song.singer_id = singer.id AND song.id in (?)" +
		" ORDER by FIELD(song.id, ?)", ids, ids)

	err := q.All(&songs)
	if err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(&songs))

} // ListGetClassicSongs default implementation.
func ListGetClassicSongs(c buffalo.Context) error {
	ids := []int{1407, 21, 571, 4302, 5274, 408, 22, 5519, 138, 13}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	songs := []ShortSong{}

	q := tx.RawQuery("SELECT song.id, song.name, song.slug, singer.name as singername, singer.slug as singerslug"+
		" FROM songs as song, singers as singer"+
		" WHERE song.singer_id = singer.id AND song.id in (?)" +
		" ORDER by FIELD(song.id, ?)", ids, ids)

	err := q.All(&songs)
	if err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(&songs))
}
