import { StyleSheet } from 'react-native';

export default StyleSheet.create({
  textTitle: {
    color: '#000',
    fontSize: 18,
    fontWeight: '500',
    marginBottom: 10,
  },
  line: {
    width: '100%',
    height: 1,
    backgroundColor: '#d2d9dfff',
  },
  appointmentContainer: {
    minHeight: 200,
    backgroundColor: '#fff',
    borderRadius: 10,
    borderWidth: 0.5,
    borderColor: '#d2d9dfff',
    shadowColor: '#6d6d6dff',
    shadowOffset: { width: 0, height: 4 },
    shadowOpacity: 0.2,
    shadowRadius: 3,
    elevation: 1,
    padding: 16,
    paddingVertical: 24,
  },
});
