package model



type Role string

const (
	RoleCustomer Role = "CUSTOMER"
	RoleStaff Role = "STAFF"
	RoleTicketManager Role = "TICKET_MANAGER"
	RoleAdmin Role = "ADMIN"
)



type CreateUserInput struct {
	DisplayName string `json:"displayName"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role *Role `json:"role"`
}
