//init object globally
var ship= null;
var vertCount = 0;
var horizCount = 0;
function init(){
	ship=document.getElementById("ship");				
	ship.style.position='relative';
	ship.style.left='0px';
	ship.style.top='0px';
}
function getKeyAndMove(e){				
	var key_code=e.which||e.keyCode;
	switch(key_code){
		case 37: //left arrow key
			moveLeft();
			break;
		case 38: //Up arrow key
			moveUp();
			break;
		case 39: //right arrow key
			moveRight();
			break;
		case 40: //down arrow key
			moveDown();
			break;						
	}
}
function moveLeft(){

	if(horizCount > 0) {
		ship.style.left=parseInt(ship.style.left)-4 +'vw';
		horizCount -= 1;
	}	
}
function moveUp(){

	if(vertCount > 0) {
		ship.style.top=parseInt(ship.style.top)-4 +'vw';
		vertCount -= 1;
	}	
}
	
function moveRight(){

	if(horizCount < 9) {
		ship.style.left=parseInt(ship.style.left)+4 +'vw';
		horizCount += 1;
	}
}
function moveDown(){
	if(vertCount < 9) {
		ship.style.top=parseInt(ship.style.top)+4 +'vw';
		vertCount += 1;
	}
}
window.onload=init;
