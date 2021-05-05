
 /*
 // Initialize and add the map
 function initMap() {
  // The location 
  const positionPoint = { lat: 47.218371, lng: -1.553621
  };
  // The map, centered at Uluru
  const map = new google.maps.Map(document.getElementById("map"), {
    zoom: 3,
    center: positionPoint,
  });
  // The marker, positioned at Uluru
  const marker = new google.maps.Marker({
    position: positionPoint,
    map: map,

  });


  addMarker({lat: 11.5564, lng: 104.9282})
  addMarker({lat: 47.0667 , lng: -0.8833})
  addMarker({lat: 48.866667, lng:  2.333333})



  function addMarker(coords){
    let marker = new google.maps.Marker({
      position: coords,
      map:map
    })
  }
*/






function initMap() {
  const map = new google.maps.Map(document.getElementById("map"), {
    zoom: 4,
    center: { lat: 46.227638
      , lng: 2.213749 },
  });
  const geocoder = new google.maps.Geocoder();
  document.getElementById("submit").addEventListener("click", () => {
    geocodeAddress(geocoder, map);
  });
}

function geocodeAddress(geocoder, resultsMap) {

  const test =  window.location.search.split("=")[2];
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
