package tools

import "log"

func ShortURL(url string) (string, error) {
	log.Println("SHORTENER", url+"SHORT")
	return url+"SHORT", nil
}
