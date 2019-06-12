module.exports = {
  devServer: {
    proxy: {
      '^/login': {
        target: 'https://auth.jacob.run/',
        ws: true,
      },
      '^/booking': {
        target: 'http://localhost:3000/',
        ws: true,
      },
      '^/room': {
        target: 'http://localhost:3000/',
        ws: true,
      },
      '^/b/r': {
        target: 'http://localhost:3000/',
        ws: true,
      },
      '^/r/d': {
        target: 'http://localhost:3000/',
        ws: true,
      },
      '^/user/loggedinuser': {
        target: 'http://localhost:3000/',
        ws: true,
      },
      //Local Routes for Tablet without auth
      '^/tablet/booking': {
        target: 'http://localhost:3000/',
        pathRewrite: {'^/tablet/booking/': '/booking/'},
        changeOrigin: true,
        ws: true,
      },
      '^/tablet/b/r': {
        target: 'http://localhost:3000/',
        pathRewrite: {'^/tablet/b/r/': '/b/r/'},
        changeOrigin: true,
        ws: true,
      },

    }
  }
}