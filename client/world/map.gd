@tool
extends ParallaxBackground

@export var map_boundary:= 2000:
	set(boundary):
		map_boundary = boundary
		if Engine.is_editor_hint():
			change_regions(map_boundary, map_border_width)

@export var map_border_width:= 20:
	set(border_width):
		map_border_width = border_width
		if Engine.is_editor_hint():
			change_regions(map_boundary, map_border_width)

func _ready():
	if not Engine.is_editor_hint():
		change_regions(map_boundary, map_border_width)

func change_regions(boundary: int, border_width: int) -> void:
	var size = Vector2(boundary, boundary)
	
	$ParallaxLayer/Sprite.region_rect.size = size
	$ParallaxLayer2/Sprite.region_rect.size = size
	
	var vertical_shape = RectangleShape2D.new()
	vertical_shape.size = Vector2(border_width, boundary)
	
	var horizontal_shape = RectangleShape2D.new()
	horizontal_shape.size = Vector2(boundary + (border_width * 2), border_width)
	
	$BorderTop.position = Vector2(0, ((boundary / 2) + border_width / 2) * -1)
	$BorderTop/CollisionShape2D.set_shape(horizontal_shape)
	
	$BorderRight.position = Vector2(((boundary / 2) + border_width / 2), 0)
	$BorderRight/CollisionShape2D.set_shape(vertical_shape)

	$BorderBottom.position = Vector2(0, (boundary / 2) + (border_width / 2))
	$BorderBottom/CollisionShape2D.set_shape(horizontal_shape)
	
	$BorderLeft.position = Vector2(((boundary / 2) + border_width / 2) * -1, 0)
	$BorderLeft/CollisionShape2D.set_shape(vertical_shape)
