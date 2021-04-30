//print the artists on the page
function getArtist() {
    let artistsData = [...searchArtistsData];
    let band;
    let nameBand;
    let nameArtists;
    let creationDate;
    let firstDate;
    for (let i =0; i < artistsData.length;i++) {
        band = document.createElement('div')
        band.classList.add('band')
        nameBand = document.createElement('h2')
        nameBand.innerText = artistsData[i]['name']
        nameArtists = document.createElement('p')
        nameArtists.innerText = artistsData[i]['members']
        creationDate = document.createElement('p')
        creationDate.innerText =artistsData[i]['creationDate']
        firstDate = document.createElement('p')
        firstDate.innerText = artistsData[i]['firstAlbum']
        band.appendChild(nameBand)
        band.appendChild(nameArtists)
        band.appendChild(creationDate)
        band.appendChild(firstDate)
        document.querySelector('main').appendChild(band)
    }
     
}
