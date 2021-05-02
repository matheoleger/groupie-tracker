//const fetchSearchArtists = "/api/artists"

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
        nameArtists.innerText = artistsData[i]['name']
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

/*

/*const loadArtistsData =  (reponse) => {
    artistsData = [...reponse]
    console.log('res : ')
    console.log(artistsData)
}


fetch(fetchSearchArtists)
.then(response => response.json())
.then(loadArtistsData)*/


const showMore = (artistID) => {
    // const artist = document.querySelector(artistID)

    fetch(`/artists/showmore?id=${artistID}`)
    .then(response => response.json())
    .then(displayMore)
}

const displayMore = (artistInformations) => {
    console.log("ceci est artistInfor.. : ")
    console.log(artistInformations)

    
}
