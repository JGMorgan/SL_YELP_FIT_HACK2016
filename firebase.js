var Firebase = require("firebase");
var myFirebaseRef = new Firebase("https://decisionmadeeasy.firebaseio.com/");

var user_name = "Leon";
var user_answer = "1";
var newUser = myFirebaseRef.child("Users");
function NewUser(Firebase ref){
	
	ref.set({
		name: user_name,
		answer1: user_answer
	});
	
	
}