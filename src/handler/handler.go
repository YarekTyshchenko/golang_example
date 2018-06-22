package handler

type Request struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Picture []byte `json:"picture"`
}

type Response struct {
	Id     string `json:"id"`
}

type Storage interface {
	Store(request Request) (string, error)
}

type Handler struct {
	storage Storage
}

func New(s Storage) Handler {
	return Handler{
		storage: s,
	}
}

func (h *Handler) Handle(r Request) (Response, error) {
	id, err := h.storage.Store(r)
	if err != nil {
		return Response{}, err
	}
	return Response{Id:id}, nil
}