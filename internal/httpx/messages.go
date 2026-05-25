package httpx

// Common response messages shared across handlers. These are generic
// HTTP-boundary strings (bind failures, auth, fallthrough errors); they are
// produced where a response is written and never matched with errors.Is — so
// they are plain constants, not error sentinels. Domain-specific messages
// belong in their own package.
const (
	MsgInvalidBody   = "invalid request body"
	MsgUnauthorized  = "unauthorized"
	MsgInternalError = "internal error"
)
