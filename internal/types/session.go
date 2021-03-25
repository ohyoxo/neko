package types

import "net/http"

type SessionState struct {
	IsConnected bool `json:"is_connected"`
	IsWatching  bool `json:"is_watching"`
}

type Session interface {
	ID() string
	Profile() MemberProfile
	State() SessionState
	IsHost() bool

	// websocket
	SetWebSocketPeer(websocketPeer WebSocketPeer)
	SetWebSocketConnected(websocketPeer WebSocketPeer, connected bool)
	Send(v interface{}) error
	Disconnect(reason string) error

	// webrtc
	SetWebRTCPeer(webrtcPeer WebRTCPeer)
	SetWebRTCConnected(webrtcPeer WebRTCPeer, connected bool)
	GetWebRTCPeer() WebRTCPeer
}

type SessionManager interface {
	Create(id string, profile MemberProfile) (Session, string, error)
	Update(id string, profile MemberProfile) error
	Delete(id string) error
	Get(id string) (Session, bool)
	GetByToken(token string) (Session, bool)
	List() []Session

	SetHost(host Session)
	GetHost() Session
	ClearHost()

	Broadcast(v interface{}, exclude interface{})
	AdminBroadcast(v interface{}, exclude interface{})

	OnCreated(listener func(session Session))
	OnDeleted(listener func(session Session))
	OnConnected(listener func(session Session))
	OnDisconnected(listener func(session Session))
	OnProfileChanged(listener func(session Session))
	OnStateChanged(listener func(session Session))
	OnHostChanged(listener func(session Session))

	ImplicitHosting() bool

	CookieSetToken(w http.ResponseWriter, token string)
	CookieClearToken(w http.ResponseWriter, r *http.Request)
	Authenticate(r *http.Request) (Session, error)
}
