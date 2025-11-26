import { StyleSheet } from "react-native";

export default StyleSheet.create({
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
  emergencyPerson: {
    backgroundColor: '#ffc2afff',
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
    shadowOffset: { width: 2, height: 4 }, // Gölge konumu
    shadowOpacity: 0.55, // Gölge saydamlığı
    shadowRadius: 2,
  },
  container: {
    marginVertical: 10,
    backgroundColor: '#fff',
  },
});
