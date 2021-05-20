// generated by goxsd; DO NOT EDIT

package mvk14201912

import "time"

// СписокРешений is generated from an XSD element.
type СписокРешений struct {
	ВерсияФормата           string                  `xml:"ВерсияФормата"`
	ДатаСписка              time.Time               `xml:"ДатаСписка"`
	ДатаПредыдущегоСписка   time.Time               `xml:"ДатаПредыдущегоСписка"`
	СписокАктуальныхРешений СписокАктуальныхРешений `xml:"СписокАктуальныхРешений"`
	СписокОтмененныхРешений СписокОтмененныхРешений `xml:"СписокОтмененныхРешений"`
}

// СписокАктуальныхРешений is generated from an XSD element.
type СписокАктуальныхРешений struct {
	Решение []Решение `xml:"Решение"`
}

// Решение is generated from an XSD element.
type Решение struct {
	ТипРешения      ТипРешения      `xml:"ТипРешения"`
	НомерРешения    string          `xml:"НомерРешения"`
	ДатаРешения     time.Time       `xml:"ДатаРешения"`
	Орган           string          `xml:"Орган"`
	ВидРешения      ВидРешения      `xml:"ВидРешения"`
	СписокСубъектов СписокСубъектов `xml:"СписокСубъектов"`
}

// ТипРешения is generated from an XSD element.
type ТипРешения struct {
	Идентификатор int    `xml:"Идентификатор"`
	Наименование  string `xml:"Наименование"`
}

// ВидРешения is generated from an XSD element.
type ВидРешения struct {
	Идентификатор int    `xml:"Идентификатор"`
	Наименование  string `xml:"Наименование"`
}

// СписокСубъектов is generated from an XSD element.
type СписокСубъектов struct {
	Субъект []Субъект `xml:"Субъект"`
}

// Субъект is generated from an XSD element.
type Субъект struct {
	ИдСубъекта        string        `xml:"ИдСубъекта"`
	ТипСубъекта       ТипСубъекта   `xml:"ТипСубъекта"`
	СписокАдресов     СписокАдресов `xml:"СписокАдресов"`
	РешениеПоСубъекту string        `xml:"РешениеПоСубъекту"`
	Примечание        string        `xml:"Примечание"`
}

// ТипСубъекта is generated from an XSD element.
type ТипСубъекта struct {
	Идентификатор int    `xml:"Идентификатор"`
	Наименование  string `xml:"Наименование"`
}

// СписокАдресов is generated from an XSD element.
type СписокАдресов struct {
	Адрес []Адрес `xml:"Адрес"`
}

// Адрес is generated from an XSD element.
type Адрес struct {
	НормАдрес            НормАдрес `xml:"НормАдрес"`
	ТипАдреса            ТипАдреса `xml:"ТипАдреса"`
	ТекстАдреса          string    `xml:"ТекстАдреса"`
	ИдентификаторФИАС    string    `xml:"ИдентификаторФИАС"`
	УровеньАдреснОбъекта int       `xml:"УровеньАдреснОбъекта"`
	Страна               Страна    `xml:"Страна"`
	Индекс               string    `xml:"Индекс"`
	ОКАТО                string    `xml:"ОКАТО"`
	Регион               string    `xml:"Регион"`
	АвтономныйОкруг      string    `xml:"АвтономныйОкруг"`
	Район                string    `xml:"Район"`
	Город                string    `xml:"Город"`
	ВнутригородскРайон   string    `xml:"ВнутригородскРайон"`
	НаселПункт           string    `xml:"НаселПункт"`
	Улица                string    `xml:"Улица"`
	Дом                  string    `xml:"Дом"`
	Корпус               string    `xml:"Корпус"`
	Строение             string    `xml:"Строение"`
	Помещение            string    `xml:"Помещение"`
}

// НормАдрес is generated from an XSD element.
type НормАдрес struct {
	Идентификатор int    `xml:"Идентификатор"`
	Наименование  string `xml:"Наименование"`
}

// ТипАдреса is generated from an XSD element.
type ТипАдреса struct {
	Идентификатор int    `xml:"Идентификатор"`
	Наименование  string `xml:"Наименование"`
}

// Страна is generated from an XSD element.
type Страна struct {
	Наименование string `xml:"Наименование"`
	Код          string `xml:"Код"`
}

// СписокОтмененныхРешений is generated from an XSD element.
type СписокОтмененныхРешений struct {
	ОтмененноеРешение []ОтмененноеРешение `xml:"ОтмененноеРешение"`
}

// ОтмененноеРешение is generated from an XSD element.
type ОтмененноеРешение struct {
	ТипРешения      ТипРешения      `xml:"ТипРешения"`
	НомерРешения    string          `xml:"НомерРешения"`
	ДатаРешения     time.Time       `xml:"ДатаРешения"`
	Орган           string          `xml:"Орган"`
	ВидРешения      ВидРешения      `xml:"ВидРешения"`
	СписокСубъектов СписокСубъектов `xml:"СписокСубъектов"`
}