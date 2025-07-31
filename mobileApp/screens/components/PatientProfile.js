import { View, Text } from 'react-native';
import styles from './styles/PatientProfileStyle';

export default function PatientProfile({ userData }) {
  const formatDate = (dateString) => {
    if (!dateString) return "N/A";
    const date = new Date(dateString);
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const year = date.getFullYear();
    return `${day}.${month}.${year}`;
  };

  return (
    <View style={styles.profil}>
      <Text style={styles.isim}>{userData.Name}</Text>
      
      <View style={styles.infoContainer}>
        <View style={styles.infoGrid}>
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>🆔 TC No:</Text>
            <Text style={styles.infoValue}>{userData.TCNumber}</Text>
          </View>
          
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>📅 Doğum:</Text>
            <Text style={styles.infoValue}>{formatDate(userData.BirthDate)}</Text>
          </View>
          
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>👤 Cinsiyet:</Text>
            <Text style={styles.infoValue}>{userData.Gender}</Text>
          </View>
          
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>🩸 Kan Grubu:</Text>
            <Text style={styles.infoValue}>{userData.BloodType}</Text>
          </View>
          
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>📏 Boy/Kilo:</Text>
            <Text style={styles.infoValue}>{userData.Height}cm / {userData.Weight}kg</Text>
          </View>
          
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>📞 Telefon:</Text>
            <Text style={styles.infoValue}>{userData.Phone}</Text>
          </View>
        </View>
      
        <View style={styles.addressRow}>
          <Text style={styles.infoLabel}>📍 Adres:</Text>
          <Text style={styles.infoValue}>{userData.Address}</Text>
        </View>
      </View>
    </View>
  );
}
