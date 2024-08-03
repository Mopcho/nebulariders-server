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
		print("YOU ARE NOW LOGGED!")

func _on_submit():
	print("POST /api/auth/login")
	var url = "{BASE_URL}/api/auth/login".format({"BASE_URL": BASE_URL})
	
	var login_data = JSON.new().stringify(
		{ 
			"email": $EmailInput.text,
			"password": $PasswordInput.text 
		}
	)
	
	print(login_data)
	
	var result = http_request.request(url, JSON_HEADERS, HTTPClient.METHOD_POST, login_data);
	
	print(result);
