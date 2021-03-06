package entities

import "time"

type AgendaElement struct {
	ElementType   string    `json:"elementType"`
	ID            uint      `json:"id"`
	UserID        uint      `json:"userId"`
	Title         string    `json:"title"`
	PaymentAmount float64   `json:"paymentAmount"`
	PaymentDate   time.Time `json:"paymentDate"`
}

type AgendaElementTransformable interface {
	Transform() AgendaElement
}

type PeriodicAgendaElementsTransformable interface {
	Transform() AgendaElement
	TransformWithPeriod(from time.Time, to time.Time) []AgendaElement
}
