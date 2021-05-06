const search = () => {
    const searchbar = document.querySelector('#searchbar');
    const listAutocomp = document.querySelector('.list-autocomp')

    //pour enlever la liste quand elle est vide
    listAutocomp.style.display = "";
    //pour aggrandir la case de recherche
    searchbar.classList.add('biggest-searching')

    if (searchbar.value.length == 0) {

        searchbar.classList.remove('biggest-searching')
        listAutocomp.style.display = "none";
    
    } else {
        fetch(`/search?search=${searchbar.value}`)
        .then(response => response.json())
        .then((data) => {
            console.log(data)
            
            removeAllChildNodes(listAutocomp)
            for (element of data) {

                displayResult(element)
            }

            document.querySelectorAll('.list-element').forEach(item => {
                item.addEventListener('click', event => {

                    let actualLoc = item.querySelectorAll("span")

                        if (item.classList.contains('locations')) {
                            window.location.href=`/locations?searchId=${item.id}&value=${actualLoc[1].textContent}`
                        } else {
                            window.location.href=`/artists/?searchId=${item.id}`
                        }
                })
            })
        })
    }

}

const createGlobalElement = (divForSpan, spanValue, element, contentData) => {

    //add image
    if(element != undefined && contentData != undefined) {
        let contentImage = document.createElement('img');
        contentImage.setAttribute("src", element.Image);
        contentData.appendChild(contentImage);
    }

    //add key or value
    let contentSpan = document.createElement('span');
    let textKey = document.createTextNode(spanValue);
    contentSpan.appendChild(textKey);

    if (spanValue == "Locations") {
        contentData.classList.add('locations')
    }

    divForSpan.appendChild(contentSpan)
    
}

const displayResult = (element) => {
    const listAutocomp = document.querySelector('.list-autocomp');


    for (const [key, value] of Object.entries(element)) {

        
        if (value == "" || value == 0 || value == null || key == "Id" || key == "Image"){
            continue
        } else if (Array.isArray(value)) {
                    
            //add members or locations
            for (el of value) {
                let contentData = document.createElement("div")
                contentData.setAttribute('id', element.Id)
                contentData.classList.add("list-element")

                let divForSpan = document.createElement('div')
                divForSpan.classList.add("div-of-span")


                createGlobalElement(divForSpan, key, element, contentData)
                createGlobalElement(divForSpan, el)

                contentData.append(divForSpan)
                listAutocomp.append(contentData)
            }

        } else {
            let contentData = document.createElement("div")
            contentData.setAttribute('id', element.Id)
            contentData.classList.add("list-element")

            let divForSpan = document.createElement('div')
            divForSpan.classList.add("div-of-span")

            createGlobalElement(divForSpan, key, element, contentData)
            createGlobalElement(divForSpan, value)

            contentData.append(divForSpan)
            listAutocomp.append(contentData)
        }

    } 
}

//pour enlever tous les elements HTML d'un element
function removeAllChildNodes(parent) {
    while (parent.firstChild) {
        parent.removeChild(parent.firstChild);
    }
}

