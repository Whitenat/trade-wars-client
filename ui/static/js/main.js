$(document).ready(function() {

});

function getName() {
	var username = $("#username").val()
	window.location.href='/navigation'
	document.getElementById("showUser").innerHTML = username;
}
