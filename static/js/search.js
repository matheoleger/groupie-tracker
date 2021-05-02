// let fetchElements = ["/api/artists", "/api/locations"]
// // let fetchArtists = "/api/artists"
// // let fetchLocations = "/api/locations"

// let searchArtists;
// let searchLocations;

// for (element of fetchElements) {
//     console.log(element)
//     const loadDataSearch = (responseData) => {

//         // let testURL = URL.createObjectURL(responseData)
//         // console.log(testURL)
        
//         if (element == "/api/artists") {
//             searchArtists = [...responseData]
//         } else{
//             searchLocations = {...responseData}
//         }
           
//     }

//     fetch(element)
//     .then(response => response.json())
//     .then(loadDataSearch)
// }

// console.log(searchArtists)
// console.log(searchLocations)




// let fetchSearchElements = ["/api/artists", "/api/locations"]
// // let fetchArtists = "/api/artists"
// // let fetchLocations = "/api/locations"

// let searchArtists;
// let searchLocations;

// for (element of fetchSearchElements) {
//     console.log(element)

//     actualFetchEl = element;
//     console.log("actualFetch in search " + actualFetchEl)
//     file = "search";

//     fetch(element)
//     .then(response => response.json())
//     .then(loadData)

//     console.log("quand cest log ca ?")

//     // console.log("test bruh: " + responseData)

//     // if (element == "/api/artists") {
//     //     searchArtists = [...responseData]
//     // } else{
//     //     searchLocations = {...responseData}
//     // }

//     // console.log("test bruh12: " + responseData)
// }

// console.log("test bruh: " + searchArtists)
// console.log(searchLocations)


// let fetchSearchElements = ["/api/artists", "/api/locations"]
let searchArtistsData;
let searchLocationsData;
let i = 0
// file = "search"

const loadSearchData = (response) => {
    // console.log("response:")
    // console.log(response)

    // if (i == 0) {
    //     searchArtistsData = [...response]
    //     console.log('this is')
    //     console.log(searchArtistsData)
    // } else {
    //     searchLocationsData = {...response}
    //     console.log('this is')
    //     console.log(searchLocationsData)
    // }
    // console.log(i)
    // i++

    searchArtistsData = [...response];

}

fetch("/api/artists")
.then(response => response.json())
.then(loadSearchData)

// fetch(fetchSearchElements[1])
// .then(response => response.json())
// .then(loadSearchData)

fetch("/api/locations")
.then(response => response.json())
.then((response) => {
    searchLocationsData = {...response}
    console.log(searchLocationsData)
})


const search = () => {
    // console.log(k.key)
    
    const searchbar = document.querySelector('#searchbar');
    const listAutocomp = document.querySelector('.list-autocomp')

    //pour enlever la liste quand elle est vide
    listAutocomp.style.display = "";

    //pour aggrandir la case de recherche
    searchbar.classList.add('biggest-searching')

    // let rightElements = searchArtistsData.filter(artist => {

    //     const regex = RegExp(searchbar.value.toUpperCase());

    //     for (let i = 0; i < artist.members.length; i++) {
    //         if (artist.members[i].toUpperCase().match(regex)) {

    //         }
    //     }

    //     return artist.name.toUpperCase().match(regex)
    // })

    // console.log(rightElements)

    removeAllChildNodes(listAutocomp)

    if (searchbar.value.length == 0) {

        searchbar.classList.remove('biggest-searching')
        listAutocomp.style.display = "none";
        
        // rightElements = []
    }

    const regex = RegExp(searchbar.value.toUpperCase());

    for (element of searchArtistsData) {
        
        // let obj = {
        //     name: "",
        //     members: "",
        //     firstAlbum: 0,
        //     creationDate: "0",
        //     location: "",
        // }

        let obj = Object.create(element)
        

        if(regex.test(element.name.toUpperCase())) {
            obj.name = element.name
        }

        if(regex.test(element.firstAlbum.toUpperCase())) {
            obj.firstAlbum = element.firstAlbum
        }

        if(regex.test(element.creationDate.toString().toUpperCase())) {
            obj.creationDate = element.creationDate
        }

        

        for(member of element.members) {
            if(regex.test(member.toString().toUpperCase())) {
                // obj.members.push(member)
                console.log(member)

                let newMembers = []
                newMembers.push(member)
                obj.members = newMembers
            }
        }

        
        
        // obj.firstAlbum = "19 701"
        if (Object.keys(obj).length != 0) {
            console.log(obj)
            obj.image = element.image
            displayAnElement(obj)
        }
        

    }

    for (element of searchLocationsData.index) {
        console.log(element)


        for (eachLocation of element.locations) {

            let isSameLocRegex = RegExp(eachLocation)

            let objLoc = {};
            if(regex.test(eachLocation.toString().toUpperCase())) {
               objLoc = {locations: eachLocation}
            }

            if (Object.keys(objLoc).length != 0) {
                objLoc.image = "../static/img/pin-w.png"
                displayAnElement(objLoc)
            }
        }
    }

    
    // for (element of rightElements) {
    //     displayAnElement(element)
    // }


    // console.log(searchbar.value)

    // let regex = RegExp(searchbar.value.toUpperCase());

    // // if (regex.test(el.children[1].textContent.toUpperCase()) == true) {

    // for (element of searchArtistsData) {
    //     let divEl = document.getElementById(element.id)

    //     if (regex.test(element.name.toUpperCase()) == true) {
    //         if(divEl == null || element.id != divEl.id) {
    //             displayAnElement(element)
    //         }
            
    //     } else {

    //         if (divEl != null) {
    //             divEl.remove()
    //         }
            
    //     }
    // }



}

const displayAnElement = (element) => {

    //list of the elements
    const listAutocomp = document.querySelector('.list-autocomp')
    const listLocations = document.querySelectorAll('.locations')
    
    listLocations.filter()

    for (const [key, value] of Object.entries(element)) {

        // div of one element
        let divEl = document.createElement("div")
        divEl.setAttribute('id', element.id)
        divEl.classList.add("list-element")

        if (key == "image") {
            continue
        }

        let contentImage = document.createElement('img');
        contentImage.setAttribute("src", element.image);
        divEl.appendChild(contentImage);


        let divForSpan = document.createElement('div')
        divForSpan.classList.add("div-of-span") 

        let contentKeySpan = document.createElement('span');
        let textKey = document.createTextNode(key);
        contentKeySpan.appendChild(textKey);
        
        let contentTextSpan = document.createElement('span');
        let textEl = document.createTextNode(value);
        if (element.locations != "undefined") {
            // console.log("je suis a " + element.locations)
            contentTextSpan.classList.add('locations')
        }
        contentTextSpan.appendChild(textEl);


        divForSpan.appendChild(contentKeySpan)
        divForSpan.appendChild(contentTextSpan)
        divEl.appendChild(divForSpan)
        listAutocomp.append(divEl)
    }

}



// const displayAnElement = (element) => {
//     console.log(element.name)

//     //list of the elements
//     const listAutocomp = document.querySelector('.list-autocomp')

//     // div of one element
//     let divEl = document.createElement("div")
//     divEl.setAttribute('id', element.id)
//     divEl.classList.add("list-element")

//     // content of one element
//     contentImage = document.createElement('img');
//     contentImage.setAttribute("src", element.image);
//     divEl.appendChild(contentImage);


//     divForSpan = document.createElement('div')
//     divForSpan.classList.add("div-of-span") 

//     contentKeySpan = document.createElement('span');
//     let textKey = document.createTextNode(Object.keys(element)[2]);
//     contentKeySpan.appendChild(textKey);
    
//     contentTextSpan = document.createElement('span');
//     let textEl = document.createTextNode(element.name);
//     contentTextSpan.appendChild(textEl);

//     divForSpan.appendChild(contentKeySpan)
//     divForSpan.appendChild(contentTextSpan)
//     divEl.appendChild(divForSpan)
//     listAutocomp.append(divEl)
// }

function removeAllChildNodes(parent) {
    while (parent.firstChild) {
        parent.removeChild(parent.firstChild);
    }
}

// searchbar = document.querySelector('#searchbar')

// searchbar.addEventListener('keyup', search)






// searchArtists = [...responseData]
// console.log(searchArtists)

// fetch(fetchLocations)
// .then(response => response.json())
// .then(loadData)

// searchLocations = [...responseData]
// console.log(searchLocations)

