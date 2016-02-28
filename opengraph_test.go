package opengraph

import (
	"net/http"
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

var tt3 = `<head>
<title>The Rock (1996)</title>
<meta property="og:title" content="The Rock">
<meta property="fb:app_id" content="115109575169727">
<meta property="og:type" content="video.movie">
<meta property="og:url" content="http://www.imdb.com/title/tt0117500/">
<meta property="og:image" content="http://ia.media-imdb.com/images/rock.jpg">
</head>`

var tt4 = `
<head>
<title>Apple Watch Hermès</title>
<meta name="og:title" content="Découvrez la Apple Watch Hermès.">
<meta name="og:site_name" content="Apple Watch Hermès">
<meta name="" content="website">
<meta property="og:locale" name="og:country" content="fr">
<meta name="og:url" content="http://hermes.com/applewatchhermes">
<meta itemprop="name" content="Découvrez la Apple Watch Hermès.">`

func TestExport(t *testing.T) {

	var ogTests = []struct {
		doc      string
		prefix   string
		expected []MetaData
	}{
		{
			"",
			"og",
			[]MetaData{},
		},
		{
			tt1,
			"og",
			[]MetaData{
				{"title", "The Rock", "og"},
				{"type", "video.movie", "og"},
				{"url", "http://www.imdb.com/title/tt0117500/", "og"},
				{"image", "http://ia.media-imdb.com/images/rock.jpg", "og"},
			},
		},
		{
			tt1,
			"fb",
			[]MetaData{
				{"app_id", "115109575169727", "fb"},
			},
		},
		{
			tt1,
			"",
			[]MetaData{
				{"title", "The Rock", "og"},
				{"app_id", "115109575169727", "fb"},
				{"type", "video.movie", "og"},
				{"url", "http://www.imdb.com/title/tt0117500/", "og"},
				{"image", "http://ia.media-imdb.com/images/rock.jpg", "og"},
			},
		},
		{
			tt2,
			"og",
			[]MetaData{
				{"title", "The Rock", "og"},
				{"url", "http://www.imdb.com/title/tt0117500/", "og"},
			},
		},
		{
			tt3,
			"og",
			[]MetaData{
				{"title", "The Rock", "og"},
				{"type", "video.movie", "og"},
				{"url", "http://www.imdb.com/title/tt0117500/", "og"},
				{"image", "http://ia.media-imdb.com/images/rock.jpg", "og"},
			},
		},
		{
			tt4,
			"og",
			[]MetaData{
				{"title", "Découvrez la Apple Watch Hermès.", "og"},
				{"site_name", "Apple Watch Hermès", "og"},
				{"locale", "fr", "og"},
				{"url", "http://hermes.com/applewatchhermes", "og"},
			},
		},
	}

	for ti, tt := range ogTests {
		og, err := ExtractPrefix(strings.NewReader(tt.doc), tt.prefix)
		if err != nil {
			t.Errorf("Test %d err: %s", ti, err.Error())
		}

		if len(og) != len(tt.expected) {
			t.Fatalf("Test %d got: %+v, expected: %+v", ti, og, tt.expected)
		}

		for i, _ := range og {
			if og[i].Property != tt.expected[i].Property || og[i].Content != tt.expected[i].Content {
				t.Errorf("Test %d got: %+v, expected: %+v", ti, og[i], tt.expected[i])
			}
		}
	}
}

func BenchmarkExport(b *testing.B) {
	b.StopTimer()
	rdr := strings.NewReader(tt1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Extract(rdr)
	}
}

func BenchmarkExport2(b *testing.B) {
	b.StopTimer()
	r, err := http.Get("http://www.imdb.com/title/tt0118715/")
	if err != nil {
		b.Skipf("skipping, net/http err: %s", err.Error())
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Extract(r.Body)
	}
}
