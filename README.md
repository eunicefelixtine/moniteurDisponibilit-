#  Go Uptime Monitor CLI

Un moniteur de disponibilité d'API et de sites web en ligne de commande (CLI), développé en **Go**. 

Ce projet utilise la puissance de la concurrence en Go (*goroutines* et *channels*) pour vérifier plusieurs URL en parallèle de manière ultra-rapide.

---

##  Fonctionnalités

* **Vérification concurrente :** Analyse des dizaines de sites simultanément grâce aux *goroutines*.
* **Mesure de performance :** Affiche le temps de réponse précis (en ms) pour chaque requête.
* **Gestion des timeouts :** Interrompt automatiquement les requêtes qui mettent plus de 5 secondes à répondre.
* **Léger et optimisé :** Utilise les requêtes HTTP `HEAD` au lieu de `GET` pour réduire la consommation de bande passante.
* **Gestion claire des erreurs :** Distingue les erreurs réseau (DNS, Timeout) des erreurs de statut HTTP (404, 500, etc.).

---

##  Prérequis

* **Go** (version 1.18 ou supérieure recommandée)  
  *Pour vérifier si Go est installé :* `go version`

---

##  Installation & Utilisation

### 1. Cloner le projet ou naviguer dans le dossier

```bash
git clone [https://github.com/votre-nom-utilisateur/uptime-monitor-go.git](https://github.com/votre-nom-utilisateur/moniteurDiponibilit-.git)
cd moniteurDiponibilit-
