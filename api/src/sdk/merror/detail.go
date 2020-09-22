package merror

const (
	DVConflict     string = "conflict"      // unique...
	DVMalformed    string = "malformed"     // email format,  ip address format...
	DVInvalid      string = "invalid"       // minumum/maximum value/lenght...
	DVRequired     string = "required"      // missing in request...
	DVExpired      string = "expired"       // expired duration...
	DVForbidden    string = "forbidden"     // forbidden to update...
	DVInternal     string = "internal"      // internal error occured
	DVLocked       string = "locked"        // cannot be updated
	DVNotFound     string = "not_found"     // correspondance has not been found
	DVNotSupported string = "not_supported" // not handled by the running implementation
	DVTimedOut     string = "timed_out"     // something... timed out
	DVUnauthorized string = "unauthorized"  // authorization is missing
	DVUnknown      string = "unknown"       // unknown detail code
	DVNoCode       string = "no_code"       // no specific code
)
