# Serie A team passwords

I had a bit of fun today poking around the [Pwned Passwords data breaches](https://haveibeenpwned.com/Passwords).

I wanted to see which Serie A team had the most fans from the number of leaked passwords that
can be reasonably associated to a specific team.

For example, "forzajuventus" is clearly a password we can associate with Juventus fans.

On the other hand, many teams carry the name of their city and some cities are big/famous enough that
non-football fans may use their name as a password. For example, Napoli or Verona.

Here are the results:

| Team        | Count  |
| ----------- | ------ |
| Juventus    | 306731 |
| Milan       | 140521 |
| Roma        | 83457  |
| Inter       | 67153  |
| Lazio       | 33965  |
| Fiorentina  | 32972  |
| Napoli      | 28375  |
| Verona      | 15232  |
| Sampdoria   | 14954  |
| Atalanta    | 14209  |
| Genoa       | 11509  |
| Torino      | 7210   |
| Salernitana | 3839   |
| Udinese     | 2835   |
| Spezia      | 1311   |
| Empoli      | 1127   |
| Sassuolo    | 751    |
| Bologna     | 690    |
| Cagliari    | 662    |
| Venezia     | 90     |

## Reproduce or extend

If you want to reproduce these results, or run a similar analysis you will need a copy
of the Pwned Passwords breaches.

I have written a small Go script which you can find in [main.go](./main.go). This script
creates a few passwords for each team and calculates the SHA1 hash to check against the
password dump with `ripgrep`.
