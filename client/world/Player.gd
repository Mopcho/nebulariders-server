extends CharacterBody2D

@export var speed = 200
var target_position = self.position

func _physics_process(delta):
	$DebugLabel.text = "X: {x}, Y: {y}\nTarget X: {tx}, Y: {ty}".format(
	{
		"x": snapped(self.position.x, 0.01),
		"y": snapped(self.position.y, 0.01),
		"tx": snapped(target_position.x, 0.01),
		"ty": snapped(target_position.y, 0.01)
	})
	
	if (Input.is_action_pressed("left_mouse")):
		target_position = get_global_mouse_position()
		$SpriteWrapper.look_at(target_position)
	
	var direction = (target_position - global_position).normalized()
	var distance = global_position.distance_to(target_position)

	if distance > 2:
		var velocity = direction * speed * delta
		move_and_collide(velocity)
