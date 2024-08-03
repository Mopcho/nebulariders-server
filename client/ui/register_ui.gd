extends Control

var http_request = HTTPRequest.new()

const JSON_HEADERS = ["Content-Type: application/json"];
const BASE_URL = "http://localhost:8081";

func _ready():
	add_child(http_request)
	http_request.request_completed.connect(self._on_request_completed)

func _on_request_completed(result, response_code, headers, body):
	var json = JSON.parse_string(body.get_string_from_utf8())
	
	if response_code == 200:
		get_tree().change_scene_to_file("res://ui/login_ui.tscn")

func _on_submit():
	print("POST /api/auth/register")
	var url = "{BASE_URL}/api/auth/register".format({"BASE_URL": BASE_URL})
	
	var register_data = JSON.new().stringify(
		{ 
			"username": $UsernameInput.text,
			"email": $EmailInput.text,
			"password": $PasswordInput.text 
		}
	)
	
	print(register_data)
	
	var result = http_request.request(url, JSON_HEADERS, HTTPClient.METHOD_POST, register_data);
	
	print(result);
