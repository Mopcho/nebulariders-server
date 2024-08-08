extends Node2D

const PLAYER_ID = "b6023e05-5aec-40b8-9500-9d6f3b949d7c";

var OtherPlayer = preload("res://world/OtherPlayer.tscn")

var ENTITIES = {}

func _on_world_websocket_packet(packed: PackedByteArray):
	var message = JSON.parse_string(packed.get_string_from_utf8())
	var players = message["players"]
	
	for playerId in players:
		if (playerId == PLAYER_ID):
			continue
		if !ENTITIES.has(playerId):
			var otherPlayerInstance = OtherPlayer.instantiate()
			ENTITIES[playerId] = otherPlayerInstance
			add_child(otherPlayerInstance)
		var server_player = players[playerId]
		var local_player = ENTITIES[playerId]
		local_player.position = Vector2(server_player["x"], server_player["y"])
