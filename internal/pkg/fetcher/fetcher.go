package fetcher

import (
	"context"
	"errors"
	"fmt"
	"github/IAD/zacks/internal/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func NewFetcher(timeout time.Duration) *Fetcher {
	return &Fetcher{
		timeout: timeout,
	}
}

type Fetcher struct {
	timeout time.Duration
}

func (f *Fetcher) GetRating(ctx context.Context, ticker string) (*models.Rating, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()

	return f.fetch(ctxWithTimeout, ticker)
}

func (zf *Fetcher) fetch(ctx context.Context, ticker string) (*models.Rating, error) {
	url := fmt.Sprintf("http://zrsa-dev.zacks.com/stock/quote/%s", ticker)

	req, _ := http.NewRequest("GET", url, nil)

	req.WithContext(ctx)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Referer", url)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86,M_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/76.0.3809.100 Chrome/76.0.3809.100 Safari/537.36")
	req.Header.Add("Sec-Fetch-Mode", "cors")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data, err := parse(string(body))
	if err != nil {
		if err.Error() != "TickerNotExists" {
			return nil, nil
		}
		return nil, err
	}

	if data != nil {
		data.Ticker = ticker
	}

	return data, nil
}

func parse(html string) (*models.Rating, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	log.Println(doc.Find(".gs-snippet").Text())
	if strings.Contains(doc.Find("title").Text(), "Zacks Search Results") {
		return nil, errors.New("TickerNotExists")
	}

	name := strings.TrimSpace(doc.Find("#quote_ribbon_v2 > header a").Text())

	rankText := strings.TrimSpace(doc.Find(".rank_container_right .zr_rankbox").Before(".rank_chip").Text())
	rank := int64(0)
	ranks := strings.Split(rankText, "-")
	if len(ranks) == 2 {
		rank, err = strconv.ParseInt(ranks[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Can't parse rank amount. Value: %s. Err: %s", ranks[0], err.Error())
		}
	}

	scoreValueText := ""
	scoreValue := int64(0)
	scoreGrowthText := ""
	scoreGrowth := int64(0)
	scoreMomentumText := ""
	scoreMomentum := int64(0)
	scoreVGMText := ""
	scoreVGM := int64(0)
	scores := strings.TrimSpace(doc.Find(".rank_container_right .composite_group").Text())
	scoresList := strings.Split(scores, "|")
	for _, score := range scoresList {
		if strings.Contains(score, "Value:") {
			scoreValueText = strings.TrimSpace(strings.Replace(score, "Value:", "", 1))
			scoreValue = letterToNumber(scoreValueText)
			continue
		}
		if strings.Contains(score, "Growth:") {
			scoreGrowthText = strings.TrimSpace(strings.Replace(score, "Growth:", "", 1))
			scoreGrowth = letterToNumber(scoreGrowthText)
			continue
		}
		if strings.Contains(score, "Momentum:") {
			scoreMomentumText = strings.TrimSpace(strings.Replace(score, "Momentum:", "", 1))
			scoreMomentum = letterToNumber(scoreMomentumText)
			continue
		}
		if strings.Contains(score, "VGM:") {
			scoreVGMText = strings.TrimSpace(strings.Replace(score, "VGM:", "", 1))
			scoreVGM = letterToNumber(scoreVGMText)
			continue
		}
	}

	dividendAmount := float64(0)
	dividendPercent := float64(0)
	beta := float64(0)
	forwardPE := float64(0)
	pEGRatio := float64(0)

	err = nil
	doc.Find("#stock_activity .abut_bottom tr").Each(func(i int, s *goquery.Selection) {
		cols := s.Find("td")
		if cols.Length() != 2 {
			err = errors.New("Expected 2 td elements for the selector #stock_key_earnings .abut_bottom tr")
			return
		}

		if cols.First().Text() == "Dividend" {
			values := strings.Split(strings.TrimSpace(cols.Last().Text()), "(")
			if len(values) != 2 {
				err = fmt.Errorf("Expected 2 dividend sections but have %v", values)
			}

			dividendAmount, err = strconv.ParseFloat(strings.TrimSpace(values[0]), 64)
			if err != nil {
				return
			}

			dividendPercent, err = strconv.ParseFloat(
				strings.Replace(strings.TrimSpace(values[1]), "%)", "", 1),
				64,
			)
			if err != nil {
				return
			}
		}

		if cols.First().Text() == "Beta" {
			value := strings.TrimSpace(cols.Last().Text())
			beta, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return
			}
		}
	})

	doc.Find("#stock_key_earnings .abut_bottom tr").Each(func(i int, s *goquery.Selection) {
		cols := s.Find("td")
		if cols.Length() != 2 {
			err = errors.New("Expected 2 td elements for the selector #stock_key_earnings .abut_bottom tr")
			return
		}

		if cols.First().Text() == "Forward PE" {
			value := strings.TrimSpace(cols.Last().Text())
			forwardPE, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return
			}
		}

		if cols.First().Text() == "PEG Ratio" {
			value := strings.TrimSpace(cols.Last().Text())
			pEGRatio, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return
			}
		}
	})

	return &models.Rating{
		Name:              name,
		RankText:          rankText,
		Rank:              rank,
		ScoreValueText:    scoreValueText,
		ScoreValue:        scoreValue,
		ScoreGrowthText:   scoreGrowthText,
		ScoreGrowth:       scoreGrowth,
		ScoreMomentumText: scoreMomentumText,
		ScoreMomentum:     scoreMomentum,
		ScoreVGMText:      scoreVGMText,
		ScoreVGM:          scoreVGM,
		DividendAmount:    dividendAmount,
		DividendPercent:   dividendPercent,
		ForwardPE:         forwardPE,
		Beta:              beta,
		PEGRatio:          pEGRatio,
	}, nil
}

func letterToNumber(str string) int64 {
	if str == "A" {
		return 1
	}
	if str == "B" {
		return 2
	}
	if str == "C" {
		return 3
	}
	if str == "D" {
		return 4
	}
	if str == "E" {
		return 5
	}
	return 0
}
