extends Control

signal submit(email: String, password: String)
signal register_here_pressed

func _ready():
	$EmailInput.grab_focus()

func _on_submit():
	submit.emit($EmailInput.text, $PasswordInput.text)

func _on_register_here():
	register_here_pressed.emit()

func _on_gui_input_forbid_tab(event, next_node, prev_node):
	if (event.is_action_pressed("ui_focus_prev")):
		get_node(prev_node).grab_focus()
		self.accept_event()
	elif (event.is_action_pressed("ui_focus_next")):
		get_node(next_node).grab_focus()
		self.accept_event()
