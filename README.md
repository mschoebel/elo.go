# Elo.go

Calculate [Elo ratings](http://en.wikipedia.org/wiki/Elo_rating_system) for small tournaments.


## Usage

Create a tournament result description in JSON format. Example:

```
{
  "participants": {
    "A": "Alice",
    "B": "Bob",
    "C": "Charly"
  },
  "results": [
    {
      "player1": "A",
      "player2": "B",
      "score": 1.0
    },
    {
      "player1": "A",
      "player2": "C",
      "score": 0.5
    },
    {
      "player1": "B",
      "player2": "C",
      "score": 0.0
    }
  ]
}
```

Build and compile `elo.go`. Output:

```
 A vs.  B = 1.0
 A vs.  C = 0.5
 B vs.  C = 0.0

 1. Charly               (1015)
 2. Alice                (1014)
 3. Bob                  (971)
```
