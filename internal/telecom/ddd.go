package telecom

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"leadphone-validator/internal/models"
)

var dddDatabase = []models.DDDInfo{
	{Code: "11", State: "SP", StateName: "São Paulo", Region: "Sudeste", MajorCities: []string{"São Paulo", "Guarulhos", "Santo André", "São Bernardo do Campo", "Osasco", "Mogi das Cruzes", "Barueri", "Jundiaí"}, Timezone: "America/Sao_Paulo"},
	{Code: "12", State: "SP", StateName: "São Paulo", Region: "Sudeste", MajorCities: []string{"São José dos Campos", "Taubaté", "Jacareí", "Pindamonhangaba", "Guaratinguetá", "Caraguatatuba", "Ubatuba", "Ilhabela"}, Timezone: "America/Sao_Paulo"},
	{Code: "13", State: "SP", StateName: "São Paulo", Region: "Sudeste", MajorCities: []string{"Santos", "São Vicente", "Praia Grande", "Guarujá", "Cubatão", "Itanhaém", "Peruíbe", "Registro"}, Timezone: "America/Sao_Paulo"},
	{Code: "14", State: "SP", StateName: "São Paulo", Region: "Sudeste", MajorCities: []string{"Bauru", "Marília", "Jaú", "Botucatu", "Ourinhos", "Avaré", "Lençóis Paulista", "Lins"}, Timezone: "America/Sao_Paulo"},
	{Code: "15", State: "SP", StateName: "São Paulo", Region: "Sudeste", MajorCities: []string{"Sorocaba", "Itapetininga", "Itapeva", "Tatuí", "Votorantim", "Salto de Pirapora", "Capão Bonito"}, Timezone: "America/Sao_Paulo"},
	{Code: "16", State: "SP", StateName: "São Paulo", Region: "Sudeste", MajorCities: []string{"Ribeirão Preto", "Franca", "Araraquara", "São Carlos", "Sertãozinho", "Jaboticabal", "Matão", "Bebedouro"}, Timezone: "America/Sao_Paulo"},
	{Code: "17", State: "SP", StateName: "São Paulo", Region: "Sudeste", MajorCities: []string{"São José do Rio Preto", "Catanduva", "Barretos", "Votuporanga", "Fernandópolis", "Jales", "Olímpia"}, Timezone: "America/Sao_Paulo"},
	{Code: "18", State: "SP", StateName: "São Paulo", Region: "Sudeste", MajorCities: []string{"Presidente Prudente", "Araçatuba", "Birigui", "Assis", "Adamantina", "Dracena", "Penápolis"}, Timezone: "America/Sao_Paulo"},
	{Code: "19", State: "SP", StateName: "São Paulo", Region: "Sudeste", MajorCities: []string{"Campinas", "Piracicaba", "Limeira", "Sumaré", "Americana", "Hortolândia", "Rio Claro", "Indaiatuba", "Santa Bárbara d'Oeste"}, Timezone: "America/Sao_Paulo"},
	{Code: "21", State: "RJ", StateName: "Rio de Janeiro", Region: "Sudeste", MajorCities: []string{"Rio de Janeiro", "Niterói", "São Gonçalo", "Duque de Caxias", "Nova Iguaçu", "Belford Roxo", "São João de Meriti", "Magé"}, Timezone: "America/Sao_Paulo"},
	{Code: "22", State: "RJ", StateName: "Rio de Janeiro", Region: "Sudeste", MajorCities: []string{"Campos dos Goytacazes", "Macaé", "Cabo Frio", "Nova Friburgo", "Rio das Ostras", "Araruama", "Itaperuna", "Armação dos Búzios"}, Timezone: "America/Sao_Paulo"},
	{Code: "24", State: "RJ", StateName: "Rio de Janeiro", Region: "Sudeste", MajorCities: []string{"Volta Redonda", "Petrópolis", "Barra Mansa", "Angra dos Reis", "Teresópolis", "Resende", "Três Rios", "Paraty"}, Timezone: "America/Sao_Paulo"},
	{Code: "27", State: "ES", StateName: "Espírito Santo", Region: "Sudeste", MajorCities: []string{"Vitória", "Vila Velha", "Serra", "Cariacica", "Linhares", "Colatina", "Guarapari", "Aracruz", "São Mateus"}, Timezone: "America/Sao_Paulo"},
	{Code: "28", State: "ES", StateName: "Espírito Santo", Region: "Sudeste", MajorCities: []string{"Cachoeiro de Itapemirim", "Marataízes", "Castelo", "Alegre", "Itapemirim", "Guaçuí", "Iúna"}, Timezone: "America/Sao_Paulo"},
	{Code: "31", State: "MG", StateName: "Minas Gerais", Region: "Sudeste", MajorCities: []string{"Belo Horizonte", "Contagem", "Betim", "Ribeirão das Neves", "Santa Luzia", "Sete Lagoas", "Ibirité", "Ipatinga", "Coronel Fabriciano"}, Timezone: "America/Sao_Paulo"},
	{Code: "32", State: "MG", StateName: "Minas Gerais", Region: "Sudeste", MajorCities: []string{"Juiz de Fora", "Barbacena", "Ubá", "Muriaé", "São João del-Rei", "Cataguases", "Viçosa", "Ponte Nova"}, Timezone: "America/Sao_Paulo"},
	{Code: "33", State: "MG", StateName: "Minas Gerais", Region: "Sudeste", MajorCities: []string{"Governador Valadares", "Teófilo Otoni", "Caratinga", "Manhuaçu", "Nanuque", "Almenara", "Capelinha"}, Timezone: "America/Sao_Paulo"},
	{Code: "34", State: "MG", StateName: "Minas Gerais", Region: "Sudeste", MajorCities: []string{"Uberlândia", "Uberaba", "Patos de Minas", "Araguari", "Araxá", "Ituiutaba", "Frutal", "Patrocínio"}, Timezone: "America/Sao_Paulo"},
	{Code: "35", State: "MG", StateName: "Minas Gerais", Region: "Sudeste", MajorCities: []string{"Poços de Caldas", "Pouso Alegre", "Varginha", "Passos", "Lavras", "Itajubá", "Alfenas", "Três Corações"}, Timezone: "America/Sao_Paulo"},
	{Code: "37", State: "MG", StateName: "Minas Gerais", Region: "Sudeste", MajorCities: []string{"Divinópolis", "Itaúna", "Nova Serrana", "Pará de Minas", "Formiga", "Bom Despacho", "Campo Belo"}, Timezone: "America/Sao_Paulo"},
	{Code: "38", State: "MG", StateName: "Minas Gerais", Region: "Sudeste", MajorCities: []string{"Montes Claros", "Curvelo", "Diamantina", "Januária", "Pirapora", "Unaí", "Paracatu", "Salinas"}, Timezone: "America/Sao_Paulo"},
	{Code: "41", State: "PR", StateName: "Paraná", Region: "Sul", MajorCities: []string{"Curitiba", "São José dos Pinhais", "Colombo", "Araucária", "Pinhais", "Campo Largo", "Paranaguá", "Fazenda Rio Grande"}, Timezone: "America/Sao_Paulo"},
	{Code: "42", State: "PR", StateName: "Paraná", Region: "Sul", MajorCities: []string{"Ponta Grossa", "Guarapuava", "Castro", "Telêmaco Borba", "Irati", "União da Vitória", "Prudentópolis"}, Timezone: "America/Sao_Paulo"},
	{Code: "43", State: "PR", StateName: "Paraná", Region: "Sul", MajorCities: []string{"Londrina", "Apucarana", "Arapongas", "Cambé", "Rolândia", "Cornélio Procópio", "Jacarezinho", "Ivaiporã"}, Timezone: "America/Sao_Paulo"},
	{Code: "44", State: "PR", StateName: "Paraná", Region: "Sul", MajorCities: []string{"Maringá", "Campo Mourão", "Paranavaí", "Cianorte", "Umuarama", "Sarandi", "Paiçandu", "Loanda"}, Timezone: "America/Sao_Paulo"},
	{Code: "45", State: "PR", StateName: "Paraná", Region: "Sul", MajorCities: []string{"Cascavel", "Foz do Iguaçu", "Toledo", "Medianeira", "Marechal Cândido Rondon", "Santa Helena", "Matelândia"}, Timezone: "America/Sao_Paulo"},
	{Code: "46", State: "PR", StateName: "Paraná", Region: "Sul", MajorCities: []string{"Francisco Beltrão", "Pato Branco", "Palmas", "Dois Vizinhos", "Chopinzinho", "Coronel Vivida"}, Timezone: "America/Sao_Paulo"},
	{Code: "47", State: "SC", StateName: "Santa Catarina", Region: "Sul", MajorCities: []string{"Joinville", "Blumenau", "Itajaí", "Balneário Camboriú", "Jaraguá do Sul", "Brusque", "Navegantes", "Camboriú", "São Bento do Sul"}, Timezone: "America/Sao_Paulo"},
	{Code: "48", State: "SC", StateName: "Santa Catarina", Region: "Sul", MajorCities: []string{"Florianópolis", "São José", "Palhoça", "Criciúma", "Tubarão", "Biguaçu", "Içara", "Laguna", "Imbituba", "Araranguá"}, Timezone: "America/Sao_Paulo"},
	{Code: "49", State: "SC", StateName: "Santa Catarina", Region: "Sul", MajorCities: []string{"Chapecó", "Lages", "Caçador", "Concórdia", "Joaçaba", "São Miguel do Oeste", "Videira", "Xanxerê", "Curitibanos"}, Timezone: "America/Sao_Paulo"},
	{Code: "51", State: "RS", StateName: "Rio Grande do Sul", Region: "Sul", MajorCities: []string{"Porto Alegre", "Canoas", "Novo Hamburgo", "São Leopoldo", "Gravataí", "Viamão", "Alvorada", "Cachoeirinha", "Santa Cruz do Sul", "Torres"}, Timezone: "America/Sao_Paulo"},
	{Code: "53", State: "RS", StateName: "Rio Grande do Sul", Region: "Sul", MajorCities: []string{"Pelotas", "Rio Grande", "Bagé", "Santana do Livramento", "Camaquã", "Jaguarão", "São Lourenço do Sul"}, Timezone: "America/Sao_Paulo"},
	{Code: "54", State: "RS", StateName: "Rio Grande do Sul", Region: "Sul", MajorCities: []string{"Caxias do Sul", "Passo Fundo", "Bento Gonçalves", "Erechim", "Farroupilha", "Vacaria", "Gramado", "Canela", "Carazinho"}, Timezone: "America/Sao_Paulo"},
	{Code: "55", State: "RS", StateName: "Rio Grande do Sul", Region: "Sul", MajorCities: []string{"Santa Maria", "Uruguaiana", "Ijuí", "Santo Ângelo", "Cruz Alta", "Alegrete", "São Borja", "Santa Rosa"}, Timezone: "America/Sao_Paulo"},
	{Code: "61", State: "DF", StateName: "Distrito Federal", Region: "Centro-Oeste", MajorCities: []string{"Brasília", "Ceilândia", "Taguatinga", "Samambaia", "Plano Piloto", "Águas Claras", "Luziânia", "Valparaíso de Goiás", "Formosa"}, Timezone: "America/Sao_Paulo"},
	{Code: "62", State: "GO", StateName: "Goiás", Region: "Centro-Oeste", MajorCities: []string{"Goiânia", "Aparecida de Goiânia", "Anápolis", "Trindade", "Senador Canedo", "Itumbiara", "Caldas Novas", "Catalão", "Jataí"}, Timezone: "America/Sao_Paulo"},
	{Code: "63", State: "TO", StateName: "Tocantins", Region: "Norte", MajorCities: []string{"Palmas", "Araguaína", "Gurupi", "Porto Nacional", "Paraíso do Tocantins", "Colinas do Tocantins"}, Timezone: "America/Araguaina"},
	{Code: "64", State: "GO", StateName: "Goiás", Region: "Centro-Oeste", MajorCities: []string{"Rio Verde", "Itumbiara", "Jataí", "Caldas Novas", "Morrinhos", "Mineiros", "Quirinópolis", "Santa Helena de Goiás"}, Timezone: "America/Sao_Paulo"},
	{Code: "65", State: "MT", StateName: "Mato Grosso", Region: "Centro-Oeste", MajorCities: []string{"Cuiabá", "Várzea Grande", "Tangará da Serra", "Cáceres", "Poconé", "Barra do Bugres", "Lucas do Rio Verde"}, Timezone: "America/Cuiaba"},
	{Code: "66", State: "MT", StateName: "Mato Grosso", Region: "Centro-Oeste", MajorCities: []string{"Rondonópolis", "Sinop", "Sorriso", "Primavera do Leste", "Barra do Garças", "Alta Floresta", "Nova Mutum"}, Timezone: "America/Cuiaba"},
	{Code: "67", State: "MS", StateName: "Mato Grosso do Sul", Region: "Centro-Oeste", MajorCities: []string{"Campo Grande", "Dourados", "Três Lagoas", "Corumbá", "Ponta Porã", "Naviraí", "Nova Andradina", "Sidrolândia"}, Timezone: "America/Campo_Grande"},
	{Code: "68", State: "AC", StateName: "Acre", Region: "Norte", MajorCities: []string{"Rio Branco", "Cruzeiro do Sul", "Sena Madureira", "Tarauacá", "Feijó", "Brasiléia"}, Timezone: "America/Rio_Branco"},
	{Code: "69", State: "RO", StateName: "Rondônia", Region: "Norte", MajorCities: []string{"Porto Velho", "Ji-Paraná", "Ariquemes", "Vilhena", "Cacoal", "Jaru", "Rolim de Moura", "Guajará-Mirim"}, Timezone: "America/Porto_Velho"},
	{Code: "71", State: "BA", StateName: "Bahia", Region: "Nordeste", MajorCities: []string{"Salvador", "Camaçari", "Lauro de Freitas", "Simões Filho", "Candeias", "Dias d'Ávila", "Mata de São João"}, Timezone: "America/Bahia"},
	{Code: "73", State: "BA", StateName: "Bahia", Region: "Nordeste", MajorCities: []string{"Ilhéus", "Itabuna", "Porto Seguro", "Teixeira de Freitas", "Eunápolis", "Valença", "Jequié", "Canavieiras"}, Timezone: "America/Bahia"},
	{Code: "74", State: "BA", StateName: "Bahia", Region: "Nordeste", MajorCities: []string{"Juazeiro", "Jacobina", "Senhor do Bonfim", "Irecê", "Campo Formoso", "Casa Nova", "Xique-Xique"}, Timezone: "America/Bahia"},
	{Code: "75", State: "BA", StateName: "Bahia", Region: "Nordeste", MajorCities: []string{"Feira de Santana", "Alagoinhas", "Serrinha", "Santo Antônio de Jesus", "Cruz das Almas", "Valente", "Conceição do Coité"}, Timezone: "America/Bahia"},
	{Code: "77", State: "BA", StateName: "Bahia", Region: "Nordeste", MajorCities: []string{"Vitória da Conquista", "Barreiras", "Guanambi", "Luís Eduardo Magalhães", "Brumado", "Bom Jesus da Lapa", "Itapetinga"}, Timezone: "America/Bahia"},
	{Code: "79", State: "SE", StateName: "Sergipe", Region: "Nordeste", MajorCities: []string{"Aracaju", "Nossa Senhora do Socorro", "Lagarto", "Itabaiana", "São Cristóvão", "Estância", "Propriá"}, Timezone: "America/Maceio"},
	{Code: "81", State: "PE", StateName: "Pernambuco", Region: "Nordeste", MajorCities: []string{"Recife", "Jaboatão dos Guararapes", "Olinda", "Paulista", "Caruaru", "Cabo de Santo Agostinho", "Camaragibe", "Garanhuns", "Vitória de Santo Antão"}, Timezone: "America/Recife"},
	{Code: "82", State: "AL", StateName: "Alagoas", Region: "Nordeste", MajorCities: []string{"Maceió", "Arapiraca", "Rio Largo", "Palmeira dos Índios", "União dos Palmares", "Penedo", "São Miguel dos Campos"}, Timezone: "America/Maceio"},
	{Code: "83", State: "PB", StateName: "Paraíba", Region: "Nordeste", MajorCities: []string{"João Pessoa", "Campina Grande", "Santa Rita", "Patos", "Bayeux", "Sousa", "Cajazeiras", "Cabedelo", "Guarabira"}, Timezone: "America/Fortaleza"},
	{Code: "84", State: "RN", StateName: "Rio Grande do Norte", Region: "Nordeste", MajorCities: []string{"Natal", "Mossoró", "Parnamirim", "São Gonçalo do Amarante", "Ceará-Mirim", "Caicó", "Açu", "Macau"}, Timezone: "America/Fortaleza"},
	{Code: "85", State: "CE", StateName: "Ceará", Region: "Nordeste", MajorCities: []string{"Fortaleza", "Caucaia", "Maracanaú", "Maranguape", "Aquiraz", "Cascavel", "Pacatuba", "Itapipoca", "Horizonte"}, Timezone: "America/Fortaleza"},
	{Code: "86", State: "PI", StateName: "Piauí", Region: "Nordeste", MajorCities: []string{"Teresina", "Parnaíba", "Piripiri", "Campo Maior", "Barras", "União", "Altos", "Esperantina"}, Timezone: "America/Fortaleza"},
	{Code: "87", State: "PE", StateName: "Pernambuco", Region: "Nordeste", MajorCities: []string{"Petrolina", "Caruaru", "Garanhuns", "Serra Talhada", "Arcoverde", "Salgueiro", "Ouricuri", "Afogados da Ingazeira"}, Timezone: "America/Recife"},
	{Code: "88", State: "CE", StateName: "Ceará", Region: "Nordeste", MajorCities: []string{"Juazeiro do Norte", "Sobral", "Crato", "Itapipoca", "Iguatu", "Quixadá", "Quixeramobim", "Russas", "Canindé", "Crateús"}, Timezone: "America/Fortaleza"},
	{Code: "89", State: "PI", StateName: "Piauí", Region: "Nordeste", MajorCities: []string{"Picos", "Floriano", "São Raimundo Nonato", "Oeiras", "Corrente", "Bom Jesus", "Paulistana"}, Timezone: "America/Fortaleza"},
	{Code: "91", State: "PA", StateName: "Pará", Region: "Norte", MajorCities: []string{"Belém", "Ananindeua", "Castanhal", "Abaetetuba", "Marituba", "Barcarena", "Bragança", "Capanema", "Santa Izabel do Pará"}, Timezone: "America/Belem"},
	{Code: "92", State: "AM", StateName: "Amazonas", Region: "Norte", MajorCities: []string{"Manaus", "Parintins", "Itacoatiara", "Manacapuru", "Coari", "Maués", "Tefé", "Iranduba"}, Timezone: "America/Manaus"},
	{Code: "93", State: "PA", StateName: "Pará", Region: "Norte", MajorCities: []string{"Santarém", "Altamira", "Itaituba", "Oriximiná", "Alenquer", "Monte Alegre", "Óbidos"}, Timezone: "America/Santarem"},
	{Code: "94", State: "PA", StateName: "Pará", Region: "Norte", MajorCities: []string{"Marabá", "Parauapebas", "Tucuruí", "Redenção", "Canaã dos Carajás", "Xinguara", "Ourilândia do Norte"}, Timezone: "America/Belem"},
	{Code: "95", State: "RR", StateName: "Roraima", Region: "Norte", MajorCities: []string{"Boa Vista", "Rorainópolis", "Caracaraí", "Pacaraima", "Cantá", "Mucajaí"}, Timezone: "America/Boa_Vista"},
	{Code: "96", State: "AP", StateName: "Amapá", Region: "Norte", MajorCities: []string{"Macapá", "Santana", "Laranjal do Jari", "Oiapoque", "Porto Grande", "Mazagão"}, Timezone: "America/Belem"},
	{Code: "97", State: "AM", StateName: "Amazonas", Region: "Norte", MajorCities: []string{"Tefé", "Tabatinga", "Coari", "Lábrea", "Eirunepé", "São Gabriel da Cachoeira", "Humaitá", "Borba"}, Timezone: "America/Manaus"},
	{Code: "98", State: "MA", StateName: "Maranhão", Region: "Nordeste", MajorCities: []string{"São Luís", "São José de Ribamar", "Paço do Lumiar", "Pinheiro", "Santa Inês", "Chapadinha", "Itapecuru Mirim"}, Timezone: "America/Fortaleza"},
	{Code: "99", State: "MA", StateName: "Maranhão", Region: "Nordeste", MajorCities: []string{"Imperatriz", "Caxias", "Timon", "Codó", "Açailândia", "Bacabal", "Balsas", "Barra do Corda", "Grajaú"}, Timezone: "America/Fortaleza"},
}

func stripAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return strings.ToLower(strings.TrimSpace(result))
}

func LookupDDD(code string) (models.DDDInfo, bool) {
	cleanCode := strings.TrimSpace(code)
	for _, item := range dddDatabase {
		if item.Code == cleanCode {
			return item, true
		}
	}
	return models.DDDInfo{}, false
}

func SearchDDD(query string) []models.DDDInfo {
	q := stripAccents(query)
	if q == "" {
		return dddDatabase
	}

	var results []models.DDDInfo
	for _, item := range dddDatabase {
		if item.Code == q || strings.ToLower(item.State) == q || stripAccents(item.StateName) == q {
			results = append(results, item)
			continue
		}

		matchedCity := false
		for _, city := range item.MajorCities {
			if strings.Contains(stripAccents(city), q) {
				matchedCity = true
				break
			}
		}

		if matchedCity {
			results = append(results, item)
		}
	}

	return results
}
