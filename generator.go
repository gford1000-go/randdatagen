package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/Pallinder/go-randomdata"
)

func NewGenerator(recordCount int, country2Letter string) *generator {
	return &generator{
		recordCount: recordCount,
		country:     country2Letter,
	}
}

type generator struct {
	recordCount int
	country     string
}

func (g *generator) createIndividual() []string {
	gender := randomdata.Number(0, 2)

	return []string{
		randomdata.Title(gender),
		randomdata.FirstName(gender),
		randomdata.LastName(),
	}
}

func (g *generator) createAddress() []string {
	return []string{
		fmt.Sprintf("%v", randomdata.Number(1, 250)),
		randomdata.StreetForCountry(g.country),
		randomdata.ProvinceForCountry(g.country),
		randomdata.PostalCode(g.country),
	}
}

func (g *generator) createDOB(minAge int) string {
	currentYear := time.Now().Year()

	return fmt.Sprintf("%02v-%v-%v",
		randomdata.Number(1, 32),
		randomdata.Month()[0:3],
		randomdata.Number(currentYear-120, currentYear-minAge))
}

func (g *generator) createMobile() string {
	operator := []string{"07107", "07400", "07414", "07493", "07431", "07510", "07726"}[randomdata.Number(0, 7)]

	return fmt.Sprintf("+44 %v %v %v",
		operator,
		randomdata.Number(0, 1000),
		randomdata.Number(0, 1000))
}

func (g *generator) createEmail(individual []string) string {

	adddot := []string{".", ""}[randomdata.Number(0, 2)]
	addnum := []string{"", fmt.Sprintf("%v", randomdata.Number(10, 1000))}[randomdata.Number(0, 2)]
	provider := []string{"outlook.com", "yahoo.com", "bt.com", "example.com", "test.com"}[randomdata.Number(0, 5)]

	return fmt.Sprintf("%v%v%v%v@%v",
		strings.ToLower(individual[1][:1]),
		adddot,
		strings.ToLower(individual[2]),
		addnum,
		provider)
}

func (g *generator) Create(w io.Writer) error {

	for i := 0; i < g.recordCount; i++ {
		individual := g.createIndividual()

		record := []string{}
		record = append(record, individual...)
		record = append(record, g.createEmail(individual))
		record = append(record, g.createDOB(18))
		record = append(record, g.createAddress()...)
		record = append(record, g.createMobile())
		record = append(record, randomdata.PhoneNumber())

		_, err := w.Write([]byte(strings.Join(record, ",") + "\n"))
		if err != nil {
			return fmt.Errorf("record %v: %v", i, err)
		}
	}

	return nil
}