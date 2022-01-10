package requests

type IDUriReq struct {
	ID int `uri:"id" binding:"required"`
}
