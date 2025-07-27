import { $api } from "@/api-client/example";
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

  // Use TanStack Query hook for fetching item by ID
  const { data: item, isLoading } = $api.useQuery(
    "get",
    "/example/items/{id}",
    {
      params: { path: { id: route.params?.id || "" } },
    },
    {
      enabled: !!route.params?.id, // Only run query if ID is available
    },
  );

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
