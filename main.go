package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Team = struct {
	Name          string
	MainNicknames []string
	Year          int
	Others        []string
	// Some teams carry the name of their city and some cities are big/famous enough that
	// non football fans may end up using it as a password. This flag disables searching for
	// the raw name
	ExcludeBecauseCity bool
}

var teams = []Team{
	{Name: "juventus", MainNicknames: []string{"juve"}, Year: 1897, Others: []string{"5maggio", "finoallafine", "delpiero", "delpiero10"}},
	{Name: "inter", Year: 1908, Others: []string{"interfc", "fcinter", "triplete", "triplete2010", "maistatiinb"}},
	{Name: "milan", Year: 1899, Others: []string{"acmilan", "milanac", "maldini", "maldini3"}},
	{Name: "napoli", Year: 1926, Others: []string{"sscnapoli"}, ExcludeBecauseCity: true},
	{Name: "atalanta", MainNicknames: []string{"dea"}, Year: 1907},
	{Name: "roma", MainNicknames: []string{"magica"}, Year: 1927, Others: []string{"dajeroma", "asroma", "asroma1927", "laziomerda", "francescototti", "francescototti10"}, ExcludeBecauseCity: true},
	{Name: "lazio", Year: 1900, Others: []string{"26maggio", "sslazio", "lulic71", "romamerda"}},
	{Name: "fiorentina", MainNicknames: []string{"viola"}, Year: 1926},
	{Name: "torino", MainNicknames: []string{"toro"}, Year: 1906, ExcludeBecauseCity: true},
	{Name: "bologna", Year: 1909, ExcludeBecauseCity: true},
	{Name: "cagliari", Year: 1920, ExcludeBecauseCity: true},
	{Name: "verona", MainNicknames: []string{"hellas"}, Year: 1903, Others: []string{"hellasverona", "hellasverona1903"}, ExcludeBecauseCity: true},
	{Name: "genoa", MainNicknames: []string{"grifone"}, Year: 1893, Others: []string{"doriamerda"}},
	{Name: "udinese", Year: 1896},
	{Name: "sampdoria", MainNicknames: []string{"samp", "doria"}, Year: 1946},
	{Name: "venezia", Year: 1907, ExcludeBecauseCity: true},
	{Name: "empoli", Year: 1920},
	{Name: "sassuolo", Year: 1920},
	{Name: "spezia", Year: 1906},
	{Name: "salernitana", Year: 1919},
}

var results = map[string]int{}

func sha1String(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func findFrequency(team Team, password string) {
	cmd := exec.Command("rg", "--no-filename", "--no-line-number", sha1String(password))
	output, err := cmd.Output()
	if err != nil {
		return
	}
	count_s := strings.Split(string(output), ":")[1]
	count, _ := strconv.Atoi(strings.TrimSpace(count_s))
	fmt.Println(password, count)
	results[team.Name] += count
}

func main() {
	for _, team := range teams {
		cases := []string{
			"forza" + team.Name,
			"forza" + strings.Title(team.Name),
			"Forza" + strings.Title(team.Name),
			team.Name + fmt.Sprint(team.Year),
			"forza" + team.Name + fmt.Sprint(team.Year),
			"forza" + strings.Title(team.Name) + fmt.Sprint(team.Year),
			"Forza" + strings.Title(team.Name) + fmt.Sprint(team.Year),
		}
		if !team.ExcludeBecauseCity {
			cases = append(cases,
				team.Name,
				strings.Title(team.Name),
			)
		}
		if len(team.MainNicknames) > 0 {
			for _, nick := range team.MainNicknames {
				cases = append(cases, nick,
					"forza"+nick,
					"forza"+strings.Title(nick),
					"Forza"+strings.Title(nick),
					nick+fmt.Sprint(team.Year),
					"forza"+nick+fmt.Sprint(team.Year),
					"forza"+strings.Title(nick)+fmt.Sprint(team.Year),
					"Forza"+strings.Title(nick)+fmt.Sprint(team.Year))
			}
		}
		if len(team.Others) > 0 {
			for _, o := range team.Others {
				cases = append(cases, o)
			}
		}
		for _, c := range cases {
			findFrequency(team, c)
		}
	}
	fmt.Println(results)
}
