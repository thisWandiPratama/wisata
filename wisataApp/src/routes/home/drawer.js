import * as React from 'react';
import { createDrawerNavigator } from '@react-navigation/drawer';

const Drawer = createDrawerNavigator();

import Profile_Admin from '../../components/drawerContent';
import Tab_Admin from './stack';

export default function Drawer_Admin() {
  return (
      <Drawer.Navigator  screenOptions={{ headerShown: false }} drawerContent={(props)=> <Profile_Admin {...props}/>}>
        <Drawer.Screen name="Tab_Admin" component={Tab_Admin} />
      </Drawer.Navigator>
  );
}