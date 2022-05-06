# Serie A team passwords

This repo contains the code I used to perform the analysis for my blog post
[Estimating the size of Serie A fanbases using Pwned Passwords](https://tommasoamici.com/blog/estimating-the-size-of-serie-a-fanbases-using-pwned-passwords)

I had a bit of fun today poking around the [Pwned Passwords data breaches](https://haveibeenpwned.com/Passwords).

I wanted to see which Serie A team had the most fans from the number of leaked passwords that
can be reasonably associated with a specific team.

For example, "forzajuventus" is clearly a password we can associate with Juventus fans.

On the other hand, many teams carry the name of their city and some cities are big/famous
enough that non-football fans may use their name as a password. For example, Napoli or Verona.

## Reproduce or extend

If you want to reproduce these results, or run a similar analysis you will need a copy
of the Pwned Passwords breaches. I dropped passwords that occurred fewer than 16 times to
work with a more manageable 1.5GB text file.

```sh
sed -i '33000000,$d' pwned-passwords-sha1-ordered-by-count-v8.txt
```

I have then written a small Go script which you can find in [main.go](./main.go). This script
creates a few passwords for each team and calculates the SHA1 hash to check against the
password dump with `ripgrep`.

## Results

| Team        | Password          |   Count |
| ----------- | ----------------- | ------: |
| Juventus    |                   | 306,731 |
|             | juventus          | 213,620 |
|             | delpiero          |  41,087 |
|             | forzajuve         |  15,975 |
|             | delpiero10        |  14,009 |
|             | Juventus          |   8,336 |
|             | juve              |   4,349 |
|             | juve1897          |   3,392 |
|             | juventus1897      |   3,118 |
|             | forzajuventus     |   1,572 |
|             | 5maggio           |     582 |
|             | finoallafine      |     340 |
|             | forzajuve1897     |     152 |
|             | ForzaJuve1897     |      83 |
|             | ForzaJuve         |      65 |
|             | ForzaJuventus     |      18 |
|             | forzajuventus1897 |      17 |
|             | forzaJuve         |      16 |
| Milan       |                   | 140,521 |
|             | acmilan           |  46,991 |
|             | milan             |  30,808 |
|             | forzamilan        |  26,944 |
|             | maldini           |  13,979 |
|             | milan1899         |   9,939 |
|             | milanac           |   5,034 |
|             | maldini3          |   4,902 |
|             | Milan             |   1,332 |
|             | forzamilan1899    |     347 |
|             | ForzaMilan        |     199 |
|             | forzaMilan        |      26 |
|             | ForzaMilan1899    |      20 |
| Roma        |                   |  83,457 |
|             | asroma            |  25,803 |
|             | forzaroma         |  25,011 |
|             | magica            |  11,102 |
|             | totti10           |  13,425 |
|             | asroma1927        |   9,968 |
|             | totti             |   5,306 |
|             | roma1927          |   4,294 |
|             | laziomerda        |   2,568 |
|             | francescototti    |   2,098 |
|             | dajeroma          |     969 |
|             | forzaroma1927     |     807 |
|             | francescototti10  |     566 |
|             | ForzaRoma         |     120 |
|             | forzamagica       |      95 |
|             | magica1927        |      56 |
| Inter       |                   |  67,153 |
|             | inter             |  23,297 |
|             | forzainter        |  21,941 |
|             | inter1908         |  14,417 |
|             | fcinter           |   3,477 |
|             | interfc           |   1,735 |
|             | triplete          |     998 |
|             | Inter             |     506 |
|             | forzainter1908    |     366 |
|             | triplete2010      |     244 |
|             | ForzaInter        |      98 |
|             | maistatiinb       |      30 |
|             | ForzaInter1908    |      26 |
|             | forzaInter        |      18 |
| Lazio       |                   |  33,965 |
|             | sslazio           |   9,624 |
|             | forzalazio        |   8,846 |
|             | lazio1900         |   7,074 |
|             | lazio             |   4,905 |
|             | romamerda         |   2,246 |
|             | 26maggio          |     660 |
|             | forzalazio1900    |     434 |
|             | Lazio             |      77 |
|             | lulic71           |      45 |
|             | ForzaLazio        |      27 |
|             | ForzaLazio1900    |      27 |
| Fiorentina  |                   |  32,972 |
|             | fiorentina        |  20,003 |
|             | viola             |   8,352 |
|             | forzaviola        |   2,764 |
|             | Fiorentina        |     745 |
|             | fiorentina1926    |     523 |
|             | viola1926         |     449 |
|             | forzafiorentina   |     103 |
|             | forzaviola1926    |      33 |
| Napoli      |                   |  28,375 |
|             | forzanapoli       |  14,879 |
|             | napoli1926        |  11,216 |
|             | sscnapoli         |   1,377 |
|             | forzanapoli1926   |     852 |
|             | ForzaNapoli       |      51 |
| Verona      |                   |  15,232 |
|             | hellas            |  12,911 |
|             | hellasverona      |   1,135 |
|             | hellas1903        |     762 |
|             | verona1903        |     204 |
|             | hellasverona1903  |     105 |
|             | forzahellas       |      94 |
|             | forzaverona       |      21 |
| Sampdoria   |                   |  14,954 |
|             | sampdoria         |  11,644 |
|             | doria             |     901 |
|             | samp              |     527 |
|             | samp1946          |     471 |
|             | sampdoria1946     |     399 |
|             | forzasamp         |     372 |
|             | Sampdoria         |     290 |
|             | forzadoria        |     264 |
|             | forzasampdoria    |      48 |
|             | doria1946         |      38 |
| Atalanta    |                   |  14,209 |
|             | atalanta          |   9,003 |
|             | dea               |   3,378 |
|             | atalanta1907      |   1,122 |
|             | Atalanta          |     375 |
|             | dea1907           |     258 |
|             | forzaatalanta     |      56 |
|             | forzadea          |      17 |
| Genoa       |                   |  11,509 |
|             | genoa1893         |   4,973 |
|             | grifone           |   2,755 |
|             | forzagenoa        |   1,635 |
|             | genoa             |   1,282 |
|             | doriamerda        |     524 |
|             | grifone1893       |     132 |
|             | forzagenoa1893    |     102 |
|             | Genoa             |      54 |
|             | forzagrifone      |      52 |
| Torino      |                   |   7,210 |
|             | forzatoro         |   3,798 |
|             | toro              |   2,525 |
|             | torino1906        |     486 |
|             | toro1906          |     259 |
|             | forzatorino       |      63 |
|             | ForzaToro         |      43 |
|             | forzatoro1906     |      36 |
| Salernitana |                   |   3,839 |
|             | salernitana       |   3,143 |
|             | salernitana1919   |     557 |
|             | Salernitana       |     101 |
|             | forzasalernitana  |      38 |
| Udinese     |                   |   2,835 |
|             | udinese           |   2,558 |
|             | forzaudinese      |     115 |
|             | Udinese           |      93 |
|             | udinese1896       |      69 |
| Spezia      |                   |   1,311 |
|             | spezia            |   1,033 |
|             | forzaspezia       |     179 |
|             | spezia1906        |      99 |
| Empoli      |                   |   1,127 |
|             | empoli            |     916 |
|             | forzaempoli       |     124 |
|             | empoli1920        |      59 |
|             | Empoli            |      28 |
| Sassuolo    |                   |     751 |
|             | sassuolo          |     713 |
|             | forzasassuolo     |      17 |
|             | Sassuolo          |      21 |
| Bologna     |                   |     690 |
|             | forzabologna      |     423 |
|             | bologna1909       |     267 |
| Cagliari    |                   |     662 |
|             | forzacagliari     |     532 |
|             | cagliari1920      |     130 |
| Venezia     |                   |      90 |
|             | forzavenezia      |      46 |
|             | venezia1907       |      44 |
