
 
  // Initialize and add the map
  function initMap() {
    // The location 
    const positionPoint = { lat: 49.218371, lng: -1.56 };
    // The map, centered at Uluru
    const map = new google.maps.Map(document.getElementById("map"), {
      zoom: 10,
      center: positionPoint,
    });
    // The marker, positioned at Uluru
    const marker = new google.maps.Marker({
      position: positionPoint,
      map: map,
    });
  }
