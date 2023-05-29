package data

import "fdsim/enums"

func GetNames(country enums.Country) []string {
	switch country {
	case enums.IT:
		return italianNames
	case enums.EN:
		return englishNames
	case enums.FR:
		return frenchNames
	case enums.ES:
		return spanishNames
	case enums.DE:
		return germanNames
	}

	return italianNames
}

func GetSurnames(country enums.Country) []string {
	switch country {
	case enums.IT:
		return italianSurnames
	case enums.EN:
		return englishSurnames
	case enums.FR:
		return frenchSurnames
	case enums.ES:
		return spanishSurnames
	case enums.DE:
		return germanSurnames
	}

	return italianSurnames
}

var italianNames = []string{
	"Luca", "Leonardo", "Mattia", "Andrea", "Gabriele", "Francesco", "Alessandro", "Lorenzo", "Riccardo", "Edoardo",
	"Davide", "Simone", "Giovanni", "Michele", "Matteo", "Marco", "Filippo", "Christian", "Nicola", "Federico",
	"Antonio", "Giacomo", "Pietro", "Raffaele", "Emanuele", "Samuele", "Gianluca", "Daniele", "Tommaso", "Alessio",
	"Enrico", "Vincenzo", "Cristian", "Diego", "Manuel", "Gianmarco", "Giuseppe", "Stefano", "Alberto", "Salvatore",
	"Massimiliano", "Gianluigi", "Angelo", "Emiliano", "Marcello", "Roberto", "Dario", "Fabio", "Luigi", "Ciro",
	"Mirko", "Loris", "Samuel", "Ettore", "Jacopo", "Maurizio", "Domenico", "Rocco", "Gabriel", "Niccolò",
	"Ruggero", "Alessandro", "Antonino", "Paolo", "Luciano", "Cesare", "Renato", "Sergio", "Omar", "Gustavo",
}

var italianSurnames = []string{
	"Rossi", "Bianchi", "Romano", "Colombo", "Ricci", "Marino", "Greco", "Conti", "Esposito", "Rizzo",
	"Ferrari", "Barbieri", "Galli", "Martini", "Fontana", "Leone", "Lombardi", "Moretti", "Santoro", "Mancini",
	"Costa", "Giordano", "Rinaldi", "Caruso", "Ferrara", "Pagano", "Villa", "Conte", "Ferraro", "De Luca",
	"Bruno", "Serra", "Mariani", "Riva", "Fabbri", "Marchetti", "Gatti", "Rizzo", "Silvestri", "Sartori",
	"Negri", "Ferri", "Bianco", "Basile", "Pellegrini", "Marchi", "Testa", "Piras", "Monti", "Palumbo",
	"Barone", "Franco", "Costantino", "Valenti", "Santini", "D'Amico", "Coppola", "Marini", "Vitale", "Sorrentino",
	"Sala", "Donati", "D'Angelo", "Palmieri", "De Angelis", "Cattaneo", "Gallo", "Battaglia", "Piazza", "Guerra",
	"Ruggeri", "Piras", "Rizzi", "Montanari", "Caputo", "Santucci", "Longo", "De Rosa", "Giuliani", "Amato",
	"Carbone", "Benedetti", "Farina", "Grasso", "Morelli", "Mazza", "Costa", "Gallo", "Rossetti", "Donati",
	"Rizzi", "Montanari",
}

var frenchNames = []string{
	"Lucas", "Hugo", "Gabriel", "Louis", "Ethan", "Arthur", "Paul", "Léo", "Nathan", "Jules",
	"Tom", "Théo", "Raphaël", "Mathis", "Noah", "Enzo", "Liam", "Adam", "Antoine", "Maxime",
	"Baptiste", "Alexandre", "Victor", "Nolan", "Matéo", "Timéo", "Clément", "Julien", "Thomas", "Maxence",
	"Gabin", "Samuel", "Axel", "Rayan", "Romain", "Luc", "Tristan", "Léonard", "Noé", "Malo",
	"Evan", "Sacha", "Thibault", "Maël", "Quentin", "Oscar", "Valentin", "Corentin", "Mathys", "Émile",
	"Antonin", "Ilyès", "Guillaume", "Julian", "Léandre", "Charles", "Raphaël", "Nils", "Marceau", "Rémi",
	"Robin", "Loan", "Esteban", "Augustin", "Justin", "Kyllian", "Marius", "Dorian", "Bastien", "Aymeric",
	"Yanis", "Mathieu", "Thibaut", "Noham", "Erwan", "Tiago", "Eden",
}

var frenchSurnames = []string{
	"Martin", "Bernard", "Thomas", "Petit", "Robert", "Richard", "Durand", "Dubois", "Moreau", "Simon",
	"Laurent", "Lefebvre", "Michel", "Garcia", "David", "Bertrand", "Roux", "Vincent", "Fournier", "Morel",
	"Girard", "André", "Lefèvre", "Mercier", "Dupont", "Lambert", "Bonnet", "François", "Martinez", "Legrand",
	"Garnier", "Faure", "Rousseau", "Blanc", "Guerin", "Muller", "Henry", "Roussel", "Nicolas", "Perrin",
	"Morin", "Mathieu", "Clement", "Gauthier", "Dumont", "Lopez", "Fontaine", "Chevalier", "Robin", "Masson",
	"Sanchez", "Meyer", "Dupuis", "Berger", "Gauthier", "Gonzalez", "Fernandez", "Philippe", "Lemoine", "Leroy",
	"Riviere", "Leclerc", "Bourgeois", "Royer", "Dupré", "Caron", "Colin", "Marchand", "Rodriguez", "Aubert",
	"Renard", "Lucas", "Marty", "Fleury", "Benoit", "Huet", "Barbier", "Brun", "Carpentier", "Roche",
	"Schmitt", "Dumas", "Lemaire", "Picard", "Roger", "Fabre", "Guillaume", "Moulin", "Duval", "Bourgeois",
}

var germanNames = []string{
	"Peter", "Michael", "Thomas", "Andreas", "Stefan", "Christian", "Markus", "Daniel", "Jürgen", "Martin",
	"Frank", "Hans", "Klaus", "Uwe", "Wolfgang", "Dirk", "Oliver", "Sven", "Alexander", "Jan",
	"Torsten", "Matthias", "Marco", "Holger", "Ralf", "Sascha", "Jens", "Marcus", "Udo", "Heinz",
	"Patrick", "Rainer", "Sebastian", "David", "Kai", "Tobias", "Karl", "Bernd", "Erik", "Rolf",
	"Dominik", "Steffen", "Björn", "Detlef", "Dieter", "Lars", "Hermann", "Axel", "Mathias", "Georg",
	"Stephan", "Norbert", "Günter", "Joachim", "Tim", "Thorsten", "Werner", "Ulf", "Johannes", "Volker",
	"Gerd", "Fritz", "Marcel", "Helmut", "Max", "Walter", "Richard", "Paul", "Emil", "Leonard",
	"Benjamin", "Willi", "Rüdiger", "Friedrich", "Herbert", "Nico", "Erich", "Alfred", "Armin", "René",
	"Philipp", "André", "Harry", "Ullrich", "Manfred", "Rudolf", "Olaf", "Erwin", "Mario", "Enrico",
	"Hannes", "Robin", "Jörg", "Samuel", "Simon", "Dominic",
}

var germanSurnames = []string{
	"Müller", "Schmidt", "Schneider", "Fischer", "Weber", "Meyer", "Wagner", "Becker", "Schulz", "Hoffmann",
	"Schäfer", "Koch", "Bauer", "Richter", "Klein", "Wolf", "Schröder", "Neumann", "Schwarz", "Zimmermann",
	"Braun", "Krüger", "Hofmann", "Hartmann", "Lange", "Schmitt", "Werner", "Schmitz", "Krause", "Meier",
	"Lehmann", "Schmid", "Schulze", "Maier", "Köhler", "Herrmann", "König", "Walter", "Mayer", "Huber",
	"Kaiser", "Fuchs", "Peters", "Lang", "Scholz", "Möller", "Weiß", "Jung", "Hahn", "Schubert",
	"Vogel", "Friedrich", "Keller", "Günther", "Frank", "Berger", "Roth", "Beck", "Lorenz", "Baumann",
	"Franke", "Albrecht", "Schuster", "Simon", "Ludwig", "Böhm", "Winter", "Stein", "Brandt", "Haas",
	"Schreiber", "Graf", "Jäger", "Sommer", "Seidel", "Heinrich", "Schulte", "Kühn", "Ziegler", "Kraft",
	"Böhme", "Ebert", "Kramer", "Bergmann", "Pfeiffer", "Kirsch", "Schilling", "Thiel", "Behrens", "Busch",
}

var englishNames = []string{
	"James", "John", "Robert", "Michael", "William", "David", "Joseph", "Charles", "Thomas", "Daniel", "Matthew",
	"Anthony", "Donald", "Mark", "Paul", "Steven", "Andrew", "Kenneth", "George", "Joshua", "Kevin", "Brian",
	"Edward", "Ronald", "Timothy", "Jason", "Jeffrey", "Ryan", "Jacob", "Gary", "Nicholas", "Eric", "Stephen",
	"Jonathan", "Larry", "Justin", "Scott", "Brandon", "Benjamin", "Samuel", "Frank", "Gregory", "Patrick",
	"Raymond", "Alexander", "Jack", "Dennis", "Jerry", "Tyler", "Aaron", "Henry", "Douglas", "Peter", "Walter",
	"Arthur", "Kyle", "Carl", "Harold", "Jeremy", "Keith", "Roger", "Gerald", "Terry", "Lawrence", "Sean",
	"Jesse", "Christian", "Ethan", "Austin", "Joe", "Albert", "Jared", "Billy", "Bruce", "Ralph", "Bryan",
	"Billy", "Louis", "Eugene", "Roy", "Wayne", "Alan", "Juan", "Russell", "Fred", "Randy", "Philip", "Howard",
	"Vincent", "Bobby", "Johnny", "Clarence", "Travis", "Craig", "Jimmy", "Johnny",
}

var englishSurnames = []string{
	"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis", "Miller", "Wilson", "Moore", "Taylor",
	"Anderson", "Thomas", "Jackson", "White", "Harris", "Martin", "Thompson", "Garcia", "Martinez", "Robinson",
	"Clark", "Rodriguez", "Lewis", "Lee", "Walker", "Hall", "Allen", "Young", "Hernandez", "King",
	"Wright", "Lopez", "Hill", "Scott", "Green", "Adams", "Baker", "Gonzalez", "Nelson", "Carter",
	"Mitchell", "Perez", "Roberts", "Turner", "Phillips", "Campbell", "Parker", "Evans", "Edwards", "Collins",
	"Stewart", "Sanchez", "Morris", "Rogers", "Reed", "Cook", "Morgan", "Bell", "Murphy", "Bailey",
	"Rivera", "Cooper", "Richardson", "Cox", "Howard", "Ward", "Torres", "Peterson", "Gray", "Ramirez",
	"James", "Watson", "Brooks", "Kelly", "Sanders", "Price", "Bennett", "Wood", "Barnes", "Ross",
	"Henderson", "Coleman", "Jenkins", "Perry", "Powell", "Long", "Patterson", "Hughes", "Flores", "Washington",
	"Simmons", "Foster", "Butler", "Bryant", "Alexander", "Russell", "Griffin", "Hayes", "Murray", "Ford",
}

var spanishNames = []string{
	"José", "Antonio", "Manuel", "Francisco", "Juan", "David", "Javier", "Jesús", "Miguel", "Ángel",
	"Pedro", "Luis", "Carlos", "Alberto", "Rafael", "Fernando", "Daniel", "Pablo", "Sergio", "Andrés",
	"Diego", "Mariano", "Ignacio", "Ramón", "Víctor", "Rubén", "Adrián", "Alfredo", "Enrique", "Óscar",
	"Roberto", "Ricardo", "Héctor", "Emilio", "Federico", "Agustín", "César", "Eduardo", "Gregorio", "Jaime",
	"Arturo", "Samuel", "Ernesto", "Tomás", "Gonzalo", "Salvador", "Raúl", "Eugenio", "Alex", "Jorge",
	"Gabriel", "Marcos", "Rodrigo", "Hugo", "Joel", "Felipe", "Gustavo", "Alejandro", "Lucas", "Hernán",
	"Benjamín", "Ismael", "Gilberto", "Adolfo", "Lorenzo", "Efrén", "Octavio", "Nicolás", "Abelardo", "Leonardo",
	"Julio", "Omar", "Armando", "Josué", "René", "Marcial", "Ezequiel", "Mario", "Ulises", "Rolando",
	"Leopoldo", "Teodoro", "Joselito", "Albert", "Carmelo", "Diego Armando", "Fidel", "Rafael Ángel", "Álvaro", "Félix",
	"Damián", "Maximiliano", "Simón", "Martín", "Guillermo", "Noé", "Roberto Carlos", "Cristian", "Ángel Gabriel", "Nelson",
}

var spanishSurnames = []string{
	"García", "González", "Rodríguez", "Fernández", "López", "Martínez", "Sánchez", "Pérez", "Gómez", "Martín",
	"Jiménez", "Hernández", "Díaz", "Moreno", "Álvarez", "Romero", "Muñoz", "Ortega", "Ramos", "Rubio",
	"Morales", "Ortiz", "Delgado", "Castro", "Silva", "Torres", "Vargas", "Flores", "Núñez", "Cruz",
	"Molina", "Soto", "Rojas", "Navarro", "Guerrero", "Herrera", "Ramírez", "Medina", "Aguilar", "Vega",
	"Acosta", "Benítez", "Mendoza", "Ríos", "Castillo", "Del Valle", "Peña", "Rivas", "Cabrera", "Campos",
	"Camacho", "Santana", "Cortés", "Reyes", "Cordero", "Sosa", "Del Río", "Velasco", "Ponce", "Estévez",
	"Cano", "Valdez", "Villarreal", "Mora", "Cervantes", "Domínguez", "Villanueva", "Fuentes", "Escobar", "Lara",
}
