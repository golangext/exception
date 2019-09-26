package exception

const httpStatusInternalServerError = 500

var uncatchable = newType("uncatchable", nil, Traits{
	HttpStatusCode: httpStatusInternalServerError,
	Description:    "An uncatchable internal only exception",
	DisplayToUser:  False,
	AllowThrow:     True,
	AllowSubclass:  False,
})

var All = newType("All", nil, Traits{
	HttpStatusCode: httpStatusInternalServerError,
	Description:    "An exception",
	DisplayToUser:  False,
	AllowThrow:     False,
	AllowSubclass:  False,
})

var Runtime = All.extend("Runtime", Traits{})

var Panic = Runtime.extend("Panic", Traits{
	Description: "An unwrapped or caught generic application panic",
})

var NotImplemented = Runtime.extend("NotImplemented", Traits{
	Description:   "The method is not implemented",
	AllowThrow:    True,
	AllowSubclass: True,
})

var Bounds = Runtime.extend("Bounds", Traits{
	Description: "An invalid data access",
})

var NullPointer = Bounds.extend("NullPointer", Traits{
	Description: "An null pointer exception from the runtime",
})

var IO = All.extend("IO", Traits{
	Description:   "Input/output error",
	AllowThrow:    True,
	AllowSubclass: True,
})
