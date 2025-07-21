import { client } from "@/api-client/example";
import { useCallback, useEffect, useState } from "react";
import type { components } from "@/schema/example";
import { ActivityIndicator, FlatList, Text, View } from "react-native";
import { useNavigation } from "@react-navigation/native";
import { Screens } from "./constants";
import type { Params } from "./type";
import type { NativeStackNavigationProp } from "@react-navigation/native-stack";

type NavigationProp = NativeStackNavigationProp<
  Params,
  (typeof Screens)["ItemDetails"]
>;

export default function ItemsScreen() {
  const navigation = useNavigation<NavigationProp>();

  const [isLoading, setLoading] = useState(true);
  const [items, setItems] = useState<
    components["schemas"]["Item"][] | undefined
  >(undefined);

  const getExampleItems = useCallback(async () => {
    try {
      const response = await client.GET("/example/items");
      setItems(response.data);
    } catch (error) {
      console.error("Error fetching example items:", error);
    }
    setLoading(false);
  }, []);

  useEffect(() => {
    getExampleItems();
  }, [getExampleItems]);

  const handleItemPress = (id: string) => {
    navigation.navigate(Screens.ItemDetails, { id });
  };

  return (
    <View style={{ flex: 1, padding: 24 }}>
      {isLoading ? (
        <ActivityIndicator />
      ) : (
        <FlatList
          data={items}
          keyExtractor={(item) => item.id.toString()}
          renderItem={({ item }) => (
            <View
              style={{
                padding: 20,
                borderBottomWidth: 1,
                borderBottomColor: "#ccc",
              }}
            >
              <Text
                style={{ fontSize: 18, color: "blue" }}
                onPress={() => handleItemPress(item.id.toString())}
              >
                {item.name}
              </Text>
            </View>
          )}
        />
      )}
    </View>
  );
}
