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

    //for the "style"
    const showMoreElement = document.querySelector('.bg-showmore')
    showMoreElement.style.display = "flex";

    const globalContent = document.querySelector(".global-content")
    const memberContent = document.querySelector(".member-content")
    const locationContent = document.querySelector('.location-content')

    removeAllChildNodes(globalContent)
    removeAllChildNodes(memberContent)
    removeAllChildNodes(locationContent)
}

const displayMore = (artistInformations) => {
    console.log("ceci est artistInfor.. : ")
    console.log(artistInformations)

    const globalContent = document.querySelector(".global-content")
    const memberContent = document.querySelector(".member-content")
    const locationContent = document.querySelector('.location-content')

    // removeAllChildNodes(globalContent)
    // removeAllChildNodes(memberContent)
    // removeAllChildNodes(locationContent)

    console.log(artistInformations.Artist.name)

    //Global Content
    for (const [key, value] of Object.entries(artistInformations.Artist)) {

        console.log(key)
        console.log(value)

        if (key == "Image") {
            let contentImage = document.createElement('img');
            contentImage.setAttribute("src", value);
            globalContent.append(contentImage);
            continue
        } else if (key == "id" || key == "Members") {
            continue
        }
        
        let divElement = document.createElement('div')
        divElement.classList.add('global-data')

       
        let keySpan = document.createElement('span')
        let textKey = document.createTextNode(key);
        keySpan.classList.add('key-span-showmore')
        keySpan.appendChild(textKey);

        let valueSpan = document.createElement('span')
        let textValue = document.createTextNode(value)
        valueSpan.appendChild(textValue)

        divElement.appendChild(keySpan)
        divElement.appendChild(valueSpan)
        globalContent.append(divElement)
    }

    for (member of artistInformations.Artist.Members) {
        let valueSpan = document.createElement('span')
        let textValue = document.createTextNode(member)
        valueSpan.appendChild(textValue)

        memberContent.append(valueSpan)
    }

    for(const [key, values] of Object.entries(artistInformations.Relation.DatesLocations)) {
        let locationDiv = document.createElement('div')
        locationDiv.classList.add('location-dates-content')

        let locationKey = document.createElement('h2')
        let locationKeyText = document.createTextNode(key)
        
        locationKey.appendChild(locationKeyText)
        locationDiv.append(locationKey)

        for (value of values) {
            let valueSpan = document.createElement('span')
            let textValue = document.createTextNode(value)
            valueSpan.appendChild(textValue)
            locationDiv.append(valueSpan)
        }

        locationContent.append(locationDiv)
    }
}

// for the "style"
const buttonClose = document.querySelector('.more-close')
buttonClose.addEventListener('click', () => {

    const showMoreElement = document.querySelector('.bg-showmore')
    showMoreElement.style.display = "none";
})
