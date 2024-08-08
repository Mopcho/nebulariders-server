extends Control

signal submit(username: String, email: String, password: String)
signal sign_here_pressed

func _ready():
	$UsernameInput.grab_focus()

func _on_submit():
	submit.emit($UsernameInput.text, $EmailInput.text, $PasswordInput.text)

func _on_sign_here():
	sign_here_pressed.emit()

func _on_gui_input_forbid_tab(event, next_node, prev_node):
	if (event.is_action_pressed("ui_focus_prev")):
		get_node(prev_node).grab_focus()
		self.accept_event()
	elif (event.is_action_pressed("ui_focus_next")):
		get_node(next_node).grab_focus()
		self.accept_event()
