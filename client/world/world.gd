extends Node2D

const WEBSOCKET_BASE_URL = "ws://localhost:8080";
const TOKEN = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMyMzQ4NjQsInN1YiI6ImI2MDIzZTA1LTVhZWMtNDBiOC05NTAwLTlkNmYzYjk0OWQ3YyIsInVzZXJuYW1lIjoic3RvamFuIn0.iEWwITC66OYJC_ONaQZg576bFhzXziiWX3wAKTsJJQA'

var socket = WebSocketPeer.new()

var socket_is_open = false

signal websocket_open(socket: WebSocketPeer)
signal websocket_packet(packed: PackedByteArray)
signal websocket_closed(code: int, reason: String)

# Called when the node enters the scene tree for the first time.
func _ready():
	var url = "%s/ws?token=%s" % [WEBSOCKET_BASE_URL, TOKEN]
	socket.connect_to_url(url)

# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	socket.poll()
	var state = socket.get_ready_state()
	if state == WebSocketPeer.STATE_OPEN:
		if (!socket_is_open):
			socket_is_open = true
			self.websocket_open.emit(socket)
		while socket.get_available_packet_count():
			self.websocket_packet.emit(socket.get_packet())
	elif state == WebSocketPeer.STATE_CLOSING:
		# Keep polling to achieve proper close.
		pass
	elif state == WebSocketPeer.STATE_CLOSED:
		var code = socket.get_close_code()
		var reason = socket.get_close_reason()
		self.websocket_closed.emit(code, reason)
		print("WebSocket closed with code: %d, reason %s. Clean: %s" % [code, reason, code != -1])
		set_process(false) # Stop processing.
