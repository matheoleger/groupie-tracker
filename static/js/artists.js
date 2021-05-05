const showMore = (artistID) => {

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

// for the search values (redirection)
if (window.location.href.indexOf("?") >= 0) {
    console.log("coucou")
    showMore(window.location.search.split("=")[1])
}

// Filter
let oneArtist = document.querySelector('#one-artist');
let group = document.querySelector('#group');
let nbrOfMembers = document.querySelector('#number-of-members');
let creationDateFilter = document.querySelector('#creation-date-filter');
let firstAlbumFilter = document.querySelector('#first-album-filter');

let oneArtistVal = true;
let groupVal = true;
let nbrOfMembersVal = nbrOfMembers.value;
let creationDateFilterVal = creationDateFilter.value;
let firstAlbumFilterVal = firstAlbumFilter.value;

oneArtist.addEventListener('click', () => {
    if(oneArtist.checked == true) {
        oneArtistVal = true
    } else {
        oneArtistVal = false
    }

    filter()    
})

group.addEventListener('click', () => {
    if(group.checked == true) {
        groupVal = true
    } else {
        groupVal = false
    }

    filter()
})

nbrOfMembers.addEventListener('click', () => {
    

    nbrOfMembersVal = nbrOfMembers.value

    //for the "style"
    const labelOfMem = document.querySelector('#nbr-of-mem-val');

    if (nbrOfMembers.value != 0) {
        labelOfMem.textContent = `≥${nbrOfMembers.value}`
    } else {
        labelOfMem.textContent = "-"
    }
    

    filter()
})

creationDateFilter.addEventListener('click', () => {
    creationDateFilterVal = creationDateFilter.value
    console.log(creationDateFilterVal)

    filter()
})

firstAlbumFilter.addEventListener('click', () => {
    firstAlbumFilterVal = firstAlbumFilter.value
    console.log(firstAlbumFilterVal)

    filter()
})

const filter = () => {
    fetch(`/artists/filter?artist=${oneArtistVal}&group=${groupVal}&members=${nbrOfMembersVal}&creationDates=${creationDateFilterVal}&firstAlbum=${firstAlbumFilterVal}`)
    .then(response => response.json())
    .then(displayFilterEl)
}

const displayFilterEl = (filteredElements) => {
    const divArtistsContent = document.querySelector('.artists-el')
    removeAllChildNodes(divArtistsContent)

    let htmlBuff = ""

    for (element of filteredElements) {
        htmlBuff += `
        <div class="artist">
            <img src="${element.Image}" class="artists-img">
            <div class="artist-text-el">
                <span class="artist-name">${element.name}</span>
                <div> 
                    <span class="key-artist">Creation date</span><span>${element.CreationDate}</span>
                </div>
                <div> 
                    <span class="key-artist">First Album</span><span>${element.FirstAlbum}</span>                  
                </div>     
                <button id="${element.id}" class="btn-showmore" onclick="showMore(this.id)">Voir plus...</button>
            </div>   
        </div>
        `
        divArtistsContent.innerHTML = htmlBuff
    }

    //for the "style"
    const artistsContent = document.querySelector(".artists-el")

    let nbrRows = Math.ceil(filteredElements.length / 3)

    artistsContent.style.gridTemplateRows = `repeat(${nbrRows}, 200px)`


}

