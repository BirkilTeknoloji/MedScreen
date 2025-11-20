import { View, Text, StyleSheet, Image } from 'react-native';
import styles from './styles/PatientProfileStyle';

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
  const userBd = formatDate(userData.BirthDate);
  const created = formatDate(userData.CreatedAt);
  return (
    <View style={cusStyle.profil}>
      <View style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
        <View style={{ flexDirection: 'row', gap: 24 }}>
          <View style={cusStyle.avatarInfo}>
            <View>
              <Image
                source={{
                  uri: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQpd4mJRIUwqgE8D_Z2znANEbtiz4GhI4M8NQ&s',
                }}
                style={cusStyle.profileImage}
              />
            </View>
          </View>
          <View>
            <View style={cusStyle.container}>
              <View>
                <Text style={cusStyle.textName}>{userData.Name}</Text>
                <Text>Hasta TC: {userData.TCNumber}</Text>
              </View>
            </View>
            <View style={cusStyle.row}>
              <View>
                <Text style={cusStyle.infoText}>📅 Doğum Tarihi</Text>
                <Text style={cusStyle.infoText2}>{userBd}</Text>
              </View>
              <View>
                <Text style={cusStyle.infoText}>Yaş</Text>
                <Text style={cusStyle.infoText2}>23</Text>
              </View>
              <View>
                <Text style={cusStyle.infoText}>Cinsiyet</Text>
                <Text style={cusStyle.infoText2}>{userData.Gender}</Text>
              </View>
              <View>
                <Text style={cusStyle.infoText}>Kan Grubu</Text>
                <Text style={cusStyle.infoText2}>{userData.BloodType}</Text>
              </View>
              <View>
                <Text style={cusStyle.infoText}>Boy(cm)</Text>
                <Text style={cusStyle.infoText2}>{userData.Height}</Text>
              </View>
              <View>
                <Text style={cusStyle.infoText}>Kilo(kg)</Text>
                <Text style={cusStyle.infoText2}>{userData.Weight}</Text>
              </View>
            </View>
          </View>
        </View>
        <View>
          <View>
            <Text
              style={[
                cusStyle.isActive,
                { color: '#166534', fontWeight: '500' },
              ]}
            >
              Active
            </Text>
            <Text style={{ textAlign: 'right' }}>Last Visit:{created}</Text>
            <Text style={{ textAlign: 'right' }}>
              Next Appointment:30.10.2025
            </Text>
          </View>
        </View>
      </View>
      <View style={cusStyle.line}></View>
      <View style={cusStyle.profilePerson}>
        <View>
          <Text style={cusStyle.infoText}>Telefon Numarası</Text>
          <Text style={cusStyle.infoText2}>{userData.Phone}</Text>
        </View>
        <View>
          <Text style={cusStyle.infoText}>Email Adresi</Text>
          <Text style={cusStyle.infoText2}>{userData.Phone}</Text>
        </View>
        <View>
          <Text style={cusStyle.infoText}>Adres</Text>
          <Text style={cusStyle.infoText2}>{userData.Address}</Text>
        </View>
      </View>
    </View>
  );
}

const cusStyle = StyleSheet.create({
  profilePerson: {
    flexDirection: 'row',
    justifyContent: 'space-evenly',
    alignItems: 'flex-start',
    padding: 8,
  },
  line: {
    width: '100%',
    height: 2,
    backgroundColor: '#d3d3d3ff',
    borderRadius: 12,
  },
  isActive: {
    backgroundColor: '#dcfce7',
    padding: 5,
    borderRadius: 12,
    alignSelf: 'flex-start',
  },
  infoText: {
    fontSize: 14,
    fontWeight: '400',
    color: '#000',
    opacity: 0.6,
  },
  infoText2: {
    fontSize: 18,
    fontWeight: '600',
    fontStyle: 'italic',
    color: '#000',
  },
  row: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    marginBottom: 8,
    gap: 24,
  },
  cell: {
    flex: 1,
    paddingHorizontal: 10,
  },
  label: {
    fontSize: 12,
    color: '#6c757d',
    marginBottom: 2,
  },
  value: {
    fontSize: 16,
    fontWeight: '600',
    color: '#212529',
  },
  textName: {
    fontWeight: 'bold',
    fontSize: 42,
  },
  avatarInfo: {
    flexDirection: 'row',
  },
  profileImage: {
    width: 100,
    height: 100,
    borderRadius: 50, // yuvarlak görünüm istiyorsan
    borderWidth: 2,
    borderColor: '#ddd',
  },
  profil: {
    marginBottom: 20,
    backgroundColor: '#f8f9fa',
    borderRadius: 10,
    padding: 15,
    borderWidth: 1,
    borderColor: '#e9ecef',
    shadowColor: '#0000',
    shadowOffset: { width:2, height: 4 }, // Gölge konumu
    shadowOpacity: 0.55, // Gölge saydamlığı
    shadowRadius: 2,
  },
  container: {
    marginVertical: 10,
    backgroundColor: '#fff',
  },
});
