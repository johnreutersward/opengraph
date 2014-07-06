package opengraph

import (
	"strings"
	"testing"
)

var tt1 = `<head>
<title>The Rock (1996)</title>
<meta property="og:title" content="The Rock" />
<meta property="fb:app_id" content="115109575169727" />
<meta property="og:type" content="video.movie" />
<meta property="og:url" content="http://www.imdb.com/title/tt0117500/" />
<meta property="og:image" content="http://ia.media-imdb.com/images/rock.jpg" />
</head>`

var tt2 = `<head>
<title>The Rock (1996)</title>
<meta property="og:title" content="The Rock" />
<meta property="fb:app_id" content="115109575169727" />
<meta property="" content="video.movie" />
<meta property="og:url" content="http://www.imdb.com/title/tt0117500/" />
<meta property="og:image" content="" />
</head>`

func TestExport(t *testing.T) {

	var ogTests = []struct {
		doc      string
		ns       string
		expected []MetaData
	}{
		{
			"",
			"og:",
			[]MetaData{},
		},
		{
			tt1,
			"og:",
			[]MetaData{
				{"title", "The Rock"},
				{"type", "video.movie"},
				{"url", "http://www.imdb.com/title/tt0117500/"},
				{"image", "http://ia.media-imdb.com/images/rock.jpg"},
			},
		},
		{
			tt1,
			"fb:",
			[]MetaData{
				{"app_id", "115109575169727"},
			},
		},
		{
			tt1,
			"",
			[]MetaData{
				{"og:title", "The Rock"},
				{"fb:app_id", "115109575169727"},
				{"og:type", "video.movie"},
				{"og:url", "http://www.imdb.com/title/tt0117500/"},
				{"og:image", "http://ia.media-imdb.com/images/rock.jpg"},
			},
		},
		{
			tt2,
			"og:",
			[]MetaData{
				{"title", "The Rock"},
				{"url", "http://www.imdb.com/title/tt0117500/"},
			},
		},
	}

	for _, tt := range ogTests {
		Namespace = tt.ns
		og, err := Extract(strings.NewReader(tt.doc))

		if err != nil {
			t.Errorf("err: %s", err.Error())
		}

		if len(og) != len(tt.expected) {
			t.Fatalf("got: %+v, expected: %+v", og, tt.expected)
		}

		for i, _ := range og {
			if og[i].Property != tt.expected[i].Property || og[i].Content != tt.expected[i].Content {
				t.Errorf("got: %+v, expected: %+v", og[i], tt.expected[i])
			}
		}
	}
}

func BenchmarkExport(b *testing.B) {
	rdr := strings.NewReader(tt1)
	for i := 0; i < b.N; i++ {
		Extract(rdr)
	}
}
