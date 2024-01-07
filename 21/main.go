// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// реализация паттерна представлена в виде адаптера для
// структуры, работающей с json файлами, который реализует
// интерфейс для работы с xml файлами.
// паттерн позволяет легко конвертировать интерфейс структуры,
// в ожидаемый клиентом

type XMLSender interface {
	SendXMLData()
}

type XMLDocument struct {

}

func (x *XMLDocument) SendXMLData() {
	fmt.Println("xml sended from xmldocument")
}

type JSONDocument struct {

}

func (j *JSONDocument) ConvertToXML() {
	fmt.Println("json converted to xml")
}

type JSONDocumentAdapter struct {
	adapter *JSONDocument
}

func (j *JSONDocumentAdapter) SendXMLData() {
	j.adapter.ConvertToXML()
	fmt.Println("xml sended from json adapter")
}

func main() {
	var service XMLSender = &JSONDocumentAdapter{
		adapter: &JSONDocument{},
	}

	service.SendXMLData()
}