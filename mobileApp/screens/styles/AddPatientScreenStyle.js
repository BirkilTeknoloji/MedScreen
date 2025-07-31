import { StyleSheet } from 'react-native';

export default StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#f0f4f7',
        padding: 20,
        justifyContent: 'center',
    },
    nfcImage: {
        width: 140,
        height: 140,
        alignSelf: 'center',
        marginBottom: 30,
    },
    infoText: {
        fontSize: 20,
        fontWeight: 'bold',
        textAlign: 'center',
        marginBottom: 20,
    },
    infoBox: {
        backgroundColor: '#fff',
        padding: 15,
        borderRadius: 10,
        marginBottom: 20,
        borderWidth: 1,
        borderColor: '#ccc',
    },
    label: {
        fontWeight: 'bold',
        fontSize: 16,
        marginTop: 10,
    },
    value: {
        fontSize: 16,
        color: '#333',
    },
    statusText: {
        fontSize: 14,
        marginTop: 15,
        textAlign: 'center',
        color: '#666',
    },
});