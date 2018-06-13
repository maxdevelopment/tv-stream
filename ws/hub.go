package ws

type hub struct {
	Clients      map[string]*Client
	register     chan *Client
	unregister   chan *Client
}

var H = hub{
	Clients:      make(map[string]*Client),
	register:     make(chan *Client),
	unregister:   make(chan *Client),
}

func (h *hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Clients[client.Id] = client

		case client := <-h.unregister:
			if _, ok := h.Clients[client.Id]; ok {
				delete(h.Clients, client.Id)
			}
		}
	}
}
