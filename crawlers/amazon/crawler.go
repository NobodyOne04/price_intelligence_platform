package amazon

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"crawlers/common"
	"github.com/gocolly/colly/v2"
)

type AmazonParser struct{}

func NewParser() *AmazonParser {
	return &AmazonParser{}
}

func (p *AmazonParser) Name() string {
	return "amazon"
}

func (p *AmazonParser) Parse(keyword string) ([]common.Result, error) {
	var results []common.Result

	cSearch, cProduct := newCollectors()
	setupSearchCollector(cSearch, cProduct)
	setupProductCollector(cProduct, &results)

	log.Println("🎯 Start scraping Amazon...")
	err := cSearch.Visit(fmt.Sprintf("https://www.amazon.com/s?k=%s", keyword))
	if err != nil {
		return nil, fmt.Errorf("failed to visit search page: %w", err)
	}

	cSearch.Wait()
	cProduct.Wait()

	log.Printf("Scrape done! Found %d products\n", len(results))
	return results, nil
}

const Parallelism = 15
const RandomDelay = 3
const domain = "www.amazon.com"

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/122.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 Safari/605.1.15",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 Chrome/119.0.0.0 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 16_5 like Mac OS X) AppleWebKit/605.1.15 Mobile/15E148",
	"Mozilla/5.0 (iPad; CPU OS 15_6 like Mac OS X) AppleWebKit/605.1.15 Version/15.0 Mobile/15E148 Safari/604.1",
}

var acceptLanguages = []string{
	"en-US,en;q=0.9",
	"en-GB,en;q=0.8",
	"de-DE,de;q=0.9,en;q=0.8",
	"fr-FR,fr;q=0.9,en;q=0.8",
}

var referers = []string{
	"https://www.google.com/",
	"https://www.bing.com/",
	"https://duckduckgo.com/",
}

func newCollectors() (*colly.Collector, *colly.Collector) {
	cSearch := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.UserAgent(userAgents[rand.Intn(len(userAgents))]),
		colly.Async(true),
	)

	cSearch.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: Parallelism,
		RandomDelay: RandomDelay * time.Second,
	})

	cProduct := cSearch.Clone()
	return cSearch, cProduct
}

func setupSearchCollector(cSearch, cProduct *colly.Collector) {
	cSearch.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", acceptLanguages[rand.Intn(len(acceptLanguages))])
		r.Headers.Set("Referer", referers[rand.Intn(len(referers))])
	})

	cSearch.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.Contains(href, "/dp/") {
			link := e.Request.AbsoluteURL(href)
			if strings.Contains(link, "slredirect") || strings.Contains(link, "redirect") {
				return
			}
			_ = cProduct.Visit(link)
		}
	})

	cSearch.OnError(func(r *colly.Response, err error) {
		log.Printf("Error %d on %s: %v\n", r.StatusCode, r.Request.URL, err)
	})
}

func setupProductCollector(cProduct *colly.Collector, results *[]common.Result) {
	cProduct.OnHTML("body", func(e *colly.HTMLElement) {
		title := e.DOM.Find("#productTitle").Text()
		price := e.DOM.Find("span.a-price > span.a-offscreen").First().Text()
		seller := e.DOM.Find("#merchant-info").Text()
		delivery := e.DOM.Find("#delivery-message span").First().Text()

		*results = append(*results, common.Result{
			Title:     strings.TrimSpace(title),
			SoldBy:    strings.TrimSpace(seller),
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
			Source:    "amazon",
		})

		log.Printf(
			"%s — %s - %s - %s\n",
			strings.TrimSpace(title),
			strings.TrimSpace(price),
			strings.TrimSpace(seller),
			strings.TrimSpace(delivery),
		)
	})

	cProduct.OnError(func(r *colly.Response, err error) {
		log.Printf("Product error %d on %s: %v\n", r.StatusCode, r.Request.URL, err)
	})
}

