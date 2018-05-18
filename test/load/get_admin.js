import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
	vus: 300,
	duration: "5s",
};

export default function() {
	// Login as user
	var sessionUrl = "http://127.0.0.1:3000/api/v1/admin/auth";
	var sessionPayload = JSON.stringify({ email: "Robin", password: "secret" });
	var sessionParams =  { headers: { "Content-Type": "application/json" } }
	var sessionRes = http.post(sessionUrl, sessionPayload, sessionParams);
	check(sessionRes, {
		"Authenticate": (r) => r.status === 200
	});
	
	// Show messages 
	var messageUrl = "http://127.0.0.1:3000/api/v1/admins";
	var messageParams =  { headers: { "Content-Type": "application/json", "hermods-cookie": String(sessionRes.headers["hermods-cookie"])} }
	var messageRes = http.get(messageUrl, messageParams);
	check(messageRes, {
		"Get Admins": (r) => r.status === 200
	});

	//var obj = JSON.parse(String(messageRes.body))
	//console.log(obj[0].body);
	sleep(1);
};

