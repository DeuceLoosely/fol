

list = document.getElementById("memes")

for (i = 0; i < 6; i++) {

var entry = document.createElement('li');

var elem = document.createElement("img");
elem.src = 'images/'+ i +'.png';

entry.appendChild(elem);


list.appendChild(entry);

} 


