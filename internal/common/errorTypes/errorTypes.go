package errorTypes

import "errors"

var (
	ErrJSONUnmarshalling  = errors.New("error while trying to unmarshal json")
	ErrDBDataInsertion    = errors.New("error while inserting new data to db table")
	ErrDBDataDeletion     = errors.New("error while deleting db table data")
	ErrDBDataReception    = errors.New("error while receipting db data")
	ErrDBDataModification = errors.New("error while changing db data")
	ErrBadRequestData     = errors.New("error bad request data")
)
