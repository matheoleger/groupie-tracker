 // Fonction qui va initialiser la map
function initMap() {
  const map = new google.maps.Map(document.getElementById("map"), {
    zoom: 7,
    center: { lat: 46.227638
      , lng: 2.213749 },
  });
  const geocoder = new google.maps.Geocoder();
 if (window.location.href.indexOf("?") >= 0){
   geocodeAddress(geocoder, map);
 }
}

  // Cette fonction de traduction -> nom de ville en coordonnées GPS -> géolocaliser un lieu.

function geocodeAddress(geocoder, resultsMap) {

  const test =  window.location.search.split("=")[2]; // recupère la value dans l' url qui correspond a la ville
  geocoder.geocode({ address: test}, (results, status) => {
    if (status === "OK") {
      resultsMap.setCenter(results[0].geometry.location);
      new google.maps.Marker({
        map: resultsMap,
        position: results[0].geometry.location,
      });
    } else {
      alert("Geocode was not successful for the following reason: " + status);
    }
  });
}
