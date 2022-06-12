import * as React from 'react';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

import Homedata from '../../containers/home'
import ByCategory from '../../components/bycategory'
import ListItinerary from '../../components/listItinerary'
import DetailItinerary from '../../components/detailItinerary'
import DetailTourist from '../../components/detailTourist'
import AddItinerary from '../../components/addItinerary'
import DetailRoute from '../../components/detailRoute'
import MapItinerary from '../../components/mapItinerary'
import Profile from '../../components/profile'

const Stack = createNativeStackNavigator();

function Routes() {
    return (
            <Stack.Navigator initialRouteName='Homedata' screenOptions={{ headerShown: false }}>
                <Stack.Screen name="Homedata" component={Homedata} />
                <Stack.Screen name="ByCategory" component={ByCategory} />
                <Stack.Screen name="ListItinerary" component={ListItinerary} />
                <Stack.Screen name="DetailItinerary" component={DetailItinerary} />
                <Stack.Screen name="DetailTourist" component={DetailTourist} />
                <Stack.Screen name="AddItinerary" component={AddItinerary} />
                <Stack.Screen name="DetailRoute" component={DetailRoute} />
                <Stack.Screen name="MapItinerary" component={MapItinerary} />
                <Stack.Screen name="Profile" component={Profile} />
            </Stack.Navigator>
    );
}

export default Routes;