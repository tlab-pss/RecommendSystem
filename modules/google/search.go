package google

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Search(query string) (*[]Page, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	engineID := os.Getenv("CUSTOM_SEARCH_ENGINE_ID")

	url := "https://www.googleapis.com/customsearch/v1?cx=" + engineID + "&key=" + apiKey + "&q=" + query

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "www.googleapis.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	var r io.Reader = res.Body

	cr := CustomSearchResponse{}

	if err := json.NewDecoder(r).Decode(&cr); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}

	fmt.Println(cr.Items)
	p := []Page{}

	for _, item := range cr.Items {
		var thumbnail string
		if len(item.Pagemap.CseThumbnail) == 0 {
			thumbnail = ""
		} else {
			thumbnail = item.Pagemap.CseThumbnail[0].Src
		}

		p = append(p, Page{
			URL:         item.Link,
			Title:       item.Title,
			Description: item.Snippet,
			Thumbnail:   thumbnail,
		})
	}

	return &p, nil
}

type Page struct {
	URL         string
	Title       string
	Description string
	Thumbnail   string
}

type CustomSearchResponse struct {
	Kind string `json:"kind"`
	URL  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		Request []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"request"`
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"nextPage"`
	} `json:"queries"`
	Context struct {
		Title string `json:"title"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []struct {
		Kind             string `json:"kind"`
		Title            string `json:"title"`
		HTMLTitle        string `json:"htmlTitle"`
		Link             string `json:"link"`
		DisplayLink      string `json:"displayLink"`
		Snippet          string `json:"snippet"`
		HTMLSnippet      string `json:"htmlSnippet"`
		CacheID          string `json:"cacheId,omitempty"`
		FormattedURL     string `json:"formattedUrl"`
		HTMLFormattedURL string `json:"htmlFormattedUrl"`
		Pagemap          struct {
			CseThumbnail []struct {
				Width  string `json:"width"`
				Height string `json:"height"`
				Src    string `json:"src"`
			} `json:"cse_thumbnail"`
			Metatags []struct {
				Referrer           string `json:"referrer"`
				OgImage            string `json:"og:image"`
				FormatDetection    string `json:"format-detection"`
				Viewport           string `json:"viewport"`
				OgLocale           string `json:"og:locale"`
				OgType             string `json:"og:type"`
				OgTitle            string `json:"og:title"`
				OgDescription      string `json:"og:description"`
				OgSiteName         string `json:"og:site_name"`
				OgURL              string `json:"og:url"`
				Author             string `json:"author"`
				Copyright          string `json:"copyright"`
				FbAppID            string `json:"fb:app_id"`
				TwitterSite        string `json:"twitter:site"`
				TwitterTitle       string `json:"twitter:title"`
				TwitterDescription string `json:"twitter:description"`
				TwitterCard        string `json:"twitter:card"`
				TwitterImage       string `json:"twitter:image"`
				AlIosAppStoreID    string `json:"al:ios:app_store_id"`
				AlIosAppName       string `json:"al:ios:app_name"`
				AlIosURL           string `json:"al:ios:url"`
				AlAndroidURL       string `json:"al:android:url"`
				AlAndroidAppName   string `json:"al:android:app_name"`
				AlAndroidPackage   string `json:"al:android:package"`
				AlWebURL           string `json:"al:web:url"`
			} `json:"metatags"`
			Breadcrumb []struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"breadcrumb"`
			CseImage []struct {
				Src string `json:"src"`
			} `json:"cse_image"`
			Videoobject []struct {
				URL              string `json:"url"`
				Name             string `json:"name"`
				Description      string `json:"description"`
				Paid             string `json:"paid"`
				Channelid        string `json:"channelid"`
				Videoid          string `json:"videoid"`
				Duration         string `json:"duration"`
				Unlisted         string `json:"unlisted"`
				Thumbnailurl     string `json:"thumbnailurl"`
				Embedurl         string `json:"embedurl"`
				Playertype       string `json:"playertype"`
				Width            string `json:"width"`
				Height           string `json:"height"`
				Isfamilyfriendly string `json:"isfamilyfriendly"`
				Regionsallowed   string `json:"regionsallowed"`
				Interactioncount string `json:"interactioncount"`
				Datepublished    string `json:"datepublished"`
				Uploaddate       string `json:"uploaddate"`
				Genre            string `json:"genre"`
			} `json:"videoobject"`
			Imageobject []struct {
				URL    string `json:"url"`
				Width  string `json:"width"`
				Height string `json:"height"`
			} `json:"imageobject"`
			Person []struct {
				URL string `json:"url"`
			} `json:"person"`
		} `json:"pagemap"`
	} `json:"items"`
}
