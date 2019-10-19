package interactors

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

// base64エンコードされた画像を引数に渡してね！
func Discriminate(i string) error {

	v, err := requestVisionAI(i)

	if err != nil {
		return err
	}

	var labels []string
	for _, l := range v.Responses[0].LabelAnnotations {
		labels = append(labels, l.Description)
	}

	if contains(labels, []string{"Food", "Dish"}) {
		var labels []string // 80%以上のlabel
		for _, l := range v.Responses[0].LabelAnnotations {
			if l.Score > 0.8 {
				labels = append(labels, l.Description)
			}
		}

		entity := v.Responses[0].WebDetection.WebEntities[0].Description // 最高ポイントのentity
		if err := storeFood(labels, entity, i); err != nil {
			return err
		}
	} else {
		// todo: 現状食べ物以外はない
		return nil
	}

	return nil
}

func requestVisionAI(i string) (*visionAIResponse, error) {
	var v = visionAIResponse{}

	bByte, _ := json.Marshal(map[string]interface{}{
		"requests": map[string]interface{}{
			"image": map[string]interface{}{
				"content": i,
			},
			"features": []map[string]interface{}{
				{
					"type":       "LABEL_DETECTION",
					"maxResults": 5,
				},
				{
					"type":       "WEB_DETECTION",
					"maxResults": 5,
				},
			},
		},
	})

	req, err := http.NewRequest("POST", "https://vision.googleapis.com/v1/images:annotate?key="+os.Getenv("GOOGLE_API_KEY"), bytes.NewReader(bByte))

	if err != nil {
		return &v, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return &v, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return &v, err
	}

	return &v, nil
}

func storeFood(l []string, e string, i string) error {
	bByte, _ := json.Marshal(map[string]interface{}{
		"menu":              e,
		"calorie":           0.0, // todo: カロリーどうやって求めようかな
		"photo":             i,
		"labels":            l,
		"small_category_id": 1, // todo: 小カテゴリIDの紐付けめんど...
	})

	req, err := http.NewRequest("POST", "pd/api/hotpepper/intake", bytes.NewReader(bByte))

	if err != nil {
		return err
	}

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

type visionAIResponse struct {
	Responses []struct {
		LabelAnnotations []struct {
			Mid         string  `json:"mid"`
			Description string  `json:"description"`
			Score       float64 `json:"score"`
			Topicality  float64 `json:"topicality"`
		} `json:"labelAnnotations"`
		WebDetection struct {
			WebEntities []struct {
				EntityID    string  `json:"entityId"`
				Score       float64 `json:"score"`
				Description string  `json:"description"`
			} `json:"webEntities"`
			FullMatchingImages []struct {
				URL string `json:"url"`
			} `json:"fullMatchingImages"`
			PartialMatchingImages []struct {
				URL string `json:"url"`
			} `json:"partialMatchingImages"`
			PagesWithMatchingImages []struct {
				URL                   string `json:"url"`
				PageTitle             string `json:"pageTitle"`
				PartialMatchingImages []struct {
					URL string `json:"url"`
				} `json:"partialMatchingImages,omitempty"`
				FullMatchingImages []struct {
					URL string `json:"url"`
				} `json:"fullMatchingImages,omitempty"`
			} `json:"pagesWithMatchingImages"`
			VisuallySimilarImages []struct {
				URL string `json:"url"`
			} `json:"visuallySimilarImages"`
			BestGuessLabels []struct {
				Label        string `json:"label"`
				LanguageCode string `json:"languageCode"`
			} `json:"bestGuessLabels"`
		} `json:"webDetection"`
	} `json:"responses"`
}

// todo: 全体のutil作ってそこに移したい
func contains(s []string, n []string) bool {
	for _, e := range s {
		for _, v := range n {
			if e == v {
				return true
			}
		}
	}
	return false
}
