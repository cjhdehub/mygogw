package schema

type MsgPack struct {
	MsgType string
	MsgContent string
	Msg interface{}
}

type RegisterRequest struct {
	VP string
}

type RegisterResponse struct {
	ClientId string
	Status string
}

type OpenConnRequest struct {
	ConnId string

	//ROLE_READER/ROLE_WRITER/ROLE_QUERY_CONNID
	Role string

	//OPERATOR_CLOSE/OPERATOR_DATA_TRANSFER
	Operator string
}

type OpenConnResponse struct {
	ConnId string

	//STATUS_SUCCESS/STATUS_FAILED
	Status string

	//OPERATOR_CLOSE/OPERATOR_DATA_TRANSFER
	Operator string
}