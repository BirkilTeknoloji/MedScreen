import { StyleSheet } from 'react-native';

export default StyleSheet.create({
  tabBar: { 
    flexDirection: 'row', 
    marginBottom: 10 
  },
  tabButton: {
    flex: 1,
    paddingVertical: 10,
    backgroundColor: '#eee',
    marginHorizontal: 2,
    alignItems: 'center',
    borderRadius: 4,
  },
  tabActive: {
    backgroundColor: '#87cefa',
  },
  tabText: { 
    fontWeight: 'bold',
    fontSize: 18, 
  },
});