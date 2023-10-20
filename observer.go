package main

import "fmt"

//visa application opening

//subject - application
//concrete subject - visa
//observer - applicant
//concrete observer - human


type Visa struct {
	Available bool
	Country string
	ApplicantList []Applicant
}

type Application interface {
	Apply(applicant Applicant)
	Withdraw(applicant Applicant)
	NotifyAll()
}

func NewVisa(country string) *Visa {
	return &Visa{
		Country: country,
	}
}

func (v *Visa) OpenApplications() {
	fmt.Printf("%s has opened visa applications! \n", v.Country)
	v.Available=true
	v.NotifyAll()
}

func (v *Visa) NotifyAll() {
	for _, applicant := range v.ApplicantList {
		applicant.Update(v.Country)
	}
}

func (v *Visa) Apply(applicant Applicant) {
	v.ApplicantList=append(v.ApplicantList, applicant)
}

func (v *Visa) Withdraw(applicant Applicant) {
	v.ApplicantList= removeFromList(v.ApplicantList, applicant)
}

func removeFromList(applicantList []Applicant, applicant Applicant) []Applicant {
	indexToRemove := -1
	for i, a := range applicantList {
		if a.GetID() == applicant.GetID() {
			indexToRemove = i
			break
		}
	}
	if indexToRemove != -1 {
		applicantList = append(applicantList[:indexToRemove], applicantList[indexToRemove+1:]...)
	}
	return applicantList
}

type Human struct {
	Name string
	PassportId string
}

func (h *Human) Update(country string) {
    fmt.Printf("%s with PassportId: %s can now apply for Visa in %s\n", h.Name, h.PassportId , country)
}

func (h *Human) GetID() string {
    return h.PassportId
}

type Applicant interface {
	Update(string)
	GetID() string
}

func main(){
	visa1 := NewVisa("Italy")
	visa2 := NewVisa("Spain")

	human1 := &Human{Name:"Elbasy", PassportId: "777"}
	human2 := &Human{Name:"Rabotyaga", PassportId: "123"}

	visa1.Apply(human1)
	visa1.Apply(human2)
	visa1.Withdraw(human2)
	visa2.Apply(human2)

	visa1.OpenApplications()
	visa2.OpenApplications()
}