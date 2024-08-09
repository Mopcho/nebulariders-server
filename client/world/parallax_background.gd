@tool
extends ParallaxBackground

@export var boundary: int:
	set(new_boundary):
		boundary = new_boundary
		print('gay')

func _process(delta):
	$ParallaxLayer/Sprite1.region_rect.size = Vector2(boundary,boundary)
	$ParallaxLayer2/Sprite2.region_rect.size = Vector2(boundary,boundary)
