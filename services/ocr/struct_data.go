package ocr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Data is a nested struct in ocr response
type Data struct {
	Name                string               `json:"Name" xml:"Name"`
	BankName            string               `json:"BankName" xml:"BankName"`
	IDNumber            string               `json:"IDNumber" xml:"IDNumber"`
	Date                string               `json:"Date" xml:"Date"`
	Seat                string               `json:"Seat" xml:"Seat"`
	FileContent         string               `json:"FileContent" xml:"FileContent"`
	CardNumber          string               `json:"CardNumber" xml:"CardNumber"`
	VinCode             string               `json:"VinCode" xml:"VinCode"`
	Angle               float64              `json:"Angle" xml:"Angle"`
	BirthPlace          string               `json:"BirthPlace" xml:"BirthPlace"`
	Nationality         string               `json:"Nationality" xml:"Nationality"`
	NativePlace         string               `json:"NativePlace" xml:"NativePlace"`
	ValidDate           string               `json:"ValidDate" xml:"ValidDate"`
	Destination         string               `json:"Destination" xml:"Destination"`
	Level               string               `json:"Level" xml:"Level"`
	Price               float64              `json:"Price" xml:"Price"`
	Relation            string               `json:"Relation" xml:"Relation"`
	Number              string               `json:"Number" xml:"Number"`
	DepartureStation    string               `json:"DepartureStation" xml:"DepartureStation"`
	Gender              string               `json:"Gender" xml:"Gender"`
	BirthDate           string               `json:"BirthDate" xml:"BirthDate"`
	OfficePhoneNumbers  []string             `json:"OfficePhoneNumbers" xml:"OfficePhoneNumbers"`
	Departments         []string             `json:"Departments" xml:"Departments"`
	Companies           []string             `json:"Companies" xml:"Companies"`
	Titles              []string             `json:"Titles" xml:"Titles"`
	Emails              []string             `json:"Emails" xml:"Emails"`
	CellPhoneNumbers    []string             `json:"CellPhoneNumbers" xml:"CellPhoneNumbers"`
	Addresses           []string             `json:"Addresses" xml:"Addresses"`
	BackResult          BackResult           `json:"BackResult" xml:"BackResult"`
	TitleArea           TitleArea            `json:"TitleArea" xml:"TitleArea"`
	FaceResult          FaceResult           `json:"FaceResult" xml:"FaceResult"`
	FrontResult         FrontResult          `json:"FrontResult" xml:"FrontResult"`
	UndertakeStampAreas []UndertakeStampArea `json:"UndertakeStampAreas" xml:"UndertakeStampAreas"`
	Tables              []Table              `json:"Tables" xml:"Tables"`
	InvalidStampAreas   []InvalidStampArea   `json:"InvalidStampAreas" xml:"InvalidStampAreas"`
	Invoices            []Invoice            `json:"Invoices" xml:"Invoices"`
	RegisterStampAreas  []RegisterStampArea  `json:"RegisterStampAreas" xml:"RegisterStampAreas"`
	Results             []Result             `json:"Results" xml:"Results"`
	OtherStampAreas     []OtherStampArea     `json:"OtherStampAreas" xml:"OtherStampAreas"`
	Plates              []Plate              `json:"Plates" xml:"Plates"`
}
