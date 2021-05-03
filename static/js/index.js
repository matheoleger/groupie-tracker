// file = "index"

const loadIndexData = (response) => {
    console.log(response)
}

fetch("/api/artists/1")
.then(response => response.json())
.then(loadData)