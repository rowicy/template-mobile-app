import React from "react";
import { View, StyleSheet, Dimensions } from "react-native";
// @ts-ignore - react-native-maps has TypeScript compatibility issues with strict mode
import MapView, { Marker, Polyline } from "react-native-maps";

// Sample coordinates for the route
const tokyoTower = {
  latitude: 35.6586,
  longitude: 139.7454,
  latitudeDelta: 0.05,
  longitudeDelta: 0.05,
};

const convenienceStore1 = {
  latitude: 35.6762,
  longitude: 139.7654,
};

const convenienceStore2 = {
  latitude: 35.6895,
  longitude: 139.7456,
};

const tokyoSkytree = {
  latitude: 35.7101,
  longitude: 139.8107,
};

// Route coordinates for the polyline
const routeCoordinates = [
  tokyoTower,
  convenienceStore1,
  convenienceStore2,
  tokyoSkytree,
];

export default function MapsScreen() {
  return (
    <View style={styles.container}>
      <MapView style={styles.map} initialRegion={tokyoTower}>
        {/* Markers for each location */}
        <Marker
          coordinate={tokyoTower}
          title="Tokyo Tower"
          description="Starting point"
          pinColor="green"
        />
        <Marker
          coordinate={convenienceStore1}
          title="Convenience Store 1"
          description="First stop"
          pinColor="blue"
        />
        <Marker
          coordinate={convenienceStore2}
          title="Convenience Store 2"
          description="Second stop"
          pinColor="blue"
        />
        <Marker
          coordinate={tokyoSkytree}
          title="Tokyo Skytree"
          description="Final destination"
          pinColor="red"
        />

        {/* Route polyline */}
        <Polyline
          coordinates={routeCoordinates}
          strokeColor="#FF6B6B"
          strokeWidth={3}
          lineCap="round"
          lineJoin="round"
        />
      </MapView>
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
