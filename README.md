# Pokedex CLI

Pokedex CLI e' un progetto che ho utilizzato per imparare a usare API in Go e a gestire richieste HTTP. L'applicazione interroga la [PokeAPI](https://pokeapi.co/) per recuperare aree della mappa, Pokemon presenti in una determinata zona e dettagli dei Pokemon catturati.

Il progetto e' una CLI interattiva: dopo l'avvio, mostra un prompt `Pokedex >` da cui e' possibile eseguire i comandi disponibili.

## Funzionamento

L'applicazione usa il package `net/http` di Go per inviare richieste HTTP alla PokeAPI e `encoding/json` per convertire le risposte JSON in strutture Go.

Durante l'esecuzione viene mantenuta una cache in memoria con durata di 5 minuti. In questo modo, se la stessa risorsa viene richiesta piu' volte, l'app puo' riutilizzare la risposta gia' ottenuta invece di effettuare una nuova chiamata HTTP.

La CLI permette di:

- consultare le aree della mappa disponibili;
- navigare avanti e indietro tra le pagine delle aree;
- esplorare una specifica area per vedere quali Pokemon contiene;
- provare a catturare un Pokemon;
- ispezionare i dettagli dei Pokemon catturati;
- visualizzare il proprio Pokedex.

## Requisiti

- Go 1.25.5 o superiore
- Connessione internet per raggiungere la PokeAPI

## Avvio

Per eseguire l'applicazione:

```bash
go run .
```

Per eseguire i test:

```bash
go test ./...
```

## Comandi

| Comando | Descrizione |
| --- | --- |
| `help` | Mostra l'elenco dei comandi disponibili. |
| `map` | Mostra 20 aree della mappa. Ripetendo il comando si passa alla pagina successiva. |
| `mapb` | Mostra le 20 aree precedenti della mappa. |
| `explore <area>` | Mostra i Pokemon presenti in una specifica area. |
| `catch <pokemon>` | Prova a catturare un Pokemon. La probabilita' dipende dalla sua esperienza base. |
| `inspect <pokemon>` | Mostra i dettagli di un Pokemon gia' catturato. |
| `pokedex` | Mostra l'elenco dei Pokemon catturati. |
| `exit` | Chiude l'applicazione. |

## Esempio di funzionamento

```text
$ go run .
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
 - tentacool
 - tentacruel
 - staryu
 - magikarp
 - gyarados
 - wingull
 - pelipper
 - shellos
 - gastrodon
 - finneon
 - lumineon

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
  -hp: 35
  -attack: 55
  -defense: 40
  -special-attack: 50
  -special-defense: 50
  -speed: 90
Types:
  - electric

Pokedex > pokedex
Your Pokedex:
 - pikachu

Pokedex > exit
Closing the Pokedex... Goodbye!
```

Nota: il comando `catch` usa una probabilita' di cattura, quindi lo stesso Pokemon potrebbe anche scappare.
