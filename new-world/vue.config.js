module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  devServer: {
    open: process.env.ENABLE_DEV_SERVER,
    proxy: {
      '/download': {
        target: process.env.DEV_SERVER,
        secure: false,
        changeOrigin: true,
      },
      '/file': {
        target: process.env.DEV_SERVER,
        secure: false,
        changeOrigin: true,
      },
    }
  }
}