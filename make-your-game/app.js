const grid = document.querySelector('.grid')
const timer = document.getElementById("Timer");
const scoreDisplay = document.querySelector('.score')
const statusDIsplay = document.querySelector(".status")
const Restart=document.getElementById('Restart')
const heart = document.getElementById("heart");
let currentShooterIndex = 202
let width = 15
let direction = 1
let invadersId
let isGameOver = false
let isWinner =false
let isPaused = false
let goingRight = true
let aliensRemoved = []
let score = 0
let time =0
let live =3
let Kill = false
let startCountingId;
Restart.onclick = () => {
  location.reload();
};
  
function AllGame(){
for (let i = 0; i < 225; i++) {
  const square = document.createElement('div')
  grid.appendChild(square)
}

const squares = Array.from(document.querySelectorAll('.grid div'))

let alienInvaders = [
  0,1,2,3,4,5,6,7,8,9,
  15,16,17,18,19,20,21,22,23,24,
  30,31,32,33,34,35,36,37,38,39
]
const alienRestart = alienInvaders
  document.addEventListener('keydown', function(e) {
  if (e.key === 'p' || e.key === 'P') {
    isPaused = !isPaused; // Изменяем состояние паузы при каждом нажатии клавиши "p"
    if (Kill || isWinner){
      return 
    }
    if (!isPaused) {

      moveInvaders(); // Если пауза снята, продолжаем движение пришельцев
    }
  }
});

function draw() {
  for (let i = 0; i < alienInvaders.length; i++) {
    if(!aliensRemoved.includes(i)) {
      squares[alienInvaders[i]].classList.add('invader');
    }
  }
}

draw()

function startCounting() {
  if (!isPaused || isWinner) {
    time = time + 1;
    timer.innerHTML = time;
  }
  startCountingId = setTimeout(startCounting, 1000);
}
startCountingId = setTimeout(startCounting, 1000);

function remove() {
  for (let i = 0; i < alienInvaders.length; i++) {
    squares[alienInvaders[i]].classList.remove('invader')
  }
}

squares[currentShooterIndex].classList.add('shooter')


function moveShooter(e) {
  if (isPaused || isWinner) {
    return;
  }
  squares[currentShooterIndex].classList.remove('shooter')
  switch(e.key) {
    case 'ArrowLeft':
      if (currentShooterIndex % width !== 0 && !isPaused) currentShooterIndex -=1
      break
      case 'ArrowRight' :
      if (currentShooterIndex % width < width -1 && !isPaused) currentShooterIndex +=1
      break
  }
  squares[currentShooterIndex].classList.add('shooter')
  
}
document.addEventListener('keydown', moveShooter)


function moveInvaders() {
  if (isPaused || isWinner) {
    if(isWinner){
      clearTimeout(startCountingId)
    }
    return;
  }
  const leftEdge = alienInvaders[0] % width === 0
  const rightEdge = alienInvaders[alienInvaders.length - 1] % width === width -1
  remove()

  if (rightEdge && goingRight) {
    for (let i = 0; i < alienInvaders.length; i++) {
      alienInvaders[i] += width +1
      direction = -1
      goingRight = false
    }
  }

  if(leftEdge && !goingRight) {
    for (let i = 0; i < alienInvaders.length; i++) {
      alienInvaders[i] += width -1
      direction = 1
      goingRight = true
    }
  }

  for (let i = 0; i < alienInvaders.length; i++) {
    alienInvaders[i] += direction
  }

  draw()

  if (squares[currentShooterIndex].classList.contains('invader', 'shooter')) {
    showGameOver();
    return;
  }
  const invaderRows = new Set(alienInvaders.map(index => Math.floor(index / width)));
  const shooterRow = Math.floor(currentShooterIndex / width);
  if (invaderRows.has(shooterRow)) {
    showGameOver();
    return; // Exit the function early to prevent further movement if game over.
  }
  for (let i = 0; i < alienInvaders.length; i++) {
    if(alienInvaders[i] >= (squares.length)) {
      showGameOver()
      return 
    }
  }
  if (aliensRemoved.length === alienInvaders.length) {
    statusDIsplay.innerHTML = '';
    const img =document.createElement('img');
    img.src='WIN.png'
    statusDIsplay.appendChild(img)
    isWinner = true
    clearTimeout(startCountingId)
    }
}
invadersId = setInterval(moveInvaders, 300)


function showGameOver() {
  if (live <= 0 || isWinner) {
    return;
  }
  
  Kill = true;
  live--; // Subtract 1 from live
  statusDIsplay.innerHTML = ''; // Очищаем содержимое, чтобы удалить текст
  const img = document.createElement('img');
  img.src = 'LOSE.png'; // Указываем путь к изображению
  clearTimeout(startCountingId);
  statusDIsplay.appendChild(img); // Добавляем изображение внутрь элемента с классом "results"
  document.removeEventListener('keydown', moveShooter);
  document.removeEventListener('keydown', shoot);
  clearInterval(invadersId);

  // Remove one heart (img element) if there are hearts left
  if (heart.children.length > 0) {
    heart.removeChild(heart.lastElementChild);
  }
  if (live <= 0) {
    isGameOver = true;
  }
}


function shoot(e) {
  if (isPaused || Kill || isWinner){
    return;
  }
  let laserId
  let currentLaserIndex = currentShooterIndex
  function moveLaser() {
    squares[currentLaserIndex].classList.remove('laser')
    currentLaserIndex -= width
    squares[currentLaserIndex].classList.add('laser')

    if (squares[currentLaserIndex].classList.contains('invader')) {
      squares[currentLaserIndex].classList.remove('laser')
      squares[currentLaserIndex].classList.remove('invader')
      squares[currentLaserIndex].classList.add('boom')

      setTimeout(()=> squares[currentLaserIndex].classList.remove('boom'), 300)
      clearInterval(laserId)

      const alienRemoved = alienInvaders.indexOf(currentLaserIndex)
      aliensRemoved.push(alienRemoved)
      score++
      scoreDisplay.innerHTML = score
    }
  }
  switch(e.key) {
    case 'ArrowUp':
      laserId = setInterval(moveLaser, 100)
  }
  
}

document.addEventListener('keydown', shoot)


function resetGrid() {
  squares.forEach(square => {
    square.classList.remove('invader', 'shooter', 'laser', 'boom');
  });


  alienInvaders.forEach(index => {
    squares[index].classList.remove('invader');
  });

  aliensRemoved = [];
  currentShooterIndex = 202;
  direction = 1;
  goingRight = true;
  time = 0;
  // live = 3;
  Kill = false;
  statusDIsplay.innerHTML = ''; // Очищаем содержимое, чтобы удалить текст
  squares[currentShooterIndex].classList.add('shooter');
  alienInvaders = alienRestart;
  scoreDisplay.innerHTML = score;
  timer.innerHTML = time;
}

function resetAllGame() {
  if (isGameOver || isWinner) {
    return;
  }
  resetGrid(); // Reset the grid to its default state
  AllGame(); // Start the game again

  // Remove the 'R' or 'r' key event listener when live is 0 to prevent further restarts
  if (live === 0) {
    document.removeEventListener('keydown', resetAllGame);
  }
}
// Listen for the 'Space' key to reset the game
document.addEventListener('keydown', function(e) {
  if (e.key === 'R' || e.key === 'r') {
    if (Kill){
    resetAllGame(); // Restart the game when Space key is pressed
  }
}
});
}
// Start the game for the first time
AllGame();