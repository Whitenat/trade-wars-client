//init object globally
var ship= null;
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
	if(ship.style.left > '0 em') {
		ship.style.left=parseInt(ship.style.left)-4 +'em';
	}	
}
// function moveLeft(){
// 	ship.style.left=parseInt(ship.style.left)-7 +'em';	
// }
function moveUp(){
	ship.style.top=parseInt(ship.style.top)-4 +'em';
}
function moveRight(){
	ship.style.left=parseInt(ship.style.left)+4 +'em';
}
function moveDown(){
	ship.style.top=parseInt(ship.style.top)+4 +'em';
}
window.onload=init;
