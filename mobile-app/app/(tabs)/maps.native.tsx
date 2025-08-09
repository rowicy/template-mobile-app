import React from "react";
import { View, StyleSheet, Dimensions, ActivityIndicator } from "react-native";
import { LeafletView, MapShapeType } from "react-native-leaflet-view";

// Sample coordinates for the route
const tokyoTower = {
  latitude: 35.6586,
  longitude: 139.7454,
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
  const mapLayers = [
    {
      attribution:
        '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
      baseLayerIsChecked: true,
      baseLayerName: "OpenStreetMap",
      url: "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
    },
  ];

  const mapMarkers = [
    {
      position: tokyoTower,
      icon: "🏗️",
      size: [32, 32],
      title: "Tokyo Tower - Starting point",
      id: "tokyo-tower",
    },
    {
      position: convenienceStore1,
      icon: "🏪",
      size: [32, 32],
      title: "Convenience Store 1 - First stop",
      id: "store-1",
    },
    {
      position: convenienceStore2,
      icon: "🏪",
      size: [32, 32],
      title: "Convenience Store 2 - Second stop",
      id: "store-2",
    },
    {
      position: tokyoSkytree,
      icon: "🗼",
      size: [32, 32],
      title: "Tokyo Skytree - Final destination",
      id: "tokyo-skytree",
    },
  ];

  const mapShapes = [
    {
      shapeType: MapShapeType.POLYLINE,
      color: "#FF6B6B",
      id: "route",
      positions: routeCoordinates,
    },
  ];

  return (
    <View style={styles.container}>
      <LeafletView
        doDebug={false}
        renderLoading={() => <ActivityIndicator size="large" />}
        mapLayers={mapLayers}
        mapMarkers={mapMarkers}
        mapShapes={mapShapes}
        mapCenterPosition={tokyoTower}
        zoom={12}
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
