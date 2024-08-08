extends Node2D

const JSON_HEADERS = ["Content-Type: application/json"];
const BASE_URL = "http://localhost:8081";

var register_request = HTTPRequest.new()
var login_request = HTTPRequest.new();

func _ready():
	add_child(register_request);
	add_child(login_request);
	
	register_request.connect("request_completed", self._on_register_request_completed)
	login_request.connect("request_completed", self._on_login_request_completed)
	$LoginUI.show();

func _on_login_ui_register_here_pressed():
	$RegisterUI.show()
	$LoginUI.hide()

func _on_register_ui_sign_here_pressed():
	$LoginUI.show()
	$RegisterUI.hide()

func _on_register_ui_submit(username: String, email: String, password: String):
	print("POST /api/auth/register")
	var url = "{BASE_URL}/api/auth/register".format({"BASE_URL": BASE_URL})
	
	var register_data = JSON.new().stringify(
		{ 
			"username": username,
			"email": email,
			"password": password
		}
	)
	
	print(register_data)
	
	register_request.request(url, JSON_HEADERS, HTTPClient.METHOD_POST, register_data);

func _on_login_ui_submit(email: String, password: String):
	print("POST /api/auth/login")
	var url = "{BASE_URL}/api/auth/login".format({"BASE_URL": BASE_URL})
	
	var login_data = JSON.new().stringify(
		{ 
			"email": email,
			"password": password
		}
	)
	
	print(login_data)
	
	login_request.request(url, JSON_HEADERS, HTTPClient.METHOD_POST, login_data);

func _on_register_request_completed(result, response_code, headers, body):
	print(JSON.parse_string(body.get_string_from_utf8()))
	if (response_code == 200):
		$LoginUI.show()
		$RegisterUI.hide();
		print("REGISTRETED")

func _on_login_request_completed(result, response_code, headers, body):
	print(JSON.parse_string(body.get_string_from_utf8()))
	if (response_code == 200):
		get_tree().change_scene_to_file("res://world/world.tscn");
