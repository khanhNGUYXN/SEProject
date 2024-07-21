document.addEventListener("DOMContentLoaded", function () {
    const addForm = document.getElementById('addMovieForm');
    addForm.onsubmit = function (event) {
        event.preventDefault();
        const formData = new FormData(addForm);
        fetch('/movies', {
            method: 'POST',
            body: formData
        }).then(response => response.json())
            .then(data => {
                console.log('Success:', data);
                alert('Movie added!');
                location.reload(); // Reload to see new movie in the list
            })
            .catch((error) => {
                console.error('Error:', error);
                alert('Failed to add movie');
            });
    };
});

function deleteMovie(movieID) {
    fetch(`/movies/${movieID}`, {
        method: 'DELETE',
    }).then(response => {
        if (response.ok) {
            alert('Movie deleted!');
            location.reload(); // Reload to update the list
        } else {
            alert('Failed to delete movie');
        }
    });
}
