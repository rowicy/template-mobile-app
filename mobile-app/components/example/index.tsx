import { createNativeStackNavigator } from "@react-navigation/native-stack";
import ItemsScreen from "./items-screen";
import ItemDetailsScreen from "./item-details-screen";
import { Screens } from "./constants";
import type { Params } from "./type";

const Stack = createNativeStackNavigator<Params>();

export default function Example() {
  return (
    <Stack.Navigator initialRouteName={Screens.Items}>
      <Stack.Screen name={Screens.Items} component={ItemsScreen} />
      <Stack.Screen name={Screens.ItemDetails} component={ItemDetailsScreen} />
    </Stack.Navigator>
  );
}
