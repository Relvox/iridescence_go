package servers

type BadRequestError string

func ToBadRequestError(err error) BadRequestError { return BadRequestError(err.Error()) }

func (e BadRequestError) Error() string { return string(e) }

type InternalError string

func ToInternalError(err error) InternalError { return InternalError(err.Error()) }

func (e InternalError) Error() string { return string(e) }

type PanicError string

func ToPanicError(err error) PanicError { return PanicError(err.Error()) }

func (e PanicError) Error() string { return string(e) }
