
# 1 introduction

golang use UTC(Universal Time Coordinated, 世界协调时间) default. But we usually use CST(China Standard Time, 中国标准时间)。CST is UTC +8:00 in general。

## 1.1 golang use UTC default

```go
    time := time.Now()
    // default UTC
    loc, err := time.LoadLocation("") 
```

## 1.2 golang load location by specified name

```go
    // CST when in China
    loc, err := time.LoadLocation("Local")
    // PDT America Los Angeles
    loc, err := time.LoadLocation("America/Los_Angeles")
    // CST China Chongqing
    loc, _:= time.LoadLocation("Asia/Chongqing")
```

you can find all timezone names in `$GOROOT/lib/time/zoneinfo.zip` showed below.
The tree below show you all the location names which could be recognized by golang `time.LoadLocation` to convert location name to `time.Location`. 

- **location names tree**
```shell
├── Africa
│   ├── Abidjan
│   ├── Accra
│   ├── Addis_Ababa
│   ├── Algiers
│   ├── Asmara
│   ├── Asmera
│   ├── Bamako
│   ├── Bangui
│   ├── Banjul
│   ├── Bissau
│   ├── Blantyre
│   ├── Brazzaville
│   ├── Bujumbura
│   ├── Cairo
│   ├── Casablanca
│   ├── Ceuta
│   ├── Conakry
│   ├── Dakar
│   ├── Dar_es_Salaam
│   ├── Djibouti
│   ├── Douala
│   ├── El_Aaiun
│   ├── Freetown
│   ├── Gaborone
│   ├── Harare
│   ├── Johannesburg
│   ├── Juba
│   ├── Kampala
│   ├── Khartoum
│   ├── Kigali
│   ├── Kinshasa
│   ├── Lagos
│   ├── Libreville
│   ├── Lome
│   ├── Luanda
│   ├── Lubumbashi
│   ├── Lusaka
│   ├── Malabo
│   ├── Maputo
│   ├── Maseru
│   ├── Mbabane
│   ├── Mogadishu
│   ├── Monrovia
│   ├── Nairobi
│   ├── Ndjamena
│   ├── Niamey
│   ├── Nouakchott
│   ├── Ouagadougou
│   ├── Porto-Novo
│   ├── Sao_Tome
│   ├── Timbuktu
│   ├── Tripoli
│   ├── Tunis
│   └── Windhoek
├── America
│   ├── Adak
│   ├── Anchorage
│   ├── Anguilla
│   ├── Antigua
│   ├── Araguaina
│   ├── Argentina
│   │   ├── Buenos_Aires
│   │   ├── Catamarca
│   │   ├── ComodRivadavia
│   │   ├── Cordoba
│   │   ├── Jujuy
│   │   ├── La_Rioja
│   │   ├── Mendoza
│   │   ├── Rio_Gallegos
│   │   ├── Salta
│   │   ├── San_Juan
│   │   ├── San_Luis
│   │   ├── Tucuman
│   │   └── Ushuaia
│   ├── Aruba
│   ├── Asuncion
│   ├── Atikokan
│   ├── Atka
│   ├── Bahia
│   ├── Bahia_Banderas
│   ├── Barbados
│   ├── Belem
│   ├── Belize
│   ├── Blanc-Sablon
│   ├── Boa_Vista
│   ├── Bogota
│   ├── Boise
│   ├── Buenos_Aires
│   ├── Cambridge_Bay
│   ├── Campo_Grande
│   ├── Cancun
│   ├── Caracas
│   ├── Catamarca
│   ├── Cayenne
│   ├── Cayman
│   ├── Chicago
│   ├── Chihuahua
│   ├── Coral_Harbour
│   ├── Cordoba
│   ├── Costa_Rica
│   ├── Creston
│   ├── Cuiaba
│   ├── Curacao
│   ├── Danmarkshavn
│   ├── Dawson
│   ├── Dawson_Creek
│   ├── Denver
│   ├── Detroit
│   ├── Dominica
│   ├── Edmonton
│   ├── Eirunepe
│   ├── El_Salvador
│   ├── Ensenada
│   ├── Fortaleza
│   ├── Fort_Nelson
│   ├── Fort_Wayne
│   ├── Glace_Bay
│   ├── Godthab
│   ├── Goose_Bay
│   ├── Grand_Turk
│   ├── Grenada
│   ├── Guadeloupe
│   ├── Guatemala
│   ├── Guayaquil
│   ├── Guyana
│   ├── Halifax
│   ├── Havana
│   ├── Hermosillo
│   ├── Indiana
│   │   ├── Indianapolis
│   │   ├── Knox
│   │   ├── Marengo
│   │   ├── Petersburg
│   │   ├── Tell_City
│   │   ├── Vevay
│   │   ├── Vincennes
│   │   └── Winamac
│   ├── Indianapolis
│   ├── Inuvik
│   ├── Iqaluit
│   ├── Jamaica
│   ├── Jujuy
│   ├── Juneau
│   ├── Kentucky
│   │   ├── Louisville
│   │   └── Monticello
│   ├── Knox_IN
│   ├── Kralendijk
│   ├── La_Paz
│   ├── Lima
│   ├── Los_Angeles
│   ├── Louisville
│   ├── Lower_Princes
│   ├── Maceio
│   ├── Managua
│   ├── Manaus
│   ├── Marigot
│   ├── Martinique
│   ├── Matamoros
│   ├── Mazatlan
│   ├── Mendoza
│   ├── Menominee
│   ├── Merida
│   ├── Metlakatla
│   ├── Mexico_City
│   ├── Miquelon
│   ├── Moncton
│   ├── Monterrey
│   ├── Montevideo
│   ├── Montreal
│   ├── Montserrat
│   ├── Nassau
│   ├── New_York
│   ├── Nipigon
│   ├── Nome
│   ├── Noronha
│   ├── North_Dakota
│   │   ├── Beulah
│   │   ├── Center
│   │   └── New_Salem
│   ├── Ojinaga
│   ├── Panama
│   ├── Pangnirtung
│   ├── Paramaribo
│   ├── Phoenix
│   ├── Port-au-Prince
│   ├── Porto_Acre
│   ├── Port_of_Spain
│   ├── Porto_Velho
│   ├── Puerto_Rico
│   ├── Punta_Arenas
│   ├── Rainy_River
│   ├── Rankin_Inlet
│   ├── Recife
│   ├── Regina
│   ├── Resolute
│   ├── Rio_Branco
│   ├── Rosario
│   ├── Santa_Isabel
│   ├── Santarem
│   ├── Santiago
│   ├── Santo_Domingo
│   ├── Sao_Paulo
│   ├── Scoresbysund
│   ├── Shiprock
│   ├── Sitka
│   ├── St_Barthelemy
│   ├── St_Johns
│   ├── St_Kitts
│   ├── St_Lucia
│   ├── St_Thomas
│   ├── St_Vincent
│   ├── Swift_Current
│   ├── Tegucigalpa
│   ├── Thule
│   ├── Thunder_Bay
│   ├── Tijuana
│   ├── Toronto
│   ├── Tortola
│   ├── Vancouver
│   ├── Virgin
│   ├── Whitehorse
│   ├── Winnipeg
│   ├── Yakutat
│   └── Yellowknife
├── Antarctica
│   ├── Casey
│   ├── Davis
│   ├── DumontDUrville
│   ├── Macquarie
│   ├── Mawson
│   ├── McMurdo
│   ├── Palmer
│   ├── Rothera
│   ├── South_Pole
│   ├── Syowa
│   ├── Troll
│   └── Vostok
├── Arctic
│   └── Longyearbyen
├── Asia
│   ├── Aden
│   ├── Almaty
│   ├── Amman
│   ├── Anadyr
│   ├── Aqtau
│   ├── Aqtobe
│   ├── Ashgabat
│   ├── Ashkhabad
│   ├── Atyrau
│   ├── Baghdad
│   ├── Bahrain
│   ├── Baku
│   ├── Bangkok
│   ├── Barnaul
│   ├── Beirut
│   ├── Bishkek
│   ├── Brunei
│   ├── Calcutta
│   ├── Chita
│   ├── Choibalsan
│   ├── Chongqing
│   ├── Chungking
│   ├── Colombo
│   ├── Dacca
│   ├── Damascus
│   ├── Dhaka
│   ├── Dili
│   ├── Dubai
│   ├── Dushanbe
│   ├── Famagusta
│   ├── Gaza
│   ├── Harbin
│   ├── Hebron
│   ├── Ho_Chi_Minh
│   ├── Hong_Kong
│   ├── Hovd
│   ├── Irkutsk
│   ├── Istanbul
│   ├── Jakarta
│   ├── Jayapura
│   ├── Jerusalem
│   ├── Kabul
│   ├── Kamchatka
│   ├── Karachi
│   ├── Kashgar
│   ├── Kathmandu
│   ├── Katmandu
│   ├── Khandyga
│   ├── Kolkata
│   ├── Krasnoyarsk
│   ├── Kuala_Lumpur
│   ├── Kuching
│   ├── Kuwait
│   ├── Macao
│   ├── Macau
│   ├── Magadan
│   ├── Makassar
│   ├── Manila
│   ├── Muscat
│   ├── Nicosia
│   ├── Novokuznetsk
│   ├── Novosibirsk
│   ├── Omsk
│   ├── Oral
│   ├── Phnom_Penh
│   ├── Pontianak
│   ├── Pyongyang
│   ├── Qatar
│   ├── Qostanay
│   ├── Qyzylorda
│   ├── Rangoon
│   ├── Riyadh
│   ├── Saigon
│   ├── Sakhalin
│   ├── Samarkand
│   ├── Seoul
│   ├── Shanghai
│   ├── Singapore
│   ├── Srednekolymsk
│   ├── Taipei
│   ├── Tashkent
│   ├── Tbilisi
│   ├── Tehran
│   ├── Tel_Aviv
│   ├── Thimbu
│   ├── Thimphu
│   ├── Tokyo
│   ├── Tomsk
│   ├── Ujung_Pandang
│   ├── Ulaanbaatar
│   ├── Ulan_Bator
│   ├── Urumqi
│   ├── Ust-Nera
│   ├── Vientiane
│   ├── Vladivostok
│   ├── Yakutsk
│   ├── Yangon
│   ├── Yekaterinburg
│   └── Yerevan
├── Atlantic
│   ├── Azores
│   ├── Bermuda
│   ├── Canary
│   ├── Cape_Verde
│   ├── Faeroe
│   ├── Faroe
│   ├── Jan_Mayen
│   ├── Madeira
│   ├── Reykjavik
│   ├── South_Georgia
│   ├── Stanley
│   └── St_Helena
├── Australia
│   ├── ACT
│   ├── Adelaide
│   ├── Brisbane
│   ├── Broken_Hill
│   ├── Canberra
│   ├── Currie
│   ├── Darwin
│   ├── Eucla
│   ├── Hobart
│   ├── LHI
│   ├── Lindeman
│   ├── Lord_Howe
│   ├── Melbourne
│   ├── North
│   ├── NSW
│   ├── Perth
│   ├── Queensland
│   ├── South
│   ├── Sydney
│   ├── Tasmania
│   ├── Victoria
│   ├── West
│   └── Yancowinna
├── Brazil
│   ├── Acre
│   ├── DeNoronha
│   ├── East
│   └── West
├── Canada
│   ├── Atlantic
│   ├── Central
│   ├── Eastern
│   ├── Mountain
│   ├── Newfoundland
│   ├── Pacific
│   ├── Saskatchewan
│   └── Yukon
├── CET
├── Chile
│   ├── Continental
│   └── EasterIsland
├── CST6CDT
├── Cuba
├── EET
├── Egypt
├── Eire
├── EST
├── EST5EDT
├── Etc
│   ├── GMT
│   ├── GMT0
│   ├── GMT-0
│   ├── GMT+0
│   ├── GMT-1
│   ├── GMT+1
│   ├── GMT-10
│   ├── GMT+10
│   ├── GMT-11
│   ├── GMT+11
│   ├── GMT-12
│   ├── GMT+12
│   ├── GMT-13
│   ├── GMT-14
│   ├── GMT-2
│   ├── GMT+2
│   ├── GMT-3
│   ├── GMT+3
│   ├── GMT-4
│   ├── GMT+4
│   ├── GMT-5
│   ├── GMT+5
│   ├── GMT-6
│   ├── GMT+6
│   ├── GMT-7
│   ├── GMT+7
│   ├── GMT-8
│   ├── GMT+8
│   ├── GMT-9
│   ├── GMT+9
│   ├── Greenwich
│   ├── UCT
│   ├── Universal
│   ├── UTC
│   └── Zulu
├── Europe
│   ├── Amsterdam
│   ├── Andorra
│   ├── Astrakhan
│   ├── Athens
│   ├── Belfast
│   ├── Belgrade
│   ├── Berlin
│   ├── Bratislava
│   ├── Brussels
│   ├── Bucharest
│   ├── Budapest
│   ├── Busingen
│   ├── Chisinau
│   ├── Copenhagen
│   ├── Dublin
│   ├── Gibraltar
│   ├── Guernsey
│   ├── Helsinki
│   ├── Isle_of_Man
│   ├── Istanbul
│   ├── Jersey
│   ├── Kaliningrad
│   ├── Kiev
│   ├── Kirov
│   ├── Lisbon
│   ├── Ljubljana
│   ├── London
│   ├── Luxembourg
│   ├── Madrid
│   ├── Malta
│   ├── Mariehamn
│   ├── Minsk
│   ├── Monaco
│   ├── Moscow
│   ├── Nicosia
│   ├── Oslo
│   ├── Paris
│   ├── Podgorica
│   ├── Prague
│   ├── Riga
│   ├── Rome
│   ├── Samara
│   ├── San_Marino
│   ├── Sarajevo
│   ├── Saratov
│   ├── Simferopol
│   ├── Skopje
│   ├── Sofia
│   ├── Stockholm
│   ├── Tallinn
│   ├── Tirane
│   ├── Tiraspol
│   ├── Ulyanovsk
│   ├── Uzhgorod
│   ├── Vaduz
│   ├── Vatican
│   ├── Vienna
│   ├── Vilnius
│   ├── Volgograd
│   ├── Warsaw
│   ├── Zagreb
│   ├── Zaporozhye
│   └── Zurich
├── Factory
├── GB
├── GB-Eire
├── GMT
├── GMT0
├── GMT-0
├── GMT+0
├── Greenwich
├── Hongkong
├── HST
├── Iceland
├── Indian
│   ├── Antananarivo
│   ├── Chagos
│   ├── Christmas
│   ├── Cocos
│   ├── Comoro
│   ├── Kerguelen
│   ├── Mahe
│   ├── Maldives
│   ├── Mauritius
│   ├── Mayotte
│   └── Reunion
├── Iran
├── Israel
├── Jamaica
├── Japan
├── Kwajalein
├── Libya
├── locationnames.md
├── MET
├── Mexico
│   ├── BajaNorte
│   ├── BajaSur
│   └── General
├── MST
├── MST7MDT
├── Navajo
├── NZ
├── NZ-CHAT
├── Pacific
│   ├── Apia
│   ├── Auckland
│   ├── Bougainville
│   ├── Chatham
│   ├── Chuuk
│   ├── Easter
│   ├── Efate
│   ├── Enderbury
│   ├── Fakaofo
│   ├── Fiji
│   ├── Funafuti
│   ├── Galapagos
│   ├── Gambier
│   ├── Guadalcanal
│   ├── Guam
│   ├── Honolulu
│   ├── Johnston
│   ├── Kiritimati
│   ├── Kosrae
│   ├── Kwajalein
│   ├── Majuro
│   ├── Marquesas
│   ├── Midway
│   ├── Nauru
│   ├── Niue
│   ├── Norfolk
│   ├── Noumea
│   ├── Pago_Pago
│   ├── Palau
│   ├── Pitcairn
│   ├── Pohnpei
│   ├── Ponape
│   ├── Port_Moresby
│   ├── Rarotonga
│   ├── Saipan
│   ├── Samoa
│   ├── Tahiti
│   ├── Tarawa
│   ├── Tongatapu
│   ├── Truk
│   ├── Wake
│   ├── Wallis
│   └── Yap
├── Poland
├── Portugal
├── PRC
├── PST8PDT
├── ROC
├── ROK
├── Singapore
├── Turkey
├── UCT
├── Universal
├── US
│   ├── Alaska
│   ├── Aleutian
│   ├── Arizona
│   ├── Central
│   ├── Eastern
│   ├── East-Indiana
│   ├── Hawaii
│   ├── Indiana-Starke
│   ├── Michigan
│   ├── Mountain
│   ├── Pacific
│   └── Samoa
├── UTC
├── WET
├── W-SU
└── Zulu

```

# 2 Example

If you want to get Location China, you could use 

```go
    location, _ := time.LoadLocation("Asia/Shanghai")
    value := "2019-08-12 19:57:23"
    pattern := "2006-01-02 15:04:05"
    result, _ := ParseTimeInLocation(pattern, value, location)

```
