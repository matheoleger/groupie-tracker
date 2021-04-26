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


let fetchSearchElements = ["/api/artists", "/api/locations"]
let searchArtistsData;
let searchLocationsData;
let i = 0
// file = "search"

const loadSearchData = (response) => {
    // console.log("response:")
    // console.log(response)

    if (i == 1) {
        searchArtistsData = [...response]
        console.log('this is')
        console.log(searchArtistsData)
    } else {
        searchLocationsData = {...response}
        console.log('this is')
        console.log(searchLocationsData)
    }
    console.log(i)
    i++
}

fetch(fetchSearchElements[0])
.then(response => response.json())
.then(loadSearchData)

fetch(fetchSearchElements[1])
.then(response => response.json())
.then(loadSearchData)


const search = (k) => {
    // console.log(k.key)
    
    let searchbar = document.querySelector('#searchbar');
    console.log(searchbar.value)

    let regex = RegExp(searchbar.value.toUpperCase());

    // if (regex.test(el.children[1].textContent.toUpperCase()) == true) {

    for (element of searchArtistsData) {
        let divEl = document.getElementById(element.id)

        if (regex.test(element.name.toUpperCase()) == true) {
            if(divEl == null || element.id != divEl.id) {
                displayAnElement(element)
            }
            
        } else {

            if (divEl != null) {
                divEl.remove()
            }
            
        }
    }

}

const displayAnElement = (element) => {
    console.log(element.name)

    const listAutocomp = document.querySelector('.list-autocomp')
    let divEl = document.createElement("div")
    divEl.setAttribute('id', element.id)
    let textEl = document.createTextNode(element.name)

    divEl.appendChild(textEl)
    listAutocomp.append(divEl)
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

