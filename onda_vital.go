package ondavital

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"regexp"
	"strings"
)

const hostLink = "https://es.wikipedia.org"
const wikipediaLink = hostLink + "/w/index.php?profile=default&search="

func Search(query string) (string, error) {
	if query == "" {
		return "", nil
	}

	url := wikipediaLink + url.QueryEscape(query)
	return crawlLink(url, nil)
}

func crawlLink(link string, c chan string) (string, error) {

	doc, _ := goquery.NewDocument(link)

	jQueryString := ".infobox tr td:contains('España')"
	innerText := doc.Find(jQueryString).Text()

	if innerText == "" {

		/* we are on the searched page */
		if doc.Find(".ambox .mw-disambig").Text() == "" {
			if c != nil {
				c <- "" //error
			}
			return "", ovError{"ERROR FINDING PAGE", "", []string{""}}
		}

		/* test if we are on  a suggestion page */
		jQueryLinks := "#mw-content-text ul a"
		links := doc.Find(jQueryLinks).Map(func(i int, a *goquery.Selection) string {
			linkPath, _ := a.Attr("href")
			return hostLink + linkPath
		})

		if len(links) > 0 {
			if c != nil {
				c <- "" //error
			}
			return recurseLinks(links)
		}
	}

	re := regexp.MustCompile(`(.*\))?(.*)\(España\)`)
	results := re.FindStringSubmatch(innerText)

	if len(results) > 1 {
		/* here we can encounter spain in either the first or the second position*/
		nombre := results[2]
		nombre = strings.TrimSpace(nombre)

		if c != nil {
			c <- nombre
		}
		return nombre, nil
	}

	if c != nil {
		c <- "" //error
	}
	return "", ovError{"ERROR PARSING REGEXP", innerText, results}
}

func recurseLinks(links []string) (string, error) {

	c := make(chan string, len(links))
	/* send go routines...*/
	for _, v := range links {
		go crawlLink(v, c)
	}

	/* wait for them */
	var result string
	for range links {
		result = <-c
		if result != "" {
			return result, nil //win
		}
	}

	return "", ovError{"ERROR FINDING PAGE", "", []string{""}}
}
