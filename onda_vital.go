package ondavital

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"regexp"
)

const host_link = "https://es.wikipedia.org"
const wikipedia_link = host_link + "/w/index.php?profile=default&search="

func Search(query string) (string, error) {
	url := wikipedia_link + url.QueryEscape(query)
	return crawlLink(url, nil)
}

func crawlLink(link string, c chan string) (string, error) {

	doc, _ := goquery.NewDocument(link)

	jQueryString := ".infobox tr td:contains('España')"
	inner_text := doc.Find(jQueryString).Text()

	if inner_text == "" {

		if doc.Find(".ambox").Text() == "" {
			if c != nil {
				c <- "" //error
			}
			return "", ovError{"ERROR FINDING PAGE", "", []string{""}}
		}

		jQueryLinks := "#mw-content-text ul a"
		links := doc.Find(jQueryLinks).Map(func(i int, a *goquery.Selection) string {
			link_path, _ := a.Attr("href")
			return host_link + link_path
		})

		if len(links) > 0 {
			if c != nil {
				c <- "" //error
			}
			return recurseLinks(links)
		}
	}

	re := regexp.MustCompile(`(.*\))?(.*)\(España\)`)
	results := re.FindStringSubmatch(inner_text)

	if len(results) > 1 {
		/* here we can encounter spain in either the first or the second position*/
		nombre := results[2]

		if c != nil {
			c <- nombre
		}
		return nombre, nil
	}

	if c != nil {
		c <- "" //error
	}
	return "", ovError{"ERROR PARSING REGEXP", inner_text, results}
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
