package data

import "fdsim/enums"

func GetCities(country enums.Country) []string {
	switch country {
	case enums.IT:
		return italianCities
	case enums.EN:
		return englishCities
	case enums.FR:
		return frenchCities
	case enums.ES:
		return spanishCities
	case enums.DE:
		return germanCities
	}

	return italianCities
}

var italianCities = []string{
	"Roma", "Milano", "Napoli", "Torino", "Palermo", "Genoa", "Bologna", "Firenze", "Bari", "Catania",
	"Venice", "Verona", "Messina", "Padua", "Trieste", "Taranto", "Brescia", "Reggio Calabria", "Modena", "Prato",
	"Parma", "Perugia", "Livorno", "Ravenna", "Cagliari", "Foggia", "Reggio Emilia", "Salerno", "Rimini", "Ferrara",
}

var englishCities = []string{
	"London", "Birmingham", "Manchester", "Leeds", "Newcastle", "Sheffield", "Liverpool", "Bristol", "Nottingham", "Southampton",
	"Leicester", "Coventry", "Bradford", "Stoke-on-Trent", "Wolverhampton", "Plymouth", "Derby", "Reading", "Bolton", "Bournemouth",
	"Brighton", "Huddersfield", "Middlesbrough", "Blackpool", "Peterborough", "Swansea", "Oxford", "Ipswich", "Cambridge", "Wigan",
}

var spanishCities = []string{
	"Madrid", "Barcelona", "Valencia", "Seville", "Zaragoza", "Málaga", "Murcia", "Palma", "Las Palmas de Gran Canaria", "Bilbao",
	"Alicante", "Córdoba", "Valladolid", "Vigo", "Gijón", "Hospitalet de Llobregat", "A Coruña", "Vitoria-Gasteiz", "Granada", "Elche",
	"Oviedo", "Badalona", "Terrassa", "Cartagena", "Jerez de la Frontera", "Sabadell", "Móstoles", "Santa Cruz de Tenerife", "Alcalá de Henares", "Pamplona",
}

var germanCities = []string{
	"Berlin", "Hamburg", "Munich", "Cologne", "Frankfurt", "Stuttgart", "Düsseldorf", "Dortmund", "Essen", "Leipzig",
	"Bremen", "Dresden", "Hanover", "Nuremberg", "Duisburg", "Bochum", "Wuppertal", "Bielefeld", "Bonn", "Münster",
	"Karlsruhe", "Mannheim", "Augsburg", "Wiesbaden", "Gelsenkirchen", "Mönchengladbach", "Braunschweig", "Chemnitz", "Kiel", "Aachen",
}
var frenchCities = []string{
	"Paris", "Marseille", "Lyon", "Toulouse", "Nice", "Nantes", "Strasbourg", "Montpellier", "Bordeaux", "Lille",
	"Rennes", "Reims", "Le Havre", "Cergy-Pontoise", "Saint-Étienne", "Toulon", "Angers", "Grenoble", "Dijon", "Nîmes",
	"Aix-en-Provence", "Saint-Quentin-en-Yvelines", "Brest", "Le Mans", "Amiens", "Tours", "Limoges", "Clermont-Ferrand", "Villeurbanne", "Besançon",
}
