import React, {Component} from 'react';
import {View, StyleSheet, Image, ActivityIndicator, Alert} from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import Routes from './src/routes';
import 'react-native-gesture-handler';

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
          fcmRegistered: false,
        };
      }
    
    render() {
        return (
            <View style={{ flex: 1 }}>
                <NavigationContainer>
                    <Routes />
                </NavigationContainer>
            </View>
        );
    }
}

export default App