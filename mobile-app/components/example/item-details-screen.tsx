import { client } from "@/api-client/example";
import { useCallback, useEffect, useState } from "react";
import type { components } from "@/schema/example";
import { ActivityIndicator, Text, View } from "react-native";
import { useNavigation } from "@react-navigation/native";
import { Screens } from "./constants";
import type { Params } from "./type";
import type {
  NativeStackNavigationProp,
  NativeStackScreenProps,
} from "@react-navigation/native-stack";

type NavigationProp = NativeStackNavigationProp<
  Params,
  (typeof Screens)["Items"]
>;

type ItemDetailsScreenProps = NativeStackScreenProps<
  Params,
  (typeof Screens)["ItemDetails"]
>;

export default function ItemDetailsScreen({ route }: ItemDetailsScreenProps) {
  const navigation = useNavigation<NavigationProp>();
  const [isLoading, setLoading] = useState(true);
  const [item, setItem] = useState<components["schemas"]["Item"] | undefined>(
    undefined,
  );

  const getExampleItemById = useCallback(async (id: string) => {
    try {
      const response = await client.GET("/example/items/{id}", {
        params: { path: { id } },
      });
      setItem(response.data);
    } catch (error) {
      console.error("Error fetching example item by ID:", error);
    }
    setLoading(false);
  }, []);

  useEffect(() => {
    if (route.params?.id) {
      console.log("Fetching item with ID:", route.params.id);
      getExampleItemById(route.params.id);
    }
  }, [route, getExampleItemById]);

  return (
    <View>
      {isLoading ? (
        <ActivityIndicator />
      ) : (
        <Text style={{ padding: 20, color: "red", fontSize: 18 }}>
          <dl>
            <dt>ID:</dt>
            <dd>{item?.id}</dd>
          </dl>
          <dl>
            <dt>Name:</dt>
            <dd>{item?.name}</dd>
          </dl>
          <dl>
            <dt>Price:</dt>
            <dd>{item?.price}</dd>
          </dl>
        </Text>
      )}

      <Text
        style={{
          fontSize: 18,
          color: "blue",
          textAlign: "center",
          marginTop: 20,
        }}
        onPress={() => navigation.navigate(Screens.Items)}
      >
        Back to Items List
      </Text>
    </View>
  );
}
