fetch("/api/artists")
.then((response) => response.json())
.then(loadData)


function loadData(response) {
    console.log(response)
}

// console.log(JSON.parse(locationData))

// console.log(JSON.parse(data))

// console.log(locationData)


// let JSONData = data.split("]")

// console.log(JSONData)

// console.log(JSON.parse(JSONData[0]))

// function getData() {
    
// }