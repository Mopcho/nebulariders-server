extends ParallaxBackground

@export var scroll_speed: Vector2 = Vector2(20,0);

# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	self.scroll_offset += scroll_speed * delta;
