module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  devServer: {
    open: true,
    proxy: {
      '/download': {
        target: process.env.VUE_APP_SERVER,
        secure: false,
        changeOrigin: true,
      },
      '/file': {
        target: process.env.VUE_APP_SERVER,
        secure: false,
        changeOrigin: true,
      },
    }
  }
}