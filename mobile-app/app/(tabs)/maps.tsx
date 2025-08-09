import React from "react";
import { View, StyleSheet, Dimensions } from "react-native";

export default function MapsScreen() {
  return (
    <View style={styles.container}>
      <iframe
        style={styles.map}
        src={`data:text/html;charset=utf-8,${encodeURIComponent(`
<!DOCTYPE html>
<html>
<head>
    <title>OpenStreetMap</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css" integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY=" crossorigin=""/>
    <script src="https://unpkg.com/leaflet@1.9.4/dist/leaflet.js" integrity="sha256-20nQCchB9co0qIjJZRGuk2/Z9VM+kNiyxNV1lvTlZBo=" crossorigin=""></script>
    <style>
        body { margin: 0; padding: 0; }
        #map { height: 100vh; width: 100vw; }
    </style>
</head>
<body>
    <div id="map"></div>
    <script>
        // Sample coordinates for the route
        const tokyoTower = [35.6586, 139.7454];
        const convenienceStore1 = [35.6762, 139.7654];
        const convenienceStore2 = [35.6895, 139.7456];
        const tokyoSkytree = [35.7101, 139.8107];

        // Initialize the map
        const map = L.map('map').setView(tokyoTower, 12);

        // Add OpenStreetMap tiles
        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        }).addTo(map);

        // Add markers
        L.marker(tokyoTower).addTo(map)
            .bindPopup('<strong>Tokyo Tower</strong><br>Starting point');
        
        L.marker(convenienceStore1).addTo(map)
            .bindPopup('<strong>Convenience Store 1</strong><br>First stop');
        
        L.marker(convenienceStore2).addTo(map)
            .bindPopup('<strong>Convenience Store 2</strong><br>Second stop');
        
        L.marker(tokyoSkytree).addTo(map)
            .bindPopup('<strong>Tokyo Skytree</strong><br>Final destination');

        // Add route polyline
        const routeCoordinates = [tokyoTower, convenienceStore1, convenienceStore2, tokyoSkytree];
        L.polyline(routeCoordinates, {color: '#FF6B6B', weight: 3}).addTo(map);
    </script>
</body>
</html>
        `)}`}
        title="OpenStreetMap"
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  map: {
    width: Dimensions.get("window").width,
    height: Dimensions.get("window").height,
  },
});
