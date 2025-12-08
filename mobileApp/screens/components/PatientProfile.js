import { View, Text, StyleSheet, Image } from 'react-native';
import styles from '../styles/PatientProfileStyle';

export default function PatientProfile({ userData }) {
  const formatDate = dateString => {
    if (!dateString) return 'N/A';
    const date = new Date(dateString);
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const year = date.getFullYear();
    console.log(year);

    return `${day}.${month}.${year}`;
  };
  const userBd = formatDate(userData.birth_date);
  const created = formatDate(userData.CreatedAt);
  return (
    <View style={styles.profil}>
      <View style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
        <View style={{ flexDirection: 'row', gap: 24 }}>
          <View style={styles.avatarInfo}>
            <View>
              <Image
                source={{
                  uri: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQpd4mJRIUwqgE8D_Z2znANEbtiz4GhI4M8NQ&s',
                }}
                style={styles.profileImage}
              />
            </View>
          </View>
          <View>
            <View style={styles.container}>
              <View>
                <Text style={styles.textName}>
                  {userData.first_name} {userData.last_name}
                </Text>
                <Text>Hasta TC: {userData.tc_number}</Text>
              </View>
            </View>
            <View style={styles.row}>
              <View>
                <Text style={styles.infoText}>📅 Doğum Tarihi</Text>
                <Text style={styles.infoText2}>{userBd}</Text>
              </View>
              <View>
                <Text style={styles.infoText}>Yaş</Text>
                <Text style={styles.infoText2}>23</Text>
              </View>
              <View>
                <Text style={styles.infoText}>Cinsiyet</Text>
                <Text style={styles.infoText2}>{userData.gender}</Text>
              </View>
              <View>
                <Text style={styles.infoText}>Kan Grubu</Text>
                <Text style={styles.infoText2}>{userData.blood_type}</Text>
              </View>
              <View>
                <Text style={styles.infoText}>Boy(cm)</Text>
                <Text style={styles.infoText2}>{userData.height}</Text>
              </View>
              <View>
                <Text style={styles.infoText}>Kilo(kg)</Text>
                <Text style={styles.infoText2}>{userData.weight}</Text>
              </View>
            </View>
          </View>
        </View>
        <View style={styles.row}>
          {/* <UserDataCom
            title="Doktor İletişim Bilgileri"
            name={userData.primary_doctor.first_name + ' ' + userData.primary_doctor.last_name}
            phone={userData.primary_doctor.phone} 
            color={'#1b8b05ff'}
            bgColor={"#b0ffa0ff"}
          />
          <UserDataCom 
            title="Acil Durumda İletişime Geçilecek Kişi" 
            name={userData.emergency_contact_name} 
            phone={userData.emergency_contact_phone} 
            color={'#dd612fff'}
            bgColor={"#ffb3b0ff"}
            />
          */}
        </View>
      </View>
      <View style={styles.line}></View>
      <View style={styles.profilePerson}>
        <View>
          <Text style={styles.infoText}>Telefon Numarası</Text>
          <Text style={styles.infoText2}>{userData.phone}</Text>
        </View>
        <View>
          <Text style={styles.infoText}>Email Adresi</Text>
          <Text style={styles.infoText2}>{userData.email}</Text>
        </View>
        <View>
          <Text style={styles.infoText}>Adres</Text>
          <Text style={styles.infoText2}>{userData.address}</Text>
        </View>
      </View>
    </View>
  );
}
