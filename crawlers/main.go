package main

import (
	"crawlers/common"
	"crawlers/registry"
	"log"
	"sync"
)

func main() {
	parsers := registry.All()

	var wg sync.WaitGroup

	for _, parser := range parsers {
		log.Printf("Starting %s parser...\n", parser.Name())
		wg.Add(1)
		go func(p common.Parser) {
			defer wg.Done()

			results, err := p.Parse("laptop")
			if err != nil {
				log.Printf("Error parsing %s: %v\n", p.Name(), err)
				return
			}

			log.Printf("%s: %d results\n", p.Name(), len(results))
		}(parser)

		wg.Wait()
	}
}

