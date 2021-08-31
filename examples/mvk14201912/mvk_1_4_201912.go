// generated by goxsd; DO NOT EDIT

package mvk14201912

import "time"

type (
	AnyURIXML             string
	BooleanXML            bool
	DateTimeXSDType       time.Time
	DateXSDType           time.Time
	DecimalXML            float64
	DurationXML           time.Duration
	Float64XML            float64
	IntXML                int
	IntegerXML            int
	LanguageXML           string
	LongIntXML            int64
	NameXML               string
	NonNegativeIntegerXML uint
	NormalizedStringXML   string
	PositiveIntegerXML    uint
	ShortIntXML           int16
	StringXML             string
	TimeXSDType           time.Time
	TokenXML              string
	UnsignedShortIntXML   uint16
)

// СписокРешений is generated from an XSD element.
type СписокРешений struct {
	ВерсияФормата           StringXML                   `xml:"ВерсияФормата"`
	ДатаСписка              DateXSDType                 `xml:"ДатаСписка"`
	ДатаПредыдущегоСписка   *DateXSDType                `xml:"ДатаПредыдущегоСписка,omitempty"`
	СписокАктуальныхРешений СписокАктуальныхРешенийТип  `xml:"СписокАктуальныхРешений"`
	СписокОтмененныхРешений *СписокОтмененныхРешенийТип `xml:"СписокОтмененныхРешений,omitempty"`
}

// СписокАктуальныхРешенийТип is generated from an XSD element.
type СписокАктуальныхРешенийТип struct {
	Решение []РешениеТип `xml:"Решение"`
}

// РешениеТип is generated from an XSD element.
type РешениеТип struct {
	ТипРешения      СправочникТип       `xml:"ТипРешения"`
	НомерРешения    StringXML           `xml:"НомерРешения"`
	ДатаРешения     DateXSDType         `xml:"ДатаРешения"`
	Орган           StringXML           `xml:"Орган"`
	ВидРешения      СправочникТип       `xml:"ВидРешения"`
	СписокСубъектов *СписокСубъектовТип `xml:"СписокСубъектов,omitempty"`
}

// СправочникТип is generated from an XSD element.
type СправочникТип struct {
	Идентификатор LongIntXML `xml:"Идентификатор"`
	Наименование  StringXML  `xml:"Наименование"`
}

// СписокСубъектовТип is generated from an XSD element.
type СписокСубъектовТип struct {
	Субъект []СубъектТип `xml:"Субъект,omitempty"`
}

// СубъектТип is generated from an XSD element.
type СубъектТип struct {
	ИдСубъекта        StringXML         `xml:"ИдСубъекта"`
	ТипСубъекта       СправочникТип     `xml:"ТипСубъекта"`
	СписокАдресов     *СписокАдресовТип `xml:"СписокАдресов,omitempty"`
	РешениеПоСубъекту *StringXML        `xml:"РешениеПоСубъекту,omitempty"`
	Примечание        *StringXML        `xml:"Примечание,omitempty"`
	ФЛ                *ФЛТип            `xml:"ФЛ,omitempty"`
	ЮЛ                *ЮЛТип            `xml:"ЮЛ,omitempty"`
}

// СписокАдресовТип is generated from an XSD element.
type СписокАдресовТип struct {
	Адрес []АдресТип `xml:"Адрес"`
}

// АдресТип is generated from an XSD element.
type АдресТип struct {
	НормАдрес            *СправочникТип `xml:"НормАдрес,omitempty"`
	ТипАдреса            СправочникТип  `xml:"ТипАдреса"`
	ТекстАдреса          StringXML      `xml:"ТекстАдреса"`
	ИдентификаторФИАС    *StringXML     `xml:"ИдентификаторФИАС,omitempty"`
	УровеньАдреснОбъекта *IntXML        `xml:"УровеньАдреснОбъекта,omitempty"`
	Страна               *СтранаТип     `xml:"Страна,omitempty"`
	Индекс               *StringXML     `xml:"Индекс,omitempty"`
	ОКАТО                *StringXML     `xml:"ОКАТО,omitempty"`
	Регион               *StringXML     `xml:"Регион,omitempty"`
	АвтономныйОкруг      *StringXML     `xml:"АвтономныйОкруг,omitempty"`
	Район                *StringXML     `xml:"Район,omitempty"`
	Город                *StringXML     `xml:"Город,omitempty"`
	ВнутригородскРайон   *StringXML     `xml:"ВнутригородскРайон,omitempty"`
	НаселПункт           *StringXML     `xml:"НаселПункт,omitempty"`
	Улица                *StringXML     `xml:"Улица,omitempty"`
	Дом                  *StringXML     `xml:"Дом,omitempty"`
	Корпус               *StringXML     `xml:"Корпус,omitempty"`
	Строение             *StringXML     `xml:"Строение,omitempty"`
	Помещение            *StringXML     `xml:"Помещение,omitempty"`
}

// СтранаТип is generated from an XSD element.
type СтранаТип struct {
	Наименование StringXML  `xml:"Наименование"`
	Код          *StringXML `xml:"Код,omitempty"`
}

// ФЛТип is generated from an XSD element.
type ФЛТип struct {
	ФИО                  StringXML                  `xml:"ФИО"`
	Фамилия              *StringXML                 `xml:"Фамилия,omitempty"`
	Имя                  *StringXML                 `xml:"Имя,omitempty"`
	Отчество             *StringXML                 `xml:"Отчество,omitempty"`
	ФИОЛат               *StringXML                 `xml:"ФИОЛат,omitempty"`
	ДатаРождения         *DateXSDType               `xml:"ДатаРождения,omitempty"`
	ГодРождения          *StringXML                 `xml:"ГодРождения,omitempty"`
	МестоРождения        *StringXML                 `xml:"МестоРождения,omitempty"`
	ИНН                  *StringXML                 `xml:"ИНН,omitempty"`
	СНИЛС                *StringXML                 `xml:"СНИЛС,omitempty"`
	СписокДокументов     *СписокДокументовТип       `xml:"СписокДокументов,omitempty"`
	СписокДрНаименований *СписокДрНаименованийФЛТип `xml:"СписокДрНаименований,omitempty"`
}

// СписокДокументовТип is generated from an XSD element.
type СписокДокументовТип struct {
	Документ []ДокументТип `xml:"Документ,omitempty"`
}

// ДокументТип is generated from an XSD element.
type ДокументТип struct {
	ТипДокумента            СправочникТип `xml:"ТипДокумента"`
	Серия                   *StringXML    `xml:"Серия,omitempty"`
	Номер                   *StringXML    `xml:"Номер,omitempty"`
	ОрганВыдачи             *StringXML    `xml:"ОрганВыдачи,omitempty"`
	ДатаВыдачи              *DateXSDType  `xml:"ДатаВыдачи,omitempty"`
	ДатаС                   *DateXSDType  `xml:"ДатаС,omitempty"`
	ДатаПо                  *DateXSDType  `xml:"ДатаПо,omitempty"`
	ПризнакДействительности *BooleanXML   `xml:"ПризнакДействительности,omitempty"`
}

// СписокДрНаименованийФЛТип is generated from an XSD element.
type СписокДрНаименованийФЛТип struct {
	ДрНаименование []ДрНаименованиеФЛТип `xml:"ДрНаименование,omitempty"`
}

// ДрНаименованиеФЛТип is generated from an XSD element.
type ДрНаименованиеФЛТип struct{}

// ЮЛТип is generated from an XSD element.
type ЮЛТип struct {
	НаименованиеЛат      *StringXML                 `xml:"НаименованиеЛат,omitempty"`
	ДатаРегистрации      *DateXSDType               `xml:"ДатаРегистрации,omitempty"`
	МестоРегистрации     *StringXML                 `xml:"МестоРегистрации,omitempty"`
	ОГРН                 *StringXML                 `xml:"ОГРН,omitempty"`
	СписокДрНаименований *СписокДрНаименованийЮЛТип `xml:"СписокДрНаименований,omitempty"`
}

// СписокДрНаименованийЮЛТип is generated from an XSD element.
type СписокДрНаименованийЮЛТип struct {
	ДрНаименование []ДрНаименованиеЮЛТип `xml:"ДрНаименование"`
}

// ДрНаименованиеЮЛТип is generated from an XSD element.
type ДрНаименованиеЮЛТип struct{}

// СписокОтмененныхРешенийТип is generated from an XSD element.
type СписокОтмененныхРешенийТип struct {
	ОтмененноеРешение []РешениеТип `xml:"ОтмененноеРешение,omitempty"`
}
