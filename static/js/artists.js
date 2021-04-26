const fetchSearchArtists = "/api/artists"
const artistsDAta;

const loadArtistsData =  (reponse) => {
    artistsDAta = [...reponse]
    console.log('res : ' + artistsDAta)
}