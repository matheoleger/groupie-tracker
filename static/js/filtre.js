const howMany = (artistID) => {
    // const artist = document.querySelector(artistID)

    fetch(`/artists/howMany?id=${artistID}`)
    .then(response => response.json())
    .then(displayMore)
}