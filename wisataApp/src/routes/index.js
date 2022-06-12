import * as React from 'react';
import { createNativeStackNavigator } from '@react-navigation/native-stack';


import Splashscreen from '../auths/splah'
import Login from '../auths/login'
import Register from '../auths/register'
import Home from './home/drawer';

const Stack = createNativeStackNavigator();

function Routes() {
    return (
            <Stack.Navigator screenOptions={{ headerShown: false }}>
                <Stack.Screen name="Splashscreen" component={Splashscreen} />
                <Stack.Screen name="Login" component={Login} />
                <Stack.Screen name="Register" component={Register} />
                <Stack.Screen name="Home" component={Home} />
            </Stack.Navigator>
    );
}

export default Routes;