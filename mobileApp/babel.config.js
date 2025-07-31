module.exports = {
  presets: [
    '@react-native/babel-preset', // Yeni RN için önerilen preset
  ],
  plugins: [
    'react-native-reanimated/plugin', // Vision Camera için gerekli
    ['module:react-native-dotenv', {
      moduleName: '@env',
      path: '.env',
      blocklist: null,
      allowlist: null,
      safe: false,
      allowUndefined: true,
    }]
  ],
};
