module.exports = {
  devServer: {
    proxy: {
      '^/login': {
        target: 'https://auth.jacob.run/',
        ws: true,
      },
      '^/booking': {
        target: 'https://auth.jacob.run/proxy/roomookBackend/q/',
        ws: true,
      },
      '^/room': {
        target: 'https://auth.jacob.run/proxy/roomookBackend/q/',
        ws: true,
      },
      '^/b/r': {
        target: 'https://auth.jacob.run/proxy/roomookBackend/q/',
        ws: true,
      },
      '^/r/d': {
        target: 'https://auth.jacob.run/proxy/roomookBackend/q/',
        ws: true,
      },
      '^/user/loggedinuser': {
        target: 'https://auth.jacob.run/proxy/roomookBackend/q/',
        ws: true,
      },

    }
  }
}