
const bookList = document.getElementById('book-list');
const addBookForm = document.getElementById('add-book-form');
const addBookBtn = document.getElementById('add-book-btn');
const updateBookForm = document.getElementById('update-book-form');
const updateBookBtn = document.getElementById('update-book-btn');

// Get books from API
function getBooks() {
    fetch('http://localhost:8080/books/all_books')
        .then(response => response.json())
        .then(data => {
            bookList.innerHTML = '';
            data.forEach(book => {
                const bookListItem = document.createElement('li');
                bookListItem.textContent = `${book.title} by ${book.author}`;
                bookListItem.id = book.id;
                bookList.appendChild(bookListItem);
            });
        })
        .catch(error => console.error('Error:', error));
}

// Add book to API
function addBook(event) {
    event.preventDefault();
    const title = document.getElementById('title').value;
    const author = document.getElementById('author').value;
    const book = { title, author };

    fetch('http://localhost:8080/books/create', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(book),
    })
        .then(response => response.json())
        .then(data => {
            getBooks();
            addBookForm.reset();
        })
        .catch(error => console.error('Error:', error));
}

// Update book in API
function updateBook(event) {
    event.preventDefault();
    const id = document.getElementById('update-id').value;
    const title = document.getElementById('update-title').value;
    const author = document.getElementById('update-author').value;
    const book = { title, author };

    fetch(`http://localhost:8080/books/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(book),
    })
        .then(response => response.json())
        .then(data => {
            getBooks();
            updateBookForm.reset();
        })
        .catch(error => console.error('Error:', error));
}

// Delete book from API
function deleteBook(event) {
    event.preventDefault();
    const id = event.target.parentNode.id;

    fetch(`http://localhost:8080/books/${id}`, {
        method: 'DELETE',
    })
        .then(response => response.json())
        .then(data => {
            getBooks();
        })
        .catch(error => console.error('Error:', error));
}

// Event listeners
addBookBtn.addEventListener('click', addBook);
updateBookBtn.addEventListener('click', updateBook);

bookList.addEventListener('click', (event) => {
    if (event.target.tagName === 'BUTTON') {
        deleteBook(event);
    }
});

// Initialize book list
getBooks();