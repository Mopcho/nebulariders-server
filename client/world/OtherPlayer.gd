extends CharacterBody2D

var server_position: Vector2

func setup(username: String):
	$Label.text = username
	
func _physics_process(delta):
	self.position = self.position.lerp(server_position, delta * 4.0)
	
func move_to_position(new_position: Vector2):
	server_position = new_position
	$SpriteWrapper.look_at(new_position)
