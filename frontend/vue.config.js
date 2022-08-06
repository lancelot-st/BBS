module.exports = {
    assetsDir: "static",
    devServer: {
        proxy: {
            '/api': {
              target: 'localhost:8080',
              changeOrigin: true,
            }
        }
    }
  }
