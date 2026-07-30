package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Resultat conserve les informations de santé d'un site
type Resultat struct {
	URL        string
	EnLigne    bool
	StatutCode int
	Duree      time.Duration
	Erreur     error
}

// verifierSite effectue une requête HTTP HEAD/GET sur une URL donnée
func verifierSite(url string, wg *sync.WaitGroup, ch chan<- Resultat) {
	defer wg.Done() // Indique au WaitGroup que cette goroutine est terminée à la fin

	debut := time.Now()

	// On configure un client HTTP avec un délai d'attente (timeout) de 5 secondes
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	// On utilise Head pour éviter de télécharger tout le contenu de la page (plus rapide)
	resp, err := client.Head(url)
	duree := time.Since(debut)

	if err != nil {
		ch <- Resultat{
			URL:     url,
			EnLigne: false,
			Duree:   duree,
			Erreur:  err,
		}
		return
	}
	defer resp.Body.Close()

	// On considère le site en ligne si le statut est entre 200 et 399
	enLigne := resp.StatusCode >= 200 && resp.StatusCode < 400

	ch <- Resultat{
		URL:        url,
		EnLigne:    enLigne,
		StatutCode: resp.StatusCode,
		Duree:      duree,
		Erreur:     nil,
	}
}

func main() {
	// Liste des URL à vérifier
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://httpbin.org/status/404", // Exemple d'erreur 404
		"https://httpbin.org/delay/2",   // Exemple avec du délai
		"https://un-domaine-qui-n-existe-pas-12345.com",
	}

	var wg sync.WaitGroup
	canalResultats := make(chan Resultat, len(urls))

	fmt.Printf("Début de la vérification de %d sites\n\n", len(urls))

	// Lancement d'une goroutine par URL pour un traitement en parallèle
	for _, url := range urls {
		wg.Add(1)
		go verifierSite(url, &wg, canalResultats)
	}

	// Goroutine anonyme pour fermer le canal une fois toutes les requêtes finies
	go func() {
		wg.Wait()
		close(canalResultats)
	}()

	// Lecture et affichage des résultats depuis le canal
	for res := range canalResultats {
		if res.EnLigne {
			fmt.Printf("[EN LIGNE]    %-45s | Statut: %d | Temps: %v\n", res.URL, res.StatutCode, res.Duree.Round(time.Millisecond))
		} else if res.Erreur != nil {
			fmt.Printf("[HORS LIGNE]  %-45s | Erreur: %v\n", res.URL, res.Erreur)
		} else {
			fmt.Printf("[ERREUR HTTP] %-45s | Statut: %d | Temps: %v\n", res.URL, res.StatutCode, res.Duree.Round(time.Millisecond))
		}
	}

	fmt.Println("\n Vérification terminée")
}
