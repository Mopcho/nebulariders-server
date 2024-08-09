extends Node2D

var OtherPlayer = preload("res://world/OtherPlayer.tscn")

var ENTITIES = {}

func _on_world_websocket_packet(packed: PackedByteArray):
	var message = JSON.parse_string(packed.get_string_from_utf8())
	var players = message["data"]["players"]
	
	for playerId in players:
		if !ENTITIES.has(playerId):
			print(message)
			var playerData = players[playerId];
			var otherPlayerInstance = OtherPlayer.instantiate()
			otherPlayerInstance.setup(playerData["username"])
			ENTITIES[playerId] = { "instance": otherPlayerInstance, "data": playerData }
			add_child(otherPlayerInstance)
		var server_player = players[playerId]
		var local_player = ENTITIES[playerId]
		var next_position = Vector2(server_player["x"], server_player["y"])
		
		local_player.instance.move_to_position(next_position)
