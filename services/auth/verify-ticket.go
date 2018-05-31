package auth

const verifyTicketKey = keyPrefix + "verify_ticket"

func SetVerifyTicket(ticket string) error {
	return redisClient.Set(verifyTicketKey, ticket, 0).Err()
}

func GetVerifyTicket() (string, error) {
	return redisClient.Get(verifyTicketKey).Result()
}