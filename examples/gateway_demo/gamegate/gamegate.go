package gamegate

type gameGate struct {
}

func (g *gameGate) handleConnection() error {
	return nil
}

func (g *gameGate) handleDisconnect() {

}

func (g *gameGate) checkAuth() *AuthData {
	return nil
}

func (g *gameGate) kick(userId string) {

}

func (g *gameGate) clientCount() {
}

func (g *gameGate) startHeartbeat() {
}

func (g *gameGate) stopHeartbeat() {
}

func (g *gameGate) checkAlive() {

}

func (g *gameGate) broadcast() {

}

func (g *gameGate) notify(userId string, data interface{}) {

}

func (g *gameGate) streamHandler() {

}
