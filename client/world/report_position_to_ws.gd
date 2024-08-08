extends Node2D

var network_tick_timer = Timer.new();
var player: CharacterBody2D;
var last_position: Vector2;
var socket: WebSocketPeer;

func _ready():
	player = get_node(get_meta("PlayerNodePath"))
	
	add_child(network_tick_timer);
	network_tick_timer.connect("timeout", self._on_network_tick)
	network_tick_timer.wait_time = 0.1;
	network_tick_timer.autostart = false
	network_tick_timer.start()

func _on_network_tick():
	var new_position = player.position;
	if (last_position.x != new_position.x || last_position.y != new_position.y):
		send_position_to_server(new_position)
	
	last_position = new_position;

func send_position_to_server(position):
	var data = { "type": "position", "x": snapped(position.x, 0.1), "y": snapped(position.y, 0.1) }
	var packet = JSON.stringify(data)
	if (socket):
		socket.put_packet(packet.to_utf8_buffer())

func _on_world_websocket_open(socket):
	self.socket = socket
