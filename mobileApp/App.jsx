import React, { useEffect } from 'react';
import { SafeAreaView } from 'react-native';
import Orientation from 'react-native-orientation-locker';

import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

import HomeScreen from './screens/HomeScreen';
import PatientScreen from './screens/PatientScreen';
// import PersonnelScreen from './screens/PersonnelScreen';

const Stack = createNativeStackNavigator();

export default function App() {
  useEffect(() => {
    Orientation.lockToLandscapeLeft(); // Landscape'a sabitler
    return () => {
      Orientation.unlockAllOrientations(); // Component unmount olduğunda serbest bırakır
    };
  }, []);

  return (
    <SafeAreaView style={{ flex: 1 }}>
      <NavigationContainer>
        <Stack.Navigator initialRouteName="Home" screenOptions={{ headerShown: false }}>
          <Stack.Screen name="Home" component={HomeScreen} />
          <Stack.Screen name="PatientScreen" component={PatientScreen} />
          {/* <Stack.Screen name="PersonnelScreen" component={PersonnelScreen} /> */}
        </Stack.Navigator>
      </NavigationContainer>
    </SafeAreaView>
  );
}
