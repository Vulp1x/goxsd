// generated by goxsd; DO NOT EDIT

package f311sfc0v512

// Файл is generated from an XSD element.
type Файл struct {
	ИдФайл    string   `xml:"ИдФайл,attr"`
	ТипИнф    string   `xml:"ТипИнф,attr"`
	ВерсПрог  string   `xml:"ВерсПрог,attr"`
	ТелОтпр   string   `xml:"ТелОтпр,attr"`
	ДолжнОтпр string   `xml:"ДолжнОтпр,attr"`
	ФамОтпр   string   `xml:"ФамОтпр,attr"`
	КолДок    string   `xml:"КолДок,attr"`
	ВерсФорм  string   `xml:"ВерсФорм,attr"`
	Документ  Документ `xml:"Документ"`
}

// Документ is generated from an XSD element.
type Документ struct {
	ИдДок    string   `xml:"ИдДок,attr"`
	КНД      string   `xml:"КНД,attr"`
	КодНОБ   string   `xml:"КодНОБ,attr"`
	НомСооб  string   `xml:"НомСооб,attr"`
	ТипСооб  string   `xml:"ТипСооб,attr"`
	ДолжнПрБ string   `xml:"ДолжнПрБ,attr"`
	ФамПрБ   string   `xml:"ФамПрБ,attr"`
	ТелБанка string   `xml:"ТелБанка,attr"`
	ДатаСооб string   `xml:"ДатаСооб,attr"`
	СвБанк   СвБанк   `xml:"СвБанк"`
	СвНП     СвНП     `xml:"СвНП"`
	СвСчет   СвСчет   `xml:"СвСчет"`
	Совлад   []Совлад `xml:"Совлад"`
}

// СвБанк is generated from an XSD element.
type СвБанк struct {
	РегНом string `xml:"РегНом,attr"`
	НомФ   string `xml:"НомФ,attr"`
	БИК    string `xml:"БИК,attr"`
	НаимКО string `xml:"НаимКО,attr"`
	ИННКО  string `xml:"ИННКО,attr"`
	КППКО  string `xml:"КППКО,attr"`
	ОГРНКО string `xml:"ОГРНКО,attr"`
}

// СвНП is generated from an XSD element.
type СвНП struct {
	КодЛица string `xml:"КодЛица,attr"`
	НПФЛ    НПФЛ   `xml:"НПФЛ"`
}

// НПФЛ is generated from an XSD element.
type НПФЛ struct {
	ИННФЛ     string `xml:"ИННФЛ,attr"`
	ДатаРожд  string `xml:"ДатаРожд,attr"`
	МестоРожд string `xml:"МестоРожд,attr"`
	КодДУЛ    string `xml:"КодДУЛ,attr"`
	СерНомДок string `xml:"СерНомДок,attr"`
	ДатаДок   string `xml:"ДатаДок,attr"`
	ФИОФЛ     ФИОФЛ  `xml:"ФИОФЛ"`
}

// ФИОФЛ is generated from an XSD element.
type ФИОФЛ struct {
	Фамилия  string `xml:"Фамилия,attr"`
	Имя      string `xml:"Имя,attr"`
	Отчество string `xml:"Отчество,attr"`
}

// СвСчет is generated from an XSD element.
type СвСчет struct {
	НомСч      string `xml:"НомСч,attr"`
	ДатаОткрСч string `xml:"ДатаОткрСч,attr"`
	КодСч      string `xml:"КодСч,attr"`
	ВалСч      string `xml:"ВалСч,attr"`
	НазнСч     string `xml:"НазнСч,attr"`
	КолСовл    string `xml:"КолСовл,attr"`
}

// Совлад is generated from an XSD element.
type Совлад struct {
	КодЛица string `xml:"КодЛица,attr"`
	НПФЛ    НПФЛ   `xml:"НПФЛ"`
}
