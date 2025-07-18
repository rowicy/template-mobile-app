import { createNativeStackNavigator } from "@react-navigation/native-stack";
import ItemsScreen from "./ItemsScreen";
import ItemDetailsScreen from "./ItemDetailsScreen";
import { Screens } from "./constants";
import type { Params } from "./types";

const Stack = createNativeStackNavigator<Params>();

export default function Example() {
  return (
    <Stack.Navigator initialRouteName={Screens.Items}>
      <Stack.Screen name={Screens.Items} component={ItemsScreen} />
      <Stack.Screen name={Screens.ItemDetails} component={ItemDetailsScreen} />
    </Stack.Navigator>
  );
}
