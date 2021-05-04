const filtre = (artistID) => {
    // const artist = document.querySelector(artistID)
    console.log(artistID)

    fetch(`/artists/howMany`)
    .then(response => response.json())
    .then(displayFiltre)
}