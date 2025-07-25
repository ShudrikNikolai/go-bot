package usecase

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

func YandexDisk(href string, msgId string) string {
	u := os.Getenv("YANDEX_CLOUD_URL")

	params := url.Values{}
	params.Set("path", os.Getenv("YANDEX_BACKET")+msgId)
	params.Set("url", href)

	logrus.Print(params.Encode())
	req, err := http.NewRequest("POST", u+"?"+params.Encode(), nil)
	if err != nil {
		logrus.Warning("Error creating request:", err)
	}

	req.Header.Set("Authorization", "OAuth "+os.Getenv("YANDEX_TOKEN"))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Warning("Error making request:", err)
		return "Error making request"
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Warning("Unexpected status code:", resp.StatusCode)
		return "Unexpected status code:" + strconv.Itoa(resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Warning("Error reading response:", err)
		return "Error reading response"
	}

	// Print response
	logrus.Printf("Response: %s", body)

	return "ok"
}
