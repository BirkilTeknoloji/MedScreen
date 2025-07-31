import { View, Button } from 'react-native';
import { useNavigation } from '@react-navigation/native';
import styles from './styles/ActionButtonsStyle';

export default function ActionButtons() {
  const navigation = useNavigation();

  const handleNavigate = (screen) => navigation.navigate(screen);

  return (
    <View style={styles.buttonGroup}>
      <Button 
        title="📷 QR Kod Okut" 
        onPress={() => handleNavigate('QrScannerScreen')} 
      />
      <Button 
        title="↩️ Çıkış Yap" 
        onPress={() => handleNavigate('Home')} 
      />
    </View>
  );
}
