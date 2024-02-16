package request

type EquipmentReservingRequest struct {
	UserId               int    `json:"userId" validate:"required"`
	ReservationStartTime string `json:"reservationStartTime" validate:"required"`
	ReservationEndTime   string `json:"reservationEndTime" validate:"required"`
	ActivityStartTime    string `json:"activityStartTime" validate:"required"`
	ActivityEndTime      string `json:"activityEndTime" validate:"required"`
}
