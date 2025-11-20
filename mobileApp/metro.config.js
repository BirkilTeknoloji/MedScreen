const { getDefaultConfig, mergeConfig } = require('@react-native/metro-config');

const config = {
  resolver: {
    assetExts: ['bin', 'txt', 'jpg', 'png', 'json', 'mp4', 'mov'],
  },
};

module.exports = mergeConfig(getDefaultConfig(__dirname), config);