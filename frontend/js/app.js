document.getElementById('showRegister').addEventListener('click', function() {
    document.getElementById('auth').style.display = 'none';
    document.getElementById('register').style.display = 'block';
});

document.getElementById('showLogin').addEventListener('click', function() {
    document.getElementById('register').style.display = 'none';
    document.getElementById('auth').style.display = 'block';
});

document.getElementById('authForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    
    fetch('/api/auth', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            document.getElementById('auth').style.display = 'none';
            document.getElementById('reservation').style.display = 'block';
            loadReservations();
        } else {
            alert('Authentication failed');
        }
    });
});

document.getElementById('registerForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const username = document.getElementById('regUsername').value;
    const password = document.getElementById('regPassword').value;
    
    fetch('/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            alert('Registration successful. Please log in.');
            document.getElementById('register').style.display = 'none';
            document.getElementById('auth').style.display = 'block';
        } else {
            alert('Registration failed');
        }
    });
});

document.getElementById('searchButton').addEventListener('click', function() {
    const criteria = document.getElementById('searchCriteria').value;
    fetch(`/api/cars?criteria=${criteria}`)
    .then(response => response.json())
    .then(cars => {
        const carList = document.getElementById('carList');
        carList.innerHTML = '';
        cars.forEach(car => {
            const carCard = document.createElement('div');
            carCard.className = 'car-card';
            carCard.innerHTML = `<h3>${car.model}</h3><p>${car.details}</p><button onclick="reserveCar(${car.id})">Reserve</button>`;
            carList.appendChild(carCard);
        });
    });
});

function loadReservations() {
    fetch('/api/reservations')
    .then(response => response.json())
    .then(reservations => {
        const reservationList = document.getElementById('reservationList');
        reservationList.innerHTML = '';
        reservations.forEach(reservation => {
            const li = document.createElement('li');
            li.textContent = reservation.details;
            reservationList.appendChild(li);
        });
    });
}

function reserveCar(carId) {
    fetch('/api/reservations', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ car_id: carId })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            loadReservations();
        } else {
            alert('Reservation failed');
        }
    });
}