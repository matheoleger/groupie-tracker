const fetchSearchArtists = "/api/artists"
let artistsDAta;

const loadArtistsData =  (reponse) => {
    artistsDAta = [...reponse]
    console.log('res : ')
    console.log(artistsDAta)
}

fetch(fetchSearchArtists)
.then(response => response.json())
.then(loadArtistsData)